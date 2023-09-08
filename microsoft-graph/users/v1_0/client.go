package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/user"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useractivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useractivityhistoryitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useractivityhistoryitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useragreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationemailmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationfido2method"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationpasswordmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationphonemethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationsoftwareoathmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationtemporaryaccesspassmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationwindowshelloforbusinessmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userauthenticationwindowshelloforbusinessmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendargroupcalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatlastmessagepreview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatpermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatpinnedmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchatpinnedmessagemessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchattab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userchattabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfolderchildfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfolderchildfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfolderchildfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usercreatedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userdevicemanagementtroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userdirectreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useremployeeexperience"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useremployeeexperiencelearningcourseactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usereventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usereventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usereventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usereventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usereventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usereventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usereventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userfollowedsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinferenceclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinferenceclassificationoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsightshared"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsightsharedlastsharedmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsightsharedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsighttrending"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsighttrendingresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsightused"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userinsightusedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamallchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamchanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamincomingchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteaminstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteaminstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteaminstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteammember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteampermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamprimarychanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduleoffershiftrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduleopenshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduleopenshiftchangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduleschedulinggroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduleshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduleswapshiftschangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduletimeoffreason"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduletimeoffrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamscheduletimesoff"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamtag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamtagmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userjoinedteamtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userlicensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfolderchildfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfolderchildfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfolderchildfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfolderchildfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermailfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanagedappregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermanager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermemberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useroauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotenotebooksectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenoteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotepagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotepageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotepageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenoteresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenoteresourcecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronenotesectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronlinemeeting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronlinemeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronlinemeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useronlinemeetingattendeereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useroutlook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/useroutlookmastercategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userowneddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userpeople"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplanner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannertaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannertaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannertaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userplannertaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userpresence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userregistereddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userscopedrolememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usersetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usersettingshiftpreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userteamwork"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userteamworkassociatedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userteamworkassociatedteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userteamworkinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userteamworkinstalledappchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userteamworkinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/userteamworkinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolistextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolisttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolisttaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolisttaskattachmentsession"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolisttaskattachmentsessioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolisttaskchecklistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolisttaskextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertodolisttasklinkedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/v1_0/usertransitivememberof"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	User                                                        *user.UserClient
	UserActivity                                                *useractivity.UserActivityClient
	UserActivityHistoryItem                                     *useractivityhistoryitem.UserActivityHistoryItemClient
	UserActivityHistoryItemActivity                             *useractivityhistoryitemactivity.UserActivityHistoryItemActivityClient
	UserAgreementAcceptance                                     *useragreementacceptance.UserAgreementAcceptanceClient
	UserAppRoleAssignment                                       *userapproleassignment.UserAppRoleAssignmentClient
	UserAuthentication                                          *userauthentication.UserAuthenticationClient
	UserAuthenticationEmailMethod                               *userauthenticationemailmethod.UserAuthenticationEmailMethodClient
	UserAuthenticationFido2Method                               *userauthenticationfido2method.UserAuthenticationFido2MethodClient
	UserAuthenticationMethod                                    *userauthenticationmethod.UserAuthenticationMethodClient
	UserAuthenticationMicrosoftAuthenticatorMethod              *userauthenticationmicrosoftauthenticatormethod.UserAuthenticationMicrosoftAuthenticatorMethodClient
	UserAuthenticationMicrosoftAuthenticatorMethodDevice        *userauthenticationmicrosoftauthenticatormethoddevice.UserAuthenticationMicrosoftAuthenticatorMethodDeviceClient
	UserAuthenticationOperation                                 *userauthenticationoperation.UserAuthenticationOperationClient
	UserAuthenticationPasswordMethod                            *userauthenticationpasswordmethod.UserAuthenticationPasswordMethodClient
	UserAuthenticationPhoneMethod                               *userauthenticationphonemethod.UserAuthenticationPhoneMethodClient
	UserAuthenticationSoftwareOathMethod                        *userauthenticationsoftwareoathmethod.UserAuthenticationSoftwareOathMethodClient
	UserAuthenticationTemporaryAccessPassMethod                 *userauthenticationtemporaryaccesspassmethod.UserAuthenticationTemporaryAccessPassMethodClient
	UserAuthenticationWindowsHelloForBusinessMethod             *userauthenticationwindowshelloforbusinessmethod.UserAuthenticationWindowsHelloForBusinessMethodClient
	UserAuthenticationWindowsHelloForBusinessMethodDevice       *userauthenticationwindowshelloforbusinessmethoddevice.UserAuthenticationWindowsHelloForBusinessMethodDeviceClient
	UserCalendar                                                *usercalendar.UserCalendarClient
	UserCalendarCalendarPermission                              *usercalendarcalendarpermission.UserCalendarCalendarPermissionClient
	UserCalendarCalendarView                                    *usercalendarcalendarview.UserCalendarCalendarViewClient
	UserCalendarCalendarViewAttachment                          *usercalendarcalendarviewattachment.UserCalendarCalendarViewAttachmentClient
	UserCalendarCalendarViewCalendar                            *usercalendarcalendarviewcalendar.UserCalendarCalendarViewCalendarClient
	UserCalendarCalendarViewExtension                           *usercalendarcalendarviewextension.UserCalendarCalendarViewExtensionClient
	UserCalendarCalendarViewInstance                            *usercalendarcalendarviewinstance.UserCalendarCalendarViewInstanceClient
	UserCalendarCalendarViewInstanceAttachment                  *usercalendarcalendarviewinstanceattachment.UserCalendarCalendarViewInstanceAttachmentClient
	UserCalendarCalendarViewInstanceCalendar                    *usercalendarcalendarviewinstancecalendar.UserCalendarCalendarViewInstanceCalendarClient
	UserCalendarCalendarViewInstanceExtension                   *usercalendarcalendarviewinstanceextension.UserCalendarCalendarViewInstanceExtensionClient
	UserCalendarEvent                                           *usercalendarevent.UserCalendarEventClient
	UserCalendarEventAttachment                                 *usercalendareventattachment.UserCalendarEventAttachmentClient
	UserCalendarEventCalendar                                   *usercalendareventcalendar.UserCalendarEventCalendarClient
	UserCalendarEventExtension                                  *usercalendareventextension.UserCalendarEventExtensionClient
	UserCalendarEventInstance                                   *usercalendareventinstance.UserCalendarEventInstanceClient
	UserCalendarEventInstanceAttachment                         *usercalendareventinstanceattachment.UserCalendarEventInstanceAttachmentClient
	UserCalendarEventInstanceCalendar                           *usercalendareventinstancecalendar.UserCalendarEventInstanceCalendarClient
	UserCalendarEventInstanceExtension                          *usercalendareventinstanceextension.UserCalendarEventInstanceExtensionClient
	UserCalendarGroup                                           *usercalendargroup.UserCalendarGroupClient
	UserCalendarGroupCalendar                                   *usercalendargroupcalendar.UserCalendarGroupCalendarClient
	UserCalendarGroupCalendarCalendarPermission                 *usercalendargroupcalendarcalendarpermission.UserCalendarGroupCalendarCalendarPermissionClient
	UserCalendarGroupCalendarCalendarView                       *usercalendargroupcalendarcalendarview.UserCalendarGroupCalendarCalendarViewClient
	UserCalendarGroupCalendarCalendarViewAttachment             *usercalendargroupcalendarcalendarviewattachment.UserCalendarGroupCalendarCalendarViewAttachmentClient
	UserCalendarGroupCalendarCalendarViewCalendar               *usercalendargroupcalendarcalendarviewcalendar.UserCalendarGroupCalendarCalendarViewCalendarClient
	UserCalendarGroupCalendarCalendarViewExtension              *usercalendargroupcalendarcalendarviewextension.UserCalendarGroupCalendarCalendarViewExtensionClient
	UserCalendarGroupCalendarCalendarViewInstance               *usercalendargroupcalendarcalendarviewinstance.UserCalendarGroupCalendarCalendarViewInstanceClient
	UserCalendarGroupCalendarCalendarViewInstanceAttachment     *usercalendargroupcalendarcalendarviewinstanceattachment.UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient
	UserCalendarGroupCalendarCalendarViewInstanceCalendar       *usercalendargroupcalendarcalendarviewinstancecalendar.UserCalendarGroupCalendarCalendarViewInstanceCalendarClient
	UserCalendarGroupCalendarCalendarViewInstanceExtension      *usercalendargroupcalendarcalendarviewinstanceextension.UserCalendarGroupCalendarCalendarViewInstanceExtensionClient
	UserCalendarGroupCalendarEvent                              *usercalendargroupcalendarevent.UserCalendarGroupCalendarEventClient
	UserCalendarGroupCalendarEventAttachment                    *usercalendargroupcalendareventattachment.UserCalendarGroupCalendarEventAttachmentClient
	UserCalendarGroupCalendarEventCalendar                      *usercalendargroupcalendareventcalendar.UserCalendarGroupCalendarEventCalendarClient
	UserCalendarGroupCalendarEventExtension                     *usercalendargroupcalendareventextension.UserCalendarGroupCalendarEventExtensionClient
	UserCalendarGroupCalendarEventInstance                      *usercalendargroupcalendareventinstance.UserCalendarGroupCalendarEventInstanceClient
	UserCalendarGroupCalendarEventInstanceAttachment            *usercalendargroupcalendareventinstanceattachment.UserCalendarGroupCalendarEventInstanceAttachmentClient
	UserCalendarGroupCalendarEventInstanceCalendar              *usercalendargroupcalendareventinstancecalendar.UserCalendarGroupCalendarEventInstanceCalendarClient
	UserCalendarGroupCalendarEventInstanceExtension             *usercalendargroupcalendareventinstanceextension.UserCalendarGroupCalendarEventInstanceExtensionClient
	UserCalendarView                                            *usercalendarview.UserCalendarViewClient
	UserCalendarViewAttachment                                  *usercalendarviewattachment.UserCalendarViewAttachmentClient
	UserCalendarViewCalendar                                    *usercalendarviewcalendar.UserCalendarViewCalendarClient
	UserCalendarViewExtension                                   *usercalendarviewextension.UserCalendarViewExtensionClient
	UserCalendarViewInstance                                    *usercalendarviewinstance.UserCalendarViewInstanceClient
	UserCalendarViewInstanceAttachment                          *usercalendarviewinstanceattachment.UserCalendarViewInstanceAttachmentClient
	UserCalendarViewInstanceCalendar                            *usercalendarviewinstancecalendar.UserCalendarViewInstanceCalendarClient
	UserCalendarViewInstanceExtension                           *usercalendarviewinstanceextension.UserCalendarViewInstanceExtensionClient
	UserChat                                                    *userchat.UserChatClient
	UserChatInstalledApp                                        *userchatinstalledapp.UserChatInstalledAppClient
	UserChatInstalledAppTeamsApp                                *userchatinstalledappteamsapp.UserChatInstalledAppTeamsAppClient
	UserChatInstalledAppTeamsAppDefinition                      *userchatinstalledappteamsappdefinition.UserChatInstalledAppTeamsAppDefinitionClient
	UserChatLastMessagePreview                                  *userchatlastmessagepreview.UserChatLastMessagePreviewClient
	UserChatMember                                              *userchatmember.UserChatMemberClient
	UserChatMessage                                             *userchatmessage.UserChatMessageClient
	UserChatMessageHostedContent                                *userchatmessagehostedcontent.UserChatMessageHostedContentClient
	UserChatMessageReply                                        *userchatmessagereply.UserChatMessageReplyClient
	UserChatMessageReplyHostedContent                           *userchatmessagereplyhostedcontent.UserChatMessageReplyHostedContentClient
	UserChatPermissionGrant                                     *userchatpermissiongrant.UserChatPermissionGrantClient
	UserChatPinnedMessage                                       *userchatpinnedmessage.UserChatPinnedMessageClient
	UserChatPinnedMessageMessage                                *userchatpinnedmessagemessage.UserChatPinnedMessageMessageClient
	UserChatTab                                                 *userchattab.UserChatTabClient
	UserChatTabTeamsApp                                         *userchattabteamsapp.UserChatTabTeamsAppClient
	UserContact                                                 *usercontact.UserContactClient
	UserContactExtension                                        *usercontactextension.UserContactExtensionClient
	UserContactFolder                                           *usercontactfolder.UserContactFolderClient
	UserContactFolderChildFolder                                *usercontactfolderchildfolder.UserContactFolderChildFolderClient
	UserContactFolderChildFolderContact                         *usercontactfolderchildfoldercontact.UserContactFolderChildFolderContactClient
	UserContactFolderChildFolderContactExtension                *usercontactfolderchildfoldercontactextension.UserContactFolderChildFolderContactExtensionClient
	UserContactFolderChildFolderContactPhoto                    *usercontactfolderchildfoldercontactphoto.UserContactFolderChildFolderContactPhotoClient
	UserContactFolderContact                                    *usercontactfoldercontact.UserContactFolderContactClient
	UserContactFolderContactExtension                           *usercontactfoldercontactextension.UserContactFolderContactExtensionClient
	UserContactFolderContactPhoto                               *usercontactfoldercontactphoto.UserContactFolderContactPhotoClient
	UserContactPhoto                                            *usercontactphoto.UserContactPhotoClient
	UserCreatedObject                                           *usercreatedobject.UserCreatedObjectClient
	UserDeviceManagementTroubleshootingEvent                    *userdevicemanagementtroubleshootingevent.UserDeviceManagementTroubleshootingEventClient
	UserDirectReport                                            *userdirectreport.UserDirectReportClient
	UserDrive                                                   *userdrive.UserDriveClient
	UserEmployeeExperience                                      *useremployeeexperience.UserEmployeeExperienceClient
	UserEmployeeExperienceLearningCourseActivity                *useremployeeexperiencelearningcourseactivity.UserEmployeeExperienceLearningCourseActivityClient
	UserEvent                                                   *userevent.UserEventClient
	UserEventAttachment                                         *usereventattachment.UserEventAttachmentClient
	UserEventCalendar                                           *usereventcalendar.UserEventCalendarClient
	UserEventExtension                                          *usereventextension.UserEventExtensionClient
	UserEventInstance                                           *usereventinstance.UserEventInstanceClient
	UserEventInstanceAttachment                                 *usereventinstanceattachment.UserEventInstanceAttachmentClient
	UserEventInstanceCalendar                                   *usereventinstancecalendar.UserEventInstanceCalendarClient
	UserEventInstanceExtension                                  *usereventinstanceextension.UserEventInstanceExtensionClient
	UserExtension                                               *userextension.UserExtensionClient
	UserFollowedSite                                            *userfollowedsite.UserFollowedSiteClient
	UserInferenceClassification                                 *userinferenceclassification.UserInferenceClassificationClient
	UserInferenceClassificationOverride                         *userinferenceclassificationoverride.UserInferenceClassificationOverrideClient
	UserInsight                                                 *userinsight.UserInsightClient
	UserInsightShared                                           *userinsightshared.UserInsightSharedClient
	UserInsightSharedLastSharedMethod                           *userinsightsharedlastsharedmethod.UserInsightSharedLastSharedMethodClient
	UserInsightSharedResource                                   *userinsightsharedresource.UserInsightSharedResourceClient
	UserInsightTrending                                         *userinsighttrending.UserInsightTrendingClient
	UserInsightTrendingResource                                 *userinsighttrendingresource.UserInsightTrendingResourceClient
	UserInsightUsed                                             *userinsightused.UserInsightUsedClient
	UserInsightUsedResource                                     *userinsightusedresource.UserInsightUsedResourceClient
	UserJoinedTeam                                              *userjoinedteam.UserJoinedTeamClient
	UserJoinedTeamAllChannel                                    *userjoinedteamallchannel.UserJoinedTeamAllChannelClient
	UserJoinedTeamChannel                                       *userjoinedteamchannel.UserJoinedTeamChannelClient
	UserJoinedTeamChannelFilesFolder                            *userjoinedteamchannelfilesfolder.UserJoinedTeamChannelFilesFolderClient
	UserJoinedTeamChannelFilesFolderContent                     *userjoinedteamchannelfilesfoldercontent.UserJoinedTeamChannelFilesFolderContentClient
	UserJoinedTeamChannelMember                                 *userjoinedteamchannelmember.UserJoinedTeamChannelMemberClient
	UserJoinedTeamChannelMessage                                *userjoinedteamchannelmessage.UserJoinedTeamChannelMessageClient
	UserJoinedTeamChannelMessageHostedContent                   *userjoinedteamchannelmessagehostedcontent.UserJoinedTeamChannelMessageHostedContentClient
	UserJoinedTeamChannelMessageReply                           *userjoinedteamchannelmessagereply.UserJoinedTeamChannelMessageReplyClient
	UserJoinedTeamChannelMessageReplyHostedContent              *userjoinedteamchannelmessagereplyhostedcontent.UserJoinedTeamChannelMessageReplyHostedContentClient
	UserJoinedTeamChannelSharedWithTeam                         *userjoinedteamchannelsharedwithteam.UserJoinedTeamChannelSharedWithTeamClient
	UserJoinedTeamChannelSharedWithTeamAllowedMember            *userjoinedteamchannelsharedwithteamallowedmember.UserJoinedTeamChannelSharedWithTeamAllowedMemberClient
	UserJoinedTeamChannelSharedWithTeamTeam                     *userjoinedteamchannelsharedwithteamteam.UserJoinedTeamChannelSharedWithTeamTeamClient
	UserJoinedTeamChannelTab                                    *userjoinedteamchanneltab.UserJoinedTeamChannelTabClient
	UserJoinedTeamChannelTabTeamsApp                            *userjoinedteamchanneltabteamsapp.UserJoinedTeamChannelTabTeamsAppClient
	UserJoinedTeamGroup                                         *userjoinedteamgroup.UserJoinedTeamGroupClient
	UserJoinedTeamIncomingChannel                               *userjoinedteamincomingchannel.UserJoinedTeamIncomingChannelClient
	UserJoinedTeamInstalledApp                                  *userjoinedteaminstalledapp.UserJoinedTeamInstalledAppClient
	UserJoinedTeamInstalledAppTeamsApp                          *userjoinedteaminstalledappteamsapp.UserJoinedTeamInstalledAppTeamsAppClient
	UserJoinedTeamInstalledAppTeamsAppDefinition                *userjoinedteaminstalledappteamsappdefinition.UserJoinedTeamInstalledAppTeamsAppDefinitionClient
	UserJoinedTeamMember                                        *userjoinedteammember.UserJoinedTeamMemberClient
	UserJoinedTeamOperation                                     *userjoinedteamoperation.UserJoinedTeamOperationClient
	UserJoinedTeamPermissionGrant                               *userjoinedteampermissiongrant.UserJoinedTeamPermissionGrantClient
	UserJoinedTeamPhoto                                         *userjoinedteamphoto.UserJoinedTeamPhotoClient
	UserJoinedTeamPrimaryChannel                                *userjoinedteamprimarychannel.UserJoinedTeamPrimaryChannelClient
	UserJoinedTeamPrimaryChannelFilesFolder                     *userjoinedteamprimarychannelfilesfolder.UserJoinedTeamPrimaryChannelFilesFolderClient
	UserJoinedTeamPrimaryChannelFilesFolderContent              *userjoinedteamprimarychannelfilesfoldercontent.UserJoinedTeamPrimaryChannelFilesFolderContentClient
	UserJoinedTeamPrimaryChannelMember                          *userjoinedteamprimarychannelmember.UserJoinedTeamPrimaryChannelMemberClient
	UserJoinedTeamPrimaryChannelMessage                         *userjoinedteamprimarychannelmessage.UserJoinedTeamPrimaryChannelMessageClient
	UserJoinedTeamPrimaryChannelMessageHostedContent            *userjoinedteamprimarychannelmessagehostedcontent.UserJoinedTeamPrimaryChannelMessageHostedContentClient
	UserJoinedTeamPrimaryChannelMessageReply                    *userjoinedteamprimarychannelmessagereply.UserJoinedTeamPrimaryChannelMessageReplyClient
	UserJoinedTeamPrimaryChannelMessageReplyHostedContent       *userjoinedteamprimarychannelmessagereplyhostedcontent.UserJoinedTeamPrimaryChannelMessageReplyHostedContentClient
	UserJoinedTeamPrimaryChannelSharedWithTeam                  *userjoinedteamprimarychannelsharedwithteam.UserJoinedTeamPrimaryChannelSharedWithTeamClient
	UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMember     *userjoinedteamprimarychannelsharedwithteamallowedmember.UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient
	UserJoinedTeamPrimaryChannelSharedWithTeamTeam              *userjoinedteamprimarychannelsharedwithteamteam.UserJoinedTeamPrimaryChannelSharedWithTeamTeamClient
	UserJoinedTeamPrimaryChannelTab                             *userjoinedteamprimarychanneltab.UserJoinedTeamPrimaryChannelTabClient
	UserJoinedTeamPrimaryChannelTabTeamsApp                     *userjoinedteamprimarychanneltabteamsapp.UserJoinedTeamPrimaryChannelTabTeamsAppClient
	UserJoinedTeamSchedule                                      *userjoinedteamschedule.UserJoinedTeamScheduleClient
	UserJoinedTeamScheduleOfferShiftRequest                     *userjoinedteamscheduleoffershiftrequest.UserJoinedTeamScheduleOfferShiftRequestClient
	UserJoinedTeamScheduleOpenShift                             *userjoinedteamscheduleopenshift.UserJoinedTeamScheduleOpenShiftClient
	UserJoinedTeamScheduleOpenShiftChangeRequest                *userjoinedteamscheduleopenshiftchangerequest.UserJoinedTeamScheduleOpenShiftChangeRequestClient
	UserJoinedTeamScheduleSchedulingGroup                       *userjoinedteamscheduleschedulinggroup.UserJoinedTeamScheduleSchedulingGroupClient
	UserJoinedTeamScheduleShift                                 *userjoinedteamscheduleshift.UserJoinedTeamScheduleShiftClient
	UserJoinedTeamScheduleSwapShiftsChangeRequest               *userjoinedteamscheduleswapshiftschangerequest.UserJoinedTeamScheduleSwapShiftsChangeRequestClient
	UserJoinedTeamScheduleTimeOffReason                         *userjoinedteamscheduletimeoffreason.UserJoinedTeamScheduleTimeOffReasonClient
	UserJoinedTeamScheduleTimeOffRequest                        *userjoinedteamscheduletimeoffrequest.UserJoinedTeamScheduleTimeOffRequestClient
	UserJoinedTeamScheduleTimesOff                              *userjoinedteamscheduletimesoff.UserJoinedTeamScheduleTimesOffClient
	UserJoinedTeamTag                                           *userjoinedteamtag.UserJoinedTeamTagClient
	UserJoinedTeamTagMember                                     *userjoinedteamtagmember.UserJoinedTeamTagMemberClient
	UserJoinedTeamTemplate                                      *userjoinedteamtemplate.UserJoinedTeamTemplateClient
	UserLicenseDetail                                           *userlicensedetail.UserLicenseDetailClient
	UserMailFolder                                              *usermailfolder.UserMailFolderClient
	UserMailFolderChildFolder                                   *usermailfolderchildfolder.UserMailFolderChildFolderClient
	UserMailFolderChildFolderMessage                            *usermailfolderchildfoldermessage.UserMailFolderChildFolderMessageClient
	UserMailFolderChildFolderMessageAttachment                  *usermailfolderchildfoldermessageattachment.UserMailFolderChildFolderMessageAttachmentClient
	UserMailFolderChildFolderMessageExtension                   *usermailfolderchildfoldermessageextension.UserMailFolderChildFolderMessageExtensionClient
	UserMailFolderChildFolderMessageRule                        *usermailfolderchildfoldermessagerule.UserMailFolderChildFolderMessageRuleClient
	UserMailFolderMessage                                       *usermailfoldermessage.UserMailFolderMessageClient
	UserMailFolderMessageAttachment                             *usermailfoldermessageattachment.UserMailFolderMessageAttachmentClient
	UserMailFolderMessageExtension                              *usermailfoldermessageextension.UserMailFolderMessageExtensionClient
	UserMailFolderMessageRule                                   *usermailfoldermessagerule.UserMailFolderMessageRuleClient
	UserMailboxSetting                                          *usermailboxsetting.UserMailboxSettingClient
	UserManagedAppRegistration                                  *usermanagedappregistration.UserManagedAppRegistrationClient
	UserManagedDevice                                           *usermanageddevice.UserManagedDeviceClient
	UserManagedDeviceDeviceCategory                             *usermanageddevicedevicecategory.UserManagedDeviceDeviceCategoryClient
	UserManagedDeviceDeviceCompliancePolicyState                *usermanageddevicedevicecompliancepolicystate.UserManagedDeviceDeviceCompliancePolicyStateClient
	UserManagedDeviceDeviceConfigurationState                   *usermanageddevicedeviceconfigurationstate.UserManagedDeviceDeviceConfigurationStateClient
	UserManagedDeviceLogCollectionRequest                       *usermanageddevicelogcollectionrequest.UserManagedDeviceLogCollectionRequestClient
	UserManagedDeviceUser                                       *usermanageddeviceuser.UserManagedDeviceUserClient
	UserManagedDeviceWindowsProtectionState                     *usermanageddevicewindowsprotectionstate.UserManagedDeviceWindowsProtectionStateClient
	UserManagedDeviceWindowsProtectionStateDetectedMalwareState *usermanageddevicewindowsprotectionstatedetectedmalwarestate.UserManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	UserManager                                                 *usermanager.UserManagerClient
	UserMemberOf                                                *usermemberof.UserMemberOfClient
	UserMessage                                                 *usermessage.UserMessageClient
	UserMessageAttachment                                       *usermessageattachment.UserMessageAttachmentClient
	UserMessageExtension                                        *usermessageextension.UserMessageExtensionClient
	UserOauth2PermissionGrant                                   *useroauth2permissiongrant.UserOauth2PermissionGrantClient
	UserOnenote                                                 *useronenote.UserOnenoteClient
	UserOnenoteNotebook                                         *useronenotenotebook.UserOnenoteNotebookClient
	UserOnenoteNotebookSection                                  *useronenotenotebooksection.UserOnenoteNotebookSectionClient
	UserOnenoteNotebookSectionGroup                             *useronenotenotebooksectiongroup.UserOnenoteNotebookSectionGroupClient
	UserOnenoteNotebookSectionGroupParentNotebook               *useronenotenotebooksectiongroupparentnotebook.UserOnenoteNotebookSectionGroupParentNotebookClient
	UserOnenoteNotebookSectionGroupParentSectionGroup           *useronenotenotebooksectiongroupparentsectiongroup.UserOnenoteNotebookSectionGroupParentSectionGroupClient
	UserOnenoteNotebookSectionGroupSection                      *useronenotenotebooksectiongroupsection.UserOnenoteNotebookSectionGroupSectionClient
	UserOnenoteNotebookSectionGroupSectionGroup                 *useronenotenotebooksectiongroupsectiongroup.UserOnenoteNotebookSectionGroupSectionGroupClient
	UserOnenoteNotebookSectionGroupSectionPage                  *useronenotenotebooksectiongroupsectionpage.UserOnenoteNotebookSectionGroupSectionPageClient
	UserOnenoteNotebookSectionGroupSectionPageContent           *useronenotenotebooksectiongroupsectionpagecontent.UserOnenoteNotebookSectionGroupSectionPageContentClient
	UserOnenoteNotebookSectionGroupSectionPageParentNotebook    *useronenotenotebooksectiongroupsectionpageparentnotebook.UserOnenoteNotebookSectionGroupSectionPageParentNotebookClient
	UserOnenoteNotebookSectionGroupSectionPageParentSection     *useronenotenotebooksectiongroupsectionpageparentsection.UserOnenoteNotebookSectionGroupSectionPageParentSectionClient
	UserOnenoteNotebookSectionGroupSectionParentNotebook        *useronenotenotebooksectiongroupsectionparentnotebook.UserOnenoteNotebookSectionGroupSectionParentNotebookClient
	UserOnenoteNotebookSectionGroupSectionParentSectionGroup    *useronenotenotebooksectiongroupsectionparentsectiongroup.UserOnenoteNotebookSectionGroupSectionParentSectionGroupClient
	UserOnenoteNotebookSectionPage                              *useronenotenotebooksectionpage.UserOnenoteNotebookSectionPageClient
	UserOnenoteNotebookSectionPageContent                       *useronenotenotebooksectionpagecontent.UserOnenoteNotebookSectionPageContentClient
	UserOnenoteNotebookSectionPageParentNotebook                *useronenotenotebooksectionpageparentnotebook.UserOnenoteNotebookSectionPageParentNotebookClient
	UserOnenoteNotebookSectionPageParentSection                 *useronenotenotebooksectionpageparentsection.UserOnenoteNotebookSectionPageParentSectionClient
	UserOnenoteNotebookSectionParentNotebook                    *useronenotenotebooksectionparentnotebook.UserOnenoteNotebookSectionParentNotebookClient
	UserOnenoteNotebookSectionParentSectionGroup                *useronenotenotebooksectionparentsectiongroup.UserOnenoteNotebookSectionParentSectionGroupClient
	UserOnenoteOperation                                        *useronenoteoperation.UserOnenoteOperationClient
	UserOnenotePage                                             *useronenotepage.UserOnenotePageClient
	UserOnenotePageContent                                      *useronenotepagecontent.UserOnenotePageContentClient
	UserOnenotePageParentNotebook                               *useronenotepageparentnotebook.UserOnenotePageParentNotebookClient
	UserOnenotePageParentSection                                *useronenotepageparentsection.UserOnenotePageParentSectionClient
	UserOnenoteResource                                         *useronenoteresource.UserOnenoteResourceClient
	UserOnenoteResourceContent                                  *useronenoteresourcecontent.UserOnenoteResourceContentClient
	UserOnenoteSection                                          *useronenotesection.UserOnenoteSectionClient
	UserOnenoteSectionGroup                                     *useronenotesectiongroup.UserOnenoteSectionGroupClient
	UserOnenoteSectionGroupParentNotebook                       *useronenotesectiongroupparentnotebook.UserOnenoteSectionGroupParentNotebookClient
	UserOnenoteSectionGroupParentSectionGroup                   *useronenotesectiongroupparentsectiongroup.UserOnenoteSectionGroupParentSectionGroupClient
	UserOnenoteSectionGroupSection                              *useronenotesectiongroupsection.UserOnenoteSectionGroupSectionClient
	UserOnenoteSectionGroupSectionGroup                         *useronenotesectiongroupsectiongroup.UserOnenoteSectionGroupSectionGroupClient
	UserOnenoteSectionGroupSectionPage                          *useronenotesectiongroupsectionpage.UserOnenoteSectionGroupSectionPageClient
	UserOnenoteSectionGroupSectionPageContent                   *useronenotesectiongroupsectionpagecontent.UserOnenoteSectionGroupSectionPageContentClient
	UserOnenoteSectionGroupSectionPageParentNotebook            *useronenotesectiongroupsectionpageparentnotebook.UserOnenoteSectionGroupSectionPageParentNotebookClient
	UserOnenoteSectionGroupSectionPageParentSection             *useronenotesectiongroupsectionpageparentsection.UserOnenoteSectionGroupSectionPageParentSectionClient
	UserOnenoteSectionGroupSectionParentNotebook                *useronenotesectiongroupsectionparentnotebook.UserOnenoteSectionGroupSectionParentNotebookClient
	UserOnenoteSectionGroupSectionParentSectionGroup            *useronenotesectiongroupsectionparentsectiongroup.UserOnenoteSectionGroupSectionParentSectionGroupClient
	UserOnenoteSectionPage                                      *useronenotesectionpage.UserOnenoteSectionPageClient
	UserOnenoteSectionPageContent                               *useronenotesectionpagecontent.UserOnenoteSectionPageContentClient
	UserOnenoteSectionPageParentNotebook                        *useronenotesectionpageparentnotebook.UserOnenoteSectionPageParentNotebookClient
	UserOnenoteSectionPageParentSection                         *useronenotesectionpageparentsection.UserOnenoteSectionPageParentSectionClient
	UserOnenoteSectionParentNotebook                            *useronenotesectionparentnotebook.UserOnenoteSectionParentNotebookClient
	UserOnenoteSectionParentSectionGroup                        *useronenotesectionparentsectiongroup.UserOnenoteSectionParentSectionGroupClient
	UserOnlineMeeting                                           *useronlinemeeting.UserOnlineMeetingClient
	UserOnlineMeetingAttendanceReport                           *useronlinemeetingattendancereport.UserOnlineMeetingAttendanceReportClient
	UserOnlineMeetingAttendanceReportAttendanceRecord           *useronlinemeetingattendancereportattendancerecord.UserOnlineMeetingAttendanceReportAttendanceRecordClient
	UserOnlineMeetingAttendeeReport                             *useronlinemeetingattendeereport.UserOnlineMeetingAttendeeReportClient
	UserOutlook                                                 *useroutlook.UserOutlookClient
	UserOutlookMasterCategory                                   *useroutlookmastercategory.UserOutlookMasterCategoryClient
	UserOwnedDevice                                             *userowneddevice.UserOwnedDeviceClient
	UserOwnedObject                                             *userownedobject.UserOwnedObjectClient
	UserPeople                                                  *userpeople.UserPeopleClient
	UserPhoto                                                   *userphoto.UserPhotoClient
	UserPlanner                                                 *userplanner.UserPlannerClient
	UserPlannerPlan                                             *userplannerplan.UserPlannerPlanClient
	UserPlannerPlanBucket                                       *userplannerplanbucket.UserPlannerPlanBucketClient
	UserPlannerPlanBucketTask                                   *userplannerplanbuckettask.UserPlannerPlanBucketTaskClient
	UserPlannerPlanBucketTaskAssignedToTaskBoardFormat          *userplannerplanbuckettaskassignedtotaskboardformat.UserPlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	UserPlannerPlanBucketTaskBucketTaskBoardFormat              *userplannerplanbuckettaskbuckettaskboardformat.UserPlannerPlanBucketTaskBucketTaskBoardFormatClient
	UserPlannerPlanBucketTaskDetail                             *userplannerplanbuckettaskdetail.UserPlannerPlanBucketTaskDetailClient
	UserPlannerPlanBucketTaskProgressTaskBoardFormat            *userplannerplanbuckettaskprogresstaskboardformat.UserPlannerPlanBucketTaskProgressTaskBoardFormatClient
	UserPlannerPlanDetail                                       *userplannerplandetail.UserPlannerPlanDetailClient
	UserPlannerPlanTask                                         *userplannerplantask.UserPlannerPlanTaskClient
	UserPlannerPlanTaskAssignedToTaskBoardFormat                *userplannerplantaskassignedtotaskboardformat.UserPlannerPlanTaskAssignedToTaskBoardFormatClient
	UserPlannerPlanTaskBucketTaskBoardFormat                    *userplannerplantaskbuckettaskboardformat.UserPlannerPlanTaskBucketTaskBoardFormatClient
	UserPlannerPlanTaskDetail                                   *userplannerplantaskdetail.UserPlannerPlanTaskDetailClient
	UserPlannerPlanTaskProgressTaskBoardFormat                  *userplannerplantaskprogresstaskboardformat.UserPlannerPlanTaskProgressTaskBoardFormatClient
	UserPlannerTask                                             *userplannertask.UserPlannerTaskClient
	UserPlannerTaskAssignedToTaskBoardFormat                    *userplannertaskassignedtotaskboardformat.UserPlannerTaskAssignedToTaskBoardFormatClient
	UserPlannerTaskBucketTaskBoardFormat                        *userplannertaskbuckettaskboardformat.UserPlannerTaskBucketTaskBoardFormatClient
	UserPlannerTaskDetail                                       *userplannertaskdetail.UserPlannerTaskDetailClient
	UserPlannerTaskProgressTaskBoardFormat                      *userplannertaskprogresstaskboardformat.UserPlannerTaskProgressTaskBoardFormatClient
	UserPresence                                                *userpresence.UserPresenceClient
	UserRegisteredDevice                                        *userregistereddevice.UserRegisteredDeviceClient
	UserScopedRoleMemberOf                                      *userscopedrolememberof.UserScopedRoleMemberOfClient
	UserSetting                                                 *usersetting.UserSettingClient
	UserSettingShiftPreference                                  *usersettingshiftpreference.UserSettingShiftPreferenceClient
	UserTeamwork                                                *userteamwork.UserTeamworkClient
	UserTeamworkAssociatedTeam                                  *userteamworkassociatedteam.UserTeamworkAssociatedTeamClient
	UserTeamworkAssociatedTeamTeam                              *userteamworkassociatedteamteam.UserTeamworkAssociatedTeamTeamClient
	UserTeamworkInstalledApp                                    *userteamworkinstalledapp.UserTeamworkInstalledAppClient
	UserTeamworkInstalledAppChat                                *userteamworkinstalledappchat.UserTeamworkInstalledAppChatClient
	UserTeamworkInstalledAppTeamsApp                            *userteamworkinstalledappteamsapp.UserTeamworkInstalledAppTeamsAppClient
	UserTeamworkInstalledAppTeamsAppDefinition                  *userteamworkinstalledappteamsappdefinition.UserTeamworkInstalledAppTeamsAppDefinitionClient
	UserTodo                                                    *usertodo.UserTodoClient
	UserTodoList                                                *usertodolist.UserTodoListClient
	UserTodoListExtension                                       *usertodolistextension.UserTodoListExtensionClient
	UserTodoListTask                                            *usertodolisttask.UserTodoListTaskClient
	UserTodoListTaskAttachment                                  *usertodolisttaskattachment.UserTodoListTaskAttachmentClient
	UserTodoListTaskAttachmentSession                           *usertodolisttaskattachmentsession.UserTodoListTaskAttachmentSessionClient
	UserTodoListTaskAttachmentSessionContent                    *usertodolisttaskattachmentsessioncontent.UserTodoListTaskAttachmentSessionContentClient
	UserTodoListTaskChecklistItem                               *usertodolisttaskchecklistitem.UserTodoListTaskChecklistItemClient
	UserTodoListTaskExtension                                   *usertodolisttaskextension.UserTodoListTaskExtensionClient
	UserTodoListTaskLinkedResource                              *usertodolisttasklinkedresource.UserTodoListTaskLinkedResourceClient
	UserTransitiveMemberOf                                      *usertransitivememberof.UserTransitiveMemberOfClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	userActivityClient, err := useractivity.NewUserActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserActivity client: %+v", err)
	}
	configureFunc(userActivityClient.Client)

	userActivityHistoryItemActivityClient, err := useractivityhistoryitemactivity.NewUserActivityHistoryItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserActivityHistoryItemActivity client: %+v", err)
	}
	configureFunc(userActivityHistoryItemActivityClient.Client)

	userActivityHistoryItemClient, err := useractivityhistoryitem.NewUserActivityHistoryItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserActivityHistoryItem client: %+v", err)
	}
	configureFunc(userActivityHistoryItemClient.Client)

	userAgreementAcceptanceClient, err := useragreementacceptance.NewUserAgreementAcceptanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAgreementAcceptance client: %+v", err)
	}
	configureFunc(userAgreementAcceptanceClient.Client)

	userAppRoleAssignmentClient, err := userapproleassignment.NewUserAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAppRoleAssignment client: %+v", err)
	}
	configureFunc(userAppRoleAssignmentClient.Client)

	userAuthenticationClient, err := userauthentication.NewUserAuthenticationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthentication client: %+v", err)
	}
	configureFunc(userAuthenticationClient.Client)

	userAuthenticationEmailMethodClient, err := userauthenticationemailmethod.NewUserAuthenticationEmailMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationEmailMethod client: %+v", err)
	}
	configureFunc(userAuthenticationEmailMethodClient.Client)

	userAuthenticationFido2MethodClient, err := userauthenticationfido2method.NewUserAuthenticationFido2MethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationFido2Method client: %+v", err)
	}
	configureFunc(userAuthenticationFido2MethodClient.Client)

	userAuthenticationMethodClient, err := userauthenticationmethod.NewUserAuthenticationMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationMethod client: %+v", err)
	}
	configureFunc(userAuthenticationMethodClient.Client)

	userAuthenticationMicrosoftAuthenticatorMethodClient, err := userauthenticationmicrosoftauthenticatormethod.NewUserAuthenticationMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationMicrosoftAuthenticatorMethod client: %+v", err)
	}
	configureFunc(userAuthenticationMicrosoftAuthenticatorMethodClient.Client)

	userAuthenticationMicrosoftAuthenticatorMethodDeviceClient, err := userauthenticationmicrosoftauthenticatormethoddevice.NewUserAuthenticationMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationMicrosoftAuthenticatorMethodDevice client: %+v", err)
	}
	configureFunc(userAuthenticationMicrosoftAuthenticatorMethodDeviceClient.Client)

	userAuthenticationOperationClient, err := userauthenticationoperation.NewUserAuthenticationOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationOperation client: %+v", err)
	}
	configureFunc(userAuthenticationOperationClient.Client)

	userAuthenticationPasswordMethodClient, err := userauthenticationpasswordmethod.NewUserAuthenticationPasswordMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationPasswordMethod client: %+v", err)
	}
	configureFunc(userAuthenticationPasswordMethodClient.Client)

	userAuthenticationPhoneMethodClient, err := userauthenticationphonemethod.NewUserAuthenticationPhoneMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationPhoneMethod client: %+v", err)
	}
	configureFunc(userAuthenticationPhoneMethodClient.Client)

	userAuthenticationSoftwareOathMethodClient, err := userauthenticationsoftwareoathmethod.NewUserAuthenticationSoftwareOathMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationSoftwareOathMethod client: %+v", err)
	}
	configureFunc(userAuthenticationSoftwareOathMethodClient.Client)

	userAuthenticationTemporaryAccessPassMethodClient, err := userauthenticationtemporaryaccesspassmethod.NewUserAuthenticationTemporaryAccessPassMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationTemporaryAccessPassMethod client: %+v", err)
	}
	configureFunc(userAuthenticationTemporaryAccessPassMethodClient.Client)

	userAuthenticationWindowsHelloForBusinessMethodClient, err := userauthenticationwindowshelloforbusinessmethod.NewUserAuthenticationWindowsHelloForBusinessMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationWindowsHelloForBusinessMethod client: %+v", err)
	}
	configureFunc(userAuthenticationWindowsHelloForBusinessMethodClient.Client)

	userAuthenticationWindowsHelloForBusinessMethodDeviceClient, err := userauthenticationwindowshelloforbusinessmethoddevice.NewUserAuthenticationWindowsHelloForBusinessMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationWindowsHelloForBusinessMethodDevice client: %+v", err)
	}
	configureFunc(userAuthenticationWindowsHelloForBusinessMethodDeviceClient.Client)

	userCalendarCalendarPermissionClient, err := usercalendarcalendarpermission.NewUserCalendarCalendarPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarPermission client: %+v", err)
	}
	configureFunc(userCalendarCalendarPermissionClient.Client)

	userCalendarCalendarViewAttachmentClient, err := usercalendarcalendarviewattachment.NewUserCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewAttachment client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewAttachmentClient.Client)

	userCalendarCalendarViewCalendarClient, err := usercalendarcalendarviewcalendar.NewUserCalendarCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewCalendar client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewCalendarClient.Client)

	userCalendarCalendarViewClient, err := usercalendarcalendarview.NewUserCalendarCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarView client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewClient.Client)

	userCalendarCalendarViewExtensionClient, err := usercalendarcalendarviewextension.NewUserCalendarCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExtension client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExtensionClient.Client)

	userCalendarCalendarViewInstanceAttachmentClient, err := usercalendarcalendarviewinstanceattachment.NewUserCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceAttachmentClient.Client)

	userCalendarCalendarViewInstanceCalendarClient, err := usercalendarcalendarviewinstancecalendar.NewUserCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceCalendarClient.Client)

	userCalendarCalendarViewInstanceClient, err := usercalendarcalendarviewinstance.NewUserCalendarCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstance client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceClient.Client)

	userCalendarCalendarViewInstanceExtensionClient, err := usercalendarcalendarviewinstanceextension.NewUserCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceExtensionClient.Client)

	userCalendarClient, err := usercalendar.NewUserCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendar client: %+v", err)
	}
	configureFunc(userCalendarClient.Client)

	userCalendarEventAttachmentClient, err := usercalendareventattachment.NewUserCalendarEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventAttachment client: %+v", err)
	}
	configureFunc(userCalendarEventAttachmentClient.Client)

	userCalendarEventCalendarClient, err := usercalendareventcalendar.NewUserCalendarEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventCalendar client: %+v", err)
	}
	configureFunc(userCalendarEventCalendarClient.Client)

	userCalendarEventClient, err := usercalendarevent.NewUserCalendarEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEvent client: %+v", err)
	}
	configureFunc(userCalendarEventClient.Client)

	userCalendarEventExtensionClient, err := usercalendareventextension.NewUserCalendarEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExtension client: %+v", err)
	}
	configureFunc(userCalendarEventExtensionClient.Client)

	userCalendarEventInstanceAttachmentClient, err := usercalendareventinstanceattachment.NewUserCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceAttachmentClient.Client)

	userCalendarEventInstanceCalendarClient, err := usercalendareventinstancecalendar.NewUserCalendarEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceCalendarClient.Client)

	userCalendarEventInstanceClient, err := usercalendareventinstance.NewUserCalendarEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstance client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceClient.Client)

	userCalendarEventInstanceExtensionClient, err := usercalendareventinstanceextension.NewUserCalendarEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceExtensionClient.Client)

	userCalendarGroupCalendarCalendarPermissionClient, err := usercalendargroupcalendarcalendarpermission.NewUserCalendarGroupCalendarCalendarPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarPermission client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarPermissionClient.Client)

	userCalendarGroupCalendarCalendarViewAttachmentClient, err := usercalendargroupcalendarcalendarviewattachment.NewUserCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewAttachmentClient.Client)

	userCalendarGroupCalendarCalendarViewCalendarClient, err := usercalendargroupcalendarcalendarviewcalendar.NewUserCalendarGroupCalendarCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewCalendarClient.Client)

	userCalendarGroupCalendarCalendarViewClient, err := usercalendargroupcalendarcalendarview.NewUserCalendarGroupCalendarCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarView client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewClient.Client)

	userCalendarGroupCalendarCalendarViewExtensionClient, err := usercalendargroupcalendarcalendarviewextension.NewUserCalendarGroupCalendarCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExtensionClient.Client)

	userCalendarGroupCalendarCalendarViewInstanceAttachmentClient, err := usercalendargroupcalendarcalendarviewinstanceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceAttachmentClient.Client)

	userCalendarGroupCalendarCalendarViewInstanceCalendarClient, err := usercalendargroupcalendarcalendarviewinstancecalendar.NewUserCalendarGroupCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceCalendarClient.Client)

	userCalendarGroupCalendarCalendarViewInstanceClient, err := usercalendargroupcalendarcalendarviewinstance.NewUserCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstance client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceClient.Client)

	userCalendarGroupCalendarCalendarViewInstanceExtensionClient, err := usercalendargroupcalendarcalendarviewinstanceextension.NewUserCalendarGroupCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceExtensionClient.Client)

	userCalendarGroupCalendarClient, err := usercalendargroupcalendar.NewUserCalendarGroupCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarClient.Client)

	userCalendarGroupCalendarEventAttachmentClient, err := usercalendargroupcalendareventattachment.NewUserCalendarGroupCalendarEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventAttachmentClient.Client)

	userCalendarGroupCalendarEventCalendarClient, err := usercalendargroupcalendareventcalendar.NewUserCalendarGroupCalendarEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventCalendarClient.Client)

	userCalendarGroupCalendarEventClient, err := usercalendargroupcalendarevent.NewUserCalendarGroupCalendarEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEvent client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventClient.Client)

	userCalendarGroupCalendarEventExtensionClient, err := usercalendargroupcalendareventextension.NewUserCalendarGroupCalendarEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExtensionClient.Client)

	userCalendarGroupCalendarEventInstanceAttachmentClient, err := usercalendargroupcalendareventinstanceattachment.NewUserCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceAttachmentClient.Client)

	userCalendarGroupCalendarEventInstanceCalendarClient, err := usercalendargroupcalendareventinstancecalendar.NewUserCalendarGroupCalendarEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceCalendarClient.Client)

	userCalendarGroupCalendarEventInstanceClient, err := usercalendargroupcalendareventinstance.NewUserCalendarGroupCalendarEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstance client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceClient.Client)

	userCalendarGroupCalendarEventInstanceExtensionClient, err := usercalendargroupcalendareventinstanceextension.NewUserCalendarGroupCalendarEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceExtensionClient.Client)

	userCalendarGroupClient, err := usercalendargroup.NewUserCalendarGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroup client: %+v", err)
	}
	configureFunc(userCalendarGroupClient.Client)

	userCalendarViewAttachmentClient, err := usercalendarviewattachment.NewUserCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewAttachment client: %+v", err)
	}
	configureFunc(userCalendarViewAttachmentClient.Client)

	userCalendarViewCalendarClient, err := usercalendarviewcalendar.NewUserCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewCalendar client: %+v", err)
	}
	configureFunc(userCalendarViewCalendarClient.Client)

	userCalendarViewClient, err := usercalendarview.NewUserCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarView client: %+v", err)
	}
	configureFunc(userCalendarViewClient.Client)

	userCalendarViewExtensionClient, err := usercalendarviewextension.NewUserCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExtension client: %+v", err)
	}
	configureFunc(userCalendarViewExtensionClient.Client)

	userCalendarViewInstanceAttachmentClient, err := usercalendarviewinstanceattachment.NewUserCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceAttachmentClient.Client)

	userCalendarViewInstanceCalendarClient, err := usercalendarviewinstancecalendar.NewUserCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceCalendarClient.Client)

	userCalendarViewInstanceClient, err := usercalendarviewinstance.NewUserCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstance client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceClient.Client)

	userCalendarViewInstanceExtensionClient, err := usercalendarviewinstanceextension.NewUserCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceExtensionClient.Client)

	userChatClient, err := userchat.NewUserChatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChat client: %+v", err)
	}
	configureFunc(userChatClient.Client)

	userChatInstalledAppClient, err := userchatinstalledapp.NewUserChatInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatInstalledApp client: %+v", err)
	}
	configureFunc(userChatInstalledAppClient.Client)

	userChatInstalledAppTeamsAppClient, err := userchatinstalledappteamsapp.NewUserChatInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(userChatInstalledAppTeamsAppClient.Client)

	userChatInstalledAppTeamsAppDefinitionClient, err := userchatinstalledappteamsappdefinition.NewUserChatInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(userChatInstalledAppTeamsAppDefinitionClient.Client)

	userChatLastMessagePreviewClient, err := userchatlastmessagepreview.NewUserChatLastMessagePreviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatLastMessagePreview client: %+v", err)
	}
	configureFunc(userChatLastMessagePreviewClient.Client)

	userChatMemberClient, err := userchatmember.NewUserChatMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatMember client: %+v", err)
	}
	configureFunc(userChatMemberClient.Client)

	userChatMessageClient, err := userchatmessage.NewUserChatMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatMessage client: %+v", err)
	}
	configureFunc(userChatMessageClient.Client)

	userChatMessageHostedContentClient, err := userchatmessagehostedcontent.NewUserChatMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatMessageHostedContent client: %+v", err)
	}
	configureFunc(userChatMessageHostedContentClient.Client)

	userChatMessageReplyClient, err := userchatmessagereply.NewUserChatMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatMessageReply client: %+v", err)
	}
	configureFunc(userChatMessageReplyClient.Client)

	userChatMessageReplyHostedContentClient, err := userchatmessagereplyhostedcontent.NewUserChatMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(userChatMessageReplyHostedContentClient.Client)

	userChatPermissionGrantClient, err := userchatpermissiongrant.NewUserChatPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatPermissionGrant client: %+v", err)
	}
	configureFunc(userChatPermissionGrantClient.Client)

	userChatPinnedMessageClient, err := userchatpinnedmessage.NewUserChatPinnedMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatPinnedMessage client: %+v", err)
	}
	configureFunc(userChatPinnedMessageClient.Client)

	userChatPinnedMessageMessageClient, err := userchatpinnedmessagemessage.NewUserChatPinnedMessageMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatPinnedMessageMessage client: %+v", err)
	}
	configureFunc(userChatPinnedMessageMessageClient.Client)

	userChatTabClient, err := userchattab.NewUserChatTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatTab client: %+v", err)
	}
	configureFunc(userChatTabClient.Client)

	userChatTabTeamsAppClient, err := userchattabteamsapp.NewUserChatTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatTabTeamsApp client: %+v", err)
	}
	configureFunc(userChatTabTeamsAppClient.Client)

	userClient, err := user.NewUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building User client: %+v", err)
	}
	configureFunc(userClient.Client)

	userContactClient, err := usercontact.NewUserContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContact client: %+v", err)
	}
	configureFunc(userContactClient.Client)

	userContactExtensionClient, err := usercontactextension.NewUserContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactExtension client: %+v", err)
	}
	configureFunc(userContactExtensionClient.Client)

	userContactFolderChildFolderClient, err := usercontactfolderchildfolder.NewUserContactFolderChildFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolderChildFolder client: %+v", err)
	}
	configureFunc(userContactFolderChildFolderClient.Client)

	userContactFolderChildFolderContactClient, err := usercontactfolderchildfoldercontact.NewUserContactFolderChildFolderContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolderChildFolderContact client: %+v", err)
	}
	configureFunc(userContactFolderChildFolderContactClient.Client)

	userContactFolderChildFolderContactExtensionClient, err := usercontactfolderchildfoldercontactextension.NewUserContactFolderChildFolderContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolderChildFolderContactExtension client: %+v", err)
	}
	configureFunc(userContactFolderChildFolderContactExtensionClient.Client)

	userContactFolderChildFolderContactPhotoClient, err := usercontactfolderchildfoldercontactphoto.NewUserContactFolderChildFolderContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolderChildFolderContactPhoto client: %+v", err)
	}
	configureFunc(userContactFolderChildFolderContactPhotoClient.Client)

	userContactFolderClient, err := usercontactfolder.NewUserContactFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolder client: %+v", err)
	}
	configureFunc(userContactFolderClient.Client)

	userContactFolderContactClient, err := usercontactfoldercontact.NewUserContactFolderContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolderContact client: %+v", err)
	}
	configureFunc(userContactFolderContactClient.Client)

	userContactFolderContactExtensionClient, err := usercontactfoldercontactextension.NewUserContactFolderContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolderContactExtension client: %+v", err)
	}
	configureFunc(userContactFolderContactExtensionClient.Client)

	userContactFolderContactPhotoClient, err := usercontactfoldercontactphoto.NewUserContactFolderContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactFolderContactPhoto client: %+v", err)
	}
	configureFunc(userContactFolderContactPhotoClient.Client)

	userContactPhotoClient, err := usercontactphoto.NewUserContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserContactPhoto client: %+v", err)
	}
	configureFunc(userContactPhotoClient.Client)

	userCreatedObjectClient, err := usercreatedobject.NewUserCreatedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCreatedObject client: %+v", err)
	}
	configureFunc(userCreatedObjectClient.Client)

	userDeviceManagementTroubleshootingEventClient, err := userdevicemanagementtroubleshootingevent.NewUserDeviceManagementTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceManagementTroubleshootingEvent client: %+v", err)
	}
	configureFunc(userDeviceManagementTroubleshootingEventClient.Client)

	userDirectReportClient, err := userdirectreport.NewUserDirectReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDirectReport client: %+v", err)
	}
	configureFunc(userDirectReportClient.Client)

	userDriveClient, err := userdrive.NewUserDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDrive client: %+v", err)
	}
	configureFunc(userDriveClient.Client)

	userEmployeeExperienceClient, err := useremployeeexperience.NewUserEmployeeExperienceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEmployeeExperience client: %+v", err)
	}
	configureFunc(userEmployeeExperienceClient.Client)

	userEmployeeExperienceLearningCourseActivityClient, err := useremployeeexperiencelearningcourseactivity.NewUserEmployeeExperienceLearningCourseActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEmployeeExperienceLearningCourseActivity client: %+v", err)
	}
	configureFunc(userEmployeeExperienceLearningCourseActivityClient.Client)

	userEventAttachmentClient, err := usereventattachment.NewUserEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventAttachment client: %+v", err)
	}
	configureFunc(userEventAttachmentClient.Client)

	userEventCalendarClient, err := usereventcalendar.NewUserEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventCalendar client: %+v", err)
	}
	configureFunc(userEventCalendarClient.Client)

	userEventClient, err := userevent.NewUserEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEvent client: %+v", err)
	}
	configureFunc(userEventClient.Client)

	userEventExtensionClient, err := usereventextension.NewUserEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExtension client: %+v", err)
	}
	configureFunc(userEventExtensionClient.Client)

	userEventInstanceAttachmentClient, err := usereventinstanceattachment.NewUserEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstanceAttachment client: %+v", err)
	}
	configureFunc(userEventInstanceAttachmentClient.Client)

	userEventInstanceCalendarClient, err := usereventinstancecalendar.NewUserEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstanceCalendar client: %+v", err)
	}
	configureFunc(userEventInstanceCalendarClient.Client)

	userEventInstanceClient, err := usereventinstance.NewUserEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstance client: %+v", err)
	}
	configureFunc(userEventInstanceClient.Client)

	userEventInstanceExtensionClient, err := usereventinstanceextension.NewUserEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstanceExtension client: %+v", err)
	}
	configureFunc(userEventInstanceExtensionClient.Client)

	userExtensionClient, err := userextension.NewUserExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserExtension client: %+v", err)
	}
	configureFunc(userExtensionClient.Client)

	userFollowedSiteClient, err := userfollowedsite.NewUserFollowedSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserFollowedSite client: %+v", err)
	}
	configureFunc(userFollowedSiteClient.Client)

	userInferenceClassificationClient, err := userinferenceclassification.NewUserInferenceClassificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInferenceClassification client: %+v", err)
	}
	configureFunc(userInferenceClassificationClient.Client)

	userInferenceClassificationOverrideClient, err := userinferenceclassificationoverride.NewUserInferenceClassificationOverrideClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInferenceClassificationOverride client: %+v", err)
	}
	configureFunc(userInferenceClassificationOverrideClient.Client)

	userInsightClient, err := userinsight.NewUserInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsight client: %+v", err)
	}
	configureFunc(userInsightClient.Client)

	userInsightSharedClient, err := userinsightshared.NewUserInsightSharedClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightShared client: %+v", err)
	}
	configureFunc(userInsightSharedClient.Client)

	userInsightSharedLastSharedMethodClient, err := userinsightsharedlastsharedmethod.NewUserInsightSharedLastSharedMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightSharedLastSharedMethod client: %+v", err)
	}
	configureFunc(userInsightSharedLastSharedMethodClient.Client)

	userInsightSharedResourceClient, err := userinsightsharedresource.NewUserInsightSharedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightSharedResource client: %+v", err)
	}
	configureFunc(userInsightSharedResourceClient.Client)

	userInsightTrendingClient, err := userinsighttrending.NewUserInsightTrendingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightTrending client: %+v", err)
	}
	configureFunc(userInsightTrendingClient.Client)

	userInsightTrendingResourceClient, err := userinsighttrendingresource.NewUserInsightTrendingResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightTrendingResource client: %+v", err)
	}
	configureFunc(userInsightTrendingResourceClient.Client)

	userInsightUsedClient, err := userinsightused.NewUserInsightUsedClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightUsed client: %+v", err)
	}
	configureFunc(userInsightUsedClient.Client)

	userInsightUsedResourceClient, err := userinsightusedresource.NewUserInsightUsedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInsightUsedResource client: %+v", err)
	}
	configureFunc(userInsightUsedResourceClient.Client)

	userJoinedTeamAllChannelClient, err := userjoinedteamallchannel.NewUserJoinedTeamAllChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamAllChannel client: %+v", err)
	}
	configureFunc(userJoinedTeamAllChannelClient.Client)

	userJoinedTeamChannelClient, err := userjoinedteamchannel.NewUserJoinedTeamChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannel client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelClient.Client)

	userJoinedTeamChannelFilesFolderClient, err := userjoinedteamchannelfilesfolder.NewUserJoinedTeamChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelFilesFolder client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelFilesFolderClient.Client)

	userJoinedTeamChannelFilesFolderContentClient, err := userjoinedteamchannelfilesfoldercontent.NewUserJoinedTeamChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelFilesFolderContentClient.Client)

	userJoinedTeamChannelMemberClient, err := userjoinedteamchannelmember.NewUserJoinedTeamChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelMember client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelMemberClient.Client)

	userJoinedTeamChannelMessageClient, err := userjoinedteamchannelmessage.NewUserJoinedTeamChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelMessage client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelMessageClient.Client)

	userJoinedTeamChannelMessageHostedContentClient, err := userjoinedteamchannelmessagehostedcontent.NewUserJoinedTeamChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelMessageHostedContentClient.Client)

	userJoinedTeamChannelMessageReplyClient, err := userjoinedteamchannelmessagereply.NewUserJoinedTeamChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelMessageReply client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelMessageReplyClient.Client)

	userJoinedTeamChannelMessageReplyHostedContentClient, err := userjoinedteamchannelmessagereplyhostedcontent.NewUserJoinedTeamChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelMessageReplyHostedContentClient.Client)

	userJoinedTeamChannelSharedWithTeamAllowedMemberClient, err := userjoinedteamchannelsharedwithteamallowedmember.NewUserJoinedTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelSharedWithTeamAllowedMemberClient.Client)

	userJoinedTeamChannelSharedWithTeamClient, err := userjoinedteamchannelsharedwithteam.NewUserJoinedTeamChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelSharedWithTeamClient.Client)

	userJoinedTeamChannelSharedWithTeamTeamClient, err := userjoinedteamchannelsharedwithteamteam.NewUserJoinedTeamChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelSharedWithTeamTeamClient.Client)

	userJoinedTeamChannelTabClient, err := userjoinedteamchanneltab.NewUserJoinedTeamChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelTab client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelTabClient.Client)

	userJoinedTeamChannelTabTeamsAppClient, err := userjoinedteamchanneltabteamsapp.NewUserJoinedTeamChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(userJoinedTeamChannelTabTeamsAppClient.Client)

	userJoinedTeamClient, err := userjoinedteam.NewUserJoinedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeam client: %+v", err)
	}
	configureFunc(userJoinedTeamClient.Client)

	userJoinedTeamGroupClient, err := userjoinedteamgroup.NewUserJoinedTeamGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamGroup client: %+v", err)
	}
	configureFunc(userJoinedTeamGroupClient.Client)

	userJoinedTeamIncomingChannelClient, err := userjoinedteamincomingchannel.NewUserJoinedTeamIncomingChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamIncomingChannel client: %+v", err)
	}
	configureFunc(userJoinedTeamIncomingChannelClient.Client)

	userJoinedTeamInstalledAppClient, err := userjoinedteaminstalledapp.NewUserJoinedTeamInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamInstalledApp client: %+v", err)
	}
	configureFunc(userJoinedTeamInstalledAppClient.Client)

	userJoinedTeamInstalledAppTeamsAppClient, err := userjoinedteaminstalledappteamsapp.NewUserJoinedTeamInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(userJoinedTeamInstalledAppTeamsAppClient.Client)

	userJoinedTeamInstalledAppTeamsAppDefinitionClient, err := userjoinedteaminstalledappteamsappdefinition.NewUserJoinedTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(userJoinedTeamInstalledAppTeamsAppDefinitionClient.Client)

	userJoinedTeamMemberClient, err := userjoinedteammember.NewUserJoinedTeamMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamMember client: %+v", err)
	}
	configureFunc(userJoinedTeamMemberClient.Client)

	userJoinedTeamOperationClient, err := userjoinedteamoperation.NewUserJoinedTeamOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamOperation client: %+v", err)
	}
	configureFunc(userJoinedTeamOperationClient.Client)

	userJoinedTeamPermissionGrantClient, err := userjoinedteampermissiongrant.NewUserJoinedTeamPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPermissionGrant client: %+v", err)
	}
	configureFunc(userJoinedTeamPermissionGrantClient.Client)

	userJoinedTeamPhotoClient, err := userjoinedteamphoto.NewUserJoinedTeamPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPhoto client: %+v", err)
	}
	configureFunc(userJoinedTeamPhotoClient.Client)

	userJoinedTeamPrimaryChannelClient, err := userjoinedteamprimarychannel.NewUserJoinedTeamPrimaryChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannel client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelClient.Client)

	userJoinedTeamPrimaryChannelFilesFolderClient, err := userjoinedteamprimarychannelfilesfolder.NewUserJoinedTeamPrimaryChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelFilesFolder client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelFilesFolderClient.Client)

	userJoinedTeamPrimaryChannelFilesFolderContentClient, err := userjoinedteamprimarychannelfilesfoldercontent.NewUserJoinedTeamPrimaryChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelFilesFolderContentClient.Client)

	userJoinedTeamPrimaryChannelMemberClient, err := userjoinedteamprimarychannelmember.NewUserJoinedTeamPrimaryChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelMember client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelMemberClient.Client)

	userJoinedTeamPrimaryChannelMessageClient, err := userjoinedteamprimarychannelmessage.NewUserJoinedTeamPrimaryChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelMessage client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelMessageClient.Client)

	userJoinedTeamPrimaryChannelMessageHostedContentClient, err := userjoinedteamprimarychannelmessagehostedcontent.NewUserJoinedTeamPrimaryChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelMessageHostedContentClient.Client)

	userJoinedTeamPrimaryChannelMessageReplyClient, err := userjoinedteamprimarychannelmessagereply.NewUserJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelMessageReply client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelMessageReplyClient.Client)

	userJoinedTeamPrimaryChannelMessageReplyHostedContentClient, err := userjoinedteamprimarychannelmessagereplyhostedcontent.NewUserJoinedTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelMessageReplyHostedContentClient.Client)

	userJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient, err := userjoinedteamprimarychannelsharedwithteamallowedmember.NewUserJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient.Client)

	userJoinedTeamPrimaryChannelSharedWithTeamClient, err := userjoinedteamprimarychannelsharedwithteam.NewUserJoinedTeamPrimaryChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelSharedWithTeamClient.Client)

	userJoinedTeamPrimaryChannelSharedWithTeamTeamClient, err := userjoinedteamprimarychannelsharedwithteamteam.NewUserJoinedTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelSharedWithTeamTeamClient.Client)

	userJoinedTeamPrimaryChannelTabClient, err := userjoinedteamprimarychanneltab.NewUserJoinedTeamPrimaryChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelTab client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelTabClient.Client)

	userJoinedTeamPrimaryChannelTabTeamsAppClient, err := userjoinedteamprimarychanneltabteamsapp.NewUserJoinedTeamPrimaryChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamPrimaryChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(userJoinedTeamPrimaryChannelTabTeamsAppClient.Client)

	userJoinedTeamScheduleClient, err := userjoinedteamschedule.NewUserJoinedTeamScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamSchedule client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleClient.Client)

	userJoinedTeamScheduleOfferShiftRequestClient, err := userjoinedteamscheduleoffershiftrequest.NewUserJoinedTeamScheduleOfferShiftRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleOfferShiftRequest client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleOfferShiftRequestClient.Client)

	userJoinedTeamScheduleOpenShiftChangeRequestClient, err := userjoinedteamscheduleopenshiftchangerequest.NewUserJoinedTeamScheduleOpenShiftChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleOpenShiftChangeRequest client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleOpenShiftChangeRequestClient.Client)

	userJoinedTeamScheduleOpenShiftClient, err := userjoinedteamscheduleopenshift.NewUserJoinedTeamScheduleOpenShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleOpenShift client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleOpenShiftClient.Client)

	userJoinedTeamScheduleSchedulingGroupClient, err := userjoinedteamscheduleschedulinggroup.NewUserJoinedTeamScheduleSchedulingGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleSchedulingGroup client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleSchedulingGroupClient.Client)

	userJoinedTeamScheduleShiftClient, err := userjoinedteamscheduleshift.NewUserJoinedTeamScheduleShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleShift client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleShiftClient.Client)

	userJoinedTeamScheduleSwapShiftsChangeRequestClient, err := userjoinedteamscheduleswapshiftschangerequest.NewUserJoinedTeamScheduleSwapShiftsChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleSwapShiftsChangeRequest client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleSwapShiftsChangeRequestClient.Client)

	userJoinedTeamScheduleTimeOffReasonClient, err := userjoinedteamscheduletimeoffreason.NewUserJoinedTeamScheduleTimeOffReasonClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleTimeOffReason client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleTimeOffReasonClient.Client)

	userJoinedTeamScheduleTimeOffRequestClient, err := userjoinedteamscheduletimeoffrequest.NewUserJoinedTeamScheduleTimeOffRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleTimeOffRequest client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleTimeOffRequestClient.Client)

	userJoinedTeamScheduleTimesOffClient, err := userjoinedteamscheduletimesoff.NewUserJoinedTeamScheduleTimesOffClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamScheduleTimesOff client: %+v", err)
	}
	configureFunc(userJoinedTeamScheduleTimesOffClient.Client)

	userJoinedTeamTagClient, err := userjoinedteamtag.NewUserJoinedTeamTagClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamTag client: %+v", err)
	}
	configureFunc(userJoinedTeamTagClient.Client)

	userJoinedTeamTagMemberClient, err := userjoinedteamtagmember.NewUserJoinedTeamTagMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamTagMember client: %+v", err)
	}
	configureFunc(userJoinedTeamTagMemberClient.Client)

	userJoinedTeamTemplateClient, err := userjoinedteamtemplate.NewUserJoinedTeamTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeamTemplate client: %+v", err)
	}
	configureFunc(userJoinedTeamTemplateClient.Client)

	userLicenseDetailClient, err := userlicensedetail.NewUserLicenseDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserLicenseDetail client: %+v", err)
	}
	configureFunc(userLicenseDetailClient.Client)

	userMailFolderChildFolderClient, err := usermailfolderchildfolder.NewUserMailFolderChildFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolder client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderClient.Client)

	userMailFolderChildFolderMessageAttachmentClient, err := usermailfolderchildfoldermessageattachment.NewUserMailFolderChildFolderMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolderMessageAttachment client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderMessageAttachmentClient.Client)

	userMailFolderChildFolderMessageClient, err := usermailfolderchildfoldermessage.NewUserMailFolderChildFolderMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolderMessage client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderMessageClient.Client)

	userMailFolderChildFolderMessageExtensionClient, err := usermailfolderchildfoldermessageextension.NewUserMailFolderChildFolderMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolderMessageExtension client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderMessageExtensionClient.Client)

	userMailFolderChildFolderMessageRuleClient, err := usermailfolderchildfoldermessagerule.NewUserMailFolderChildFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolderMessageRule client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderMessageRuleClient.Client)

	userMailFolderClient, err := usermailfolder.NewUserMailFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolder client: %+v", err)
	}
	configureFunc(userMailFolderClient.Client)

	userMailFolderMessageAttachmentClient, err := usermailfoldermessageattachment.NewUserMailFolderMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderMessageAttachment client: %+v", err)
	}
	configureFunc(userMailFolderMessageAttachmentClient.Client)

	userMailFolderMessageClient, err := usermailfoldermessage.NewUserMailFolderMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderMessage client: %+v", err)
	}
	configureFunc(userMailFolderMessageClient.Client)

	userMailFolderMessageExtensionClient, err := usermailfoldermessageextension.NewUserMailFolderMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderMessageExtension client: %+v", err)
	}
	configureFunc(userMailFolderMessageExtensionClient.Client)

	userMailFolderMessageRuleClient, err := usermailfoldermessagerule.NewUserMailFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderMessageRule client: %+v", err)
	}
	configureFunc(userMailFolderMessageRuleClient.Client)

	userMailboxSettingClient, err := usermailboxsetting.NewUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailboxSetting client: %+v", err)
	}
	configureFunc(userMailboxSettingClient.Client)

	userManagedAppRegistrationClient, err := usermanagedappregistration.NewUserManagedAppRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedAppRegistration client: %+v", err)
	}
	configureFunc(userManagedAppRegistrationClient.Client)

	userManagedDeviceClient, err := usermanageddevice.NewUserManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDevice client: %+v", err)
	}
	configureFunc(userManagedDeviceClient.Client)

	userManagedDeviceDeviceCategoryClient, err := usermanageddevicedevicecategory.NewUserManagedDeviceDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceDeviceCategory client: %+v", err)
	}
	configureFunc(userManagedDeviceDeviceCategoryClient.Client)

	userManagedDeviceDeviceCompliancePolicyStateClient, err := usermanageddevicedevicecompliancepolicystate.NewUserManagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceDeviceCompliancePolicyState client: %+v", err)
	}
	configureFunc(userManagedDeviceDeviceCompliancePolicyStateClient.Client)

	userManagedDeviceDeviceConfigurationStateClient, err := usermanageddevicedeviceconfigurationstate.NewUserManagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceDeviceConfigurationState client: %+v", err)
	}
	configureFunc(userManagedDeviceDeviceConfigurationStateClient.Client)

	userManagedDeviceLogCollectionRequestClient, err := usermanageddevicelogcollectionrequest.NewUserManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(userManagedDeviceLogCollectionRequestClient.Client)

	userManagedDeviceUserClient, err := usermanageddeviceuser.NewUserManagedDeviceUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceUser client: %+v", err)
	}
	configureFunc(userManagedDeviceUserClient.Client)

	userManagedDeviceWindowsProtectionStateClient, err := usermanageddevicewindowsprotectionstate.NewUserManagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceWindowsProtectionState client: %+v", err)
	}
	configureFunc(userManagedDeviceWindowsProtectionStateClient.Client)

	userManagedDeviceWindowsProtectionStateDetectedMalwareStateClient, err := usermanageddevicewindowsprotectionstatedetectedmalwarestate.NewUserManagedDeviceWindowsProtectionStateDetectedMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceWindowsProtectionStateDetectedMalwareState client: %+v", err)
	}
	configureFunc(userManagedDeviceWindowsProtectionStateDetectedMalwareStateClient.Client)

	userManagerClient, err := usermanager.NewUserManagerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManager client: %+v", err)
	}
	configureFunc(userManagerClient.Client)

	userMemberOfClient, err := usermemberof.NewUserMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMemberOf client: %+v", err)
	}
	configureFunc(userMemberOfClient.Client)

	userMessageAttachmentClient, err := usermessageattachment.NewUserMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMessageAttachment client: %+v", err)
	}
	configureFunc(userMessageAttachmentClient.Client)

	userMessageClient, err := usermessage.NewUserMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMessage client: %+v", err)
	}
	configureFunc(userMessageClient.Client)

	userMessageExtensionClient, err := usermessageextension.NewUserMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMessageExtension client: %+v", err)
	}
	configureFunc(userMessageExtensionClient.Client)

	userOauth2PermissionGrantClient, err := useroauth2permissiongrant.NewUserOauth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOauth2PermissionGrant client: %+v", err)
	}
	configureFunc(userOauth2PermissionGrantClient.Client)

	userOnenoteClient, err := useronenote.NewUserOnenoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenote client: %+v", err)
	}
	configureFunc(userOnenoteClient.Client)

	userOnenoteNotebookClient, err := useronenotenotebook.NewUserOnenoteNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebook client: %+v", err)
	}
	configureFunc(userOnenoteNotebookClient.Client)

	userOnenoteNotebookSectionClient, err := useronenotenotebooksection.NewUserOnenoteNotebookSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSection client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionClient.Client)

	userOnenoteNotebookSectionGroupClient, err := useronenotenotebooksectiongroup.NewUserOnenoteNotebookSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupClient.Client)

	userOnenoteNotebookSectionGroupParentNotebookClient, err := useronenotenotebooksectiongroupparentnotebook.NewUserOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupParentNotebookClient.Client)

	userOnenoteNotebookSectionGroupParentSectionGroupClient, err := useronenotenotebooksectiongroupparentsectiongroup.NewUserOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupParentSectionGroupClient.Client)

	userOnenoteNotebookSectionGroupSectionClient, err := useronenotenotebooksectiongroupsection.NewUserOnenoteNotebookSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSection client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionClient.Client)

	userOnenoteNotebookSectionGroupSectionGroupClient, err := useronenotenotebooksectiongroupsectiongroup.NewUserOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionGroupClient.Client)

	userOnenoteNotebookSectionGroupSectionPageClient, err := useronenotenotebooksectiongroupsectionpage.NewUserOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionPageClient.Client)

	userOnenoteNotebookSectionGroupSectionPageContentClient, err := useronenotenotebooksectiongroupsectionpagecontent.NewUserOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionPageContentClient.Client)

	userOnenoteNotebookSectionGroupSectionPageParentNotebookClient, err := useronenotenotebooksectiongroupsectionpageparentnotebook.NewUserOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionPageParentNotebookClient.Client)

	userOnenoteNotebookSectionGroupSectionPageParentSectionClient, err := useronenotenotebooksectiongroupsectionpageparentsection.NewUserOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionPageParentSectionClient.Client)

	userOnenoteNotebookSectionGroupSectionParentNotebookClient, err := useronenotenotebooksectiongroupsectionparentnotebook.NewUserOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionParentNotebookClient.Client)

	userOnenoteNotebookSectionGroupSectionParentSectionGroupClient, err := useronenotenotebooksectiongroupsectionparentsectiongroup.NewUserOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionGroupSectionParentSectionGroupClient.Client)

	userOnenoteNotebookSectionPageClient, err := useronenotenotebooksectionpage.NewUserOnenoteNotebookSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionPage client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionPageClient.Client)

	userOnenoteNotebookSectionPageContentClient, err := useronenotenotebooksectionpagecontent.NewUserOnenoteNotebookSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionPageContent client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionPageContentClient.Client)

	userOnenoteNotebookSectionPageParentNotebookClient, err := useronenotenotebooksectionpageparentnotebook.NewUserOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionPageParentNotebookClient.Client)

	userOnenoteNotebookSectionPageParentSectionClient, err := useronenotenotebooksectionpageparentsection.NewUserOnenoteNotebookSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionPageParentSection client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionPageParentSectionClient.Client)

	userOnenoteNotebookSectionParentNotebookClient, err := useronenotenotebooksectionparentnotebook.NewUserOnenoteNotebookSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionParentNotebookClient.Client)

	userOnenoteNotebookSectionParentSectionGroupClient, err := useronenotenotebooksectionparentsectiongroup.NewUserOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteNotebookSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteNotebookSectionParentSectionGroupClient.Client)

	userOnenoteOperationClient, err := useronenoteoperation.NewUserOnenoteOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteOperation client: %+v", err)
	}
	configureFunc(userOnenoteOperationClient.Client)

	userOnenotePageClient, err := useronenotepage.NewUserOnenotePageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenotePage client: %+v", err)
	}
	configureFunc(userOnenotePageClient.Client)

	userOnenotePageContentClient, err := useronenotepagecontent.NewUserOnenotePageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenotePageContent client: %+v", err)
	}
	configureFunc(userOnenotePageContentClient.Client)

	userOnenotePageParentNotebookClient, err := useronenotepageparentnotebook.NewUserOnenotePageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenotePageParentNotebook client: %+v", err)
	}
	configureFunc(userOnenotePageParentNotebookClient.Client)

	userOnenotePageParentSectionClient, err := useronenotepageparentsection.NewUserOnenotePageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenotePageParentSection client: %+v", err)
	}
	configureFunc(userOnenotePageParentSectionClient.Client)

	userOnenoteResourceClient, err := useronenoteresource.NewUserOnenoteResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteResource client: %+v", err)
	}
	configureFunc(userOnenoteResourceClient.Client)

	userOnenoteResourceContentClient, err := useronenoteresourcecontent.NewUserOnenoteResourceContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteResourceContent client: %+v", err)
	}
	configureFunc(userOnenoteResourceContentClient.Client)

	userOnenoteSectionClient, err := useronenotesection.NewUserOnenoteSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSection client: %+v", err)
	}
	configureFunc(userOnenoteSectionClient.Client)

	userOnenoteSectionGroupClient, err := useronenotesectiongroup.NewUserOnenoteSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupClient.Client)

	userOnenoteSectionGroupParentNotebookClient, err := useronenotesectiongroupparentnotebook.NewUserOnenoteSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupParentNotebookClient.Client)

	userOnenoteSectionGroupParentSectionGroupClient, err := useronenotesectiongroupparentsectiongroup.NewUserOnenoteSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupParentSectionGroupClient.Client)

	userOnenoteSectionGroupSectionClient, err := useronenotesectiongroupsection.NewUserOnenoteSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSection client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionClient.Client)

	userOnenoteSectionGroupSectionGroupClient, err := useronenotesectiongroupsectiongroup.NewUserOnenoteSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionGroupClient.Client)

	userOnenoteSectionGroupSectionPageClient, err := useronenotesectiongroupsectionpage.NewUserOnenoteSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionPageClient.Client)

	userOnenoteSectionGroupSectionPageContentClient, err := useronenotesectiongroupsectionpagecontent.NewUserOnenoteSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionPageContentClient.Client)

	userOnenoteSectionGroupSectionPageParentNotebookClient, err := useronenotesectiongroupsectionpageparentnotebook.NewUserOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionPageParentNotebookClient.Client)

	userOnenoteSectionGroupSectionPageParentSectionClient, err := useronenotesectiongroupsectionpageparentsection.NewUserOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionPageParentSectionClient.Client)

	userOnenoteSectionGroupSectionParentNotebookClient, err := useronenotesectiongroupsectionparentnotebook.NewUserOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionParentNotebookClient.Client)

	userOnenoteSectionGroupSectionParentSectionGroupClient, err := useronenotesectiongroupsectionparentsectiongroup.NewUserOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteSectionGroupSectionParentSectionGroupClient.Client)

	userOnenoteSectionPageClient, err := useronenotesectionpage.NewUserOnenoteSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionPage client: %+v", err)
	}
	configureFunc(userOnenoteSectionPageClient.Client)

	userOnenoteSectionPageContentClient, err := useronenotesectionpagecontent.NewUserOnenoteSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionPageContent client: %+v", err)
	}
	configureFunc(userOnenoteSectionPageContentClient.Client)

	userOnenoteSectionPageParentNotebookClient, err := useronenotesectionpageparentnotebook.NewUserOnenoteSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteSectionPageParentNotebookClient.Client)

	userOnenoteSectionPageParentSectionClient, err := useronenotesectionpageparentsection.NewUserOnenoteSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionPageParentSection client: %+v", err)
	}
	configureFunc(userOnenoteSectionPageParentSectionClient.Client)

	userOnenoteSectionParentNotebookClient, err := useronenotesectionparentnotebook.NewUserOnenoteSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionParentNotebook client: %+v", err)
	}
	configureFunc(userOnenoteSectionParentNotebookClient.Client)

	userOnenoteSectionParentSectionGroupClient, err := useronenotesectionparentsectiongroup.NewUserOnenoteSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnenoteSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(userOnenoteSectionParentSectionGroupClient.Client)

	userOnlineMeetingAttendanceReportAttendanceRecordClient, err := useronlinemeetingattendancereportattendancerecord.NewUserOnlineMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingAttendanceReportAttendanceRecord client: %+v", err)
	}
	configureFunc(userOnlineMeetingAttendanceReportAttendanceRecordClient.Client)

	userOnlineMeetingAttendanceReportClient, err := useronlinemeetingattendancereport.NewUserOnlineMeetingAttendanceReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingAttendanceReport client: %+v", err)
	}
	configureFunc(userOnlineMeetingAttendanceReportClient.Client)

	userOnlineMeetingAttendeeReportClient, err := useronlinemeetingattendeereport.NewUserOnlineMeetingAttendeeReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingAttendeeReport client: %+v", err)
	}
	configureFunc(userOnlineMeetingAttendeeReportClient.Client)

	userOnlineMeetingClient, err := useronlinemeeting.NewUserOnlineMeetingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeeting client: %+v", err)
	}
	configureFunc(userOnlineMeetingClient.Client)

	userOutlookClient, err := useroutlook.NewUserOutlookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlook client: %+v", err)
	}
	configureFunc(userOutlookClient.Client)

	userOutlookMasterCategoryClient, err := useroutlookmastercategory.NewUserOutlookMasterCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookMasterCategory client: %+v", err)
	}
	configureFunc(userOutlookMasterCategoryClient.Client)

	userOwnedDeviceClient, err := userowneddevice.NewUserOwnedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOwnedDevice client: %+v", err)
	}
	configureFunc(userOwnedDeviceClient.Client)

	userOwnedObjectClient, err := userownedobject.NewUserOwnedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOwnedObject client: %+v", err)
	}
	configureFunc(userOwnedObjectClient.Client)

	userPeopleClient, err := userpeople.NewUserPeopleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPeople client: %+v", err)
	}
	configureFunc(userPeopleClient.Client)

	userPhotoClient, err := userphoto.NewUserPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPhoto client: %+v", err)
	}
	configureFunc(userPhotoClient.Client)

	userPlannerClient, err := userplanner.NewUserPlannerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlanner client: %+v", err)
	}
	configureFunc(userPlannerClient.Client)

	userPlannerPlanBucketClient, err := userplannerplanbucket.NewUserPlannerPlanBucketClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanBucket client: %+v", err)
	}
	configureFunc(userPlannerPlanBucketClient.Client)

	userPlannerPlanBucketTaskAssignedToTaskBoardFormatClient, err := userplannerplanbuckettaskassignedtotaskboardformat.NewUserPlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanBucketTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerPlanBucketTaskAssignedToTaskBoardFormatClient.Client)

	userPlannerPlanBucketTaskBucketTaskBoardFormatClient, err := userplannerplanbuckettaskbuckettaskboardformat.NewUserPlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanBucketTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerPlanBucketTaskBucketTaskBoardFormatClient.Client)

	userPlannerPlanBucketTaskClient, err := userplannerplanbuckettask.NewUserPlannerPlanBucketTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanBucketTask client: %+v", err)
	}
	configureFunc(userPlannerPlanBucketTaskClient.Client)

	userPlannerPlanBucketTaskDetailClient, err := userplannerplanbuckettaskdetail.NewUserPlannerPlanBucketTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanBucketTaskDetail client: %+v", err)
	}
	configureFunc(userPlannerPlanBucketTaskDetailClient.Client)

	userPlannerPlanBucketTaskProgressTaskBoardFormatClient, err := userplannerplanbuckettaskprogresstaskboardformat.NewUserPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanBucketTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerPlanBucketTaskProgressTaskBoardFormatClient.Client)

	userPlannerPlanClient, err := userplannerplan.NewUserPlannerPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlan client: %+v", err)
	}
	configureFunc(userPlannerPlanClient.Client)

	userPlannerPlanDetailClient, err := userplannerplandetail.NewUserPlannerPlanDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanDetail client: %+v", err)
	}
	configureFunc(userPlannerPlanDetailClient.Client)

	userPlannerPlanTaskAssignedToTaskBoardFormatClient, err := userplannerplantaskassignedtotaskboardformat.NewUserPlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerPlanTaskAssignedToTaskBoardFormatClient.Client)

	userPlannerPlanTaskBucketTaskBoardFormatClient, err := userplannerplantaskbuckettaskboardformat.NewUserPlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerPlanTaskBucketTaskBoardFormatClient.Client)

	userPlannerPlanTaskClient, err := userplannerplantask.NewUserPlannerPlanTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanTask client: %+v", err)
	}
	configureFunc(userPlannerPlanTaskClient.Client)

	userPlannerPlanTaskDetailClient, err := userplannerplantaskdetail.NewUserPlannerPlanTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanTaskDetail client: %+v", err)
	}
	configureFunc(userPlannerPlanTaskDetailClient.Client)

	userPlannerPlanTaskProgressTaskBoardFormatClient, err := userplannerplantaskprogresstaskboardformat.NewUserPlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerPlanTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerPlanTaskProgressTaskBoardFormatClient.Client)

	userPlannerTaskAssignedToTaskBoardFormatClient, err := userplannertaskassignedtotaskboardformat.NewUserPlannerTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerTaskAssignedToTaskBoardFormatClient.Client)

	userPlannerTaskBucketTaskBoardFormatClient, err := userplannertaskbuckettaskboardformat.NewUserPlannerTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerTaskBucketTaskBoardFormatClient.Client)

	userPlannerTaskClient, err := userplannertask.NewUserPlannerTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerTask client: %+v", err)
	}
	configureFunc(userPlannerTaskClient.Client)

	userPlannerTaskDetailClient, err := userplannertaskdetail.NewUserPlannerTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerTaskDetail client: %+v", err)
	}
	configureFunc(userPlannerTaskDetailClient.Client)

	userPlannerTaskProgressTaskBoardFormatClient, err := userplannertaskprogresstaskboardformat.NewUserPlannerTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(userPlannerTaskProgressTaskBoardFormatClient.Client)

	userPresenceClient, err := userpresence.NewUserPresenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPresence client: %+v", err)
	}
	configureFunc(userPresenceClient.Client)

	userRegisteredDeviceClient, err := userregistereddevice.NewUserRegisteredDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserRegisteredDevice client: %+v", err)
	}
	configureFunc(userRegisteredDeviceClient.Client)

	userScopedRoleMemberOfClient, err := userscopedrolememberof.NewUserScopedRoleMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserScopedRoleMemberOf client: %+v", err)
	}
	configureFunc(userScopedRoleMemberOfClient.Client)

	userSettingClient, err := usersetting.NewUserSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSetting client: %+v", err)
	}
	configureFunc(userSettingClient.Client)

	userSettingShiftPreferenceClient, err := usersettingshiftpreference.NewUserSettingShiftPreferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSettingShiftPreference client: %+v", err)
	}
	configureFunc(userSettingShiftPreferenceClient.Client)

	userTeamworkAssociatedTeamClient, err := userteamworkassociatedteam.NewUserTeamworkAssociatedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTeamworkAssociatedTeam client: %+v", err)
	}
	configureFunc(userTeamworkAssociatedTeamClient.Client)

	userTeamworkAssociatedTeamTeamClient, err := userteamworkassociatedteamteam.NewUserTeamworkAssociatedTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTeamworkAssociatedTeamTeam client: %+v", err)
	}
	configureFunc(userTeamworkAssociatedTeamTeamClient.Client)

	userTeamworkClient, err := userteamwork.NewUserTeamworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTeamwork client: %+v", err)
	}
	configureFunc(userTeamworkClient.Client)

	userTeamworkInstalledAppChatClient, err := userteamworkinstalledappchat.NewUserTeamworkInstalledAppChatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTeamworkInstalledAppChat client: %+v", err)
	}
	configureFunc(userTeamworkInstalledAppChatClient.Client)

	userTeamworkInstalledAppClient, err := userteamworkinstalledapp.NewUserTeamworkInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTeamworkInstalledApp client: %+v", err)
	}
	configureFunc(userTeamworkInstalledAppClient.Client)

	userTeamworkInstalledAppTeamsAppClient, err := userteamworkinstalledappteamsapp.NewUserTeamworkInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTeamworkInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(userTeamworkInstalledAppTeamsAppClient.Client)

	userTeamworkInstalledAppTeamsAppDefinitionClient, err := userteamworkinstalledappteamsappdefinition.NewUserTeamworkInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTeamworkInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(userTeamworkInstalledAppTeamsAppDefinitionClient.Client)

	userTodoClient, err := usertodo.NewUserTodoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodo client: %+v", err)
	}
	configureFunc(userTodoClient.Client)

	userTodoListClient, err := usertodolist.NewUserTodoListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoList client: %+v", err)
	}
	configureFunc(userTodoListClient.Client)

	userTodoListExtensionClient, err := usertodolistextension.NewUserTodoListExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListExtension client: %+v", err)
	}
	configureFunc(userTodoListExtensionClient.Client)

	userTodoListTaskAttachmentClient, err := usertodolisttaskattachment.NewUserTodoListTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListTaskAttachment client: %+v", err)
	}
	configureFunc(userTodoListTaskAttachmentClient.Client)

	userTodoListTaskAttachmentSessionClient, err := usertodolisttaskattachmentsession.NewUserTodoListTaskAttachmentSessionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListTaskAttachmentSession client: %+v", err)
	}
	configureFunc(userTodoListTaskAttachmentSessionClient.Client)

	userTodoListTaskAttachmentSessionContentClient, err := usertodolisttaskattachmentsessioncontent.NewUserTodoListTaskAttachmentSessionContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListTaskAttachmentSessionContent client: %+v", err)
	}
	configureFunc(userTodoListTaskAttachmentSessionContentClient.Client)

	userTodoListTaskChecklistItemClient, err := usertodolisttaskchecklistitem.NewUserTodoListTaskChecklistItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListTaskChecklistItem client: %+v", err)
	}
	configureFunc(userTodoListTaskChecklistItemClient.Client)

	userTodoListTaskClient, err := usertodolisttask.NewUserTodoListTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListTask client: %+v", err)
	}
	configureFunc(userTodoListTaskClient.Client)

	userTodoListTaskExtensionClient, err := usertodolisttaskextension.NewUserTodoListTaskExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListTaskExtension client: %+v", err)
	}
	configureFunc(userTodoListTaskExtensionClient.Client)

	userTodoListTaskLinkedResourceClient, err := usertodolisttasklinkedresource.NewUserTodoListTaskLinkedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTodoListTaskLinkedResource client: %+v", err)
	}
	configureFunc(userTodoListTaskLinkedResourceClient.Client)

	userTransitiveMemberOfClient, err := usertransitivememberof.NewUserTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTransitiveMemberOf client: %+v", err)
	}
	configureFunc(userTransitiveMemberOfClient.Client)

	return &Client{
		User:                                           userClient,
		UserActivity:                                   userActivityClient,
		UserActivityHistoryItem:                        userActivityHistoryItemClient,
		UserActivityHistoryItemActivity:                userActivityHistoryItemActivityClient,
		UserAgreementAcceptance:                        userAgreementAcceptanceClient,
		UserAppRoleAssignment:                          userAppRoleAssignmentClient,
		UserAuthentication:                             userAuthenticationClient,
		UserAuthenticationEmailMethod:                  userAuthenticationEmailMethodClient,
		UserAuthenticationFido2Method:                  userAuthenticationFido2MethodClient,
		UserAuthenticationMethod:                       userAuthenticationMethodClient,
		UserAuthenticationMicrosoftAuthenticatorMethod: userAuthenticationMicrosoftAuthenticatorMethodClient,
		UserAuthenticationMicrosoftAuthenticatorMethodDevice:  userAuthenticationMicrosoftAuthenticatorMethodDeviceClient,
		UserAuthenticationOperation:                           userAuthenticationOperationClient,
		UserAuthenticationPasswordMethod:                      userAuthenticationPasswordMethodClient,
		UserAuthenticationPhoneMethod:                         userAuthenticationPhoneMethodClient,
		UserAuthenticationSoftwareOathMethod:                  userAuthenticationSoftwareOathMethodClient,
		UserAuthenticationTemporaryAccessPassMethod:           userAuthenticationTemporaryAccessPassMethodClient,
		UserAuthenticationWindowsHelloForBusinessMethod:       userAuthenticationWindowsHelloForBusinessMethodClient,
		UserAuthenticationWindowsHelloForBusinessMethodDevice: userAuthenticationWindowsHelloForBusinessMethodDeviceClient,
		UserCalendar:                                            userCalendarClient,
		UserCalendarCalendarPermission:                          userCalendarCalendarPermissionClient,
		UserCalendarCalendarView:                                userCalendarCalendarViewClient,
		UserCalendarCalendarViewAttachment:                      userCalendarCalendarViewAttachmentClient,
		UserCalendarCalendarViewCalendar:                        userCalendarCalendarViewCalendarClient,
		UserCalendarCalendarViewExtension:                       userCalendarCalendarViewExtensionClient,
		UserCalendarCalendarViewInstance:                        userCalendarCalendarViewInstanceClient,
		UserCalendarCalendarViewInstanceAttachment:              userCalendarCalendarViewInstanceAttachmentClient,
		UserCalendarCalendarViewInstanceCalendar:                userCalendarCalendarViewInstanceCalendarClient,
		UserCalendarCalendarViewInstanceExtension:               userCalendarCalendarViewInstanceExtensionClient,
		UserCalendarEvent:                                       userCalendarEventClient,
		UserCalendarEventAttachment:                             userCalendarEventAttachmentClient,
		UserCalendarEventCalendar:                               userCalendarEventCalendarClient,
		UserCalendarEventExtension:                              userCalendarEventExtensionClient,
		UserCalendarEventInstance:                               userCalendarEventInstanceClient,
		UserCalendarEventInstanceAttachment:                     userCalendarEventInstanceAttachmentClient,
		UserCalendarEventInstanceCalendar:                       userCalendarEventInstanceCalendarClient,
		UserCalendarEventInstanceExtension:                      userCalendarEventInstanceExtensionClient,
		UserCalendarGroup:                                       userCalendarGroupClient,
		UserCalendarGroupCalendar:                               userCalendarGroupCalendarClient,
		UserCalendarGroupCalendarCalendarPermission:             userCalendarGroupCalendarCalendarPermissionClient,
		UserCalendarGroupCalendarCalendarView:                   userCalendarGroupCalendarCalendarViewClient,
		UserCalendarGroupCalendarCalendarViewAttachment:         userCalendarGroupCalendarCalendarViewAttachmentClient,
		UserCalendarGroupCalendarCalendarViewCalendar:           userCalendarGroupCalendarCalendarViewCalendarClient,
		UserCalendarGroupCalendarCalendarViewExtension:          userCalendarGroupCalendarCalendarViewExtensionClient,
		UserCalendarGroupCalendarCalendarViewInstance:           userCalendarGroupCalendarCalendarViewInstanceClient,
		UserCalendarGroupCalendarCalendarViewInstanceAttachment: userCalendarGroupCalendarCalendarViewInstanceAttachmentClient,
		UserCalendarGroupCalendarCalendarViewInstanceCalendar:   userCalendarGroupCalendarCalendarViewInstanceCalendarClient,
		UserCalendarGroupCalendarCalendarViewInstanceExtension:  userCalendarGroupCalendarCalendarViewInstanceExtensionClient,
		UserCalendarGroupCalendarEvent:                          userCalendarGroupCalendarEventClient,
		UserCalendarGroupCalendarEventAttachment:                userCalendarGroupCalendarEventAttachmentClient,
		UserCalendarGroupCalendarEventCalendar:                  userCalendarGroupCalendarEventCalendarClient,
		UserCalendarGroupCalendarEventExtension:                 userCalendarGroupCalendarEventExtensionClient,
		UserCalendarGroupCalendarEventInstance:                  userCalendarGroupCalendarEventInstanceClient,
		UserCalendarGroupCalendarEventInstanceAttachment:        userCalendarGroupCalendarEventInstanceAttachmentClient,
		UserCalendarGroupCalendarEventInstanceCalendar:          userCalendarGroupCalendarEventInstanceCalendarClient,
		UserCalendarGroupCalendarEventInstanceExtension:         userCalendarGroupCalendarEventInstanceExtensionClient,
		UserCalendarView:                                        userCalendarViewClient,
		UserCalendarViewAttachment:                              userCalendarViewAttachmentClient,
		UserCalendarViewCalendar:                                userCalendarViewCalendarClient,
		UserCalendarViewExtension:                               userCalendarViewExtensionClient,
		UserCalendarViewInstance:                                userCalendarViewInstanceClient,
		UserCalendarViewInstanceAttachment:                      userCalendarViewInstanceAttachmentClient,
		UserCalendarViewInstanceCalendar:                        userCalendarViewInstanceCalendarClient,
		UserCalendarViewInstanceExtension:                       userCalendarViewInstanceExtensionClient,
		UserChat:                                                userChatClient,
		UserChatInstalledApp:                                    userChatInstalledAppClient,
		UserChatInstalledAppTeamsApp:                            userChatInstalledAppTeamsAppClient,
		UserChatInstalledAppTeamsAppDefinition:                  userChatInstalledAppTeamsAppDefinitionClient,
		UserChatLastMessagePreview:                              userChatLastMessagePreviewClient,
		UserChatMember:                                          userChatMemberClient,
		UserChatMessage:                                         userChatMessageClient,
		UserChatMessageHostedContent:                            userChatMessageHostedContentClient,
		UserChatMessageReply:                                    userChatMessageReplyClient,
		UserChatMessageReplyHostedContent:                       userChatMessageReplyHostedContentClient,
		UserChatPermissionGrant:                                 userChatPermissionGrantClient,
		UserChatPinnedMessage:                                   userChatPinnedMessageClient,
		UserChatPinnedMessageMessage:                            userChatPinnedMessageMessageClient,
		UserChatTab:                                             userChatTabClient,
		UserChatTabTeamsApp:                                     userChatTabTeamsAppClient,
		UserContact:                                             userContactClient,
		UserContactExtension:                                    userContactExtensionClient,
		UserContactFolder:                                       userContactFolderClient,
		UserContactFolderChildFolder:                            userContactFolderChildFolderClient,
		UserContactFolderChildFolderContact:                     userContactFolderChildFolderContactClient,
		UserContactFolderChildFolderContactExtension:            userContactFolderChildFolderContactExtensionClient,
		UserContactFolderChildFolderContactPhoto:                userContactFolderChildFolderContactPhotoClient,
		UserContactFolderContact:                                userContactFolderContactClient,
		UserContactFolderContactExtension:                       userContactFolderContactExtensionClient,
		UserContactFolderContactPhoto:                           userContactFolderContactPhotoClient,
		UserContactPhoto:                                        userContactPhotoClient,
		UserCreatedObject:                                       userCreatedObjectClient,
		UserDeviceManagementTroubleshootingEvent:                userDeviceManagementTroubleshootingEventClient,
		UserDirectReport:                                        userDirectReportClient,
		UserDrive:                                               userDriveClient,
		UserEmployeeExperience:                                  userEmployeeExperienceClient,
		UserEmployeeExperienceLearningCourseActivity:            userEmployeeExperienceLearningCourseActivityClient,
		UserEvent:                                                   userEventClient,
		UserEventAttachment:                                         userEventAttachmentClient,
		UserEventCalendar:                                           userEventCalendarClient,
		UserEventExtension:                                          userEventExtensionClient,
		UserEventInstance:                                           userEventInstanceClient,
		UserEventInstanceAttachment:                                 userEventInstanceAttachmentClient,
		UserEventInstanceCalendar:                                   userEventInstanceCalendarClient,
		UserEventInstanceExtension:                                  userEventInstanceExtensionClient,
		UserExtension:                                               userExtensionClient,
		UserFollowedSite:                                            userFollowedSiteClient,
		UserInferenceClassification:                                 userInferenceClassificationClient,
		UserInferenceClassificationOverride:                         userInferenceClassificationOverrideClient,
		UserInsight:                                                 userInsightClient,
		UserInsightShared:                                           userInsightSharedClient,
		UserInsightSharedLastSharedMethod:                           userInsightSharedLastSharedMethodClient,
		UserInsightSharedResource:                                   userInsightSharedResourceClient,
		UserInsightTrending:                                         userInsightTrendingClient,
		UserInsightTrendingResource:                                 userInsightTrendingResourceClient,
		UserInsightUsed:                                             userInsightUsedClient,
		UserInsightUsedResource:                                     userInsightUsedResourceClient,
		UserJoinedTeam:                                              userJoinedTeamClient,
		UserJoinedTeamAllChannel:                                    userJoinedTeamAllChannelClient,
		UserJoinedTeamChannel:                                       userJoinedTeamChannelClient,
		UserJoinedTeamChannelFilesFolder:                            userJoinedTeamChannelFilesFolderClient,
		UserJoinedTeamChannelFilesFolderContent:                     userJoinedTeamChannelFilesFolderContentClient,
		UserJoinedTeamChannelMember:                                 userJoinedTeamChannelMemberClient,
		UserJoinedTeamChannelMessage:                                userJoinedTeamChannelMessageClient,
		UserJoinedTeamChannelMessageHostedContent:                   userJoinedTeamChannelMessageHostedContentClient,
		UserJoinedTeamChannelMessageReply:                           userJoinedTeamChannelMessageReplyClient,
		UserJoinedTeamChannelMessageReplyHostedContent:              userJoinedTeamChannelMessageReplyHostedContentClient,
		UserJoinedTeamChannelSharedWithTeam:                         userJoinedTeamChannelSharedWithTeamClient,
		UserJoinedTeamChannelSharedWithTeamAllowedMember:            userJoinedTeamChannelSharedWithTeamAllowedMemberClient,
		UserJoinedTeamChannelSharedWithTeamTeam:                     userJoinedTeamChannelSharedWithTeamTeamClient,
		UserJoinedTeamChannelTab:                                    userJoinedTeamChannelTabClient,
		UserJoinedTeamChannelTabTeamsApp:                            userJoinedTeamChannelTabTeamsAppClient,
		UserJoinedTeamGroup:                                         userJoinedTeamGroupClient,
		UserJoinedTeamIncomingChannel:                               userJoinedTeamIncomingChannelClient,
		UserJoinedTeamInstalledApp:                                  userJoinedTeamInstalledAppClient,
		UserJoinedTeamInstalledAppTeamsApp:                          userJoinedTeamInstalledAppTeamsAppClient,
		UserJoinedTeamInstalledAppTeamsAppDefinition:                userJoinedTeamInstalledAppTeamsAppDefinitionClient,
		UserJoinedTeamMember:                                        userJoinedTeamMemberClient,
		UserJoinedTeamOperation:                                     userJoinedTeamOperationClient,
		UserJoinedTeamPermissionGrant:                               userJoinedTeamPermissionGrantClient,
		UserJoinedTeamPhoto:                                         userJoinedTeamPhotoClient,
		UserJoinedTeamPrimaryChannel:                                userJoinedTeamPrimaryChannelClient,
		UserJoinedTeamPrimaryChannelFilesFolder:                     userJoinedTeamPrimaryChannelFilesFolderClient,
		UserJoinedTeamPrimaryChannelFilesFolderContent:              userJoinedTeamPrimaryChannelFilesFolderContentClient,
		UserJoinedTeamPrimaryChannelMember:                          userJoinedTeamPrimaryChannelMemberClient,
		UserJoinedTeamPrimaryChannelMessage:                         userJoinedTeamPrimaryChannelMessageClient,
		UserJoinedTeamPrimaryChannelMessageHostedContent:            userJoinedTeamPrimaryChannelMessageHostedContentClient,
		UserJoinedTeamPrimaryChannelMessageReply:                    userJoinedTeamPrimaryChannelMessageReplyClient,
		UserJoinedTeamPrimaryChannelMessageReplyHostedContent:       userJoinedTeamPrimaryChannelMessageReplyHostedContentClient,
		UserJoinedTeamPrimaryChannelSharedWithTeam:                  userJoinedTeamPrimaryChannelSharedWithTeamClient,
		UserJoinedTeamPrimaryChannelSharedWithTeamAllowedMember:     userJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient,
		UserJoinedTeamPrimaryChannelSharedWithTeamTeam:              userJoinedTeamPrimaryChannelSharedWithTeamTeamClient,
		UserJoinedTeamPrimaryChannelTab:                             userJoinedTeamPrimaryChannelTabClient,
		UserJoinedTeamPrimaryChannelTabTeamsApp:                     userJoinedTeamPrimaryChannelTabTeamsAppClient,
		UserJoinedTeamSchedule:                                      userJoinedTeamScheduleClient,
		UserJoinedTeamScheduleOfferShiftRequest:                     userJoinedTeamScheduleOfferShiftRequestClient,
		UserJoinedTeamScheduleOpenShift:                             userJoinedTeamScheduleOpenShiftClient,
		UserJoinedTeamScheduleOpenShiftChangeRequest:                userJoinedTeamScheduleOpenShiftChangeRequestClient,
		UserJoinedTeamScheduleSchedulingGroup:                       userJoinedTeamScheduleSchedulingGroupClient,
		UserJoinedTeamScheduleShift:                                 userJoinedTeamScheduleShiftClient,
		UserJoinedTeamScheduleSwapShiftsChangeRequest:               userJoinedTeamScheduleSwapShiftsChangeRequestClient,
		UserJoinedTeamScheduleTimeOffReason:                         userJoinedTeamScheduleTimeOffReasonClient,
		UserJoinedTeamScheduleTimeOffRequest:                        userJoinedTeamScheduleTimeOffRequestClient,
		UserJoinedTeamScheduleTimesOff:                              userJoinedTeamScheduleTimesOffClient,
		UserJoinedTeamTag:                                           userJoinedTeamTagClient,
		UserJoinedTeamTagMember:                                     userJoinedTeamTagMemberClient,
		UserJoinedTeamTemplate:                                      userJoinedTeamTemplateClient,
		UserLicenseDetail:                                           userLicenseDetailClient,
		UserMailFolder:                                              userMailFolderClient,
		UserMailFolderChildFolder:                                   userMailFolderChildFolderClient,
		UserMailFolderChildFolderMessage:                            userMailFolderChildFolderMessageClient,
		UserMailFolderChildFolderMessageAttachment:                  userMailFolderChildFolderMessageAttachmentClient,
		UserMailFolderChildFolderMessageExtension:                   userMailFolderChildFolderMessageExtensionClient,
		UserMailFolderChildFolderMessageRule:                        userMailFolderChildFolderMessageRuleClient,
		UserMailFolderMessage:                                       userMailFolderMessageClient,
		UserMailFolderMessageAttachment:                             userMailFolderMessageAttachmentClient,
		UserMailFolderMessageExtension:                              userMailFolderMessageExtensionClient,
		UserMailFolderMessageRule:                                   userMailFolderMessageRuleClient,
		UserMailboxSetting:                                          userMailboxSettingClient,
		UserManagedAppRegistration:                                  userManagedAppRegistrationClient,
		UserManagedDevice:                                           userManagedDeviceClient,
		UserManagedDeviceDeviceCategory:                             userManagedDeviceDeviceCategoryClient,
		UserManagedDeviceDeviceCompliancePolicyState:                userManagedDeviceDeviceCompliancePolicyStateClient,
		UserManagedDeviceDeviceConfigurationState:                   userManagedDeviceDeviceConfigurationStateClient,
		UserManagedDeviceLogCollectionRequest:                       userManagedDeviceLogCollectionRequestClient,
		UserManagedDeviceUser:                                       userManagedDeviceUserClient,
		UserManagedDeviceWindowsProtectionState:                     userManagedDeviceWindowsProtectionStateClient,
		UserManagedDeviceWindowsProtectionStateDetectedMalwareState: userManagedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		UserManager:                     userManagerClient,
		UserMemberOf:                    userMemberOfClient,
		UserMessage:                     userMessageClient,
		UserMessageAttachment:           userMessageAttachmentClient,
		UserMessageExtension:            userMessageExtensionClient,
		UserOauth2PermissionGrant:       userOauth2PermissionGrantClient,
		UserOnenote:                     userOnenoteClient,
		UserOnenoteNotebook:             userOnenoteNotebookClient,
		UserOnenoteNotebookSection:      userOnenoteNotebookSectionClient,
		UserOnenoteNotebookSectionGroup: userOnenoteNotebookSectionGroupClient,
		UserOnenoteNotebookSectionGroupParentNotebook:            userOnenoteNotebookSectionGroupParentNotebookClient,
		UserOnenoteNotebookSectionGroupParentSectionGroup:        userOnenoteNotebookSectionGroupParentSectionGroupClient,
		UserOnenoteNotebookSectionGroupSection:                   userOnenoteNotebookSectionGroupSectionClient,
		UserOnenoteNotebookSectionGroupSectionGroup:              userOnenoteNotebookSectionGroupSectionGroupClient,
		UserOnenoteNotebookSectionGroupSectionPage:               userOnenoteNotebookSectionGroupSectionPageClient,
		UserOnenoteNotebookSectionGroupSectionPageContent:        userOnenoteNotebookSectionGroupSectionPageContentClient,
		UserOnenoteNotebookSectionGroupSectionPageParentNotebook: userOnenoteNotebookSectionGroupSectionPageParentNotebookClient,
		UserOnenoteNotebookSectionGroupSectionPageParentSection:  userOnenoteNotebookSectionGroupSectionPageParentSectionClient,
		UserOnenoteNotebookSectionGroupSectionParentNotebook:     userOnenoteNotebookSectionGroupSectionParentNotebookClient,
		UserOnenoteNotebookSectionGroupSectionParentSectionGroup: userOnenoteNotebookSectionGroupSectionParentSectionGroupClient,
		UserOnenoteNotebookSectionPage:                           userOnenoteNotebookSectionPageClient,
		UserOnenoteNotebookSectionPageContent:                    userOnenoteNotebookSectionPageContentClient,
		UserOnenoteNotebookSectionPageParentNotebook:             userOnenoteNotebookSectionPageParentNotebookClient,
		UserOnenoteNotebookSectionPageParentSection:              userOnenoteNotebookSectionPageParentSectionClient,
		UserOnenoteNotebookSectionParentNotebook:                 userOnenoteNotebookSectionParentNotebookClient,
		UserOnenoteNotebookSectionParentSectionGroup:             userOnenoteNotebookSectionParentSectionGroupClient,
		UserOnenoteOperation:                                     userOnenoteOperationClient,
		UserOnenotePage:                                          userOnenotePageClient,
		UserOnenotePageContent:                                   userOnenotePageContentClient,
		UserOnenotePageParentNotebook:                            userOnenotePageParentNotebookClient,
		UserOnenotePageParentSection:                             userOnenotePageParentSectionClient,
		UserOnenoteResource:                                      userOnenoteResourceClient,
		UserOnenoteResourceContent:                               userOnenoteResourceContentClient,
		UserOnenoteSection:                                       userOnenoteSectionClient,
		UserOnenoteSectionGroup:                                  userOnenoteSectionGroupClient,
		UserOnenoteSectionGroupParentNotebook:                    userOnenoteSectionGroupParentNotebookClient,
		UserOnenoteSectionGroupParentSectionGroup:                userOnenoteSectionGroupParentSectionGroupClient,
		UserOnenoteSectionGroupSection:                           userOnenoteSectionGroupSectionClient,
		UserOnenoteSectionGroupSectionGroup:                      userOnenoteSectionGroupSectionGroupClient,
		UserOnenoteSectionGroupSectionPage:                       userOnenoteSectionGroupSectionPageClient,
		UserOnenoteSectionGroupSectionPageContent:                userOnenoteSectionGroupSectionPageContentClient,
		UserOnenoteSectionGroupSectionPageParentNotebook:         userOnenoteSectionGroupSectionPageParentNotebookClient,
		UserOnenoteSectionGroupSectionPageParentSection:          userOnenoteSectionGroupSectionPageParentSectionClient,
		UserOnenoteSectionGroupSectionParentNotebook:             userOnenoteSectionGroupSectionParentNotebookClient,
		UserOnenoteSectionGroupSectionParentSectionGroup:         userOnenoteSectionGroupSectionParentSectionGroupClient,
		UserOnenoteSectionPage:                                   userOnenoteSectionPageClient,
		UserOnenoteSectionPageContent:                            userOnenoteSectionPageContentClient,
		UserOnenoteSectionPageParentNotebook:                     userOnenoteSectionPageParentNotebookClient,
		UserOnenoteSectionPageParentSection:                      userOnenoteSectionPageParentSectionClient,
		UserOnenoteSectionParentNotebook:                         userOnenoteSectionParentNotebookClient,
		UserOnenoteSectionParentSectionGroup:                     userOnenoteSectionParentSectionGroupClient,
		UserOnlineMeeting:                                        userOnlineMeetingClient,
		UserOnlineMeetingAttendanceReport:                        userOnlineMeetingAttendanceReportClient,
		UserOnlineMeetingAttendanceReportAttendanceRecord:        userOnlineMeetingAttendanceReportAttendanceRecordClient,
		UserOnlineMeetingAttendeeReport:                          userOnlineMeetingAttendeeReportClient,
		UserOutlook:                                              userOutlookClient,
		UserOutlookMasterCategory:                                userOutlookMasterCategoryClient,
		UserOwnedDevice:                                          userOwnedDeviceClient,
		UserOwnedObject:                                          userOwnedObjectClient,
		UserPeople:                                               userPeopleClient,
		UserPhoto:                                                userPhotoClient,
		UserPlanner:                                              userPlannerClient,
		UserPlannerPlan:                                          userPlannerPlanClient,
		UserPlannerPlanBucket:                                    userPlannerPlanBucketClient,
		UserPlannerPlanBucketTask:                                userPlannerPlanBucketTaskClient,
		UserPlannerPlanBucketTaskAssignedToTaskBoardFormat:       userPlannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		UserPlannerPlanBucketTaskBucketTaskBoardFormat:           userPlannerPlanBucketTaskBucketTaskBoardFormatClient,
		UserPlannerPlanBucketTaskDetail:                          userPlannerPlanBucketTaskDetailClient,
		UserPlannerPlanBucketTaskProgressTaskBoardFormat:         userPlannerPlanBucketTaskProgressTaskBoardFormatClient,
		UserPlannerPlanDetail:                                    userPlannerPlanDetailClient,
		UserPlannerPlanTask:                                      userPlannerPlanTaskClient,
		UserPlannerPlanTaskAssignedToTaskBoardFormat:             userPlannerPlanTaskAssignedToTaskBoardFormatClient,
		UserPlannerPlanTaskBucketTaskBoardFormat:                 userPlannerPlanTaskBucketTaskBoardFormatClient,
		UserPlannerPlanTaskDetail:                                userPlannerPlanTaskDetailClient,
		UserPlannerPlanTaskProgressTaskBoardFormat:               userPlannerPlanTaskProgressTaskBoardFormatClient,
		UserPlannerTask:                                          userPlannerTaskClient,
		UserPlannerTaskAssignedToTaskBoardFormat:                 userPlannerTaskAssignedToTaskBoardFormatClient,
		UserPlannerTaskBucketTaskBoardFormat:                     userPlannerTaskBucketTaskBoardFormatClient,
		UserPlannerTaskDetail:                                    userPlannerTaskDetailClient,
		UserPlannerTaskProgressTaskBoardFormat:                   userPlannerTaskProgressTaskBoardFormatClient,
		UserPresence:                                             userPresenceClient,
		UserRegisteredDevice:                                     userRegisteredDeviceClient,
		UserScopedRoleMemberOf:                                   userScopedRoleMemberOfClient,
		UserSetting:                                              userSettingClient,
		UserSettingShiftPreference:                               userSettingShiftPreferenceClient,
		UserTeamwork:                                             userTeamworkClient,
		UserTeamworkAssociatedTeam:                               userTeamworkAssociatedTeamClient,
		UserTeamworkAssociatedTeamTeam:                           userTeamworkAssociatedTeamTeamClient,
		UserTeamworkInstalledApp:                                 userTeamworkInstalledAppClient,
		UserTeamworkInstalledAppChat:                             userTeamworkInstalledAppChatClient,
		UserTeamworkInstalledAppTeamsApp:                         userTeamworkInstalledAppTeamsAppClient,
		UserTeamworkInstalledAppTeamsAppDefinition:               userTeamworkInstalledAppTeamsAppDefinitionClient,
		UserTodo:                                 userTodoClient,
		UserTodoList:                             userTodoListClient,
		UserTodoListExtension:                    userTodoListExtensionClient,
		UserTodoListTask:                         userTodoListTaskClient,
		UserTodoListTaskAttachment:               userTodoListTaskAttachmentClient,
		UserTodoListTaskAttachmentSession:        userTodoListTaskAttachmentSessionClient,
		UserTodoListTaskAttachmentSessionContent: userTodoListTaskAttachmentSessionContentClient,
		UserTodoListTaskChecklistItem:            userTodoListTaskChecklistItemClient,
		UserTodoListTaskExtension:                userTodoListTaskExtensionClient,
		UserTodoListTaskLinkedResource:           userTodoListTaskLinkedResourceClient,
		UserTransitiveMemberOf:                   userTransitiveMemberOfClient,
	}, nil
}
