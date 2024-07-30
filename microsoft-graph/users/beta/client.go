package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/activity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/activityhistoryitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/activityhistoryitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/agreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/analytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/analyticactivitystatistic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/appconsentrequestsforapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/appconsentrequestsforapprovaluserconsentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/appconsentrequestsforapprovaluserconsentrequestapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/appconsentrequestsforapprovaluserconsentrequestapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/approleassignedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/approleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/approval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/approvalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationemailmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationfido2method"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationpasswordlessmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationpasswordlessmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationpasswordmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationphonemethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationplatformcredentialmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationplatformcredentialmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationsigninpreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationsoftwareoathmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationtemporaryaccesspassmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationwindowshelloforbusinessmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationwindowshelloforbusinessmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendargroupcalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/calendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatlastmessagepreview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatpermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatpinnedmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chatpinnedmessagemessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chattab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/chattabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/cloudpc"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfolderchildfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfolderchildfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfolderchildfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/contactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/createdobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/device"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/devicecommand"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/devicecommandresponsepayload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceenrollmentconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceenrollmentconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/devicemanagementtroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/devicememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceregisteredowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceregistereduser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/devicetransitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/deviceusageright"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/directreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/employeeexperience"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/employeeexperiencelearningcourseactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/event"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/eventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/extension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/followedsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/inferenceclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/inferenceclassificationoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionbitlocker"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionbitlockerrecoverykey"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectiondatalosspreventionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionpolicylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionsensitivitylabelsublabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionsensitivitypolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionthreatassessmentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionthreatassessmentrequestresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insightshared"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insightsharedlastsharedmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insightsharedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insighttrending"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insighttrendingresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insightused"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/insightusedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/joinedgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/joinedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/licensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfoldermessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderchildfolderuserconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfoldermessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mailfolderuserconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/managedappregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddeviceassignmentfilterevaluationstatusdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedetectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicehealthscriptstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicemanageddevicemobileappconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicesecuritybaselinestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicesecuritybaselinestatesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/memberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/message"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/messageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/messageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/messagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mobileappintentandstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mobileapptroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/mobileapptroubleshootingeventapplogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/notification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/oauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeeting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingalternativerecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingattendeereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingbroadcastrecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingmeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingmeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingrecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingrecordingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingregistrationcustomquestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingregistrationregistrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingtranscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingtranscriptcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeetingtranscriptmetadatacontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlookmastercategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskfoldertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskfoldertaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskgrouptaskfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskgrouptaskfoldertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/outlooktaskgrouptaskfoldertaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/owneddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/ownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/people"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/photo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/planner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerall"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerfavoriteplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerrecentplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerrosterplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannertaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannertaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannertaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannertaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/presence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileaccount"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileaddress"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileanniversary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileaward"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilecertification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileeducationalactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileemail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileinterest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilelanguage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilename"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilepatent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilephone"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileproject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilepublication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profileskill"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilewebaccount"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/profilewebsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/registereddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/scopedrolememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/security"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotectionlabelpolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotectionsensitivitylabelparent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/serviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/setting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingcontactmergesuggestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingiteminsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingregionalandlanguagesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingshiftpreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/sponsor"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamwork"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamworkassociatedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamworkassociatedteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamworkinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamworkinstalledappchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamworkinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/teamworkinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolistextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolisttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolisttaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolisttaskattachmentsession"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolisttaskattachmentsessioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolisttaskchecklistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolisttaskextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/todolisttasklinkedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/transitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/transitivereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usageright"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/user"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/virtualevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/virtualeventwebinar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/windowsinformationprotectiondeviceregistration"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	Activity                                                                 *activity.ActivityClient
	ActivityHistoryItem                                                      *activityhistoryitem.ActivityHistoryItemClient
	ActivityHistoryItemActivity                                              *activityhistoryitemactivity.ActivityHistoryItemActivityClient
	AgreementAcceptance                                                      *agreementacceptance.AgreementAcceptanceClient
	Analytic                                                                 *analytic.AnalyticClient
	AnalyticActivityStatistic                                                *analyticactivitystatistic.AnalyticActivityStatisticClient
	AppConsentRequestsForApproval                                            *appconsentrequestsforapproval.AppConsentRequestsForApprovalClient
	AppConsentRequestsForApprovalUserConsentRequest                          *appconsentrequestsforapprovaluserconsentrequest.AppConsentRequestsForApprovalUserConsentRequestClient
	AppConsentRequestsForApprovalUserConsentRequestApproval                  *appconsentrequestsforapprovaluserconsentrequestapproval.AppConsentRequestsForApprovalUserConsentRequestApprovalClient
	AppConsentRequestsForApprovalUserConsentRequestApprovalStep              *appconsentrequestsforapprovaluserconsentrequestapprovalstep.AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient
	AppRoleAssignedResource                                                  *approleassignedresource.AppRoleAssignedResourceClient
	AppRoleAssignment                                                        *approleassignment.AppRoleAssignmentClient
	Approval                                                                 *approval.ApprovalClient
	ApprovalStep                                                             *approvalstep.ApprovalStepClient
	Authentication                                                           *authentication.AuthenticationClient
	AuthenticationEmailMethod                                                *authenticationemailmethod.AuthenticationEmailMethodClient
	AuthenticationFido2Method                                                *authenticationfido2method.AuthenticationFido2MethodClient
	AuthenticationMethod                                                     *authenticationmethod.AuthenticationMethodClient
	AuthenticationMicrosoftAuthenticatorMethod                               *authenticationmicrosoftauthenticatormethod.AuthenticationMicrosoftAuthenticatorMethodClient
	AuthenticationMicrosoftAuthenticatorMethodDevice                         *authenticationmicrosoftauthenticatormethoddevice.AuthenticationMicrosoftAuthenticatorMethodDeviceClient
	AuthenticationOperation                                                  *authenticationoperation.AuthenticationOperationClient
	AuthenticationPasswordMethod                                             *authenticationpasswordmethod.AuthenticationPasswordMethodClient
	AuthenticationPasswordlessMicrosoftAuthenticatorMethod                   *authenticationpasswordlessmicrosoftauthenticatormethod.AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient
	AuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice             *authenticationpasswordlessmicrosoftauthenticatormethoddevice.AuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient
	AuthenticationPhoneMethod                                                *authenticationphonemethod.AuthenticationPhoneMethodClient
	AuthenticationPlatformCredentialMethod                                   *authenticationplatformcredentialmethod.AuthenticationPlatformCredentialMethodClient
	AuthenticationPlatformCredentialMethodDevice                             *authenticationplatformcredentialmethoddevice.AuthenticationPlatformCredentialMethodDeviceClient
	AuthenticationSignInPreference                                           *authenticationsigninpreference.AuthenticationSignInPreferenceClient
	AuthenticationSoftwareOathMethod                                         *authenticationsoftwareoathmethod.AuthenticationSoftwareOathMethodClient
	AuthenticationTemporaryAccessPassMethod                                  *authenticationtemporaryaccesspassmethod.AuthenticationTemporaryAccessPassMethodClient
	AuthenticationWindowsHelloForBusinessMethod                              *authenticationwindowshelloforbusinessmethod.AuthenticationWindowsHelloForBusinessMethodClient
	AuthenticationWindowsHelloForBusinessMethodDevice                        *authenticationwindowshelloforbusinessmethoddevice.AuthenticationWindowsHelloForBusinessMethodDeviceClient
	Calendar                                                                 *calendar.CalendarClient
	CalendarCalendarPermission                                               *calendarcalendarpermission.CalendarCalendarPermissionClient
	CalendarCalendarView                                                     *calendarcalendarview.CalendarCalendarViewClient
	CalendarCalendarViewAttachment                                           *calendarcalendarviewattachment.CalendarCalendarViewAttachmentClient
	CalendarCalendarViewCalendar                                             *calendarcalendarviewcalendar.CalendarCalendarViewCalendarClient
	CalendarCalendarViewExceptionOccurrence                                  *calendarcalendarviewexceptionoccurrence.CalendarCalendarViewExceptionOccurrenceClient
	CalendarCalendarViewExceptionOccurrenceAttachment                        *calendarcalendarviewexceptionoccurrenceattachment.CalendarCalendarViewExceptionOccurrenceAttachmentClient
	CalendarCalendarViewExceptionOccurrenceCalendar                          *calendarcalendarviewexceptionoccurrencecalendar.CalendarCalendarViewExceptionOccurrenceCalendarClient
	CalendarCalendarViewExceptionOccurrenceExtension                         *calendarcalendarviewexceptionoccurrenceextension.CalendarCalendarViewExceptionOccurrenceExtensionClient
	CalendarCalendarViewExceptionOccurrenceInstance                          *calendarcalendarviewexceptionoccurrenceinstance.CalendarCalendarViewExceptionOccurrenceInstanceClient
	CalendarCalendarViewExceptionOccurrenceInstanceAttachment                *calendarcalendarviewexceptionoccurrenceinstanceattachment.CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	CalendarCalendarViewExceptionOccurrenceInstanceCalendar                  *calendarcalendarviewexceptionoccurrenceinstancecalendar.CalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	CalendarCalendarViewExceptionOccurrenceInstanceExtension                 *calendarcalendarviewexceptionoccurrenceinstanceextension.CalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	CalendarCalendarViewExtension                                            *calendarcalendarviewextension.CalendarCalendarViewExtensionClient
	CalendarCalendarViewInstance                                             *calendarcalendarviewinstance.CalendarCalendarViewInstanceClient
	CalendarCalendarViewInstanceAttachment                                   *calendarcalendarviewinstanceattachment.CalendarCalendarViewInstanceAttachmentClient
	CalendarCalendarViewInstanceCalendar                                     *calendarcalendarviewinstancecalendar.CalendarCalendarViewInstanceCalendarClient
	CalendarCalendarViewInstanceExceptionOccurrence                          *calendarcalendarviewinstanceexceptionoccurrence.CalendarCalendarViewInstanceExceptionOccurrenceClient
	CalendarCalendarViewInstanceExceptionOccurrenceAttachment                *calendarcalendarviewinstanceexceptionoccurrenceattachment.CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	CalendarCalendarViewInstanceExceptionOccurrenceCalendar                  *calendarcalendarviewinstanceexceptionoccurrencecalendar.CalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	CalendarCalendarViewInstanceExceptionOccurrenceExtension                 *calendarcalendarviewinstanceexceptionoccurrenceextension.CalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	CalendarCalendarViewInstanceExtension                                    *calendarcalendarviewinstanceextension.CalendarCalendarViewInstanceExtensionClient
	CalendarEvent                                                            *calendarevent.CalendarEventClient
	CalendarEventAttachment                                                  *calendareventattachment.CalendarEventAttachmentClient
	CalendarEventCalendar                                                    *calendareventcalendar.CalendarEventCalendarClient
	CalendarEventExceptionOccurrence                                         *calendareventexceptionoccurrence.CalendarEventExceptionOccurrenceClient
	CalendarEventExceptionOccurrenceAttachment                               *calendareventexceptionoccurrenceattachment.CalendarEventExceptionOccurrenceAttachmentClient
	CalendarEventExceptionOccurrenceCalendar                                 *calendareventexceptionoccurrencecalendar.CalendarEventExceptionOccurrenceCalendarClient
	CalendarEventExceptionOccurrenceExtension                                *calendareventexceptionoccurrenceextension.CalendarEventExceptionOccurrenceExtensionClient
	CalendarEventExceptionOccurrenceInstance                                 *calendareventexceptionoccurrenceinstance.CalendarEventExceptionOccurrenceInstanceClient
	CalendarEventExceptionOccurrenceInstanceAttachment                       *calendareventexceptionoccurrenceinstanceattachment.CalendarEventExceptionOccurrenceInstanceAttachmentClient
	CalendarEventExceptionOccurrenceInstanceCalendar                         *calendareventexceptionoccurrenceinstancecalendar.CalendarEventExceptionOccurrenceInstanceCalendarClient
	CalendarEventExceptionOccurrenceInstanceExtension                        *calendareventexceptionoccurrenceinstanceextension.CalendarEventExceptionOccurrenceInstanceExtensionClient
	CalendarEventExtension                                                   *calendareventextension.CalendarEventExtensionClient
	CalendarEventInstance                                                    *calendareventinstance.CalendarEventInstanceClient
	CalendarEventInstanceAttachment                                          *calendareventinstanceattachment.CalendarEventInstanceAttachmentClient
	CalendarEventInstanceCalendar                                            *calendareventinstancecalendar.CalendarEventInstanceCalendarClient
	CalendarEventInstanceExceptionOccurrence                                 *calendareventinstanceexceptionoccurrence.CalendarEventInstanceExceptionOccurrenceClient
	CalendarEventInstanceExceptionOccurrenceAttachment                       *calendareventinstanceexceptionoccurrenceattachment.CalendarEventInstanceExceptionOccurrenceAttachmentClient
	CalendarEventInstanceExceptionOccurrenceCalendar                         *calendareventinstanceexceptionoccurrencecalendar.CalendarEventInstanceExceptionOccurrenceCalendarClient
	CalendarEventInstanceExceptionOccurrenceExtension                        *calendareventinstanceexceptionoccurrenceextension.CalendarEventInstanceExceptionOccurrenceExtensionClient
	CalendarEventInstanceExtension                                           *calendareventinstanceextension.CalendarEventInstanceExtensionClient
	CalendarGroup                                                            *calendargroup.CalendarGroupClient
	CalendarGroupCalendar                                                    *calendargroupcalendar.CalendarGroupCalendarClient
	CalendarGroupCalendarCalendarPermission                                  *calendargroupcalendarcalendarpermission.CalendarGroupCalendarCalendarPermissionClient
	CalendarGroupCalendarCalendarView                                        *calendargroupcalendarcalendarview.CalendarGroupCalendarCalendarViewClient
	CalendarGroupCalendarCalendarViewAttachment                              *calendargroupcalendarcalendarviewattachment.CalendarGroupCalendarCalendarViewAttachmentClient
	CalendarGroupCalendarCalendarViewCalendar                                *calendargroupcalendarcalendarviewcalendar.CalendarGroupCalendarCalendarViewCalendarClient
	CalendarGroupCalendarCalendarViewExceptionOccurrence                     *calendargroupcalendarcalendarviewexceptionoccurrence.CalendarGroupCalendarCalendarViewExceptionOccurrenceClient
	CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment           *calendargroupcalendarcalendarviewexceptionoccurrenceattachment.CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient
	CalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar             *calendargroupcalendarcalendarviewexceptionoccurrencecalendar.CalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient
	CalendarGroupCalendarCalendarViewExceptionOccurrenceExtension            *calendargroupcalendarcalendarviewexceptionoccurrenceextension.CalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient
	CalendarGroupCalendarCalendarViewExceptionOccurrenceInstance             *calendargroupcalendarcalendarviewexceptionoccurrenceinstance.CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient
	CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment   *calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar     *calendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar.CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension    *calendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension.CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	CalendarGroupCalendarCalendarViewExtension                               *calendargroupcalendarcalendarviewextension.CalendarGroupCalendarCalendarViewExtensionClient
	CalendarGroupCalendarCalendarViewInstance                                *calendargroupcalendarcalendarviewinstance.CalendarGroupCalendarCalendarViewInstanceClient
	CalendarGroupCalendarCalendarViewInstanceAttachment                      *calendargroupcalendarcalendarviewinstanceattachment.CalendarGroupCalendarCalendarViewInstanceAttachmentClient
	CalendarGroupCalendarCalendarViewInstanceCalendar                        *calendargroupcalendarcalendarviewinstancecalendar.CalendarGroupCalendarCalendarViewInstanceCalendarClient
	CalendarGroupCalendarCalendarViewInstanceExceptionOccurrence             *calendargroupcalendarcalendarviewinstanceexceptionoccurrence.CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient
	CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment   *calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar     *calendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar.CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension    *calendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension.CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	CalendarGroupCalendarCalendarViewInstanceExtension                       *calendargroupcalendarcalendarviewinstanceextension.CalendarGroupCalendarCalendarViewInstanceExtensionClient
	CalendarGroupCalendarEvent                                               *calendargroupcalendarevent.CalendarGroupCalendarEventClient
	CalendarGroupCalendarEventAttachment                                     *calendargroupcalendareventattachment.CalendarGroupCalendarEventAttachmentClient
	CalendarGroupCalendarEventCalendar                                       *calendargroupcalendareventcalendar.CalendarGroupCalendarEventCalendarClient
	CalendarGroupCalendarEventExceptionOccurrence                            *calendargroupcalendareventexceptionoccurrence.CalendarGroupCalendarEventExceptionOccurrenceClient
	CalendarGroupCalendarEventExceptionOccurrenceAttachment                  *calendargroupcalendareventexceptionoccurrenceattachment.CalendarGroupCalendarEventExceptionOccurrenceAttachmentClient
	CalendarGroupCalendarEventExceptionOccurrenceCalendar                    *calendargroupcalendareventexceptionoccurrencecalendar.CalendarGroupCalendarEventExceptionOccurrenceCalendarClient
	CalendarGroupCalendarEventExceptionOccurrenceExtension                   *calendargroupcalendareventexceptionoccurrenceextension.CalendarGroupCalendarEventExceptionOccurrenceExtensionClient
	CalendarGroupCalendarEventExceptionOccurrenceInstance                    *calendargroupcalendareventexceptionoccurrenceinstance.CalendarGroupCalendarEventExceptionOccurrenceInstanceClient
	CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment          *calendargroupcalendareventexceptionoccurrenceinstanceattachment.CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient
	CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar            *calendargroupcalendareventexceptionoccurrenceinstancecalendar.CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient
	CalendarGroupCalendarEventExceptionOccurrenceInstanceExtension           *calendargroupcalendareventexceptionoccurrenceinstanceextension.CalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient
	CalendarGroupCalendarEventExtension                                      *calendargroupcalendareventextension.CalendarGroupCalendarEventExtensionClient
	CalendarGroupCalendarEventInstance                                       *calendargroupcalendareventinstance.CalendarGroupCalendarEventInstanceClient
	CalendarGroupCalendarEventInstanceAttachment                             *calendargroupcalendareventinstanceattachment.CalendarGroupCalendarEventInstanceAttachmentClient
	CalendarGroupCalendarEventInstanceCalendar                               *calendargroupcalendareventinstancecalendar.CalendarGroupCalendarEventInstanceCalendarClient
	CalendarGroupCalendarEventInstanceExceptionOccurrence                    *calendargroupcalendareventinstanceexceptionoccurrence.CalendarGroupCalendarEventInstanceExceptionOccurrenceClient
	CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment          *calendargroupcalendareventinstanceexceptionoccurrenceattachment.CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient
	CalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar            *calendargroupcalendareventinstanceexceptionoccurrencecalendar.CalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient
	CalendarGroupCalendarEventInstanceExceptionOccurrenceExtension           *calendargroupcalendareventinstanceexceptionoccurrenceextension.CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient
	CalendarGroupCalendarEventInstanceExtension                              *calendargroupcalendareventinstanceextension.CalendarGroupCalendarEventInstanceExtensionClient
	CalendarView                                                             *calendarview.CalendarViewClient
	CalendarViewAttachment                                                   *calendarviewattachment.CalendarViewAttachmentClient
	CalendarViewCalendar                                                     *calendarviewcalendar.CalendarViewCalendarClient
	CalendarViewExceptionOccurrence                                          *calendarviewexceptionoccurrence.CalendarViewExceptionOccurrenceClient
	CalendarViewExceptionOccurrenceAttachment                                *calendarviewexceptionoccurrenceattachment.CalendarViewExceptionOccurrenceAttachmentClient
	CalendarViewExceptionOccurrenceCalendar                                  *calendarviewexceptionoccurrencecalendar.CalendarViewExceptionOccurrenceCalendarClient
	CalendarViewExceptionOccurrenceExtension                                 *calendarviewexceptionoccurrenceextension.CalendarViewExceptionOccurrenceExtensionClient
	CalendarViewExceptionOccurrenceInstance                                  *calendarviewexceptionoccurrenceinstance.CalendarViewExceptionOccurrenceInstanceClient
	CalendarViewExceptionOccurrenceInstanceAttachment                        *calendarviewexceptionoccurrenceinstanceattachment.CalendarViewExceptionOccurrenceInstanceAttachmentClient
	CalendarViewExceptionOccurrenceInstanceCalendar                          *calendarviewexceptionoccurrenceinstancecalendar.CalendarViewExceptionOccurrenceInstanceCalendarClient
	CalendarViewExceptionOccurrenceInstanceExtension                         *calendarviewexceptionoccurrenceinstanceextension.CalendarViewExceptionOccurrenceInstanceExtensionClient
	CalendarViewExtension                                                    *calendarviewextension.CalendarViewExtensionClient
	CalendarViewInstance                                                     *calendarviewinstance.CalendarViewInstanceClient
	CalendarViewInstanceAttachment                                           *calendarviewinstanceattachment.CalendarViewInstanceAttachmentClient
	CalendarViewInstanceCalendar                                             *calendarviewinstancecalendar.CalendarViewInstanceCalendarClient
	CalendarViewInstanceExceptionOccurrence                                  *calendarviewinstanceexceptionoccurrence.CalendarViewInstanceExceptionOccurrenceClient
	CalendarViewInstanceExceptionOccurrenceAttachment                        *calendarviewinstanceexceptionoccurrenceattachment.CalendarViewInstanceExceptionOccurrenceAttachmentClient
	CalendarViewInstanceExceptionOccurrenceCalendar                          *calendarviewinstanceexceptionoccurrencecalendar.CalendarViewInstanceExceptionOccurrenceCalendarClient
	CalendarViewInstanceExceptionOccurrenceExtension                         *calendarviewinstanceexceptionoccurrenceextension.CalendarViewInstanceExceptionOccurrenceExtensionClient
	CalendarViewInstanceExtension                                            *calendarviewinstanceextension.CalendarViewInstanceExtensionClient
	Chat                                                                     *chat.ChatClient
	ChatInstalledApp                                                         *chatinstalledapp.ChatInstalledAppClient
	ChatInstalledAppTeamsApp                                                 *chatinstalledappteamsapp.ChatInstalledAppTeamsAppClient
	ChatInstalledAppTeamsAppDefinition                                       *chatinstalledappteamsappdefinition.ChatInstalledAppTeamsAppDefinitionClient
	ChatLastMessagePreview                                                   *chatlastmessagepreview.ChatLastMessagePreviewClient
	ChatMember                                                               *chatmember.ChatMemberClient
	ChatMessage                                                              *chatmessage.ChatMessageClient
	ChatMessageHostedContent                                                 *chatmessagehostedcontent.ChatMessageHostedContentClient
	ChatMessageReply                                                         *chatmessagereply.ChatMessageReplyClient
	ChatMessageReplyHostedContent                                            *chatmessagereplyhostedcontent.ChatMessageReplyHostedContentClient
	ChatOperation                                                            *chatoperation.ChatOperationClient
	ChatPermissionGrant                                                      *chatpermissiongrant.ChatPermissionGrantClient
	ChatPinnedMessage                                                        *chatpinnedmessage.ChatPinnedMessageClient
	ChatPinnedMessageMessage                                                 *chatpinnedmessagemessage.ChatPinnedMessageMessageClient
	ChatTab                                                                  *chattab.ChatTabClient
	ChatTabTeamsApp                                                          *chattabteamsapp.ChatTabTeamsAppClient
	CloudPC                                                                  *cloudpc.CloudPCClient
	Contact                                                                  *contact.ContactClient
	ContactExtension                                                         *contactextension.ContactExtensionClient
	ContactFolder                                                            *contactfolder.ContactFolderClient
	ContactFolderChildFolder                                                 *contactfolderchildfolder.ContactFolderChildFolderClient
	ContactFolderChildFolderContact                                          *contactfolderchildfoldercontact.ContactFolderChildFolderContactClient
	ContactFolderChildFolderContactExtension                                 *contactfolderchildfoldercontactextension.ContactFolderChildFolderContactExtensionClient
	ContactFolderChildFolderContactPhoto                                     *contactfolderchildfoldercontactphoto.ContactFolderChildFolderContactPhotoClient
	ContactFolderContact                                                     *contactfoldercontact.ContactFolderContactClient
	ContactFolderContactExtension                                            *contactfoldercontactextension.ContactFolderContactExtensionClient
	ContactFolderContactPhoto                                                *contactfoldercontactphoto.ContactFolderContactPhotoClient
	ContactPhoto                                                             *contactphoto.ContactPhotoClient
	CreatedObject                                                            *createdobject.CreatedObjectClient
	Device                                                                   *device.DeviceClient
	DeviceCommand                                                            *devicecommand.DeviceCommandClient
	DeviceCommandResponsepayload                                             *devicecommandresponsepayload.DeviceCommandResponsepayloadClient
	DeviceEnrollmentConfiguration                                            *deviceenrollmentconfiguration.DeviceEnrollmentConfigurationClient
	DeviceEnrollmentConfigurationAssignment                                  *deviceenrollmentconfigurationassignment.DeviceEnrollmentConfigurationAssignmentClient
	DeviceExtension                                                          *deviceextension.DeviceExtensionClient
	DeviceManagementTroubleshootingEvent                                     *devicemanagementtroubleshootingevent.DeviceManagementTroubleshootingEventClient
	DeviceMemberOf                                                           *devicememberof.DeviceMemberOfClient
	DeviceRegisteredOwner                                                    *deviceregisteredowner.DeviceRegisteredOwnerClient
	DeviceRegisteredUser                                                     *deviceregistereduser.DeviceRegisteredUserClient
	DeviceTransitiveMemberOf                                                 *devicetransitivememberof.DeviceTransitiveMemberOfClient
	DeviceUsageRight                                                         *deviceusageright.DeviceUsageRightClient
	DirectReport                                                             *directreport.DirectReportClient
	Drive                                                                    *drive.DriveClient
	EmployeeExperience                                                       *employeeexperience.EmployeeExperienceClient
	EmployeeExperienceLearningCourseActivity                                 *employeeexperiencelearningcourseactivity.EmployeeExperienceLearningCourseActivityClient
	Event                                                                    *event.EventClient
	EventAttachment                                                          *eventattachment.EventAttachmentClient
	EventCalendar                                                            *eventcalendar.EventCalendarClient
	EventExceptionOccurrence                                                 *eventexceptionoccurrence.EventExceptionOccurrenceClient
	EventExceptionOccurrenceAttachment                                       *eventexceptionoccurrenceattachment.EventExceptionOccurrenceAttachmentClient
	EventExceptionOccurrenceCalendar                                         *eventexceptionoccurrencecalendar.EventExceptionOccurrenceCalendarClient
	EventExceptionOccurrenceExtension                                        *eventexceptionoccurrenceextension.EventExceptionOccurrenceExtensionClient
	EventExceptionOccurrenceInstance                                         *eventexceptionoccurrenceinstance.EventExceptionOccurrenceInstanceClient
	EventExceptionOccurrenceInstanceAttachment                               *eventexceptionoccurrenceinstanceattachment.EventExceptionOccurrenceInstanceAttachmentClient
	EventExceptionOccurrenceInstanceCalendar                                 *eventexceptionoccurrenceinstancecalendar.EventExceptionOccurrenceInstanceCalendarClient
	EventExceptionOccurrenceInstanceExtension                                *eventexceptionoccurrenceinstanceextension.EventExceptionOccurrenceInstanceExtensionClient
	EventExtension                                                           *eventextension.EventExtensionClient
	EventInstance                                                            *eventinstance.EventInstanceClient
	EventInstanceAttachment                                                  *eventinstanceattachment.EventInstanceAttachmentClient
	EventInstanceCalendar                                                    *eventinstancecalendar.EventInstanceCalendarClient
	EventInstanceExceptionOccurrence                                         *eventinstanceexceptionoccurrence.EventInstanceExceptionOccurrenceClient
	EventInstanceExceptionOccurrenceAttachment                               *eventinstanceexceptionoccurrenceattachment.EventInstanceExceptionOccurrenceAttachmentClient
	EventInstanceExceptionOccurrenceCalendar                                 *eventinstanceexceptionoccurrencecalendar.EventInstanceExceptionOccurrenceCalendarClient
	EventInstanceExceptionOccurrenceExtension                                *eventinstanceexceptionoccurrenceextension.EventInstanceExceptionOccurrenceExtensionClient
	EventInstanceExtension                                                   *eventinstanceextension.EventInstanceExtensionClient
	Extension                                                                *extension.ExtensionClient
	FollowedSite                                                             *followedsite.FollowedSiteClient
	InferenceClassification                                                  *inferenceclassification.InferenceClassificationClient
	InferenceClassificationOverride                                          *inferenceclassificationoverride.InferenceClassificationOverrideClient
	InformationProtection                                                    *informationprotection.InformationProtectionClient
	InformationProtectionBitlocker                                           *informationprotectionbitlocker.InformationProtectionBitlockerClient
	InformationProtectionBitlockerRecoveryKey                                *informationprotectionbitlockerrecoverykey.InformationProtectionBitlockerRecoveryKeyClient
	InformationProtectionDataLossPreventionPolicy                            *informationprotectiondatalosspreventionpolicy.InformationProtectionDataLossPreventionPolicyClient
	InformationProtectionPolicy                                              *informationprotectionpolicy.InformationProtectionPolicyClient
	InformationProtectionPolicyLabel                                         *informationprotectionpolicylabel.InformationProtectionPolicyLabelClient
	InformationProtectionSensitivityLabel                                    *informationprotectionsensitivitylabel.InformationProtectionSensitivityLabelClient
	InformationProtectionSensitivityLabelSublabel                            *informationprotectionsensitivitylabelsublabel.InformationProtectionSensitivityLabelSublabelClient
	InformationProtectionSensitivityPolicySetting                            *informationprotectionsensitivitypolicysetting.InformationProtectionSensitivityPolicySettingClient
	InformationProtectionThreatAssessmentRequest                             *informationprotectionthreatassessmentrequest.InformationProtectionThreatAssessmentRequestClient
	InformationProtectionThreatAssessmentRequestResult                       *informationprotectionthreatassessmentrequestresult.InformationProtectionThreatAssessmentRequestResultClient
	Insight                                                                  *insight.InsightClient
	InsightShared                                                            *insightshared.InsightSharedClient
	InsightSharedLastSharedMethod                                            *insightsharedlastsharedmethod.InsightSharedLastSharedMethodClient
	InsightSharedResource                                                    *insightsharedresource.InsightSharedResourceClient
	InsightTrending                                                          *insighttrending.InsightTrendingClient
	InsightTrendingResource                                                  *insighttrendingresource.InsightTrendingResourceClient
	InsightUsed                                                              *insightused.InsightUsedClient
	InsightUsedResource                                                      *insightusedresource.InsightUsedResourceClient
	JoinedGroup                                                              *joinedgroup.JoinedGroupClient
	JoinedTeam                                                               *joinedteam.JoinedTeamClient
	LicenseDetail                                                            *licensedetail.LicenseDetailClient
	MailFolder                                                               *mailfolder.MailFolderClient
	MailFolderChildFolder                                                    *mailfolderchildfolder.MailFolderChildFolderClient
	MailFolderChildFolderMessage                                             *mailfolderchildfoldermessage.MailFolderChildFolderMessageClient
	MailFolderChildFolderMessageAttachment                                   *mailfolderchildfoldermessageattachment.MailFolderChildFolderMessageAttachmentClient
	MailFolderChildFolderMessageExtension                                    *mailfolderchildfoldermessageextension.MailFolderChildFolderMessageExtensionClient
	MailFolderChildFolderMessageMention                                      *mailfolderchildfoldermessagemention.MailFolderChildFolderMessageMentionClient
	MailFolderChildFolderMessageRule                                         *mailfolderchildfoldermessagerule.MailFolderChildFolderMessageRuleClient
	MailFolderChildFolderUserConfiguration                                   *mailfolderchildfolderuserconfiguration.MailFolderChildFolderUserConfigurationClient
	MailFolderMessage                                                        *mailfoldermessage.MailFolderMessageClient
	MailFolderMessageAttachment                                              *mailfoldermessageattachment.MailFolderMessageAttachmentClient
	MailFolderMessageExtension                                               *mailfoldermessageextension.MailFolderMessageExtensionClient
	MailFolderMessageMention                                                 *mailfoldermessagemention.MailFolderMessageMentionClient
	MailFolderMessageRule                                                    *mailfoldermessagerule.MailFolderMessageRuleClient
	MailFolderUserConfiguration                                              *mailfolderuserconfiguration.MailFolderUserConfigurationClient
	MailboxSetting                                                           *mailboxsetting.MailboxSettingClient
	ManagedAppRegistration                                                   *managedappregistration.ManagedAppRegistrationClient
	ManagedDevice                                                            *manageddevice.ManagedDeviceClient
	ManagedDeviceAssignmentFilterEvaluationStatusDetail                      *manageddeviceassignmentfilterevaluationstatusdetail.ManagedDeviceAssignmentFilterEvaluationStatusDetailClient
	ManagedDeviceDetectedApp                                                 *manageddevicedetectedapp.ManagedDeviceDetectedAppClient
	ManagedDeviceDeviceCategory                                              *manageddevicedevicecategory.ManagedDeviceDeviceCategoryClient
	ManagedDeviceDeviceCompliancePolicyState                                 *manageddevicedevicecompliancepolicystate.ManagedDeviceDeviceCompliancePolicyStateClient
	ManagedDeviceDeviceConfigurationState                                    *manageddevicedeviceconfigurationstate.ManagedDeviceDeviceConfigurationStateClient
	ManagedDeviceDeviceHealthScriptState                                     *manageddevicedevicehealthscriptstate.ManagedDeviceDeviceHealthScriptStateClient
	ManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId *manageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid.ManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient
	ManagedDeviceLogCollectionRequest                                        *manageddevicelogcollectionrequest.ManagedDeviceLogCollectionRequestClient
	ManagedDeviceManagedDeviceMobileAppConfigurationState                    *manageddevicemanageddevicemobileappconfigurationstate.ManagedDeviceManagedDeviceMobileAppConfigurationStateClient
	ManagedDeviceSecurityBaselineState                                       *manageddevicesecuritybaselinestate.ManagedDeviceSecurityBaselineStateClient
	ManagedDeviceSecurityBaselineStateSettingState                           *manageddevicesecuritybaselinestatesettingstate.ManagedDeviceSecurityBaselineStateSettingStateClient
	ManagedDeviceUser                                                        *manageddeviceuser.ManagedDeviceUserClient
	ManagedDeviceWindowsProtectionState                                      *manageddevicewindowsprotectionstate.ManagedDeviceWindowsProtectionStateClient
	ManagedDeviceWindowsProtectionStateDetectedMalwareState                  *manageddevicewindowsprotectionstatedetectedmalwarestate.ManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	Manager                                                                  *manager.ManagerClient
	MemberOf                                                                 *memberof.MemberOfClient
	Message                                                                  *message.MessageClient
	MessageAttachment                                                        *messageattachment.MessageAttachmentClient
	MessageExtension                                                         *messageextension.MessageExtensionClient
	MessageMention                                                           *messagemention.MessageMentionClient
	MobileAppIntentAndState                                                  *mobileappintentandstate.MobileAppIntentAndStateClient
	MobileAppTroubleshootingEvent                                            *mobileapptroubleshootingevent.MobileAppTroubleshootingEventClient
	MobileAppTroubleshootingEventAppLogCollectionRequest                     *mobileapptroubleshootingeventapplogcollectionrequest.MobileAppTroubleshootingEventAppLogCollectionRequestClient
	Notification                                                             *notification.NotificationClient
	Oauth2PermissionGrant                                                    *oauth2permissiongrant.Oauth2PermissionGrantClient
	OnlineMeeting                                                            *onlinemeeting.OnlineMeetingClient
	OnlineMeetingAlternativeRecording                                        *onlinemeetingalternativerecording.OnlineMeetingAlternativeRecordingClient
	OnlineMeetingAttendanceReport                                            *onlinemeetingattendancereport.OnlineMeetingAttendanceReportClient
	OnlineMeetingAttendanceReportAttendanceRecord                            *onlinemeetingattendancereportattendancerecord.OnlineMeetingAttendanceReportAttendanceRecordClient
	OnlineMeetingAttendeeReport                                              *onlinemeetingattendeereport.OnlineMeetingAttendeeReportClient
	OnlineMeetingBroadcastRecording                                          *onlinemeetingbroadcastrecording.OnlineMeetingBroadcastRecordingClient
	OnlineMeetingMeetingAttendanceReport                                     *onlinemeetingmeetingattendancereport.OnlineMeetingMeetingAttendanceReportClient
	OnlineMeetingMeetingAttendanceReportAttendanceRecord                     *onlinemeetingmeetingattendancereportattendancerecord.OnlineMeetingMeetingAttendanceReportAttendanceRecordClient
	OnlineMeetingRecording                                                   *onlinemeetingrecording.OnlineMeetingRecordingClient
	OnlineMeetingRecordingContent                                            *onlinemeetingrecordingcontent.OnlineMeetingRecordingContentClient
	OnlineMeetingRegistration                                                *onlinemeetingregistration.OnlineMeetingRegistrationClient
	OnlineMeetingRegistrationCustomQuestion                                  *onlinemeetingregistrationcustomquestion.OnlineMeetingRegistrationCustomQuestionClient
	OnlineMeetingRegistrationRegistrant                                      *onlinemeetingregistrationregistrant.OnlineMeetingRegistrationRegistrantClient
	OnlineMeetingTranscript                                                  *onlinemeetingtranscript.OnlineMeetingTranscriptClient
	OnlineMeetingTranscriptContent                                           *onlinemeetingtranscriptcontent.OnlineMeetingTranscriptContentClient
	OnlineMeetingTranscriptMetadataContent                                   *onlinemeetingtranscriptmetadatacontent.OnlineMeetingTranscriptMetadataContentClient
	Outlook                                                                  *outlook.OutlookClient
	OutlookMasterCategory                                                    *outlookmastercategory.OutlookMasterCategoryClient
	OutlookTask                                                              *outlooktask.OutlookTaskClient
	OutlookTaskAttachment                                                    *outlooktaskattachment.OutlookTaskAttachmentClient
	OutlookTaskFolder                                                        *outlooktaskfolder.OutlookTaskFolderClient
	OutlookTaskFolderTask                                                    *outlooktaskfoldertask.OutlookTaskFolderTaskClient
	OutlookTaskFolderTaskAttachment                                          *outlooktaskfoldertaskattachment.OutlookTaskFolderTaskAttachmentClient
	OutlookTaskGroup                                                         *outlooktaskgroup.OutlookTaskGroupClient
	OutlookTaskGroupTaskFolder                                               *outlooktaskgrouptaskfolder.OutlookTaskGroupTaskFolderClient
	OutlookTaskGroupTaskFolderTask                                           *outlooktaskgrouptaskfoldertask.OutlookTaskGroupTaskFolderTaskClient
	OutlookTaskGroupTaskFolderTaskAttachment                                 *outlooktaskgrouptaskfoldertaskattachment.OutlookTaskGroupTaskFolderTaskAttachmentClient
	OwnedDevice                                                              *owneddevice.OwnedDeviceClient
	OwnedObject                                                              *ownedobject.OwnedObjectClient
	PendingAccessReviewInstance                                              *pendingaccessreviewinstance.PendingAccessReviewInstanceClient
	PendingAccessReviewInstanceContactedReviewer                             *pendingaccessreviewinstancecontactedreviewer.PendingAccessReviewInstanceContactedReviewerClient
	PendingAccessReviewInstanceDecision                                      *pendingaccessreviewinstancedecision.PendingAccessReviewInstanceDecisionClient
	PendingAccessReviewInstanceDecisionInsight                               *pendingaccessreviewinstancedecisioninsight.PendingAccessReviewInstanceDecisionInsightClient
	PendingAccessReviewInstanceDecisionInstance                              *pendingaccessreviewinstancedecisioninstance.PendingAccessReviewInstanceDecisionInstanceClient
	PendingAccessReviewInstanceDecisionInstanceContactedReviewer             *pendingaccessreviewinstancedecisioninstancecontactedreviewer.PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient
	PendingAccessReviewInstanceDecisionInstanceDefinition                    *pendingaccessreviewinstancedecisioninstancedefinition.PendingAccessReviewInstanceDecisionInstanceDefinitionClient
	PendingAccessReviewInstanceDecisionInstanceStage                         *pendingaccessreviewinstancedecisioninstancestage.PendingAccessReviewInstanceDecisionInstanceStageClient
	PendingAccessReviewInstanceDefinition                                    *pendingaccessreviewinstancedefinition.PendingAccessReviewInstanceDefinitionClient
	PendingAccessReviewInstanceStage                                         *pendingaccessreviewinstancestage.PendingAccessReviewInstanceStageClient
	PendingAccessReviewInstanceStageDecision                                 *pendingaccessreviewinstancestagedecision.PendingAccessReviewInstanceStageDecisionClient
	PendingAccessReviewInstanceStageDecisionInsight                          *pendingaccessreviewinstancestagedecisioninsight.PendingAccessReviewInstanceStageDecisionInsightClient
	PendingAccessReviewInstanceStageDecisionInstance                         *pendingaccessreviewinstancestagedecisioninstance.PendingAccessReviewInstanceStageDecisionInstanceClient
	PendingAccessReviewInstanceStageDecisionInstanceContactedReviewer        *pendingaccessreviewinstancestagedecisioninstancecontactedreviewer.PendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient
	PendingAccessReviewInstanceStageDecisionInstanceDefinition               *pendingaccessreviewinstancestagedecisioninstancedefinition.PendingAccessReviewInstanceStageDecisionInstanceDefinitionClient
	People                                                                   *people.PeopleClient
	PermissionGrant                                                          *permissiongrant.PermissionGrantClient
	Photo                                                                    *photo.PhotoClient
	Planner                                                                  *planner.PlannerClient
	PlannerAll                                                               *plannerall.PlannerAllClient
	PlannerFavoritePlan                                                      *plannerfavoriteplan.PlannerFavoritePlanClient
	PlannerPlan                                                              *plannerplan.PlannerPlanClient
	PlannerPlanBucket                                                        *plannerplanbucket.PlannerPlanBucketClient
	PlannerPlanBucketTask                                                    *plannerplanbuckettask.PlannerPlanBucketTaskClient
	PlannerPlanBucketTaskAssignedToTaskBoardFormat                           *plannerplanbuckettaskassignedtotaskboardformat.PlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	PlannerPlanBucketTaskBucketTaskBoardFormat                               *plannerplanbuckettaskbuckettaskboardformat.PlannerPlanBucketTaskBucketTaskBoardFormatClient
	PlannerPlanBucketTaskDetail                                              *plannerplanbuckettaskdetail.PlannerPlanBucketTaskDetailClient
	PlannerPlanBucketTaskProgressTaskBoardFormat                             *plannerplanbuckettaskprogresstaskboardformat.PlannerPlanBucketTaskProgressTaskBoardFormatClient
	PlannerPlanDetail                                                        *plannerplandetail.PlannerPlanDetailClient
	PlannerPlanTask                                                          *plannerplantask.PlannerPlanTaskClient
	PlannerPlanTaskAssignedToTaskBoardFormat                                 *plannerplantaskassignedtotaskboardformat.PlannerPlanTaskAssignedToTaskBoardFormatClient
	PlannerPlanTaskBucketTaskBoardFormat                                     *plannerplantaskbuckettaskboardformat.PlannerPlanTaskBucketTaskBoardFormatClient
	PlannerPlanTaskDetail                                                    *plannerplantaskdetail.PlannerPlanTaskDetailClient
	PlannerPlanTaskProgressTaskBoardFormat                                   *plannerplantaskprogresstaskboardformat.PlannerPlanTaskProgressTaskBoardFormatClient
	PlannerRecentPlan                                                        *plannerrecentplan.PlannerRecentPlanClient
	PlannerRosterPlan                                                        *plannerrosterplan.PlannerRosterPlanClient
	PlannerTask                                                              *plannertask.PlannerTaskClient
	PlannerTaskAssignedToTaskBoardFormat                                     *plannertaskassignedtotaskboardformat.PlannerTaskAssignedToTaskBoardFormatClient
	PlannerTaskBucketTaskBoardFormat                                         *plannertaskbuckettaskboardformat.PlannerTaskBucketTaskBoardFormatClient
	PlannerTaskDetail                                                        *plannertaskdetail.PlannerTaskDetailClient
	PlannerTaskProgressTaskBoardFormat                                       *plannertaskprogresstaskboardformat.PlannerTaskProgressTaskBoardFormatClient
	Presence                                                                 *presence.PresenceClient
	Profile                                                                  *profile.ProfileClient
	ProfileAccount                                                           *profileaccount.ProfileAccountClient
	ProfileAddress                                                           *profileaddress.ProfileAddressClient
	ProfileAnniversary                                                       *profileanniversary.ProfileAnniversaryClient
	ProfileAward                                                             *profileaward.ProfileAwardClient
	ProfileCertification                                                     *profilecertification.ProfileCertificationClient
	ProfileEducationalActivity                                               *profileeducationalactivity.ProfileEducationalActivityClient
	ProfileEmail                                                             *profileemail.ProfileEmailClient
	ProfileInterest                                                          *profileinterest.ProfileInterestClient
	ProfileLanguage                                                          *profilelanguage.ProfileLanguageClient
	ProfileName                                                              *profilename.ProfileNameClient
	ProfileNote                                                              *profilenote.ProfileNoteClient
	ProfilePatent                                                            *profilepatent.ProfilePatentClient
	ProfilePhone                                                             *profilephone.ProfilePhoneClient
	ProfilePosition                                                          *profileposition.ProfilePositionClient
	ProfileProject                                                           *profileproject.ProfileProjectClient
	ProfilePublication                                                       *profilepublication.ProfilePublicationClient
	ProfileSkill                                                             *profileskill.ProfileSkillClient
	ProfileWebAccount                                                        *profilewebaccount.ProfileWebAccountClient
	ProfileWebsite                                                           *profilewebsite.ProfileWebsiteClient
	RegisteredDevice                                                         *registereddevice.RegisteredDeviceClient
	ScopedRoleMemberOf                                                       *scopedrolememberof.ScopedRoleMemberOfClient
	Security                                                                 *security.SecurityClient
	SecurityInformationProtection                                            *securityinformationprotection.SecurityInformationProtectionClient
	SecurityInformationProtectionLabelPolicySetting                          *securityinformationprotectionlabelpolicysetting.SecurityInformationProtectionLabelPolicySettingClient
	SecurityInformationProtectionSensitivityLabel                            *securityinformationprotectionsensitivitylabel.SecurityInformationProtectionSensitivityLabelClient
	SecurityInformationProtectionSensitivityLabelParent                      *securityinformationprotectionsensitivitylabelparent.SecurityInformationProtectionSensitivityLabelParentClient
	ServiceProvisioningError                                                 *serviceprovisioningerror.ServiceProvisioningErrorClient
	Setting                                                                  *setting.SettingClient
	SettingContactMergeSuggestion                                            *settingcontactmergesuggestion.SettingContactMergeSuggestionClient
	SettingItemInsight                                                       *settingiteminsight.SettingItemInsightClient
	SettingRegionalAndLanguageSetting                                        *settingregionalandlanguagesetting.SettingRegionalAndLanguageSettingClient
	SettingShiftPreference                                                   *settingshiftpreference.SettingShiftPreferenceClient
	Sponsor                                                                  *sponsor.SponsorClient
	Teamwork                                                                 *teamwork.TeamworkClient
	TeamworkAssociatedTeam                                                   *teamworkassociatedteam.TeamworkAssociatedTeamClient
	TeamworkAssociatedTeamTeam                                               *teamworkassociatedteamteam.TeamworkAssociatedTeamTeamClient
	TeamworkInstalledApp                                                     *teamworkinstalledapp.TeamworkInstalledAppClient
	TeamworkInstalledAppChat                                                 *teamworkinstalledappchat.TeamworkInstalledAppChatClient
	TeamworkInstalledAppTeamsApp                                             *teamworkinstalledappteamsapp.TeamworkInstalledAppTeamsAppClient
	TeamworkInstalledAppTeamsAppDefinition                                   *teamworkinstalledappteamsappdefinition.TeamworkInstalledAppTeamsAppDefinitionClient
	Todo                                                                     *todo.TodoClient
	TodoList                                                                 *todolist.TodoListClient
	TodoListExtension                                                        *todolistextension.TodoListExtensionClient
	TodoListTask                                                             *todolisttask.TodoListTaskClient
	TodoListTaskAttachment                                                   *todolisttaskattachment.TodoListTaskAttachmentClient
	TodoListTaskAttachmentSession                                            *todolisttaskattachmentsession.TodoListTaskAttachmentSessionClient
	TodoListTaskAttachmentSessionContent                                     *todolisttaskattachmentsessioncontent.TodoListTaskAttachmentSessionContentClient
	TodoListTaskChecklistItem                                                *todolisttaskchecklistitem.TodoListTaskChecklistItemClient
	TodoListTaskExtension                                                    *todolisttaskextension.TodoListTaskExtensionClient
	TodoListTaskLinkedResource                                               *todolisttasklinkedresource.TodoListTaskLinkedResourceClient
	TransitiveMemberOf                                                       *transitivememberof.TransitiveMemberOfClient
	TransitiveReport                                                         *transitivereport.TransitiveReportClient
	UsageRight                                                               *usageright.UsageRightClient
	User                                                                     *user.UserClient
	VirtualEvent                                                             *virtualevent.VirtualEventClient
	VirtualEventWebinar                                                      *virtualeventwebinar.VirtualEventWebinarClient
	WindowsInformationProtectionDeviceRegistration                           *windowsinformationprotectiondeviceregistration.WindowsInformationProtectionDeviceRegistrationClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	activityClient, err := activity.NewActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Activity client: %+v", err)
	}
	configureFunc(activityClient.Client)

	activityHistoryItemActivityClient, err := activityhistoryitemactivity.NewActivityHistoryItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActivityHistoryItemActivity client: %+v", err)
	}
	configureFunc(activityHistoryItemActivityClient.Client)

	activityHistoryItemClient, err := activityhistoryitem.NewActivityHistoryItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ActivityHistoryItem client: %+v", err)
	}
	configureFunc(activityHistoryItemClient.Client)

	agreementAcceptanceClient, err := agreementacceptance.NewAgreementAcceptanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AgreementAcceptance client: %+v", err)
	}
	configureFunc(agreementAcceptanceClient.Client)

	analyticActivityStatisticClient, err := analyticactivitystatistic.NewAnalyticActivityStatisticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AnalyticActivityStatistic client: %+v", err)
	}
	configureFunc(analyticActivityStatisticClient.Client)

	analyticClient, err := analytic.NewAnalyticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Analytic client: %+v", err)
	}
	configureFunc(analyticClient.Client)

	appConsentRequestsForApprovalClient, err := appconsentrequestsforapproval.NewAppConsentRequestsForApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentRequestsForApproval client: %+v", err)
	}
	configureFunc(appConsentRequestsForApprovalClient.Client)

	appConsentRequestsForApprovalUserConsentRequestApprovalClient, err := appconsentrequestsforapprovaluserconsentrequestapproval.NewAppConsentRequestsForApprovalUserConsentRequestApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentRequestsForApprovalUserConsentRequestApproval client: %+v", err)
	}
	configureFunc(appConsentRequestsForApprovalUserConsentRequestApprovalClient.Client)

	appConsentRequestsForApprovalUserConsentRequestApprovalStepClient, err := appconsentrequestsforapprovaluserconsentrequestapprovalstep.NewAppConsentRequestsForApprovalUserConsentRequestApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentRequestsForApprovalUserConsentRequestApprovalStep client: %+v", err)
	}
	configureFunc(appConsentRequestsForApprovalUserConsentRequestApprovalStepClient.Client)

	appConsentRequestsForApprovalUserConsentRequestClient, err := appconsentrequestsforapprovaluserconsentrequest.NewAppConsentRequestsForApprovalUserConsentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentRequestsForApprovalUserConsentRequest client: %+v", err)
	}
	configureFunc(appConsentRequestsForApprovalUserConsentRequestClient.Client)

	appRoleAssignedResourceClient, err := approleassignedresource.NewAppRoleAssignedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppRoleAssignedResource client: %+v", err)
	}
	configureFunc(appRoleAssignedResourceClient.Client)

	appRoleAssignmentClient, err := approleassignment.NewAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppRoleAssignment client: %+v", err)
	}
	configureFunc(appRoleAssignmentClient.Client)

	approvalClient, err := approval.NewApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Approval client: %+v", err)
	}
	configureFunc(approvalClient.Client)

	approvalStepClient, err := approvalstep.NewApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApprovalStep client: %+v", err)
	}
	configureFunc(approvalStepClient.Client)

	authenticationClient, err := authentication.NewAuthenticationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Authentication client: %+v", err)
	}
	configureFunc(authenticationClient.Client)

	authenticationEmailMethodClient, err := authenticationemailmethod.NewAuthenticationEmailMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationEmailMethod client: %+v", err)
	}
	configureFunc(authenticationEmailMethodClient.Client)

	authenticationFido2MethodClient, err := authenticationfido2method.NewAuthenticationFido2MethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationFido2Method client: %+v", err)
	}
	configureFunc(authenticationFido2MethodClient.Client)

	authenticationMethodClient, err := authenticationmethod.NewAuthenticationMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMethod client: %+v", err)
	}
	configureFunc(authenticationMethodClient.Client)

	authenticationMicrosoftAuthenticatorMethodClient, err := authenticationmicrosoftauthenticatormethod.NewAuthenticationMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMicrosoftAuthenticatorMethod client: %+v", err)
	}
	configureFunc(authenticationMicrosoftAuthenticatorMethodClient.Client)

	authenticationMicrosoftAuthenticatorMethodDeviceClient, err := authenticationmicrosoftauthenticatormethoddevice.NewAuthenticationMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationMicrosoftAuthenticatorMethodDevice client: %+v", err)
	}
	configureFunc(authenticationMicrosoftAuthenticatorMethodDeviceClient.Client)

	authenticationOperationClient, err := authenticationoperation.NewAuthenticationOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationOperation client: %+v", err)
	}
	configureFunc(authenticationOperationClient.Client)

	authenticationPasswordMethodClient, err := authenticationpasswordmethod.NewAuthenticationPasswordMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationPasswordMethod client: %+v", err)
	}
	configureFunc(authenticationPasswordMethodClient.Client)

	authenticationPasswordlessMicrosoftAuthenticatorMethodClient, err := authenticationpasswordlessmicrosoftauthenticatormethod.NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationPasswordlessMicrosoftAuthenticatorMethod client: %+v", err)
	}
	configureFunc(authenticationPasswordlessMicrosoftAuthenticatorMethodClient.Client)

	authenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient, err := authenticationpasswordlessmicrosoftauthenticatormethoddevice.NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice client: %+v", err)
	}
	configureFunc(authenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient.Client)

	authenticationPhoneMethodClient, err := authenticationphonemethod.NewAuthenticationPhoneMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationPhoneMethod client: %+v", err)
	}
	configureFunc(authenticationPhoneMethodClient.Client)

	authenticationPlatformCredentialMethodClient, err := authenticationplatformcredentialmethod.NewAuthenticationPlatformCredentialMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationPlatformCredentialMethod client: %+v", err)
	}
	configureFunc(authenticationPlatformCredentialMethodClient.Client)

	authenticationPlatformCredentialMethodDeviceClient, err := authenticationplatformcredentialmethoddevice.NewAuthenticationPlatformCredentialMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationPlatformCredentialMethodDevice client: %+v", err)
	}
	configureFunc(authenticationPlatformCredentialMethodDeviceClient.Client)

	authenticationSignInPreferenceClient, err := authenticationsigninpreference.NewAuthenticationSignInPreferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationSignInPreference client: %+v", err)
	}
	configureFunc(authenticationSignInPreferenceClient.Client)

	authenticationSoftwareOathMethodClient, err := authenticationsoftwareoathmethod.NewAuthenticationSoftwareOathMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationSoftwareOathMethod client: %+v", err)
	}
	configureFunc(authenticationSoftwareOathMethodClient.Client)

	authenticationTemporaryAccessPassMethodClient, err := authenticationtemporaryaccesspassmethod.NewAuthenticationTemporaryAccessPassMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationTemporaryAccessPassMethod client: %+v", err)
	}
	configureFunc(authenticationTemporaryAccessPassMethodClient.Client)

	authenticationWindowsHelloForBusinessMethodClient, err := authenticationwindowshelloforbusinessmethod.NewAuthenticationWindowsHelloForBusinessMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationWindowsHelloForBusinessMethod client: %+v", err)
	}
	configureFunc(authenticationWindowsHelloForBusinessMethodClient.Client)

	authenticationWindowsHelloForBusinessMethodDeviceClient, err := authenticationwindowshelloforbusinessmethoddevice.NewAuthenticationWindowsHelloForBusinessMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationWindowsHelloForBusinessMethodDevice client: %+v", err)
	}
	configureFunc(authenticationWindowsHelloForBusinessMethodDeviceClient.Client)

	calendarCalendarPermissionClient, err := calendarcalendarpermission.NewCalendarCalendarPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarPermission client: %+v", err)
	}
	configureFunc(calendarCalendarPermissionClient.Client)

	calendarCalendarViewAttachmentClient, err := calendarcalendarviewattachment.NewCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewAttachment client: %+v", err)
	}
	configureFunc(calendarCalendarViewAttachmentClient.Client)

	calendarCalendarViewCalendarClient, err := calendarcalendarviewcalendar.NewCalendarCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewCalendar client: %+v", err)
	}
	configureFunc(calendarCalendarViewCalendarClient.Client)

	calendarCalendarViewClient, err := calendarcalendarview.NewCalendarCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarView client: %+v", err)
	}
	configureFunc(calendarCalendarViewClient.Client)

	calendarCalendarViewExceptionOccurrenceAttachmentClient, err := calendarcalendarviewexceptionoccurrenceattachment.NewCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceAttachmentClient.Client)

	calendarCalendarViewExceptionOccurrenceCalendarClient, err := calendarcalendarviewexceptionoccurrencecalendar.NewCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceCalendarClient.Client)

	calendarCalendarViewExceptionOccurrenceClient, err := calendarcalendarviewexceptionoccurrence.NewCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceClient.Client)

	calendarCalendarViewExceptionOccurrenceExtensionClient, err := calendarcalendarviewexceptionoccurrenceextension.NewCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceExtensionClient.Client)

	calendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := calendarcalendarviewexceptionoccurrenceinstanceattachment.NewCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	calendarCalendarViewExceptionOccurrenceInstanceCalendarClient, err := calendarcalendarviewexceptionoccurrenceinstancecalendar.NewCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	calendarCalendarViewExceptionOccurrenceInstanceClient, err := calendarcalendarviewexceptionoccurrenceinstance.NewCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceInstanceClient.Client)

	calendarCalendarViewExceptionOccurrenceInstanceExtensionClient, err := calendarcalendarviewexceptionoccurrenceinstanceextension.NewCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(calendarCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

	calendarCalendarViewExtensionClient, err := calendarcalendarviewextension.NewCalendarCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewExtension client: %+v", err)
	}
	configureFunc(calendarCalendarViewExtensionClient.Client)

	calendarCalendarViewInstanceAttachmentClient, err := calendarcalendarviewinstanceattachment.NewCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceAttachmentClient.Client)

	calendarCalendarViewInstanceCalendarClient, err := calendarcalendarviewinstancecalendar.NewCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceCalendarClient.Client)

	calendarCalendarViewInstanceClient, err := calendarcalendarviewinstance.NewCalendarCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstance client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceClient.Client)

	calendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := calendarcalendarviewinstanceexceptionoccurrenceattachment.NewCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	calendarCalendarViewInstanceExceptionOccurrenceCalendarClient, err := calendarcalendarviewinstanceexceptionoccurrencecalendar.NewCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	calendarCalendarViewInstanceExceptionOccurrenceClient, err := calendarcalendarviewinstanceexceptionoccurrence.NewCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceExceptionOccurrenceClient.Client)

	calendarCalendarViewInstanceExceptionOccurrenceExtensionClient, err := calendarcalendarviewinstanceexceptionoccurrenceextension.NewCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

	calendarCalendarViewInstanceExtensionClient, err := calendarcalendarviewinstanceextension.NewCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(calendarCalendarViewInstanceExtensionClient.Client)

	calendarClient, err := calendar.NewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Calendar client: %+v", err)
	}
	configureFunc(calendarClient.Client)

	calendarEventAttachmentClient, err := calendareventattachment.NewCalendarEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventAttachment client: %+v", err)
	}
	configureFunc(calendarEventAttachmentClient.Client)

	calendarEventCalendarClient, err := calendareventcalendar.NewCalendarEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventCalendar client: %+v", err)
	}
	configureFunc(calendarEventCalendarClient.Client)

	calendarEventClient, err := calendarevent.NewCalendarEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEvent client: %+v", err)
	}
	configureFunc(calendarEventClient.Client)

	calendarEventExceptionOccurrenceAttachmentClient, err := calendareventexceptionoccurrenceattachment.NewCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceAttachmentClient.Client)

	calendarEventExceptionOccurrenceCalendarClient, err := calendareventexceptionoccurrencecalendar.NewCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceCalendarClient.Client)

	calendarEventExceptionOccurrenceClient, err := calendareventexceptionoccurrence.NewCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceClient.Client)

	calendarEventExceptionOccurrenceExtensionClient, err := calendareventexceptionoccurrenceextension.NewCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceExtensionClient.Client)

	calendarEventExceptionOccurrenceInstanceAttachmentClient, err := calendareventexceptionoccurrenceinstanceattachment.NewCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceInstanceAttachmentClient.Client)

	calendarEventExceptionOccurrenceInstanceCalendarClient, err := calendareventexceptionoccurrenceinstancecalendar.NewCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceInstanceCalendarClient.Client)

	calendarEventExceptionOccurrenceInstanceClient, err := calendareventexceptionoccurrenceinstance.NewCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceInstanceClient.Client)

	calendarEventExceptionOccurrenceInstanceExtensionClient, err := calendareventexceptionoccurrenceinstanceextension.NewCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(calendarEventExceptionOccurrenceInstanceExtensionClient.Client)

	calendarEventExtensionClient, err := calendareventextension.NewCalendarEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventExtension client: %+v", err)
	}
	configureFunc(calendarEventExtensionClient.Client)

	calendarEventInstanceAttachmentClient, err := calendareventinstanceattachment.NewCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarEventInstanceAttachmentClient.Client)

	calendarEventInstanceCalendarClient, err := calendareventinstancecalendar.NewCalendarEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarEventInstanceCalendarClient.Client)

	calendarEventInstanceClient, err := calendareventinstance.NewCalendarEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstance client: %+v", err)
	}
	configureFunc(calendarEventInstanceClient.Client)

	calendarEventInstanceExceptionOccurrenceAttachmentClient, err := calendareventinstanceexceptionoccurrenceattachment.NewCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarEventInstanceExceptionOccurrenceAttachmentClient.Client)

	calendarEventInstanceExceptionOccurrenceCalendarClient, err := calendareventinstanceexceptionoccurrencecalendar.NewCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarEventInstanceExceptionOccurrenceCalendarClient.Client)

	calendarEventInstanceExceptionOccurrenceClient, err := calendareventinstanceexceptionoccurrence.NewCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarEventInstanceExceptionOccurrenceClient.Client)

	calendarEventInstanceExceptionOccurrenceExtensionClient, err := calendareventinstanceexceptionoccurrenceextension.NewCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarEventInstanceExceptionOccurrenceExtensionClient.Client)

	calendarEventInstanceExtensionClient, err := calendareventinstanceextension.NewCalendarEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarEventInstanceExtension client: %+v", err)
	}
	configureFunc(calendarEventInstanceExtensionClient.Client)

	calendarGroupCalendarCalendarPermissionClient, err := calendargroupcalendarcalendarpermission.NewCalendarGroupCalendarCalendarPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarPermission client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarPermissionClient.Client)

	calendarGroupCalendarCalendarViewAttachmentClient, err := calendargroupcalendarcalendarviewattachment.NewCalendarGroupCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewAttachmentClient.Client)

	calendarGroupCalendarCalendarViewCalendarClient, err := calendargroupcalendarcalendarviewcalendar.NewCalendarGroupCalendarCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewCalendarClient.Client)

	calendarGroupCalendarCalendarViewClient, err := calendargroupcalendarcalendarview.NewCalendarGroupCalendarCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarView client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient, err := calendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient, err := calendargroupcalendarcalendarviewexceptionoccurrencecalendar.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceClient, err := calendargroupcalendarcalendarviewexceptionoccurrence.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient, err := calendargroupcalendarcalendarviewexceptionoccurrenceextension.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, err := calendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient, err := calendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.Client)

	calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient, err := calendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension.NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

	calendarGroupCalendarCalendarViewExtensionClient, err := calendargroupcalendarcalendarviewextension.NewCalendarGroupCalendarCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewExtensionClient.Client)

	calendarGroupCalendarCalendarViewInstanceAttachmentClient, err := calendargroupcalendarcalendarviewinstanceattachment.NewCalendarGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceAttachmentClient.Client)

	calendarGroupCalendarCalendarViewInstanceCalendarClient, err := calendargroupcalendarcalendarviewinstancecalendar.NewCalendarGroupCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceCalendarClient.Client)

	calendarGroupCalendarCalendarViewInstanceClient, err := calendargroupcalendarcalendarviewinstance.NewCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstance client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceClient.Client)

	calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, err := calendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar.NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient, err := calendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.Client)

	calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, err := calendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension.NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

	calendarGroupCalendarCalendarViewInstanceExtensionClient, err := calendargroupcalendarcalendarviewinstanceextension.NewCalendarGroupCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarCalendarViewInstanceExtensionClient.Client)

	calendarGroupCalendarClient, err := calendargroupcalendar.NewCalendarGroupCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarClient.Client)

	calendarGroupCalendarEventAttachmentClient, err := calendargroupcalendareventattachment.NewCalendarGroupCalendarEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventAttachmentClient.Client)

	calendarGroupCalendarEventCalendarClient, err := calendargroupcalendareventcalendar.NewCalendarGroupCalendarEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventCalendarClient.Client)

	calendarGroupCalendarEventClient, err := calendargroupcalendarevent.NewCalendarGroupCalendarEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEvent client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceAttachmentClient, err := calendargroupcalendareventexceptionoccurrenceattachment.NewCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceAttachmentClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceCalendarClient, err := calendargroupcalendareventexceptionoccurrencecalendar.NewCalendarGroupCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceCalendarClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceClient, err := calendargroupcalendareventexceptionoccurrence.NewCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceExtensionClient, err := calendargroupcalendareventexceptionoccurrenceextension.NewCalendarGroupCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceExtensionClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient, err := calendargroupcalendareventexceptionoccurrenceinstanceattachment.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient, err := calendargroupcalendareventexceptionoccurrenceinstancecalendar.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceInstanceClient, err := calendargroupcalendareventexceptionoccurrenceinstance.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceInstanceClient.Client)

	calendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient, err := calendargroupcalendareventexceptionoccurrenceinstanceextension.NewCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient.Client)

	calendarGroupCalendarEventExtensionClient, err := calendargroupcalendareventextension.NewCalendarGroupCalendarEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventExtensionClient.Client)

	calendarGroupCalendarEventInstanceAttachmentClient, err := calendargroupcalendareventinstanceattachment.NewCalendarGroupCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceAttachmentClient.Client)

	calendarGroupCalendarEventInstanceCalendarClient, err := calendargroupcalendareventinstancecalendar.NewCalendarGroupCalendarEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceCalendarClient.Client)

	calendarGroupCalendarEventInstanceClient, err := calendargroupcalendareventinstance.NewCalendarGroupCalendarEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstance client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceClient.Client)

	calendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient, err := calendargroupcalendareventinstanceexceptionoccurrenceattachment.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.Client)

	calendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient, err := calendargroupcalendareventinstanceexceptionoccurrencecalendar.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient.Client)

	calendarGroupCalendarEventInstanceExceptionOccurrenceClient, err := calendargroupcalendareventinstanceexceptionoccurrence.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceExceptionOccurrenceClient.Client)

	calendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient, err := calendargroupcalendareventinstanceexceptionoccurrenceextension.NewCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient.Client)

	calendarGroupCalendarEventInstanceExtensionClient, err := calendargroupcalendareventinstanceextension.NewCalendarGroupCalendarEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroupCalendarEventInstanceExtension client: %+v", err)
	}
	configureFunc(calendarGroupCalendarEventInstanceExtensionClient.Client)

	calendarGroupClient, err := calendargroup.NewCalendarGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarGroup client: %+v", err)
	}
	configureFunc(calendarGroupClient.Client)

	calendarViewAttachmentClient, err := calendarviewattachment.NewCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewAttachment client: %+v", err)
	}
	configureFunc(calendarViewAttachmentClient.Client)

	calendarViewCalendarClient, err := calendarviewcalendar.NewCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewCalendar client: %+v", err)
	}
	configureFunc(calendarViewCalendarClient.Client)

	calendarViewClient, err := calendarview.NewCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarView client: %+v", err)
	}
	configureFunc(calendarViewClient.Client)

	calendarViewExceptionOccurrenceAttachmentClient, err := calendarviewexceptionoccurrenceattachment.NewCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceAttachmentClient.Client)

	calendarViewExceptionOccurrenceCalendarClient, err := calendarviewexceptionoccurrencecalendar.NewCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceCalendarClient.Client)

	calendarViewExceptionOccurrenceClient, err := calendarviewexceptionoccurrence.NewCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceClient.Client)

	calendarViewExceptionOccurrenceExtensionClient, err := calendarviewexceptionoccurrenceextension.NewCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceExtensionClient.Client)

	calendarViewExceptionOccurrenceInstanceAttachmentClient, err := calendarviewexceptionoccurrenceinstanceattachment.NewCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	calendarViewExceptionOccurrenceInstanceCalendarClient, err := calendarviewexceptionoccurrenceinstancecalendar.NewCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	calendarViewExceptionOccurrenceInstanceClient, err := calendarviewexceptionoccurrenceinstance.NewCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceInstanceClient.Client)

	calendarViewExceptionOccurrenceInstanceExtensionClient, err := calendarviewexceptionoccurrenceinstanceextension.NewCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(calendarViewExceptionOccurrenceInstanceExtensionClient.Client)

	calendarViewExtensionClient, err := calendarviewextension.NewCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewExtension client: %+v", err)
	}
	configureFunc(calendarViewExtensionClient.Client)

	calendarViewInstanceAttachmentClient, err := calendarviewinstanceattachment.NewCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(calendarViewInstanceAttachmentClient.Client)

	calendarViewInstanceCalendarClient, err := calendarviewinstancecalendar.NewCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(calendarViewInstanceCalendarClient.Client)

	calendarViewInstanceClient, err := calendarviewinstance.NewCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstance client: %+v", err)
	}
	configureFunc(calendarViewInstanceClient.Client)

	calendarViewInstanceExceptionOccurrenceAttachmentClient, err := calendarviewinstanceexceptionoccurrenceattachment.NewCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(calendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	calendarViewInstanceExceptionOccurrenceCalendarClient, err := calendarviewinstanceexceptionoccurrencecalendar.NewCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(calendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	calendarViewInstanceExceptionOccurrenceClient, err := calendarviewinstanceexceptionoccurrence.NewCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(calendarViewInstanceExceptionOccurrenceClient.Client)

	calendarViewInstanceExceptionOccurrenceExtensionClient, err := calendarviewinstanceexceptionoccurrenceextension.NewCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(calendarViewInstanceExceptionOccurrenceExtensionClient.Client)

	calendarViewInstanceExtensionClient, err := calendarviewinstanceextension.NewCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(calendarViewInstanceExtensionClient.Client)

	chatClient, err := chat.NewChatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Chat client: %+v", err)
	}
	configureFunc(chatClient.Client)

	chatInstalledAppClient, err := chatinstalledapp.NewChatInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatInstalledApp client: %+v", err)
	}
	configureFunc(chatInstalledAppClient.Client)

	chatInstalledAppTeamsAppClient, err := chatinstalledappteamsapp.NewChatInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(chatInstalledAppTeamsAppClient.Client)

	chatInstalledAppTeamsAppDefinitionClient, err := chatinstalledappteamsappdefinition.NewChatInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(chatInstalledAppTeamsAppDefinitionClient.Client)

	chatLastMessagePreviewClient, err := chatlastmessagepreview.NewChatLastMessagePreviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatLastMessagePreview client: %+v", err)
	}
	configureFunc(chatLastMessagePreviewClient.Client)

	chatMemberClient, err := chatmember.NewChatMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatMember client: %+v", err)
	}
	configureFunc(chatMemberClient.Client)

	chatMessageClient, err := chatmessage.NewChatMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatMessage client: %+v", err)
	}
	configureFunc(chatMessageClient.Client)

	chatMessageHostedContentClient, err := chatmessagehostedcontent.NewChatMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatMessageHostedContent client: %+v", err)
	}
	configureFunc(chatMessageHostedContentClient.Client)

	chatMessageReplyClient, err := chatmessagereply.NewChatMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatMessageReply client: %+v", err)
	}
	configureFunc(chatMessageReplyClient.Client)

	chatMessageReplyHostedContentClient, err := chatmessagereplyhostedcontent.NewChatMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(chatMessageReplyHostedContentClient.Client)

	chatOperationClient, err := chatoperation.NewChatOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatOperation client: %+v", err)
	}
	configureFunc(chatOperationClient.Client)

	chatPermissionGrantClient, err := chatpermissiongrant.NewChatPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatPermissionGrant client: %+v", err)
	}
	configureFunc(chatPermissionGrantClient.Client)

	chatPinnedMessageClient, err := chatpinnedmessage.NewChatPinnedMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatPinnedMessage client: %+v", err)
	}
	configureFunc(chatPinnedMessageClient.Client)

	chatPinnedMessageMessageClient, err := chatpinnedmessagemessage.NewChatPinnedMessageMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatPinnedMessageMessage client: %+v", err)
	}
	configureFunc(chatPinnedMessageMessageClient.Client)

	chatTabClient, err := chattab.NewChatTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatTab client: %+v", err)
	}
	configureFunc(chatTabClient.Client)

	chatTabTeamsAppClient, err := chattabteamsapp.NewChatTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChatTabTeamsApp client: %+v", err)
	}
	configureFunc(chatTabTeamsAppClient.Client)

	cloudPCClient, err := cloudpc.NewCloudPCClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPC client: %+v", err)
	}
	configureFunc(cloudPCClient.Client)

	contactClient, err := contact.NewContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Contact client: %+v", err)
	}
	configureFunc(contactClient.Client)

	contactExtensionClient, err := contactextension.NewContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactExtension client: %+v", err)
	}
	configureFunc(contactExtensionClient.Client)

	contactFolderChildFolderClient, err := contactfolderchildfolder.NewContactFolderChildFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolderChildFolder client: %+v", err)
	}
	configureFunc(contactFolderChildFolderClient.Client)

	contactFolderChildFolderContactClient, err := contactfolderchildfoldercontact.NewContactFolderChildFolderContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolderChildFolderContact client: %+v", err)
	}
	configureFunc(contactFolderChildFolderContactClient.Client)

	contactFolderChildFolderContactExtensionClient, err := contactfolderchildfoldercontactextension.NewContactFolderChildFolderContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolderChildFolderContactExtension client: %+v", err)
	}
	configureFunc(contactFolderChildFolderContactExtensionClient.Client)

	contactFolderChildFolderContactPhotoClient, err := contactfolderchildfoldercontactphoto.NewContactFolderChildFolderContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolderChildFolderContactPhoto client: %+v", err)
	}
	configureFunc(contactFolderChildFolderContactPhotoClient.Client)

	contactFolderClient, err := contactfolder.NewContactFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolder client: %+v", err)
	}
	configureFunc(contactFolderClient.Client)

	contactFolderContactClient, err := contactfoldercontact.NewContactFolderContactClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolderContact client: %+v", err)
	}
	configureFunc(contactFolderContactClient.Client)

	contactFolderContactExtensionClient, err := contactfoldercontactextension.NewContactFolderContactExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolderContactExtension client: %+v", err)
	}
	configureFunc(contactFolderContactExtensionClient.Client)

	contactFolderContactPhotoClient, err := contactfoldercontactphoto.NewContactFolderContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactFolderContactPhoto client: %+v", err)
	}
	configureFunc(contactFolderContactPhotoClient.Client)

	contactPhotoClient, err := contactphoto.NewContactPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContactPhoto client: %+v", err)
	}
	configureFunc(contactPhotoClient.Client)

	createdObjectClient, err := createdobject.NewCreatedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreatedObject client: %+v", err)
	}
	configureFunc(createdObjectClient.Client)

	deviceClient, err := device.NewDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Device client: %+v", err)
	}
	configureFunc(deviceClient.Client)

	deviceCommandClient, err := devicecommand.NewDeviceCommandClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCommand client: %+v", err)
	}
	configureFunc(deviceCommandClient.Client)

	deviceCommandResponsepayloadClient, err := devicecommandresponsepayload.NewDeviceCommandResponsepayloadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCommandResponsepayload client: %+v", err)
	}
	configureFunc(deviceCommandResponsepayloadClient.Client)

	deviceEnrollmentConfigurationAssignmentClient, err := deviceenrollmentconfigurationassignment.NewDeviceEnrollmentConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceEnrollmentConfigurationAssignment client: %+v", err)
	}
	configureFunc(deviceEnrollmentConfigurationAssignmentClient.Client)

	deviceEnrollmentConfigurationClient, err := deviceenrollmentconfiguration.NewDeviceEnrollmentConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceEnrollmentConfiguration client: %+v", err)
	}
	configureFunc(deviceEnrollmentConfigurationClient.Client)

	deviceExtensionClient, err := deviceextension.NewDeviceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceExtension client: %+v", err)
	}
	configureFunc(deviceExtensionClient.Client)

	deviceManagementTroubleshootingEventClient, err := devicemanagementtroubleshootingevent.NewDeviceManagementTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementTroubleshootingEvent client: %+v", err)
	}
	configureFunc(deviceManagementTroubleshootingEventClient.Client)

	deviceMemberOfClient, err := devicememberof.NewDeviceMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceMemberOf client: %+v", err)
	}
	configureFunc(deviceMemberOfClient.Client)

	deviceRegisteredOwnerClient, err := deviceregisteredowner.NewDeviceRegisteredOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceRegisteredOwner client: %+v", err)
	}
	configureFunc(deviceRegisteredOwnerClient.Client)

	deviceRegisteredUserClient, err := deviceregistereduser.NewDeviceRegisteredUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceRegisteredUser client: %+v", err)
	}
	configureFunc(deviceRegisteredUserClient.Client)

	deviceTransitiveMemberOfClient, err := devicetransitivememberof.NewDeviceTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceTransitiveMemberOf client: %+v", err)
	}
	configureFunc(deviceTransitiveMemberOfClient.Client)

	deviceUsageRightClient, err := deviceusageright.NewDeviceUsageRightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceUsageRight client: %+v", err)
	}
	configureFunc(deviceUsageRightClient.Client)

	directReportClient, err := directreport.NewDirectReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectReport client: %+v", err)
	}
	configureFunc(directReportClient.Client)

	driveClient, err := drive.NewDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Drive client: %+v", err)
	}
	configureFunc(driveClient.Client)

	employeeExperienceClient, err := employeeexperience.NewEmployeeExperienceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EmployeeExperience client: %+v", err)
	}
	configureFunc(employeeExperienceClient.Client)

	employeeExperienceLearningCourseActivityClient, err := employeeexperiencelearningcourseactivity.NewEmployeeExperienceLearningCourseActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EmployeeExperienceLearningCourseActivity client: %+v", err)
	}
	configureFunc(employeeExperienceLearningCourseActivityClient.Client)

	eventAttachmentClient, err := eventattachment.NewEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventAttachment client: %+v", err)
	}
	configureFunc(eventAttachmentClient.Client)

	eventCalendarClient, err := eventcalendar.NewEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventCalendar client: %+v", err)
	}
	configureFunc(eventCalendarClient.Client)

	eventClient, err := event.NewEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Event client: %+v", err)
	}
	configureFunc(eventClient.Client)

	eventExceptionOccurrenceAttachmentClient, err := eventexceptionoccurrenceattachment.NewEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceAttachmentClient.Client)

	eventExceptionOccurrenceCalendarClient, err := eventexceptionoccurrencecalendar.NewEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceCalendarClient.Client)

	eventExceptionOccurrenceClient, err := eventexceptionoccurrence.NewEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrence client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceClient.Client)

	eventExceptionOccurrenceExtensionClient, err := eventexceptionoccurrenceextension.NewEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceExtensionClient.Client)

	eventExceptionOccurrenceInstanceAttachmentClient, err := eventexceptionoccurrenceinstanceattachment.NewEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceInstanceAttachmentClient.Client)

	eventExceptionOccurrenceInstanceCalendarClient, err := eventexceptionoccurrenceinstancecalendar.NewEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceInstanceCalendarClient.Client)

	eventExceptionOccurrenceInstanceClient, err := eventexceptionoccurrenceinstance.NewEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceInstanceClient.Client)

	eventExceptionOccurrenceInstanceExtensionClient, err := eventexceptionoccurrenceinstanceextension.NewEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(eventExceptionOccurrenceInstanceExtensionClient.Client)

	eventExtensionClient, err := eventextension.NewEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventExtension client: %+v", err)
	}
	configureFunc(eventExtensionClient.Client)

	eventInstanceAttachmentClient, err := eventinstanceattachment.NewEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstanceAttachment client: %+v", err)
	}
	configureFunc(eventInstanceAttachmentClient.Client)

	eventInstanceCalendarClient, err := eventinstancecalendar.NewEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstanceCalendar client: %+v", err)
	}
	configureFunc(eventInstanceCalendarClient.Client)

	eventInstanceClient, err := eventinstance.NewEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstance client: %+v", err)
	}
	configureFunc(eventInstanceClient.Client)

	eventInstanceExceptionOccurrenceAttachmentClient, err := eventinstanceexceptionoccurrenceattachment.NewEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(eventInstanceExceptionOccurrenceAttachmentClient.Client)

	eventInstanceExceptionOccurrenceCalendarClient, err := eventinstanceexceptionoccurrencecalendar.NewEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(eventInstanceExceptionOccurrenceCalendarClient.Client)

	eventInstanceExceptionOccurrenceClient, err := eventinstanceexceptionoccurrence.NewEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(eventInstanceExceptionOccurrenceClient.Client)

	eventInstanceExceptionOccurrenceExtensionClient, err := eventinstanceexceptionoccurrenceextension.NewEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(eventInstanceExceptionOccurrenceExtensionClient.Client)

	eventInstanceExtensionClient, err := eventinstanceextension.NewEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EventInstanceExtension client: %+v", err)
	}
	configureFunc(eventInstanceExtensionClient.Client)

	extensionClient, err := extension.NewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extension client: %+v", err)
	}
	configureFunc(extensionClient.Client)

	followedSiteClient, err := followedsite.NewFollowedSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FollowedSite client: %+v", err)
	}
	configureFunc(followedSiteClient.Client)

	inferenceClassificationClient, err := inferenceclassification.NewInferenceClassificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InferenceClassification client: %+v", err)
	}
	configureFunc(inferenceClassificationClient.Client)

	inferenceClassificationOverrideClient, err := inferenceclassificationoverride.NewInferenceClassificationOverrideClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InferenceClassificationOverride client: %+v", err)
	}
	configureFunc(inferenceClassificationOverrideClient.Client)

	informationProtectionBitlockerClient, err := informationprotectionbitlocker.NewInformationProtectionBitlockerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionBitlocker client: %+v", err)
	}
	configureFunc(informationProtectionBitlockerClient.Client)

	informationProtectionBitlockerRecoveryKeyClient, err := informationprotectionbitlockerrecoverykey.NewInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionBitlockerRecoveryKey client: %+v", err)
	}
	configureFunc(informationProtectionBitlockerRecoveryKeyClient.Client)

	informationProtectionClient, err := informationprotection.NewInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtection client: %+v", err)
	}
	configureFunc(informationProtectionClient.Client)

	informationProtectionDataLossPreventionPolicyClient, err := informationprotectiondatalosspreventionpolicy.NewInformationProtectionDataLossPreventionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionDataLossPreventionPolicy client: %+v", err)
	}
	configureFunc(informationProtectionDataLossPreventionPolicyClient.Client)

	informationProtectionPolicyClient, err := informationprotectionpolicy.NewInformationProtectionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionPolicy client: %+v", err)
	}
	configureFunc(informationProtectionPolicyClient.Client)

	informationProtectionPolicyLabelClient, err := informationprotectionpolicylabel.NewInformationProtectionPolicyLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionPolicyLabel client: %+v", err)
	}
	configureFunc(informationProtectionPolicyLabelClient.Client)

	informationProtectionSensitivityLabelClient, err := informationprotectionsensitivitylabel.NewInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(informationProtectionSensitivityLabelClient.Client)

	informationProtectionSensitivityLabelSublabelClient, err := informationprotectionsensitivitylabelsublabel.NewInformationProtectionSensitivityLabelSublabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionSensitivityLabelSublabel client: %+v", err)
	}
	configureFunc(informationProtectionSensitivityLabelSublabelClient.Client)

	informationProtectionSensitivityPolicySettingClient, err := informationprotectionsensitivitypolicysetting.NewInformationProtectionSensitivityPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionSensitivityPolicySetting client: %+v", err)
	}
	configureFunc(informationProtectionSensitivityPolicySettingClient.Client)

	informationProtectionThreatAssessmentRequestClient, err := informationprotectionthreatassessmentrequest.NewInformationProtectionThreatAssessmentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionThreatAssessmentRequest client: %+v", err)
	}
	configureFunc(informationProtectionThreatAssessmentRequestClient.Client)

	informationProtectionThreatAssessmentRequestResultClient, err := informationprotectionthreatassessmentrequestresult.NewInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionThreatAssessmentRequestResult client: %+v", err)
	}
	configureFunc(informationProtectionThreatAssessmentRequestResultClient.Client)

	insightClient, err := insight.NewInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Insight client: %+v", err)
	}
	configureFunc(insightClient.Client)

	insightSharedClient, err := insightshared.NewInsightSharedClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InsightShared client: %+v", err)
	}
	configureFunc(insightSharedClient.Client)

	insightSharedLastSharedMethodClient, err := insightsharedlastsharedmethod.NewInsightSharedLastSharedMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InsightSharedLastSharedMethod client: %+v", err)
	}
	configureFunc(insightSharedLastSharedMethodClient.Client)

	insightSharedResourceClient, err := insightsharedresource.NewInsightSharedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InsightSharedResource client: %+v", err)
	}
	configureFunc(insightSharedResourceClient.Client)

	insightTrendingClient, err := insighttrending.NewInsightTrendingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InsightTrending client: %+v", err)
	}
	configureFunc(insightTrendingClient.Client)

	insightTrendingResourceClient, err := insighttrendingresource.NewInsightTrendingResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InsightTrendingResource client: %+v", err)
	}
	configureFunc(insightTrendingResourceClient.Client)

	insightUsedClient, err := insightused.NewInsightUsedClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InsightUsed client: %+v", err)
	}
	configureFunc(insightUsedClient.Client)

	insightUsedResourceClient, err := insightusedresource.NewInsightUsedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InsightUsedResource client: %+v", err)
	}
	configureFunc(insightUsedResourceClient.Client)

	joinedGroupClient, err := joinedgroup.NewJoinedGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedGroup client: %+v", err)
	}
	configureFunc(joinedGroupClient.Client)

	joinedTeamClient, err := joinedteam.NewJoinedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeam client: %+v", err)
	}
	configureFunc(joinedTeamClient.Client)

	licenseDetailClient, err := licensedetail.NewLicenseDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LicenseDetail client: %+v", err)
	}
	configureFunc(licenseDetailClient.Client)

	mailFolderChildFolderClient, err := mailfolderchildfolder.NewMailFolderChildFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolder client: %+v", err)
	}
	configureFunc(mailFolderChildFolderClient.Client)

	mailFolderChildFolderMessageAttachmentClient, err := mailfolderchildfoldermessageattachment.NewMailFolderChildFolderMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolderMessageAttachment client: %+v", err)
	}
	configureFunc(mailFolderChildFolderMessageAttachmentClient.Client)

	mailFolderChildFolderMessageClient, err := mailfolderchildfoldermessage.NewMailFolderChildFolderMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolderMessage client: %+v", err)
	}
	configureFunc(mailFolderChildFolderMessageClient.Client)

	mailFolderChildFolderMessageExtensionClient, err := mailfolderchildfoldermessageextension.NewMailFolderChildFolderMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolderMessageExtension client: %+v", err)
	}
	configureFunc(mailFolderChildFolderMessageExtensionClient.Client)

	mailFolderChildFolderMessageMentionClient, err := mailfolderchildfoldermessagemention.NewMailFolderChildFolderMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolderMessageMention client: %+v", err)
	}
	configureFunc(mailFolderChildFolderMessageMentionClient.Client)

	mailFolderChildFolderMessageRuleClient, err := mailfolderchildfoldermessagerule.NewMailFolderChildFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolderMessageRule client: %+v", err)
	}
	configureFunc(mailFolderChildFolderMessageRuleClient.Client)

	mailFolderChildFolderUserConfigurationClient, err := mailfolderchildfolderuserconfiguration.NewMailFolderChildFolderUserConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolderUserConfiguration client: %+v", err)
	}
	configureFunc(mailFolderChildFolderUserConfigurationClient.Client)

	mailFolderClient, err := mailfolder.NewMailFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolder client: %+v", err)
	}
	configureFunc(mailFolderClient.Client)

	mailFolderMessageAttachmentClient, err := mailfoldermessageattachment.NewMailFolderMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderMessageAttachment client: %+v", err)
	}
	configureFunc(mailFolderMessageAttachmentClient.Client)

	mailFolderMessageClient, err := mailfoldermessage.NewMailFolderMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderMessage client: %+v", err)
	}
	configureFunc(mailFolderMessageClient.Client)

	mailFolderMessageExtensionClient, err := mailfoldermessageextension.NewMailFolderMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderMessageExtension client: %+v", err)
	}
	configureFunc(mailFolderMessageExtensionClient.Client)

	mailFolderMessageMentionClient, err := mailfoldermessagemention.NewMailFolderMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderMessageMention client: %+v", err)
	}
	configureFunc(mailFolderMessageMentionClient.Client)

	mailFolderMessageRuleClient, err := mailfoldermessagerule.NewMailFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderMessageRule client: %+v", err)
	}
	configureFunc(mailFolderMessageRuleClient.Client)

	mailFolderUserConfigurationClient, err := mailfolderuserconfiguration.NewMailFolderUserConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderUserConfiguration client: %+v", err)
	}
	configureFunc(mailFolderUserConfigurationClient.Client)

	mailboxSettingClient, err := mailboxsetting.NewMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailboxSetting client: %+v", err)
	}
	configureFunc(mailboxSettingClient.Client)

	managedAppRegistrationClient, err := managedappregistration.NewManagedAppRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedAppRegistration client: %+v", err)
	}
	configureFunc(managedAppRegistrationClient.Client)

	managedDeviceAssignmentFilterEvaluationStatusDetailClient, err := manageddeviceassignmentfilterevaluationstatusdetail.NewManagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceAssignmentFilterEvaluationStatusDetail client: %+v", err)
	}
	configureFunc(managedDeviceAssignmentFilterEvaluationStatusDetailClient.Client)

	managedDeviceClient, err := manageddevice.NewManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDevice client: %+v", err)
	}
	configureFunc(managedDeviceClient.Client)

	managedDeviceDetectedAppClient, err := manageddevicedetectedapp.NewManagedDeviceDetectedAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDetectedApp client: %+v", err)
	}
	configureFunc(managedDeviceDetectedAppClient.Client)

	managedDeviceDeviceCategoryClient, err := manageddevicedevicecategory.NewManagedDeviceDeviceCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceCategory client: %+v", err)
	}
	configureFunc(managedDeviceDeviceCategoryClient.Client)

	managedDeviceDeviceCompliancePolicyStateClient, err := manageddevicedevicecompliancepolicystate.NewManagedDeviceDeviceCompliancePolicyStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceCompliancePolicyState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceCompliancePolicyStateClient.Client)

	managedDeviceDeviceConfigurationStateClient, err := manageddevicedeviceconfigurationstate.NewManagedDeviceDeviceConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceConfigurationState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceConfigurationStateClient.Client)

	managedDeviceDeviceHealthScriptStateClient, err := manageddevicedevicehealthscriptstate.NewManagedDeviceDeviceHealthScriptStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceHealthScriptState client: %+v", err)
	}
	configureFunc(managedDeviceDeviceHealthScriptStateClient.Client)

	managedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient, err := manageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid.NewManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId client: %+v", err)
	}
	configureFunc(managedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient.Client)

	managedDeviceLogCollectionRequestClient, err := manageddevicelogcollectionrequest.NewManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(managedDeviceLogCollectionRequestClient.Client)

	managedDeviceManagedDeviceMobileAppConfigurationStateClient, err := manageddevicemanageddevicemobileappconfigurationstate.NewManagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceManagedDeviceMobileAppConfigurationState client: %+v", err)
	}
	configureFunc(managedDeviceManagedDeviceMobileAppConfigurationStateClient.Client)

	managedDeviceSecurityBaselineStateClient, err := manageddevicesecuritybaselinestate.NewManagedDeviceSecurityBaselineStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceSecurityBaselineState client: %+v", err)
	}
	configureFunc(managedDeviceSecurityBaselineStateClient.Client)

	managedDeviceSecurityBaselineStateSettingStateClient, err := manageddevicesecuritybaselinestatesettingstate.NewManagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceSecurityBaselineStateSettingState client: %+v", err)
	}
	configureFunc(managedDeviceSecurityBaselineStateSettingStateClient.Client)

	managedDeviceUserClient, err := manageddeviceuser.NewManagedDeviceUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceUser client: %+v", err)
	}
	configureFunc(managedDeviceUserClient.Client)

	managedDeviceWindowsProtectionStateClient, err := manageddevicewindowsprotectionstate.NewManagedDeviceWindowsProtectionStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceWindowsProtectionState client: %+v", err)
	}
	configureFunc(managedDeviceWindowsProtectionStateClient.Client)

	managedDeviceWindowsProtectionStateDetectedMalwareStateClient, err := manageddevicewindowsprotectionstatedetectedmalwarestate.NewManagedDeviceWindowsProtectionStateDetectedMalwareStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceWindowsProtectionStateDetectedMalwareState client: %+v", err)
	}
	configureFunc(managedDeviceWindowsProtectionStateDetectedMalwareStateClient.Client)

	managerClient, err := manager.NewManagerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Manager client: %+v", err)
	}
	configureFunc(managerClient.Client)

	memberOfClient, err := memberof.NewMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MemberOf client: %+v", err)
	}
	configureFunc(memberOfClient.Client)

	messageAttachmentClient, err := messageattachment.NewMessageAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MessageAttachment client: %+v", err)
	}
	configureFunc(messageAttachmentClient.Client)

	messageClient, err := message.NewMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Message client: %+v", err)
	}
	configureFunc(messageClient.Client)

	messageExtensionClient, err := messageextension.NewMessageExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MessageExtension client: %+v", err)
	}
	configureFunc(messageExtensionClient.Client)

	messageMentionClient, err := messagemention.NewMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MessageMention client: %+v", err)
	}
	configureFunc(messageMentionClient.Client)

	mobileAppIntentAndStateClient, err := mobileappintentandstate.NewMobileAppIntentAndStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppIntentAndState client: %+v", err)
	}
	configureFunc(mobileAppIntentAndStateClient.Client)

	mobileAppTroubleshootingEventAppLogCollectionRequestClient, err := mobileapptroubleshootingeventapplogcollectionrequest.NewMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppTroubleshootingEventAppLogCollectionRequest client: %+v", err)
	}
	configureFunc(mobileAppTroubleshootingEventAppLogCollectionRequestClient.Client)

	mobileAppTroubleshootingEventClient, err := mobileapptroubleshootingevent.NewMobileAppTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MobileAppTroubleshootingEvent client: %+v", err)
	}
	configureFunc(mobileAppTroubleshootingEventClient.Client)

	notificationClient, err := notification.NewNotificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Notification client: %+v", err)
	}
	configureFunc(notificationClient.Client)

	oauth2PermissionGrantClient, err := oauth2permissiongrant.NewOauth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Oauth2PermissionGrant client: %+v", err)
	}
	configureFunc(oauth2PermissionGrantClient.Client)

	onlineMeetingAlternativeRecordingClient, err := onlinemeetingalternativerecording.NewOnlineMeetingAlternativeRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingAlternativeRecording client: %+v", err)
	}
	configureFunc(onlineMeetingAlternativeRecordingClient.Client)

	onlineMeetingAttendanceReportAttendanceRecordClient, err := onlinemeetingattendancereportattendancerecord.NewOnlineMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingAttendanceReportAttendanceRecord client: %+v", err)
	}
	configureFunc(onlineMeetingAttendanceReportAttendanceRecordClient.Client)

	onlineMeetingAttendanceReportClient, err := onlinemeetingattendancereport.NewOnlineMeetingAttendanceReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingAttendanceReport client: %+v", err)
	}
	configureFunc(onlineMeetingAttendanceReportClient.Client)

	onlineMeetingAttendeeReportClient, err := onlinemeetingattendeereport.NewOnlineMeetingAttendeeReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingAttendeeReport client: %+v", err)
	}
	configureFunc(onlineMeetingAttendeeReportClient.Client)

	onlineMeetingBroadcastRecordingClient, err := onlinemeetingbroadcastrecording.NewOnlineMeetingBroadcastRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingBroadcastRecording client: %+v", err)
	}
	configureFunc(onlineMeetingBroadcastRecordingClient.Client)

	onlineMeetingClient, err := onlinemeeting.NewOnlineMeetingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeeting client: %+v", err)
	}
	configureFunc(onlineMeetingClient.Client)

	onlineMeetingMeetingAttendanceReportAttendanceRecordClient, err := onlinemeetingmeetingattendancereportattendancerecord.NewOnlineMeetingMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingMeetingAttendanceReportAttendanceRecord client: %+v", err)
	}
	configureFunc(onlineMeetingMeetingAttendanceReportAttendanceRecordClient.Client)

	onlineMeetingMeetingAttendanceReportClient, err := onlinemeetingmeetingattendancereport.NewOnlineMeetingMeetingAttendanceReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingMeetingAttendanceReport client: %+v", err)
	}
	configureFunc(onlineMeetingMeetingAttendanceReportClient.Client)

	onlineMeetingRecordingClient, err := onlinemeetingrecording.NewOnlineMeetingRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingRecording client: %+v", err)
	}
	configureFunc(onlineMeetingRecordingClient.Client)

	onlineMeetingRecordingContentClient, err := onlinemeetingrecordingcontent.NewOnlineMeetingRecordingContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingRecordingContent client: %+v", err)
	}
	configureFunc(onlineMeetingRecordingContentClient.Client)

	onlineMeetingRegistrationClient, err := onlinemeetingregistration.NewOnlineMeetingRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingRegistration client: %+v", err)
	}
	configureFunc(onlineMeetingRegistrationClient.Client)

	onlineMeetingRegistrationCustomQuestionClient, err := onlinemeetingregistrationcustomquestion.NewOnlineMeetingRegistrationCustomQuestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingRegistrationCustomQuestion client: %+v", err)
	}
	configureFunc(onlineMeetingRegistrationCustomQuestionClient.Client)

	onlineMeetingRegistrationRegistrantClient, err := onlinemeetingregistrationregistrant.NewOnlineMeetingRegistrationRegistrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingRegistrationRegistrant client: %+v", err)
	}
	configureFunc(onlineMeetingRegistrationRegistrantClient.Client)

	onlineMeetingTranscriptClient, err := onlinemeetingtranscript.NewOnlineMeetingTranscriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingTranscript client: %+v", err)
	}
	configureFunc(onlineMeetingTranscriptClient.Client)

	onlineMeetingTranscriptContentClient, err := onlinemeetingtranscriptcontent.NewOnlineMeetingTranscriptContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingTranscriptContent client: %+v", err)
	}
	configureFunc(onlineMeetingTranscriptContentClient.Client)

	onlineMeetingTranscriptMetadataContentClient, err := onlinemeetingtranscriptmetadatacontent.NewOnlineMeetingTranscriptMetadataContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeetingTranscriptMetadataContent client: %+v", err)
	}
	configureFunc(onlineMeetingTranscriptMetadataContentClient.Client)

	outlookClient, err := outlook.NewOutlookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Outlook client: %+v", err)
	}
	configureFunc(outlookClient.Client)

	outlookMasterCategoryClient, err := outlookmastercategory.NewOutlookMasterCategoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookMasterCategory client: %+v", err)
	}
	configureFunc(outlookMasterCategoryClient.Client)

	outlookTaskAttachmentClient, err := outlooktaskattachment.NewOutlookTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskAttachment client: %+v", err)
	}
	configureFunc(outlookTaskAttachmentClient.Client)

	outlookTaskClient, err := outlooktask.NewOutlookTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTask client: %+v", err)
	}
	configureFunc(outlookTaskClient.Client)

	outlookTaskFolderClient, err := outlooktaskfolder.NewOutlookTaskFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskFolder client: %+v", err)
	}
	configureFunc(outlookTaskFolderClient.Client)

	outlookTaskFolderTaskAttachmentClient, err := outlooktaskfoldertaskattachment.NewOutlookTaskFolderTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskFolderTaskAttachment client: %+v", err)
	}
	configureFunc(outlookTaskFolderTaskAttachmentClient.Client)

	outlookTaskFolderTaskClient, err := outlooktaskfoldertask.NewOutlookTaskFolderTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskFolderTask client: %+v", err)
	}
	configureFunc(outlookTaskFolderTaskClient.Client)

	outlookTaskGroupClient, err := outlooktaskgroup.NewOutlookTaskGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskGroup client: %+v", err)
	}
	configureFunc(outlookTaskGroupClient.Client)

	outlookTaskGroupTaskFolderClient, err := outlooktaskgrouptaskfolder.NewOutlookTaskGroupTaskFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskGroupTaskFolder client: %+v", err)
	}
	configureFunc(outlookTaskGroupTaskFolderClient.Client)

	outlookTaskGroupTaskFolderTaskAttachmentClient, err := outlooktaskgrouptaskfoldertaskattachment.NewOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskGroupTaskFolderTaskAttachment client: %+v", err)
	}
	configureFunc(outlookTaskGroupTaskFolderTaskAttachmentClient.Client)

	outlookTaskGroupTaskFolderTaskClient, err := outlooktaskgrouptaskfoldertask.NewOutlookTaskGroupTaskFolderTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutlookTaskGroupTaskFolderTask client: %+v", err)
	}
	configureFunc(outlookTaskGroupTaskFolderTaskClient.Client)

	ownedDeviceClient, err := owneddevice.NewOwnedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OwnedDevice client: %+v", err)
	}
	configureFunc(ownedDeviceClient.Client)

	ownedObjectClient, err := ownedobject.NewOwnedObjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OwnedObject client: %+v", err)
	}
	configureFunc(ownedObjectClient.Client)

	pendingAccessReviewInstanceClient, err := pendingaccessreviewinstance.NewPendingAccessReviewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstance client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceClient.Client)

	pendingAccessReviewInstanceContactedReviewerClient, err := pendingaccessreviewinstancecontactedreviewer.NewPendingAccessReviewInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceContactedReviewerClient.Client)

	pendingAccessReviewInstanceDecisionClient, err := pendingaccessreviewinstancedecision.NewPendingAccessReviewInstanceDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDecision client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDecisionClient.Client)

	pendingAccessReviewInstanceDecisionInsightClient, err := pendingaccessreviewinstancedecisioninsight.NewPendingAccessReviewInstanceDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDecisionInsight client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDecisionInsightClient.Client)

	pendingAccessReviewInstanceDecisionInstanceClient, err := pendingaccessreviewinstancedecisioninstance.NewPendingAccessReviewInstanceDecisionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDecisionInstance client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDecisionInstanceClient.Client)

	pendingAccessReviewInstanceDecisionInstanceContactedReviewerClient, err := pendingaccessreviewinstancedecisioninstancecontactedreviewer.NewPendingAccessReviewInstanceDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDecisionInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDecisionInstanceContactedReviewerClient.Client)

	pendingAccessReviewInstanceDecisionInstanceDefinitionClient, err := pendingaccessreviewinstancedecisioninstancedefinition.NewPendingAccessReviewInstanceDecisionInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDecisionInstanceDefinition client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDecisionInstanceDefinitionClient.Client)

	pendingAccessReviewInstanceDecisionInstanceStageClient, err := pendingaccessreviewinstancedecisioninstancestage.NewPendingAccessReviewInstanceDecisionInstanceStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDecisionInstanceStage client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDecisionInstanceStageClient.Client)

	pendingAccessReviewInstanceDefinitionClient, err := pendingaccessreviewinstancedefinition.NewPendingAccessReviewInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDefinition client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDefinitionClient.Client)

	pendingAccessReviewInstanceStageClient, err := pendingaccessreviewinstancestage.NewPendingAccessReviewInstanceStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStage client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageClient.Client)

	pendingAccessReviewInstanceStageDecisionClient, err := pendingaccessreviewinstancestagedecision.NewPendingAccessReviewInstanceStageDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStageDecision client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageDecisionClient.Client)

	pendingAccessReviewInstanceStageDecisionInsightClient, err := pendingaccessreviewinstancestagedecisioninsight.NewPendingAccessReviewInstanceStageDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStageDecisionInsight client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageDecisionInsightClient.Client)

	pendingAccessReviewInstanceStageDecisionInstanceClient, err := pendingaccessreviewinstancestagedecisioninstance.NewPendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStageDecisionInstance client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageDecisionInstanceClient.Client)

	pendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient, err := pendingaccessreviewinstancestagedecisioninstancecontactedreviewer.NewPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStageDecisionInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient.Client)

	pendingAccessReviewInstanceStageDecisionInstanceDefinitionClient, err := pendingaccessreviewinstancestagedecisioninstancedefinition.NewPendingAccessReviewInstanceStageDecisionInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStageDecisionInstanceDefinition client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageDecisionInstanceDefinitionClient.Client)

	peopleClient, err := people.NewPeopleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building People client: %+v", err)
	}
	configureFunc(peopleClient.Client)

	permissionGrantClient, err := permissiongrant.NewPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PermissionGrant client: %+v", err)
	}
	configureFunc(permissionGrantClient.Client)

	photoClient, err := photo.NewPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Photo client: %+v", err)
	}
	configureFunc(photoClient.Client)

	plannerAllClient, err := plannerall.NewPlannerAllClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerAll client: %+v", err)
	}
	configureFunc(plannerAllClient.Client)

	plannerClient, err := planner.NewPlannerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Planner client: %+v", err)
	}
	configureFunc(plannerClient.Client)

	plannerFavoritePlanClient, err := plannerfavoriteplan.NewPlannerFavoritePlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerFavoritePlan client: %+v", err)
	}
	configureFunc(plannerFavoritePlanClient.Client)

	plannerPlanBucketClient, err := plannerplanbucket.NewPlannerPlanBucketClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanBucket client: %+v", err)
	}
	configureFunc(plannerPlanBucketClient.Client)

	plannerPlanBucketTaskAssignedToTaskBoardFormatClient, err := plannerplanbuckettaskassignedtotaskboardformat.NewPlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanBucketTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerPlanBucketTaskAssignedToTaskBoardFormatClient.Client)

	plannerPlanBucketTaskBucketTaskBoardFormatClient, err := plannerplanbuckettaskbuckettaskboardformat.NewPlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanBucketTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerPlanBucketTaskBucketTaskBoardFormatClient.Client)

	plannerPlanBucketTaskClient, err := plannerplanbuckettask.NewPlannerPlanBucketTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanBucketTask client: %+v", err)
	}
	configureFunc(plannerPlanBucketTaskClient.Client)

	plannerPlanBucketTaskDetailClient, err := plannerplanbuckettaskdetail.NewPlannerPlanBucketTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanBucketTaskDetail client: %+v", err)
	}
	configureFunc(plannerPlanBucketTaskDetailClient.Client)

	plannerPlanBucketTaskProgressTaskBoardFormatClient, err := plannerplanbuckettaskprogresstaskboardformat.NewPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanBucketTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerPlanBucketTaskProgressTaskBoardFormatClient.Client)

	plannerPlanClient, err := plannerplan.NewPlannerPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlan client: %+v", err)
	}
	configureFunc(plannerPlanClient.Client)

	plannerPlanDetailClient, err := plannerplandetail.NewPlannerPlanDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanDetail client: %+v", err)
	}
	configureFunc(plannerPlanDetailClient.Client)

	plannerPlanTaskAssignedToTaskBoardFormatClient, err := plannerplantaskassignedtotaskboardformat.NewPlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerPlanTaskAssignedToTaskBoardFormatClient.Client)

	plannerPlanTaskBucketTaskBoardFormatClient, err := plannerplantaskbuckettaskboardformat.NewPlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerPlanTaskBucketTaskBoardFormatClient.Client)

	plannerPlanTaskClient, err := plannerplantask.NewPlannerPlanTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanTask client: %+v", err)
	}
	configureFunc(plannerPlanTaskClient.Client)

	plannerPlanTaskDetailClient, err := plannerplantaskdetail.NewPlannerPlanTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanTaskDetail client: %+v", err)
	}
	configureFunc(plannerPlanTaskDetailClient.Client)

	plannerPlanTaskProgressTaskBoardFormatClient, err := plannerplantaskprogresstaskboardformat.NewPlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerPlanTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerPlanTaskProgressTaskBoardFormatClient.Client)

	plannerRecentPlanClient, err := plannerrecentplan.NewPlannerRecentPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerRecentPlan client: %+v", err)
	}
	configureFunc(plannerRecentPlanClient.Client)

	plannerRosterPlanClient, err := plannerrosterplan.NewPlannerRosterPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerRosterPlan client: %+v", err)
	}
	configureFunc(plannerRosterPlanClient.Client)

	plannerTaskAssignedToTaskBoardFormatClient, err := plannertaskassignedtotaskboardformat.NewPlannerTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerTaskAssignedToTaskBoardFormatClient.Client)

	plannerTaskBucketTaskBoardFormatClient, err := plannertaskbuckettaskboardformat.NewPlannerTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerTaskBucketTaskBoardFormatClient.Client)

	plannerTaskClient, err := plannertask.NewPlannerTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerTask client: %+v", err)
	}
	configureFunc(plannerTaskClient.Client)

	plannerTaskDetailClient, err := plannertaskdetail.NewPlannerTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerTaskDetail client: %+v", err)
	}
	configureFunc(plannerTaskDetailClient.Client)

	plannerTaskProgressTaskBoardFormatClient, err := plannertaskprogresstaskboardformat.NewPlannerTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(plannerTaskProgressTaskBoardFormatClient.Client)

	presenceClient, err := presence.NewPresenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Presence client: %+v", err)
	}
	configureFunc(presenceClient.Client)

	profileAccountClient, err := profileaccount.NewProfileAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileAccount client: %+v", err)
	}
	configureFunc(profileAccountClient.Client)

	profileAddressClient, err := profileaddress.NewProfileAddressClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileAddress client: %+v", err)
	}
	configureFunc(profileAddressClient.Client)

	profileAnniversaryClient, err := profileanniversary.NewProfileAnniversaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileAnniversary client: %+v", err)
	}
	configureFunc(profileAnniversaryClient.Client)

	profileAwardClient, err := profileaward.NewProfileAwardClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileAward client: %+v", err)
	}
	configureFunc(profileAwardClient.Client)

	profileCertificationClient, err := profilecertification.NewProfileCertificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileCertification client: %+v", err)
	}
	configureFunc(profileCertificationClient.Client)

	profileClient, err := profile.NewProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Profile client: %+v", err)
	}
	configureFunc(profileClient.Client)

	profileEducationalActivityClient, err := profileeducationalactivity.NewProfileEducationalActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileEducationalActivity client: %+v", err)
	}
	configureFunc(profileEducationalActivityClient.Client)

	profileEmailClient, err := profileemail.NewProfileEmailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileEmail client: %+v", err)
	}
	configureFunc(profileEmailClient.Client)

	profileInterestClient, err := profileinterest.NewProfileInterestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileInterest client: %+v", err)
	}
	configureFunc(profileInterestClient.Client)

	profileLanguageClient, err := profilelanguage.NewProfileLanguageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileLanguage client: %+v", err)
	}
	configureFunc(profileLanguageClient.Client)

	profileNameClient, err := profilename.NewProfileNameClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileName client: %+v", err)
	}
	configureFunc(profileNameClient.Client)

	profileNoteClient, err := profilenote.NewProfileNoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileNote client: %+v", err)
	}
	configureFunc(profileNoteClient.Client)

	profilePatentClient, err := profilepatent.NewProfilePatentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfilePatent client: %+v", err)
	}
	configureFunc(profilePatentClient.Client)

	profilePhoneClient, err := profilephone.NewProfilePhoneClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfilePhone client: %+v", err)
	}
	configureFunc(profilePhoneClient.Client)

	profilePositionClient, err := profileposition.NewProfilePositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfilePosition client: %+v", err)
	}
	configureFunc(profilePositionClient.Client)

	profileProjectClient, err := profileproject.NewProfileProjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileProject client: %+v", err)
	}
	configureFunc(profileProjectClient.Client)

	profilePublicationClient, err := profilepublication.NewProfilePublicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfilePublication client: %+v", err)
	}
	configureFunc(profilePublicationClient.Client)

	profileSkillClient, err := profileskill.NewProfileSkillClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileSkill client: %+v", err)
	}
	configureFunc(profileSkillClient.Client)

	profileWebAccountClient, err := profilewebaccount.NewProfileWebAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileWebAccount client: %+v", err)
	}
	configureFunc(profileWebAccountClient.Client)

	profileWebsiteClient, err := profilewebsite.NewProfileWebsiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProfileWebsite client: %+v", err)
	}
	configureFunc(profileWebsiteClient.Client)

	registeredDeviceClient, err := registereddevice.NewRegisteredDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegisteredDevice client: %+v", err)
	}
	configureFunc(registeredDeviceClient.Client)

	scopedRoleMemberOfClient, err := scopedrolememberof.NewScopedRoleMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScopedRoleMemberOf client: %+v", err)
	}
	configureFunc(scopedRoleMemberOfClient.Client)

	securityClient, err := security.NewSecurityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Security client: %+v", err)
	}
	configureFunc(securityClient.Client)

	securityInformationProtectionClient, err := securityinformationprotection.NewSecurityInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityInformationProtection client: %+v", err)
	}
	configureFunc(securityInformationProtectionClient.Client)

	securityInformationProtectionLabelPolicySettingClient, err := securityinformationprotectionlabelpolicysetting.NewSecurityInformationProtectionLabelPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityInformationProtectionLabelPolicySetting client: %+v", err)
	}
	configureFunc(securityInformationProtectionLabelPolicySettingClient.Client)

	securityInformationProtectionSensitivityLabelClient, err := securityinformationprotectionsensitivitylabel.NewSecurityInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityInformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(securityInformationProtectionSensitivityLabelClient.Client)

	securityInformationProtectionSensitivityLabelParentClient, err := securityinformationprotectionsensitivitylabelparent.NewSecurityInformationProtectionSensitivityLabelParentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityInformationProtectionSensitivityLabelParent client: %+v", err)
	}
	configureFunc(securityInformationProtectionSensitivityLabelParentClient.Client)

	serviceProvisioningErrorClient, err := serviceprovisioningerror.NewServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceProvisioningError client: %+v", err)
	}
	configureFunc(serviceProvisioningErrorClient.Client)

	settingClient, err := setting.NewSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Setting client: %+v", err)
	}
	configureFunc(settingClient.Client)

	settingContactMergeSuggestionClient, err := settingcontactmergesuggestion.NewSettingContactMergeSuggestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingContactMergeSuggestion client: %+v", err)
	}
	configureFunc(settingContactMergeSuggestionClient.Client)

	settingItemInsightClient, err := settingiteminsight.NewSettingItemInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingItemInsight client: %+v", err)
	}
	configureFunc(settingItemInsightClient.Client)

	settingRegionalAndLanguageSettingClient, err := settingregionalandlanguagesetting.NewSettingRegionalAndLanguageSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingRegionalAndLanguageSetting client: %+v", err)
	}
	configureFunc(settingRegionalAndLanguageSettingClient.Client)

	settingShiftPreferenceClient, err := settingshiftpreference.NewSettingShiftPreferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingShiftPreference client: %+v", err)
	}
	configureFunc(settingShiftPreferenceClient.Client)

	sponsorClient, err := sponsor.NewSponsorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Sponsor client: %+v", err)
	}
	configureFunc(sponsorClient.Client)

	teamworkAssociatedTeamClient, err := teamworkassociatedteam.NewTeamworkAssociatedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamworkAssociatedTeam client: %+v", err)
	}
	configureFunc(teamworkAssociatedTeamClient.Client)

	teamworkAssociatedTeamTeamClient, err := teamworkassociatedteamteam.NewTeamworkAssociatedTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamworkAssociatedTeamTeam client: %+v", err)
	}
	configureFunc(teamworkAssociatedTeamTeamClient.Client)

	teamworkClient, err := teamwork.NewTeamworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Teamwork client: %+v", err)
	}
	configureFunc(teamworkClient.Client)

	teamworkInstalledAppChatClient, err := teamworkinstalledappchat.NewTeamworkInstalledAppChatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamworkInstalledAppChat client: %+v", err)
	}
	configureFunc(teamworkInstalledAppChatClient.Client)

	teamworkInstalledAppClient, err := teamworkinstalledapp.NewTeamworkInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamworkInstalledApp client: %+v", err)
	}
	configureFunc(teamworkInstalledAppClient.Client)

	teamworkInstalledAppTeamsAppClient, err := teamworkinstalledappteamsapp.NewTeamworkInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamworkInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(teamworkInstalledAppTeamsAppClient.Client)

	teamworkInstalledAppTeamsAppDefinitionClient, err := teamworkinstalledappteamsappdefinition.NewTeamworkInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamworkInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(teamworkInstalledAppTeamsAppDefinitionClient.Client)

	todoClient, err := todo.NewTodoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Todo client: %+v", err)
	}
	configureFunc(todoClient.Client)

	todoListClient, err := todolist.NewTodoListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoList client: %+v", err)
	}
	configureFunc(todoListClient.Client)

	todoListExtensionClient, err := todolistextension.NewTodoListExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListExtension client: %+v", err)
	}
	configureFunc(todoListExtensionClient.Client)

	todoListTaskAttachmentClient, err := todolisttaskattachment.NewTodoListTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListTaskAttachment client: %+v", err)
	}
	configureFunc(todoListTaskAttachmentClient.Client)

	todoListTaskAttachmentSessionClient, err := todolisttaskattachmentsession.NewTodoListTaskAttachmentSessionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListTaskAttachmentSession client: %+v", err)
	}
	configureFunc(todoListTaskAttachmentSessionClient.Client)

	todoListTaskAttachmentSessionContentClient, err := todolisttaskattachmentsessioncontent.NewTodoListTaskAttachmentSessionContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListTaskAttachmentSessionContent client: %+v", err)
	}
	configureFunc(todoListTaskAttachmentSessionContentClient.Client)

	todoListTaskChecklistItemClient, err := todolisttaskchecklistitem.NewTodoListTaskChecklistItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListTaskChecklistItem client: %+v", err)
	}
	configureFunc(todoListTaskChecklistItemClient.Client)

	todoListTaskClient, err := todolisttask.NewTodoListTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListTask client: %+v", err)
	}
	configureFunc(todoListTaskClient.Client)

	todoListTaskExtensionClient, err := todolisttaskextension.NewTodoListTaskExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListTaskExtension client: %+v", err)
	}
	configureFunc(todoListTaskExtensionClient.Client)

	todoListTaskLinkedResourceClient, err := todolisttasklinkedresource.NewTodoListTaskLinkedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TodoListTaskLinkedResource client: %+v", err)
	}
	configureFunc(todoListTaskLinkedResourceClient.Client)

	transitiveMemberOfClient, err := transitivememberof.NewTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TransitiveMemberOf client: %+v", err)
	}
	configureFunc(transitiveMemberOfClient.Client)

	transitiveReportClient, err := transitivereport.NewTransitiveReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TransitiveReport client: %+v", err)
	}
	configureFunc(transitiveReportClient.Client)

	usageRightClient, err := usageright.NewUsageRightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UsageRight client: %+v", err)
	}
	configureFunc(usageRightClient.Client)

	userClient, err := user.NewUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building User client: %+v", err)
	}
	configureFunc(userClient.Client)

	virtualEventClient, err := virtualevent.NewVirtualEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEvent client: %+v", err)
	}
	configureFunc(virtualEventClient.Client)

	virtualEventWebinarClient, err := virtualeventwebinar.NewVirtualEventWebinarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualEventWebinar client: %+v", err)
	}
	configureFunc(virtualEventWebinarClient.Client)

	windowsInformationProtectionDeviceRegistrationClient, err := windowsinformationprotectiondeviceregistration.NewWindowsInformationProtectionDeviceRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WindowsInformationProtectionDeviceRegistration client: %+v", err)
	}
	configureFunc(windowsInformationProtectionDeviceRegistrationClient.Client)

	return &Client{
		Activity:                      activityClient,
		ActivityHistoryItem:           activityHistoryItemClient,
		ActivityHistoryItemActivity:   activityHistoryItemActivityClient,
		AgreementAcceptance:           agreementAcceptanceClient,
		Analytic:                      analyticClient,
		AnalyticActivityStatistic:     analyticActivityStatisticClient,
		AppConsentRequestsForApproval: appConsentRequestsForApprovalClient,
		AppConsentRequestsForApprovalUserConsentRequest:              appConsentRequestsForApprovalUserConsentRequestClient,
		AppConsentRequestsForApprovalUserConsentRequestApproval:      appConsentRequestsForApprovalUserConsentRequestApprovalClient,
		AppConsentRequestsForApprovalUserConsentRequestApprovalStep:  appConsentRequestsForApprovalUserConsentRequestApprovalStepClient,
		AppRoleAssignedResource:                                      appRoleAssignedResourceClient,
		AppRoleAssignment:                                            appRoleAssignmentClient,
		Approval:                                                     approvalClient,
		ApprovalStep:                                                 approvalStepClient,
		Authentication:                                               authenticationClient,
		AuthenticationEmailMethod:                                    authenticationEmailMethodClient,
		AuthenticationFido2Method:                                    authenticationFido2MethodClient,
		AuthenticationMethod:                                         authenticationMethodClient,
		AuthenticationMicrosoftAuthenticatorMethod:                   authenticationMicrosoftAuthenticatorMethodClient,
		AuthenticationMicrosoftAuthenticatorMethodDevice:             authenticationMicrosoftAuthenticatorMethodDeviceClient,
		AuthenticationOperation:                                      authenticationOperationClient,
		AuthenticationPasswordMethod:                                 authenticationPasswordMethodClient,
		AuthenticationPasswordlessMicrosoftAuthenticatorMethod:       authenticationPasswordlessMicrosoftAuthenticatorMethodClient,
		AuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice: authenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient,
		AuthenticationPhoneMethod:                                    authenticationPhoneMethodClient,
		AuthenticationPlatformCredentialMethod:                       authenticationPlatformCredentialMethodClient,
		AuthenticationPlatformCredentialMethodDevice:                 authenticationPlatformCredentialMethodDeviceClient,
		AuthenticationSignInPreference:                               authenticationSignInPreferenceClient,
		AuthenticationSoftwareOathMethod:                             authenticationSoftwareOathMethodClient,
		AuthenticationTemporaryAccessPassMethod:                      authenticationTemporaryAccessPassMethodClient,
		AuthenticationWindowsHelloForBusinessMethod:                  authenticationWindowsHelloForBusinessMethodClient,
		AuthenticationWindowsHelloForBusinessMethodDevice:            authenticationWindowsHelloForBusinessMethodDeviceClient,
		Calendar:                                                               calendarClient,
		CalendarCalendarPermission:                                             calendarCalendarPermissionClient,
		CalendarCalendarView:                                                   calendarCalendarViewClient,
		CalendarCalendarViewAttachment:                                         calendarCalendarViewAttachmentClient,
		CalendarCalendarViewCalendar:                                           calendarCalendarViewCalendarClient,
		CalendarCalendarViewExceptionOccurrence:                                calendarCalendarViewExceptionOccurrenceClient,
		CalendarCalendarViewExceptionOccurrenceAttachment:                      calendarCalendarViewExceptionOccurrenceAttachmentClient,
		CalendarCalendarViewExceptionOccurrenceCalendar:                        calendarCalendarViewExceptionOccurrenceCalendarClient,
		CalendarCalendarViewExceptionOccurrenceExtension:                       calendarCalendarViewExceptionOccurrenceExtensionClient,
		CalendarCalendarViewExceptionOccurrenceInstance:                        calendarCalendarViewExceptionOccurrenceInstanceClient,
		CalendarCalendarViewExceptionOccurrenceInstanceAttachment:              calendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		CalendarCalendarViewExceptionOccurrenceInstanceCalendar:                calendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		CalendarCalendarViewExceptionOccurrenceInstanceExtension:               calendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		CalendarCalendarViewExtension:                                          calendarCalendarViewExtensionClient,
		CalendarCalendarViewInstance:                                           calendarCalendarViewInstanceClient,
		CalendarCalendarViewInstanceAttachment:                                 calendarCalendarViewInstanceAttachmentClient,
		CalendarCalendarViewInstanceCalendar:                                   calendarCalendarViewInstanceCalendarClient,
		CalendarCalendarViewInstanceExceptionOccurrence:                        calendarCalendarViewInstanceExceptionOccurrenceClient,
		CalendarCalendarViewInstanceExceptionOccurrenceAttachment:              calendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		CalendarCalendarViewInstanceExceptionOccurrenceCalendar:                calendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		CalendarCalendarViewInstanceExceptionOccurrenceExtension:               calendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		CalendarCalendarViewInstanceExtension:                                  calendarCalendarViewInstanceExtensionClient,
		CalendarEvent:                                                          calendarEventClient,
		CalendarEventAttachment:                                                calendarEventAttachmentClient,
		CalendarEventCalendar:                                                  calendarEventCalendarClient,
		CalendarEventExceptionOccurrence:                                       calendarEventExceptionOccurrenceClient,
		CalendarEventExceptionOccurrenceAttachment:                             calendarEventExceptionOccurrenceAttachmentClient,
		CalendarEventExceptionOccurrenceCalendar:                               calendarEventExceptionOccurrenceCalendarClient,
		CalendarEventExceptionOccurrenceExtension:                              calendarEventExceptionOccurrenceExtensionClient,
		CalendarEventExceptionOccurrenceInstance:                               calendarEventExceptionOccurrenceInstanceClient,
		CalendarEventExceptionOccurrenceInstanceAttachment:                     calendarEventExceptionOccurrenceInstanceAttachmentClient,
		CalendarEventExceptionOccurrenceInstanceCalendar:                       calendarEventExceptionOccurrenceInstanceCalendarClient,
		CalendarEventExceptionOccurrenceInstanceExtension:                      calendarEventExceptionOccurrenceInstanceExtensionClient,
		CalendarEventExtension:                                                 calendarEventExtensionClient,
		CalendarEventInstance:                                                  calendarEventInstanceClient,
		CalendarEventInstanceAttachment:                                        calendarEventInstanceAttachmentClient,
		CalendarEventInstanceCalendar:                                          calendarEventInstanceCalendarClient,
		CalendarEventInstanceExceptionOccurrence:                               calendarEventInstanceExceptionOccurrenceClient,
		CalendarEventInstanceExceptionOccurrenceAttachment:                     calendarEventInstanceExceptionOccurrenceAttachmentClient,
		CalendarEventInstanceExceptionOccurrenceCalendar:                       calendarEventInstanceExceptionOccurrenceCalendarClient,
		CalendarEventInstanceExceptionOccurrenceExtension:                      calendarEventInstanceExceptionOccurrenceExtensionClient,
		CalendarEventInstanceExtension:                                         calendarEventInstanceExtensionClient,
		CalendarGroup:                                                          calendarGroupClient,
		CalendarGroupCalendar:                                                  calendarGroupCalendarClient,
		CalendarGroupCalendarCalendarPermission:                                calendarGroupCalendarCalendarPermissionClient,
		CalendarGroupCalendarCalendarView:                                      calendarGroupCalendarCalendarViewClient,
		CalendarGroupCalendarCalendarViewAttachment:                            calendarGroupCalendarCalendarViewAttachmentClient,
		CalendarGroupCalendarCalendarViewCalendar:                              calendarGroupCalendarCalendarViewCalendarClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrence:                   calendarGroupCalendarCalendarViewExceptionOccurrenceClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment:         calendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar:           calendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrenceExtension:          calendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrenceInstance:           calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment: calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar:   calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension:  calendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		CalendarGroupCalendarCalendarViewExtension:                             calendarGroupCalendarCalendarViewExtensionClient,
		CalendarGroupCalendarCalendarViewInstance:                              calendarGroupCalendarCalendarViewInstanceClient,
		CalendarGroupCalendarCalendarViewInstanceAttachment:                    calendarGroupCalendarCalendarViewInstanceAttachmentClient,
		CalendarGroupCalendarCalendarViewInstanceCalendar:                      calendarGroupCalendarCalendarViewInstanceCalendarClient,
		CalendarGroupCalendarCalendarViewInstanceExceptionOccurrence:           calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient,
		CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment: calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar:   calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension:  calendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		CalendarGroupCalendarCalendarViewInstanceExtension:                     calendarGroupCalendarCalendarViewInstanceExtensionClient,
		CalendarGroupCalendarEvent:                                             calendarGroupCalendarEventClient,
		CalendarGroupCalendarEventAttachment:                                   calendarGroupCalendarEventAttachmentClient,
		CalendarGroupCalendarEventCalendar:                                     calendarGroupCalendarEventCalendarClient,
		CalendarGroupCalendarEventExceptionOccurrence:                          calendarGroupCalendarEventExceptionOccurrenceClient,
		CalendarGroupCalendarEventExceptionOccurrenceAttachment:                calendarGroupCalendarEventExceptionOccurrenceAttachmentClient,
		CalendarGroupCalendarEventExceptionOccurrenceCalendar:                  calendarGroupCalendarEventExceptionOccurrenceCalendarClient,
		CalendarGroupCalendarEventExceptionOccurrenceExtension:                 calendarGroupCalendarEventExceptionOccurrenceExtensionClient,
		CalendarGroupCalendarEventExceptionOccurrenceInstance:                  calendarGroupCalendarEventExceptionOccurrenceInstanceClient,
		CalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment:        calendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient,
		CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar:          calendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient,
		CalendarGroupCalendarEventExceptionOccurrenceInstanceExtension:         calendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient,
		CalendarGroupCalendarEventExtension:                                    calendarGroupCalendarEventExtensionClient,
		CalendarGroupCalendarEventInstance:                                     calendarGroupCalendarEventInstanceClient,
		CalendarGroupCalendarEventInstanceAttachment:                           calendarGroupCalendarEventInstanceAttachmentClient,
		CalendarGroupCalendarEventInstanceCalendar:                             calendarGroupCalendarEventInstanceCalendarClient,
		CalendarGroupCalendarEventInstanceExceptionOccurrence:                  calendarGroupCalendarEventInstanceExceptionOccurrenceClient,
		CalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment:        calendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient,
		CalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar:          calendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient,
		CalendarGroupCalendarEventInstanceExceptionOccurrenceExtension:         calendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient,
		CalendarGroupCalendarEventInstanceExtension:                            calendarGroupCalendarEventInstanceExtensionClient,
		CalendarView:                                       calendarViewClient,
		CalendarViewAttachment:                             calendarViewAttachmentClient,
		CalendarViewCalendar:                               calendarViewCalendarClient,
		CalendarViewExceptionOccurrence:                    calendarViewExceptionOccurrenceClient,
		CalendarViewExceptionOccurrenceAttachment:          calendarViewExceptionOccurrenceAttachmentClient,
		CalendarViewExceptionOccurrenceCalendar:            calendarViewExceptionOccurrenceCalendarClient,
		CalendarViewExceptionOccurrenceExtension:           calendarViewExceptionOccurrenceExtensionClient,
		CalendarViewExceptionOccurrenceInstance:            calendarViewExceptionOccurrenceInstanceClient,
		CalendarViewExceptionOccurrenceInstanceAttachment:  calendarViewExceptionOccurrenceInstanceAttachmentClient,
		CalendarViewExceptionOccurrenceInstanceCalendar:    calendarViewExceptionOccurrenceInstanceCalendarClient,
		CalendarViewExceptionOccurrenceInstanceExtension:   calendarViewExceptionOccurrenceInstanceExtensionClient,
		CalendarViewExtension:                              calendarViewExtensionClient,
		CalendarViewInstance:                               calendarViewInstanceClient,
		CalendarViewInstanceAttachment:                     calendarViewInstanceAttachmentClient,
		CalendarViewInstanceCalendar:                       calendarViewInstanceCalendarClient,
		CalendarViewInstanceExceptionOccurrence:            calendarViewInstanceExceptionOccurrenceClient,
		CalendarViewInstanceExceptionOccurrenceAttachment:  calendarViewInstanceExceptionOccurrenceAttachmentClient,
		CalendarViewInstanceExceptionOccurrenceCalendar:    calendarViewInstanceExceptionOccurrenceCalendarClient,
		CalendarViewInstanceExceptionOccurrenceExtension:   calendarViewInstanceExceptionOccurrenceExtensionClient,
		CalendarViewInstanceExtension:                      calendarViewInstanceExtensionClient,
		Chat:                                               chatClient,
		ChatInstalledApp:                                   chatInstalledAppClient,
		ChatInstalledAppTeamsApp:                           chatInstalledAppTeamsAppClient,
		ChatInstalledAppTeamsAppDefinition:                 chatInstalledAppTeamsAppDefinitionClient,
		ChatLastMessagePreview:                             chatLastMessagePreviewClient,
		ChatMember:                                         chatMemberClient,
		ChatMessage:                                        chatMessageClient,
		ChatMessageHostedContent:                           chatMessageHostedContentClient,
		ChatMessageReply:                                   chatMessageReplyClient,
		ChatMessageReplyHostedContent:                      chatMessageReplyHostedContentClient,
		ChatOperation:                                      chatOperationClient,
		ChatPermissionGrant:                                chatPermissionGrantClient,
		ChatPinnedMessage:                                  chatPinnedMessageClient,
		ChatPinnedMessageMessage:                           chatPinnedMessageMessageClient,
		ChatTab:                                            chatTabClient,
		ChatTabTeamsApp:                                    chatTabTeamsAppClient,
		CloudPC:                                            cloudPCClient,
		Contact:                                            contactClient,
		ContactExtension:                                   contactExtensionClient,
		ContactFolder:                                      contactFolderClient,
		ContactFolderChildFolder:                           contactFolderChildFolderClient,
		ContactFolderChildFolderContact:                    contactFolderChildFolderContactClient,
		ContactFolderChildFolderContactExtension:           contactFolderChildFolderContactExtensionClient,
		ContactFolderChildFolderContactPhoto:               contactFolderChildFolderContactPhotoClient,
		ContactFolderContact:                               contactFolderContactClient,
		ContactFolderContactExtension:                      contactFolderContactExtensionClient,
		ContactFolderContactPhoto:                          contactFolderContactPhotoClient,
		ContactPhoto:                                       contactPhotoClient,
		CreatedObject:                                      createdObjectClient,
		Device:                                             deviceClient,
		DeviceCommand:                                      deviceCommandClient,
		DeviceCommandResponsepayload:                       deviceCommandResponsepayloadClient,
		DeviceEnrollmentConfiguration:                      deviceEnrollmentConfigurationClient,
		DeviceEnrollmentConfigurationAssignment:            deviceEnrollmentConfigurationAssignmentClient,
		DeviceExtension:                                    deviceExtensionClient,
		DeviceManagementTroubleshootingEvent:               deviceManagementTroubleshootingEventClient,
		DeviceMemberOf:                                     deviceMemberOfClient,
		DeviceRegisteredOwner:                              deviceRegisteredOwnerClient,
		DeviceRegisteredUser:                               deviceRegisteredUserClient,
		DeviceTransitiveMemberOf:                           deviceTransitiveMemberOfClient,
		DeviceUsageRight:                                   deviceUsageRightClient,
		DirectReport:                                       directReportClient,
		Drive:                                              driveClient,
		EmployeeExperience:                                 employeeExperienceClient,
		EmployeeExperienceLearningCourseActivity:           employeeExperienceLearningCourseActivityClient,
		Event:                                              eventClient,
		EventAttachment:                                    eventAttachmentClient,
		EventCalendar:                                      eventCalendarClient,
		EventExceptionOccurrence:                           eventExceptionOccurrenceClient,
		EventExceptionOccurrenceAttachment:                 eventExceptionOccurrenceAttachmentClient,
		EventExceptionOccurrenceCalendar:                   eventExceptionOccurrenceCalendarClient,
		EventExceptionOccurrenceExtension:                  eventExceptionOccurrenceExtensionClient,
		EventExceptionOccurrenceInstance:                   eventExceptionOccurrenceInstanceClient,
		EventExceptionOccurrenceInstanceAttachment:         eventExceptionOccurrenceInstanceAttachmentClient,
		EventExceptionOccurrenceInstanceCalendar:           eventExceptionOccurrenceInstanceCalendarClient,
		EventExceptionOccurrenceInstanceExtension:          eventExceptionOccurrenceInstanceExtensionClient,
		EventExtension:                                     eventExtensionClient,
		EventInstance:                                      eventInstanceClient,
		EventInstanceAttachment:                            eventInstanceAttachmentClient,
		EventInstanceCalendar:                              eventInstanceCalendarClient,
		EventInstanceExceptionOccurrence:                   eventInstanceExceptionOccurrenceClient,
		EventInstanceExceptionOccurrenceAttachment:         eventInstanceExceptionOccurrenceAttachmentClient,
		EventInstanceExceptionOccurrenceCalendar:           eventInstanceExceptionOccurrenceCalendarClient,
		EventInstanceExceptionOccurrenceExtension:          eventInstanceExceptionOccurrenceExtensionClient,
		EventInstanceExtension:                             eventInstanceExtensionClient,
		Extension:                                          extensionClient,
		FollowedSite:                                       followedSiteClient,
		InferenceClassification:                            inferenceClassificationClient,
		InferenceClassificationOverride:                    inferenceClassificationOverrideClient,
		InformationProtection:                              informationProtectionClient,
		InformationProtectionBitlocker:                     informationProtectionBitlockerClient,
		InformationProtectionBitlockerRecoveryKey:          informationProtectionBitlockerRecoveryKeyClient,
		InformationProtectionDataLossPreventionPolicy:      informationProtectionDataLossPreventionPolicyClient,
		InformationProtectionPolicy:                        informationProtectionPolicyClient,
		InformationProtectionPolicyLabel:                   informationProtectionPolicyLabelClient,
		InformationProtectionSensitivityLabel:              informationProtectionSensitivityLabelClient,
		InformationProtectionSensitivityLabelSublabel:      informationProtectionSensitivityLabelSublabelClient,
		InformationProtectionSensitivityPolicySetting:      informationProtectionSensitivityPolicySettingClient,
		InformationProtectionThreatAssessmentRequest:       informationProtectionThreatAssessmentRequestClient,
		InformationProtectionThreatAssessmentRequestResult: informationProtectionThreatAssessmentRequestResultClient,
		Insight:                                insightClient,
		InsightShared:                          insightSharedClient,
		InsightSharedLastSharedMethod:          insightSharedLastSharedMethodClient,
		InsightSharedResource:                  insightSharedResourceClient,
		InsightTrending:                        insightTrendingClient,
		InsightTrendingResource:                insightTrendingResourceClient,
		InsightUsed:                            insightUsedClient,
		InsightUsedResource:                    insightUsedResourceClient,
		JoinedGroup:                            joinedGroupClient,
		JoinedTeam:                             joinedTeamClient,
		LicenseDetail:                          licenseDetailClient,
		MailFolder:                             mailFolderClient,
		MailFolderChildFolder:                  mailFolderChildFolderClient,
		MailFolderChildFolderMessage:           mailFolderChildFolderMessageClient,
		MailFolderChildFolderMessageAttachment: mailFolderChildFolderMessageAttachmentClient,
		MailFolderChildFolderMessageExtension:  mailFolderChildFolderMessageExtensionClient,
		MailFolderChildFolderMessageMention:    mailFolderChildFolderMessageMentionClient,
		MailFolderChildFolderMessageRule:       mailFolderChildFolderMessageRuleClient,
		MailFolderChildFolderUserConfiguration: mailFolderChildFolderUserConfigurationClient,
		MailFolderMessage:                      mailFolderMessageClient,
		MailFolderMessageAttachment:            mailFolderMessageAttachmentClient,
		MailFolderMessageExtension:             mailFolderMessageExtensionClient,
		MailFolderMessageMention:               mailFolderMessageMentionClient,
		MailFolderMessageRule:                  mailFolderMessageRuleClient,
		MailFolderUserConfiguration:            mailFolderUserConfigurationClient,
		MailboxSetting:                         mailboxSettingClient,
		ManagedAppRegistration:                 managedAppRegistrationClient,
		ManagedDevice:                          managedDeviceClient,
		ManagedDeviceAssignmentFilterEvaluationStatusDetail:                      managedDeviceAssignmentFilterEvaluationStatusDetailClient,
		ManagedDeviceDetectedApp:                                                 managedDeviceDetectedAppClient,
		ManagedDeviceDeviceCategory:                                              managedDeviceDeviceCategoryClient,
		ManagedDeviceDeviceCompliancePolicyState:                                 managedDeviceDeviceCompliancePolicyStateClient,
		ManagedDeviceDeviceConfigurationState:                                    managedDeviceDeviceConfigurationStateClient,
		ManagedDeviceDeviceHealthScriptState:                                     managedDeviceDeviceHealthScriptStateClient,
		ManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId: managedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient,
		ManagedDeviceLogCollectionRequest:                                        managedDeviceLogCollectionRequestClient,
		ManagedDeviceManagedDeviceMobileAppConfigurationState:                    managedDeviceManagedDeviceMobileAppConfigurationStateClient,
		ManagedDeviceSecurityBaselineState:                                       managedDeviceSecurityBaselineStateClient,
		ManagedDeviceSecurityBaselineStateSettingState:                           managedDeviceSecurityBaselineStateSettingStateClient,
		ManagedDeviceUser:                                                        managedDeviceUserClient,
		ManagedDeviceWindowsProtectionState:                                      managedDeviceWindowsProtectionStateClient,
		ManagedDeviceWindowsProtectionStateDetectedMalwareState:                  managedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		Manager:                       managerClient,
		MemberOf:                      memberOfClient,
		Message:                       messageClient,
		MessageAttachment:             messageAttachmentClient,
		MessageExtension:              messageExtensionClient,
		MessageMention:                messageMentionClient,
		MobileAppIntentAndState:       mobileAppIntentAndStateClient,
		MobileAppTroubleshootingEvent: mobileAppTroubleshootingEventClient,
		MobileAppTroubleshootingEventAppLogCollectionRequest: mobileAppTroubleshootingEventAppLogCollectionRequestClient,
		Notification:                                                      notificationClient,
		Oauth2PermissionGrant:                                             oauth2PermissionGrantClient,
		OnlineMeeting:                                                     onlineMeetingClient,
		OnlineMeetingAlternativeRecording:                                 onlineMeetingAlternativeRecordingClient,
		OnlineMeetingAttendanceReport:                                     onlineMeetingAttendanceReportClient,
		OnlineMeetingAttendanceReportAttendanceRecord:                     onlineMeetingAttendanceReportAttendanceRecordClient,
		OnlineMeetingAttendeeReport:                                       onlineMeetingAttendeeReportClient,
		OnlineMeetingBroadcastRecording:                                   onlineMeetingBroadcastRecordingClient,
		OnlineMeetingMeetingAttendanceReport:                              onlineMeetingMeetingAttendanceReportClient,
		OnlineMeetingMeetingAttendanceReportAttendanceRecord:              onlineMeetingMeetingAttendanceReportAttendanceRecordClient,
		OnlineMeetingRecording:                                            onlineMeetingRecordingClient,
		OnlineMeetingRecordingContent:                                     onlineMeetingRecordingContentClient,
		OnlineMeetingRegistration:                                         onlineMeetingRegistrationClient,
		OnlineMeetingRegistrationCustomQuestion:                           onlineMeetingRegistrationCustomQuestionClient,
		OnlineMeetingRegistrationRegistrant:                               onlineMeetingRegistrationRegistrantClient,
		OnlineMeetingTranscript:                                           onlineMeetingTranscriptClient,
		OnlineMeetingTranscriptContent:                                    onlineMeetingTranscriptContentClient,
		OnlineMeetingTranscriptMetadataContent:                            onlineMeetingTranscriptMetadataContentClient,
		Outlook:                                                           outlookClient,
		OutlookMasterCategory:                                             outlookMasterCategoryClient,
		OutlookTask:                                                       outlookTaskClient,
		OutlookTaskAttachment:                                             outlookTaskAttachmentClient,
		OutlookTaskFolder:                                                 outlookTaskFolderClient,
		OutlookTaskFolderTask:                                             outlookTaskFolderTaskClient,
		OutlookTaskFolderTaskAttachment:                                   outlookTaskFolderTaskAttachmentClient,
		OutlookTaskGroup:                                                  outlookTaskGroupClient,
		OutlookTaskGroupTaskFolder:                                        outlookTaskGroupTaskFolderClient,
		OutlookTaskGroupTaskFolderTask:                                    outlookTaskGroupTaskFolderTaskClient,
		OutlookTaskGroupTaskFolderTaskAttachment:                          outlookTaskGroupTaskFolderTaskAttachmentClient,
		OwnedDevice:                                                       ownedDeviceClient,
		OwnedObject:                                                       ownedObjectClient,
		PendingAccessReviewInstance:                                       pendingAccessReviewInstanceClient,
		PendingAccessReviewInstanceContactedReviewer:                      pendingAccessReviewInstanceContactedReviewerClient,
		PendingAccessReviewInstanceDecision:                               pendingAccessReviewInstanceDecisionClient,
		PendingAccessReviewInstanceDecisionInsight:                        pendingAccessReviewInstanceDecisionInsightClient,
		PendingAccessReviewInstanceDecisionInstance:                       pendingAccessReviewInstanceDecisionInstanceClient,
		PendingAccessReviewInstanceDecisionInstanceContactedReviewer:      pendingAccessReviewInstanceDecisionInstanceContactedReviewerClient,
		PendingAccessReviewInstanceDecisionInstanceDefinition:             pendingAccessReviewInstanceDecisionInstanceDefinitionClient,
		PendingAccessReviewInstanceDecisionInstanceStage:                  pendingAccessReviewInstanceDecisionInstanceStageClient,
		PendingAccessReviewInstanceDefinition:                             pendingAccessReviewInstanceDefinitionClient,
		PendingAccessReviewInstanceStage:                                  pendingAccessReviewInstanceStageClient,
		PendingAccessReviewInstanceStageDecision:                          pendingAccessReviewInstanceStageDecisionClient,
		PendingAccessReviewInstanceStageDecisionInsight:                   pendingAccessReviewInstanceStageDecisionInsightClient,
		PendingAccessReviewInstanceStageDecisionInstance:                  pendingAccessReviewInstanceStageDecisionInstanceClient,
		PendingAccessReviewInstanceStageDecisionInstanceContactedReviewer: pendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient,
		PendingAccessReviewInstanceStageDecisionInstanceDefinition:        pendingAccessReviewInstanceStageDecisionInstanceDefinitionClient,
		People:                peopleClient,
		PermissionGrant:       permissionGrantClient,
		Photo:                 photoClient,
		Planner:               plannerClient,
		PlannerAll:            plannerAllClient,
		PlannerFavoritePlan:   plannerFavoritePlanClient,
		PlannerPlan:           plannerPlanClient,
		PlannerPlanBucket:     plannerPlanBucketClient,
		PlannerPlanBucketTask: plannerPlanBucketTaskClient,
		PlannerPlanBucketTaskAssignedToTaskBoardFormat: plannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		PlannerPlanBucketTaskBucketTaskBoardFormat:     plannerPlanBucketTaskBucketTaskBoardFormatClient,
		PlannerPlanBucketTaskDetail:                    plannerPlanBucketTaskDetailClient,
		PlannerPlanBucketTaskProgressTaskBoardFormat:   plannerPlanBucketTaskProgressTaskBoardFormatClient,
		PlannerPlanDetail:                              plannerPlanDetailClient,
		PlannerPlanTask:                                plannerPlanTaskClient,
		PlannerPlanTaskAssignedToTaskBoardFormat:       plannerPlanTaskAssignedToTaskBoardFormatClient,
		PlannerPlanTaskBucketTaskBoardFormat:           plannerPlanTaskBucketTaskBoardFormatClient,
		PlannerPlanTaskDetail:                          plannerPlanTaskDetailClient,
		PlannerPlanTaskProgressTaskBoardFormat:         plannerPlanTaskProgressTaskBoardFormatClient,
		PlannerRecentPlan:                              plannerRecentPlanClient,
		PlannerRosterPlan:                              plannerRosterPlanClient,
		PlannerTask:                                    plannerTaskClient,
		PlannerTaskAssignedToTaskBoardFormat:           plannerTaskAssignedToTaskBoardFormatClient,
		PlannerTaskBucketTaskBoardFormat:               plannerTaskBucketTaskBoardFormatClient,
		PlannerTaskDetail:                              plannerTaskDetailClient,
		PlannerTaskProgressTaskBoardFormat:             plannerTaskProgressTaskBoardFormatClient,
		Presence:                                       presenceClient,
		Profile:                                        profileClient,
		ProfileAccount:                                 profileAccountClient,
		ProfileAddress:                                 profileAddressClient,
		ProfileAnniversary:                             profileAnniversaryClient,
		ProfileAward:                                   profileAwardClient,
		ProfileCertification:                           profileCertificationClient,
		ProfileEducationalActivity:                     profileEducationalActivityClient,
		ProfileEmail:                                   profileEmailClient,
		ProfileInterest:                                profileInterestClient,
		ProfileLanguage:                                profileLanguageClient,
		ProfileName:                                    profileNameClient,
		ProfileNote:                                    profileNoteClient,
		ProfilePatent:                                  profilePatentClient,
		ProfilePhone:                                   profilePhoneClient,
		ProfilePosition:                                profilePositionClient,
		ProfileProject:                                 profileProjectClient,
		ProfilePublication:                             profilePublicationClient,
		ProfileSkill:                                   profileSkillClient,
		ProfileWebAccount:                              profileWebAccountClient,
		ProfileWebsite:                                 profileWebsiteClient,
		RegisteredDevice:                               registeredDeviceClient,
		ScopedRoleMemberOf:                             scopedRoleMemberOfClient,
		Security:                                       securityClient,
		SecurityInformationProtection:                  securityInformationProtectionClient,
		SecurityInformationProtectionLabelPolicySetting:     securityInformationProtectionLabelPolicySettingClient,
		SecurityInformationProtectionSensitivityLabel:       securityInformationProtectionSensitivityLabelClient,
		SecurityInformationProtectionSensitivityLabelParent: securityInformationProtectionSensitivityLabelParentClient,
		ServiceProvisioningError:                            serviceProvisioningErrorClient,
		Setting:                                             settingClient,
		SettingContactMergeSuggestion:                       settingContactMergeSuggestionClient,
		SettingItemInsight:                                  settingItemInsightClient,
		SettingRegionalAndLanguageSetting:                   settingRegionalAndLanguageSettingClient,
		SettingShiftPreference:                              settingShiftPreferenceClient,
		Sponsor:                                             sponsorClient,
		Teamwork:                                            teamworkClient,
		TeamworkAssociatedTeam:                              teamworkAssociatedTeamClient,
		TeamworkAssociatedTeamTeam:                          teamworkAssociatedTeamTeamClient,
		TeamworkInstalledApp:                                teamworkInstalledAppClient,
		TeamworkInstalledAppChat:                            teamworkInstalledAppChatClient,
		TeamworkInstalledAppTeamsApp:                        teamworkInstalledAppTeamsAppClient,
		TeamworkInstalledAppTeamsAppDefinition:              teamworkInstalledAppTeamsAppDefinitionClient,
		Todo:                                                todoClient,
		TodoList:                                            todoListClient,
		TodoListExtension:                                   todoListExtensionClient,
		TodoListTask:                                        todoListTaskClient,
		TodoListTaskAttachment:                              todoListTaskAttachmentClient,
		TodoListTaskAttachmentSession:                       todoListTaskAttachmentSessionClient,
		TodoListTaskAttachmentSessionContent:                todoListTaskAttachmentSessionContentClient,
		TodoListTaskChecklistItem:                           todoListTaskChecklistItemClient,
		TodoListTaskExtension:                               todoListTaskExtensionClient,
		TodoListTaskLinkedResource:                          todoListTaskLinkedResourceClient,
		TransitiveMemberOf:                                  transitiveMemberOfClient,
		TransitiveReport:                                    transitiveReportClient,
		UsageRight:                                          usageRightClient,
		User:                                                userClient,
		VirtualEvent:                                        virtualEventClient,
		VirtualEventWebinar:                                 virtualEventWebinarClient,
		WindowsInformationProtectionDeviceRegistration:      windowsInformationProtectionDeviceRegistrationClient,
	}, nil
}
