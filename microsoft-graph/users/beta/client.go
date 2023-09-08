package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/user"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useractivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useractivityhistoryitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useractivityhistoryitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useragreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useranalytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useranalyticactivitystatistic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userappconsentrequestsforapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userappconsentrequestsforapprovaluserconsentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userappconsentrequestsforapprovaluserconsentrequestapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userappconsentrequestsforapprovaluserconsentrequestapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userapproleassignedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationemailmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationfido2method"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationpasswordlessmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationpasswordlessmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationpasswordmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationphonemethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationsoftwareoathmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationtemporaryaccesspassmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationwindowshelloforbusinessmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userauthenticationwindowshelloforbusinessmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendargroupcalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatlastmessagepreview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatpermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatpinnedmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchatpinnedmessagemessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchattab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userchattabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercloudpc"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfolderchildfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfolderchildfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfolderchildfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usercreatedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdevicecommand"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdevicecommandresponsepayload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdeviceenrollmentconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdeviceenrollmentconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdeviceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdevicemanagementtroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdevicememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdeviceregisteredowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdeviceregistereduser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdevicetransitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdeviceusageright"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdirectreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useremployeeexperience"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useremployeeexperiencelearningcourseactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usereventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userfollowedsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinferenceclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinferenceclassificationoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionbitlocker"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionbitlockerrecoverykey"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectiondatalosspreventionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionpolicylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionsensitivitylabelsublabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionsensitivitypolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionthreatassessmentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinformationprotectionthreatassessmentrequestresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsightshared"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsightsharedlastsharedmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsightsharedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsighttrending"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsighttrendingresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsightused"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userinsightusedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userjoinedgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userjoinedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userlicensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderchildfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderchildfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderchildfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderchildfoldermessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderchildfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderchildfolderuserconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfoldermessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermailfolderuserconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanagedappregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddeviceassignmentfilterevaluationstatusdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicedetectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicedevicehealthscriptstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicemanageddevicemobileappconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicesecuritybaselinestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicesecuritybaselinestatesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermanager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermemberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermobileappintentandstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermobileapptroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usermobileapptroubleshootingeventapplogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usernotification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotenotebooksectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenoteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotepagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotepageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotepageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenoteresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenoteresourcecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronenotesectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeeting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingalternativerecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingattendeereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingbroadcastrecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingmeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingmeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingrecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingrecordingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingregistrationcustomquestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingregistrationregistrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingtranscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingtranscriptcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useronlinemeetingtranscriptmetadatacontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlookmastercategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskfoldertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskfoldertaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskgrouptaskfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskgrouptaskfoldertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/useroutlooktaskgrouptaskfoldertaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userowneddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecisioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecisioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecisioninstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecisioninstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecisioninstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedecisioninstancestagedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestagedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestagedecisioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestagedecisioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestagedecisioninstancedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestagedecisioninstancedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpendingaccessreviewinstancestagedecisioninstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpeople"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplanner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerall"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerfavoriteplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerrecentplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannerrosterplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannertaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannertaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannertaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userplannertaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userpresence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileaccount"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileaddress"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileanniversary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileaward"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilecertification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileeducationalactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileemail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileinterest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilelanguage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilename"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilepatent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilephone"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileproject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilepublication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofileskill"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilewebaccount"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userprofilewebsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userregistereddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userscopedrolememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersecurity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersecurityinformationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersecurityinformationprotectionlabelpolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersecurityinformationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersecurityinformationprotectionsensitivitylabelparent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersettingcontactmergesuggestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersettingiteminsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersettingregionalandlanguagesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersettingshiftpreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usersponsor"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userteamwork"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userteamworkassociatedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userteamworkassociatedteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userteamworkinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userteamworkinstalledappchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userteamworkinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userteamworkinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolistextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolisttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolisttaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolisttaskattachmentsession"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolisttaskattachmentsessioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolisttaskchecklistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolisttaskextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertodolisttasklinkedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertransitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/usertransitivereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userusageright"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/userwindowsinformationprotectiondeviceregistration"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	User                                                                         *user.UserClient
	UserActivity                                                                 *useractivity.UserActivityClient
	UserActivityHistoryItem                                                      *useractivityhistoryitem.UserActivityHistoryItemClient
	UserActivityHistoryItemActivity                                              *useractivityhistoryitemactivity.UserActivityHistoryItemActivityClient
	UserAgreementAcceptance                                                      *useragreementacceptance.UserAgreementAcceptanceClient
	UserAnalytic                                                                 *useranalytic.UserAnalyticClient
	UserAnalyticActivityStatistic                                                *useranalyticactivitystatistic.UserAnalyticActivityStatisticClient
	UserAppConsentRequestsForApproval                                            *userappconsentrequestsforapproval.UserAppConsentRequestsForApprovalClient
	UserAppConsentRequestsForApprovalUserConsentRequest                          *userappconsentrequestsforapprovaluserconsentrequest.UserAppConsentRequestsForApprovalUserConsentRequestClient
	UserAppConsentRequestsForApprovalUserConsentRequestApproval                  *userappconsentrequestsforapprovaluserconsentrequestapproval.UserAppConsentRequestsForApprovalUserConsentRequestApprovalClient
	UserAppConsentRequestsForApprovalUserConsentRequestApprovalStep              *userappconsentrequestsforapprovaluserconsentrequestapprovalstep.UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient
	UserAppRoleAssignedResource                                                  *userapproleassignedresource.UserAppRoleAssignedResourceClient
	UserAppRoleAssignment                                                        *userapproleassignment.UserAppRoleAssignmentClient
	UserApproval                                                                 *userapproval.UserApprovalClient
	UserApprovalStep                                                             *userapprovalstep.UserApprovalStepClient
	UserAuthentication                                                           *userauthentication.UserAuthenticationClient
	UserAuthenticationEmailMethod                                                *userauthenticationemailmethod.UserAuthenticationEmailMethodClient
	UserAuthenticationFido2Method                                                *userauthenticationfido2method.UserAuthenticationFido2MethodClient
	UserAuthenticationMethod                                                     *userauthenticationmethod.UserAuthenticationMethodClient
	UserAuthenticationMicrosoftAuthenticatorMethod                               *userauthenticationmicrosoftauthenticatormethod.UserAuthenticationMicrosoftAuthenticatorMethodClient
	UserAuthenticationMicrosoftAuthenticatorMethodDevice                         *userauthenticationmicrosoftauthenticatormethoddevice.UserAuthenticationMicrosoftAuthenticatorMethodDeviceClient
	UserAuthenticationOperation                                                  *userauthenticationoperation.UserAuthenticationOperationClient
	UserAuthenticationPasswordMethod                                             *userauthenticationpasswordmethod.UserAuthenticationPasswordMethodClient
	UserAuthenticationPasswordlessMicrosoftAuthenticatorMethod                   *userauthenticationpasswordlessmicrosoftauthenticatormethod.UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient
	UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice             *userauthenticationpasswordlessmicrosoftauthenticatormethoddevice.UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient
	UserAuthenticationPhoneMethod                                                *userauthenticationphonemethod.UserAuthenticationPhoneMethodClient
	UserAuthenticationSoftwareOathMethod                                         *userauthenticationsoftwareoathmethod.UserAuthenticationSoftwareOathMethodClient
	UserAuthenticationTemporaryAccessPassMethod                                  *userauthenticationtemporaryaccesspassmethod.UserAuthenticationTemporaryAccessPassMethodClient
	UserAuthenticationWindowsHelloForBusinessMethod                              *userauthenticationwindowshelloforbusinessmethod.UserAuthenticationWindowsHelloForBusinessMethodClient
	UserAuthenticationWindowsHelloForBusinessMethodDevice                        *userauthenticationwindowshelloforbusinessmethoddevice.UserAuthenticationWindowsHelloForBusinessMethodDeviceClient
	UserCalendar                                                                 *usercalendar.UserCalendarClient
	UserCalendarCalendarPermission                                               *usercalendarcalendarpermission.UserCalendarCalendarPermissionClient
	UserCalendarCalendarView                                                     *usercalendarcalendarview.UserCalendarCalendarViewClient
	UserCalendarCalendarViewAttachment                                           *usercalendarcalendarviewattachment.UserCalendarCalendarViewAttachmentClient
	UserCalendarCalendarViewCalendar                                             *usercalendarcalendarviewcalendar.UserCalendarCalendarViewCalendarClient
	UserCalendarCalendarViewExceptionOccurrence                                  *usercalendarcalendarviewexceptionoccurrence.UserCalendarCalendarViewExceptionOccurrenceClient
	UserCalendarCalendarViewExceptionOccurrenceAttachment                        *usercalendarcalendarviewexceptionoccurrenceattachment.UserCalendarCalendarViewExceptionOccurrenceAttachmentClient
	UserCalendarCalendarViewExceptionOccurrenceCalendar                          *usercalendarcalendarviewexceptionoccurrencecalendar.UserCalendarCalendarViewExceptionOccurrenceCalendarClient
	UserCalendarCalendarViewExceptionOccurrenceExtension                         *usercalendarcalendarviewexceptionoccurrenceextension.UserCalendarCalendarViewExceptionOccurrenceExtensionClient
	UserCalendarCalendarViewExceptionOccurrenceInstance                          *usercalendarcalendarviewexceptionoccurrenceinstance.UserCalendarCalendarViewExceptionOccurrenceInstanceClient
	UserCalendarCalendarViewExceptionOccurrenceInstanceAttachment                *usercalendarcalendarviewexceptionoccurrenceinstanceattachment.UserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	UserCalendarCalendarViewExceptionOccurrenceInstanceCalendar                  *usercalendarcalendarviewexceptionoccurrenceinstancecalendar.UserCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	UserCalendarCalendarViewExceptionOccurrenceInstanceExtension                 *usercalendarcalendarviewexceptionoccurrenceinstanceextension.UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	UserCalendarCalendarViewExtension                                            *usercalendarcalendarviewextension.UserCalendarCalendarViewExtensionClient
	UserCalendarCalendarViewInstance                                             *usercalendarcalendarviewinstance.UserCalendarCalendarViewInstanceClient
	UserCalendarCalendarViewInstanceAttachment                                   *usercalendarcalendarviewinstanceattachment.UserCalendarCalendarViewInstanceAttachmentClient
	UserCalendarCalendarViewInstanceCalendar                                     *usercalendarcalendarviewinstancecalendar.UserCalendarCalendarViewInstanceCalendarClient
	UserCalendarCalendarViewInstanceExceptionOccurrence                          *usercalendarcalendarviewinstanceexceptionoccurrence.UserCalendarCalendarViewInstanceExceptionOccurrenceClient
	UserCalendarCalendarViewInstanceExceptionOccurrenceAttachment                *usercalendarcalendarviewinstanceexceptionoccurrenceattachment.UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	UserCalendarCalendarViewInstanceExceptionOccurrenceCalendar                  *usercalendarcalendarviewinstanceexceptionoccurrencecalendar.UserCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	UserCalendarCalendarViewInstanceExceptionOccurrenceExtension                 *usercalendarcalendarviewinstanceexceptionoccurrenceextension.UserCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	UserCalendarCalendarViewInstanceExtension                                    *usercalendarcalendarviewinstanceextension.UserCalendarCalendarViewInstanceExtensionClient
	UserCalendarEvent                                                            *usercalendarevent.UserCalendarEventClient
	UserCalendarEventAttachment                                                  *usercalendareventattachment.UserCalendarEventAttachmentClient
	UserCalendarEventCalendar                                                    *usercalendareventcalendar.UserCalendarEventCalendarClient
	UserCalendarEventExceptionOccurrence                                         *usercalendareventexceptionoccurrence.UserCalendarEventExceptionOccurrenceClient
	UserCalendarEventExceptionOccurrenceAttachment                               *usercalendareventexceptionoccurrenceattachment.UserCalendarEventExceptionOccurrenceAttachmentClient
	UserCalendarEventExceptionOccurrenceCalendar                                 *usercalendareventexceptionoccurrencecalendar.UserCalendarEventExceptionOccurrenceCalendarClient
	UserCalendarEventExceptionOccurrenceExtension                                *usercalendareventexceptionoccurrenceextension.UserCalendarEventExceptionOccurrenceExtensionClient
	UserCalendarEventExceptionOccurrenceInstance                                 *usercalendareventexceptionoccurrenceinstance.UserCalendarEventExceptionOccurrenceInstanceClient
	UserCalendarEventExceptionOccurrenceInstanceAttachment                       *usercalendareventexceptionoccurrenceinstanceattachment.UserCalendarEventExceptionOccurrenceInstanceAttachmentClient
	UserCalendarEventExceptionOccurrenceInstanceCalendar                         *usercalendareventexceptionoccurrenceinstancecalendar.UserCalendarEventExceptionOccurrenceInstanceCalendarClient
	UserCalendarEventExceptionOccurrenceInstanceExtension                        *usercalendareventexceptionoccurrenceinstanceextension.UserCalendarEventExceptionOccurrenceInstanceExtensionClient
	UserCalendarEventExtension                                                   *usercalendareventextension.UserCalendarEventExtensionClient
	UserCalendarEventInstance                                                    *usercalendareventinstance.UserCalendarEventInstanceClient
	UserCalendarEventInstanceAttachment                                          *usercalendareventinstanceattachment.UserCalendarEventInstanceAttachmentClient
	UserCalendarEventInstanceCalendar                                            *usercalendareventinstancecalendar.UserCalendarEventInstanceCalendarClient
	UserCalendarEventInstanceExceptionOccurrence                                 *usercalendareventinstanceexceptionoccurrence.UserCalendarEventInstanceExceptionOccurrenceClient
	UserCalendarEventInstanceExceptionOccurrenceAttachment                       *usercalendareventinstanceexceptionoccurrenceattachment.UserCalendarEventInstanceExceptionOccurrenceAttachmentClient
	UserCalendarEventInstanceExceptionOccurrenceCalendar                         *usercalendareventinstanceexceptionoccurrencecalendar.UserCalendarEventInstanceExceptionOccurrenceCalendarClient
	UserCalendarEventInstanceExceptionOccurrenceExtension                        *usercalendareventinstanceexceptionoccurrenceextension.UserCalendarEventInstanceExceptionOccurrenceExtensionClient
	UserCalendarEventInstanceExtension                                           *usercalendareventinstanceextension.UserCalendarEventInstanceExtensionClient
	UserCalendarGroup                                                            *usercalendargroup.UserCalendarGroupClient
	UserCalendarGroupCalendar                                                    *usercalendargroupcalendar.UserCalendarGroupCalendarClient
	UserCalendarGroupCalendarCalendarPermission                                  *usercalendargroupcalendarcalendarpermission.UserCalendarGroupCalendarCalendarPermissionClient
	UserCalendarGroupCalendarCalendarView                                        *usercalendargroupcalendarcalendarview.UserCalendarGroupCalendarCalendarViewClient
	UserCalendarGroupCalendarCalendarViewAttachment                              *usercalendargroupcalendarcalendarviewattachment.UserCalendarGroupCalendarCalendarViewAttachmentClient
	UserCalendarGroupCalendarCalendarViewCalendar                                *usercalendargroupcalendarcalendarviewcalendar.UserCalendarGroupCalendarCalendarViewCalendarClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrence                     *usercalendargroupcalendarcalendarviewexceptionoccurrence.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment           *usercalendargroupcalendarcalendarviewexceptionoccurrenceattachment.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar             *usercalendargroupcalendarcalendarviewexceptionoccurrencecalendar.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtension            *usercalendargroupcalendarcalendarviewexceptionoccurrenceextension.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance             *usercalendargroupcalendarcalendarviewexceptionoccurrenceinstance.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment   *usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar     *usercalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension    *usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension.UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	UserCalendarGroupCalendarCalendarViewExtension                               *usercalendargroupcalendarcalendarviewextension.UserCalendarGroupCalendarCalendarViewExtensionClient
	UserCalendarGroupCalendarCalendarViewInstance                                *usercalendargroupcalendarcalendarviewinstance.UserCalendarGroupCalendarCalendarViewInstanceClient
	UserCalendarGroupCalendarCalendarViewInstanceAttachment                      *usercalendargroupcalendarcalendarviewinstanceattachment.UserCalendarGroupCalendarCalendarViewInstanceAttachmentClient
	UserCalendarGroupCalendarCalendarViewInstanceCalendar                        *usercalendargroupcalendarcalendarviewinstancecalendar.UserCalendarGroupCalendarCalendarViewInstanceCalendarClient
	UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence             *usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient
	UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment   *usercalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar     *usercalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar.UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension    *usercalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension.UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	UserCalendarGroupCalendarCalendarViewInstanceExtension                       *usercalendargroupcalendarcalendarviewinstanceextension.UserCalendarGroupCalendarCalendarViewInstanceExtensionClient
	UserCalendarGroupCalendarEvent                                               *usercalendargroupcalendarevent.UserCalendarGroupCalendarEventClient
	UserCalendarGroupCalendarEventAttachment                                     *usercalendargroupcalendareventattachment.UserCalendarGroupCalendarEventAttachmentClient
	UserCalendarGroupCalendarEventCalendar                                       *usercalendargroupcalendareventcalendar.UserCalendarGroupCalendarEventCalendarClient
	UserCalendarGroupCalendarEventExceptionOccurrence                            *usercalendargroupcalendareventexceptionoccurrence.UserCalendarGroupCalendarEventExceptionOccurrenceClient
	UserCalendarGroupCalendarEventExceptionOccurrenceAttachment                  *usercalendargroupcalendareventexceptionoccurrenceattachment.UserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient
	UserCalendarGroupCalendarEventExceptionOccurrenceCalendar                    *usercalendargroupcalendareventexceptionoccurrencecalendar.UserCalendarGroupCalendarEventExceptionOccurrenceCalendarClient
	UserCalendarGroupCalendarEventExceptionOccurrenceExtension                   *usercalendargroupcalendareventexceptionoccurrenceextension.UserCalendarGroupCalendarEventExceptionOccurrenceExtensionClient
	UserCalendarGroupCalendarEventExceptionOccurrenceInstance                    *usercalendargroupcalendareventexceptionoccurrenceinstance.UserCalendarGroupCalendarEventExceptionOccurrenceInstanceClient
	UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment          *usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient
	UserCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar            *usercalendargroupcalendareventexceptionoccurrenceinstancecalendar.UserCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient
	UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtension           *usercalendargroupcalendareventexceptionoccurrenceinstanceextension.UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient
	UserCalendarGroupCalendarEventExtension                                      *usercalendargroupcalendareventextension.UserCalendarGroupCalendarEventExtensionClient
	UserCalendarGroupCalendarEventInstance                                       *usercalendargroupcalendareventinstance.UserCalendarGroupCalendarEventInstanceClient
	UserCalendarGroupCalendarEventInstanceAttachment                             *usercalendargroupcalendareventinstanceattachment.UserCalendarGroupCalendarEventInstanceAttachmentClient
	UserCalendarGroupCalendarEventInstanceCalendar                               *usercalendargroupcalendareventinstancecalendar.UserCalendarGroupCalendarEventInstanceCalendarClient
	UserCalendarGroupCalendarEventInstanceExceptionOccurrence                    *usercalendargroupcalendareventinstanceexceptionoccurrence.UserCalendarGroupCalendarEventInstanceExceptionOccurrenceClient
	UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment          *usercalendargroupcalendareventinstanceexceptionoccurrenceattachment.UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient
	UserCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar            *usercalendargroupcalendareventinstanceexceptionoccurrencecalendar.UserCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient
	UserCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension           *usercalendargroupcalendareventinstanceexceptionoccurrenceextension.UserCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient
	UserCalendarGroupCalendarEventInstanceExtension                              *usercalendargroupcalendareventinstanceextension.UserCalendarGroupCalendarEventInstanceExtensionClient
	UserCalendarView                                                             *usercalendarview.UserCalendarViewClient
	UserCalendarViewAttachment                                                   *usercalendarviewattachment.UserCalendarViewAttachmentClient
	UserCalendarViewCalendar                                                     *usercalendarviewcalendar.UserCalendarViewCalendarClient
	UserCalendarViewExceptionOccurrence                                          *usercalendarviewexceptionoccurrence.UserCalendarViewExceptionOccurrenceClient
	UserCalendarViewExceptionOccurrenceAttachment                                *usercalendarviewexceptionoccurrenceattachment.UserCalendarViewExceptionOccurrenceAttachmentClient
	UserCalendarViewExceptionOccurrenceCalendar                                  *usercalendarviewexceptionoccurrencecalendar.UserCalendarViewExceptionOccurrenceCalendarClient
	UserCalendarViewExceptionOccurrenceExtension                                 *usercalendarviewexceptionoccurrenceextension.UserCalendarViewExceptionOccurrenceExtensionClient
	UserCalendarViewExceptionOccurrenceInstance                                  *usercalendarviewexceptionoccurrenceinstance.UserCalendarViewExceptionOccurrenceInstanceClient
	UserCalendarViewExceptionOccurrenceInstanceAttachment                        *usercalendarviewexceptionoccurrenceinstanceattachment.UserCalendarViewExceptionOccurrenceInstanceAttachmentClient
	UserCalendarViewExceptionOccurrenceInstanceCalendar                          *usercalendarviewexceptionoccurrenceinstancecalendar.UserCalendarViewExceptionOccurrenceInstanceCalendarClient
	UserCalendarViewExceptionOccurrenceInstanceExtension                         *usercalendarviewexceptionoccurrenceinstanceextension.UserCalendarViewExceptionOccurrenceInstanceExtensionClient
	UserCalendarViewExtension                                                    *usercalendarviewextension.UserCalendarViewExtensionClient
	UserCalendarViewInstance                                                     *usercalendarviewinstance.UserCalendarViewInstanceClient
	UserCalendarViewInstanceAttachment                                           *usercalendarviewinstanceattachment.UserCalendarViewInstanceAttachmentClient
	UserCalendarViewInstanceCalendar                                             *usercalendarviewinstancecalendar.UserCalendarViewInstanceCalendarClient
	UserCalendarViewInstanceExceptionOccurrence                                  *usercalendarviewinstanceexceptionoccurrence.UserCalendarViewInstanceExceptionOccurrenceClient
	UserCalendarViewInstanceExceptionOccurrenceAttachment                        *usercalendarviewinstanceexceptionoccurrenceattachment.UserCalendarViewInstanceExceptionOccurrenceAttachmentClient
	UserCalendarViewInstanceExceptionOccurrenceCalendar                          *usercalendarviewinstanceexceptionoccurrencecalendar.UserCalendarViewInstanceExceptionOccurrenceCalendarClient
	UserCalendarViewInstanceExceptionOccurrenceExtension                         *usercalendarviewinstanceexceptionoccurrenceextension.UserCalendarViewInstanceExceptionOccurrenceExtensionClient
	UserCalendarViewInstanceExtension                                            *usercalendarviewinstanceextension.UserCalendarViewInstanceExtensionClient
	UserChat                                                                     *userchat.UserChatClient
	UserChatInstalledApp                                                         *userchatinstalledapp.UserChatInstalledAppClient
	UserChatInstalledAppTeamsApp                                                 *userchatinstalledappteamsapp.UserChatInstalledAppTeamsAppClient
	UserChatInstalledAppTeamsAppDefinition                                       *userchatinstalledappteamsappdefinition.UserChatInstalledAppTeamsAppDefinitionClient
	UserChatLastMessagePreview                                                   *userchatlastmessagepreview.UserChatLastMessagePreviewClient
	UserChatMember                                                               *userchatmember.UserChatMemberClient
	UserChatMessage                                                              *userchatmessage.UserChatMessageClient
	UserChatMessageHostedContent                                                 *userchatmessagehostedcontent.UserChatMessageHostedContentClient
	UserChatMessageReply                                                         *userchatmessagereply.UserChatMessageReplyClient
	UserChatMessageReplyHostedContent                                            *userchatmessagereplyhostedcontent.UserChatMessageReplyHostedContentClient
	UserChatOperation                                                            *userchatoperation.UserChatOperationClient
	UserChatPermissionGrant                                                      *userchatpermissiongrant.UserChatPermissionGrantClient
	UserChatPinnedMessage                                                        *userchatpinnedmessage.UserChatPinnedMessageClient
	UserChatPinnedMessageMessage                                                 *userchatpinnedmessagemessage.UserChatPinnedMessageMessageClient
	UserChatTab                                                                  *userchattab.UserChatTabClient
	UserChatTabTeamsApp                                                          *userchattabteamsapp.UserChatTabTeamsAppClient
	UserCloudPC                                                                  *usercloudpc.UserCloudPCClient
	UserContact                                                                  *usercontact.UserContactClient
	UserContactExtension                                                         *usercontactextension.UserContactExtensionClient
	UserContactFolder                                                            *usercontactfolder.UserContactFolderClient
	UserContactFolderChildFolder                                                 *usercontactfolderchildfolder.UserContactFolderChildFolderClient
	UserContactFolderChildFolderContact                                          *usercontactfolderchildfoldercontact.UserContactFolderChildFolderContactClient
	UserContactFolderChildFolderContactExtension                                 *usercontactfolderchildfoldercontactextension.UserContactFolderChildFolderContactExtensionClient
	UserContactFolderChildFolderContactPhoto                                     *usercontactfolderchildfoldercontactphoto.UserContactFolderChildFolderContactPhotoClient
	UserContactFolderContact                                                     *usercontactfoldercontact.UserContactFolderContactClient
	UserContactFolderContactExtension                                            *usercontactfoldercontactextension.UserContactFolderContactExtensionClient
	UserContactFolderContactPhoto                                                *usercontactfoldercontactphoto.UserContactFolderContactPhotoClient
	UserContactPhoto                                                             *usercontactphoto.UserContactPhotoClient
	UserCreatedObject                                                            *usercreatedobject.UserCreatedObjectClient
	UserDevice                                                                   *userdevice.UserDeviceClient
	UserDeviceCommand                                                            *userdevicecommand.UserDeviceCommandClient
	UserDeviceCommandResponsepayload                                             *userdevicecommandresponsepayload.UserDeviceCommandResponsepayloadClient
	UserDeviceEnrollmentConfiguration                                            *userdeviceenrollmentconfiguration.UserDeviceEnrollmentConfigurationClient
	UserDeviceEnrollmentConfigurationAssignment                                  *userdeviceenrollmentconfigurationassignment.UserDeviceEnrollmentConfigurationAssignmentClient
	UserDeviceExtension                                                          *userdeviceextension.UserDeviceExtensionClient
	UserDeviceManagementTroubleshootingEvent                                     *userdevicemanagementtroubleshootingevent.UserDeviceManagementTroubleshootingEventClient
	UserDeviceMemberOf                                                           *userdevicememberof.UserDeviceMemberOfClient
	UserDeviceRegisteredOwner                                                    *userdeviceregisteredowner.UserDeviceRegisteredOwnerClient
	UserDeviceRegisteredUser                                                     *userdeviceregistereduser.UserDeviceRegisteredUserClient
	UserDeviceTransitiveMemberOf                                                 *userdevicetransitivememberof.UserDeviceTransitiveMemberOfClient
	UserDeviceUsageRight                                                         *userdeviceusageright.UserDeviceUsageRightClient
	UserDirectReport                                                             *userdirectreport.UserDirectReportClient
	UserDrive                                                                    *userdrive.UserDriveClient
	UserEmployeeExperience                                                       *useremployeeexperience.UserEmployeeExperienceClient
	UserEmployeeExperienceLearningCourseActivity                                 *useremployeeexperiencelearningcourseactivity.UserEmployeeExperienceLearningCourseActivityClient
	UserEvent                                                                    *userevent.UserEventClient
	UserEventAttachment                                                          *usereventattachment.UserEventAttachmentClient
	UserEventCalendar                                                            *usereventcalendar.UserEventCalendarClient
	UserEventExceptionOccurrence                                                 *usereventexceptionoccurrence.UserEventExceptionOccurrenceClient
	UserEventExceptionOccurrenceAttachment                                       *usereventexceptionoccurrenceattachment.UserEventExceptionOccurrenceAttachmentClient
	UserEventExceptionOccurrenceCalendar                                         *usereventexceptionoccurrencecalendar.UserEventExceptionOccurrenceCalendarClient
	UserEventExceptionOccurrenceExtension                                        *usereventexceptionoccurrenceextension.UserEventExceptionOccurrenceExtensionClient
	UserEventExceptionOccurrenceInstance                                         *usereventexceptionoccurrenceinstance.UserEventExceptionOccurrenceInstanceClient
	UserEventExceptionOccurrenceInstanceAttachment                               *usereventexceptionoccurrenceinstanceattachment.UserEventExceptionOccurrenceInstanceAttachmentClient
	UserEventExceptionOccurrenceInstanceCalendar                                 *usereventexceptionoccurrenceinstancecalendar.UserEventExceptionOccurrenceInstanceCalendarClient
	UserEventExceptionOccurrenceInstanceExtension                                *usereventexceptionoccurrenceinstanceextension.UserEventExceptionOccurrenceInstanceExtensionClient
	UserEventExtension                                                           *usereventextension.UserEventExtensionClient
	UserEventInstance                                                            *usereventinstance.UserEventInstanceClient
	UserEventInstanceAttachment                                                  *usereventinstanceattachment.UserEventInstanceAttachmentClient
	UserEventInstanceCalendar                                                    *usereventinstancecalendar.UserEventInstanceCalendarClient
	UserEventInstanceExceptionOccurrence                                         *usereventinstanceexceptionoccurrence.UserEventInstanceExceptionOccurrenceClient
	UserEventInstanceExceptionOccurrenceAttachment                               *usereventinstanceexceptionoccurrenceattachment.UserEventInstanceExceptionOccurrenceAttachmentClient
	UserEventInstanceExceptionOccurrenceCalendar                                 *usereventinstanceexceptionoccurrencecalendar.UserEventInstanceExceptionOccurrenceCalendarClient
	UserEventInstanceExceptionOccurrenceExtension                                *usereventinstanceexceptionoccurrenceextension.UserEventInstanceExceptionOccurrenceExtensionClient
	UserEventInstanceExtension                                                   *usereventinstanceextension.UserEventInstanceExtensionClient
	UserExtension                                                                *userextension.UserExtensionClient
	UserFollowedSite                                                             *userfollowedsite.UserFollowedSiteClient
	UserInferenceClassification                                                  *userinferenceclassification.UserInferenceClassificationClient
	UserInferenceClassificationOverride                                          *userinferenceclassificationoverride.UserInferenceClassificationOverrideClient
	UserInformationProtection                                                    *userinformationprotection.UserInformationProtectionClient
	UserInformationProtectionBitlocker                                           *userinformationprotectionbitlocker.UserInformationProtectionBitlockerClient
	UserInformationProtectionBitlockerRecoveryKey                                *userinformationprotectionbitlockerrecoverykey.UserInformationProtectionBitlockerRecoveryKeyClient
	UserInformationProtectionDataLossPreventionPolicy                            *userinformationprotectiondatalosspreventionpolicy.UserInformationProtectionDataLossPreventionPolicyClient
	UserInformationProtectionPolicy                                              *userinformationprotectionpolicy.UserInformationProtectionPolicyClient
	UserInformationProtectionPolicyLabel                                         *userinformationprotectionpolicylabel.UserInformationProtectionPolicyLabelClient
	UserInformationProtectionSensitivityLabel                                    *userinformationprotectionsensitivitylabel.UserInformationProtectionSensitivityLabelClient
	UserInformationProtectionSensitivityLabelSublabel                            *userinformationprotectionsensitivitylabelsublabel.UserInformationProtectionSensitivityLabelSublabelClient
	UserInformationProtectionSensitivityPolicySetting                            *userinformationprotectionsensitivitypolicysetting.UserInformationProtectionSensitivityPolicySettingClient
	UserInformationProtectionThreatAssessmentRequest                             *userinformationprotectionthreatassessmentrequest.UserInformationProtectionThreatAssessmentRequestClient
	UserInformationProtectionThreatAssessmentRequestResult                       *userinformationprotectionthreatassessmentrequestresult.UserInformationProtectionThreatAssessmentRequestResultClient
	UserInsight                                                                  *userinsight.UserInsightClient
	UserInsightShared                                                            *userinsightshared.UserInsightSharedClient
	UserInsightSharedLastSharedMethod                                            *userinsightsharedlastsharedmethod.UserInsightSharedLastSharedMethodClient
	UserInsightSharedResource                                                    *userinsightsharedresource.UserInsightSharedResourceClient
	UserInsightTrending                                                          *userinsighttrending.UserInsightTrendingClient
	UserInsightTrendingResource                                                  *userinsighttrendingresource.UserInsightTrendingResourceClient
	UserInsightUsed                                                              *userinsightused.UserInsightUsedClient
	UserInsightUsedResource                                                      *userinsightusedresource.UserInsightUsedResourceClient
	UserJoinedGroup                                                              *userjoinedgroup.UserJoinedGroupClient
	UserJoinedTeam                                                               *userjoinedteam.UserJoinedTeamClient
	UserLicenseDetail                                                            *userlicensedetail.UserLicenseDetailClient
	UserMailFolder                                                               *usermailfolder.UserMailFolderClient
	UserMailFolderChildFolder                                                    *usermailfolderchildfolder.UserMailFolderChildFolderClient
	UserMailFolderChildFolderMessage                                             *usermailfolderchildfoldermessage.UserMailFolderChildFolderMessageClient
	UserMailFolderChildFolderMessageAttachment                                   *usermailfolderchildfoldermessageattachment.UserMailFolderChildFolderMessageAttachmentClient
	UserMailFolderChildFolderMessageExtension                                    *usermailfolderchildfoldermessageextension.UserMailFolderChildFolderMessageExtensionClient
	UserMailFolderChildFolderMessageMention                                      *usermailfolderchildfoldermessagemention.UserMailFolderChildFolderMessageMentionClient
	UserMailFolderChildFolderMessageRule                                         *usermailfolderchildfoldermessagerule.UserMailFolderChildFolderMessageRuleClient
	UserMailFolderChildFolderUserConfiguration                                   *usermailfolderchildfolderuserconfiguration.UserMailFolderChildFolderUserConfigurationClient
	UserMailFolderMessage                                                        *usermailfoldermessage.UserMailFolderMessageClient
	UserMailFolderMessageAttachment                                              *usermailfoldermessageattachment.UserMailFolderMessageAttachmentClient
	UserMailFolderMessageExtension                                               *usermailfoldermessageextension.UserMailFolderMessageExtensionClient
	UserMailFolderMessageMention                                                 *usermailfoldermessagemention.UserMailFolderMessageMentionClient
	UserMailFolderMessageRule                                                    *usermailfoldermessagerule.UserMailFolderMessageRuleClient
	UserMailFolderUserConfiguration                                              *usermailfolderuserconfiguration.UserMailFolderUserConfigurationClient
	UserMailboxSetting                                                           *usermailboxsetting.UserMailboxSettingClient
	UserManagedAppRegistration                                                   *usermanagedappregistration.UserManagedAppRegistrationClient
	UserManagedDevice                                                            *usermanageddevice.UserManagedDeviceClient
	UserManagedDeviceAssignmentFilterEvaluationStatusDetail                      *usermanageddeviceassignmentfilterevaluationstatusdetail.UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient
	UserManagedDeviceDetectedApp                                                 *usermanageddevicedetectedapp.UserManagedDeviceDetectedAppClient
	UserManagedDeviceDeviceCategory                                              *usermanageddevicedevicecategory.UserManagedDeviceDeviceCategoryClient
	UserManagedDeviceDeviceCompliancePolicyState                                 *usermanageddevicedevicecompliancepolicystate.UserManagedDeviceDeviceCompliancePolicyStateClient
	UserManagedDeviceDeviceConfigurationState                                    *usermanageddevicedeviceconfigurationstate.UserManagedDeviceDeviceConfigurationStateClient
	UserManagedDeviceDeviceHealthScriptState                                     *usermanageddevicedevicehealthscriptstate.UserManagedDeviceDeviceHealthScriptStateClient
	UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId *usermanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid.UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient
	UserManagedDeviceLogCollectionRequest                                        *usermanageddevicelogcollectionrequest.UserManagedDeviceLogCollectionRequestClient
	UserManagedDeviceManagedDeviceMobileAppConfigurationState                    *usermanageddevicemanageddevicemobileappconfigurationstate.UserManagedDeviceManagedDeviceMobileAppConfigurationStateClient
	UserManagedDeviceSecurityBaselineState                                       *usermanageddevicesecuritybaselinestate.UserManagedDeviceSecurityBaselineStateClient
	UserManagedDeviceSecurityBaselineStateSettingState                           *usermanageddevicesecuritybaselinestatesettingstate.UserManagedDeviceSecurityBaselineStateSettingStateClient
	UserManagedDeviceUser                                                        *usermanageddeviceuser.UserManagedDeviceUserClient
	UserManagedDeviceWindowsProtectionState                                      *usermanageddevicewindowsprotectionstate.UserManagedDeviceWindowsProtectionStateClient
	UserManagedDeviceWindowsProtectionStateDetectedMalwareState                  *usermanageddevicewindowsprotectionstatedetectedmalwarestate.UserManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	UserManager                                                                  *usermanager.UserManagerClient
	UserMemberOf                                                                 *usermemberof.UserMemberOfClient
	UserMessage                                                                  *usermessage.UserMessageClient
	UserMessageAttachment                                                        *usermessageattachment.UserMessageAttachmentClient
	UserMessageExtension                                                         *usermessageextension.UserMessageExtensionClient
	UserMessageMention                                                           *usermessagemention.UserMessageMentionClient
	UserMobileAppIntentAndState                                                  *usermobileappintentandstate.UserMobileAppIntentAndStateClient
	UserMobileAppTroubleshootingEvent                                            *usermobileapptroubleshootingevent.UserMobileAppTroubleshootingEventClient
	UserMobileAppTroubleshootingEventAppLogCollectionRequest                     *usermobileapptroubleshootingeventapplogcollectionrequest.UserMobileAppTroubleshootingEventAppLogCollectionRequestClient
	UserNotification                                                             *usernotification.UserNotificationClient
	UserOauth2PermissionGrant                                                    *useroauth2permissiongrant.UserOauth2PermissionGrantClient
	UserOnenote                                                                  *useronenote.UserOnenoteClient
	UserOnenoteNotebook                                                          *useronenotenotebook.UserOnenoteNotebookClient
	UserOnenoteNotebookSection                                                   *useronenotenotebooksection.UserOnenoteNotebookSectionClient
	UserOnenoteNotebookSectionGroup                                              *useronenotenotebooksectiongroup.UserOnenoteNotebookSectionGroupClient
	UserOnenoteNotebookSectionGroupParentNotebook                                *useronenotenotebooksectiongroupparentnotebook.UserOnenoteNotebookSectionGroupParentNotebookClient
	UserOnenoteNotebookSectionGroupParentSectionGroup                            *useronenotenotebooksectiongroupparentsectiongroup.UserOnenoteNotebookSectionGroupParentSectionGroupClient
	UserOnenoteNotebookSectionGroupSection                                       *useronenotenotebooksectiongroupsection.UserOnenoteNotebookSectionGroupSectionClient
	UserOnenoteNotebookSectionGroupSectionGroup                                  *useronenotenotebooksectiongroupsectiongroup.UserOnenoteNotebookSectionGroupSectionGroupClient
	UserOnenoteNotebookSectionGroupSectionPage                                   *useronenotenotebooksectiongroupsectionpage.UserOnenoteNotebookSectionGroupSectionPageClient
	UserOnenoteNotebookSectionGroupSectionPageContent                            *useronenotenotebooksectiongroupsectionpagecontent.UserOnenoteNotebookSectionGroupSectionPageContentClient
	UserOnenoteNotebookSectionGroupSectionPageParentNotebook                     *useronenotenotebooksectiongroupsectionpageparentnotebook.UserOnenoteNotebookSectionGroupSectionPageParentNotebookClient
	UserOnenoteNotebookSectionGroupSectionPageParentSection                      *useronenotenotebooksectiongroupsectionpageparentsection.UserOnenoteNotebookSectionGroupSectionPageParentSectionClient
	UserOnenoteNotebookSectionGroupSectionParentNotebook                         *useronenotenotebooksectiongroupsectionparentnotebook.UserOnenoteNotebookSectionGroupSectionParentNotebookClient
	UserOnenoteNotebookSectionGroupSectionParentSectionGroup                     *useronenotenotebooksectiongroupsectionparentsectiongroup.UserOnenoteNotebookSectionGroupSectionParentSectionGroupClient
	UserOnenoteNotebookSectionPage                                               *useronenotenotebooksectionpage.UserOnenoteNotebookSectionPageClient
	UserOnenoteNotebookSectionPageContent                                        *useronenotenotebooksectionpagecontent.UserOnenoteNotebookSectionPageContentClient
	UserOnenoteNotebookSectionPageParentNotebook                                 *useronenotenotebooksectionpageparentnotebook.UserOnenoteNotebookSectionPageParentNotebookClient
	UserOnenoteNotebookSectionPageParentSection                                  *useronenotenotebooksectionpageparentsection.UserOnenoteNotebookSectionPageParentSectionClient
	UserOnenoteNotebookSectionParentNotebook                                     *useronenotenotebooksectionparentnotebook.UserOnenoteNotebookSectionParentNotebookClient
	UserOnenoteNotebookSectionParentSectionGroup                                 *useronenotenotebooksectionparentsectiongroup.UserOnenoteNotebookSectionParentSectionGroupClient
	UserOnenoteOperation                                                         *useronenoteoperation.UserOnenoteOperationClient
	UserOnenotePage                                                              *useronenotepage.UserOnenotePageClient
	UserOnenotePageContent                                                       *useronenotepagecontent.UserOnenotePageContentClient
	UserOnenotePageParentNotebook                                                *useronenotepageparentnotebook.UserOnenotePageParentNotebookClient
	UserOnenotePageParentSection                                                 *useronenotepageparentsection.UserOnenotePageParentSectionClient
	UserOnenoteResource                                                          *useronenoteresource.UserOnenoteResourceClient
	UserOnenoteResourceContent                                                   *useronenoteresourcecontent.UserOnenoteResourceContentClient
	UserOnenoteSection                                                           *useronenotesection.UserOnenoteSectionClient
	UserOnenoteSectionGroup                                                      *useronenotesectiongroup.UserOnenoteSectionGroupClient
	UserOnenoteSectionGroupParentNotebook                                        *useronenotesectiongroupparentnotebook.UserOnenoteSectionGroupParentNotebookClient
	UserOnenoteSectionGroupParentSectionGroup                                    *useronenotesectiongroupparentsectiongroup.UserOnenoteSectionGroupParentSectionGroupClient
	UserOnenoteSectionGroupSection                                               *useronenotesectiongroupsection.UserOnenoteSectionGroupSectionClient
	UserOnenoteSectionGroupSectionGroup                                          *useronenotesectiongroupsectiongroup.UserOnenoteSectionGroupSectionGroupClient
	UserOnenoteSectionGroupSectionPage                                           *useronenotesectiongroupsectionpage.UserOnenoteSectionGroupSectionPageClient
	UserOnenoteSectionGroupSectionPageContent                                    *useronenotesectiongroupsectionpagecontent.UserOnenoteSectionGroupSectionPageContentClient
	UserOnenoteSectionGroupSectionPageParentNotebook                             *useronenotesectiongroupsectionpageparentnotebook.UserOnenoteSectionGroupSectionPageParentNotebookClient
	UserOnenoteSectionGroupSectionPageParentSection                              *useronenotesectiongroupsectionpageparentsection.UserOnenoteSectionGroupSectionPageParentSectionClient
	UserOnenoteSectionGroupSectionParentNotebook                                 *useronenotesectiongroupsectionparentnotebook.UserOnenoteSectionGroupSectionParentNotebookClient
	UserOnenoteSectionGroupSectionParentSectionGroup                             *useronenotesectiongroupsectionparentsectiongroup.UserOnenoteSectionGroupSectionParentSectionGroupClient
	UserOnenoteSectionPage                                                       *useronenotesectionpage.UserOnenoteSectionPageClient
	UserOnenoteSectionPageContent                                                *useronenotesectionpagecontent.UserOnenoteSectionPageContentClient
	UserOnenoteSectionPageParentNotebook                                         *useronenotesectionpageparentnotebook.UserOnenoteSectionPageParentNotebookClient
	UserOnenoteSectionPageParentSection                                          *useronenotesectionpageparentsection.UserOnenoteSectionPageParentSectionClient
	UserOnenoteSectionParentNotebook                                             *useronenotesectionparentnotebook.UserOnenoteSectionParentNotebookClient
	UserOnenoteSectionParentSectionGroup                                         *useronenotesectionparentsectiongroup.UserOnenoteSectionParentSectionGroupClient
	UserOnlineMeeting                                                            *useronlinemeeting.UserOnlineMeetingClient
	UserOnlineMeetingAlternativeRecording                                        *useronlinemeetingalternativerecording.UserOnlineMeetingAlternativeRecordingClient
	UserOnlineMeetingAttendanceReport                                            *useronlinemeetingattendancereport.UserOnlineMeetingAttendanceReportClient
	UserOnlineMeetingAttendanceReportAttendanceRecord                            *useronlinemeetingattendancereportattendancerecord.UserOnlineMeetingAttendanceReportAttendanceRecordClient
	UserOnlineMeetingAttendeeReport                                              *useronlinemeetingattendeereport.UserOnlineMeetingAttendeeReportClient
	UserOnlineMeetingBroadcastRecording                                          *useronlinemeetingbroadcastrecording.UserOnlineMeetingBroadcastRecordingClient
	UserOnlineMeetingMeetingAttendanceReport                                     *useronlinemeetingmeetingattendancereport.UserOnlineMeetingMeetingAttendanceReportClient
	UserOnlineMeetingMeetingAttendanceReportAttendanceRecord                     *useronlinemeetingmeetingattendancereportattendancerecord.UserOnlineMeetingMeetingAttendanceReportAttendanceRecordClient
	UserOnlineMeetingRecording                                                   *useronlinemeetingrecording.UserOnlineMeetingRecordingClient
	UserOnlineMeetingRecordingContent                                            *useronlinemeetingrecordingcontent.UserOnlineMeetingRecordingContentClient
	UserOnlineMeetingRegistration                                                *useronlinemeetingregistration.UserOnlineMeetingRegistrationClient
	UserOnlineMeetingRegistrationCustomQuestion                                  *useronlinemeetingregistrationcustomquestion.UserOnlineMeetingRegistrationCustomQuestionClient
	UserOnlineMeetingRegistrationRegistrant                                      *useronlinemeetingregistrationregistrant.UserOnlineMeetingRegistrationRegistrantClient
	UserOnlineMeetingTranscript                                                  *useronlinemeetingtranscript.UserOnlineMeetingTranscriptClient
	UserOnlineMeetingTranscriptContent                                           *useronlinemeetingtranscriptcontent.UserOnlineMeetingTranscriptContentClient
	UserOnlineMeetingTranscriptMetadataContent                                   *useronlinemeetingtranscriptmetadatacontent.UserOnlineMeetingTranscriptMetadataContentClient
	UserOutlook                                                                  *useroutlook.UserOutlookClient
	UserOutlookMasterCategory                                                    *useroutlookmastercategory.UserOutlookMasterCategoryClient
	UserOutlookTask                                                              *useroutlooktask.UserOutlookTaskClient
	UserOutlookTaskAttachment                                                    *useroutlooktaskattachment.UserOutlookTaskAttachmentClient
	UserOutlookTaskFolder                                                        *useroutlooktaskfolder.UserOutlookTaskFolderClient
	UserOutlookTaskFolderTask                                                    *useroutlooktaskfoldertask.UserOutlookTaskFolderTaskClient
	UserOutlookTaskFolderTaskAttachment                                          *useroutlooktaskfoldertaskattachment.UserOutlookTaskFolderTaskAttachmentClient
	UserOutlookTaskGroup                                                         *useroutlooktaskgroup.UserOutlookTaskGroupClient
	UserOutlookTaskGroupTaskFolder                                               *useroutlooktaskgrouptaskfolder.UserOutlookTaskGroupTaskFolderClient
	UserOutlookTaskGroupTaskFolderTask                                           *useroutlooktaskgrouptaskfoldertask.UserOutlookTaskGroupTaskFolderTaskClient
	UserOutlookTaskGroupTaskFolderTaskAttachment                                 *useroutlooktaskgrouptaskfoldertaskattachment.UserOutlookTaskGroupTaskFolderTaskAttachmentClient
	UserOwnedDevice                                                              *userowneddevice.UserOwnedDeviceClient
	UserOwnedObject                                                              *userownedobject.UserOwnedObjectClient
	UserPendingAccessReviewInstance                                              *userpendingaccessreviewinstance.UserPendingAccessReviewInstanceClient
	UserPendingAccessReviewInstanceContactedReviewer                             *userpendingaccessreviewinstancecontactedreviewer.UserPendingAccessReviewInstanceContactedReviewerClient
	UserPendingAccessReviewInstanceDecision                                      *userpendingaccessreviewinstancedecision.UserPendingAccessReviewInstanceDecisionClient
	UserPendingAccessReviewInstanceDecisionInsight                               *userpendingaccessreviewinstancedecisioninsight.UserPendingAccessReviewInstanceDecisionInsightClient
	UserPendingAccessReviewInstanceDecisionInstance                              *userpendingaccessreviewinstancedecisioninstance.UserPendingAccessReviewInstanceDecisionInstanceClient
	UserPendingAccessReviewInstanceDecisionInstanceContactedReviewer             *userpendingaccessreviewinstancedecisioninstancecontactedreviewer.UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient
	UserPendingAccessReviewInstanceDecisionInstanceDefinition                    *userpendingaccessreviewinstancedecisioninstancedefinition.UserPendingAccessReviewInstanceDecisionInstanceDefinitionClient
	UserPendingAccessReviewInstanceDecisionInstanceStage                         *userpendingaccessreviewinstancedecisioninstancestage.UserPendingAccessReviewInstanceDecisionInstanceStageClient
	UserPendingAccessReviewInstanceDecisionInstanceStageDecision                 *userpendingaccessreviewinstancedecisioninstancestagedecision.UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient
	UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsight          *userpendingaccessreviewinstancedecisioninstancestagedecisioninsight.UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient
	UserPendingAccessReviewInstanceDefinition                                    *userpendingaccessreviewinstancedefinition.UserPendingAccessReviewInstanceDefinitionClient
	UserPendingAccessReviewInstanceStage                                         *userpendingaccessreviewinstancestage.UserPendingAccessReviewInstanceStageClient
	UserPendingAccessReviewInstanceStageDecision                                 *userpendingaccessreviewinstancestagedecision.UserPendingAccessReviewInstanceStageDecisionClient
	UserPendingAccessReviewInstanceStageDecisionInsight                          *userpendingaccessreviewinstancestagedecisioninsight.UserPendingAccessReviewInstanceStageDecisionInsightClient
	UserPendingAccessReviewInstanceStageDecisionInstance                         *userpendingaccessreviewinstancestagedecisioninstance.UserPendingAccessReviewInstanceStageDecisionInstanceClient
	UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewer        *userpendingaccessreviewinstancestagedecisioninstancecontactedreviewer.UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient
	UserPendingAccessReviewInstanceStageDecisionInstanceDecision                 *userpendingaccessreviewinstancestagedecisioninstancedecision.UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient
	UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsight          *userpendingaccessreviewinstancestagedecisioninstancedecisioninsight.UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient
	UserPendingAccessReviewInstanceStageDecisionInstanceDefinition               *userpendingaccessreviewinstancestagedecisioninstancedefinition.UserPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient
	UserPeople                                                                   *userpeople.UserPeopleClient
	UserPermissionGrant                                                          *userpermissiongrant.UserPermissionGrantClient
	UserPhoto                                                                    *userphoto.UserPhotoClient
	UserPlanner                                                                  *userplanner.UserPlannerClient
	UserPlannerAll                                                               *userplannerall.UserPlannerAllClient
	UserPlannerFavoritePlan                                                      *userplannerfavoriteplan.UserPlannerFavoritePlanClient
	UserPlannerPlan                                                              *userplannerplan.UserPlannerPlanClient
	UserPlannerPlanBucket                                                        *userplannerplanbucket.UserPlannerPlanBucketClient
	UserPlannerPlanBucketTask                                                    *userplannerplanbuckettask.UserPlannerPlanBucketTaskClient
	UserPlannerPlanBucketTaskAssignedToTaskBoardFormat                           *userplannerplanbuckettaskassignedtotaskboardformat.UserPlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	UserPlannerPlanBucketTaskBucketTaskBoardFormat                               *userplannerplanbuckettaskbuckettaskboardformat.UserPlannerPlanBucketTaskBucketTaskBoardFormatClient
	UserPlannerPlanBucketTaskDetail                                              *userplannerplanbuckettaskdetail.UserPlannerPlanBucketTaskDetailClient
	UserPlannerPlanBucketTaskProgressTaskBoardFormat                             *userplannerplanbuckettaskprogresstaskboardformat.UserPlannerPlanBucketTaskProgressTaskBoardFormatClient
	UserPlannerPlanDetail                                                        *userplannerplandetail.UserPlannerPlanDetailClient
	UserPlannerPlanTask                                                          *userplannerplantask.UserPlannerPlanTaskClient
	UserPlannerPlanTaskAssignedToTaskBoardFormat                                 *userplannerplantaskassignedtotaskboardformat.UserPlannerPlanTaskAssignedToTaskBoardFormatClient
	UserPlannerPlanTaskBucketTaskBoardFormat                                     *userplannerplantaskbuckettaskboardformat.UserPlannerPlanTaskBucketTaskBoardFormatClient
	UserPlannerPlanTaskDetail                                                    *userplannerplantaskdetail.UserPlannerPlanTaskDetailClient
	UserPlannerPlanTaskProgressTaskBoardFormat                                   *userplannerplantaskprogresstaskboardformat.UserPlannerPlanTaskProgressTaskBoardFormatClient
	UserPlannerRecentPlan                                                        *userplannerrecentplan.UserPlannerRecentPlanClient
	UserPlannerRosterPlan                                                        *userplannerrosterplan.UserPlannerRosterPlanClient
	UserPlannerTask                                                              *userplannertask.UserPlannerTaskClient
	UserPlannerTaskAssignedToTaskBoardFormat                                     *userplannertaskassignedtotaskboardformat.UserPlannerTaskAssignedToTaskBoardFormatClient
	UserPlannerTaskBucketTaskBoardFormat                                         *userplannertaskbuckettaskboardformat.UserPlannerTaskBucketTaskBoardFormatClient
	UserPlannerTaskDetail                                                        *userplannertaskdetail.UserPlannerTaskDetailClient
	UserPlannerTaskProgressTaskBoardFormat                                       *userplannertaskprogresstaskboardformat.UserPlannerTaskProgressTaskBoardFormatClient
	UserPresence                                                                 *userpresence.UserPresenceClient
	UserProfile                                                                  *userprofile.UserProfileClient
	UserProfileAccount                                                           *userprofileaccount.UserProfileAccountClient
	UserProfileAddress                                                           *userprofileaddress.UserProfileAddressClient
	UserProfileAnniversary                                                       *userprofileanniversary.UserProfileAnniversaryClient
	UserProfileAward                                                             *userprofileaward.UserProfileAwardClient
	UserProfileCertification                                                     *userprofilecertification.UserProfileCertificationClient
	UserProfileEducationalActivity                                               *userprofileeducationalactivity.UserProfileEducationalActivityClient
	UserProfileEmail                                                             *userprofileemail.UserProfileEmailClient
	UserProfileInterest                                                          *userprofileinterest.UserProfileInterestClient
	UserProfileLanguage                                                          *userprofilelanguage.UserProfileLanguageClient
	UserProfileName                                                              *userprofilename.UserProfileNameClient
	UserProfileNote                                                              *userprofilenote.UserProfileNoteClient
	UserProfilePatent                                                            *userprofilepatent.UserProfilePatentClient
	UserProfilePhone                                                             *userprofilephone.UserProfilePhoneClient
	UserProfilePosition                                                          *userprofileposition.UserProfilePositionClient
	UserProfileProject                                                           *userprofileproject.UserProfileProjectClient
	UserProfilePublication                                                       *userprofilepublication.UserProfilePublicationClient
	UserProfileSkill                                                             *userprofileskill.UserProfileSkillClient
	UserProfileWebAccount                                                        *userprofilewebaccount.UserProfileWebAccountClient
	UserProfileWebsite                                                           *userprofilewebsite.UserProfileWebsiteClient
	UserRegisteredDevice                                                         *userregistereddevice.UserRegisteredDeviceClient
	UserScopedRoleMemberOf                                                       *userscopedrolememberof.UserScopedRoleMemberOfClient
	UserSecurity                                                                 *usersecurity.UserSecurityClient
	UserSecurityInformationProtection                                            *usersecurityinformationprotection.UserSecurityInformationProtectionClient
	UserSecurityInformationProtectionLabelPolicySetting                          *usersecurityinformationprotectionlabelpolicysetting.UserSecurityInformationProtectionLabelPolicySettingClient
	UserSecurityInformationProtectionSensitivityLabel                            *usersecurityinformationprotectionsensitivitylabel.UserSecurityInformationProtectionSensitivityLabelClient
	UserSecurityInformationProtectionSensitivityLabelParent                      *usersecurityinformationprotectionsensitivitylabelparent.UserSecurityInformationProtectionSensitivityLabelParentClient
	UserSetting                                                                  *usersetting.UserSettingClient
	UserSettingContactMergeSuggestion                                            *usersettingcontactmergesuggestion.UserSettingContactMergeSuggestionClient
	UserSettingItemInsight                                                       *usersettingiteminsight.UserSettingItemInsightClient
	UserSettingRegionalAndLanguageSetting                                        *usersettingregionalandlanguagesetting.UserSettingRegionalAndLanguageSettingClient
	UserSettingShiftPreference                                                   *usersettingshiftpreference.UserSettingShiftPreferenceClient
	UserSponsor                                                                  *usersponsor.UserSponsorClient
	UserTeamwork                                                                 *userteamwork.UserTeamworkClient
	UserTeamworkAssociatedTeam                                                   *userteamworkassociatedteam.UserTeamworkAssociatedTeamClient
	UserTeamworkAssociatedTeamTeam                                               *userteamworkassociatedteamteam.UserTeamworkAssociatedTeamTeamClient
	UserTeamworkInstalledApp                                                     *userteamworkinstalledapp.UserTeamworkInstalledAppClient
	UserTeamworkInstalledAppChat                                                 *userteamworkinstalledappchat.UserTeamworkInstalledAppChatClient
	UserTeamworkInstalledAppTeamsApp                                             *userteamworkinstalledappteamsapp.UserTeamworkInstalledAppTeamsAppClient
	UserTeamworkInstalledAppTeamsAppDefinition                                   *userteamworkinstalledappteamsappdefinition.UserTeamworkInstalledAppTeamsAppDefinitionClient
	UserTodo                                                                     *usertodo.UserTodoClient
	UserTodoList                                                                 *usertodolist.UserTodoListClient
	UserTodoListExtension                                                        *usertodolistextension.UserTodoListExtensionClient
	UserTodoListTask                                                             *usertodolisttask.UserTodoListTaskClient
	UserTodoListTaskAttachment                                                   *usertodolisttaskattachment.UserTodoListTaskAttachmentClient
	UserTodoListTaskAttachmentSession                                            *usertodolisttaskattachmentsession.UserTodoListTaskAttachmentSessionClient
	UserTodoListTaskAttachmentSessionContent                                     *usertodolisttaskattachmentsessioncontent.UserTodoListTaskAttachmentSessionContentClient
	UserTodoListTaskChecklistItem                                                *usertodolisttaskchecklistitem.UserTodoListTaskChecklistItemClient
	UserTodoListTaskExtension                                                    *usertodolisttaskextension.UserTodoListTaskExtensionClient
	UserTodoListTaskLinkedResource                                               *usertodolisttasklinkedresource.UserTodoListTaskLinkedResourceClient
	UserTransitiveMemberOf                                                       *usertransitivememberof.UserTransitiveMemberOfClient
	UserTransitiveReport                                                         *usertransitivereport.UserTransitiveReportClient
	UserUsageRight                                                               *userusageright.UserUsageRightClient
	UserWindowsInformationProtectionDeviceRegistration                           *userwindowsinformationprotectiondeviceregistration.UserWindowsInformationProtectionDeviceRegistrationClient
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

	userAnalyticActivityStatisticClient, err := useranalyticactivitystatistic.NewUserAnalyticActivityStatisticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAnalyticActivityStatistic client: %+v", err)
	}
	configureFunc(userAnalyticActivityStatisticClient.Client)

	userAnalyticClient, err := useranalytic.NewUserAnalyticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAnalytic client: %+v", err)
	}
	configureFunc(userAnalyticClient.Client)

	userAppConsentRequestsForApprovalClient, err := userappconsentrequestsforapproval.NewUserAppConsentRequestsForApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAppConsentRequestsForApproval client: %+v", err)
	}
	configureFunc(userAppConsentRequestsForApprovalClient.Client)

	userAppConsentRequestsForApprovalUserConsentRequestApprovalClient, err := userappconsentrequestsforapprovaluserconsentrequestapproval.NewUserAppConsentRequestsForApprovalUserConsentRequestApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAppConsentRequestsForApprovalUserConsentRequestApproval client: %+v", err)
	}
	configureFunc(userAppConsentRequestsForApprovalUserConsentRequestApprovalClient.Client)

	userAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient, err := userappconsentrequestsforapprovaluserconsentrequestapprovalstep.NewUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAppConsentRequestsForApprovalUserConsentRequestApprovalStep client: %+v", err)
	}
	configureFunc(userAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient.Client)

	userAppConsentRequestsForApprovalUserConsentRequestClient, err := userappconsentrequestsforapprovaluserconsentrequest.NewUserAppConsentRequestsForApprovalUserConsentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAppConsentRequestsForApprovalUserConsentRequest client: %+v", err)
	}
	configureFunc(userAppConsentRequestsForApprovalUserConsentRequestClient.Client)

	userAppRoleAssignedResourceClient, err := userapproleassignedresource.NewUserAppRoleAssignedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAppRoleAssignedResource client: %+v", err)
	}
	configureFunc(userAppRoleAssignedResourceClient.Client)

	userAppRoleAssignmentClient, err := userapproleassignment.NewUserAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAppRoleAssignment client: %+v", err)
	}
	configureFunc(userAppRoleAssignmentClient.Client)

	userApprovalClient, err := userapproval.NewUserApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserApproval client: %+v", err)
	}
	configureFunc(userApprovalClient.Client)

	userApprovalStepClient, err := userapprovalstep.NewUserApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserApprovalStep client: %+v", err)
	}
	configureFunc(userApprovalStepClient.Client)

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

	userAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient, err := userauthenticationpasswordlessmicrosoftauthenticatormethod.NewUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationPasswordlessMicrosoftAuthenticatorMethod client: %+v", err)
	}
	configureFunc(userAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient.Client)

	userAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient, err := userauthenticationpasswordlessmicrosoftauthenticatormethoddevice.NewUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice client: %+v", err)
	}
	configureFunc(userAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient.Client)

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

	userCalendarCalendarViewExceptionOccurrenceAttachmentClient, err := usercalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceAttachmentClient.Client)

	userCalendarCalendarViewExceptionOccurrenceCalendarClient, err := usercalendarcalendarviewexceptionoccurrencecalendar.NewUserCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceCalendarClient.Client)

	userCalendarCalendarViewExceptionOccurrenceClient, err := usercalendarcalendarviewexceptionoccurrence.NewUserCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceClient.Client)

	userCalendarCalendarViewExceptionOccurrenceExtensionClient, err := usercalendarcalendarviewexceptionoccurrenceextension.NewUserCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceExtensionClient.Client)

	userCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := usercalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	userCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, err := usercalendarcalendarviewexceptionoccurrenceinstancecalendar.NewUserCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	userCalendarCalendarViewExceptionOccurrenceInstanceClient, err := usercalendarcalendarviewexceptionoccurrenceinstance.NewUserCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceInstanceClient.Client)

	userCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient, err := usercalendarcalendarviewexceptionoccurrenceinstanceextension.NewUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

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

	userCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := usercalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	userCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, err := usercalendarcalendarviewinstanceexceptionoccurrencecalendar.NewUserCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	userCalendarCalendarViewInstanceExceptionOccurrenceClient, err := usercalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceExceptionOccurrenceClient.Client)

	userCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, err := usercalendarcalendarviewinstanceexceptionoccurrenceextension.NewUserCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

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

	userCalendarEventExceptionOccurrenceAttachmentClient, err := usercalendareventexceptionoccurrenceattachment.NewUserCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceAttachmentClient.Client)

	userCalendarEventExceptionOccurrenceCalendarClient, err := usercalendareventexceptionoccurrencecalendar.NewUserCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceCalendarClient.Client)

	userCalendarEventExceptionOccurrenceClient, err := usercalendareventexceptionoccurrence.NewUserCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceClient.Client)

	userCalendarEventExceptionOccurrenceExtensionClient, err := usercalendareventexceptionoccurrenceextension.NewUserCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceExtensionClient.Client)

	userCalendarEventExceptionOccurrenceInstanceAttachmentClient, err := usercalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceInstanceAttachmentClient.Client)

	userCalendarEventExceptionOccurrenceInstanceCalendarClient, err := usercalendareventexceptionoccurrenceinstancecalendar.NewUserCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceInstanceCalendarClient.Client)

	userCalendarEventExceptionOccurrenceInstanceClient, err := usercalendareventexceptionoccurrenceinstance.NewUserCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceInstanceClient.Client)

	userCalendarEventExceptionOccurrenceInstanceExtensionClient, err := usercalendareventexceptionoccurrenceinstanceextension.NewUserCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarEventExceptionOccurrenceInstanceExtensionClient.Client)

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

	userCalendarEventInstanceExceptionOccurrenceAttachmentClient, err := usercalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceExceptionOccurrenceAttachmentClient.Client)

	userCalendarEventInstanceExceptionOccurrenceCalendarClient, err := usercalendareventinstanceexceptionoccurrencecalendar.NewUserCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceExceptionOccurrenceCalendarClient.Client)

	userCalendarEventInstanceExceptionOccurrenceClient, err := usercalendareventinstanceexceptionoccurrence.NewUserCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceExceptionOccurrenceClient.Client)

	userCalendarEventInstanceExceptionOccurrenceExtensionClient, err := usercalendareventinstanceexceptionoccurrenceextension.NewUserCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarEventInstanceExceptionOccurrenceExtensionClient.Client)

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

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.Client)

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrencecalendar.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient.Client)

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.Client)

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrenceextension.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient.Client)

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.Client)

	userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient, err := usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension.NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

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

	userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, err := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient, err := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.Client)

	userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, err := usercalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension.NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

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

	userCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient, err := usercalendargroupcalendareventexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.Client)

	userCalendarGroupCalendarEventExceptionOccurrenceCalendarClient, err := usercalendargroupcalendareventexceptionoccurrencecalendar.NewUserCalendarGroupCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceCalendarClient.Client)

	userCalendarGroupCalendarEventExceptionOccurrenceClient, err := usercalendargroupcalendareventexceptionoccurrence.NewUserCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceClient.Client)

	userCalendarGroupCalendarEventExceptionOccurrenceExtensionClient, err := usercalendargroupcalendareventexceptionoccurrenceextension.NewUserCalendarGroupCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceExtensionClient.Client)

	userCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient, err := usercalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.Client)

	userCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient, err := usercalendargroupcalendareventexceptionoccurrenceinstancecalendar.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient.Client)

	userCalendarGroupCalendarEventExceptionOccurrenceInstanceClient, err := usercalendargroupcalendareventexceptionoccurrenceinstance.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.Client)

	userCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient, err := usercalendargroupcalendareventexceptionoccurrenceinstanceextension.NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient.Client)

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

	userCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient, err := usercalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.Client)

	userCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient, err := usercalendargroupcalendareventinstanceexceptionoccurrencecalendar.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient.Client)

	userCalendarGroupCalendarEventInstanceExceptionOccurrenceClient, err := usercalendargroupcalendareventinstanceexceptionoccurrence.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.Client)

	userCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient, err := usercalendargroupcalendareventinstanceexceptionoccurrenceextension.NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient.Client)

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

	userCalendarViewExceptionOccurrenceAttachmentClient, err := usercalendarviewexceptionoccurrenceattachment.NewUserCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceAttachmentClient.Client)

	userCalendarViewExceptionOccurrenceCalendarClient, err := usercalendarviewexceptionoccurrencecalendar.NewUserCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceCalendarClient.Client)

	userCalendarViewExceptionOccurrenceClient, err := usercalendarviewexceptionoccurrence.NewUserCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceClient.Client)

	userCalendarViewExceptionOccurrenceExtensionClient, err := usercalendarviewexceptionoccurrenceextension.NewUserCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceExtensionClient.Client)

	userCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := usercalendarviewexceptionoccurrenceinstanceattachment.NewUserCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	userCalendarViewExceptionOccurrenceInstanceCalendarClient, err := usercalendarviewexceptionoccurrenceinstancecalendar.NewUserCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	userCalendarViewExceptionOccurrenceInstanceClient, err := usercalendarviewexceptionoccurrenceinstance.NewUserCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceInstanceClient.Client)

	userCalendarViewExceptionOccurrenceInstanceExtensionClient, err := usercalendarviewexceptionoccurrenceinstanceextension.NewUserCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(userCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

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

	userCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := usercalendarviewinstanceexceptionoccurrenceattachment.NewUserCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	userCalendarViewInstanceExceptionOccurrenceCalendarClient, err := usercalendarviewinstanceexceptionoccurrencecalendar.NewUserCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	userCalendarViewInstanceExceptionOccurrenceClient, err := usercalendarviewinstanceexceptionoccurrence.NewUserCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceExceptionOccurrenceClient.Client)

	userCalendarViewInstanceExceptionOccurrenceExtensionClient, err := usercalendarviewinstanceexceptionoccurrenceextension.NewUserCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

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

	userChatOperationClient, err := userchatoperation.NewUserChatOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserChatOperation client: %+v", err)
	}
	configureFunc(userChatOperationClient.Client)

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

	userCloudPCClient, err := usercloudpc.NewUserCloudPCClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserCloudPC client: %+v", err)
	}
	configureFunc(userCloudPCClient.Client)

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

	userDeviceClient, err := userdevice.NewUserDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDevice client: %+v", err)
	}
	configureFunc(userDeviceClient.Client)

	userDeviceCommandClient, err := userdevicecommand.NewUserDeviceCommandClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceCommand client: %+v", err)
	}
	configureFunc(userDeviceCommandClient.Client)

	userDeviceCommandResponsepayloadClient, err := userdevicecommandresponsepayload.NewUserDeviceCommandResponsepayloadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceCommandResponsepayload client: %+v", err)
	}
	configureFunc(userDeviceCommandResponsepayloadClient.Client)

	userDeviceEnrollmentConfigurationAssignmentClient, err := userdeviceenrollmentconfigurationassignment.NewUserDeviceEnrollmentConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceEnrollmentConfigurationAssignment client: %+v", err)
	}
	configureFunc(userDeviceEnrollmentConfigurationAssignmentClient.Client)

	userDeviceEnrollmentConfigurationClient, err := userdeviceenrollmentconfiguration.NewUserDeviceEnrollmentConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceEnrollmentConfiguration client: %+v", err)
	}
	configureFunc(userDeviceEnrollmentConfigurationClient.Client)

	userDeviceExtensionClient, err := userdeviceextension.NewUserDeviceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceExtension client: %+v", err)
	}
	configureFunc(userDeviceExtensionClient.Client)

	userDeviceManagementTroubleshootingEventClient, err := userdevicemanagementtroubleshootingevent.NewUserDeviceManagementTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceManagementTroubleshootingEvent client: %+v", err)
	}
	configureFunc(userDeviceManagementTroubleshootingEventClient.Client)

	userDeviceMemberOfClient, err := userdevicememberof.NewUserDeviceMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceMemberOf client: %+v", err)
	}
	configureFunc(userDeviceMemberOfClient.Client)

	userDeviceRegisteredOwnerClient, err := userdeviceregisteredowner.NewUserDeviceRegisteredOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceRegisteredOwner client: %+v", err)
	}
	configureFunc(userDeviceRegisteredOwnerClient.Client)

	userDeviceRegisteredUserClient, err := userdeviceregistereduser.NewUserDeviceRegisteredUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceRegisteredUser client: %+v", err)
	}
	configureFunc(userDeviceRegisteredUserClient.Client)

	userDeviceTransitiveMemberOfClient, err := userdevicetransitivememberof.NewUserDeviceTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceTransitiveMemberOf client: %+v", err)
	}
	configureFunc(userDeviceTransitiveMemberOfClient.Client)

	userDeviceUsageRightClient, err := userdeviceusageright.NewUserDeviceUsageRightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserDeviceUsageRight client: %+v", err)
	}
	configureFunc(userDeviceUsageRightClient.Client)

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

	userEventExceptionOccurrenceAttachmentClient, err := usereventexceptionoccurrenceattachment.NewUserEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceAttachmentClient.Client)

	userEventExceptionOccurrenceCalendarClient, err := usereventexceptionoccurrencecalendar.NewUserEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceCalendarClient.Client)

	userEventExceptionOccurrenceClient, err := usereventexceptionoccurrence.NewUserEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceClient.Client)

	userEventExceptionOccurrenceExtensionClient, err := usereventexceptionoccurrenceextension.NewUserEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceExtensionClient.Client)

	userEventExceptionOccurrenceInstanceAttachmentClient, err := usereventexceptionoccurrenceinstanceattachment.NewUserEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceInstanceAttachmentClient.Client)

	userEventExceptionOccurrenceInstanceCalendarClient, err := usereventexceptionoccurrenceinstancecalendar.NewUserEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceInstanceCalendarClient.Client)

	userEventExceptionOccurrenceInstanceClient, err := usereventexceptionoccurrenceinstance.NewUserEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceInstanceClient.Client)

	userEventExceptionOccurrenceInstanceExtensionClient, err := usereventexceptionoccurrenceinstanceextension.NewUserEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(userEventExceptionOccurrenceInstanceExtensionClient.Client)

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

	userEventInstanceExceptionOccurrenceAttachmentClient, err := usereventinstanceexceptionoccurrenceattachment.NewUserEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(userEventInstanceExceptionOccurrenceAttachmentClient.Client)

	userEventInstanceExceptionOccurrenceCalendarClient, err := usereventinstanceexceptionoccurrencecalendar.NewUserEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(userEventInstanceExceptionOccurrenceCalendarClient.Client)

	userEventInstanceExceptionOccurrenceClient, err := usereventinstanceexceptionoccurrence.NewUserEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(userEventInstanceExceptionOccurrenceClient.Client)

	userEventInstanceExceptionOccurrenceExtensionClient, err := usereventinstanceexceptionoccurrenceextension.NewUserEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(userEventInstanceExceptionOccurrenceExtensionClient.Client)

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

	userInformationProtectionBitlockerClient, err := userinformationprotectionbitlocker.NewUserInformationProtectionBitlockerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionBitlocker client: %+v", err)
	}
	configureFunc(userInformationProtectionBitlockerClient.Client)

	userInformationProtectionBitlockerRecoveryKeyClient, err := userinformationprotectionbitlockerrecoverykey.NewUserInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionBitlockerRecoveryKey client: %+v", err)
	}
	configureFunc(userInformationProtectionBitlockerRecoveryKeyClient.Client)

	userInformationProtectionClient, err := userinformationprotection.NewUserInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtection client: %+v", err)
	}
	configureFunc(userInformationProtectionClient.Client)

	userInformationProtectionDataLossPreventionPolicyClient, err := userinformationprotectiondatalosspreventionpolicy.NewUserInformationProtectionDataLossPreventionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionDataLossPreventionPolicy client: %+v", err)
	}
	configureFunc(userInformationProtectionDataLossPreventionPolicyClient.Client)

	userInformationProtectionPolicyClient, err := userinformationprotectionpolicy.NewUserInformationProtectionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionPolicy client: %+v", err)
	}
	configureFunc(userInformationProtectionPolicyClient.Client)

	userInformationProtectionPolicyLabelClient, err := userinformationprotectionpolicylabel.NewUserInformationProtectionPolicyLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionPolicyLabel client: %+v", err)
	}
	configureFunc(userInformationProtectionPolicyLabelClient.Client)

	userInformationProtectionSensitivityLabelClient, err := userinformationprotectionsensitivitylabel.NewUserInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(userInformationProtectionSensitivityLabelClient.Client)

	userInformationProtectionSensitivityLabelSublabelClient, err := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelSublabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionSensitivityLabelSublabel client: %+v", err)
	}
	configureFunc(userInformationProtectionSensitivityLabelSublabelClient.Client)

	userInformationProtectionSensitivityPolicySettingClient, err := userinformationprotectionsensitivitypolicysetting.NewUserInformationProtectionSensitivityPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionSensitivityPolicySetting client: %+v", err)
	}
	configureFunc(userInformationProtectionSensitivityPolicySettingClient.Client)

	userInformationProtectionThreatAssessmentRequestClient, err := userinformationprotectionthreatassessmentrequest.NewUserInformationProtectionThreatAssessmentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionThreatAssessmentRequest client: %+v", err)
	}
	configureFunc(userInformationProtectionThreatAssessmentRequestClient.Client)

	userInformationProtectionThreatAssessmentRequestResultClient, err := userinformationprotectionthreatassessmentrequestresult.NewUserInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserInformationProtectionThreatAssessmentRequestResult client: %+v", err)
	}
	configureFunc(userInformationProtectionThreatAssessmentRequestResultClient.Client)

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

	userJoinedGroupClient, err := userjoinedgroup.NewUserJoinedGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedGroup client: %+v", err)
	}
	configureFunc(userJoinedGroupClient.Client)

	userJoinedTeamClient, err := userjoinedteam.NewUserJoinedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserJoinedTeam client: %+v", err)
	}
	configureFunc(userJoinedTeamClient.Client)

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

	userMailFolderChildFolderMessageMentionClient, err := usermailfolderchildfoldermessagemention.NewUserMailFolderChildFolderMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolderMessageMention client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderMessageMentionClient.Client)

	userMailFolderChildFolderMessageRuleClient, err := usermailfolderchildfoldermessagerule.NewUserMailFolderChildFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolderMessageRule client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderMessageRuleClient.Client)

	userMailFolderChildFolderUserConfigurationClient, err := usermailfolderchildfolderuserconfiguration.NewUserMailFolderChildFolderUserConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderChildFolderUserConfiguration client: %+v", err)
	}
	configureFunc(userMailFolderChildFolderUserConfigurationClient.Client)

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

	userMailFolderMessageMentionClient, err := usermailfoldermessagemention.NewUserMailFolderMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderMessageMention client: %+v", err)
	}
	configureFunc(userMailFolderMessageMentionClient.Client)

	userMailFolderMessageRuleClient, err := usermailfoldermessagerule.NewUserMailFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderMessageRule client: %+v", err)
	}
	configureFunc(userMailFolderMessageRuleClient.Client)

	userMailFolderUserConfigurationClient, err := usermailfolderuserconfiguration.NewUserMailFolderUserConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMailFolderUserConfiguration client: %+v", err)
	}
	configureFunc(userMailFolderUserConfigurationClient.Client)

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

	userManagedDeviceAssignmentFilterEvaluationStatusDetailClient, err := usermanageddeviceassignmentfilterevaluationstatusdetail.NewUserManagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceAssignmentFilterEvaluationStatusDetail client: %+v", err)
	}
	configureFunc(userManagedDeviceAssignmentFilterEvaluationStatusDetailClient.Client)

	userManagedDeviceClient, err := usermanageddevice.NewUserManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDevice client: %+v", err)
	}
	configureFunc(userManagedDeviceClient.Client)

	userManagedDeviceDetectedAppClient, err := usermanageddevicedetectedapp.NewUserManagedDeviceDetectedAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceDetectedApp client: %+v", err)
	}
	configureFunc(userManagedDeviceDetectedAppClient.Client)

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

	userManagedDeviceDeviceHealthScriptStateClient, err := usermanageddevicedevicehealthscriptstate.NewUserManagedDeviceDeviceHealthScriptStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceDeviceHealthScriptState client: %+v", err)
	}
	configureFunc(userManagedDeviceDeviceHealthScriptStateClient.Client)

	userManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient, err := usermanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid.NewUserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId client: %+v", err)
	}
	configureFunc(userManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient.Client)

	userManagedDeviceLogCollectionRequestClient, err := usermanageddevicelogcollectionrequest.NewUserManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(userManagedDeviceLogCollectionRequestClient.Client)

	userManagedDeviceManagedDeviceMobileAppConfigurationStateClient, err := usermanageddevicemanageddevicemobileappconfigurationstate.NewUserManagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceManagedDeviceMobileAppConfigurationState client: %+v", err)
	}
	configureFunc(userManagedDeviceManagedDeviceMobileAppConfigurationStateClient.Client)

	userManagedDeviceSecurityBaselineStateClient, err := usermanageddevicesecuritybaselinestate.NewUserManagedDeviceSecurityBaselineStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceSecurityBaselineState client: %+v", err)
	}
	configureFunc(userManagedDeviceSecurityBaselineStateClient.Client)

	userManagedDeviceSecurityBaselineStateSettingStateClient, err := usermanageddevicesecuritybaselinestatesettingstate.NewUserManagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserManagedDeviceSecurityBaselineStateSettingState client: %+v", err)
	}
	configureFunc(userManagedDeviceSecurityBaselineStateSettingStateClient.Client)

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

	userMessageMentionClient, err := usermessagemention.NewUserMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMessageMention client: %+v", err)
	}
	configureFunc(userMessageMentionClient.Client)

	userMobileAppIntentAndStateClient, err := usermobileappintentandstate.NewUserMobileAppIntentAndStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMobileAppIntentAndState client: %+v", err)
	}
	configureFunc(userMobileAppIntentAndStateClient.Client)

	userMobileAppTroubleshootingEventAppLogCollectionRequestClient, err := usermobileapptroubleshootingeventapplogcollectionrequest.NewUserMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMobileAppTroubleshootingEventAppLogCollectionRequest client: %+v", err)
	}
	configureFunc(userMobileAppTroubleshootingEventAppLogCollectionRequestClient.Client)

	userMobileAppTroubleshootingEventClient, err := usermobileapptroubleshootingevent.NewUserMobileAppTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserMobileAppTroubleshootingEvent client: %+v", err)
	}
	configureFunc(userMobileAppTroubleshootingEventClient.Client)

	userNotificationClient, err := usernotification.NewUserNotificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserNotification client: %+v", err)
	}
	configureFunc(userNotificationClient.Client)

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

	userOnlineMeetingAlternativeRecordingClient, err := useronlinemeetingalternativerecording.NewUserOnlineMeetingAlternativeRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingAlternativeRecording client: %+v", err)
	}
	configureFunc(userOnlineMeetingAlternativeRecordingClient.Client)

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

	userOnlineMeetingBroadcastRecordingClient, err := useronlinemeetingbroadcastrecording.NewUserOnlineMeetingBroadcastRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingBroadcastRecording client: %+v", err)
	}
	configureFunc(userOnlineMeetingBroadcastRecordingClient.Client)

	userOnlineMeetingClient, err := useronlinemeeting.NewUserOnlineMeetingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeeting client: %+v", err)
	}
	configureFunc(userOnlineMeetingClient.Client)

	userOnlineMeetingMeetingAttendanceReportAttendanceRecordClient, err := useronlinemeetingmeetingattendancereportattendancerecord.NewUserOnlineMeetingMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingMeetingAttendanceReportAttendanceRecord client: %+v", err)
	}
	configureFunc(userOnlineMeetingMeetingAttendanceReportAttendanceRecordClient.Client)

	userOnlineMeetingMeetingAttendanceReportClient, err := useronlinemeetingmeetingattendancereport.NewUserOnlineMeetingMeetingAttendanceReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingMeetingAttendanceReport client: %+v", err)
	}
	configureFunc(userOnlineMeetingMeetingAttendanceReportClient.Client)

	userOnlineMeetingRecordingClient, err := useronlinemeetingrecording.NewUserOnlineMeetingRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingRecording client: %+v", err)
	}
	configureFunc(userOnlineMeetingRecordingClient.Client)

	userOnlineMeetingRecordingContentClient, err := useronlinemeetingrecordingcontent.NewUserOnlineMeetingRecordingContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingRecordingContent client: %+v", err)
	}
	configureFunc(userOnlineMeetingRecordingContentClient.Client)

	userOnlineMeetingRegistrationClient, err := useronlinemeetingregistration.NewUserOnlineMeetingRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingRegistration client: %+v", err)
	}
	configureFunc(userOnlineMeetingRegistrationClient.Client)

	userOnlineMeetingRegistrationCustomQuestionClient, err := useronlinemeetingregistrationcustomquestion.NewUserOnlineMeetingRegistrationCustomQuestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingRegistrationCustomQuestion client: %+v", err)
	}
	configureFunc(userOnlineMeetingRegistrationCustomQuestionClient.Client)

	userOnlineMeetingRegistrationRegistrantClient, err := useronlinemeetingregistrationregistrant.NewUserOnlineMeetingRegistrationRegistrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingRegistrationRegistrant client: %+v", err)
	}
	configureFunc(userOnlineMeetingRegistrationRegistrantClient.Client)

	userOnlineMeetingTranscriptClient, err := useronlinemeetingtranscript.NewUserOnlineMeetingTranscriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingTranscript client: %+v", err)
	}
	configureFunc(userOnlineMeetingTranscriptClient.Client)

	userOnlineMeetingTranscriptContentClient, err := useronlinemeetingtranscriptcontent.NewUserOnlineMeetingTranscriptContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingTranscriptContent client: %+v", err)
	}
	configureFunc(userOnlineMeetingTranscriptContentClient.Client)

	userOnlineMeetingTranscriptMetadataContentClient, err := useronlinemeetingtranscriptmetadatacontent.NewUserOnlineMeetingTranscriptMetadataContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOnlineMeetingTranscriptMetadataContent client: %+v", err)
	}
	configureFunc(userOnlineMeetingTranscriptMetadataContentClient.Client)

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

	userOutlookTaskAttachmentClient, err := useroutlooktaskattachment.NewUserOutlookTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskAttachment client: %+v", err)
	}
	configureFunc(userOutlookTaskAttachmentClient.Client)

	userOutlookTaskClient, err := useroutlooktask.NewUserOutlookTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTask client: %+v", err)
	}
	configureFunc(userOutlookTaskClient.Client)

	userOutlookTaskFolderClient, err := useroutlooktaskfolder.NewUserOutlookTaskFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskFolder client: %+v", err)
	}
	configureFunc(userOutlookTaskFolderClient.Client)

	userOutlookTaskFolderTaskAttachmentClient, err := useroutlooktaskfoldertaskattachment.NewUserOutlookTaskFolderTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskFolderTaskAttachment client: %+v", err)
	}
	configureFunc(userOutlookTaskFolderTaskAttachmentClient.Client)

	userOutlookTaskFolderTaskClient, err := useroutlooktaskfoldertask.NewUserOutlookTaskFolderTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskFolderTask client: %+v", err)
	}
	configureFunc(userOutlookTaskFolderTaskClient.Client)

	userOutlookTaskGroupClient, err := useroutlooktaskgroup.NewUserOutlookTaskGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskGroup client: %+v", err)
	}
	configureFunc(userOutlookTaskGroupClient.Client)

	userOutlookTaskGroupTaskFolderClient, err := useroutlooktaskgrouptaskfolder.NewUserOutlookTaskGroupTaskFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskGroupTaskFolder client: %+v", err)
	}
	configureFunc(userOutlookTaskGroupTaskFolderClient.Client)

	userOutlookTaskGroupTaskFolderTaskAttachmentClient, err := useroutlooktaskgrouptaskfoldertaskattachment.NewUserOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskGroupTaskFolderTaskAttachment client: %+v", err)
	}
	configureFunc(userOutlookTaskGroupTaskFolderTaskAttachmentClient.Client)

	userOutlookTaskGroupTaskFolderTaskClient, err := useroutlooktaskgrouptaskfoldertask.NewUserOutlookTaskGroupTaskFolderTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserOutlookTaskGroupTaskFolderTask client: %+v", err)
	}
	configureFunc(userOutlookTaskGroupTaskFolderTaskClient.Client)

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

	userPendingAccessReviewInstanceClient, err := userpendingaccessreviewinstance.NewUserPendingAccessReviewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstance client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceClient.Client)

	userPendingAccessReviewInstanceContactedReviewerClient, err := userpendingaccessreviewinstancecontactedreviewer.NewUserPendingAccessReviewInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceContactedReviewerClient.Client)

	userPendingAccessReviewInstanceDecisionClient, err := userpendingaccessreviewinstancedecision.NewUserPendingAccessReviewInstanceDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecision client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionClient.Client)

	userPendingAccessReviewInstanceDecisionInsightClient, err := userpendingaccessreviewinstancedecisioninsight.NewUserPendingAccessReviewInstanceDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecisionInsight client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionInsightClient.Client)

	userPendingAccessReviewInstanceDecisionInstanceClient, err := userpendingaccessreviewinstancedecisioninstance.NewUserPendingAccessReviewInstanceDecisionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecisionInstance client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionInstanceClient.Client)

	userPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient, err := userpendingaccessreviewinstancedecisioninstancecontactedreviewer.NewUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecisionInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient.Client)

	userPendingAccessReviewInstanceDecisionInstanceDefinitionClient, err := userpendingaccessreviewinstancedecisioninstancedefinition.NewUserPendingAccessReviewInstanceDecisionInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecisionInstanceDefinition client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionInstanceDefinitionClient.Client)

	userPendingAccessReviewInstanceDecisionInstanceStageClient, err := userpendingaccessreviewinstancedecisioninstancestage.NewUserPendingAccessReviewInstanceDecisionInstanceStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecisionInstanceStage client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionInstanceStageClient.Client)

	userPendingAccessReviewInstanceDecisionInstanceStageDecisionClient, err := userpendingaccessreviewinstancedecisioninstancestagedecision.NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecisionInstanceStageDecision client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionInstanceStageDecisionClient.Client)

	userPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient, err := userpendingaccessreviewinstancedecisioninstancestagedecisioninsight.NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsight client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient.Client)

	userPendingAccessReviewInstanceDefinitionClient, err := userpendingaccessreviewinstancedefinition.NewUserPendingAccessReviewInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceDefinition client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceDefinitionClient.Client)

	userPendingAccessReviewInstanceStageClient, err := userpendingaccessreviewinstancestage.NewUserPendingAccessReviewInstanceStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStage client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageClient.Client)

	userPendingAccessReviewInstanceStageDecisionClient, err := userpendingaccessreviewinstancestagedecision.NewUserPendingAccessReviewInstanceStageDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStageDecision client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageDecisionClient.Client)

	userPendingAccessReviewInstanceStageDecisionInsightClient, err := userpendingaccessreviewinstancestagedecisioninsight.NewUserPendingAccessReviewInstanceStageDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStageDecisionInsight client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageDecisionInsightClient.Client)

	userPendingAccessReviewInstanceStageDecisionInstanceClient, err := userpendingaccessreviewinstancestagedecisioninstance.NewUserPendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStageDecisionInstance client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageDecisionInstanceClient.Client)

	userPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient, err := userpendingaccessreviewinstancestagedecisioninstancecontactedreviewer.NewUserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient.Client)

	userPendingAccessReviewInstanceStageDecisionInstanceDecisionClient, err := userpendingaccessreviewinstancestagedecisioninstancedecision.NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStageDecisionInstanceDecision client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageDecisionInstanceDecisionClient.Client)

	userPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient, err := userpendingaccessreviewinstancestagedecisioninstancedecisioninsight.NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsight client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient.Client)

	userPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient, err := userpendingaccessreviewinstancestagedecisioninstancedefinition.NewUserPendingAccessReviewInstanceStageDecisionInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPendingAccessReviewInstanceStageDecisionInstanceDefinition client: %+v", err)
	}
	configureFunc(userPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient.Client)

	userPeopleClient, err := userpeople.NewUserPeopleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPeople client: %+v", err)
	}
	configureFunc(userPeopleClient.Client)

	userPermissionGrantClient, err := userpermissiongrant.NewUserPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPermissionGrant client: %+v", err)
	}
	configureFunc(userPermissionGrantClient.Client)

	userPhotoClient, err := userphoto.NewUserPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPhoto client: %+v", err)
	}
	configureFunc(userPhotoClient.Client)

	userPlannerAllClient, err := userplannerall.NewUserPlannerAllClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerAll client: %+v", err)
	}
	configureFunc(userPlannerAllClient.Client)

	userPlannerClient, err := userplanner.NewUserPlannerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlanner client: %+v", err)
	}
	configureFunc(userPlannerClient.Client)

	userPlannerFavoritePlanClient, err := userplannerfavoriteplan.NewUserPlannerFavoritePlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerFavoritePlan client: %+v", err)
	}
	configureFunc(userPlannerFavoritePlanClient.Client)

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

	userPlannerRecentPlanClient, err := userplannerrecentplan.NewUserPlannerRecentPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerRecentPlan client: %+v", err)
	}
	configureFunc(userPlannerRecentPlanClient.Client)

	userPlannerRosterPlanClient, err := userplannerrosterplan.NewUserPlannerRosterPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserPlannerRosterPlan client: %+v", err)
	}
	configureFunc(userPlannerRosterPlanClient.Client)

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

	userProfileAccountClient, err := userprofileaccount.NewUserProfileAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileAccount client: %+v", err)
	}
	configureFunc(userProfileAccountClient.Client)

	userProfileAddressClient, err := userprofileaddress.NewUserProfileAddressClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileAddress client: %+v", err)
	}
	configureFunc(userProfileAddressClient.Client)

	userProfileAnniversaryClient, err := userprofileanniversary.NewUserProfileAnniversaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileAnniversary client: %+v", err)
	}
	configureFunc(userProfileAnniversaryClient.Client)

	userProfileAwardClient, err := userprofileaward.NewUserProfileAwardClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileAward client: %+v", err)
	}
	configureFunc(userProfileAwardClient.Client)

	userProfileCertificationClient, err := userprofilecertification.NewUserProfileCertificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileCertification client: %+v", err)
	}
	configureFunc(userProfileCertificationClient.Client)

	userProfileClient, err := userprofile.NewUserProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfile client: %+v", err)
	}
	configureFunc(userProfileClient.Client)

	userProfileEducationalActivityClient, err := userprofileeducationalactivity.NewUserProfileEducationalActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileEducationalActivity client: %+v", err)
	}
	configureFunc(userProfileEducationalActivityClient.Client)

	userProfileEmailClient, err := userprofileemail.NewUserProfileEmailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileEmail client: %+v", err)
	}
	configureFunc(userProfileEmailClient.Client)

	userProfileInterestClient, err := userprofileinterest.NewUserProfileInterestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileInterest client: %+v", err)
	}
	configureFunc(userProfileInterestClient.Client)

	userProfileLanguageClient, err := userprofilelanguage.NewUserProfileLanguageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileLanguage client: %+v", err)
	}
	configureFunc(userProfileLanguageClient.Client)

	userProfileNameClient, err := userprofilename.NewUserProfileNameClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileName client: %+v", err)
	}
	configureFunc(userProfileNameClient.Client)

	userProfileNoteClient, err := userprofilenote.NewUserProfileNoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileNote client: %+v", err)
	}
	configureFunc(userProfileNoteClient.Client)

	userProfilePatentClient, err := userprofilepatent.NewUserProfilePatentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfilePatent client: %+v", err)
	}
	configureFunc(userProfilePatentClient.Client)

	userProfilePhoneClient, err := userprofilephone.NewUserProfilePhoneClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfilePhone client: %+v", err)
	}
	configureFunc(userProfilePhoneClient.Client)

	userProfilePositionClient, err := userprofileposition.NewUserProfilePositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfilePosition client: %+v", err)
	}
	configureFunc(userProfilePositionClient.Client)

	userProfileProjectClient, err := userprofileproject.NewUserProfileProjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileProject client: %+v", err)
	}
	configureFunc(userProfileProjectClient.Client)

	userProfilePublicationClient, err := userprofilepublication.NewUserProfilePublicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfilePublication client: %+v", err)
	}
	configureFunc(userProfilePublicationClient.Client)

	userProfileSkillClient, err := userprofileskill.NewUserProfileSkillClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileSkill client: %+v", err)
	}
	configureFunc(userProfileSkillClient.Client)

	userProfileWebAccountClient, err := userprofilewebaccount.NewUserProfileWebAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileWebAccount client: %+v", err)
	}
	configureFunc(userProfileWebAccountClient.Client)

	userProfileWebsiteClient, err := userprofilewebsite.NewUserProfileWebsiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserProfileWebsite client: %+v", err)
	}
	configureFunc(userProfileWebsiteClient.Client)

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

	userSecurityClient, err := usersecurity.NewUserSecurityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSecurity client: %+v", err)
	}
	configureFunc(userSecurityClient.Client)

	userSecurityInformationProtectionClient, err := usersecurityinformationprotection.NewUserSecurityInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSecurityInformationProtection client: %+v", err)
	}
	configureFunc(userSecurityInformationProtectionClient.Client)

	userSecurityInformationProtectionLabelPolicySettingClient, err := usersecurityinformationprotectionlabelpolicysetting.NewUserSecurityInformationProtectionLabelPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSecurityInformationProtectionLabelPolicySetting client: %+v", err)
	}
	configureFunc(userSecurityInformationProtectionLabelPolicySettingClient.Client)

	userSecurityInformationProtectionSensitivityLabelClient, err := usersecurityinformationprotectionsensitivitylabel.NewUserSecurityInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSecurityInformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(userSecurityInformationProtectionSensitivityLabelClient.Client)

	userSecurityInformationProtectionSensitivityLabelParentClient, err := usersecurityinformationprotectionsensitivitylabelparent.NewUserSecurityInformationProtectionSensitivityLabelParentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSecurityInformationProtectionSensitivityLabelParent client: %+v", err)
	}
	configureFunc(userSecurityInformationProtectionSensitivityLabelParentClient.Client)

	userSettingClient, err := usersetting.NewUserSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSetting client: %+v", err)
	}
	configureFunc(userSettingClient.Client)

	userSettingContactMergeSuggestionClient, err := usersettingcontactmergesuggestion.NewUserSettingContactMergeSuggestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSettingContactMergeSuggestion client: %+v", err)
	}
	configureFunc(userSettingContactMergeSuggestionClient.Client)

	userSettingItemInsightClient, err := usersettingiteminsight.NewUserSettingItemInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSettingItemInsight client: %+v", err)
	}
	configureFunc(userSettingItemInsightClient.Client)

	userSettingRegionalAndLanguageSettingClient, err := usersettingregionalandlanguagesetting.NewUserSettingRegionalAndLanguageSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSettingRegionalAndLanguageSetting client: %+v", err)
	}
	configureFunc(userSettingRegionalAndLanguageSettingClient.Client)

	userSettingShiftPreferenceClient, err := usersettingshiftpreference.NewUserSettingShiftPreferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSettingShiftPreference client: %+v", err)
	}
	configureFunc(userSettingShiftPreferenceClient.Client)

	userSponsorClient, err := usersponsor.NewUserSponsorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSponsor client: %+v", err)
	}
	configureFunc(userSponsorClient.Client)

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

	userTransitiveReportClient, err := usertransitivereport.NewUserTransitiveReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserTransitiveReport client: %+v", err)
	}
	configureFunc(userTransitiveReportClient.Client)

	userUsageRightClient, err := userusageright.NewUserUsageRightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserUsageRight client: %+v", err)
	}
	configureFunc(userUsageRightClient.Client)

	userWindowsInformationProtectionDeviceRegistrationClient, err := userwindowsinformationprotectiondeviceregistration.NewUserWindowsInformationProtectionDeviceRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserWindowsInformationProtectionDeviceRegistration client: %+v", err)
	}
	configureFunc(userWindowsInformationProtectionDeviceRegistrationClient.Client)

	return &Client{
		User:                              userClient,
		UserActivity:                      userActivityClient,
		UserActivityHistoryItem:           userActivityHistoryItemClient,
		UserActivityHistoryItemActivity:   userActivityHistoryItemActivityClient,
		UserAgreementAcceptance:           userAgreementAcceptanceClient,
		UserAnalytic:                      userAnalyticClient,
		UserAnalyticActivityStatistic:     userAnalyticActivityStatisticClient,
		UserAppConsentRequestsForApproval: userAppConsentRequestsForApprovalClient,
		UserAppConsentRequestsForApprovalUserConsentRequest:              userAppConsentRequestsForApprovalUserConsentRequestClient,
		UserAppConsentRequestsForApprovalUserConsentRequestApproval:      userAppConsentRequestsForApprovalUserConsentRequestApprovalClient,
		UserAppConsentRequestsForApprovalUserConsentRequestApprovalStep:  userAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient,
		UserAppRoleAssignedResource:                                      userAppRoleAssignedResourceClient,
		UserAppRoleAssignment:                                            userAppRoleAssignmentClient,
		UserApproval:                                                     userApprovalClient,
		UserApprovalStep:                                                 userApprovalStepClient,
		UserAuthentication:                                               userAuthenticationClient,
		UserAuthenticationEmailMethod:                                    userAuthenticationEmailMethodClient,
		UserAuthenticationFido2Method:                                    userAuthenticationFido2MethodClient,
		UserAuthenticationMethod:                                         userAuthenticationMethodClient,
		UserAuthenticationMicrosoftAuthenticatorMethod:                   userAuthenticationMicrosoftAuthenticatorMethodClient,
		UserAuthenticationMicrosoftAuthenticatorMethodDevice:             userAuthenticationMicrosoftAuthenticatorMethodDeviceClient,
		UserAuthenticationOperation:                                      userAuthenticationOperationClient,
		UserAuthenticationPasswordMethod:                                 userAuthenticationPasswordMethodClient,
		UserAuthenticationPasswordlessMicrosoftAuthenticatorMethod:       userAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient,
		UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice: userAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient,
		UserAuthenticationPhoneMethod:                                    userAuthenticationPhoneMethodClient,
		UserAuthenticationSoftwareOathMethod:                             userAuthenticationSoftwareOathMethodClient,
		UserAuthenticationTemporaryAccessPassMethod:                      userAuthenticationTemporaryAccessPassMethodClient,
		UserAuthenticationWindowsHelloForBusinessMethod:                  userAuthenticationWindowsHelloForBusinessMethodClient,
		UserAuthenticationWindowsHelloForBusinessMethodDevice:            userAuthenticationWindowsHelloForBusinessMethodDeviceClient,
		UserCalendar:                                                               userCalendarClient,
		UserCalendarCalendarPermission:                                             userCalendarCalendarPermissionClient,
		UserCalendarCalendarView:                                                   userCalendarCalendarViewClient,
		UserCalendarCalendarViewAttachment:                                         userCalendarCalendarViewAttachmentClient,
		UserCalendarCalendarViewCalendar:                                           userCalendarCalendarViewCalendarClient,
		UserCalendarCalendarViewExceptionOccurrence:                                userCalendarCalendarViewExceptionOccurrenceClient,
		UserCalendarCalendarViewExceptionOccurrenceAttachment:                      userCalendarCalendarViewExceptionOccurrenceAttachmentClient,
		UserCalendarCalendarViewExceptionOccurrenceCalendar:                        userCalendarCalendarViewExceptionOccurrenceCalendarClient,
		UserCalendarCalendarViewExceptionOccurrenceExtension:                       userCalendarCalendarViewExceptionOccurrenceExtensionClient,
		UserCalendarCalendarViewExceptionOccurrenceInstance:                        userCalendarCalendarViewExceptionOccurrenceInstanceClient,
		UserCalendarCalendarViewExceptionOccurrenceInstanceAttachment:              userCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		UserCalendarCalendarViewExceptionOccurrenceInstanceCalendar:                userCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		UserCalendarCalendarViewExceptionOccurrenceInstanceExtension:               userCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		UserCalendarCalendarViewExtension:                                          userCalendarCalendarViewExtensionClient,
		UserCalendarCalendarViewInstance:                                           userCalendarCalendarViewInstanceClient,
		UserCalendarCalendarViewInstanceAttachment:                                 userCalendarCalendarViewInstanceAttachmentClient,
		UserCalendarCalendarViewInstanceCalendar:                                   userCalendarCalendarViewInstanceCalendarClient,
		UserCalendarCalendarViewInstanceExceptionOccurrence:                        userCalendarCalendarViewInstanceExceptionOccurrenceClient,
		UserCalendarCalendarViewInstanceExceptionOccurrenceAttachment:              userCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		UserCalendarCalendarViewInstanceExceptionOccurrenceCalendar:                userCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		UserCalendarCalendarViewInstanceExceptionOccurrenceExtension:               userCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		UserCalendarCalendarViewInstanceExtension:                                  userCalendarCalendarViewInstanceExtensionClient,
		UserCalendarEvent:                                                          userCalendarEventClient,
		UserCalendarEventAttachment:                                                userCalendarEventAttachmentClient,
		UserCalendarEventCalendar:                                                  userCalendarEventCalendarClient,
		UserCalendarEventExceptionOccurrence:                                       userCalendarEventExceptionOccurrenceClient,
		UserCalendarEventExceptionOccurrenceAttachment:                             userCalendarEventExceptionOccurrenceAttachmentClient,
		UserCalendarEventExceptionOccurrenceCalendar:                               userCalendarEventExceptionOccurrenceCalendarClient,
		UserCalendarEventExceptionOccurrenceExtension:                              userCalendarEventExceptionOccurrenceExtensionClient,
		UserCalendarEventExceptionOccurrenceInstance:                               userCalendarEventExceptionOccurrenceInstanceClient,
		UserCalendarEventExceptionOccurrenceInstanceAttachment:                     userCalendarEventExceptionOccurrenceInstanceAttachmentClient,
		UserCalendarEventExceptionOccurrenceInstanceCalendar:                       userCalendarEventExceptionOccurrenceInstanceCalendarClient,
		UserCalendarEventExceptionOccurrenceInstanceExtension:                      userCalendarEventExceptionOccurrenceInstanceExtensionClient,
		UserCalendarEventExtension:                                                 userCalendarEventExtensionClient,
		UserCalendarEventInstance:                                                  userCalendarEventInstanceClient,
		UserCalendarEventInstanceAttachment:                                        userCalendarEventInstanceAttachmentClient,
		UserCalendarEventInstanceCalendar:                                          userCalendarEventInstanceCalendarClient,
		UserCalendarEventInstanceExceptionOccurrence:                               userCalendarEventInstanceExceptionOccurrenceClient,
		UserCalendarEventInstanceExceptionOccurrenceAttachment:                     userCalendarEventInstanceExceptionOccurrenceAttachmentClient,
		UserCalendarEventInstanceExceptionOccurrenceCalendar:                       userCalendarEventInstanceExceptionOccurrenceCalendarClient,
		UserCalendarEventInstanceExceptionOccurrenceExtension:                      userCalendarEventInstanceExceptionOccurrenceExtensionClient,
		UserCalendarEventInstanceExtension:                                         userCalendarEventInstanceExtensionClient,
		UserCalendarGroup:                                                          userCalendarGroupClient,
		UserCalendarGroupCalendar:                                                  userCalendarGroupCalendarClient,
		UserCalendarGroupCalendarCalendarPermission:                                userCalendarGroupCalendarCalendarPermissionClient,
		UserCalendarGroupCalendarCalendarView:                                      userCalendarGroupCalendarCalendarViewClient,
		UserCalendarGroupCalendarCalendarViewAttachment:                            userCalendarGroupCalendarCalendarViewAttachmentClient,
		UserCalendarGroupCalendarCalendarViewCalendar:                              userCalendarGroupCalendarCalendarViewCalendarClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrence:                   userCalendarGroupCalendarCalendarViewExceptionOccurrenceClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment:         userCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar:           userCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtension:          userCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance:           userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment: userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar:   userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension:  userCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		UserCalendarGroupCalendarCalendarViewExtension:                             userCalendarGroupCalendarCalendarViewExtensionClient,
		UserCalendarGroupCalendarCalendarViewInstance:                              userCalendarGroupCalendarCalendarViewInstanceClient,
		UserCalendarGroupCalendarCalendarViewInstanceAttachment:                    userCalendarGroupCalendarCalendarViewInstanceAttachmentClient,
		UserCalendarGroupCalendarCalendarViewInstanceCalendar:                      userCalendarGroupCalendarCalendarViewInstanceCalendarClient,
		UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence:           userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient,
		UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment: userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar:   userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension:  userCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		UserCalendarGroupCalendarCalendarViewInstanceExtension:                     userCalendarGroupCalendarCalendarViewInstanceExtensionClient,
		UserCalendarGroupCalendarEvent:                                             userCalendarGroupCalendarEventClient,
		UserCalendarGroupCalendarEventAttachment:                                   userCalendarGroupCalendarEventAttachmentClient,
		UserCalendarGroupCalendarEventCalendar:                                     userCalendarGroupCalendarEventCalendarClient,
		UserCalendarGroupCalendarEventExceptionOccurrence:                          userCalendarGroupCalendarEventExceptionOccurrenceClient,
		UserCalendarGroupCalendarEventExceptionOccurrenceAttachment:                userCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient,
		UserCalendarGroupCalendarEventExceptionOccurrenceCalendar:                  userCalendarGroupCalendarEventExceptionOccurrenceCalendarClient,
		UserCalendarGroupCalendarEventExceptionOccurrenceExtension:                 userCalendarGroupCalendarEventExceptionOccurrenceExtensionClient,
		UserCalendarGroupCalendarEventExceptionOccurrenceInstance:                  userCalendarGroupCalendarEventExceptionOccurrenceInstanceClient,
		UserCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment:        userCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient,
		UserCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar:          userCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient,
		UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtension:         userCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient,
		UserCalendarGroupCalendarEventExtension:                                    userCalendarGroupCalendarEventExtensionClient,
		UserCalendarGroupCalendarEventInstance:                                     userCalendarGroupCalendarEventInstanceClient,
		UserCalendarGroupCalendarEventInstanceAttachment:                           userCalendarGroupCalendarEventInstanceAttachmentClient,
		UserCalendarGroupCalendarEventInstanceCalendar:                             userCalendarGroupCalendarEventInstanceCalendarClient,
		UserCalendarGroupCalendarEventInstanceExceptionOccurrence:                  userCalendarGroupCalendarEventInstanceExceptionOccurrenceClient,
		UserCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment:        userCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient,
		UserCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar:          userCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient,
		UserCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension:         userCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient,
		UserCalendarGroupCalendarEventInstanceExtension:                            userCalendarGroupCalendarEventInstanceExtensionClient,
		UserCalendarView:                                      userCalendarViewClient,
		UserCalendarViewAttachment:                            userCalendarViewAttachmentClient,
		UserCalendarViewCalendar:                              userCalendarViewCalendarClient,
		UserCalendarViewExceptionOccurrence:                   userCalendarViewExceptionOccurrenceClient,
		UserCalendarViewExceptionOccurrenceAttachment:         userCalendarViewExceptionOccurrenceAttachmentClient,
		UserCalendarViewExceptionOccurrenceCalendar:           userCalendarViewExceptionOccurrenceCalendarClient,
		UserCalendarViewExceptionOccurrenceExtension:          userCalendarViewExceptionOccurrenceExtensionClient,
		UserCalendarViewExceptionOccurrenceInstance:           userCalendarViewExceptionOccurrenceInstanceClient,
		UserCalendarViewExceptionOccurrenceInstanceAttachment: userCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		UserCalendarViewExceptionOccurrenceInstanceCalendar:   userCalendarViewExceptionOccurrenceInstanceCalendarClient,
		UserCalendarViewExceptionOccurrenceInstanceExtension:  userCalendarViewExceptionOccurrenceInstanceExtensionClient,
		UserCalendarViewExtension:                             userCalendarViewExtensionClient,
		UserCalendarViewInstance:                              userCalendarViewInstanceClient,
		UserCalendarViewInstanceAttachment:                    userCalendarViewInstanceAttachmentClient,
		UserCalendarViewInstanceCalendar:                      userCalendarViewInstanceCalendarClient,
		UserCalendarViewInstanceExceptionOccurrence:           userCalendarViewInstanceExceptionOccurrenceClient,
		UserCalendarViewInstanceExceptionOccurrenceAttachment: userCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		UserCalendarViewInstanceExceptionOccurrenceCalendar:   userCalendarViewInstanceExceptionOccurrenceCalendarClient,
		UserCalendarViewInstanceExceptionOccurrenceExtension:  userCalendarViewInstanceExceptionOccurrenceExtensionClient,
		UserCalendarViewInstanceExtension:                     userCalendarViewInstanceExtensionClient,
		UserChat:                                              userChatClient,
		UserChatInstalledApp:                                  userChatInstalledAppClient,
		UserChatInstalledAppTeamsApp:                          userChatInstalledAppTeamsAppClient,
		UserChatInstalledAppTeamsAppDefinition:                userChatInstalledAppTeamsAppDefinitionClient,
		UserChatLastMessagePreview:                            userChatLastMessagePreviewClient,
		UserChatMember:                                        userChatMemberClient,
		UserChatMessage:                                       userChatMessageClient,
		UserChatMessageHostedContent:                          userChatMessageHostedContentClient,
		UserChatMessageReply:                                  userChatMessageReplyClient,
		UserChatMessageReplyHostedContent:                     userChatMessageReplyHostedContentClient,
		UserChatOperation:                                     userChatOperationClient,
		UserChatPermissionGrant:                               userChatPermissionGrantClient,
		UserChatPinnedMessage:                                 userChatPinnedMessageClient,
		UserChatPinnedMessageMessage:                          userChatPinnedMessageMessageClient,
		UserChatTab:                                           userChatTabClient,
		UserChatTabTeamsApp:                                   userChatTabTeamsAppClient,
		UserCloudPC:                                           userCloudPCClient,
		UserContact:                                           userContactClient,
		UserContactExtension:                                  userContactExtensionClient,
		UserContactFolder:                                     userContactFolderClient,
		UserContactFolderChildFolder:                          userContactFolderChildFolderClient,
		UserContactFolderChildFolderContact:                   userContactFolderChildFolderContactClient,
		UserContactFolderChildFolderContactExtension:          userContactFolderChildFolderContactExtensionClient,
		UserContactFolderChildFolderContactPhoto:              userContactFolderChildFolderContactPhotoClient,
		UserContactFolderContact:                              userContactFolderContactClient,
		UserContactFolderContactExtension:                     userContactFolderContactExtensionClient,
		UserContactFolderContactPhoto:                         userContactFolderContactPhotoClient,
		UserContactPhoto:                                      userContactPhotoClient,
		UserCreatedObject:                                     userCreatedObjectClient,
		UserDevice:                                            userDeviceClient,
		UserDeviceCommand:                                     userDeviceCommandClient,
		UserDeviceCommandResponsepayload:                      userDeviceCommandResponsepayloadClient,
		UserDeviceEnrollmentConfiguration:                     userDeviceEnrollmentConfigurationClient,
		UserDeviceEnrollmentConfigurationAssignment:           userDeviceEnrollmentConfigurationAssignmentClient,
		UserDeviceExtension:                                   userDeviceExtensionClient,
		UserDeviceManagementTroubleshootingEvent:              userDeviceManagementTroubleshootingEventClient,
		UserDeviceMemberOf:                                    userDeviceMemberOfClient,
		UserDeviceRegisteredOwner:                             userDeviceRegisteredOwnerClient,
		UserDeviceRegisteredUser:                              userDeviceRegisteredUserClient,
		UserDeviceTransitiveMemberOf:                          userDeviceTransitiveMemberOfClient,
		UserDeviceUsageRight:                                  userDeviceUsageRightClient,
		UserDirectReport:                                      userDirectReportClient,
		UserDrive:                                             userDriveClient,
		UserEmployeeExperience:                                userEmployeeExperienceClient,
		UserEmployeeExperienceLearningCourseActivity:          userEmployeeExperienceLearningCourseActivityClient,
		UserEvent:                                              userEventClient,
		UserEventAttachment:                                    userEventAttachmentClient,
		UserEventCalendar:                                      userEventCalendarClient,
		UserEventExceptionOccurrence:                           userEventExceptionOccurrenceClient,
		UserEventExceptionOccurrenceAttachment:                 userEventExceptionOccurrenceAttachmentClient,
		UserEventExceptionOccurrenceCalendar:                   userEventExceptionOccurrenceCalendarClient,
		UserEventExceptionOccurrenceExtension:                  userEventExceptionOccurrenceExtensionClient,
		UserEventExceptionOccurrenceInstance:                   userEventExceptionOccurrenceInstanceClient,
		UserEventExceptionOccurrenceInstanceAttachment:         userEventExceptionOccurrenceInstanceAttachmentClient,
		UserEventExceptionOccurrenceInstanceCalendar:           userEventExceptionOccurrenceInstanceCalendarClient,
		UserEventExceptionOccurrenceInstanceExtension:          userEventExceptionOccurrenceInstanceExtensionClient,
		UserEventExtension:                                     userEventExtensionClient,
		UserEventInstance:                                      userEventInstanceClient,
		UserEventInstanceAttachment:                            userEventInstanceAttachmentClient,
		UserEventInstanceCalendar:                              userEventInstanceCalendarClient,
		UserEventInstanceExceptionOccurrence:                   userEventInstanceExceptionOccurrenceClient,
		UserEventInstanceExceptionOccurrenceAttachment:         userEventInstanceExceptionOccurrenceAttachmentClient,
		UserEventInstanceExceptionOccurrenceCalendar:           userEventInstanceExceptionOccurrenceCalendarClient,
		UserEventInstanceExceptionOccurrenceExtension:          userEventInstanceExceptionOccurrenceExtensionClient,
		UserEventInstanceExtension:                             userEventInstanceExtensionClient,
		UserExtension:                                          userExtensionClient,
		UserFollowedSite:                                       userFollowedSiteClient,
		UserInferenceClassification:                            userInferenceClassificationClient,
		UserInferenceClassificationOverride:                    userInferenceClassificationOverrideClient,
		UserInformationProtection:                              userInformationProtectionClient,
		UserInformationProtectionBitlocker:                     userInformationProtectionBitlockerClient,
		UserInformationProtectionBitlockerRecoveryKey:          userInformationProtectionBitlockerRecoveryKeyClient,
		UserInformationProtectionDataLossPreventionPolicy:      userInformationProtectionDataLossPreventionPolicyClient,
		UserInformationProtectionPolicy:                        userInformationProtectionPolicyClient,
		UserInformationProtectionPolicyLabel:                   userInformationProtectionPolicyLabelClient,
		UserInformationProtectionSensitivityLabel:              userInformationProtectionSensitivityLabelClient,
		UserInformationProtectionSensitivityLabelSublabel:      userInformationProtectionSensitivityLabelSublabelClient,
		UserInformationProtectionSensitivityPolicySetting:      userInformationProtectionSensitivityPolicySettingClient,
		UserInformationProtectionThreatAssessmentRequest:       userInformationProtectionThreatAssessmentRequestClient,
		UserInformationProtectionThreatAssessmentRequestResult: userInformationProtectionThreatAssessmentRequestResultClient,
		UserInsight:                                             userInsightClient,
		UserInsightShared:                                       userInsightSharedClient,
		UserInsightSharedLastSharedMethod:                       userInsightSharedLastSharedMethodClient,
		UserInsightSharedResource:                               userInsightSharedResourceClient,
		UserInsightTrending:                                     userInsightTrendingClient,
		UserInsightTrendingResource:                             userInsightTrendingResourceClient,
		UserInsightUsed:                                         userInsightUsedClient,
		UserInsightUsedResource:                                 userInsightUsedResourceClient,
		UserJoinedGroup:                                         userJoinedGroupClient,
		UserJoinedTeam:                                          userJoinedTeamClient,
		UserLicenseDetail:                                       userLicenseDetailClient,
		UserMailFolder:                                          userMailFolderClient,
		UserMailFolderChildFolder:                               userMailFolderChildFolderClient,
		UserMailFolderChildFolderMessage:                        userMailFolderChildFolderMessageClient,
		UserMailFolderChildFolderMessageAttachment:              userMailFolderChildFolderMessageAttachmentClient,
		UserMailFolderChildFolderMessageExtension:               userMailFolderChildFolderMessageExtensionClient,
		UserMailFolderChildFolderMessageMention:                 userMailFolderChildFolderMessageMentionClient,
		UserMailFolderChildFolderMessageRule:                    userMailFolderChildFolderMessageRuleClient,
		UserMailFolderChildFolderUserConfiguration:              userMailFolderChildFolderUserConfigurationClient,
		UserMailFolderMessage:                                   userMailFolderMessageClient,
		UserMailFolderMessageAttachment:                         userMailFolderMessageAttachmentClient,
		UserMailFolderMessageExtension:                          userMailFolderMessageExtensionClient,
		UserMailFolderMessageMention:                            userMailFolderMessageMentionClient,
		UserMailFolderMessageRule:                               userMailFolderMessageRuleClient,
		UserMailFolderUserConfiguration:                         userMailFolderUserConfigurationClient,
		UserMailboxSetting:                                      userMailboxSettingClient,
		UserManagedAppRegistration:                              userManagedAppRegistrationClient,
		UserManagedDevice:                                       userManagedDeviceClient,
		UserManagedDeviceAssignmentFilterEvaluationStatusDetail: userManagedDeviceAssignmentFilterEvaluationStatusDetailClient,
		UserManagedDeviceDetectedApp:                            userManagedDeviceDetectedAppClient,
		UserManagedDeviceDeviceCategory:                         userManagedDeviceDeviceCategoryClient,
		UserManagedDeviceDeviceCompliancePolicyState:            userManagedDeviceDeviceCompliancePolicyStateClient,
		UserManagedDeviceDeviceConfigurationState:               userManagedDeviceDeviceConfigurationStateClient,
		UserManagedDeviceDeviceHealthScriptState:                userManagedDeviceDeviceHealthScriptStateClient,
		UserManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId: userManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient,
		UserManagedDeviceLogCollectionRequest:                                        userManagedDeviceLogCollectionRequestClient,
		UserManagedDeviceManagedDeviceMobileAppConfigurationState:                    userManagedDeviceManagedDeviceMobileAppConfigurationStateClient,
		UserManagedDeviceSecurityBaselineState:                                       userManagedDeviceSecurityBaselineStateClient,
		UserManagedDeviceSecurityBaselineStateSettingState:                           userManagedDeviceSecurityBaselineStateSettingStateClient,
		UserManagedDeviceUser:                                                        userManagedDeviceUserClient,
		UserManagedDeviceWindowsProtectionState:                                      userManagedDeviceWindowsProtectionStateClient,
		UserManagedDeviceWindowsProtectionStateDetectedMalwareState:                  userManagedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		UserManager:                       userManagerClient,
		UserMemberOf:                      userMemberOfClient,
		UserMessage:                       userMessageClient,
		UserMessageAttachment:             userMessageAttachmentClient,
		UserMessageExtension:              userMessageExtensionClient,
		UserMessageMention:                userMessageMentionClient,
		UserMobileAppIntentAndState:       userMobileAppIntentAndStateClient,
		UserMobileAppTroubleshootingEvent: userMobileAppTroubleshootingEventClient,
		UserMobileAppTroubleshootingEventAppLogCollectionRequest: userMobileAppTroubleshootingEventAppLogCollectionRequestClient,
		UserNotification:                                         userNotificationClient,
		UserOauth2PermissionGrant:                                userOauth2PermissionGrantClient,
		UserOnenote:                                              userOnenoteClient,
		UserOnenoteNotebook:                                      userOnenoteNotebookClient,
		UserOnenoteNotebookSection:                               userOnenoteNotebookSectionClient,
		UserOnenoteNotebookSectionGroup:                          userOnenoteNotebookSectionGroupClient,
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
		UserOnlineMeetingAlternativeRecording:                    userOnlineMeetingAlternativeRecordingClient,
		UserOnlineMeetingAttendanceReport:                        userOnlineMeetingAttendanceReportClient,
		UserOnlineMeetingAttendanceReportAttendanceRecord:        userOnlineMeetingAttendanceReportAttendanceRecordClient,
		UserOnlineMeetingAttendeeReport:                          userOnlineMeetingAttendeeReportClient,
		UserOnlineMeetingBroadcastRecording:                      userOnlineMeetingBroadcastRecordingClient,
		UserOnlineMeetingMeetingAttendanceReport:                 userOnlineMeetingMeetingAttendanceReportClient,
		UserOnlineMeetingMeetingAttendanceReportAttendanceRecord: userOnlineMeetingMeetingAttendanceReportAttendanceRecordClient,
		UserOnlineMeetingRecording:                               userOnlineMeetingRecordingClient,
		UserOnlineMeetingRecordingContent:                        userOnlineMeetingRecordingContentClient,
		UserOnlineMeetingRegistration:                            userOnlineMeetingRegistrationClient,
		UserOnlineMeetingRegistrationCustomQuestion:              userOnlineMeetingRegistrationCustomQuestionClient,
		UserOnlineMeetingRegistrationRegistrant:                  userOnlineMeetingRegistrationRegistrantClient,
		UserOnlineMeetingTranscript:                              userOnlineMeetingTranscriptClient,
		UserOnlineMeetingTranscriptContent:                       userOnlineMeetingTranscriptContentClient,
		UserOnlineMeetingTranscriptMetadataContent:               userOnlineMeetingTranscriptMetadataContentClient,
		UserOutlook:                                                           userOutlookClient,
		UserOutlookMasterCategory:                                             userOutlookMasterCategoryClient,
		UserOutlookTask:                                                       userOutlookTaskClient,
		UserOutlookTaskAttachment:                                             userOutlookTaskAttachmentClient,
		UserOutlookTaskFolder:                                                 userOutlookTaskFolderClient,
		UserOutlookTaskFolderTask:                                             userOutlookTaskFolderTaskClient,
		UserOutlookTaskFolderTaskAttachment:                                   userOutlookTaskFolderTaskAttachmentClient,
		UserOutlookTaskGroup:                                                  userOutlookTaskGroupClient,
		UserOutlookTaskGroupTaskFolder:                                        userOutlookTaskGroupTaskFolderClient,
		UserOutlookTaskGroupTaskFolderTask:                                    userOutlookTaskGroupTaskFolderTaskClient,
		UserOutlookTaskGroupTaskFolderTaskAttachment:                          userOutlookTaskGroupTaskFolderTaskAttachmentClient,
		UserOwnedDevice:                                                       userOwnedDeviceClient,
		UserOwnedObject:                                                       userOwnedObjectClient,
		UserPendingAccessReviewInstance:                                       userPendingAccessReviewInstanceClient,
		UserPendingAccessReviewInstanceContactedReviewer:                      userPendingAccessReviewInstanceContactedReviewerClient,
		UserPendingAccessReviewInstanceDecision:                               userPendingAccessReviewInstanceDecisionClient,
		UserPendingAccessReviewInstanceDecisionInsight:                        userPendingAccessReviewInstanceDecisionInsightClient,
		UserPendingAccessReviewInstanceDecisionInstance:                       userPendingAccessReviewInstanceDecisionInstanceClient,
		UserPendingAccessReviewInstanceDecisionInstanceContactedReviewer:      userPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient,
		UserPendingAccessReviewInstanceDecisionInstanceDefinition:             userPendingAccessReviewInstanceDecisionInstanceDefinitionClient,
		UserPendingAccessReviewInstanceDecisionInstanceStage:                  userPendingAccessReviewInstanceDecisionInstanceStageClient,
		UserPendingAccessReviewInstanceDecisionInstanceStageDecision:          userPendingAccessReviewInstanceDecisionInstanceStageDecisionClient,
		UserPendingAccessReviewInstanceDecisionInstanceStageDecisionInsight:   userPendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient,
		UserPendingAccessReviewInstanceDefinition:                             userPendingAccessReviewInstanceDefinitionClient,
		UserPendingAccessReviewInstanceStage:                                  userPendingAccessReviewInstanceStageClient,
		UserPendingAccessReviewInstanceStageDecision:                          userPendingAccessReviewInstanceStageDecisionClient,
		UserPendingAccessReviewInstanceStageDecisionInsight:                   userPendingAccessReviewInstanceStageDecisionInsightClient,
		UserPendingAccessReviewInstanceStageDecisionInstance:                  userPendingAccessReviewInstanceStageDecisionInstanceClient,
		UserPendingAccessReviewInstanceStageDecisionInstanceContactedReviewer: userPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient,
		UserPendingAccessReviewInstanceStageDecisionInstanceDecision:          userPendingAccessReviewInstanceStageDecisionInstanceDecisionClient,
		UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsight:   userPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient,
		UserPendingAccessReviewInstanceStageDecisionInstanceDefinition:        userPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient,
		UserPeople:                userPeopleClient,
		UserPermissionGrant:       userPermissionGrantClient,
		UserPhoto:                 userPhotoClient,
		UserPlanner:               userPlannerClient,
		UserPlannerAll:            userPlannerAllClient,
		UserPlannerFavoritePlan:   userPlannerFavoritePlanClient,
		UserPlannerPlan:           userPlannerPlanClient,
		UserPlannerPlanBucket:     userPlannerPlanBucketClient,
		UserPlannerPlanBucketTask: userPlannerPlanBucketTaskClient,
		UserPlannerPlanBucketTaskAssignedToTaskBoardFormat:      userPlannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		UserPlannerPlanBucketTaskBucketTaskBoardFormat:          userPlannerPlanBucketTaskBucketTaskBoardFormatClient,
		UserPlannerPlanBucketTaskDetail:                         userPlannerPlanBucketTaskDetailClient,
		UserPlannerPlanBucketTaskProgressTaskBoardFormat:        userPlannerPlanBucketTaskProgressTaskBoardFormatClient,
		UserPlannerPlanDetail:                                   userPlannerPlanDetailClient,
		UserPlannerPlanTask:                                     userPlannerPlanTaskClient,
		UserPlannerPlanTaskAssignedToTaskBoardFormat:            userPlannerPlanTaskAssignedToTaskBoardFormatClient,
		UserPlannerPlanTaskBucketTaskBoardFormat:                userPlannerPlanTaskBucketTaskBoardFormatClient,
		UserPlannerPlanTaskDetail:                               userPlannerPlanTaskDetailClient,
		UserPlannerPlanTaskProgressTaskBoardFormat:              userPlannerPlanTaskProgressTaskBoardFormatClient,
		UserPlannerRecentPlan:                                   userPlannerRecentPlanClient,
		UserPlannerRosterPlan:                                   userPlannerRosterPlanClient,
		UserPlannerTask:                                         userPlannerTaskClient,
		UserPlannerTaskAssignedToTaskBoardFormat:                userPlannerTaskAssignedToTaskBoardFormatClient,
		UserPlannerTaskBucketTaskBoardFormat:                    userPlannerTaskBucketTaskBoardFormatClient,
		UserPlannerTaskDetail:                                   userPlannerTaskDetailClient,
		UserPlannerTaskProgressTaskBoardFormat:                  userPlannerTaskProgressTaskBoardFormatClient,
		UserPresence:                                            userPresenceClient,
		UserProfile:                                             userProfileClient,
		UserProfileAccount:                                      userProfileAccountClient,
		UserProfileAddress:                                      userProfileAddressClient,
		UserProfileAnniversary:                                  userProfileAnniversaryClient,
		UserProfileAward:                                        userProfileAwardClient,
		UserProfileCertification:                                userProfileCertificationClient,
		UserProfileEducationalActivity:                          userProfileEducationalActivityClient,
		UserProfileEmail:                                        userProfileEmailClient,
		UserProfileInterest:                                     userProfileInterestClient,
		UserProfileLanguage:                                     userProfileLanguageClient,
		UserProfileName:                                         userProfileNameClient,
		UserProfileNote:                                         userProfileNoteClient,
		UserProfilePatent:                                       userProfilePatentClient,
		UserProfilePhone:                                        userProfilePhoneClient,
		UserProfilePosition:                                     userProfilePositionClient,
		UserProfileProject:                                      userProfileProjectClient,
		UserProfilePublication:                                  userProfilePublicationClient,
		UserProfileSkill:                                        userProfileSkillClient,
		UserProfileWebAccount:                                   userProfileWebAccountClient,
		UserProfileWebsite:                                      userProfileWebsiteClient,
		UserRegisteredDevice:                                    userRegisteredDeviceClient,
		UserScopedRoleMemberOf:                                  userScopedRoleMemberOfClient,
		UserSecurity:                                            userSecurityClient,
		UserSecurityInformationProtection:                       userSecurityInformationProtectionClient,
		UserSecurityInformationProtectionLabelPolicySetting:     userSecurityInformationProtectionLabelPolicySettingClient,
		UserSecurityInformationProtectionSensitivityLabel:       userSecurityInformationProtectionSensitivityLabelClient,
		UserSecurityInformationProtectionSensitivityLabelParent: userSecurityInformationProtectionSensitivityLabelParentClient,
		UserSetting:                                             userSettingClient,
		UserSettingContactMergeSuggestion:                       userSettingContactMergeSuggestionClient,
		UserSettingItemInsight:                                  userSettingItemInsightClient,
		UserSettingRegionalAndLanguageSetting:                   userSettingRegionalAndLanguageSettingClient,
		UserSettingShiftPreference:                              userSettingShiftPreferenceClient,
		UserSponsor:                                             userSponsorClient,
		UserTeamwork:                                            userTeamworkClient,
		UserTeamworkAssociatedTeam:                              userTeamworkAssociatedTeamClient,
		UserTeamworkAssociatedTeamTeam:                          userTeamworkAssociatedTeamTeamClient,
		UserTeamworkInstalledApp:                                userTeamworkInstalledAppClient,
		UserTeamworkInstalledAppChat:                            userTeamworkInstalledAppChatClient,
		UserTeamworkInstalledAppTeamsApp:                        userTeamworkInstalledAppTeamsAppClient,
		UserTeamworkInstalledAppTeamsAppDefinition:              userTeamworkInstalledAppTeamsAppDefinitionClient,
		UserTodo:                                           userTodoClient,
		UserTodoList:                                       userTodoListClient,
		UserTodoListExtension:                              userTodoListExtensionClient,
		UserTodoListTask:                                   userTodoListTaskClient,
		UserTodoListTaskAttachment:                         userTodoListTaskAttachmentClient,
		UserTodoListTaskAttachmentSession:                  userTodoListTaskAttachmentSessionClient,
		UserTodoListTaskAttachmentSessionContent:           userTodoListTaskAttachmentSessionContentClient,
		UserTodoListTaskChecklistItem:                      userTodoListTaskChecklistItemClient,
		UserTodoListTaskExtension:                          userTodoListTaskExtensionClient,
		UserTodoListTaskLinkedResource:                     userTodoListTaskLinkedResourceClient,
		UserTransitiveMemberOf:                             userTransitiveMemberOfClient,
		UserTransitiveReport:                               userTransitiveReportClient,
		UserUsageRight:                                     userUsageRightClient,
		UserWindowsInformationProtectionDeviceRegistration: userWindowsInformationProtectionDeviceRegistrationClient,
	}, nil
}
