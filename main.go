package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	saleDatesTemplate       = "https://web.northamptoncounty.org/SheriffSale/api/Dates/GetThisYearDates/$date?_=$timestamp"
	saleDateDetailsTemplate = "https://web.northamptoncounty.org/SheriffSale/api/Listings/GetSelListing/$date?_=$timestamp"
	parcelUrlTemplate       = "https://www.ncpub.org/_web/datalets/datalet.aspx?mode=profile_nh&UseSearch=NO&pin=$parcelNumber"
	dispositionStayed       = "STAYED"
	dispositionContinuedTo  = "CONTINUED TO"
)

type sheriffSalesDatesResponse struct {
	ID              int    `json:"ID"`
	SaleDate        string `json:"SaleDate"`
	FileDate        string `json:"FileDate"`
	AdvertisingDate string `json:"AdvertisingDate"`
}

type sheriffSaleDetailResponse struct {
	ID              int         `json:"ID"`
	DocketNumber    string      `json:"DocketNumber"`
	CaseTitle       string      `json:"CaseTitle"`
	AttorneyName    string      `json:"AttorneyName"`
	SaleNumber      int         `json:"SaleNumber"`
	ParcelMap       string      `json:"ParcelMap"`
	ParcelBlock     string      `json:"ParcelBlock"`
	ParcelLot       string      `json:"ParcelLot"`
	ParcelSchool    string      `json:"ParcelSchool"`
	ParcelMunicipal string      `json:"ParcelMunicipal"`
	ParcelCounty    string      `json:"ParcelCounty"`
	ParcelNumber    string      `json:"ParcelNumber"`
	SaleAddress     string      `json:"SaleAddress"`
	Town            string      `json:"Town"`
	Disposition     string      `json:"Disposition"`
	DisplayMonth    string      `json:"DisplayMonth"`
	TextMonth       string      `json:"TextMonth"`
	DebtAmount      float64     `json:"DebtAmount"`
	AllCosts        float64     `json:"AllCosts"`
	UserEntered     string      `json:"UserEntered"`
	EnterDate       string      `json:"EnterDate"`
	CaseID          interface{} `json:"CaseID"`
	Month           int         `json:"Month"`
	Year            int         `json:"Year"`
	AttorneyRep     interface{} `json:"AttorneyRep"`
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Info("using debug log level")
	}
	currentDate := time.Now().Format("01-02-2006")
	saleDatesUrl := strings.Replace(saleDatesTemplate, "$date", currentDate, 1)
	saleDatesUrl = strings.Replace(saleDatesUrl, "$timestamp", strconv.FormatInt(time.Now().UnixMilli(), 10), 1)
	logrus.WithField("currentDate", currentDate).WithField("saleDatesUrl", saleDatesUrl).Debug("built sale dates url")
	resp, err := http.Get(saleDatesUrl)
	if err != nil {
		logrus.WithError(err).Error("failed to fetch upcoming sheriff sales")
		panic(err)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("failed to read response body")
	}
	logrus.WithField("upcoming response", string(bytes)).Debug("got upcoming sales response")
	sheriffSaleDatesResponses := &[]sheriffSalesDatesResponse{}
	err = json.Unmarshal(bytes, sheriffSaleDatesResponses)
	if err != nil {
		logrus.WithError(err).Error("failed to parse sale dates response")
	}
	parcelDetails := make([]parcelDetail, 0)
	for _, sheriffSaleDateResponse := range *sheriffSaleDatesResponses {
		saleDate, err := time.Parse("2006-01-02T00:00:00", sheriffSaleDateResponse.SaleDate)
		if err != nil {
			logrus.WithError(err).WithField("sale id", sheriffSaleDateResponse.ID).Debug("failed to parse sale date")
			continue
		}
		detailsUrl := strings.Replace(saleDateDetailsTemplate, "$date", saleDate.Format("01-02-2006"), 1)
		detailsUrl = strings.Replace(detailsUrl, "$timestamp", strconv.FormatInt(time.Now().UnixMilli(), 10), 1)

		logrus.WithField("detailsUrl", detailsUrl).Debug("generated details url")
		resp, err = http.Get(detailsUrl)
		if err != nil {
			logrus.WithError(err).WithField("sale id", sheriffSaleDateResponse.ID).Info("failed to fetch details for sale")
			continue
		}
		bytes, err = io.ReadAll(resp.Body)
		logrus.WithField("sale details response", string(bytes)).Debug("got sale details response")
		details := make([]sheriffSaleDetailResponse, 0)
		err = json.Unmarshal(bytes, &details)
		if err != nil {
			logrus.WithError(err).WithField("sale id", sheriffSaleDateResponse.ID).Info("failed to fetch details for sale")
			continue
		}
		filterSalesDetails(details)
		parcelDetails = append(parcelDetails, fetchParcelDetails(details)...)
	}
	j, err := json.MarshalIndent(parcelDetails, "", "    ")
	if err != nil {
		logrus.WithError(err).Error("failed to marshal parcel details")
		panic(err)
	}
	println(fmt.Sprintf("Found %v parcels greater than 10 acres:", len(parcelDetails)))
	println(string(j))
}

type parcelDetail struct {
	ParcelNumber string
	ParcelUrl    string
	Acres        float64
}

func fetchParcelDetails(details []sheriffSaleDetailResponse) []parcelDetail {
	result := make([]parcelDetail, 0)
	for _, detail := range details {
		url := strings.Replace(parcelUrlTemplate, "$parcelNumber", detail.ParcelNumber, 1)
		logrus.WithField("parcel url", url).Debug("generated parcel url")
		resp, err := http.Get(url)
		if err != nil {
			logrus.WithError(err).WithField("parcel number", detail.ParcelNumber).Error("failed to fetch parcel details")
			continue
		}
		acreage, err := getAcreageFromParcelHtml(resp.Body)
		if err != nil {
			logrus.WithError(err).WithField("parcel number", detail.ParcelNumber).Error("failed get acreage from parcel html")
			continue
		}
		logrus.WithField("acreage", acreage).Debug("found acreage for parcel")
		if acreage > 10 {
			result = append(result, parcelDetail{
				ParcelNumber: detail.ParcelNumber,
				ParcelUrl:    url,
				Acres:        acreage,
			})
		}
	}
	return result
}

func getAcreageFromParcelHtml(doc io.Reader) (float64, error) {
	d, err := html.Parse(doc)
	if err != nil {
		return 0, err
	}
	var f func(node *html.Node) float64
	f = func(node *html.Node) float64 {
		if node.Type == html.TextNode && node.Data == "CAMA Acres" {
			logrus.WithField("label", node.Data).Debug("found node with correct field label")
			value := node.Parent.NextSibling.FirstChild.Data
			float, err := strconv.ParseFloat(value, 64)
			if err != nil {
				logrus.WithError(err).Debug("failed to find acreage value in label's next sibling")
				return 0
			}
			return float
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			acreage := f(c)
			if acreage > 0 {
				return acreage
			}
		}
		return 0
	}
	return f(d), nil
}

func filterSalesDetails(details []sheriffSaleDetailResponse) []sheriffSaleDetailResponse {
	count := 0
	for i := len(details) - 1; i > -1; i-- {
		detail := details[i]
		if detail.Disposition == dispositionStayed || detail.Disposition == dispositionContinuedTo {
			details = append(details[:i], details[i+1:]...)
			count++
		}
	}
	logrus.WithField("number of removed details", count).Debug("finished filtering sales details")
	return details
}
