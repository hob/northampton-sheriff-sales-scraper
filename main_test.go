package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHtmlParsing(t *testing.T) {
	acreage, err := getAcreageFromParcelHtml(strings.NewReader(doc))
	assert.NoError(t, err)
	assert.Equal(t, 0.055, acreage)
}

var doc = `<!DOCTYPE HTML>
<html lang='en-US' XMLNS:MPC><head runat='server'>
	                            <title>Northampton County</title>
                                <meta http-equiv='Content-Type' content='text/html; charset=UTF-8'/>
                                <meta http-equiv='X-UA-Compatible' content='IE=9' />
	                            <meta http-equiv='Keywords' content='CAMA, Assessment Administration, Tax'>
	                            <meta http-equiv='Description' content='iasWorld'>
	                            <meta name='Copyright' content='Tyler Technologies Inc, Akanda Solutions, LLC'> 
                                <link rel='icon' href='https://www.ncpub.org/_webfavicon.ico' type='image/x-icon'>
                                <link rel='shortcut icon' href='https://www.ncpub.org/_webfavicon.ico' type='image/x-icon'>
                                <meta http-equiv='pragma' content='no-cache'>
                            </head><link rel='stylesheet' href='https://www.ncpub.org/_web/styles/font-awesome.min.css' type='text/css'><link rel='stylesheet' href='https://www.ncpub.org/_web/styles/styles.css' type='text/css'><link rel='stylesheet' href='https://www.ncpub.org/_web/styles/custom/override-NorthamptonPA.css' type='text/css'><body id='thebody' class='normalpage' ><div id='pagewrapper'><div id='divSMaint' style='display:none;border-bottom:2px solid red;width:780'><table role='presentation' cellpadding='0' cellspacing='0' width=780><tr id='trSMaint' class='MessageBar' style='height:21px;vertical-align:middle;'><td id='tdSMaint'>&nbsp;The System is currently unavailable due to maintenance.  Please check again later.</td><td id='tdTimer' align='right'></td><td align="right" valign="middle" onmouseout="ChangeImage('out')" onmouseover="ChangeImage('over')"><img id="CloseMsgBar" src="../images/icons/Set1/closeicon-normal.gif" height="14px" alt="Hide Bar" style="vertical-align:middle; border:solid 1px white; cursor:pointer" onclick="HideBar()"  />&nbsp;</td></tr></table></div><script type='text/javascript' language='javascript'>function HideBar() { document.getElementById('divSMaint').style.display='none';SoapInvokeService(WS_ACCESS_MANAGER, 'HideBar', null, false)}function ChangeImage(el) { var img = document.getElementById('CloseMsgBar'); if (el == 'out') img.src = img.src.replace('closeicon-mouseover.gif', 'closeicon-normal.gif');  else img.src = img.src.replace('closeicon-normal.gif', 'closeicon-mouseover.gif'); }</script><input type='hidden' id='hdDefHelp' value='1'><div id='outerheader'><div id='middleheader'>

<header>
    <div id='header'>
        <div class='sitelogocontainer'>
            <a href="http://www.ncpub.org/_web/forms/htmlframe.aspx?mode=content/home.htm" id="aIASWLink" target="_blank"><div class="sitelogo"></div></a>
        </div>
        <div class="HeaderSiteTitle"></div>
        <div id='rightsection' class='HeaderLinks'>
            <span id="RoleValue" class="HeaderLinks">
                    </span><span id="lbRole"></span>
            <span id="ContSignIn"><a href="../Main/Login.aspx?targetUrl=" id="LinkSignIn" onclick="return&#32;HeaderClick(this)" class="HeaderLinks">SIGN IN</a> &nbsp;&nbsp;|&nbsp;&nbsp; </span>
            <span id="contactUs"><A id="linkContactUs" class="HeaderLinks" href="../Forms/ContactUs.aspx" target="_self">CONTACT US</A>&nbsp;&nbsp;|&nbsp;&nbsp; </span>
            <span id="lbHelp"><a class="Headerlinks Help" href="javascript:showHelp('')"
                class='HeaderLinks'>
                <img src='../images/common/icons/050.png' alt='Help' /></a> &nbsp;&nbsp; </span>
            
        </div>
    </div>
</header>

<nav>
       <div id='topmenu'><ul><li id='pd_Parent' class='unsel' onmouseover='menuOver_pd("pd_189", event, this)' onmouseout='menuOut("pd_189", event, this)' onfocus='menuOver_pd("pd_189", event, this)' onfocusout='menuOut("pd_189", event, this)'><a href='../forms/htmlframe.aspx?mode=content/home.htm' target=''><span>Home<span style='font-size: 7'></span></span></a></li><li id='pd_Parent' class='sel' onmouseover='menuOver_pd("pd_2", event, this)' onmouseout='menuOut("pd_2", event, this)' onfocus='menuOver_pd("pd_2", event, this)' onfocusout='menuOut("pd_2", event, this)'><a href='../search/commonsearch.aspx?mode=address' target=''><span>Property Records<span style='font-size: 7'></span></span></a></li><li class='unsel'><a href='../forms/externallink.aspx?mode=https://www.northamptoncounty.org/pages/default.aspx' target=''><span>County Website</span></a></li></ul></div>
       <div id='secondarytopmenu'><ul></ul></div>
   </nav>
<script language='Javascript' src='../Script/Menu.js'></script>
<div id='pd_189' class='PullDownMenuLayer' style='visibility:hidden; position: absolute; width:50; top:0; left:0' onmouseover='bMenuOn=true' onmouseout='menuOut(this.id)' onfocus='bMenuOn=true' onfocusout='menuOut(this.id)'>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../forms/htmlframe.aspx?mode=content/history.htm"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;County History&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../forms/htmlframe.aspx?mode=content/millagerates.htm"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Millage Rates&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='javascript:window.open("../content/school_districts.pdf");' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;School Districts&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../content/municipality_list.pdf"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Municipalities&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../forms/externallink.aspx?mode=https://www.northamptoncounty.org/fisaff/assessmnt/pages/tax-collectors.aspx"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Tax Collectors&nbsp;</a></div></div>
<div id='pd_2' class='PullDownMenuLayer' style='visibility:hidden; position: absolute; width:50; top:0; left:0' onmouseover='bMenuOn=true' onmouseout='menuOut(this.id)' onfocus='bMenuOn=true' onfocusout='menuOut(this.id)'>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../search/commonsearch.aspx?mode=address"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Address&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../search/commonsearch.aspx?mode=parid"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Parcel&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../search/advancedsearch.aspx?mode=sales"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Sales&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../search/advancedsearch.aspx?mode=advanced"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Advanced&nbsp;</a></div>
<div class='PullDownMenuItem'><a class='PullDownMenuItem' onclick='location="../maps/mapadv.aspx"' onmouseover='this.className="MenuRowHL"' 
onmouseout='menuLinkOut(this); this.className="PullDownMenuItem"' onfocus='this.className="MenuRowHL"' 
onfocusout='menuLinkOut(this); this.className="PullDownMenuItem"' 
onmousedown='this.className="MenuRowClicked"'>&nbsp;&nbsp;Map Search&nbsp;</a></div></div>
<script type="text/javascript" src="../script/soap.js?210917101150"></script>
<script type="text/javascript" src="../script/Common.js?210917101148"></script>
<script type="text/javascript">

    function HeaderClick(link) {
        var par = window.parent;
        if (par == window) return true;
        var op = par.opener;
        if (op == null) return true;
        if (op.HeaderClick == null) return true;
        var id = link.id;
        var el = op.document.getElementById(id);
        if (el == null) return true;
        if (!confirm("This command will close the current window.\nDo you want to proceed?")) return false;

        if (document.createEvent) {
            var evt = document.createEvent('MouseEvents');
            evt.initEvent('click', true, false);
            el.dispatchEvent(evt);
        } else if (typeof el.onclick == 'function') {
            el.click();
        }

        par.close();
        op.focus();
        return false;
    }

    function SetMyHome(url) {
        var DLG_SETMYHOME = "center:yes; dialogHeight:180px; dialogWidth:240px; help:no; resizable:no; status:no; title:no";
        window.showModalDialog('../Forms/SetMyHome.aspx?url=' + url, null, DLG_SETMYHOME);
    }


    function ShowSimpleHelp(url) {
        var DLG_POPUP = "scrollbar= yes, height=600, Width=800, resizable=yes, status=no, titlebar=no";
        if (url.indexOf("?") > 0) {
            url += "&";
        } else {
            url += "?";
        }
        window.open(url, null, DLG_POPUP);
    }

    function ViewCompleted() {
        var url = "../Maintain/ReportSpool.aspx?simple=1&runningJobs=0&batchJobs=2";
        window.open(url, "Reports", "status=1, toolbar=0, height=650px, width=1000px, resizable=1, scrollbars=yes");
    }

    function PopupWindow2(url) {
        var DLG_POPUP = "dialogLeft:0; dialogTop: 0; edge: sunken; scroll: yes; center:yes; dialogHeight:750; dialogWidth:1028; help:no; resizable:yes; status:no; title:no";
        if (url.indexOf("?") > 0) {
            url += "&";
        } else {
            url += "?";
        }
        url += "frameset=1";
        window.showModalDialog(url, null, DLG_POPUP);
    }

    var gLastPopup = parseInt(Math.random() * 10);

    function PopupWindow(url, noToolbar, name) {
        var win = null;

        if (name == null)
            name = "";
        else
            name = name.replace(' ', "");

        if (win == null) {
            if (name == null || name.length == 0) name = "popup" + (gLastPopup++ % 10);
            var params;
            if (noToolbar) {
                params = "fullscreen=no, titlebar=no, location=no, menubar=no, toolbar=no, resizable=yes, scrollbars=yes, status=yes, titlebar=yes";
            } else {
                params = null;
            }
            win = window.open(url, name, params);
        }
        return win;
    }

    function setRole(val, form) {
        var url = window.location.href;
        url = url.replace("#", "");
        var i = url.indexOf("&");
        if (i != -1) {
            url = url.substr(0, i);
        }
        var pars = SoapBuildInputParam("SelectedRole", val);
        pars += SoapBuildInputParam("URL", url);
        var res = SoapInvokeService(WS_PUBLIC, "AssignSelectedRole", pars);
        if (res == -1) {
            location.href = "../Main/Home.aspx";
        }
        else {
            location.reload();
        }
    }
</script>
</div></div><div id='wrapper' class='center'><div id='sidemenu' valign='top' class='columns' style='width:200; '><nav><div class='contentpanel'><div id='sidemenu'><ul class='navigation'><li class='sel'><a href='../datalets/datalet.aspx?mode=profile_nh&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Parcel</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=owner_details&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Owner</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=ownmlt_details&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Multi-Owner</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=residential_nh&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Residential</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=commercial_nh&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Commercial</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=oby_nh&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Out Buildings</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=land&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Land</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=values_nh&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Values</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=homestead&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Homestead</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=sales_all&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Sales</span></a></li><li class='unsel'><a href='../datalets/datalet.aspx?mode=charge_text&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Tax Information</span></a></li><li class='unsel'><a href='../idoc2/photoview.aspx?UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Photos</span></a></li><li class='unsel'><a href='../datalets/sketchframe.aspx?mode=sketch.aspx&amp;UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Sketch</span></a></li><li class='unsel'><a href='../maps/map.aspx?UseSearch=no&amp;pin=L9NE4C 20 2       0310&amp;jur=048&amp;taxyr=2023&amp;LMparent=20' target=''><span>Map</span></a></li></ul></div></div></nav></div><section><div id='bodyarea' >

<script>
var bOnLoad = false;
</script>
 <script src="//code.jquery.com/jquery-3.5.1.js"></script>
 <script src="//code.jquery.com/jquery-migrate-3.3.1.js"></script>
<style>
    .holder { /*width:400px;*/ position:relative; }
    .columns { float: none !important; }
</style>

<form name="frmMain" method="POST" action="./datalet.aspx?mode=profile_nh&amp;UseSearch=NO&amp;pin=L9NE4C+20+2+++++++0310" id="frmMain" style="position:relative;&#32;top:-5px;">
<div>
<input type="hidden" name="__EVENTTARGET" id="__EVENTTARGET" value="" />
<input type="hidden" name="__EVENTARGUMENT" id="__EVENTARGUMENT" value="" />
<input type="hidden" name="__LASTFOCUS" id="__LASTFOCUS" value="" />
<input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="" />
</div>

<script type="text/javascript">
//<![CDATA[
var theForm = document.forms['frmMain'];
if (!theForm) {
    theForm = document.frmMain;
}
function __doPostBack(eventTarget, eventArgument) {
    if (!theForm.onsubmit || (theForm.onsubmit() != false)) {
        theForm.__EVENTTARGET.value = eventTarget;
        theForm.__EVENTARGUMENT.value = eventArgument;
        theForm.submit();
    }
}
//]]>
</script>


<div>

	<input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="nEtVTqNvP5zrTEzthbUbaOn3XC/Tle7KsvAAlWFRVSjUQvT+vRNIKii4BfwZABfwWN+A006PlNR27abnKWY8gAn4smYJjS3iNImzbJE5Hdrcni3ySfBI3Ilj6ht7HJ0b4rekdbTGE7v8vZ71bwcYVWnFBcg467Tpm4Q744b58ooMxhOkccRRUDKkluUgEdxpMX3Z9vjjJY6VkmAd6b1EShRBz9xSNSZvD7Oe9+6di6tQjboUF2FSZ5ac8ejX44L4q7rioWqHU9XBe4qc2NJ8e7gwRtpYbgj3eNVRLYt3vr1XngUiuwQh5sT70tpym6e6EgE2zVZXShskumoa5hto9TnoYYgvte0ixURhbyf0gdxRR/DjS/uFAPMqYI6YUYIm/njjWDmHeBAVLQTaXUki4lNrshflaQAi5uCd+knpF3N7fJAFRF4eZbp8Kk3mx3mGAeu3VrGId6jdcYwvV+KCxqc00y6bBuK4zGj3N8bMFfRuOG9HSK1Hd71+zzSllJlW" />
</div>

  <div style='position:relative; top: 5px;' class='columns'>
  <div class='contentpanel'><div style="min-width:500px; min-height:450px;max-width:1600px;   vertical-align:top;">
	 <!--Container for datalet-->  
  <table style="vertical-align:top;" width=600px cellpadding="2" cellspacing="2">
		<tr>
			<td valign="top" width="80%">
				<table border="0" cellpadding="0" cellspacing="0" width="100%" class="WidgetBar" align="center">
					<tr id="datalet_header_row">
						<td>
							<?xml version="1.0" encoding="utf-16"?><table width="100%" border="0" cellspacing="0" cellpadding="0"><tr class="DataletHeaderTop" height="1"><td class="DataletHeaderTop" width="50%" align="left" colspan="2" nowrap="1">PARID: L9NE4C 20 2       0310</td><td class="DataletHeaderTop" width="50%" align="right" colspan="2"></td></tr><tr height="1" class="BannerTabs"><td colspan="4" height="1" valign="top" width="100%"></td></tr><tr class="DataletHeaderBottom" height="1"><td class="DataletHeaderBottom" width="50%" align="left" colspan="2">EVANS SCHWANDA R,  </td><td class="DataletHeaderBottom" width="50%" align="right" colspan="2">829 SPRING GARDEN ST  </td></tr><tr><td colspan="4" height="10" valign="top" width="100%"></td></tr></table>

						</td>
					</tr> 
					<tr>
						<td>
                            <div class="holder">
						    	<div id='datalet_div_1' name='PARCEL_NH'><table align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td colspan=2 class="DataletTitleColor"><font size=-1>Parcel</font></td></tr></table><table id='Parcel' align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Property Location</td><td width='*' valign='top' align='left' class="DataletData">829       SPRING GARDEN ST</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Unit Desc</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Unit #</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">City</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">State</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Zip Code</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Neighborhood Valuation Code</td><td width='*' valign='top' align='left' class="DataletData">1004</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Trailer Description</td><td width='*' valign='top' align='left' class="DataletData"><BR><BR></td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Municipality</td><td width='*' valign='top' align='left' class="DataletData">EASTON CITY</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Classification</td><td width='*' valign='top' align='left' class="DataletData">Residential</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Land Use Code</td><td width='*' valign='top' align='left' class="DataletData">110 - Single Family, Residential</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">School District</td><td width='*' valign='top' align='left' class="DataletData">EASTON SCHOOL DIST</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Topography</td><td width='*' valign='top' align='left' class="DataletData">LEVEL</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Utilities</td><td width='*' valign='top' align='left' class="DataletData">ALL PUBLIC</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Street/Road</td><td width='*' valign='top' align='left' class="DataletData">PAVED/SIDEWALK</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Total Cards</td><td width='*' valign='top' align='left' class="DataletData">1</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Living Units</td><td width='*' valign='top' align='left' class="DataletData">1</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">CAMA Acres</td><td width='*' valign='top' align='left' class="DataletData">.055</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Homestead /Farmstead</td><td width='*' valign='top' align='left' class="DataletData">H - Homestead</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Approved?</td><td width='*' valign='top' align='left' class="DataletData">A - Approved</td></tr><tr><td colspan=2 height="1" valign="top" width='100%'></td></tr><hr></table><font size='0px'><br></font></div><div id='datalet_div_2' name='OWNER'><table align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td colspan=2 class="DataletTitleColor"><font size=-1>Parcel Mailing Address</font></td></tr></table><table id='Parcel Mailing Address' align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">In Care of</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Name(s)</td><td width='*' valign='top' align='left' class="DataletData">EVANS SCHWANDA R</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Mailing Address</td><td width='*' valign='top' align='left' class="DataletData">829 SPRING GARDEN ST</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">City, State, Zip Code</td><td width='*' valign='top' align='left' class="DataletData">EASTON, PA, 18042-3374</td></tr><tr><td colspan=2 height="1" valign="top" width='100%'></td></tr><hr></table><font size='0px'><br></font></div><div id='datalet_div_3' name='PARCEL_ALT'><table align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td colspan=2 class="DataletTitleColor"><font size=-1>Alternate Address</font></td></tr></table><table id='Alternate Address' align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Alternate Address</td><td width='*' valign='top' align='center' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">City</td><td width='*' valign='top' align='center' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">State</td><td width='*' valign='top' align='center' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Zip</td><td width='*' valign='top' align='center' class="DataletData">&nbsp;</td></tr><tr><td colspan=2 height="1" valign="top" width='100%'></td></tr><hr></table><font size='0px'><br></font></div><div id='datalet_div_4' name='ASMT_FLAGS'><table align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td colspan=2 class="DataletTitleColor"><font size=-1>ACT Flags</font></td></tr></table><table id='ACT Flags' align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Act 319/515</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">LERTA</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Act 43</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Act 66</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Act 4/149</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">KOZ</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">TIF Expiration Date</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">BID</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Millage Freeze Date</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Millage Freeze Rate</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">Veterans Exemption</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td colspan=2 height="1" valign="top" width='100%'></td></tr><hr></table><font size='0px'><br></font></div><div id='datalet_div_5' name='TAX_COLLECTOR'><table align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td colspan=2 class="DataletTitleColor"><font size=-1>Tax Collector</font></td></tr></table><table id='Tax Collector' align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">MARK LYSYNECKY, FINANCE DIRECTOR</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">123 S 3RD ST</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">EASTON PA 18042</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">&nbsp;</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">610-250-6625</td></tr><tr><td colspan=2 height="1" valign="top" width='100%'></td></tr><hr></table><font size='0px'><br></font></div><div id='datalet_div_6' name='ASSESSOR'><table align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td colspan=2 class="DataletTitleColor"><font size=-1>Assessor</font></td></tr></table><table id='Assessor' align=center width='100%' border='0' borderColor=black cellPadding=1 cellSpacing=0><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">JULIE AZZALINA</td></tr><tr><td width='40%' valign='top' align='left' class="DataletSideHeading">&nbsp;</td><td width='*' valign='top' align='left' class="DataletData">610-829-6167</td></tr><tr><td colspan=2 height="1" valign="top" width='100%'></td></tr><hr></table><font size='0px'><br></font></div>


                            </div>
						</td>
					</tr>					
					<tr>
						<td align="center">
							<div class="holder">
                                
                            </div>
						</td>
					</tr>
					<tr>
						<td align="center">
                            <div class="holder">
							    
                            </div>
						</td>
					</tr>
					<tr>
						<td align="center">
                            <div class="holder">
							    
                            </div>
						</td>
					</tr>
					
				</table>
			</td>
		
		</tr>
	</table>
	
</div></div>
    </div>

  <!--Column for Side Widget-->
  
        <div style='position:relative; top: 5px;' class='columns'>
				<div class='contentpanel'><div style="align: left; vertical-align:top; width:160px; overflow:hidden;">
	 <!--Container Right hand panel-->  
                    <table border="0" cellpadding="0" cellspacing="0" class="WidgetBar" align="center">
					<tr id="trTaxYear">
		<td class="DataletHeader" align="right">
							
<!-- record navigator: start -->


<script type="text/javascript">
    (function (window) {
        if (!Tyler) { var Tyler = {}; }
        Tyler.SyncWindow = function (paramCallback) {
            var m_isEnabled = ("False" == "True");
            var m_winRef = null;
            var m_paramCallback = paramCallback;
            var m_winName = "SyncWindow_" + "wwwncpuborg_web";
            var m_url = "../Forms/SyncWindow.aspx";
            var m_lastQueryString = "";
            var me = this;

            this.IsEnabled = function () {
                /// <summary>returns true if SyncWindow feature is enabled, false otherwise</summary>
                /// <returns>boolean</returns>
                return m_isEnabled;
            }

            this.SetCallback = function (callback) {
                /// <summary>Sets the parameter callback function to be called by the SyncWindow.Open method</summary>
                /// <returns>void</returns>
                m_paramCallback = callback;
            }

            this.Open = function () {
                /// <summary>(re)loads the SyncWindow window by passing a URL to it containing qeurystring parameters
                /// built using the parameter callback function</summary>
                /// <returns>void</returns>
                if (m_isEnabled) {
                    var parms = null;
                    if (m_paramCallback) parms = m_paramCallback();
                    if (parms != null) {
                        var qs = "";
                        for (var key in parms) {
                            if (qs.length > 0) qs += "&";
                            qs += key + "=" + parms[key];
                        }
                        if ("#" + qs != m_lastQueryString) {
                            m_winRef = window.open(m_url + "#" + qs, m_winName);
                        }
                    }
                }
            }

            this.Listen = function (el, evt) {
                /// <summary>Registers the SyncWindow object as a listener of an event</summary>
                /// <param name='el'>HTML element</param>
                /// <param name='evt'>string event name</param>
                /// <returns>void</returns>
                if (el.addEventListener) {
                    el.addEventListener(evt, this.Open);
                }
                else {
                    el.attachEvent("on" + evt, this.Open);
                }
            }

            this.GetBrowserWindowRef = function () {
                /// <summary>Returns a browser window refernce for the window that was created by the SyncWindow object</summary>
                /// <returns>browser window ref</returns>
                return m_winRef;
            }
        }
        
        window.SyncWindow = new Tyler.SyncWindow();
    }) (window);
</script>




<input name="DTLNavigator$idx" type="hidden" id="DTLNavigator_idx" value="1" />
<table class="WidgetBar" cellspacing="0" cellpadding="0" width="130" border="0">

    <tr align="middle">
        <td nowrap align="middle" width="130">
            <table cellspacing="1" cellpadding="1" width="100%" align="center" border="0">
                <tr valign="middle">
                    <td valign="middle" align="middle" nowrap="noWrap" style="width: 25px; vertical-align: middle;">
                        
                        
                    </td>
                    <td width="100%" valign="middle" align="center">
                        <input name="DTLNavigator$txtFromTo" type="text" id="DTLNavigator_txtFromTo" style="font-size:&#32;11px;&#32;color:&#32;#3d6599;&#32;width:&#32;100%;&#32;border-top-style:&#32;groove;&#32;border-right-style:&#32;groove;&#32;border-left-style:&#32;groove;&#32;border-bottom-style:&#32;groove;&#32;border-bottom-width:&#32;1px;&#32;border-top-width:&#32;1px;&#32;border-right-width:&#32;1px;&#32;border-left-width:&#32;1px;&#32;border-bottom-color:&#32;#3d6599;&#32;border-top-color:&#32;#3d6599;&#32;border-right-color:&#32;#3d6599;&#32;border-left-color:&#32;#3d6599;&#32;height:&#32;100%;&#32;background-color:&#32;white;&#32;text-align:&#32;center;&#32;vertical-align:&#32;top;" readonly="READONLY" value="1&#32;of&#32;1" />
                    </td>
                    <td valign="middle" align="middle" nowrap="noWrap" style="width: 25px; vertical-align: middle;">&nbsp;
                        
                    </td>
                </tr>
            </table>
        </td>
    </tr>


    
    

    <tr style="height: 6px;">
        <td></td>
    </tr>
    <tr>
        <td>
            <hr />
        </td>
    </tr>
    <tr align="left" style="height: 25px;">
        <td class="SideBarHeading" style="height: 8px; text-align: left;">Actions</td>
        <br />
    </tr>
    <tr align="left">
        <td></td>
    </tr>


    
    <tr valign="middle">
        
    </tr>
    
    <tr valign="middle">
        
    </tr>
    
    <tr>
        <td id="DTLNavigator_lbPrintAll" align="left">
            <a href="javascript:PrintDatalet('all')">
                <image alt="Show all data in a printer friendly version" border="no"
                    onclick="PrintDatalet()" src="../images/new/print.png" alt="Print">
                    <font class="SideBarTabs" valign="middle" ><span style="vertical-align:top;" >Printable Summary</span></font></image>
            </a>
        </td>
		
    </tr>
    
    <tr>
        <td id="DTLNavigator_lbPrintThis" valign="middle" align="left">
            <a href="javascript:PrintDatalet()">
                <image alt="Show this page in a printer friendly version" border="no" onclick="PrintDatalet()" alt="Print" src="../images/new/print.png">
                    <font class="SideBarTabs"><span style="vertical-align:top;" >Printable Version</span></font></image>
            </a>
        </td>
		
    </tr>
    

    <!--new workflow stuff starts here-->
    
    <tr>
        <td>
            <hr id="DTLNavigator_hrReports"></hr>
        </td>
    </tr>
    <tr align="left">
        <td align="left">
            

<script src="../Script/Soap.js"></script>


<script language="javascript">

    var WS_REPORTS = "../Reports/Services/RptServices.asmx";
    var isReportProcessingStarted = false;
    var isMobile = false;

    


    function IsReportProcessingStarted() {
        return isReportProcessingStarted;
    }
    
  function ReportProcessingStart() {
//debugger;
      isReportProcessingStarted = true;
      //alert("Show Message");
      document.getElementById("divMessage").style.display = ""; //loading ... 
      document.getElementById("ReportListButton").disabled = "disabled";
      //var listboxName = DTLNavigator_Report2"_ReportsListBox";
      //alert(listboxName);
      document.getElementById("DTLNavigator_Report2_ReportsListBox").disabled = "disabled";
      //---      
      document.getElementById("divReportDone").style.display = "none";
  }
  function ReportProcessingEnd()
  {      
      //debugger;
      isReportProcessingStarted = false;
      document.getElementById("divMessage").style.display = "none";//loading...
      document.getElementById("ReportListButton").disabled = "";      
      document.getElementById("DTLNavigator_Report2_ReportsListBox").disabled = "";
      document.getElementById("divReportDone").style.display = "inline";
  }
  function ReportProcessingError()
  {   
      //debugger;
      isReportProcessingStarted = false;
      document.getElementById("divMessage").style.display = "none";//loading...
      document.getElementById("ReportListButton").disabled = "";      
      document.getElementById("DTLNavigator_Report2_ReportsListBox").disabled = "";
      document.getElementById("divReportDone").style.display = "none";

  }
    

  function ReportReady() {
    //debugger;
    // 3 hide display layer
    document.getElementById("divMessage").style.display = "none";
    //enable the GO button
    document.getElementById("ReportListButton").disabled = "";
    

    if (ifrDetails.document.location.href.indexOf("ErrorMessage") > -1) {
      ShowPopup();
    }
    
    //calling the report as a WS    
    AKA_XMLNS = "http://tempuri.org/";
   
    SoapInvokeService(WS_REPORTS, "IsReportReady", null, true, null, "ReportReadyComplete");
      //alert("isReportReady= " + isReportReady);
  }
 
  function ReportReadyComplete(oHttp, methodName){
 
      if(oHttp.readyState == 4) {
      
          var isReportReady = SoapGetResult(oHttp, methodName);
          if (isReportReady && isReportReady == "true") {
              //document.getElementById("divReportReady").style.display = 'inline';
              document.getElementById("divReportDone").style.display = 'inline';
              ReportProcessingEnd();
              //OpenReport();//does not work. Window opens and cloes.
          }
      }
  }

  function ShowPopup() {

    if (document.getElementById("ifrDetails").src != "") {
      document.getElementById("divErrormessage").style.display = 'inline';
    }
  }

  function OpenReport() {
      //debugger;
      document.getElementById("divReportDone").style.display = "none";
    window.open("../reports/displayreport.aspx");
  }
</script>

<input id="hasData" value="false" type="hidden" />


<table id="Table1" cellspacing="1" cellpadding="1" border="0" width="162px" >
  <tr>
    <td style="height: 8px; text-align: left;" class="SideBarHeading">
        <label for="DTLNavigator_Report2_ReportsListBox">Reports</label>
    </td>
  </tr>
  <tr>
    <td valign="middle" >
      &nbsp;<select size="4" name="DTLNavigator$Report2$ReportsListBox" id="DTLNavigator_Report2_ReportsListBox" onDblclick="ShowReport()" Label="" style="width:99%;">
			<option value="CSVMailingList">Mailing List</option>
			<option value="rptPRCMainMultiCard">PRC Report</option>

		</select>
    </td>
  </tr>
  <tr>
    <td style="text-align: right; height: 25px" class="SideBarTabs">
      <div style="display:none; vertical-align: top; text-align:left;position:relative;left:-7px;top:8px;" id="divViewReportStatus">        
        <a onmouseover="this.style.textDecoration='underline';this.className='';" onmouseout="this.style.textDecoration='';this.className='';" onclick="OpenViewODReportsWindow()"> <span style="cursor:pointer; font-size:8px; align:left;"> View Report Status</span></a>&nbsp;
	    </div>         
         
<button type="button" id="ReportListButton" name="ReportListButton" class="Btn BtnGreen" onclick="ShowReport()" title="Show Report" style="clear:both;" >Go</button>
                <input id="ReportsListParIDs" name="ReportsListParIDs" style="width: 24px; height: 3px" type="hidden" size="1" value=""> 
    </td>
  </tr>
  <tr valign="middle">    
    <td class="BannerSubTabs" valign="middle">      
      <div style="display: none; vertical-align: top; background-color:White;" id="divMessage">
        <img  src="../images/ajax/ajax-loader-circle.gif" style="height: 22px; width: 21px" />&nbsp;<span class="BannerSubTabs">Processing...</span>

      </div>
      <div style="display: none; vertical-align: top;" id="divReportDone">
        <a title="Open Report" onclick="OpenReport();"> <b><span style="cursor:pointer">Report complete.<br /><u>Click</u> to open</b></span></a>
	    </div>
    </td>
  </tr>
</table>


<div id="divErrormessage" class="WidgetBar"  
  style="border: 2px outset #cccccc; FONT-SIZE: 11px; Z-INDEX: 10000; display:none; LEFT: 186px; WIDTH:600px; HEIGHT: 400px; COLOR: blue; POSITION: absolute; TOP: 114px; 
  background-color: #FFFFCC;">
	<table style="height:200px" cellspacing="0" cellpadding="3" border="0">
	<tr class="BannerSubTabs">
	<td align="left" width: 700px; height: 15px; style="background-color:#3d6599;" class="SideBarHeading"   ></td>
	<td align="right" width: 700px;height: 15px; class="SideBarHeading" style="background-color:#3d6599; text-align:right;" >
	  <a title="Close" 
      onclick="document.getElementById('divErrormessage').style.display='none'" 
      href="#" style="color: #FFFFFF"><b>[X]</b></a>
	</td>
	</tr>
	<tr class="BannerSubTabs">
	<td colspan="2" style="width:700px;" align="center">
	    <iframe id="ifrDetails" name="ifrDetails" src="../Config/Blank.htm" scrolling="no"
        border="0" frameborder="0" marginwidth="0" align="top" 
        onload="ReportReady()" class="BannerSubTabs" style="width:560px; height:350px; margin-top: 0px; background:inherit;" >
      </iframe>
	</td>
	
	</tr>
</table>
</div>

 
<script>
  var gReport = "";
  var ReportWindow = null;
  var useNewWindow = "true";
  function GetReportID() {
    var retVal = trim(gReport);

    if (!((retVal) && (retVal.length > 0))) {
      var reportListBox = document.getElementById("DTLNavigator_Report2_ReportsListBox");

      if (reportListBox.selectedIndex >= 0) {
        retVal = reportListBox.options[reportListBox.selectedIndex].value;
      }
      else
        return "";
    }

    return retVal;
  }

    function ShowReport() {
        if (IsReportProcessingStarted()) {
            return;
        }
    //debugger;
    var ret = true;
    var reportID = GetReportID();
    var parID;
    var taxYr = null;
    var jur = "";
    ReportProcessingStart();      
    //-----------------------------------------------------------------
    if (window.SetCheck)
      SetCheck();

        //added this call to see if at least one parcel is selected
        try
        {
            var pinCount = setSelection();
            if (pinCount == 0) {
                alert("Please select at least one parcel");
                ReportProcessingError();
                return;
            }      
        }
        catch(e){   }
        

    parID = GetParIDs();
    taxYr = GetTaxYear();
    jur = GetJur();
    
    if ((!reportID) || (reportID.length == 0)) {
        alert("Please select a report.");
        ReportProcessingError();
      document.getElementById("ReportListButton").disabled = "";
      return false;
    }

    if ((!(parID)) || (parID == '')) {
        alert("Please select at least one Parcel.");
        ReportProcessingError();
      document.getElementById("ReportListButton").disabled = "";
      return false;
    }
    //*******IAS Report integration Nov.15.2010
    if (reportID == "IAS_PRC") //IAS reports
    {
      try {
        document.getElementById("ReportListButton").disabled = "";
        SetIASReportData(parID, taxYr, jur); //post this to session and use for processing!
        var url = "../reports/tylerreports/iasreportlist.aspx?taxYr=" + taxYr; //
        var iasRptDlg = "center:yes; help:no; resizable:yes; status:no: title:yes";

        //alert("You have picked IAS Report");
        var retIas = window.showModalDialog(url, iasRptDlg);
        return false;
      }
        catch (e) {
            ReportProcessingError();
        alert("There was an error processing parcels. Please try again.");
        return false;
      }
    }
    //*******Clermont BL320OH Tax Summary IAS Report integration Nov.23.2011
    if (reportID == "CLERMONT_BL320OH") //IAS reports
    {
      try {
        document.getElementById("ReportListButton").disabled = "";
        ReportProcessingStart();
        //var queryString = "http://66.161.222.108/Reporting/getreport.aspx?parid=" + parID + "&jur="+jur + "&taxyr=" + taxYr + "&reportName=BL320OH&userId=120&roleId=1";
        //var queryString = "reportClass=" + reportID + "&parid=" + parID + "&taxyr=" + taxYr + "&jur=" + jur + "&internal=";
        //var url = queryString;
        var queryString = "reportClass=" + reportID + "&parid=" + parID + "&taxyr=" + taxYr + "&jur=" + jur + "&internal=";
        var url = '../reports/ShowReport.aspx?' + queryString;
        //alert(url);
        document.getElementById("ifrDetails").src = "";
        document.getElementById("ifrDetails").src = url;
        
        //SetIASReportData(parID, taxYr, jur); //post this to session and use for processing!        
        return false;
      }
        catch (e) {
            ReportProcessingError();
        alert("There was an error processing parcels. Please try again.");
        return false;
      }
    }
    //-----------IAS OD Report-------------------------------
    if (reportID.indexOf("TODR-") == 0) //IAS OD reports
    {
      document.getElementById("ReportListButton").disabled = "";
      if(SetIASReportData(parID, taxYr, jur)=="false")//checks for max size
      {
          ReportProcessingError();
          return;
      } //post this to session and use for processing!
      else
      {
          RunODReport(reportID, parID, jur, taxYr);
          return; //page won't refresh... 
      }

    }
    //----------------------------------------------------------------------------------------------------------------
    else if (window.ReportMaxListSizeValidate) {
      var parIDLength = parID.replace(/,$/, "").split(',').length; //removing last comma and getting length of arry      
      var result = ReportMaxListSizeValidate(reportID, parIDLength);
      if (result > 0) {
          alert("This report can only be used with " + result + " selections. Please reduce the number of items and try again.");
          ReportProcessingError();
        return false;
      }
    }

    if (window.ReportExtraInfo) {
      var isSuccess = 1;
      isSuccess = ReportExtraInfo(reportID);

      if (isSuccess != 1)
        return false;
    }

    //ReportProcessingStart();
    var queryString = "reportClass=" + reportID + "&parid=" + parID + "&taxyr=" + taxYr + "&jur=" + jur + "&internal=";
    var url = '../reports/ShowReport.aspx?' + queryString;
    //alert(url);
    document.getElementById("ifrDetails").src = "";
    //BDias. Sept.2.2010: if the query string is tool long then set the parid:jur:taxyr info in session via webService
    if (url.length > 1000) {
      var WS_REPORTS = "../Reports/Services/RptServices.asmx";
      var pars = SoapBuildInputParam("sParidList", parID);
      SoapInvokeService(WS_REPORTS, "SetCSVPins", pars, false); //calling webservice synchronously!	
      //set the parid to empty or else it won't send the request
      parID = "";
      queryString = "reportClass=" + reportID + "&parid=" + parID + "&taxyr=" + taxYr + "&jur=" + jur + "&internal=";
      url = '../reports/ShowReport.aspx?' + queryString;
    }
     document.getElementById("ifrDetails").src = url;
        //window.open(url);
     if(isMobile)
     {
         ReportProcessingEnd();
         window.open(url);
     }

    return true; //page won't refresh... BD Sept.11.2008    
  }

  function RunODReport(_reportName, parID, jur, taxYr) {
    
    ReportProcessingStart();

    var _message = "";
    
    //call web service to call report generation
    //pars = SoapBuildInputParam("moduleName", "MODULE_AC");
    pars = SoapBuildInputParam("reportName", _reportName);
    pars += SoapBuildInputParam("parid", parID);
    pars += SoapBuildInputParam("jur", jur);
    pars += SoapBuildInputParam("taxyr", taxYr);
    //alert("before call");
    var WS_REPORTS = "../Reports/Services/RptServices.asmx";
    retMessage = SoapInvokeService(WS_REPORTS, "GenerateIASODReport", pars, false); //*Not async call
    //alert(retMessage);

    if (retMessage != null && !retMessage.indexOf("Error")==0) {
      //_message += "Your Report has been submitted. Job Number: " + retMessage + "\n\r" + ". To check to see if your report is ready: Click \"View Reports\" now <br /> -OR- <br />Choose “View On Demand Report Status�? later from the Run Reports icon in the Activity Center or Manage Workflow Task toolbars";
      _message += "Your report has been submitted. Job#: " + retMessage + ".";
      //alert(_message);      
      //OpenViewODReportsWindow();
      OpenODReportsWindow(retMessage,_reportName);
    }
    else {
        ReportProcessingError();
        _message += "An error was returned by the On Demand Report Engine:\n\r"  + retMessage + ".";
      alert(_message);
    }
    
  }

  var myVar;//used for stopping the timer 
  var intervalTime = 5000;//2 seconds
  var numberOfPing = 0;
  var pingRound = 0;


  //Will try to open the report on screen.
  function OpenODReportsWindow(jobNo, reportName) {
           
      //re-initilize vars
      myVar="";//used for stopping the timer 
      intervalTime = 5000;//mili seconds
      numberOfPing = 0;
      pingRound = 0;

      myCheckFunction(jobNo, reportName);
            
  }
  function StopTimer() {
      clearInterval(myVar);      
  }

  function myCheckFunction(jobNo, reportName)
  {
      myVar=setInterval(function(){checkStatusService(jobNo, reportName)},intervalTime);//check every 2 seconds
  }

  function checkStatusService(jobNo, reportName)
  {
      //debugger;
      var retMessage = "";
      if (numberOfPing == 3 && pingRound==1)
      {
          StopTimer();
          //start it again
          intervalTime = 8000;//mili seconds
          myCheckFunction(jobNo,reportName);
          //return true;
      }
      else if (numberOfPing ==7 && pingRound == 2) {
          StopTimer();
          intervalTime = 10000;//mili seconds
          myCheckFunction(jobNo,reportName);
          //return true;
      }
      else if (numberOfPing > 17) {
          StopTimer();//Timeout Error after n number of pings
          alert("The report is taking too long to generate. Please click on the View Report Status link to check status.");
          ReportProcessingEnd();
          return true;
      }

      
      //call WS to get report file      
      pars = SoapBuildInputParam("jobNo", jobNo);
      pars += SoapBuildInputParam("reportName", reportName);
      
      //alert("before call");
      var WS_REPORTS = "../Reports/Services/RptServices.asmx";
      retMessage = SoapInvokeService(WS_REPORTS, "GetIASODReportGetFile", pars, false); //*Not async call
      //alert(retMessage);
      if (retMessage == "True") {
          StopTimer();
          var url = "../reports/showODreport.aspx?jobNo=" + jobNo + "&fileName="+reportName;
          //alert(url);
          //document.getElementById("ifrDetails").src =url;
          window.open(url);
          ReportProcessingEnd();
          document.getElementById("divReportDone").style.display = "none";//hide the click to open for OD reports
          return true;
      } else {
          if (retMessage != "False") {
              StopTimer();
              alert("No data");
              ReportProcessingEnd();
              document.getElementById("divReportDone").style.display = "none";//hide the click to open for OD reports
              return true;
          }
      }
      numberOfPing++;
      //alert("reportName=" + reportName + " retMessage= "+ retMessage + " " + numberOfPing + " pingRound="+pingRound);
  }
    
  //View On Demand Report
  function OpenViewODReportsWindow() {
       
    var url = "../maintain/ReportSpool.aspx?simple=1&runningJobs=0&batchJobs=999";
    var title = "View On Demand Report";
    var w = "1000";
    var h = "750";
    var left = (screen.width / 2) - (w / 2);
    var top = (screen.height / 2) - (h / 2);
    newwin = window.open(url, "Reports", "toolbar=no, location=no, directories=no, status=no, menubar=no, status=1, toolbar=0, height=650px, width=" + w + ", height=" + h + ", resizable=1, scrollbars=yes" + ", top=" + top + ", left=" + left);

    if (window.focus) { newwin.focus() }
    return false;
  }

  function SetIASReportData(parID, taxyr, jur) {
      //debugger;
      var maxSize = 100;
      var currentSelectedSize = parID.split(",");

      //alert("Max Size = " + maxSize + ". Current Size=" + currentSelectedSize.length);
      if(currentSelectedSize.length > maxSize)
      {
          var reportID = GetReportID();
          //alert("This report can only be used with " + maxSize + " selections. Please reduce the number of items and try again.");
          alert(reportID + " can only be generated for " + maxSize + " records at a time through the On Demand Reports Engine.  Please reduce your selected set and try again.");

          return "false";
      }
      
      try {
      var ret = "";
      var WS_REPORTS = "../Reports/Services/RptServices.asmx";
      var pars = SoapBuildInputParam("parids", parID);
      pars += SoapBuildInputParam("jur", jur);
      pars += SoapBuildInputParam("taxyr", taxyr);
      ret = SoapInvokeService(WS_REPORTS, "SetIASReportDataForDatalet", pars, false); //calling webservice synchronously!

      if (ret == null || ret == "false" || ret.length == 0) {
        alert("There was a problem in processing your data. Please try again.");
        return false;
      }
    }
    catch (e) {
        ReportProcessingError();
        alert("There was an error processing parcels. Please try again.");
      return false;
    }

  }

  function GetJur() {
    var _jur = "";
    var jurObj = "";
    try {
      jurObj = document.getElementById("hdJur");
      if (jurObj != null) {
        _jur = jurObj.value;
      }
      else {
        jurObj = document.getElementById("hdXJur");
        if (jurObj != null) _jur = jurObj.value;
      }
    }
    catch (e) {
    }
    return _jur;

  }

  function GetParIDs() {
    //debugger;
    var pin = null;



    var temp = document.getElementById("hdXPin");
    if (temp) {
      pin = temp.value;
    }
    else if (parent.document.frmMain) {
      if (parent.document.frmMain.hdSelected) {
        var pins = document.getElementById('AkaCfgResults_hdPins');

        if (pins) {
          var pagePins = getSelected("");
          setSelection();
          pins.value = pagePins;
          pin = getSelectedPins();
        }
      }
      else if (pin = parent.document.frmMain.hdPin) {
        pin = parent.document.frmMain.hdPin.value;
      }
      else {
        var pins = document.getElementById('AkaCfgResults_hdPins');

        if (pins) {
          var pagePins = getSelected("");
          setSelection();
          pins.value = pagePins;
          pin = getSelectedPins();
        }
      }
    }
    else if (parent.document.zoom)
      if (parent.document.zoom.hdXPin) {
        pin = parent.document.zoom.hdXPin.value;
      }
      else if (parent.document.zoom.hdSelected) {
        pin = parent.document.zoom.hdSelected.value;
      }
      else if (parent.document.Form1)
        pin = parent.document.Form1.hdXPin.value;

    return pin;
  }

  function GetForm() {
    var form = null;

    if (parent.document.frmMain)
      form = parent.document.frmMain;
    else if (parent.document.zoom)
      form = parent.document.zoom;
    else if (parent.document.Form1)
      form = parent.document.Form1;

    return form;
  }

  function GetTaxYear() {
      var form = GetForm();
      var taxYr;
      var navDdTaxYr = document.getElementById("DTLNavigator_ddTaxYear");

    if (form) {
      if (form.inpTaxyr)
        taxYr = form.inpTaxyr.options[form.inpTaxyr.selectedIndex];
      else if (form.ddTaxYear)
          taxYr = form.ddTaxYear.options[form.ddTaxYear.selectedIndex];
      else if(navDdTaxYr)
      {
          taxYr = navDdTaxYr.options[navDdTaxYr.selectedIndex];
      }
    }

    if (taxYr)
      return taxYr.value;
    else {
      var taxYrObj = document.getElementById("hdXTaxYr");
      if (taxYrObj != null && taxYrObj.value != "") return taxYrObj.value;

      taxYrObj = document.getElementById("hdTaxYear");
      if (taxYrObj != null) return taxYrObj.value;
      else return null;
    }

  }

  function EnableView(viewButton, val) {
    var buttonSetup = document.getElementById(viewButton);
    buttonSetup.disabled = false;

    gReport = val;
  }

//Mobile deivice logic
    //if ((navigator.userAgent.indexOf('iPhone') != -1) || (navigator.userAgent.indexOf('iPod') != -1) || (navigator.userAgent.indexOf('iPad') != -1)) {
  if( /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent) ) {

      isMobile = true;              
  } 
  else
  {
      //add style back for non-mobile
      isMobile = false;
      var reportListBox = document.getElementById("DTLNavigator_Report2_ReportsListBox");
      reportListBox.className  = "SideBarTabs";      
  }

</script>


        </td>
    </tr>
    
    <tr>
        <td>
            <hr />
        </td>
    </tr>
    <tr align="left" style="height: 25px;">
        <td class="SideBarHeading" style="height: 8px; text-align: left;">
            <div id="DTLNavigator_divExtraLinksLabel">Links</div>
        </td>
        <br />
    </tr>
    <tr>
        <td valign="middle" align="left">
            <div id="DTLNavigator_dvAnyLinks"><html>
<body>
<br><a href="../content/NHCodes_Descriptions.pdf" target ="_blank">Codes and Descriptions</a>
<br>
<br>
<a href="../content/NHSoilCodes_Descriptions.pdf" target ="_blank">Soil Codes and Descriptions</a>
<br>
<br>

</div>
        </td>
    </tr>
    

</table>
<input name="DTLNavigator$hdRecCount" type="hidden" id="DTLNavigator_hdRecCount" value="0" />
<input type="hidden" id="hdXPin" name="hdXPin" value="L9NE4C 20 2       0310" />
<input type="hidden" id="hdXJur" name="hdXJur" value="048" />
<input type="hidden" id="hdXTaxYr" name="hdXTaxYr" value="" />
<input name="DTLNavigator$hdTaxYear" type="hidden" id="DTLNavigator_hdTaxYear" />
<input type="hidden" id="hdXUserName" name="hdXUserName" value="Default" />

<script> 

 var customReportBuilderRows= 0;
 var bImplementsAjaxRecordNavigator = false;
 var hasCustomReportBuilder = false;
   
    function OpenContactUs(id) {
    var parid = document.getElementById("hdXPin").value;
    var jur = document.getElementById("hdXJur").value;
    var taxyr = document.getElementById("hdXTaxYr").value;
    var ddtaxyr = document.getElementById("ddTaxYear");
    if (ddtaxyr != null) {
        taxyr = ddtaxyr.value;
    }

    var url = '/contactus/#/contact/'+id + '/' + jur + "/" + parid + "/" + taxyr+ "/";


    window.open(url,"_blank");
}


function CheckSyncWindow() {
    if (window.SyncWindow.IsEnabled()) {
        window.SyncWindow.SetCallback(GetSyncWindowParams);
        window.SyncWindow.Open();
        if (bImplementsAjaxRecordNavigator) {
            el = document.getElementById("DTLNavigator_imageFirst");
            if (el!=null) window.SyncWindow.Listen(el, "click");
            el = document.getElementById("DTLNavigator_imagePrevious");
            if (el!=null) window.SyncWindow.Listen(el, "click");
            el = document.getElementById("DTLNavigator_imageNext");
            if (el!=null) window.SyncWindow.Listen(el, "click");
            el = document.getElementById("DTLNavigator_imageLast");
            if (el!=null) window.SyncWindow.Listen(el, "click");

        }
    }
}

function GetSyncWindowParams() {
    //debugger;
    var sPinVal = document.getElementById("hdXPin").value;
    if (sPinVal != "") {
        var sUserNameVal = document.getElementById("hdXUserName").value;
        var sJurVal = document.getElementById("hdXJur").value;
        var sTaxYearVal = document.getElementById("hdXTaxYr").value;
        var parms = {
            "USERNAME": sUserNameVal,
            "JUR": sJurVal,
            "PARID": sPinVal,
            "TAXYR": sTaxYearVal
        };
        return parms;
    }
    else {
        return null;
    }
}

//window.attachEvent("onload", CheckSyncWindow);

function OpenIMEdit(param) {
    var url = "../Maintain/SelectTransForEdit.aspx?frameset=1&" + param;
    var dlg = "center:yes; dialogHeight:500px; dialogWidth:500px; help:no; resizable:yes; status:no; title:no";
    var iMUrl = window.showModalDialog(url, null, dlg);
    if (iMUrl != null) {
        var name = "iMaintainTransaction";
        var s = "location=no, resizable=yes, status=yes, menubar=no"
        window.open(iMUrl, name, s);
    }
}

function EditThis() {
    var parid = document.getElementById("hdXPin").value;
    var jur = document.getElementById("hdXJur").value;
    var taxyr = document.getElementById("hdXTaxYr").value;
    var ddtaxyr = document.getElementById("DTLNavigator_ddTaxYear");
    if (ddtaxyr != null) {
        taxyr = ddtaxyr.value;
    }

    OpenIMEdit("pin=" + parid + "&jur=" + jur + "&taxyr=" + taxyr);
}
function CreateWorkflow() {
  var WS_WF = "../ActivityCenter/Services/WF.asmx";

  var pars = SoapBuildInputParam("recordCount", 1);
  var ret = SoapInvokeService(WS_WF, "CanUserCreateWorkflow", pars, false);
  if (ret) {
      ret = ret.replace("[br]", "\n");
      ret = ret.replace("[br]", "\n");
      ret = ret.replace("[br][br]", "\n\n");
      alert(ret);
      return;  // user is not allowed to create workflow
  }

  var parid = document.getElementById("hdXPin").value;
  var jur = document.getElementById("hdXJur").value;
  var taxyr = document.getElementById("hdXTaxYr").value;
  var ddtaxyr = document.getElementById("ddTaxYear");
  if (ddtaxyr != null) {
      taxyr = ddtaxyr.value;
  }
  //var url = "../activitycenter/eventpublisher/publishevents.aspx?parid=" + parid + "&jur=" + jur + "&taxyear=" + taxyr;
  var url = "../activitycenter/eventpublisher/publishevents.aspx?parid=" + jur + ":" + parid + ":" + taxyr;
  var dlg = "center:yes; dialogHeight:800px; dialogWidth:500px; help:no; resizable:no; status:no; title:no; scroll:off;";
  window.showModalDialog(url, null, dlg);
}
//scripts for Action Links **********************************
function NBHDSearch(pin) {
  var url = "../Search/NbhdSearch.aspx?pin=" + pin + "&fromDL=yes";
  url += "&idx=" + sIdx + "&sIndex=" + sIndex;
  //	alert(url);
  document.frmMain.hdSkip.value = 1;
  document.location.href = url;
}
function ProximitySearch(pin) {
    var url = "../Search/proximitysearch.aspx?pin=" + pin + "&fromDL=yes";
    url += "&idx=" + sIdx + "&sIndex=" + sIndex;    
    document.frmMain.hdSkip.value = 1;
    document.location.href = url;
}
function PrintDatalet(val) {
  var doc = document.frmMain;
  var url = '../Datalets/PrintDatalet.aspx?pin=' + doc.hdPin.value +
		'&gsp=' + doc.hdMode.value + '&taxyear=' + doc.hdTaxYear.value +
		'&jur=' + doc.hdJur.value + '&ownseq=' + doc.hdOwnSeq.value +
		'&card=' + doc.hdCard.value + '&roll=' + roll +
		'&State=' + doc.hdMask.value + '&item=' + doc.hdItem.value +
		'&items=' + doc.hdItems.value + '&all=' + val + "&ranks=" + ranks;
  window.open(url, 'Print', "location=no, scrollbars=yes, menubar=yes, resizable=yes, width=900");
}

function CreateComparables() {
    var url = "../Forms/Comparables.aspx?showSearch=1&sIndex=" + sIndex + "&idx=" + sIdx;
    window.location.href = url;
}

$(document).ready(function () {

    if(hasCustomReportBuilder)
    {
        getLeftMenuLinks();
    }       
});    

function getLeftMenuLinks(){    
    
    var links = [];
    if($('#sidemenu')){
        
        $('#sidemenu').find('a').each(function() {
        
            var href = $(this).attr('href');
            var dataletLabel = $(this).find('span').text();        
            if(href.indexOf('datalet.aspx') != -1)
            {
                var stringArray  = href.split("=");
                var idx = stringArray[1].indexOf('&');
                if(idx != -1)
                {
                    var dataletName = stringArray[1].substr(0, idx);
                    links.push({ Name: dataletName, Label: dataletLabel, Position: 0})                  
                }
            }
        
            if(href.indexOf('photoview.aspx' ) != -1 || href.indexOf('photoviewfixed.aspx') != -1){
                links.push({ Name: "PHOTO", Label: dataletLabel })   
            }
            if(href.indexOf('map.aspx') != -1){
                links.push({ Name: "MAP", Label: dataletLabel })   
            }
            if(href.indexOf('sketchframe.aspx') != -1){
                links.push({ Name: "SKETCH", Label: dataletLabel })   
            }
            if(href.indexOf('pictometryipa.aspx') != -1){
                links.push({ Name: "PICTOMETRY", Label: dataletLabel })   
            }
        });
    
        customReportBuilderRows = links.length;
 
        var request = $.ajax({
            dataType: "json",
            url: "../api/printformdata/links",
            method: "POST",
            data: { '': links },
            //success: openCustomReportBuilderWindow()        
        });
    }    
}

function openCustomReportBuilderWindow(){

    var doc = document.frmMain;   

    var url = '../Datalets/CustomReportBuilder/printForm.html?pin=' + doc.hdPin.value +
          '&gsp=' + doc.hdMode.value + '&taxyear=' + doc.hdTaxYear.value +
          '&jur=' + doc.hdJur.value + '&ownseq=' + doc.hdOwnSeq.value +
          '&card=' + doc.hdCard.value + '&roll=' + roll +
          '&State=' + doc.hdMask.value + '&item=' + doc.hdItem.value +
          '&items=' + doc.hdItems.value + '&all=' + "" + "&ranks=" + ranks;  
    
    var windowHeight = customReportBuilderRows*25 + 200;
    
    var options = "location=no, scrollbars=yes, menubar=yes, resizable=yes, width=500, height="+ windowHeight;

    window.open(url, 'CustomReportBuilder', options);
}

//END Action Links ******************************************

</script>


<!-- record navigator: end -->

						</td>
	</tr>
					     
				    </table>
               
</div></div>
        </div>
	

  </tr> 
  </table>

	
	<input name="hdIDX" type="hidden" id="hdIDX" />
	<input name="hdPin" type="hidden" id="hdPin" value="L9NE4C&#32;20&#32;2&#32;&#32;&#32;&#32;&#32;&#32;&#32;0310" />
    <input name="hdItemKey" type="hidden" id="hdItemKey" />
	<input name="hdTaxYear" type="hidden" id="hdTaxYear" value="2023" />
	<input name="hdJur" type="hidden" id="hdJur" value="048" />
	<input name="hdOwnSeq" type="hidden" id="hdOwnSeq" value="0" />
	<input name="hdMode" type="hidden" id="hdMode" value="PROFILE_NH" />
	<input name="hdCard" type="hidden" id="hdCard" value="1" />
	<input name="hdRoll" type="hidden" id="hdRoll" />
	<input name="hdItem" type="hidden" id="hdItem" value="1" />
	<input name="hdItems" type="hidden" id="hdItems" value="-1" />
	<input name="hdMask" type="hidden" id="hdMask" value="1" />
	<input name="hdsIndex" type="hidden" id="hdsIndex" value="-1" />
	<input type="hidden" id="hdSkip" name="hdSkip">

</td>
</tr>
</table>


<Script lang='javascript'>
<!--
function ReportMaxListSizeValidate(reportID, parIdCount)
{

	if (reportID == 'CSVMailingList')
	{
		if (parIdCount > 100)
				return 100
	}

	if (reportID == 'rptPRCMainMultiCard')
	{
		if (parIdCount > 50)
				return 50
	}

	return 0;
}
-->
</Script>
</form>


<script >
//    alert("onload");
    //window.attachEvent("onload", WindowOnLoad);
    //window.addEventListener("onload", WindowOnLoad);
    window.onload=function(){
        
        var skip = document.frmMain.hdSkip;
        if (skip != null && skip.value == "1") {
            frmMain.submit();
        }
	
        //added for map
        //alert(bOnLoad);
        if (bOnLoad) body_onload();
        if (newPinFromMap) updateFrmPinParam();
    };
        
</script>

<script>

var gPin = "L9NE4C 20 2       0310";
var thisURL = "http://www.ncpub.org/_web/datalets/datalet.aspx?mode=profile_nh&UseSearch=NO&pin=L9NE4C%2020%202%20%20%20%20%20%20%200310";
var sIndex	= "-1";
var sIdx	= "1";
var ranks	= "Datalet";
var DLG_EMAIL = "center:yes; dialogHeight:300px; dialogWidth:600px; help:no; resizable:no; status:no; title:no";
var roll = "REAL";
var newPinFromMap = false;

function showReport(el)
{
	var report = document.getElementById("Reports_currentrecs");
	var val = trim(report.value);
	if (val.length == 0)
	{
		alert("Please select report");
		return;
	}
	var url = "../Reports/ReportAV2.aspx?Pin=" + gPin + "&Report=" + val;
	window.open(url, 'w1', 'width=800, height=600');
}

function OpenIRIncident(id)
{
	var url = '../Respond/IncidentDetail.aspx?id=' + id + '&mode=link';
	
		var s;
		s = "fullscreen=no, titlebar=no, location=no, menubar=no, toolbar=no, resizable=yes, scrollbars=yes, status=yes, titlebar=yes"; 
		return window.open(url, 'OpenIRIncident', s);
}

window.focus();

</script>


</div></section></div></div><footer><div id='footerarea' ><div id='footercontent' ><div id="dvFooter" style="clear:both;&#32;text-align:center;">
<head>
<style type="text/css">
.auto-style1 {
	font-family: Verdana, Geneva, Tahoma, sans-serif;
}
</style>
</head>

<table border="0" width="100%" align="center" id="Table1">
	<tr>
		<td align="center" width="100%">
			<table border="0" class="AkandaCopyright" id="Table2">
				<tr>
					<td align="center" valign="bottom" class="auto-style1">
						<font size="2">
						<a href="http://support.iasworld.net/company.htm" target="_blank"></a>
						</font>
					</td>
					<td class="AkandaCopyright" align="center">
					<font size="2">
						<br class="auto-style1">
					</font>
					<font size="1">
						<span class="auto-style1">Data Copyright Northampton County<br><strong>Last Updated</strong>:&nbsp;</span><span id="lbLastUpdate" class="auto-style1">08/May/2023</span><span class="auto-style1">
					</span>
					</font>
					<br>
					<span class="auto-style1">
					<font size="1">
					Powered by <a href="http://www.tylertech.com">iasWorld Public Access</a>. All
					rights reserved.
					</font>
					</span>
					</td>
				</tr>
			</table>
		</td>
	</tr>
</table>
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-49285134-14', 'auto');
  ga('send', 'pageview');

</script></div>
<script>
var tmpLate = "../images/cur_template";
function getTmpLate(par)
{
	return tmpLate + par;
}
</script>


</div></div></footer></body><script language=JavaScript>window.setTimeout("window.location.replace('../forms/htmlframe.aspx?mode=content/home.htm')", 3590000);</script></html>`
