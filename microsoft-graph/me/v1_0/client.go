package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/me"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meactivityhistoryitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meactivityhistoryitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meagreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationemailmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationfido2method"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationpasswordmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationphonemethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationsoftwareoathmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationtemporaryaccesspassmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationwindowshelloforbusinessmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meauthenticationwindowshelloforbusinessmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendargroupcalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatlastmessagepreview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatpermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatpinnedmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechatpinnedmessagemessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechattab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mechattabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfolderchildfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfolderchildfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfolderchildfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mecreatedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/medevicemanagementtroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/medirectreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/medrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meemployeeexperience"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meemployeeexperiencelearningcourseactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meeventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meeventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meeventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meeventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meeventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meeventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meeventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mefollowedsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinferenceclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinferenceclassificationoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsightshared"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsightsharedlastsharedmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsightsharedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsighttrending"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsighttrendingresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsightused"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meinsightusedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamallchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamchanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamincomingchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteaminstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteaminstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteaminstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteammember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteampermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamprimarychanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduleoffershiftrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduleopenshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduleopenshiftchangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduleschedulinggroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduleshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduleswapshiftschangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduletimeoffreason"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduletimeoffrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamscheduletimesoff"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamtag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamtagmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mejoinedteamtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/melicensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfolderchildfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfolderchildfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfolderchildfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfolderchildfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memailfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanagedappregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memanager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/memessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meoauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotenotebooksectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenoteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotepagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotepageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotepageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenoteresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenoteresourcecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonenotesectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonlinemeeting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonlinemeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonlinemeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meonlinemeetingattendeereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meoutlook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meoutlookmastercategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meowneddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mepeople"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mephoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplanner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannertaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannertaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannertaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meplannertaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mepresence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meregistereddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mescopedrolememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/mesettingshiftpreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meteamwork"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meteamworkassociatedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meteamworkassociatedteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meteamworkinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meteamworkinstalledappchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meteamworkinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/meteamworkinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolistextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolisttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolisttaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolisttaskattachmentsession"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolisttaskattachmentsessioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolisttaskchecklistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolisttaskextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metodolisttasklinkedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/v1_0/metransitivememberof"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Me                                                        *me.MeClient
	MeActivity                                                *meactivity.MeActivityClient
	MeActivityHistoryItem                                     *meactivityhistoryitem.MeActivityHistoryItemClient
	MeActivityHistoryItemActivity                             *meactivityhistoryitemactivity.MeActivityHistoryItemActivityClient
	MeAgreementAcceptance                                     *meagreementacceptance.MeAgreementAcceptanceClient
	MeAppRoleAssignment                                       *meapproleassignment.MeAppRoleAssignmentClient
	MeAuthentication                                          *meauthentication.MeAuthenticationClient
	MeAuthenticationEmailMethod                               *meauthenticationemailmethod.MeAuthenticationEmailMethodClient
	MeAuthenticationFido2Method                               *meauthenticationfido2method.MeAuthenticationFido2MethodClient
	MeAuthenticationMethod                                    *meauthenticationmethod.MeAuthenticationMethodClient
	MeAuthenticationMicrosoftAuthenticatorMethod              *meauthenticationmicrosoftauthenticatormethod.MeAuthenticationMicrosoftAuthenticatorMethodClient
	MeAuthenticationMicrosoftAuthenticatorMethodDevice        *meauthenticationmicrosoftauthenticatormethoddevice.MeAuthenticationMicrosoftAuthenticatorMethodDeviceClient
	MeAuthenticationOperation                                 *meauthenticationoperation.MeAuthenticationOperationClient
	MeAuthenticationPasswordMethod                            *meauthenticationpasswordmethod.MeAuthenticationPasswordMethodClient
	MeAuthenticationPhoneMethod                               *meauthenticationphonemethod.MeAuthenticationPhoneMethodClient
	MeAuthenticationSoftwareOathMethod                        *meauthenticationsoftwareoathmethod.MeAuthenticationSoftwareOathMethodClient
	MeAuthenticationTemporaryAccessPassMethod                 *meauthenticationtemporaryaccesspassmethod.MeAuthenticationTemporaryAccessPassMethodClient
	MeAuthenticationWindowsHelloForBusinessMethod             *meauthenticationwindowshelloforbusinessmethod.MeAuthenticationWindowsHelloForBusinessMethodClient
	MeAuthenticationWindowsHelloForBusinessMethodDevice       *meauthenticationwindowshelloforbusinessmethoddevice.MeAuthenticationWindowsHelloForBusinessMethodDeviceClient
	MeCalendar                                                *mecalendar.MeCalendarClient
	MeCalendarCalendarPermission                              *mecalendarcalendarpermission.MeCalendarCalendarPermissionClient
	MeCalendarCalendarView                                    *mecalendarcalendarview.MeCalendarCalendarViewClient
	MeCalendarCalendarViewAttachment                          *mecalendarcalendarviewattachment.MeCalendarCalendarViewAttachmentClient
	MeCalendarCalendarViewCalendar                            *mecalendarcalendarviewcalendar.MeCalendarCalendarViewCalendarClient
	MeCalendarCalendarViewExtension                           *mecalendarcalendarviewextension.MeCalendarCalendarViewExtensionClient
	MeCalendarCalendarViewInstance                            *mecalendarcalendarviewinstance.MeCalendarCalendarViewInstanceClient
	MeCalendarCalendarViewInstanceAttachment                  *mecalendarcalendarviewinstanceattachment.MeCalendarCalendarViewInstanceAttachmentClient
	MeCalendarCalendarViewInstanceCalendar                    *mecalendarcalendarviewinstancecalendar.MeCalendarCalendarViewInstanceCalendarClient
	MeCalendarCalendarViewInstanceExtension                   *mecalendarcalendarviewinstanceextension.MeCalendarCalendarViewInstanceExtensionClient
	MeCalendarEvent                                           *mecalendarevent.MeCalendarEventClient
	MeCalendarEventAttachment                                 *mecalendareventattachment.MeCalendarEventAttachmentClient
	MeCalendarEventCalendar                                   *mecalendareventcalendar.MeCalendarEventCalendarClient
	MeCalendarEventExtension                                  *mecalendareventextension.MeCalendarEventExtensionClient
	MeCalendarEventInstance                                   *mecalendareventinstance.MeCalendarEventInstanceClient
	MeCalendarEventInstanceAttachment                         *mecalendareventinstanceattachment.MeCalendarEventInstanceAttachmentClient
	MeCalendarEventInstanceCalendar                           *mecalendareventinstancecalendar.MeCalendarEventInstanceCalendarClient
	MeCalendarEventInstanceExtension                          *mecalendareventinstanceextension.MeCalendarEventInstanceExtensionClient
	MeCalendarGroup                                           *mecalendargroup.MeCalendarGroupClient
	MeCalendarGroupCalendar                                   *mecalendargroupcalendar.MeCalendarGroupCalendarClient
	MeCalendarGroupCalendarCalendarPermission                 *mecalendargroupcalendarcalendarpermission.MeCalendarGroupCalendarCalendarPermissionClient
	MeCalendarGroupCalendarCalendarView                       *mecalendargroupcalendarcalendarview.MeCalendarGroupCalendarCalendarViewClient
	MeCalendarGroupCalendarCalendarViewAttachment             *mecalendargroupcalendarcalendarviewattachment.MeCalendarGroupCalendarCalendarViewAttachmentClient
	MeCalendarGroupCalendarCalendarViewCalendar               *mecalendargroupcalendarcalendarviewcalendar.MeCalendarGroupCalendarCalendarViewCalendarClient
	MeCalendarGroupCalendarCalendarViewExtension              *mecalendargroupcalendarcalendarviewextension.MeCalendarGroupCalendarCalendarViewExtensionClient
	MeCalendarGroupCalendarCalendarViewInstance               *mecalendargroupcalendarcalendarviewinstance.MeCalendarGroupCalendarCalendarViewInstanceClient
	MeCalendarGroupCalendarCalendarViewInstanceAttachment     *mecalendargroupcalendarcalendarviewinstanceattachment.MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient
	MeCalendarGroupCalendarCalendarViewInstanceCalendar       *mecalendargroupcalendarcalendarviewinstancecalendar.MeCalendarGroupCalendarCalendarViewInstanceCalendarClient
	MeCalendarGroupCalendarCalendarViewInstanceExtension      *mecalendargroupcalendarcalendarviewinstanceextension.MeCalendarGroupCalendarCalendarViewInstanceExtensionClient
	MeCalendarGroupCalendarEvent                              *mecalendargroupcalendarevent.MeCalendarGroupCalendarEventClient
	MeCalendarGroupCalendarEventAttachment                    *mecalendargroupcalendareventattachment.MeCalendarGroupCalendarEventAttachmentClient
	MeCalendarGroupCalendarEventCalendar                      *mecalendargroupcalendareventcalendar.MeCalendarGroupCalendarEventCalendarClient
	MeCalendarGroupCalendarEventExtension                     *mecalendargroupcalendareventextension.MeCalendarGroupCalendarEventExtensionClient
	MeCalendarGroupCalendarEventInstance                      *mecalendargroupcalendareventinstance.MeCalendarGroupCalendarEventInstanceClient
	MeCalendarGroupCalendarEventInstanceAttachment            *mecalendargroupcalendareventinstanceattachment.MeCalendarGroupCalendarEventInstanceAttachmentClient
	MeCalendarGroupCalendarEventInstanceCalendar              *mecalendargroupcalendareventinstancecalendar.MeCalendarGroupCalendarEventInstanceCalendarClient
	MeCalendarGroupCalendarEventInstanceExtension             *mecalendargroupcalendareventinstanceextension.MeCalendarGroupCalendarEventInstanceExtensionClient
	MeCalendarView                                            *mecalendarview.MeCalendarViewClient
	MeCalendarViewAttachment                                  *mecalendarviewattachment.MeCalendarViewAttachmentClient
	MeCalendarViewCalendar                                    *mecalendarviewcalendar.MeCalendarViewCalendarClient
	MeCalendarViewExtension                                   *mecalendarviewextension.MeCalendarViewExtensionClient
	MeCalendarViewInstance                                    *mecalendarviewinstance.MeCalendarViewInstanceClient
	MeCalendarViewInstanceAttachment                          *mecalendarviewinstanceattachment.MeCalendarViewInstanceAttachmentClient
	MeCalendarViewInstanceCalendar                            *mecalendarviewinstancecalendar.MeCalendarViewInstanceCalendarClient
	MeCalendarViewInstanceExtension                           *mecalendarviewinstanceextension.MeCalendarViewInstanceExtensionClient
	MeChat                                                    *mechat.MeChatClient
	MeChatInstalledApp                                        *mechatinstalledapp.MeChatInstalledAppClient
	MeChatInstalledAppTeamsApp                                *mechatinstalledappteamsapp.MeChatInstalledAppTeamsAppClient
	MeChatInstalledAppTeamsAppDefinition                      *mechatinstalledappteamsappdefinition.MeChatInstalledAppTeamsAppDefinitionClient
	MeChatLastMessagePreview                                  *mechatlastmessagepreview.MeChatLastMessagePreviewClient
	MeChatMember                                              *mechatmember.MeChatMemberClient
	MeChatMessage                                             *mechatmessage.MeChatMessageClient
	MeChatMessageHostedContent                                *mechatmessagehostedcontent.MeChatMessageHostedContentClient
	MeChatMessageReply                                        *mechatmessagereply.MeChatMessageReplyClient
	MeChatMessageReplyHostedContent                           *mechatmessagereplyhostedcontent.MeChatMessageReplyHostedContentClient
	MeChatPermissionGrant                                     *mechatpermissiongrant.MeChatPermissionGrantClient
	MeChatPinnedMessage                                       *mechatpinnedmessage.MeChatPinnedMessageClient
	MeChatPinnedMessageMessage                                *mechatpinnedmessagemessage.MeChatPinnedMessageMessageClient
	MeChatTab                                                 *mechattab.MeChatTabClient
	MeChatTabTeamsApp                                         *mechattabteamsapp.MeChatTabTeamsAppClient
	MeContact                                                 *mecontact.MeContactClient
	MeContactExtension                                        *mecontactextension.MeContactExtensionClient
	MeContactFolder                                           *mecontactfolder.MeContactFolderClient
	MeContactFolderChildFolder                                *mecontactfolderchildfolder.MeContactFolderChildFolderClient
	MeContactFolderChildFolderContact                         *mecontactfolderchildfoldercontact.MeContactFolderChildFolderContactClient
	MeContactFolderChildFolderContactExtension                *mecontactfolderchildfoldercontactextension.MeContactFolderChildFolderContactExtensionClient
	MeContactFolderChildFolderContactPhoto                    *mecontactfolderchildfoldercontactphoto.MeContactFolderChildFolderContactPhotoClient
	MeContactFolderContact                                    *mecontactfoldercontact.MeContactFolderContactClient
	MeContactFolderContactExtension                           *mecontactfoldercontactextension.MeContactFolderContactExtensionClient
	MeContactFolderContactPhoto                               *mecontactfoldercontactphoto.MeContactFolderContactPhotoClient
	MeContactPhoto                                            *mecontactphoto.MeContactPhotoClient
	MeCreatedObject                                           *mecreatedobject.MeCreatedObjectClient
	MeDeviceManagementTroubleshootingEvent                    *medevicemanagementtroubleshootingevent.MeDeviceManagementTroubleshootingEventClient
	MeDirectReport                                            *medirectreport.MeDirectReportClient
	MeDrive                                                   *medrive.MeDriveClient
	MeEmployeeExperience                                      *meemployeeexperience.MeEmployeeExperienceClient
	MeEmployeeExperienceLearningCourseActivity                *meemployeeexperiencelearningcourseactivity.MeEmployeeExperienceLearningCourseActivityClient
	MeEvent                                                   *meevent.MeEventClient
	MeEventAttachment                                         *meeventattachment.MeEventAttachmentClient
	MeEventCalendar                                           *meeventcalendar.MeEventCalendarClient
	MeEventExtension                                          *meeventextension.MeEventExtensionClient
	MeEventInstance                                           *meeventinstance.MeEventInstanceClient
	MeEventInstanceAttachment                                 *meeventinstanceattachment.MeEventInstanceAttachmentClient
	MeEventInstanceCalendar                                   *meeventinstancecalendar.MeEventInstanceCalendarClient
	MeEventInstanceExtension                                  *meeventinstanceextension.MeEventInstanceExtensionClient
	MeExtension                                               *meextension.MeExtensionClient
	MeFollowedSite                                            *mefollowedsite.MeFollowedSiteClient
	MeInferenceClassification                                 *meinferenceclassification.MeInferenceClassificationClient
	MeInferenceClassificationOverride                         *meinferenceclassificationoverride.MeInferenceClassificationOverrideClient
	MeInsight                                                 *meinsight.MeInsightClient
	MeInsightShared                                           *meinsightshared.MeInsightSharedClient
	MeInsightSharedLastSharedMethod                           *meinsightsharedlastsharedmethod.MeInsightSharedLastSharedMethodClient
	MeInsightSharedResource                                   *meinsightsharedresource.MeInsightSharedResourceClient
	MeInsightTrending                                         *meinsighttrending.MeInsightTrendingClient
	MeInsightTrendingResource                                 *meinsighttrendingresource.MeInsightTrendingResourceClient
	MeInsightUsed                                             *meinsightused.MeInsightUsedClient
	MeInsightUsedResource                                     *meinsightusedresource.MeInsightUsedResourceClient
	MeJoinedTeam                                              *mejoinedteam.MeJoinedTeamClient
	MeJoinedTeamAllChannel                                    *mejoinedteamallchannel.MeJoinedTeamAllChannelClient
	MeJoinedTeamChannel                                       *mejoinedteamchannel.MeJoinedTeamChannelClient
	MeJoinedTeamChannelFilesFolder                            *mejoinedteamchannelfilesfolder.MeJoinedTeamChannelFilesFolderClient
	MeJoinedTeamChannelFilesFolderContent                     *mejoinedteamchannelfilesfoldercontent.MeJoinedTeamChannelFilesFolderContentClient
	MeJoinedTeamChannelMember                                 *mejoinedteamchannelmember.MeJoinedTeamChannelMemberClient
	MeJoinedTeamChannelMessage                                *mejoinedteamchannelmessage.MeJoinedTeamChannelMessageClient
	MeJoinedTeamChannelMessageHostedContent                   *mejoinedteamchannelmessagehostedcontent.MeJoinedTeamChannelMessageHostedContentClient
	MeJoinedTeamChannelMessageReply                           *mejoinedteamchannelmessagereply.MeJoinedTeamChannelMessageReplyClient
	MeJoinedTeamChannelMessageReplyHostedContent              *mejoinedteamchannelmessagereplyhostedcontent.MeJoinedTeamChannelMessageReplyHostedContentClient
	MeJoinedTeamChannelSharedWithTeam                         *mejoinedteamchannelsharedwithteam.MeJoinedTeamChannelSharedWithTeamClient
	MeJoinedTeamChannelSharedWithTeamAllowedMember            *mejoinedteamchannelsharedwithteamallowedmember.MeJoinedTeamChannelSharedWithTeamAllowedMemberClient
	MeJoinedTeamChannelSharedWithTeamTeam                     *mejoinedteamchannelsharedwithteamteam.MeJoinedTeamChannelSharedWithTeamTeamClient
	MeJoinedTeamChannelTab                                    *mejoinedteamchanneltab.MeJoinedTeamChannelTabClient
	MeJoinedTeamChannelTabTeamsApp                            *mejoinedteamchanneltabteamsapp.MeJoinedTeamChannelTabTeamsAppClient
	MeJoinedTeamGroup                                         *mejoinedteamgroup.MeJoinedTeamGroupClient
	MeJoinedTeamIncomingChannel                               *mejoinedteamincomingchannel.MeJoinedTeamIncomingChannelClient
	MeJoinedTeamInstalledApp                                  *mejoinedteaminstalledapp.MeJoinedTeamInstalledAppClient
	MeJoinedTeamInstalledAppTeamsApp                          *mejoinedteaminstalledappteamsapp.MeJoinedTeamInstalledAppTeamsAppClient
	MeJoinedTeamInstalledAppTeamsAppDefinition                *mejoinedteaminstalledappteamsappdefinition.MeJoinedTeamInstalledAppTeamsAppDefinitionClient
	MeJoinedTeamMember                                        *mejoinedteammember.MeJoinedTeamMemberClient
	MeJoinedTeamOperation                                     *mejoinedteamoperation.MeJoinedTeamOperationClient
	MeJoinedTeamPermissionGrant                               *mejoinedteampermissiongrant.MeJoinedTeamPermissionGrantClient
	MeJoinedTeamPhoto                                         *mejoinedteamphoto.MeJoinedTeamPhotoClient
	MeJoinedTeamPrimaryChannel                                *mejoinedteamprimarychannel.MeJoinedTeamPrimaryChannelClient
	MeJoinedTeamPrimaryChannelFilesFolder                     *mejoinedteamprimarychannelfilesfolder.MeJoinedTeamPrimaryChannelFilesFolderClient
	MeJoinedTeamPrimaryChannelFilesFolderContent              *mejoinedteamprimarychannelfilesfoldercontent.MeJoinedTeamPrimaryChannelFilesFolderContentClient
	MeJoinedTeamPrimaryChannelMember                          *mejoinedteamprimarychannelmember.MeJoinedTeamPrimaryChannelMemberClient
	MeJoinedTeamPrimaryChannelMessage                         *mejoinedteamprimarychannelmessage.MeJoinedTeamPrimaryChannelMessageClient
	MeJoinedTeamPrimaryChannelMessageHostedContent            *mejoinedteamprimarychannelmessagehostedcontent.MeJoinedTeamPrimaryChannelMessageHostedContentClient
	MeJoinedTeamPrimaryChannelMessageReply                    *mejoinedteamprimarychannelmessagereply.MeJoinedTeamPrimaryChannelMessageReplyClient
	MeJoinedTeamPrimaryChannelMessageReplyHostedContent       *mejoinedteamprimarychannelmessagereplyhostedcontent.MeJoinedTeamPrimaryChannelMessageReplyHostedContentClient
	MeJoinedTeamPrimaryChannelSharedWithTeam                  *mejoinedteamprimarychannelsharedwithteam.MeJoinedTeamPrimaryChannelSharedWithTeamClient
	MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMember     *mejoinedteamprimarychannelsharedwithteamallowedmember.MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient
	MeJoinedTeamPrimaryChannelSharedWithTeamTeam              *mejoinedteamprimarychannelsharedwithteamteam.MeJoinedTeamPrimaryChannelSharedWithTeamTeamClient
	MeJoinedTeamPrimaryChannelTab                             *mejoinedteamprimarychanneltab.MeJoinedTeamPrimaryChannelTabClient
	MeJoinedTeamPrimaryChannelTabTeamsApp                     *mejoinedteamprimarychanneltabteamsapp.MeJoinedTeamPrimaryChannelTabTeamsAppClient
	MeJoinedTeamSchedule                                      *mejoinedteamschedule.MeJoinedTeamScheduleClient
	MeJoinedTeamScheduleOfferShiftRequest                     *mejoinedteamscheduleoffershiftrequest.MeJoinedTeamScheduleOfferShiftRequestClient
	MeJoinedTeamScheduleOpenShift                             *mejoinedteamscheduleopenshift.MeJoinedTeamScheduleOpenShiftClient
	MeJoinedTeamScheduleOpenShiftChangeRequest                *mejoinedteamscheduleopenshiftchangerequest.MeJoinedTeamScheduleOpenShiftChangeRequestClient
	MeJoinedTeamScheduleSchedulingGroup                       *mejoinedteamscheduleschedulinggroup.MeJoinedTeamScheduleSchedulingGroupClient
	MeJoinedTeamScheduleShift                                 *mejoinedteamscheduleshift.MeJoinedTeamScheduleShiftClient
	MeJoinedTeamScheduleSwapShiftsChangeRequest               *mejoinedteamscheduleswapshiftschangerequest.MeJoinedTeamScheduleSwapShiftsChangeRequestClient
	MeJoinedTeamScheduleTimeOffReason                         *mejoinedteamscheduletimeoffreason.MeJoinedTeamScheduleTimeOffReasonClient
	MeJoinedTeamScheduleTimeOffRequest                        *mejoinedteamscheduletimeoffrequest.MeJoinedTeamScheduleTimeOffRequestClient
	MeJoinedTeamScheduleTimesOff                              *mejoinedteamscheduletimesoff.MeJoinedTeamScheduleTimesOffClient
	MeJoinedTeamTag                                           *mejoinedteamtag.MeJoinedTeamTagClient
	MeJoinedTeamTagMember                                     *mejoinedteamtagmember.MeJoinedTeamTagMemberClient
	MeJoinedTeamTemplate                                      *mejoinedteamtemplate.MeJoinedTeamTemplateClient
	MeLicenseDetail                                           *melicensedetail.MeLicenseDetailClient
	MeMailFolder                                              *memailfolder.MeMailFolderClient
	MeMailFolderChildFolder                                   *memailfolderchildfolder.MeMailFolderChildFolderClient
	MeMailFolderChildFolderMessage                            *memailfolderchildfoldermessage.MeMailFolderChildFolderMessageClient
	MeMailFolderChildFolderMessageAttachment                  *memailfolderchildfoldermessageattachment.MeMailFolderChildFolderMessageAttachmentClient
	MeMailFolderChildFolderMessageExtension                   *memailfolderchildfoldermessageextension.MeMailFolderChildFolderMessageExtensionClient
	MeMailFolderChildFolderMessageRule                        *memailfolderchildfoldermessagerule.MeMailFolderChildFolderMessageRuleClient
	MeMailFolderMessage                                       *memailfoldermessage.MeMailFolderMessageClient
	MeMailFolderMessageAttachment                             *memailfoldermessageattachment.MeMailFolderMessageAttachmentClient
	MeMailFolderMessageExtension                              *memailfoldermessageextension.MeMailFolderMessageExtensionClient
	MeMailFolderMessageRule                                   *memailfoldermessagerule.MeMailFolderMessageRuleClient
	MeMailboxSetting                                          *memailboxsetting.MeMailboxSettingClient
	MeManagedAppRegistration                                  *memanagedappregistration.MeManagedAppRegistrationClient
	MeManagedDevice                                           *memanageddevice.MeManagedDeviceClient
	MeManagedDeviceDeviceCategory                             *memanageddevicedevicecategory.MeManagedDeviceDeviceCategoryClient
	MeManagedDeviceDeviceCompliancePolicyState                *memanageddevicedevicecompliancepolicystate.MeManagedDeviceDeviceCompliancePolicyStateClient
	MeManagedDeviceDeviceConfigurationState                   *memanageddevicedeviceconfigurationstate.MeManagedDeviceDeviceConfigurationStateClient
	MeManagedDeviceLogCollectionRequest                       *memanageddevicelogcollectionrequest.MeManagedDeviceLogCollectionRequestClient
	MeManagedDeviceUser                                       *memanageddeviceuser.MeManagedDeviceUserClient
	MeManagedDeviceWindowsProtectionState                     *memanageddevicewindowsprotectionstate.MeManagedDeviceWindowsProtectionStateClient
	MeManagedDeviceWindowsProtectionStateDetectedMalwareState *memanageddevicewindowsprotectionstatedetectedmalwarestate.MeManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	MeManager                                                 *memanager.MeManagerClient
	MeMemberOf                                                *mememberof.MeMemberOfClient
	MeMessage                                                 *memessage.MeMessageClient
	MeMessageAttachment                                       *memessageattachment.MeMessageAttachmentClient
	MeMessageExtension                                        *memessageextension.MeMessageExtensionClient
	MeOauth2PermissionGrant                                   *meoauth2permissiongrant.MeOauth2PermissionGrantClient
	MeOnenote                                                 *meonenote.MeOnenoteClient
	MeOnenoteNotebook                                         *meonenotenotebook.MeOnenoteNotebookClient
	MeOnenoteNotebookSection                                  *meonenotenotebooksection.MeOnenoteNotebookSectionClient
	MeOnenoteNotebookSectionGroup                             *meonenotenotebooksectiongroup.MeOnenoteNotebookSectionGroupClient
	MeOnenoteNotebookSectionGroupParentNotebook               *meonenotenotebooksectiongroupparentnotebook.MeOnenoteNotebookSectionGroupParentNotebookClient
	MeOnenoteNotebookSectionGroupParentSectionGroup           *meonenotenotebooksectiongroupparentsectiongroup.MeOnenoteNotebookSectionGroupParentSectionGroupClient
	MeOnenoteNotebookSectionGroupSection                      *meonenotenotebooksectiongroupsection.MeOnenoteNotebookSectionGroupSectionClient
	MeOnenoteNotebookSectionGroupSectionGroup                 *meonenotenotebooksectiongroupsectiongroup.MeOnenoteNotebookSectionGroupSectionGroupClient
	MeOnenoteNotebookSectionGroupSectionPage                  *meonenotenotebooksectiongroupsectionpage.MeOnenoteNotebookSectionGroupSectionPageClient
	MeOnenoteNotebookSectionGroupSectionPageContent           *meonenotenotebooksectiongroupsectionpagecontent.MeOnenoteNotebookSectionGroupSectionPageContentClient
	MeOnenoteNotebookSectionGroupSectionPageParentNotebook    *meonenotenotebooksectiongroupsectionpageparentnotebook.MeOnenoteNotebookSectionGroupSectionPageParentNotebookClient
	MeOnenoteNotebookSectionGroupSectionPageParentSection     *meonenotenotebooksectiongroupsectionpageparentsection.MeOnenoteNotebookSectionGroupSectionPageParentSectionClient
	MeOnenoteNotebookSectionGroupSectionParentNotebook        *meonenotenotebooksectiongroupsectionparentnotebook.MeOnenoteNotebookSectionGroupSectionParentNotebookClient
	MeOnenoteNotebookSectionGroupSectionParentSectionGroup    *meonenotenotebooksectiongroupsectionparentsectiongroup.MeOnenoteNotebookSectionGroupSectionParentSectionGroupClient
	MeOnenoteNotebookSectionPage                              *meonenotenotebooksectionpage.MeOnenoteNotebookSectionPageClient
	MeOnenoteNotebookSectionPageContent                       *meonenotenotebooksectionpagecontent.MeOnenoteNotebookSectionPageContentClient
	MeOnenoteNotebookSectionPageParentNotebook                *meonenotenotebooksectionpageparentnotebook.MeOnenoteNotebookSectionPageParentNotebookClient
	MeOnenoteNotebookSectionPageParentSection                 *meonenotenotebooksectionpageparentsection.MeOnenoteNotebookSectionPageParentSectionClient
	MeOnenoteNotebookSectionParentNotebook                    *meonenotenotebooksectionparentnotebook.MeOnenoteNotebookSectionParentNotebookClient
	MeOnenoteNotebookSectionParentSectionGroup                *meonenotenotebooksectionparentsectiongroup.MeOnenoteNotebookSectionParentSectionGroupClient
	MeOnenoteOperation                                        *meonenoteoperation.MeOnenoteOperationClient
	MeOnenotePage                                             *meonenotepage.MeOnenotePageClient
	MeOnenotePageContent                                      *meonenotepagecontent.MeOnenotePageContentClient
	MeOnenotePageParentNotebook                               *meonenotepageparentnotebook.MeOnenotePageParentNotebookClient
	MeOnenotePageParentSection                                *meonenotepageparentsection.MeOnenotePageParentSectionClient
	MeOnenoteResource                                         *meonenoteresource.MeOnenoteResourceClient
	MeOnenoteResourceContent                                  *meonenoteresourcecontent.MeOnenoteResourceContentClient
	MeOnenoteSection                                          *meonenotesection.MeOnenoteSectionClient
	MeOnenoteSectionGroup                                     *meonenotesectiongroup.MeOnenoteSectionGroupClient
	MeOnenoteSectionGroupParentNotebook                       *meonenotesectiongroupparentnotebook.MeOnenoteSectionGroupParentNotebookClient
	MeOnenoteSectionGroupParentSectionGroup                   *meonenotesectiongroupparentsectiongroup.MeOnenoteSectionGroupParentSectionGroupClient
	MeOnenoteSectionGroupSection                              *meonenotesectiongroupsection.MeOnenoteSectionGroupSectionClient
	MeOnenoteSectionGroupSectionGroup                         *meonenotesectiongroupsectiongroup.MeOnenoteSectionGroupSectionGroupClient
	MeOnenoteSectionGroupSectionPage                          *meonenotesectiongroupsectionpage.MeOnenoteSectionGroupSectionPageClient
	MeOnenoteSectionGroupSectionPageContent                   *meonenotesectiongroupsectionpagecontent.MeOnenoteSectionGroupSectionPageContentClient
	MeOnenoteSectionGroupSectionPageParentNotebook            *meonenotesectiongroupsectionpageparentnotebook.MeOnenoteSectionGroupSectionPageParentNotebookClient
	MeOnenoteSectionGroupSectionPageParentSection             *meonenotesectiongroupsectionpageparentsection.MeOnenoteSectionGroupSectionPageParentSectionClient
	MeOnenoteSectionGroupSectionParentNotebook                *meonenotesectiongroupsectionparentnotebook.MeOnenoteSectionGroupSectionParentNotebookClient
	MeOnenoteSectionGroupSectionParentSectionGroup            *meonenotesectiongroupsectionparentsectiongroup.MeOnenoteSectionGroupSectionParentSectionGroupClient
	MeOnenoteSectionPage                                      *meonenotesectionpage.MeOnenoteSectionPageClient
	MeOnenoteSectionPageContent                               *meonenotesectionpagecontent.MeOnenoteSectionPageContentClient
	MeOnenoteSectionPageParentNotebook                        *meonenotesectionpageparentnotebook.MeOnenoteSectionPageParentNotebookClient
	MeOnenoteSectionPageParentSection                         *meonenotesectionpageparentsection.MeOnenoteSectionPageParentSectionClient
	MeOnenoteSectionParentNotebook                            *meonenotesectionparentnotebook.MeOnenoteSectionParentNotebookClient
	MeOnenoteSectionParentSectionGroup                        *meonenotesectionparentsectiongroup.MeOnenoteSectionParentSectionGroupClient
	MeOnlineMeeting                                           *meonlinemeeting.MeOnlineMeetingClient
	MeOnlineMeetingAttendanceReport                           *meonlinemeetingattendancereport.MeOnlineMeetingAttendanceReportClient
	MeOnlineMeetingAttendanceReportAttendanceRecord           *meonlinemeetingattendancereportattendancerecord.MeOnlineMeetingAttendanceReportAttendanceRecordClient
	MeOnlineMeetingAttendeeReport                             *meonlinemeetingattendeereport.MeOnlineMeetingAttendeeReportClient
	MeOutlook                                                 *meoutlook.MeOutlookClient
	MeOutlookMasterCategory                                   *meoutlookmastercategory.MeOutlookMasterCategoryClient
	MeOwnedDevice                                             *meowneddevice.MeOwnedDeviceClient
	MeOwnedObject                                             *meownedobject.MeOwnedObjectClient
	MePeople                                                  *mepeople.MePeopleClient
	MePhoto                                                   *mephoto.MePhotoClient
	MePlanner                                                 *meplanner.MePlannerClient
	MePlannerPlan                                             *meplannerplan.MePlannerPlanClient
	MePlannerPlanBucket                                       *meplannerplanbucket.MePlannerPlanBucketClient
	MePlannerPlanBucketTask                                   *meplannerplanbuckettask.MePlannerPlanBucketTaskClient
	MePlannerPlanBucketTaskAssignedToTaskBoardFormat          *meplannerplanbuckettaskassignedtotaskboardformat.MePlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	MePlannerPlanBucketTaskBucketTaskBoardFormat              *meplannerplanbuckettaskbuckettaskboardformat.MePlannerPlanBucketTaskBucketTaskBoardFormatClient
	MePlannerPlanBucketTaskDetail                             *meplannerplanbuckettaskdetail.MePlannerPlanBucketTaskDetailClient
	MePlannerPlanBucketTaskProgressTaskBoardFormat            *meplannerplanbuckettaskprogresstaskboardformat.MePlannerPlanBucketTaskProgressTaskBoardFormatClient
	MePlannerPlanDetail                                       *meplannerplandetail.MePlannerPlanDetailClient
	MePlannerPlanTask                                         *meplannerplantask.MePlannerPlanTaskClient
	MePlannerPlanTaskAssignedToTaskBoardFormat                *meplannerplantaskassignedtotaskboardformat.MePlannerPlanTaskAssignedToTaskBoardFormatClient
	MePlannerPlanTaskBucketTaskBoardFormat                    *meplannerplantaskbuckettaskboardformat.MePlannerPlanTaskBucketTaskBoardFormatClient
	MePlannerPlanTaskDetail                                   *meplannerplantaskdetail.MePlannerPlanTaskDetailClient
	MePlannerPlanTaskProgressTaskBoardFormat                  *meplannerplantaskprogresstaskboardformat.MePlannerPlanTaskProgressTaskBoardFormatClient
	MePlannerTask                                             *meplannertask.MePlannerTaskClient
	MePlannerTaskAssignedToTaskBoardFormat                    *meplannertaskassignedtotaskboardformat.MePlannerTaskAssignedToTaskBoardFormatClient
	MePlannerTaskBucketTaskBoardFormat                        *meplannertaskbuckettaskboardformat.MePlannerTaskBucketTaskBoardFormatClient
	MePlannerTaskDetail                                       *meplannertaskdetail.MePlannerTaskDetailClient
	MePlannerTaskProgressTaskBoardFormat                      *meplannertaskprogresstaskboardformat.MePlannerTaskProgressTaskBoardFormatClient
	MePresence                                                *mepresence.MePresenceClient
	MeRegisteredDevice                                        *meregistereddevice.MeRegisteredDeviceClient
	MeScopedRoleMemberOf                                      *mescopedrolememberof.MeScopedRoleMemberOfClient
	MeSetting                                                 *mesetting.MeSettingClient
	MeSettingShiftPreference                                  *mesettingshiftpreference.MeSettingShiftPreferenceClient
	MeTeamwork                                                *meteamwork.MeTeamworkClient
	MeTeamworkAssociatedTeam                                  *meteamworkassociatedteam.MeTeamworkAssociatedTeamClient
	MeTeamworkAssociatedTeamTeam                              *meteamworkassociatedteamteam.MeTeamworkAssociatedTeamTeamClient
	MeTeamworkInstalledApp                                    *meteamworkinstalledapp.MeTeamworkInstalledAppClient
	MeTeamworkInstalledAppChat                                *meteamworkinstalledappchat.MeTeamworkInstalledAppChatClient
	MeTeamworkInstalledAppTeamsApp                            *meteamworkinstalledappteamsapp.MeTeamworkInstalledAppTeamsAppClient
	MeTeamworkInstalledAppTeamsAppDefinition                  *meteamworkinstalledappteamsappdefinition.MeTeamworkInstalledAppTeamsAppDefinitionClient
	MeTodo                                                    *metodo.MeTodoClient
	MeTodoList                                                *metodolist.MeTodoListClient
	MeTodoListExtension                                       *metodolistextension.MeTodoListExtensionClient
	MeTodoListTask                                            *metodolisttask.MeTodoListTaskClient
	MeTodoListTaskAttachment                                  *metodolisttaskattachment.MeTodoListTaskAttachmentClient
	MeTodoListTaskAttachmentSession                           *metodolisttaskattachmentsession.MeTodoListTaskAttachmentSessionClient
	MeTodoListTaskAttachmentSessionContent                    *metodolisttaskattachmentsessioncontent.MeTodoListTaskAttachmentSessionContentClient
	MeTodoListTaskChecklistItem                               *metodolisttaskchecklistitem.MeTodoListTaskChecklistItemClient
	MeTodoListTaskExtension                                   *metodolisttaskextension.MeTodoListTaskExtensionClient
	MeTodoListTaskLinkedResource                              *metodolisttasklinkedresource.MeTodoListTaskLinkedResourceClient
	MeTransitiveMemberOf                                      *metransitivememberof.MeTransitiveMemberOfClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	meActivityClient, err := meactivity.NewMeActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeActivity client: %+v", err)
	}
	configureFunc(meActivityClient.Client)

	meActivityHistoryItemActivityClient, err := meactivityhistoryitemactivity.NewMeActivityHistoryItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeActivityHistoryItemActivity client: %+v", err)
	}
	configureFunc(meActivityHistoryItemActivityClient.Client)

	meActivityHistoryItemClient, err := meactivityhistoryitem.NewMeActivityHistoryItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeActivityHistoryItem client: %+v", err)
	}
	configureFunc(meActivityHistoryItemClient.Client)

	meAgreementAcceptanceClient, err := meagreementacceptance.NewMeAgreementAcceptanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAgreementAcceptance client: %+v", err)
	}
	configureFunc(meAgreementAcceptanceClient.Client)

	meAppRoleAssignmentClient, err := meapproleassignment.NewMeAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAppRoleAssignment client: %+v", err)
	}
	configureFunc(meAppRoleAssignmentClient.Client)

	meAuthenticationClient, err := meauthentication.NewMeAuthenticationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthentication client: %+v", err)
	}
	configureFunc(meAuthenticationClient.Client)

	meAuthenticationEmailMethodClient, err := meauthenticationemailmethod.NewMeAuthenticationEmailMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationEmailMethod client: %+v", err)
	}
	configureFunc(meAuthenticationEmailMethodClient.Client)

	meAuthenticationFido2MethodClient, err := meauthenticationfido2method.NewMeAuthenticationFido2MethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationFido2Method client: %+v", err)
	}
	configureFunc(meAuthenticationFido2MethodClient.Client)

	meAuthenticationMethodClient, err := meauthenticationmethod.NewMeAuthenticationMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationMethod client: %+v", err)
	}
	configureFunc(meAuthenticationMethodClient.Client)

	meAuthenticationMicrosoftAuthenticatorMethodClient, err := meauthenticationmicrosoftauthenticatormethod.NewMeAuthenticationMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationMicrosoftAuthenticatorMethod client: %+v", err)
	}
	configureFunc(meAuthenticationMicrosoftAuthenticatorMethodClient.Client)

	meAuthenticationMicrosoftAuthenticatorMethodDeviceClient, err := meauthenticationmicrosoftauthenticatormethoddevice.NewMeAuthenticationMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationMicrosoftAuthenticatorMethodDevice client: %+v", err)
	}
	configureFunc(meAuthenticationMicrosoftAuthenticatorMethodDeviceClient.Client)

	meAuthenticationOperationClient, err := meauthenticationoperation.NewMeAuthenticationOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationOperation client: %+v", err)
	}
	configureFunc(meAuthenticationOperationClient.Client)

	meAuthenticationPasswordMethodClient, err := meauthenticationpasswordmethod.NewMeAuthenticationPasswordMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationPasswordMethod client: %+v", err)
	}
	configureFunc(meAuthenticationPasswordMethodClient.Client)

	meAuthenticationPhoneMethodClient, err := meauthenticationphonemethod.NewMeAuthenticationPhoneMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationPhoneMethod client: %+v", err)
	}
	configureFunc(meAuthenticationPhoneMethodClient.Client)

	meAuthenticationSoftwareOathMethodClient, err := meauthenticationsoftwareoathmethod.NewMeAuthenticationSoftwareOathMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationSoftwareOathMethod client: %+v", err)
	}
	configureFunc(meAuthenticationSoftwareOathMethodClient.Client)

	meAuthenticationTemporaryAccessPassMethodClient, err := meauthenticationtemporaryaccesspassmethod.NewMeAuthenticationTemporaryAccessPassMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationTemporaryAccessPassMethod client: %+v", err)
	}
	configureFunc(meAuthenticationTemporaryAccessPassMethodClient.Client)

	meAuthenticationWindowsHelloForBusinessMethodClient, err := meauthenticationwindowshelloforbusinessmethod.NewMeAuthenticationWindowsHelloForBusinessMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationWindowsHelloForBusinessMethod client: %+v", err)
	}
	configureFunc(meAuthenticationWindowsHelloForBusinessMethodClient.Client)

	meAuthenticationWindowsHelloForBusinessMethodDeviceClient, err := meauthenticationwindowshelloforbusinessmethoddevice.NewMeAuthenticationWindowsHelloForBusinessMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationWindowsHelloForBusinessMethodDevice client: %+v", err)
	}
	configureFunc(meAuthenticationWindowsHelloForBusinessMethodDeviceClient.Client)

	meCalendarCalendarPermissionClient, err := mecalendarcalendarpermission.NewMeCalendarCalendarPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarPermission client: %+v", err)
	}
	configureFunc(meCalendarCalendarPermissionClient.Client)

	meCalendarCalendarViewAttachmentClient, err := mecalendarcalendarviewattachment.NewMeCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewAttachment client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewAttachmentClient.Client)

	meCalendarCalendarViewCalendarClient, err := mecalendarcalendarviewcalendar.NewMeCalendarCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewCalendar client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewCalendarClient.Client)

	meCalendarCalendarViewClient, err := mecalendarcalendarview.NewMeCalendarCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarView client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewClient.Client)

	meCalendarCalendarViewExtensionClient, err := mecalendarcalendarviewextension.NewMeCalendarCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExtension client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExtensionClient.Client)

	meCalendarCalendarViewInstanceAttachmentClient, err := mecalendarcalendarviewinstanceattachment.NewMeCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceAttachmentClient.Client)

	meCalendarCalendarViewInstanceCalendarClient, err := mecalendarcalendarviewinstancecalendar.NewMeCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceCalendarClient.Client)

	meCalendarCalendarViewInstanceClient, err := mecalendarcalendarviewinstance.NewMeCalendarCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstance client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceClient.Client)

	meCalendarCalendarViewInstanceExtensionClient, err := mecalendarcalendarviewinstanceextension.NewMeCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceExtensionClient.Client)

	meCalendarClient, err := mecalendar.NewMeCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendar client: %+v", err)
	}
	configureFunc(meCalendarClient.Client)

	meCalendarEventAttachmentClient, err := mecalendareventattachment.NewMeCalendarEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventAttachment client: %+v", err)
	}
	configureFunc(meCalendarEventAttachmentClient.Client)

	meCalendarEventCalendarClient, err := mecalendareventcalendar.NewMeCalendarEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventCalendar client: %+v", err)
	}
	configureFunc(meCalendarEventCalendarClient.Client)

	meCalendarEventClient, err := mecalendarevent.NewMeCalendarEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEvent client: %+v", err)
	}
	configureFunc(meCalendarEventClient.Client)

	meCalendarEventExtensionClient, err := mecalendareventextension.NewMeCalendarEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExtension client: %+v", err)
	}
	configureFunc(meCalendarEventExtensionClient.Client)

	meCalendarEventInstanceAttachmentClient, err := mecalendareventinstanceattachment.NewMeCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceAttachmentClient.Client)

	meCalendarEventInstanceCalendarClient, err := mecalendareventinstancecalendar.NewMeCalendarEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceCalendarClient.Client)

	meCalendarEventInstanceClient, err := mecalendareventinstance.NewMeCalendarEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstance client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceClient.Client)

	meCalendarEventInstanceExtensionClient, err := mecalendareventinstanceextension.NewMeCalendarEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceExtensionClient.Client)

	meCalendarGroupCalendarCalendarPermissionClient, err := mecalendargroupcalendarcalendarpermission.NewMeCalendarGroupCalendarCalendarPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarPermission client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarPermissionClient.Client)

	meCalendarGroupCalendarCalendarViewAttachmentClient, err := mecalendargroupcalendarcalendarviewattachment.NewMeCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewAttachmentClient.Client)

	meCalendarGroupCalendarCalendarViewCalendarClient, err := mecalendargroupcalendarcalendarviewcalendar.NewMeCalendarGroupCalendarCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewCalendarClient.Client)

	meCalendarGroupCalendarCalendarViewClient, err := mecalendargroupcalendarcalendarview.NewMeCalendarGroupCalendarCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarView client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewClient.Client)

	meCalendarGroupCalendarCalendarViewExtensionClient, err := mecalendargroupcalendarcalendarviewextension.NewMeCalendarGroupCalendarCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExtensionClient.Client)

	meCalendarGroupCalendarCalendarViewInstanceAttachmentClient, err := mecalendargroupcalendarcalendarviewinstanceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceAttachmentClient.Client)

	meCalendarGroupCalendarCalendarViewInstanceCalendarClient, err := mecalendargroupcalendarcalendarviewinstancecalendar.NewMeCalendarGroupCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceCalendarClient.Client)

	meCalendarGroupCalendarCalendarViewInstanceClient, err := mecalendargroupcalendarcalendarviewinstance.NewMeCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstance client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceClient.Client)

	meCalendarGroupCalendarCalendarViewInstanceExtensionClient, err := mecalendargroupcalendarcalendarviewinstanceextension.NewMeCalendarGroupCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceExtensionClient.Client)

	meCalendarGroupCalendarClient, err := mecalendargroupcalendar.NewMeCalendarGroupCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarClient.Client)

	meCalendarGroupCalendarEventAttachmentClient, err := mecalendargroupcalendareventattachment.NewMeCalendarGroupCalendarEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventAttachmentClient.Client)

	meCalendarGroupCalendarEventCalendarClient, err := mecalendargroupcalendareventcalendar.NewMeCalendarGroupCalendarEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventCalendarClient.Client)

	meCalendarGroupCalendarEventClient, err := mecalendargroupcalendarevent.NewMeCalendarGroupCalendarEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEvent client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventClient.Client)

	meCalendarGroupCalendarEventExtensionClient, err := mecalendargroupcalendareventextension.NewMeCalendarGroupCalendarEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExtensionClient.Client)

	meCalendarGroupCalendarEventInstanceAttachmentClient, err := mecalendargroupcalendareventinstanceattachment.NewMeCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceAttachmentClient.Client)

	meCalendarGroupCalendarEventInstanceCalendarClient, err := mecalendargroupcalendareventinstancecalendar.NewMeCalendarGroupCalendarEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceCalendarClient.Client)

	meCalendarGroupCalendarEventInstanceClient, err := mecalendargroupcalendareventinstance.NewMeCalendarGroupCalendarEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstance client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceClient.Client)

	meCalendarGroupCalendarEventInstanceExtensionClient, err := mecalendargroupcalendareventinstanceextension.NewMeCalendarGroupCalendarEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceExtensionClient.Client)

	meCalendarGroupClient, err := mecalendargroup.NewMeCalendarGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroup client: %+v", err)
	}
	configureFunc(meCalendarGroupClient.Client)

	meCalendarViewAttachmentClient, err := mecalendarviewattachment.NewMeCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewAttachment client: %+v", err)
	}
	configureFunc(meCalendarViewAttachmentClient.Client)

	meCalendarViewCalendarClient, err := mecalendarviewcalendar.NewMeCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewCalendar client: %+v", err)
	}
	configureFunc(meCalendarViewCalendarClient.Client)

	meCalendarViewClient, err := mecalendarview.NewMeCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarView client: %+v", err)
	}
	configureFunc(meCalendarViewClient.Client)

	meCalendarViewExtensionClient, err := mecalendarviewextension.NewMeCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExtension client: %+v", err)
	}
	configureFunc(meCalendarViewExtensionClient.Client)

	meCalendarViewInstanceAttachmentClient, err := mecalendarviewinstanceattachment.NewMeCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceAttachmentClient.Client)

	meCalendarViewInstanceCalendarClient, err := mecalendarviewinstancecalendar.NewMeCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceCalendarClient.Client)

	meCalendarViewInstanceClient, err := mecalendarviewinstance.NewMeCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstance client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceClient.Client)

	meCalendarViewInstanceExtensionClient, err := mecalendarviewinstanceextension.NewMeCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceExtensionClient.Client)

	meChatClient, err := mechat.NewMeChatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChat client: %+v", err)
	}
	configureFunc(meChatClient.Client)

	meChatInstalledAppClient, err := mechatinstalledapp.NewMeChatInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatInstalledApp client: %+v", err)
	}
	configureFunc(meChatInstalledAppClient.Client)

	meChatInstalledAppTeamsAppClient, err := mechatinstalledappteamsapp.NewMeChatInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(meChatInstalledAppTeamsAppClient.Client)

	meChatInstalledAppTeamsAppDefinitionClient, err := mechatinstalledappteamsappdefinition.NewMeChatInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(meChatInstalledAppTeamsAppDefinitionClient.Client)

	meChatLastMessagePreviewClient, err := mechatlastmessagepreview.NewMeChatLastMessagePreviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatLastMessagePreview client: %+v", err)
	}
	configureFunc(meChatLastMessagePreviewClient.Client)

	meChatMemberClient, err := mechatmember.NewMeChatMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatMember client: %+v", err)
	}
	configureFunc(meChatMemberClient.Client)

	meChatMessageClient, err := mechatmessage.NewMeChatMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatMessage client: %+v", err)
	}
	configureFunc(meChatMessageClient.Client)

	meChatMessageHostedContentClient, err := mechatmessagehostedcontent.NewMeChatMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatMessageHostedContent client: %+v", err)
	}
	configureFunc(meChatMessageHostedContentClient.Client)

	meChatMessageReplyClient, err := mechatmessagereply.NewMeChatMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatMessageReply client: %+v", err)
	}
	configureFunc(meChatMessageReplyClient.Client)

	meChatMessageReplyHostedContentClient, err := mechatmessagereplyhostedcontent.NewMeChatMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(meChatMessageReplyHostedContentClient.Client)

	meChatPermissionGrantClient, err := mechatpermissiongrant.NewMeChatPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatPermissionGrant client: %+v", err)
	}
	configureFunc(meChatPermissionGrantClient.Client)

	meChatPinnedMessageClient, err := mechatpinnedmessage.NewMeChatPinnedMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatPinnedMessage client: %+v", err)
	}
	configureFunc(meChatPinnedMessageClient.Client)

	meChatPinnedMessageMessageClient, err := mechatpinnedmessagemessage.NewMeChatPinnedMessageMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatPinnedMessageMessage client: %+v", err)
	}
	configureFunc(meChatPinnedMessageMessageClient.Client)

	meChatTabClient, err := mechattab.NewMeChatTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatTab client: %+v", err)
	}
	configureFunc(meChatTabClient.Client)

	meChatTabTeamsAppClient, err := mechattabteamsapp.NewMeChatTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatTabTeamsApp client: %+v", err)
	}
	configureFunc(meChatTabTeamsAppClient.Client)

	meClient, err := me.NewMeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Me client: %+v", err)
	}
	configureFunc(meClient.Client)

	meContactClient, err := mecontact.NewMeContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContact client: %+v", err)
	}
	configureFunc(meContactClient.Client)

	meContactExtensionClient, err := mecontactextension.NewMeContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactExtension client: %+v", err)
	}
	configureFunc(meContactExtensionClient.Client)

	meContactFolderChildFolderClient, err := mecontactfolderchildfolder.NewMeContactFolderChildFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolderChildFolder client: %+v", err)
	}
	configureFunc(meContactFolderChildFolderClient.Client)

	meContactFolderChildFolderContactClient, err := mecontactfolderchildfoldercontact.NewMeContactFolderChildFolderContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolderChildFolderContact client: %+v", err)
	}
	configureFunc(meContactFolderChildFolderContactClient.Client)

	meContactFolderChildFolderContactExtensionClient, err := mecontactfolderchildfoldercontactextension.NewMeContactFolderChildFolderContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolderChildFolderContactExtension client: %+v", err)
	}
	configureFunc(meContactFolderChildFolderContactExtensionClient.Client)

	meContactFolderChildFolderContactPhotoClient, err := mecontactfolderchildfoldercontactphoto.NewMeContactFolderChildFolderContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolderChildFolderContactPhoto client: %+v", err)
	}
	configureFunc(meContactFolderChildFolderContactPhotoClient.Client)

	meContactFolderClient, err := mecontactfolder.NewMeContactFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolder client: %+v", err)
	}
	configureFunc(meContactFolderClient.Client)

	meContactFolderContactClient, err := mecontactfoldercontact.NewMeContactFolderContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolderContact client: %+v", err)
	}
	configureFunc(meContactFolderContactClient.Client)

	meContactFolderContactExtensionClient, err := mecontactfoldercontactextension.NewMeContactFolderContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolderContactExtension client: %+v", err)
	}
	configureFunc(meContactFolderContactExtensionClient.Client)

	meContactFolderContactPhotoClient, err := mecontactfoldercontactphoto.NewMeContactFolderContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactFolderContactPhoto client: %+v", err)
	}
	configureFunc(meContactFolderContactPhotoClient.Client)

	meContactPhotoClient, err := mecontactphoto.NewMeContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeContactPhoto client: %+v", err)
	}
	configureFunc(meContactPhotoClient.Client)

	meCreatedObjectClient, err := mecreatedobject.NewMeCreatedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCreatedObject client: %+v", err)
	}
	configureFunc(meCreatedObjectClient.Client)

	meDeviceManagementTroubleshootingEventClient, err := medevicemanagementtroubleshootingevent.NewMeDeviceManagementTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceManagementTroubleshootingEvent client: %+v", err)
	}
	configureFunc(meDeviceManagementTroubleshootingEventClient.Client)

	meDirectReportClient, err := medirectreport.NewMeDirectReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDirectReport client: %+v", err)
	}
	configureFunc(meDirectReportClient.Client)

	meDriveClient, err := medrive.NewMeDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDrive client: %+v", err)
	}
	configureFunc(meDriveClient.Client)

	meEmployeeExperienceClient, err := meemployeeexperience.NewMeEmployeeExperienceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEmployeeExperience client: %+v", err)
	}
	configureFunc(meEmployeeExperienceClient.Client)

	meEmployeeExperienceLearningCourseActivityClient, err := meemployeeexperiencelearningcourseactivity.NewMeEmployeeExperienceLearningCourseActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEmployeeExperienceLearningCourseActivity client: %+v", err)
	}
	configureFunc(meEmployeeExperienceLearningCourseActivityClient.Client)

	meEventAttachmentClient, err := meeventattachment.NewMeEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventAttachment client: %+v", err)
	}
	configureFunc(meEventAttachmentClient.Client)

	meEventCalendarClient, err := meeventcalendar.NewMeEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventCalendar client: %+v", err)
	}
	configureFunc(meEventCalendarClient.Client)

	meEventClient, err := meevent.NewMeEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEvent client: %+v", err)
	}
	configureFunc(meEventClient.Client)

	meEventExtensionClient, err := meeventextension.NewMeEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExtension client: %+v", err)
	}
	configureFunc(meEventExtensionClient.Client)

	meEventInstanceAttachmentClient, err := meeventinstanceattachment.NewMeEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstanceAttachment client: %+v", err)
	}
	configureFunc(meEventInstanceAttachmentClient.Client)

	meEventInstanceCalendarClient, err := meeventinstancecalendar.NewMeEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstanceCalendar client: %+v", err)
	}
	configureFunc(meEventInstanceCalendarClient.Client)

	meEventInstanceClient, err := meeventinstance.NewMeEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstance client: %+v", err)
	}
	configureFunc(meEventInstanceClient.Client)

	meEventInstanceExtensionClient, err := meeventinstanceextension.NewMeEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstanceExtension client: %+v", err)
	}
	configureFunc(meEventInstanceExtensionClient.Client)

	meExtensionClient, err := meextension.NewMeExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeExtension client: %+v", err)
	}
	configureFunc(meExtensionClient.Client)

	meFollowedSiteClient, err := mefollowedsite.NewMeFollowedSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeFollowedSite client: %+v", err)
	}
	configureFunc(meFollowedSiteClient.Client)

	meInferenceClassificationClient, err := meinferenceclassification.NewMeInferenceClassificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInferenceClassification client: %+v", err)
	}
	configureFunc(meInferenceClassificationClient.Client)

	meInferenceClassificationOverrideClient, err := meinferenceclassificationoverride.NewMeInferenceClassificationOverrideClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInferenceClassificationOverride client: %+v", err)
	}
	configureFunc(meInferenceClassificationOverrideClient.Client)

	meInsightClient, err := meinsight.NewMeInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsight client: %+v", err)
	}
	configureFunc(meInsightClient.Client)

	meInsightSharedClient, err := meinsightshared.NewMeInsightSharedClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsightShared client: %+v", err)
	}
	configureFunc(meInsightSharedClient.Client)

	meInsightSharedLastSharedMethodClient, err := meinsightsharedlastsharedmethod.NewMeInsightSharedLastSharedMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsightSharedLastSharedMethod client: %+v", err)
	}
	configureFunc(meInsightSharedLastSharedMethodClient.Client)

	meInsightSharedResourceClient, err := meinsightsharedresource.NewMeInsightSharedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsightSharedResource client: %+v", err)
	}
	configureFunc(meInsightSharedResourceClient.Client)

	meInsightTrendingClient, err := meinsighttrending.NewMeInsightTrendingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsightTrending client: %+v", err)
	}
	configureFunc(meInsightTrendingClient.Client)

	meInsightTrendingResourceClient, err := meinsighttrendingresource.NewMeInsightTrendingResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsightTrendingResource client: %+v", err)
	}
	configureFunc(meInsightTrendingResourceClient.Client)

	meInsightUsedClient, err := meinsightused.NewMeInsightUsedClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsightUsed client: %+v", err)
	}
	configureFunc(meInsightUsedClient.Client)

	meInsightUsedResourceClient, err := meinsightusedresource.NewMeInsightUsedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInsightUsedResource client: %+v", err)
	}
	configureFunc(meInsightUsedResourceClient.Client)

	meJoinedTeamAllChannelClient, err := mejoinedteamallchannel.NewMeJoinedTeamAllChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamAllChannel client: %+v", err)
	}
	configureFunc(meJoinedTeamAllChannelClient.Client)

	meJoinedTeamChannelClient, err := mejoinedteamchannel.NewMeJoinedTeamChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannel client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelClient.Client)

	meJoinedTeamChannelFilesFolderClient, err := mejoinedteamchannelfilesfolder.NewMeJoinedTeamChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelFilesFolder client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelFilesFolderClient.Client)

	meJoinedTeamChannelFilesFolderContentClient, err := mejoinedteamchannelfilesfoldercontent.NewMeJoinedTeamChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelFilesFolderContentClient.Client)

	meJoinedTeamChannelMemberClient, err := mejoinedteamchannelmember.NewMeJoinedTeamChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelMember client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelMemberClient.Client)

	meJoinedTeamChannelMessageClient, err := mejoinedteamchannelmessage.NewMeJoinedTeamChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelMessage client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelMessageClient.Client)

	meJoinedTeamChannelMessageHostedContentClient, err := mejoinedteamchannelmessagehostedcontent.NewMeJoinedTeamChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelMessageHostedContentClient.Client)

	meJoinedTeamChannelMessageReplyClient, err := mejoinedteamchannelmessagereply.NewMeJoinedTeamChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelMessageReply client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelMessageReplyClient.Client)

	meJoinedTeamChannelMessageReplyHostedContentClient, err := mejoinedteamchannelmessagereplyhostedcontent.NewMeJoinedTeamChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelMessageReplyHostedContentClient.Client)

	meJoinedTeamChannelSharedWithTeamAllowedMemberClient, err := mejoinedteamchannelsharedwithteamallowedmember.NewMeJoinedTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelSharedWithTeamAllowedMemberClient.Client)

	meJoinedTeamChannelSharedWithTeamClient, err := mejoinedteamchannelsharedwithteam.NewMeJoinedTeamChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelSharedWithTeamClient.Client)

	meJoinedTeamChannelSharedWithTeamTeamClient, err := mejoinedteamchannelsharedwithteamteam.NewMeJoinedTeamChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelSharedWithTeamTeamClient.Client)

	meJoinedTeamChannelTabClient, err := mejoinedteamchanneltab.NewMeJoinedTeamChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelTab client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelTabClient.Client)

	meJoinedTeamChannelTabTeamsAppClient, err := mejoinedteamchanneltabteamsapp.NewMeJoinedTeamChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(meJoinedTeamChannelTabTeamsAppClient.Client)

	meJoinedTeamClient, err := mejoinedteam.NewMeJoinedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeam client: %+v", err)
	}
	configureFunc(meJoinedTeamClient.Client)

	meJoinedTeamGroupClient, err := mejoinedteamgroup.NewMeJoinedTeamGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamGroup client: %+v", err)
	}
	configureFunc(meJoinedTeamGroupClient.Client)

	meJoinedTeamIncomingChannelClient, err := mejoinedteamincomingchannel.NewMeJoinedTeamIncomingChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamIncomingChannel client: %+v", err)
	}
	configureFunc(meJoinedTeamIncomingChannelClient.Client)

	meJoinedTeamInstalledAppClient, err := mejoinedteaminstalledapp.NewMeJoinedTeamInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamInstalledApp client: %+v", err)
	}
	configureFunc(meJoinedTeamInstalledAppClient.Client)

	meJoinedTeamInstalledAppTeamsAppClient, err := mejoinedteaminstalledappteamsapp.NewMeJoinedTeamInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(meJoinedTeamInstalledAppTeamsAppClient.Client)

	meJoinedTeamInstalledAppTeamsAppDefinitionClient, err := mejoinedteaminstalledappteamsappdefinition.NewMeJoinedTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(meJoinedTeamInstalledAppTeamsAppDefinitionClient.Client)

	meJoinedTeamMemberClient, err := mejoinedteammember.NewMeJoinedTeamMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamMember client: %+v", err)
	}
	configureFunc(meJoinedTeamMemberClient.Client)

	meJoinedTeamOperationClient, err := mejoinedteamoperation.NewMeJoinedTeamOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamOperation client: %+v", err)
	}
	configureFunc(meJoinedTeamOperationClient.Client)

	meJoinedTeamPermissionGrantClient, err := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPermissionGrant client: %+v", err)
	}
	configureFunc(meJoinedTeamPermissionGrantClient.Client)

	meJoinedTeamPhotoClient, err := mejoinedteamphoto.NewMeJoinedTeamPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPhoto client: %+v", err)
	}
	configureFunc(meJoinedTeamPhotoClient.Client)

	meJoinedTeamPrimaryChannelClient, err := mejoinedteamprimarychannel.NewMeJoinedTeamPrimaryChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannel client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelClient.Client)

	meJoinedTeamPrimaryChannelFilesFolderClient, err := mejoinedteamprimarychannelfilesfolder.NewMeJoinedTeamPrimaryChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelFilesFolder client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelFilesFolderClient.Client)

	meJoinedTeamPrimaryChannelFilesFolderContentClient, err := mejoinedteamprimarychannelfilesfoldercontent.NewMeJoinedTeamPrimaryChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelFilesFolderContentClient.Client)

	meJoinedTeamPrimaryChannelMemberClient, err := mejoinedteamprimarychannelmember.NewMeJoinedTeamPrimaryChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelMember client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelMemberClient.Client)

	meJoinedTeamPrimaryChannelMessageClient, err := mejoinedteamprimarychannelmessage.NewMeJoinedTeamPrimaryChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelMessage client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelMessageClient.Client)

	meJoinedTeamPrimaryChannelMessageHostedContentClient, err := mejoinedteamprimarychannelmessagehostedcontent.NewMeJoinedTeamPrimaryChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelMessageHostedContentClient.Client)

	meJoinedTeamPrimaryChannelMessageReplyClient, err := mejoinedteamprimarychannelmessagereply.NewMeJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelMessageReply client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelMessageReplyClient.Client)

	meJoinedTeamPrimaryChannelMessageReplyHostedContentClient, err := mejoinedteamprimarychannelmessagereplyhostedcontent.NewMeJoinedTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelMessageReplyHostedContentClient.Client)

	meJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient, err := mejoinedteamprimarychannelsharedwithteamallowedmember.NewMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient.Client)

	meJoinedTeamPrimaryChannelSharedWithTeamClient, err := mejoinedteamprimarychannelsharedwithteam.NewMeJoinedTeamPrimaryChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelSharedWithTeamClient.Client)

	meJoinedTeamPrimaryChannelSharedWithTeamTeamClient, err := mejoinedteamprimarychannelsharedwithteamteam.NewMeJoinedTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelSharedWithTeamTeamClient.Client)

	meJoinedTeamPrimaryChannelTabClient, err := mejoinedteamprimarychanneltab.NewMeJoinedTeamPrimaryChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelTab client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelTabClient.Client)

	meJoinedTeamPrimaryChannelTabTeamsAppClient, err := mejoinedteamprimarychanneltabteamsapp.NewMeJoinedTeamPrimaryChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamPrimaryChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(meJoinedTeamPrimaryChannelTabTeamsAppClient.Client)

	meJoinedTeamScheduleClient, err := mejoinedteamschedule.NewMeJoinedTeamScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamSchedule client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleClient.Client)

	meJoinedTeamScheduleOfferShiftRequestClient, err := mejoinedteamscheduleoffershiftrequest.NewMeJoinedTeamScheduleOfferShiftRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleOfferShiftRequest client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleOfferShiftRequestClient.Client)

	meJoinedTeamScheduleOpenShiftChangeRequestClient, err := mejoinedteamscheduleopenshiftchangerequest.NewMeJoinedTeamScheduleOpenShiftChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleOpenShiftChangeRequest client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleOpenShiftChangeRequestClient.Client)

	meJoinedTeamScheduleOpenShiftClient, err := mejoinedteamscheduleopenshift.NewMeJoinedTeamScheduleOpenShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleOpenShift client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleOpenShiftClient.Client)

	meJoinedTeamScheduleSchedulingGroupClient, err := mejoinedteamscheduleschedulinggroup.NewMeJoinedTeamScheduleSchedulingGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleSchedulingGroup client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleSchedulingGroupClient.Client)

	meJoinedTeamScheduleShiftClient, err := mejoinedteamscheduleshift.NewMeJoinedTeamScheduleShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleShift client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleShiftClient.Client)

	meJoinedTeamScheduleSwapShiftsChangeRequestClient, err := mejoinedteamscheduleswapshiftschangerequest.NewMeJoinedTeamScheduleSwapShiftsChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleSwapShiftsChangeRequest client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleSwapShiftsChangeRequestClient.Client)

	meJoinedTeamScheduleTimeOffReasonClient, err := mejoinedteamscheduletimeoffreason.NewMeJoinedTeamScheduleTimeOffReasonClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleTimeOffReason client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleTimeOffReasonClient.Client)

	meJoinedTeamScheduleTimeOffRequestClient, err := mejoinedteamscheduletimeoffrequest.NewMeJoinedTeamScheduleTimeOffRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleTimeOffRequest client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleTimeOffRequestClient.Client)

	meJoinedTeamScheduleTimesOffClient, err := mejoinedteamscheduletimesoff.NewMeJoinedTeamScheduleTimesOffClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamScheduleTimesOff client: %+v", err)
	}
	configureFunc(meJoinedTeamScheduleTimesOffClient.Client)

	meJoinedTeamTagClient, err := mejoinedteamtag.NewMeJoinedTeamTagClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamTag client: %+v", err)
	}
	configureFunc(meJoinedTeamTagClient.Client)

	meJoinedTeamTagMemberClient, err := mejoinedteamtagmember.NewMeJoinedTeamTagMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamTagMember client: %+v", err)
	}
	configureFunc(meJoinedTeamTagMemberClient.Client)

	meJoinedTeamTemplateClient, err := mejoinedteamtemplate.NewMeJoinedTeamTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeamTemplate client: %+v", err)
	}
	configureFunc(meJoinedTeamTemplateClient.Client)

	meLicenseDetailClient, err := melicensedetail.NewMeLicenseDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeLicenseDetail client: %+v", err)
	}
	configureFunc(meLicenseDetailClient.Client)

	meMailFolderChildFolderClient, err := memailfolderchildfolder.NewMeMailFolderChildFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolder client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderClient.Client)

	meMailFolderChildFolderMessageAttachmentClient, err := memailfolderchildfoldermessageattachment.NewMeMailFolderChildFolderMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolderMessageAttachment client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderMessageAttachmentClient.Client)

	meMailFolderChildFolderMessageClient, err := memailfolderchildfoldermessage.NewMeMailFolderChildFolderMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolderMessage client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderMessageClient.Client)

	meMailFolderChildFolderMessageExtensionClient, err := memailfolderchildfoldermessageextension.NewMeMailFolderChildFolderMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolderMessageExtension client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderMessageExtensionClient.Client)

	meMailFolderChildFolderMessageRuleClient, err := memailfolderchildfoldermessagerule.NewMeMailFolderChildFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolderMessageRule client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderMessageRuleClient.Client)

	meMailFolderClient, err := memailfolder.NewMeMailFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolder client: %+v", err)
	}
	configureFunc(meMailFolderClient.Client)

	meMailFolderMessageAttachmentClient, err := memailfoldermessageattachment.NewMeMailFolderMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderMessageAttachment client: %+v", err)
	}
	configureFunc(meMailFolderMessageAttachmentClient.Client)

	meMailFolderMessageClient, err := memailfoldermessage.NewMeMailFolderMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderMessage client: %+v", err)
	}
	configureFunc(meMailFolderMessageClient.Client)

	meMailFolderMessageExtensionClient, err := memailfoldermessageextension.NewMeMailFolderMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderMessageExtension client: %+v", err)
	}
	configureFunc(meMailFolderMessageExtensionClient.Client)

	meMailFolderMessageRuleClient, err := memailfoldermessagerule.NewMeMailFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderMessageRule client: %+v", err)
	}
	configureFunc(meMailFolderMessageRuleClient.Client)

	meMailboxSettingClient, err := memailboxsetting.NewMeMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailboxSetting client: %+v", err)
	}
	configureFunc(meMailboxSettingClient.Client)

	meManagedAppRegistrationClient, err := memanagedappregistration.NewMeManagedAppRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedAppRegistration client: %+v", err)
	}
	configureFunc(meManagedAppRegistrationClient.Client)

	meManagedDeviceClient, err := memanageddevice.NewMeManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDevice client: %+v", err)
	}
	configureFunc(meManagedDeviceClient.Client)

	meManagedDeviceDeviceCategoryClient, err := memanageddevicedevicecategory.NewMeManagedDeviceDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceDeviceCategory client: %+v", err)
	}
	configureFunc(meManagedDeviceDeviceCategoryClient.Client)

	meManagedDeviceDeviceCompliancePolicyStateClient, err := memanageddevicedevicecompliancepolicystate.NewMeManagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceDeviceCompliancePolicyState client: %+v", err)
	}
	configureFunc(meManagedDeviceDeviceCompliancePolicyStateClient.Client)

	meManagedDeviceDeviceConfigurationStateClient, err := memanageddevicedeviceconfigurationstate.NewMeManagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceDeviceConfigurationState client: %+v", err)
	}
	configureFunc(meManagedDeviceDeviceConfigurationStateClient.Client)

	meManagedDeviceLogCollectionRequestClient, err := memanageddevicelogcollectionrequest.NewMeManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(meManagedDeviceLogCollectionRequestClient.Client)

	meManagedDeviceUserClient, err := memanageddeviceuser.NewMeManagedDeviceUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceUser client: %+v", err)
	}
	configureFunc(meManagedDeviceUserClient.Client)

	meManagedDeviceWindowsProtectionStateClient, err := memanageddevicewindowsprotectionstate.NewMeManagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceWindowsProtectionState client: %+v", err)
	}
	configureFunc(meManagedDeviceWindowsProtectionStateClient.Client)

	meManagedDeviceWindowsProtectionStateDetectedMalwareStateClient, err := memanageddevicewindowsprotectionstatedetectedmalwarestate.NewMeManagedDeviceWindowsProtectionStateDetectedMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceWindowsProtectionStateDetectedMalwareState client: %+v", err)
	}
	configureFunc(meManagedDeviceWindowsProtectionStateDetectedMalwareStateClient.Client)

	meManagerClient, err := memanager.NewMeManagerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManager client: %+v", err)
	}
	configureFunc(meManagerClient.Client)

	meMemberOfClient, err := mememberof.NewMeMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMemberOf client: %+v", err)
	}
	configureFunc(meMemberOfClient.Client)

	meMessageAttachmentClient, err := memessageattachment.NewMeMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMessageAttachment client: %+v", err)
	}
	configureFunc(meMessageAttachmentClient.Client)

	meMessageClient, err := memessage.NewMeMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMessage client: %+v", err)
	}
	configureFunc(meMessageClient.Client)

	meMessageExtensionClient, err := memessageextension.NewMeMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMessageExtension client: %+v", err)
	}
	configureFunc(meMessageExtensionClient.Client)

	meOauth2PermissionGrantClient, err := meoauth2permissiongrant.NewMeOauth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOauth2PermissionGrant client: %+v", err)
	}
	configureFunc(meOauth2PermissionGrantClient.Client)

	meOnenoteClient, err := meonenote.NewMeOnenoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenote client: %+v", err)
	}
	configureFunc(meOnenoteClient.Client)

	meOnenoteNotebookClient, err := meonenotenotebook.NewMeOnenoteNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebook client: %+v", err)
	}
	configureFunc(meOnenoteNotebookClient.Client)

	meOnenoteNotebookSectionClient, err := meonenotenotebooksection.NewMeOnenoteNotebookSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSection client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionClient.Client)

	meOnenoteNotebookSectionGroupClient, err := meonenotenotebooksectiongroup.NewMeOnenoteNotebookSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupClient.Client)

	meOnenoteNotebookSectionGroupParentNotebookClient, err := meonenotenotebooksectiongroupparentnotebook.NewMeOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupParentNotebookClient.Client)

	meOnenoteNotebookSectionGroupParentSectionGroupClient, err := meonenotenotebooksectiongroupparentsectiongroup.NewMeOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupParentSectionGroupClient.Client)

	meOnenoteNotebookSectionGroupSectionClient, err := meonenotenotebooksectiongroupsection.NewMeOnenoteNotebookSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSection client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionClient.Client)

	meOnenoteNotebookSectionGroupSectionGroupClient, err := meonenotenotebooksectiongroupsectiongroup.NewMeOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionGroupClient.Client)

	meOnenoteNotebookSectionGroupSectionPageClient, err := meonenotenotebooksectiongroupsectionpage.NewMeOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionPageClient.Client)

	meOnenoteNotebookSectionGroupSectionPageContentClient, err := meonenotenotebooksectiongroupsectionpagecontent.NewMeOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionPageContentClient.Client)

	meOnenoteNotebookSectionGroupSectionPageParentNotebookClient, err := meonenotenotebooksectiongroupsectionpageparentnotebook.NewMeOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionPageParentNotebookClient.Client)

	meOnenoteNotebookSectionGroupSectionPageParentSectionClient, err := meonenotenotebooksectiongroupsectionpageparentsection.NewMeOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionPageParentSectionClient.Client)

	meOnenoteNotebookSectionGroupSectionParentNotebookClient, err := meonenotenotebooksectiongroupsectionparentnotebook.NewMeOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionParentNotebookClient.Client)

	meOnenoteNotebookSectionGroupSectionParentSectionGroupClient, err := meonenotenotebooksectiongroupsectionparentsectiongroup.NewMeOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionGroupSectionParentSectionGroupClient.Client)

	meOnenoteNotebookSectionPageClient, err := meonenotenotebooksectionpage.NewMeOnenoteNotebookSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionPage client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionPageClient.Client)

	meOnenoteNotebookSectionPageContentClient, err := meonenotenotebooksectionpagecontent.NewMeOnenoteNotebookSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionPageContent client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionPageContentClient.Client)

	meOnenoteNotebookSectionPageParentNotebookClient, err := meonenotenotebooksectionpageparentnotebook.NewMeOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionPageParentNotebookClient.Client)

	meOnenoteNotebookSectionPageParentSectionClient, err := meonenotenotebooksectionpageparentsection.NewMeOnenoteNotebookSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionPageParentSection client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionPageParentSectionClient.Client)

	meOnenoteNotebookSectionParentNotebookClient, err := meonenotenotebooksectionparentnotebook.NewMeOnenoteNotebookSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionParentNotebookClient.Client)

	meOnenoteNotebookSectionParentSectionGroupClient, err := meonenotenotebooksectionparentsectiongroup.NewMeOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteNotebookSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteNotebookSectionParentSectionGroupClient.Client)

	meOnenoteOperationClient, err := meonenoteoperation.NewMeOnenoteOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteOperation client: %+v", err)
	}
	configureFunc(meOnenoteOperationClient.Client)

	meOnenotePageClient, err := meonenotepage.NewMeOnenotePageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenotePage client: %+v", err)
	}
	configureFunc(meOnenotePageClient.Client)

	meOnenotePageContentClient, err := meonenotepagecontent.NewMeOnenotePageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenotePageContent client: %+v", err)
	}
	configureFunc(meOnenotePageContentClient.Client)

	meOnenotePageParentNotebookClient, err := meonenotepageparentnotebook.NewMeOnenotePageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenotePageParentNotebook client: %+v", err)
	}
	configureFunc(meOnenotePageParentNotebookClient.Client)

	meOnenotePageParentSectionClient, err := meonenotepageparentsection.NewMeOnenotePageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenotePageParentSection client: %+v", err)
	}
	configureFunc(meOnenotePageParentSectionClient.Client)

	meOnenoteResourceClient, err := meonenoteresource.NewMeOnenoteResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteResource client: %+v", err)
	}
	configureFunc(meOnenoteResourceClient.Client)

	meOnenoteResourceContentClient, err := meonenoteresourcecontent.NewMeOnenoteResourceContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteResourceContent client: %+v", err)
	}
	configureFunc(meOnenoteResourceContentClient.Client)

	meOnenoteSectionClient, err := meonenotesection.NewMeOnenoteSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSection client: %+v", err)
	}
	configureFunc(meOnenoteSectionClient.Client)

	meOnenoteSectionGroupClient, err := meonenotesectiongroup.NewMeOnenoteSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupClient.Client)

	meOnenoteSectionGroupParentNotebookClient, err := meonenotesectiongroupparentnotebook.NewMeOnenoteSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupParentNotebookClient.Client)

	meOnenoteSectionGroupParentSectionGroupClient, err := meonenotesectiongroupparentsectiongroup.NewMeOnenoteSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupParentSectionGroupClient.Client)

	meOnenoteSectionGroupSectionClient, err := meonenotesectiongroupsection.NewMeOnenoteSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSection client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionClient.Client)

	meOnenoteSectionGroupSectionGroupClient, err := meonenotesectiongroupsectiongroup.NewMeOnenoteSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionGroupClient.Client)

	meOnenoteSectionGroupSectionPageClient, err := meonenotesectiongroupsectionpage.NewMeOnenoteSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionPageClient.Client)

	meOnenoteSectionGroupSectionPageContentClient, err := meonenotesectiongroupsectionpagecontent.NewMeOnenoteSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionPageContentClient.Client)

	meOnenoteSectionGroupSectionPageParentNotebookClient, err := meonenotesectiongroupsectionpageparentnotebook.NewMeOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionPageParentNotebookClient.Client)

	meOnenoteSectionGroupSectionPageParentSectionClient, err := meonenotesectiongroupsectionpageparentsection.NewMeOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionPageParentSectionClient.Client)

	meOnenoteSectionGroupSectionParentNotebookClient, err := meonenotesectiongroupsectionparentnotebook.NewMeOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionParentNotebookClient.Client)

	meOnenoteSectionGroupSectionParentSectionGroupClient, err := meonenotesectiongroupsectionparentsectiongroup.NewMeOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteSectionGroupSectionParentSectionGroupClient.Client)

	meOnenoteSectionPageClient, err := meonenotesectionpage.NewMeOnenoteSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionPage client: %+v", err)
	}
	configureFunc(meOnenoteSectionPageClient.Client)

	meOnenoteSectionPageContentClient, err := meonenotesectionpagecontent.NewMeOnenoteSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionPageContent client: %+v", err)
	}
	configureFunc(meOnenoteSectionPageContentClient.Client)

	meOnenoteSectionPageParentNotebookClient, err := meonenotesectionpageparentnotebook.NewMeOnenoteSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteSectionPageParentNotebookClient.Client)

	meOnenoteSectionPageParentSectionClient, err := meonenotesectionpageparentsection.NewMeOnenoteSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionPageParentSection client: %+v", err)
	}
	configureFunc(meOnenoteSectionPageParentSectionClient.Client)

	meOnenoteSectionParentNotebookClient, err := meonenotesectionparentnotebook.NewMeOnenoteSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionParentNotebook client: %+v", err)
	}
	configureFunc(meOnenoteSectionParentNotebookClient.Client)

	meOnenoteSectionParentSectionGroupClient, err := meonenotesectionparentsectiongroup.NewMeOnenoteSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnenoteSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(meOnenoteSectionParentSectionGroupClient.Client)

	meOnlineMeetingAttendanceReportAttendanceRecordClient, err := meonlinemeetingattendancereportattendancerecord.NewMeOnlineMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingAttendanceReportAttendanceRecord client: %+v", err)
	}
	configureFunc(meOnlineMeetingAttendanceReportAttendanceRecordClient.Client)

	meOnlineMeetingAttendanceReportClient, err := meonlinemeetingattendancereport.NewMeOnlineMeetingAttendanceReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingAttendanceReport client: %+v", err)
	}
	configureFunc(meOnlineMeetingAttendanceReportClient.Client)

	meOnlineMeetingAttendeeReportClient, err := meonlinemeetingattendeereport.NewMeOnlineMeetingAttendeeReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingAttendeeReport client: %+v", err)
	}
	configureFunc(meOnlineMeetingAttendeeReportClient.Client)

	meOnlineMeetingClient, err := meonlinemeeting.NewMeOnlineMeetingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeeting client: %+v", err)
	}
	configureFunc(meOnlineMeetingClient.Client)

	meOutlookClient, err := meoutlook.NewMeOutlookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlook client: %+v", err)
	}
	configureFunc(meOutlookClient.Client)

	meOutlookMasterCategoryClient, err := meoutlookmastercategory.NewMeOutlookMasterCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookMasterCategory client: %+v", err)
	}
	configureFunc(meOutlookMasterCategoryClient.Client)

	meOwnedDeviceClient, err := meowneddevice.NewMeOwnedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOwnedDevice client: %+v", err)
	}
	configureFunc(meOwnedDeviceClient.Client)

	meOwnedObjectClient, err := meownedobject.NewMeOwnedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOwnedObject client: %+v", err)
	}
	configureFunc(meOwnedObjectClient.Client)

	mePeopleClient, err := mepeople.NewMePeopleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePeople client: %+v", err)
	}
	configureFunc(mePeopleClient.Client)

	mePhotoClient, err := mephoto.NewMePhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePhoto client: %+v", err)
	}
	configureFunc(mePhotoClient.Client)

	mePlannerClient, err := meplanner.NewMePlannerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlanner client: %+v", err)
	}
	configureFunc(mePlannerClient.Client)

	mePlannerPlanBucketClient, err := meplannerplanbucket.NewMePlannerPlanBucketClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanBucket client: %+v", err)
	}
	configureFunc(mePlannerPlanBucketClient.Client)

	mePlannerPlanBucketTaskAssignedToTaskBoardFormatClient, err := meplannerplanbuckettaskassignedtotaskboardformat.NewMePlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanBucketTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerPlanBucketTaskAssignedToTaskBoardFormatClient.Client)

	mePlannerPlanBucketTaskBucketTaskBoardFormatClient, err := meplannerplanbuckettaskbuckettaskboardformat.NewMePlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanBucketTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerPlanBucketTaskBucketTaskBoardFormatClient.Client)

	mePlannerPlanBucketTaskClient, err := meplannerplanbuckettask.NewMePlannerPlanBucketTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanBucketTask client: %+v", err)
	}
	configureFunc(mePlannerPlanBucketTaskClient.Client)

	mePlannerPlanBucketTaskDetailClient, err := meplannerplanbuckettaskdetail.NewMePlannerPlanBucketTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanBucketTaskDetail client: %+v", err)
	}
	configureFunc(mePlannerPlanBucketTaskDetailClient.Client)

	mePlannerPlanBucketTaskProgressTaskBoardFormatClient, err := meplannerplanbuckettaskprogresstaskboardformat.NewMePlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanBucketTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerPlanBucketTaskProgressTaskBoardFormatClient.Client)

	mePlannerPlanClient, err := meplannerplan.NewMePlannerPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlan client: %+v", err)
	}
	configureFunc(mePlannerPlanClient.Client)

	mePlannerPlanDetailClient, err := meplannerplandetail.NewMePlannerPlanDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanDetail client: %+v", err)
	}
	configureFunc(mePlannerPlanDetailClient.Client)

	mePlannerPlanTaskAssignedToTaskBoardFormatClient, err := meplannerplantaskassignedtotaskboardformat.NewMePlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerPlanTaskAssignedToTaskBoardFormatClient.Client)

	mePlannerPlanTaskBucketTaskBoardFormatClient, err := meplannerplantaskbuckettaskboardformat.NewMePlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerPlanTaskBucketTaskBoardFormatClient.Client)

	mePlannerPlanTaskClient, err := meplannerplantask.NewMePlannerPlanTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanTask client: %+v", err)
	}
	configureFunc(mePlannerPlanTaskClient.Client)

	mePlannerPlanTaskDetailClient, err := meplannerplantaskdetail.NewMePlannerPlanTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanTaskDetail client: %+v", err)
	}
	configureFunc(mePlannerPlanTaskDetailClient.Client)

	mePlannerPlanTaskProgressTaskBoardFormatClient, err := meplannerplantaskprogresstaskboardformat.NewMePlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerPlanTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerPlanTaskProgressTaskBoardFormatClient.Client)

	mePlannerTaskAssignedToTaskBoardFormatClient, err := meplannertaskassignedtotaskboardformat.NewMePlannerTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerTaskAssignedToTaskBoardFormatClient.Client)

	mePlannerTaskBucketTaskBoardFormatClient, err := meplannertaskbuckettaskboardformat.NewMePlannerTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerTaskBucketTaskBoardFormatClient.Client)

	mePlannerTaskClient, err := meplannertask.NewMePlannerTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerTask client: %+v", err)
	}
	configureFunc(mePlannerTaskClient.Client)

	mePlannerTaskDetailClient, err := meplannertaskdetail.NewMePlannerTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerTaskDetail client: %+v", err)
	}
	configureFunc(mePlannerTaskDetailClient.Client)

	mePlannerTaskProgressTaskBoardFormatClient, err := meplannertaskprogresstaskboardformat.NewMePlannerTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(mePlannerTaskProgressTaskBoardFormatClient.Client)

	mePresenceClient, err := mepresence.NewMePresenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePresence client: %+v", err)
	}
	configureFunc(mePresenceClient.Client)

	meRegisteredDeviceClient, err := meregistereddevice.NewMeRegisteredDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeRegisteredDevice client: %+v", err)
	}
	configureFunc(meRegisteredDeviceClient.Client)

	meScopedRoleMemberOfClient, err := mescopedrolememberof.NewMeScopedRoleMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeScopedRoleMemberOf client: %+v", err)
	}
	configureFunc(meScopedRoleMemberOfClient.Client)

	meSettingClient, err := mesetting.NewMeSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSetting client: %+v", err)
	}
	configureFunc(meSettingClient.Client)

	meSettingShiftPreferenceClient, err := mesettingshiftpreference.NewMeSettingShiftPreferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSettingShiftPreference client: %+v", err)
	}
	configureFunc(meSettingShiftPreferenceClient.Client)

	meTeamworkAssociatedTeamClient, err := meteamworkassociatedteam.NewMeTeamworkAssociatedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTeamworkAssociatedTeam client: %+v", err)
	}
	configureFunc(meTeamworkAssociatedTeamClient.Client)

	meTeamworkAssociatedTeamTeamClient, err := meteamworkassociatedteamteam.NewMeTeamworkAssociatedTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTeamworkAssociatedTeamTeam client: %+v", err)
	}
	configureFunc(meTeamworkAssociatedTeamTeamClient.Client)

	meTeamworkClient, err := meteamwork.NewMeTeamworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTeamwork client: %+v", err)
	}
	configureFunc(meTeamworkClient.Client)

	meTeamworkInstalledAppChatClient, err := meteamworkinstalledappchat.NewMeTeamworkInstalledAppChatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTeamworkInstalledAppChat client: %+v", err)
	}
	configureFunc(meTeamworkInstalledAppChatClient.Client)

	meTeamworkInstalledAppClient, err := meteamworkinstalledapp.NewMeTeamworkInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTeamworkInstalledApp client: %+v", err)
	}
	configureFunc(meTeamworkInstalledAppClient.Client)

	meTeamworkInstalledAppTeamsAppClient, err := meteamworkinstalledappteamsapp.NewMeTeamworkInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTeamworkInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(meTeamworkInstalledAppTeamsAppClient.Client)

	meTeamworkInstalledAppTeamsAppDefinitionClient, err := meteamworkinstalledappteamsappdefinition.NewMeTeamworkInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTeamworkInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(meTeamworkInstalledAppTeamsAppDefinitionClient.Client)

	meTodoClient, err := metodo.NewMeTodoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodo client: %+v", err)
	}
	configureFunc(meTodoClient.Client)

	meTodoListClient, err := metodolist.NewMeTodoListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoList client: %+v", err)
	}
	configureFunc(meTodoListClient.Client)

	meTodoListExtensionClient, err := metodolistextension.NewMeTodoListExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListExtension client: %+v", err)
	}
	configureFunc(meTodoListExtensionClient.Client)

	meTodoListTaskAttachmentClient, err := metodolisttaskattachment.NewMeTodoListTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListTaskAttachment client: %+v", err)
	}
	configureFunc(meTodoListTaskAttachmentClient.Client)

	meTodoListTaskAttachmentSessionClient, err := metodolisttaskattachmentsession.NewMeTodoListTaskAttachmentSessionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListTaskAttachmentSession client: %+v", err)
	}
	configureFunc(meTodoListTaskAttachmentSessionClient.Client)

	meTodoListTaskAttachmentSessionContentClient, err := metodolisttaskattachmentsessioncontent.NewMeTodoListTaskAttachmentSessionContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListTaskAttachmentSessionContent client: %+v", err)
	}
	configureFunc(meTodoListTaskAttachmentSessionContentClient.Client)

	meTodoListTaskChecklistItemClient, err := metodolisttaskchecklistitem.NewMeTodoListTaskChecklistItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListTaskChecklistItem client: %+v", err)
	}
	configureFunc(meTodoListTaskChecklistItemClient.Client)

	meTodoListTaskClient, err := metodolisttask.NewMeTodoListTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListTask client: %+v", err)
	}
	configureFunc(meTodoListTaskClient.Client)

	meTodoListTaskExtensionClient, err := metodolisttaskextension.NewMeTodoListTaskExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListTaskExtension client: %+v", err)
	}
	configureFunc(meTodoListTaskExtensionClient.Client)

	meTodoListTaskLinkedResourceClient, err := metodolisttasklinkedresource.NewMeTodoListTaskLinkedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTodoListTaskLinkedResource client: %+v", err)
	}
	configureFunc(meTodoListTaskLinkedResourceClient.Client)

	meTransitiveMemberOfClient, err := metransitivememberof.NewMeTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTransitiveMemberOf client: %+v", err)
	}
	configureFunc(meTransitiveMemberOfClient.Client)

	return &Client{
		Me:                            meClient,
		MeActivity:                    meActivityClient,
		MeActivityHistoryItem:         meActivityHistoryItemClient,
		MeActivityHistoryItemActivity: meActivityHistoryItemActivityClient,
		MeAgreementAcceptance:         meAgreementAcceptanceClient,
		MeAppRoleAssignment:           meAppRoleAssignmentClient,
		MeAuthentication:              meAuthenticationClient,
		MeAuthenticationEmailMethod:   meAuthenticationEmailMethodClient,
		MeAuthenticationFido2Method:   meAuthenticationFido2MethodClient,
		MeAuthenticationMethod:        meAuthenticationMethodClient,
		MeAuthenticationMicrosoftAuthenticatorMethod:        meAuthenticationMicrosoftAuthenticatorMethodClient,
		MeAuthenticationMicrosoftAuthenticatorMethodDevice:  meAuthenticationMicrosoftAuthenticatorMethodDeviceClient,
		MeAuthenticationOperation:                           meAuthenticationOperationClient,
		MeAuthenticationPasswordMethod:                      meAuthenticationPasswordMethodClient,
		MeAuthenticationPhoneMethod:                         meAuthenticationPhoneMethodClient,
		MeAuthenticationSoftwareOathMethod:                  meAuthenticationSoftwareOathMethodClient,
		MeAuthenticationTemporaryAccessPassMethod:           meAuthenticationTemporaryAccessPassMethodClient,
		MeAuthenticationWindowsHelloForBusinessMethod:       meAuthenticationWindowsHelloForBusinessMethodClient,
		MeAuthenticationWindowsHelloForBusinessMethodDevice: meAuthenticationWindowsHelloForBusinessMethodDeviceClient,
		MeCalendar:                                            meCalendarClient,
		MeCalendarCalendarPermission:                          meCalendarCalendarPermissionClient,
		MeCalendarCalendarView:                                meCalendarCalendarViewClient,
		MeCalendarCalendarViewAttachment:                      meCalendarCalendarViewAttachmentClient,
		MeCalendarCalendarViewCalendar:                        meCalendarCalendarViewCalendarClient,
		MeCalendarCalendarViewExtension:                       meCalendarCalendarViewExtensionClient,
		MeCalendarCalendarViewInstance:                        meCalendarCalendarViewInstanceClient,
		MeCalendarCalendarViewInstanceAttachment:              meCalendarCalendarViewInstanceAttachmentClient,
		MeCalendarCalendarViewInstanceCalendar:                meCalendarCalendarViewInstanceCalendarClient,
		MeCalendarCalendarViewInstanceExtension:               meCalendarCalendarViewInstanceExtensionClient,
		MeCalendarEvent:                                       meCalendarEventClient,
		MeCalendarEventAttachment:                             meCalendarEventAttachmentClient,
		MeCalendarEventCalendar:                               meCalendarEventCalendarClient,
		MeCalendarEventExtension:                              meCalendarEventExtensionClient,
		MeCalendarEventInstance:                               meCalendarEventInstanceClient,
		MeCalendarEventInstanceAttachment:                     meCalendarEventInstanceAttachmentClient,
		MeCalendarEventInstanceCalendar:                       meCalendarEventInstanceCalendarClient,
		MeCalendarEventInstanceExtension:                      meCalendarEventInstanceExtensionClient,
		MeCalendarGroup:                                       meCalendarGroupClient,
		MeCalendarGroupCalendar:                               meCalendarGroupCalendarClient,
		MeCalendarGroupCalendarCalendarPermission:             meCalendarGroupCalendarCalendarPermissionClient,
		MeCalendarGroupCalendarCalendarView:                   meCalendarGroupCalendarCalendarViewClient,
		MeCalendarGroupCalendarCalendarViewAttachment:         meCalendarGroupCalendarCalendarViewAttachmentClient,
		MeCalendarGroupCalendarCalendarViewCalendar:           meCalendarGroupCalendarCalendarViewCalendarClient,
		MeCalendarGroupCalendarCalendarViewExtension:          meCalendarGroupCalendarCalendarViewExtensionClient,
		MeCalendarGroupCalendarCalendarViewInstance:           meCalendarGroupCalendarCalendarViewInstanceClient,
		MeCalendarGroupCalendarCalendarViewInstanceAttachment: meCalendarGroupCalendarCalendarViewInstanceAttachmentClient,
		MeCalendarGroupCalendarCalendarViewInstanceCalendar:   meCalendarGroupCalendarCalendarViewInstanceCalendarClient,
		MeCalendarGroupCalendarCalendarViewInstanceExtension:  meCalendarGroupCalendarCalendarViewInstanceExtensionClient,
		MeCalendarGroupCalendarEvent:                          meCalendarGroupCalendarEventClient,
		MeCalendarGroupCalendarEventAttachment:                meCalendarGroupCalendarEventAttachmentClient,
		MeCalendarGroupCalendarEventCalendar:                  meCalendarGroupCalendarEventCalendarClient,
		MeCalendarGroupCalendarEventExtension:                 meCalendarGroupCalendarEventExtensionClient,
		MeCalendarGroupCalendarEventInstance:                  meCalendarGroupCalendarEventInstanceClient,
		MeCalendarGroupCalendarEventInstanceAttachment:        meCalendarGroupCalendarEventInstanceAttachmentClient,
		MeCalendarGroupCalendarEventInstanceCalendar:          meCalendarGroupCalendarEventInstanceCalendarClient,
		MeCalendarGroupCalendarEventInstanceExtension:         meCalendarGroupCalendarEventInstanceExtensionClient,
		MeCalendarView:                                        meCalendarViewClient,
		MeCalendarViewAttachment:                              meCalendarViewAttachmentClient,
		MeCalendarViewCalendar:                                meCalendarViewCalendarClient,
		MeCalendarViewExtension:                               meCalendarViewExtensionClient,
		MeCalendarViewInstance:                                meCalendarViewInstanceClient,
		MeCalendarViewInstanceAttachment:                      meCalendarViewInstanceAttachmentClient,
		MeCalendarViewInstanceCalendar:                        meCalendarViewInstanceCalendarClient,
		MeCalendarViewInstanceExtension:                       meCalendarViewInstanceExtensionClient,
		MeChat:                                                meChatClient,
		MeChatInstalledApp:                                    meChatInstalledAppClient,
		MeChatInstalledAppTeamsApp:                            meChatInstalledAppTeamsAppClient,
		MeChatInstalledAppTeamsAppDefinition:                  meChatInstalledAppTeamsAppDefinitionClient,
		MeChatLastMessagePreview:                              meChatLastMessagePreviewClient,
		MeChatMember:                                          meChatMemberClient,
		MeChatMessage:                                         meChatMessageClient,
		MeChatMessageHostedContent:                            meChatMessageHostedContentClient,
		MeChatMessageReply:                                    meChatMessageReplyClient,
		MeChatMessageReplyHostedContent:                       meChatMessageReplyHostedContentClient,
		MeChatPermissionGrant:                                 meChatPermissionGrantClient,
		MeChatPinnedMessage:                                   meChatPinnedMessageClient,
		MeChatPinnedMessageMessage:                            meChatPinnedMessageMessageClient,
		MeChatTab:                                             meChatTabClient,
		MeChatTabTeamsApp:                                     meChatTabTeamsAppClient,
		MeContact:                                             meContactClient,
		MeContactExtension:                                    meContactExtensionClient,
		MeContactFolder:                                       meContactFolderClient,
		MeContactFolderChildFolder:                            meContactFolderChildFolderClient,
		MeContactFolderChildFolderContact:                     meContactFolderChildFolderContactClient,
		MeContactFolderChildFolderContactExtension:            meContactFolderChildFolderContactExtensionClient,
		MeContactFolderChildFolderContactPhoto:                meContactFolderChildFolderContactPhotoClient,
		MeContactFolderContact:                                meContactFolderContactClient,
		MeContactFolderContactExtension:                       meContactFolderContactExtensionClient,
		MeContactFolderContactPhoto:                           meContactFolderContactPhotoClient,
		MeContactPhoto:                                        meContactPhotoClient,
		MeCreatedObject:                                       meCreatedObjectClient,
		MeDeviceManagementTroubleshootingEvent:                meDeviceManagementTroubleshootingEventClient,
		MeDirectReport:                                        meDirectReportClient,
		MeDrive:                                               meDriveClient,
		MeEmployeeExperience:                                  meEmployeeExperienceClient,
		MeEmployeeExperienceLearningCourseActivity:            meEmployeeExperienceLearningCourseActivityClient,
		MeEvent:                                                   meEventClient,
		MeEventAttachment:                                         meEventAttachmentClient,
		MeEventCalendar:                                           meEventCalendarClient,
		MeEventExtension:                                          meEventExtensionClient,
		MeEventInstance:                                           meEventInstanceClient,
		MeEventInstanceAttachment:                                 meEventInstanceAttachmentClient,
		MeEventInstanceCalendar:                                   meEventInstanceCalendarClient,
		MeEventInstanceExtension:                                  meEventInstanceExtensionClient,
		MeExtension:                                               meExtensionClient,
		MeFollowedSite:                                            meFollowedSiteClient,
		MeInferenceClassification:                                 meInferenceClassificationClient,
		MeInferenceClassificationOverride:                         meInferenceClassificationOverrideClient,
		MeInsight:                                                 meInsightClient,
		MeInsightShared:                                           meInsightSharedClient,
		MeInsightSharedLastSharedMethod:                           meInsightSharedLastSharedMethodClient,
		MeInsightSharedResource:                                   meInsightSharedResourceClient,
		MeInsightTrending:                                         meInsightTrendingClient,
		MeInsightTrendingResource:                                 meInsightTrendingResourceClient,
		MeInsightUsed:                                             meInsightUsedClient,
		MeInsightUsedResource:                                     meInsightUsedResourceClient,
		MeJoinedTeam:                                              meJoinedTeamClient,
		MeJoinedTeamAllChannel:                                    meJoinedTeamAllChannelClient,
		MeJoinedTeamChannel:                                       meJoinedTeamChannelClient,
		MeJoinedTeamChannelFilesFolder:                            meJoinedTeamChannelFilesFolderClient,
		MeJoinedTeamChannelFilesFolderContent:                     meJoinedTeamChannelFilesFolderContentClient,
		MeJoinedTeamChannelMember:                                 meJoinedTeamChannelMemberClient,
		MeJoinedTeamChannelMessage:                                meJoinedTeamChannelMessageClient,
		MeJoinedTeamChannelMessageHostedContent:                   meJoinedTeamChannelMessageHostedContentClient,
		MeJoinedTeamChannelMessageReply:                           meJoinedTeamChannelMessageReplyClient,
		MeJoinedTeamChannelMessageReplyHostedContent:              meJoinedTeamChannelMessageReplyHostedContentClient,
		MeJoinedTeamChannelSharedWithTeam:                         meJoinedTeamChannelSharedWithTeamClient,
		MeJoinedTeamChannelSharedWithTeamAllowedMember:            meJoinedTeamChannelSharedWithTeamAllowedMemberClient,
		MeJoinedTeamChannelSharedWithTeamTeam:                     meJoinedTeamChannelSharedWithTeamTeamClient,
		MeJoinedTeamChannelTab:                                    meJoinedTeamChannelTabClient,
		MeJoinedTeamChannelTabTeamsApp:                            meJoinedTeamChannelTabTeamsAppClient,
		MeJoinedTeamGroup:                                         meJoinedTeamGroupClient,
		MeJoinedTeamIncomingChannel:                               meJoinedTeamIncomingChannelClient,
		MeJoinedTeamInstalledApp:                                  meJoinedTeamInstalledAppClient,
		MeJoinedTeamInstalledAppTeamsApp:                          meJoinedTeamInstalledAppTeamsAppClient,
		MeJoinedTeamInstalledAppTeamsAppDefinition:                meJoinedTeamInstalledAppTeamsAppDefinitionClient,
		MeJoinedTeamMember:                                        meJoinedTeamMemberClient,
		MeJoinedTeamOperation:                                     meJoinedTeamOperationClient,
		MeJoinedTeamPermissionGrant:                               meJoinedTeamPermissionGrantClient,
		MeJoinedTeamPhoto:                                         meJoinedTeamPhotoClient,
		MeJoinedTeamPrimaryChannel:                                meJoinedTeamPrimaryChannelClient,
		MeJoinedTeamPrimaryChannelFilesFolder:                     meJoinedTeamPrimaryChannelFilesFolderClient,
		MeJoinedTeamPrimaryChannelFilesFolderContent:              meJoinedTeamPrimaryChannelFilesFolderContentClient,
		MeJoinedTeamPrimaryChannelMember:                          meJoinedTeamPrimaryChannelMemberClient,
		MeJoinedTeamPrimaryChannelMessage:                         meJoinedTeamPrimaryChannelMessageClient,
		MeJoinedTeamPrimaryChannelMessageHostedContent:            meJoinedTeamPrimaryChannelMessageHostedContentClient,
		MeJoinedTeamPrimaryChannelMessageReply:                    meJoinedTeamPrimaryChannelMessageReplyClient,
		MeJoinedTeamPrimaryChannelMessageReplyHostedContent:       meJoinedTeamPrimaryChannelMessageReplyHostedContentClient,
		MeJoinedTeamPrimaryChannelSharedWithTeam:                  meJoinedTeamPrimaryChannelSharedWithTeamClient,
		MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMember:     meJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient,
		MeJoinedTeamPrimaryChannelSharedWithTeamTeam:              meJoinedTeamPrimaryChannelSharedWithTeamTeamClient,
		MeJoinedTeamPrimaryChannelTab:                             meJoinedTeamPrimaryChannelTabClient,
		MeJoinedTeamPrimaryChannelTabTeamsApp:                     meJoinedTeamPrimaryChannelTabTeamsAppClient,
		MeJoinedTeamSchedule:                                      meJoinedTeamScheduleClient,
		MeJoinedTeamScheduleOfferShiftRequest:                     meJoinedTeamScheduleOfferShiftRequestClient,
		MeJoinedTeamScheduleOpenShift:                             meJoinedTeamScheduleOpenShiftClient,
		MeJoinedTeamScheduleOpenShiftChangeRequest:                meJoinedTeamScheduleOpenShiftChangeRequestClient,
		MeJoinedTeamScheduleSchedulingGroup:                       meJoinedTeamScheduleSchedulingGroupClient,
		MeJoinedTeamScheduleShift:                                 meJoinedTeamScheduleShiftClient,
		MeJoinedTeamScheduleSwapShiftsChangeRequest:               meJoinedTeamScheduleSwapShiftsChangeRequestClient,
		MeJoinedTeamScheduleTimeOffReason:                         meJoinedTeamScheduleTimeOffReasonClient,
		MeJoinedTeamScheduleTimeOffRequest:                        meJoinedTeamScheduleTimeOffRequestClient,
		MeJoinedTeamScheduleTimesOff:                              meJoinedTeamScheduleTimesOffClient,
		MeJoinedTeamTag:                                           meJoinedTeamTagClient,
		MeJoinedTeamTagMember:                                     meJoinedTeamTagMemberClient,
		MeJoinedTeamTemplate:                                      meJoinedTeamTemplateClient,
		MeLicenseDetail:                                           meLicenseDetailClient,
		MeMailFolder:                                              meMailFolderClient,
		MeMailFolderChildFolder:                                   meMailFolderChildFolderClient,
		MeMailFolderChildFolderMessage:                            meMailFolderChildFolderMessageClient,
		MeMailFolderChildFolderMessageAttachment:                  meMailFolderChildFolderMessageAttachmentClient,
		MeMailFolderChildFolderMessageExtension:                   meMailFolderChildFolderMessageExtensionClient,
		MeMailFolderChildFolderMessageRule:                        meMailFolderChildFolderMessageRuleClient,
		MeMailFolderMessage:                                       meMailFolderMessageClient,
		MeMailFolderMessageAttachment:                             meMailFolderMessageAttachmentClient,
		MeMailFolderMessageExtension:                              meMailFolderMessageExtensionClient,
		MeMailFolderMessageRule:                                   meMailFolderMessageRuleClient,
		MeMailboxSetting:                                          meMailboxSettingClient,
		MeManagedAppRegistration:                                  meManagedAppRegistrationClient,
		MeManagedDevice:                                           meManagedDeviceClient,
		MeManagedDeviceDeviceCategory:                             meManagedDeviceDeviceCategoryClient,
		MeManagedDeviceDeviceCompliancePolicyState:                meManagedDeviceDeviceCompliancePolicyStateClient,
		MeManagedDeviceDeviceConfigurationState:                   meManagedDeviceDeviceConfigurationStateClient,
		MeManagedDeviceLogCollectionRequest:                       meManagedDeviceLogCollectionRequestClient,
		MeManagedDeviceUser:                                       meManagedDeviceUserClient,
		MeManagedDeviceWindowsProtectionState:                     meManagedDeviceWindowsProtectionStateClient,
		MeManagedDeviceWindowsProtectionStateDetectedMalwareState: meManagedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		MeManager:                     meManagerClient,
		MeMemberOf:                    meMemberOfClient,
		MeMessage:                     meMessageClient,
		MeMessageAttachment:           meMessageAttachmentClient,
		MeMessageExtension:            meMessageExtensionClient,
		MeOauth2PermissionGrant:       meOauth2PermissionGrantClient,
		MeOnenote:                     meOnenoteClient,
		MeOnenoteNotebook:             meOnenoteNotebookClient,
		MeOnenoteNotebookSection:      meOnenoteNotebookSectionClient,
		MeOnenoteNotebookSectionGroup: meOnenoteNotebookSectionGroupClient,
		MeOnenoteNotebookSectionGroupParentNotebook:            meOnenoteNotebookSectionGroupParentNotebookClient,
		MeOnenoteNotebookSectionGroupParentSectionGroup:        meOnenoteNotebookSectionGroupParentSectionGroupClient,
		MeOnenoteNotebookSectionGroupSection:                   meOnenoteNotebookSectionGroupSectionClient,
		MeOnenoteNotebookSectionGroupSectionGroup:              meOnenoteNotebookSectionGroupSectionGroupClient,
		MeOnenoteNotebookSectionGroupSectionPage:               meOnenoteNotebookSectionGroupSectionPageClient,
		MeOnenoteNotebookSectionGroupSectionPageContent:        meOnenoteNotebookSectionGroupSectionPageContentClient,
		MeOnenoteNotebookSectionGroupSectionPageParentNotebook: meOnenoteNotebookSectionGroupSectionPageParentNotebookClient,
		MeOnenoteNotebookSectionGroupSectionPageParentSection:  meOnenoteNotebookSectionGroupSectionPageParentSectionClient,
		MeOnenoteNotebookSectionGroupSectionParentNotebook:     meOnenoteNotebookSectionGroupSectionParentNotebookClient,
		MeOnenoteNotebookSectionGroupSectionParentSectionGroup: meOnenoteNotebookSectionGroupSectionParentSectionGroupClient,
		MeOnenoteNotebookSectionPage:                           meOnenoteNotebookSectionPageClient,
		MeOnenoteNotebookSectionPageContent:                    meOnenoteNotebookSectionPageContentClient,
		MeOnenoteNotebookSectionPageParentNotebook:             meOnenoteNotebookSectionPageParentNotebookClient,
		MeOnenoteNotebookSectionPageParentSection:              meOnenoteNotebookSectionPageParentSectionClient,
		MeOnenoteNotebookSectionParentNotebook:                 meOnenoteNotebookSectionParentNotebookClient,
		MeOnenoteNotebookSectionParentSectionGroup:             meOnenoteNotebookSectionParentSectionGroupClient,
		MeOnenoteOperation:                                     meOnenoteOperationClient,
		MeOnenotePage:                                          meOnenotePageClient,
		MeOnenotePageContent:                                   meOnenotePageContentClient,
		MeOnenotePageParentNotebook:                            meOnenotePageParentNotebookClient,
		MeOnenotePageParentSection:                             meOnenotePageParentSectionClient,
		MeOnenoteResource:                                      meOnenoteResourceClient,
		MeOnenoteResourceContent:                               meOnenoteResourceContentClient,
		MeOnenoteSection:                                       meOnenoteSectionClient,
		MeOnenoteSectionGroup:                                  meOnenoteSectionGroupClient,
		MeOnenoteSectionGroupParentNotebook:                    meOnenoteSectionGroupParentNotebookClient,
		MeOnenoteSectionGroupParentSectionGroup:                meOnenoteSectionGroupParentSectionGroupClient,
		MeOnenoteSectionGroupSection:                           meOnenoteSectionGroupSectionClient,
		MeOnenoteSectionGroupSectionGroup:                      meOnenoteSectionGroupSectionGroupClient,
		MeOnenoteSectionGroupSectionPage:                       meOnenoteSectionGroupSectionPageClient,
		MeOnenoteSectionGroupSectionPageContent:                meOnenoteSectionGroupSectionPageContentClient,
		MeOnenoteSectionGroupSectionPageParentNotebook:         meOnenoteSectionGroupSectionPageParentNotebookClient,
		MeOnenoteSectionGroupSectionPageParentSection:          meOnenoteSectionGroupSectionPageParentSectionClient,
		MeOnenoteSectionGroupSectionParentNotebook:             meOnenoteSectionGroupSectionParentNotebookClient,
		MeOnenoteSectionGroupSectionParentSectionGroup:         meOnenoteSectionGroupSectionParentSectionGroupClient,
		MeOnenoteSectionPage:                                   meOnenoteSectionPageClient,
		MeOnenoteSectionPageContent:                            meOnenoteSectionPageContentClient,
		MeOnenoteSectionPageParentNotebook:                     meOnenoteSectionPageParentNotebookClient,
		MeOnenoteSectionPageParentSection:                      meOnenoteSectionPageParentSectionClient,
		MeOnenoteSectionParentNotebook:                         meOnenoteSectionParentNotebookClient,
		MeOnenoteSectionParentSectionGroup:                     meOnenoteSectionParentSectionGroupClient,
		MeOnlineMeeting:                                        meOnlineMeetingClient,
		MeOnlineMeetingAttendanceReport:                        meOnlineMeetingAttendanceReportClient,
		MeOnlineMeetingAttendanceReportAttendanceRecord:        meOnlineMeetingAttendanceReportAttendanceRecordClient,
		MeOnlineMeetingAttendeeReport:                          meOnlineMeetingAttendeeReportClient,
		MeOutlook:                                              meOutlookClient,
		MeOutlookMasterCategory:                                meOutlookMasterCategoryClient,
		MeOwnedDevice:                                          meOwnedDeviceClient,
		MeOwnedObject:                                          meOwnedObjectClient,
		MePeople:                                               mePeopleClient,
		MePhoto:                                                mePhotoClient,
		MePlanner:                                              mePlannerClient,
		MePlannerPlan:                                          mePlannerPlanClient,
		MePlannerPlanBucket:                                    mePlannerPlanBucketClient,
		MePlannerPlanBucketTask:                                mePlannerPlanBucketTaskClient,
		MePlannerPlanBucketTaskAssignedToTaskBoardFormat:       mePlannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		MePlannerPlanBucketTaskBucketTaskBoardFormat:           mePlannerPlanBucketTaskBucketTaskBoardFormatClient,
		MePlannerPlanBucketTaskDetail:                          mePlannerPlanBucketTaskDetailClient,
		MePlannerPlanBucketTaskProgressTaskBoardFormat:         mePlannerPlanBucketTaskProgressTaskBoardFormatClient,
		MePlannerPlanDetail:                                    mePlannerPlanDetailClient,
		MePlannerPlanTask:                                      mePlannerPlanTaskClient,
		MePlannerPlanTaskAssignedToTaskBoardFormat:             mePlannerPlanTaskAssignedToTaskBoardFormatClient,
		MePlannerPlanTaskBucketTaskBoardFormat:                 mePlannerPlanTaskBucketTaskBoardFormatClient,
		MePlannerPlanTaskDetail:                                mePlannerPlanTaskDetailClient,
		MePlannerPlanTaskProgressTaskBoardFormat:               mePlannerPlanTaskProgressTaskBoardFormatClient,
		MePlannerTask:                                          mePlannerTaskClient,
		MePlannerTaskAssignedToTaskBoardFormat:                 mePlannerTaskAssignedToTaskBoardFormatClient,
		MePlannerTaskBucketTaskBoardFormat:                     mePlannerTaskBucketTaskBoardFormatClient,
		MePlannerTaskDetail:                                    mePlannerTaskDetailClient,
		MePlannerTaskProgressTaskBoardFormat:                   mePlannerTaskProgressTaskBoardFormatClient,
		MePresence:                                             mePresenceClient,
		MeRegisteredDevice:                                     meRegisteredDeviceClient,
		MeScopedRoleMemberOf:                                   meScopedRoleMemberOfClient,
		MeSetting:                                              meSettingClient,
		MeSettingShiftPreference:                               meSettingShiftPreferenceClient,
		MeTeamwork:                                             meTeamworkClient,
		MeTeamworkAssociatedTeam:                               meTeamworkAssociatedTeamClient,
		MeTeamworkAssociatedTeamTeam:                           meTeamworkAssociatedTeamTeamClient,
		MeTeamworkInstalledApp:                                 meTeamworkInstalledAppClient,
		MeTeamworkInstalledAppChat:                             meTeamworkInstalledAppChatClient,
		MeTeamworkInstalledAppTeamsApp:                         meTeamworkInstalledAppTeamsAppClient,
		MeTeamworkInstalledAppTeamsAppDefinition:               meTeamworkInstalledAppTeamsAppDefinitionClient,
		MeTodo:                                                 meTodoClient,
		MeTodoList:                                             meTodoListClient,
		MeTodoListExtension:                                    meTodoListExtensionClient,
		MeTodoListTask:                                         meTodoListTaskClient,
		MeTodoListTaskAttachment:                               meTodoListTaskAttachmentClient,
		MeTodoListTaskAttachmentSession:                        meTodoListTaskAttachmentSessionClient,
		MeTodoListTaskAttachmentSessionContent:                 meTodoListTaskAttachmentSessionContentClient,
		MeTodoListTaskChecklistItem:                            meTodoListTaskChecklistItemClient,
		MeTodoListTaskExtension:                                meTodoListTaskExtensionClient,
		MeTodoListTaskLinkedResource:                           meTodoListTaskLinkedResourceClient,
		MeTransitiveMemberOf:                                   meTransitiveMemberOfClient,
	}, nil
}
