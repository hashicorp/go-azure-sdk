package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/me"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meactivityhistoryitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meactivityhistoryitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meagreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meanalytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meanalyticactivitystatistic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meappconsentrequestsforapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meappconsentrequestsforapprovaluserconsentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meappconsentrequestsforapprovaluserconsentrequestapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meappconsentrequestsforapprovaluserconsentrequestapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meapproleassignedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationemailmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationfido2method"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationpasswordlessmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationpasswordlessmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationpasswordmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationphonemethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationsoftwareoathmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationtemporaryaccesspassmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationwindowshelloforbusinessmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meauthenticationwindowshelloforbusinessmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendargroupcalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatlastmessagepreview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatpermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatpinnedmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechatpinnedmessagemessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechattab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mechattabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecloudpc"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfolderchildfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfolderchildfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfolderchildfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mecreatedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medevicecommand"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medevicecommandresponsepayload"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medeviceenrollmentconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medeviceenrollmentconfigurationassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medeviceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medevicemanagementtroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medevicememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medeviceregisteredowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medeviceregistereduser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medevicetransitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medeviceusageright"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medirectreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/medrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meemployeeexperience"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meemployeeexperiencelearningcourseactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meeventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mefollowedsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinferenceclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinferenceclassificationoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionbitlocker"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionbitlockerrecoverykey"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectiondatalosspreventionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionpolicylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionsensitivitylabelsublabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionsensitivitypolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionthreatassessmentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinformationprotectionthreatassessmentrequestresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsightshared"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsightsharedlastsharedmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsightsharedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsighttrending"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsighttrendingresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsightused"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meinsightusedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mejoinedgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mejoinedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/melicensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderchildfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderchildfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderchildfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderchildfoldermessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderchildfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderchildfolderuserconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfoldermessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memailfolderuserconfiguration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanagedappregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddeviceassignmentfilterevaluationstatusdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicedetectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicedevicehealthscriptstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicemanageddevicemobileappconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicesecuritybaselinestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicesecuritybaselinestatesettingstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memanager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memessagemention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memobileappintentandstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memobileapptroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/memobileapptroubleshootingeventapplogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/menotification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotenotebooksectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenoteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotepagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotepageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotepageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenoteresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenoteresourcecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonenotesectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeeting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingalternativerecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingattendeereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingbroadcastrecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingmeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingmeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingrecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingrecordingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingregistrationcustomquestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingregistrationregistrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingtranscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingtranscriptcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meonlinemeetingtranscriptmetadatacontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlookmastercategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskfoldertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskfoldertaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskgrouptaskfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskgrouptaskfoldertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meoutlooktaskgrouptaskfoldertaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meowneddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecisioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecisioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecisioninstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecisioninstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecisioninstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedecisioninstancestagedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestagedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestagedecisioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestagedecisioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestagedecisioninstancedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestagedecisioninstancedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mependingaccessreviewinstancestagedecisioninstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mepeople"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mepermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mephoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplanner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerall"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerfavoriteplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerrecentplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannerrosterplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannertaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannertaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannertaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meplannertaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mepresence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileaccount"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileaddress"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileanniversary"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileaward"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilecertification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileeducationalactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileemail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileinterest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilelanguage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilename"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilepatent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilephone"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileproject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilepublication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofileskill"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilewebaccount"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meprofilewebsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meregistereddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mescopedrolememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesecurity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesecurityinformationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesecurityinformationprotectionlabelpolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesecurityinformationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesecurityinformationprotectionsensitivitylabelparent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesettingcontactmergesuggestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesettingiteminsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesettingregionalandlanguagesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesettingshiftpreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mesponsor"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meteamwork"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meteamworkassociatedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meteamworkassociatedteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meteamworkinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meteamworkinstalledappchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meteamworkinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meteamworkinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolistextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolisttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolisttaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolisttaskattachmentsession"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolisttaskattachmentsessioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolisttaskchecklistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolisttaskextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metodolisttasklinkedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metransitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/metransitivereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/meusageright"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mewindowsinformationprotectiondeviceregistration"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Me                                                                         *me.MeClient
	MeActivity                                                                 *meactivity.MeActivityClient
	MeActivityHistoryItem                                                      *meactivityhistoryitem.MeActivityHistoryItemClient
	MeActivityHistoryItemActivity                                              *meactivityhistoryitemactivity.MeActivityHistoryItemActivityClient
	MeAgreementAcceptance                                                      *meagreementacceptance.MeAgreementAcceptanceClient
	MeAnalytic                                                                 *meanalytic.MeAnalyticClient
	MeAnalyticActivityStatistic                                                *meanalyticactivitystatistic.MeAnalyticActivityStatisticClient
	MeAppConsentRequestsForApproval                                            *meappconsentrequestsforapproval.MeAppConsentRequestsForApprovalClient
	MeAppConsentRequestsForApprovalUserConsentRequest                          *meappconsentrequestsforapprovaluserconsentrequest.MeAppConsentRequestsForApprovalUserConsentRequestClient
	MeAppConsentRequestsForApprovalUserConsentRequestApproval                  *meappconsentrequestsforapprovaluserconsentrequestapproval.MeAppConsentRequestsForApprovalUserConsentRequestApprovalClient
	MeAppConsentRequestsForApprovalUserConsentRequestApprovalStep              *meappconsentrequestsforapprovaluserconsentrequestapprovalstep.MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient
	MeAppRoleAssignedResource                                                  *meapproleassignedresource.MeAppRoleAssignedResourceClient
	MeAppRoleAssignment                                                        *meapproleassignment.MeAppRoleAssignmentClient
	MeApproval                                                                 *meapproval.MeApprovalClient
	MeApprovalStep                                                             *meapprovalstep.MeApprovalStepClient
	MeAuthentication                                                           *meauthentication.MeAuthenticationClient
	MeAuthenticationEmailMethod                                                *meauthenticationemailmethod.MeAuthenticationEmailMethodClient
	MeAuthenticationFido2Method                                                *meauthenticationfido2method.MeAuthenticationFido2MethodClient
	MeAuthenticationMethod                                                     *meauthenticationmethod.MeAuthenticationMethodClient
	MeAuthenticationMicrosoftAuthenticatorMethod                               *meauthenticationmicrosoftauthenticatormethod.MeAuthenticationMicrosoftAuthenticatorMethodClient
	MeAuthenticationMicrosoftAuthenticatorMethodDevice                         *meauthenticationmicrosoftauthenticatormethoddevice.MeAuthenticationMicrosoftAuthenticatorMethodDeviceClient
	MeAuthenticationOperation                                                  *meauthenticationoperation.MeAuthenticationOperationClient
	MeAuthenticationPasswordMethod                                             *meauthenticationpasswordmethod.MeAuthenticationPasswordMethodClient
	MeAuthenticationPasswordlessMicrosoftAuthenticatorMethod                   *meauthenticationpasswordlessmicrosoftauthenticatormethod.MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient
	MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice             *meauthenticationpasswordlessmicrosoftauthenticatormethoddevice.MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient
	MeAuthenticationPhoneMethod                                                *meauthenticationphonemethod.MeAuthenticationPhoneMethodClient
	MeAuthenticationSoftwareOathMethod                                         *meauthenticationsoftwareoathmethod.MeAuthenticationSoftwareOathMethodClient
	MeAuthenticationTemporaryAccessPassMethod                                  *meauthenticationtemporaryaccesspassmethod.MeAuthenticationTemporaryAccessPassMethodClient
	MeAuthenticationWindowsHelloForBusinessMethod                              *meauthenticationwindowshelloforbusinessmethod.MeAuthenticationWindowsHelloForBusinessMethodClient
	MeAuthenticationWindowsHelloForBusinessMethodDevice                        *meauthenticationwindowshelloforbusinessmethoddevice.MeAuthenticationWindowsHelloForBusinessMethodDeviceClient
	MeCalendar                                                                 *mecalendar.MeCalendarClient
	MeCalendarCalendarPermission                                               *mecalendarcalendarpermission.MeCalendarCalendarPermissionClient
	MeCalendarCalendarView                                                     *mecalendarcalendarview.MeCalendarCalendarViewClient
	MeCalendarCalendarViewAttachment                                           *mecalendarcalendarviewattachment.MeCalendarCalendarViewAttachmentClient
	MeCalendarCalendarViewCalendar                                             *mecalendarcalendarviewcalendar.MeCalendarCalendarViewCalendarClient
	MeCalendarCalendarViewExceptionOccurrence                                  *mecalendarcalendarviewexceptionoccurrence.MeCalendarCalendarViewExceptionOccurrenceClient
	MeCalendarCalendarViewExceptionOccurrenceAttachment                        *mecalendarcalendarviewexceptionoccurrenceattachment.MeCalendarCalendarViewExceptionOccurrenceAttachmentClient
	MeCalendarCalendarViewExceptionOccurrenceCalendar                          *mecalendarcalendarviewexceptionoccurrencecalendar.MeCalendarCalendarViewExceptionOccurrenceCalendarClient
	MeCalendarCalendarViewExceptionOccurrenceExtension                         *mecalendarcalendarviewexceptionoccurrenceextension.MeCalendarCalendarViewExceptionOccurrenceExtensionClient
	MeCalendarCalendarViewExceptionOccurrenceInstance                          *mecalendarcalendarviewexceptionoccurrenceinstance.MeCalendarCalendarViewExceptionOccurrenceInstanceClient
	MeCalendarCalendarViewExceptionOccurrenceInstanceAttachment                *mecalendarcalendarviewexceptionoccurrenceinstanceattachment.MeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	MeCalendarCalendarViewExceptionOccurrenceInstanceCalendar                  *mecalendarcalendarviewexceptionoccurrenceinstancecalendar.MeCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	MeCalendarCalendarViewExceptionOccurrenceInstanceExtension                 *mecalendarcalendarviewexceptionoccurrenceinstanceextension.MeCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	MeCalendarCalendarViewExtension                                            *mecalendarcalendarviewextension.MeCalendarCalendarViewExtensionClient
	MeCalendarCalendarViewInstance                                             *mecalendarcalendarviewinstance.MeCalendarCalendarViewInstanceClient
	MeCalendarCalendarViewInstanceAttachment                                   *mecalendarcalendarviewinstanceattachment.MeCalendarCalendarViewInstanceAttachmentClient
	MeCalendarCalendarViewInstanceCalendar                                     *mecalendarcalendarviewinstancecalendar.MeCalendarCalendarViewInstanceCalendarClient
	MeCalendarCalendarViewInstanceExceptionOccurrence                          *mecalendarcalendarviewinstanceexceptionoccurrence.MeCalendarCalendarViewInstanceExceptionOccurrenceClient
	MeCalendarCalendarViewInstanceExceptionOccurrenceAttachment                *mecalendarcalendarviewinstanceexceptionoccurrenceattachment.MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	MeCalendarCalendarViewInstanceExceptionOccurrenceCalendar                  *mecalendarcalendarviewinstanceexceptionoccurrencecalendar.MeCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	MeCalendarCalendarViewInstanceExceptionOccurrenceExtension                 *mecalendarcalendarviewinstanceexceptionoccurrenceextension.MeCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	MeCalendarCalendarViewInstanceExtension                                    *mecalendarcalendarviewinstanceextension.MeCalendarCalendarViewInstanceExtensionClient
	MeCalendarEvent                                                            *mecalendarevent.MeCalendarEventClient
	MeCalendarEventAttachment                                                  *mecalendareventattachment.MeCalendarEventAttachmentClient
	MeCalendarEventCalendar                                                    *mecalendareventcalendar.MeCalendarEventCalendarClient
	MeCalendarEventExceptionOccurrence                                         *mecalendareventexceptionoccurrence.MeCalendarEventExceptionOccurrenceClient
	MeCalendarEventExceptionOccurrenceAttachment                               *mecalendareventexceptionoccurrenceattachment.MeCalendarEventExceptionOccurrenceAttachmentClient
	MeCalendarEventExceptionOccurrenceCalendar                                 *mecalendareventexceptionoccurrencecalendar.MeCalendarEventExceptionOccurrenceCalendarClient
	MeCalendarEventExceptionOccurrenceExtension                                *mecalendareventexceptionoccurrenceextension.MeCalendarEventExceptionOccurrenceExtensionClient
	MeCalendarEventExceptionOccurrenceInstance                                 *mecalendareventexceptionoccurrenceinstance.MeCalendarEventExceptionOccurrenceInstanceClient
	MeCalendarEventExceptionOccurrenceInstanceAttachment                       *mecalendareventexceptionoccurrenceinstanceattachment.MeCalendarEventExceptionOccurrenceInstanceAttachmentClient
	MeCalendarEventExceptionOccurrenceInstanceCalendar                         *mecalendareventexceptionoccurrenceinstancecalendar.MeCalendarEventExceptionOccurrenceInstanceCalendarClient
	MeCalendarEventExceptionOccurrenceInstanceExtension                        *mecalendareventexceptionoccurrenceinstanceextension.MeCalendarEventExceptionOccurrenceInstanceExtensionClient
	MeCalendarEventExtension                                                   *mecalendareventextension.MeCalendarEventExtensionClient
	MeCalendarEventInstance                                                    *mecalendareventinstance.MeCalendarEventInstanceClient
	MeCalendarEventInstanceAttachment                                          *mecalendareventinstanceattachment.MeCalendarEventInstanceAttachmentClient
	MeCalendarEventInstanceCalendar                                            *mecalendareventinstancecalendar.MeCalendarEventInstanceCalendarClient
	MeCalendarEventInstanceExceptionOccurrence                                 *mecalendareventinstanceexceptionoccurrence.MeCalendarEventInstanceExceptionOccurrenceClient
	MeCalendarEventInstanceExceptionOccurrenceAttachment                       *mecalendareventinstanceexceptionoccurrenceattachment.MeCalendarEventInstanceExceptionOccurrenceAttachmentClient
	MeCalendarEventInstanceExceptionOccurrenceCalendar                         *mecalendareventinstanceexceptionoccurrencecalendar.MeCalendarEventInstanceExceptionOccurrenceCalendarClient
	MeCalendarEventInstanceExceptionOccurrenceExtension                        *mecalendareventinstanceexceptionoccurrenceextension.MeCalendarEventInstanceExceptionOccurrenceExtensionClient
	MeCalendarEventInstanceExtension                                           *mecalendareventinstanceextension.MeCalendarEventInstanceExtensionClient
	MeCalendarGroup                                                            *mecalendargroup.MeCalendarGroupClient
	MeCalendarGroupCalendar                                                    *mecalendargroupcalendar.MeCalendarGroupCalendarClient
	MeCalendarGroupCalendarCalendarPermission                                  *mecalendargroupcalendarcalendarpermission.MeCalendarGroupCalendarCalendarPermissionClient
	MeCalendarGroupCalendarCalendarView                                        *mecalendargroupcalendarcalendarview.MeCalendarGroupCalendarCalendarViewClient
	MeCalendarGroupCalendarCalendarViewAttachment                              *mecalendargroupcalendarcalendarviewattachment.MeCalendarGroupCalendarCalendarViewAttachmentClient
	MeCalendarGroupCalendarCalendarViewCalendar                                *mecalendargroupcalendarcalendarviewcalendar.MeCalendarGroupCalendarCalendarViewCalendarClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrence                     *mecalendargroupcalendarcalendarviewexceptionoccurrence.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment           *mecalendargroupcalendarcalendarviewexceptionoccurrenceattachment.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar             *mecalendargroupcalendarcalendarviewexceptionoccurrencecalendar.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtension            *mecalendargroupcalendarcalendarviewexceptionoccurrenceextension.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance             *mecalendargroupcalendarcalendarviewexceptionoccurrenceinstance.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment   *mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar     *mecalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension    *mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension.MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	MeCalendarGroupCalendarCalendarViewExtension                               *mecalendargroupcalendarcalendarviewextension.MeCalendarGroupCalendarCalendarViewExtensionClient
	MeCalendarGroupCalendarCalendarViewInstance                                *mecalendargroupcalendarcalendarviewinstance.MeCalendarGroupCalendarCalendarViewInstanceClient
	MeCalendarGroupCalendarCalendarViewInstanceAttachment                      *mecalendargroupcalendarcalendarviewinstanceattachment.MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient
	MeCalendarGroupCalendarCalendarViewInstanceCalendar                        *mecalendargroupcalendarcalendarviewinstancecalendar.MeCalendarGroupCalendarCalendarViewInstanceCalendarClient
	MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence             *mecalendargroupcalendarcalendarviewinstanceexceptionoccurrence.MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient
	MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment   *mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar     *mecalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar.MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension    *mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension.MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	MeCalendarGroupCalendarCalendarViewInstanceExtension                       *mecalendargroupcalendarcalendarviewinstanceextension.MeCalendarGroupCalendarCalendarViewInstanceExtensionClient
	MeCalendarGroupCalendarEvent                                               *mecalendargroupcalendarevent.MeCalendarGroupCalendarEventClient
	MeCalendarGroupCalendarEventAttachment                                     *mecalendargroupcalendareventattachment.MeCalendarGroupCalendarEventAttachmentClient
	MeCalendarGroupCalendarEventCalendar                                       *mecalendargroupcalendareventcalendar.MeCalendarGroupCalendarEventCalendarClient
	MeCalendarGroupCalendarEventExceptionOccurrence                            *mecalendargroupcalendareventexceptionoccurrence.MeCalendarGroupCalendarEventExceptionOccurrenceClient
	MeCalendarGroupCalendarEventExceptionOccurrenceAttachment                  *mecalendargroupcalendareventexceptionoccurrenceattachment.MeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient
	MeCalendarGroupCalendarEventExceptionOccurrenceCalendar                    *mecalendargroupcalendareventexceptionoccurrencecalendar.MeCalendarGroupCalendarEventExceptionOccurrenceCalendarClient
	MeCalendarGroupCalendarEventExceptionOccurrenceExtension                   *mecalendargroupcalendareventexceptionoccurrenceextension.MeCalendarGroupCalendarEventExceptionOccurrenceExtensionClient
	MeCalendarGroupCalendarEventExceptionOccurrenceInstance                    *mecalendargroupcalendareventexceptionoccurrenceinstance.MeCalendarGroupCalendarEventExceptionOccurrenceInstanceClient
	MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment          *mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient
	MeCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar            *mecalendargroupcalendareventexceptionoccurrenceinstancecalendar.MeCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient
	MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtension           *mecalendargroupcalendareventexceptionoccurrenceinstanceextension.MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient
	MeCalendarGroupCalendarEventExtension                                      *mecalendargroupcalendareventextension.MeCalendarGroupCalendarEventExtensionClient
	MeCalendarGroupCalendarEventInstance                                       *mecalendargroupcalendareventinstance.MeCalendarGroupCalendarEventInstanceClient
	MeCalendarGroupCalendarEventInstanceAttachment                             *mecalendargroupcalendareventinstanceattachment.MeCalendarGroupCalendarEventInstanceAttachmentClient
	MeCalendarGroupCalendarEventInstanceCalendar                               *mecalendargroupcalendareventinstancecalendar.MeCalendarGroupCalendarEventInstanceCalendarClient
	MeCalendarGroupCalendarEventInstanceExceptionOccurrence                    *mecalendargroupcalendareventinstanceexceptionoccurrence.MeCalendarGroupCalendarEventInstanceExceptionOccurrenceClient
	MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment          *mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient
	MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar            *mecalendargroupcalendareventinstanceexceptionoccurrencecalendar.MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient
	MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension           *mecalendargroupcalendareventinstanceexceptionoccurrenceextension.MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient
	MeCalendarGroupCalendarEventInstanceExtension                              *mecalendargroupcalendareventinstanceextension.MeCalendarGroupCalendarEventInstanceExtensionClient
	MeCalendarView                                                             *mecalendarview.MeCalendarViewClient
	MeCalendarViewAttachment                                                   *mecalendarviewattachment.MeCalendarViewAttachmentClient
	MeCalendarViewCalendar                                                     *mecalendarviewcalendar.MeCalendarViewCalendarClient
	MeCalendarViewExceptionOccurrence                                          *mecalendarviewexceptionoccurrence.MeCalendarViewExceptionOccurrenceClient
	MeCalendarViewExceptionOccurrenceAttachment                                *mecalendarviewexceptionoccurrenceattachment.MeCalendarViewExceptionOccurrenceAttachmentClient
	MeCalendarViewExceptionOccurrenceCalendar                                  *mecalendarviewexceptionoccurrencecalendar.MeCalendarViewExceptionOccurrenceCalendarClient
	MeCalendarViewExceptionOccurrenceExtension                                 *mecalendarviewexceptionoccurrenceextension.MeCalendarViewExceptionOccurrenceExtensionClient
	MeCalendarViewExceptionOccurrenceInstance                                  *mecalendarviewexceptionoccurrenceinstance.MeCalendarViewExceptionOccurrenceInstanceClient
	MeCalendarViewExceptionOccurrenceInstanceAttachment                        *mecalendarviewexceptionoccurrenceinstanceattachment.MeCalendarViewExceptionOccurrenceInstanceAttachmentClient
	MeCalendarViewExceptionOccurrenceInstanceCalendar                          *mecalendarviewexceptionoccurrenceinstancecalendar.MeCalendarViewExceptionOccurrenceInstanceCalendarClient
	MeCalendarViewExceptionOccurrenceInstanceExtension                         *mecalendarviewexceptionoccurrenceinstanceextension.MeCalendarViewExceptionOccurrenceInstanceExtensionClient
	MeCalendarViewExtension                                                    *mecalendarviewextension.MeCalendarViewExtensionClient
	MeCalendarViewInstance                                                     *mecalendarviewinstance.MeCalendarViewInstanceClient
	MeCalendarViewInstanceAttachment                                           *mecalendarviewinstanceattachment.MeCalendarViewInstanceAttachmentClient
	MeCalendarViewInstanceCalendar                                             *mecalendarviewinstancecalendar.MeCalendarViewInstanceCalendarClient
	MeCalendarViewInstanceExceptionOccurrence                                  *mecalendarviewinstanceexceptionoccurrence.MeCalendarViewInstanceExceptionOccurrenceClient
	MeCalendarViewInstanceExceptionOccurrenceAttachment                        *mecalendarviewinstanceexceptionoccurrenceattachment.MeCalendarViewInstanceExceptionOccurrenceAttachmentClient
	MeCalendarViewInstanceExceptionOccurrenceCalendar                          *mecalendarviewinstanceexceptionoccurrencecalendar.MeCalendarViewInstanceExceptionOccurrenceCalendarClient
	MeCalendarViewInstanceExceptionOccurrenceExtension                         *mecalendarviewinstanceexceptionoccurrenceextension.MeCalendarViewInstanceExceptionOccurrenceExtensionClient
	MeCalendarViewInstanceExtension                                            *mecalendarviewinstanceextension.MeCalendarViewInstanceExtensionClient
	MeChat                                                                     *mechat.MeChatClient
	MeChatInstalledApp                                                         *mechatinstalledapp.MeChatInstalledAppClient
	MeChatInstalledAppTeamsApp                                                 *mechatinstalledappteamsapp.MeChatInstalledAppTeamsAppClient
	MeChatInstalledAppTeamsAppDefinition                                       *mechatinstalledappteamsappdefinition.MeChatInstalledAppTeamsAppDefinitionClient
	MeChatLastMessagePreview                                                   *mechatlastmessagepreview.MeChatLastMessagePreviewClient
	MeChatMember                                                               *mechatmember.MeChatMemberClient
	MeChatMessage                                                              *mechatmessage.MeChatMessageClient
	MeChatMessageHostedContent                                                 *mechatmessagehostedcontent.MeChatMessageHostedContentClient
	MeChatMessageReply                                                         *mechatmessagereply.MeChatMessageReplyClient
	MeChatMessageReplyHostedContent                                            *mechatmessagereplyhostedcontent.MeChatMessageReplyHostedContentClient
	MeChatOperation                                                            *mechatoperation.MeChatOperationClient
	MeChatPermissionGrant                                                      *mechatpermissiongrant.MeChatPermissionGrantClient
	MeChatPinnedMessage                                                        *mechatpinnedmessage.MeChatPinnedMessageClient
	MeChatPinnedMessageMessage                                                 *mechatpinnedmessagemessage.MeChatPinnedMessageMessageClient
	MeChatTab                                                                  *mechattab.MeChatTabClient
	MeChatTabTeamsApp                                                          *mechattabteamsapp.MeChatTabTeamsAppClient
	MeCloudPC                                                                  *mecloudpc.MeCloudPCClient
	MeContact                                                                  *mecontact.MeContactClient
	MeContactExtension                                                         *mecontactextension.MeContactExtensionClient
	MeContactFolder                                                            *mecontactfolder.MeContactFolderClient
	MeContactFolderChildFolder                                                 *mecontactfolderchildfolder.MeContactFolderChildFolderClient
	MeContactFolderChildFolderContact                                          *mecontactfolderchildfoldercontact.MeContactFolderChildFolderContactClient
	MeContactFolderChildFolderContactExtension                                 *mecontactfolderchildfoldercontactextension.MeContactFolderChildFolderContactExtensionClient
	MeContactFolderChildFolderContactPhoto                                     *mecontactfolderchildfoldercontactphoto.MeContactFolderChildFolderContactPhotoClient
	MeContactFolderContact                                                     *mecontactfoldercontact.MeContactFolderContactClient
	MeContactFolderContactExtension                                            *mecontactfoldercontactextension.MeContactFolderContactExtensionClient
	MeContactFolderContactPhoto                                                *mecontactfoldercontactphoto.MeContactFolderContactPhotoClient
	MeContactPhoto                                                             *mecontactphoto.MeContactPhotoClient
	MeCreatedObject                                                            *mecreatedobject.MeCreatedObjectClient
	MeDevice                                                                   *medevice.MeDeviceClient
	MeDeviceCommand                                                            *medevicecommand.MeDeviceCommandClient
	MeDeviceCommandResponsepayload                                             *medevicecommandresponsepayload.MeDeviceCommandResponsepayloadClient
	MeDeviceEnrollmentConfiguration                                            *medeviceenrollmentconfiguration.MeDeviceEnrollmentConfigurationClient
	MeDeviceEnrollmentConfigurationAssignment                                  *medeviceenrollmentconfigurationassignment.MeDeviceEnrollmentConfigurationAssignmentClient
	MeDeviceExtension                                                          *medeviceextension.MeDeviceExtensionClient
	MeDeviceManagementTroubleshootingEvent                                     *medevicemanagementtroubleshootingevent.MeDeviceManagementTroubleshootingEventClient
	MeDeviceMemberOf                                                           *medevicememberof.MeDeviceMemberOfClient
	MeDeviceRegisteredOwner                                                    *medeviceregisteredowner.MeDeviceRegisteredOwnerClient
	MeDeviceRegisteredUser                                                     *medeviceregistereduser.MeDeviceRegisteredUserClient
	MeDeviceTransitiveMemberOf                                                 *medevicetransitivememberof.MeDeviceTransitiveMemberOfClient
	MeDeviceUsageRight                                                         *medeviceusageright.MeDeviceUsageRightClient
	MeDirectReport                                                             *medirectreport.MeDirectReportClient
	MeDrive                                                                    *medrive.MeDriveClient
	MeEmployeeExperience                                                       *meemployeeexperience.MeEmployeeExperienceClient
	MeEmployeeExperienceLearningCourseActivity                                 *meemployeeexperiencelearningcourseactivity.MeEmployeeExperienceLearningCourseActivityClient
	MeEvent                                                                    *meevent.MeEventClient
	MeEventAttachment                                                          *meeventattachment.MeEventAttachmentClient
	MeEventCalendar                                                            *meeventcalendar.MeEventCalendarClient
	MeEventExceptionOccurrence                                                 *meeventexceptionoccurrence.MeEventExceptionOccurrenceClient
	MeEventExceptionOccurrenceAttachment                                       *meeventexceptionoccurrenceattachment.MeEventExceptionOccurrenceAttachmentClient
	MeEventExceptionOccurrenceCalendar                                         *meeventexceptionoccurrencecalendar.MeEventExceptionOccurrenceCalendarClient
	MeEventExceptionOccurrenceExtension                                        *meeventexceptionoccurrenceextension.MeEventExceptionOccurrenceExtensionClient
	MeEventExceptionOccurrenceInstance                                         *meeventexceptionoccurrenceinstance.MeEventExceptionOccurrenceInstanceClient
	MeEventExceptionOccurrenceInstanceAttachment                               *meeventexceptionoccurrenceinstanceattachment.MeEventExceptionOccurrenceInstanceAttachmentClient
	MeEventExceptionOccurrenceInstanceCalendar                                 *meeventexceptionoccurrenceinstancecalendar.MeEventExceptionOccurrenceInstanceCalendarClient
	MeEventExceptionOccurrenceInstanceExtension                                *meeventexceptionoccurrenceinstanceextension.MeEventExceptionOccurrenceInstanceExtensionClient
	MeEventExtension                                                           *meeventextension.MeEventExtensionClient
	MeEventInstance                                                            *meeventinstance.MeEventInstanceClient
	MeEventInstanceAttachment                                                  *meeventinstanceattachment.MeEventInstanceAttachmentClient
	MeEventInstanceCalendar                                                    *meeventinstancecalendar.MeEventInstanceCalendarClient
	MeEventInstanceExceptionOccurrence                                         *meeventinstanceexceptionoccurrence.MeEventInstanceExceptionOccurrenceClient
	MeEventInstanceExceptionOccurrenceAttachment                               *meeventinstanceexceptionoccurrenceattachment.MeEventInstanceExceptionOccurrenceAttachmentClient
	MeEventInstanceExceptionOccurrenceCalendar                                 *meeventinstanceexceptionoccurrencecalendar.MeEventInstanceExceptionOccurrenceCalendarClient
	MeEventInstanceExceptionOccurrenceExtension                                *meeventinstanceexceptionoccurrenceextension.MeEventInstanceExceptionOccurrenceExtensionClient
	MeEventInstanceExtension                                                   *meeventinstanceextension.MeEventInstanceExtensionClient
	MeExtension                                                                *meextension.MeExtensionClient
	MeFollowedSite                                                             *mefollowedsite.MeFollowedSiteClient
	MeInferenceClassification                                                  *meinferenceclassification.MeInferenceClassificationClient
	MeInferenceClassificationOverride                                          *meinferenceclassificationoverride.MeInferenceClassificationOverrideClient
	MeInformationProtection                                                    *meinformationprotection.MeInformationProtectionClient
	MeInformationProtectionBitlocker                                           *meinformationprotectionbitlocker.MeInformationProtectionBitlockerClient
	MeInformationProtectionBitlockerRecoveryKey                                *meinformationprotectionbitlockerrecoverykey.MeInformationProtectionBitlockerRecoveryKeyClient
	MeInformationProtectionDataLossPreventionPolicy                            *meinformationprotectiondatalosspreventionpolicy.MeInformationProtectionDataLossPreventionPolicyClient
	MeInformationProtectionPolicy                                              *meinformationprotectionpolicy.MeInformationProtectionPolicyClient
	MeInformationProtectionPolicyLabel                                         *meinformationprotectionpolicylabel.MeInformationProtectionPolicyLabelClient
	MeInformationProtectionSensitivityLabel                                    *meinformationprotectionsensitivitylabel.MeInformationProtectionSensitivityLabelClient
	MeInformationProtectionSensitivityLabelSublabel                            *meinformationprotectionsensitivitylabelsublabel.MeInformationProtectionSensitivityLabelSublabelClient
	MeInformationProtectionSensitivityPolicySetting                            *meinformationprotectionsensitivitypolicysetting.MeInformationProtectionSensitivityPolicySettingClient
	MeInformationProtectionThreatAssessmentRequest                             *meinformationprotectionthreatassessmentrequest.MeInformationProtectionThreatAssessmentRequestClient
	MeInformationProtectionThreatAssessmentRequestResult                       *meinformationprotectionthreatassessmentrequestresult.MeInformationProtectionThreatAssessmentRequestResultClient
	MeInsight                                                                  *meinsight.MeInsightClient
	MeInsightShared                                                            *meinsightshared.MeInsightSharedClient
	MeInsightSharedLastSharedMethod                                            *meinsightsharedlastsharedmethod.MeInsightSharedLastSharedMethodClient
	MeInsightSharedResource                                                    *meinsightsharedresource.MeInsightSharedResourceClient
	MeInsightTrending                                                          *meinsighttrending.MeInsightTrendingClient
	MeInsightTrendingResource                                                  *meinsighttrendingresource.MeInsightTrendingResourceClient
	MeInsightUsed                                                              *meinsightused.MeInsightUsedClient
	MeInsightUsedResource                                                      *meinsightusedresource.MeInsightUsedResourceClient
	MeJoinedGroup                                                              *mejoinedgroup.MeJoinedGroupClient
	MeJoinedTeam                                                               *mejoinedteam.MeJoinedTeamClient
	MeLicenseDetail                                                            *melicensedetail.MeLicenseDetailClient
	MeMailFolder                                                               *memailfolder.MeMailFolderClient
	MeMailFolderChildFolder                                                    *memailfolderchildfolder.MeMailFolderChildFolderClient
	MeMailFolderChildFolderMessage                                             *memailfolderchildfoldermessage.MeMailFolderChildFolderMessageClient
	MeMailFolderChildFolderMessageAttachment                                   *memailfolderchildfoldermessageattachment.MeMailFolderChildFolderMessageAttachmentClient
	MeMailFolderChildFolderMessageExtension                                    *memailfolderchildfoldermessageextension.MeMailFolderChildFolderMessageExtensionClient
	MeMailFolderChildFolderMessageMention                                      *memailfolderchildfoldermessagemention.MeMailFolderChildFolderMessageMentionClient
	MeMailFolderChildFolderMessageRule                                         *memailfolderchildfoldermessagerule.MeMailFolderChildFolderMessageRuleClient
	MeMailFolderChildFolderUserConfiguration                                   *memailfolderchildfolderuserconfiguration.MeMailFolderChildFolderUserConfigurationClient
	MeMailFolderMessage                                                        *memailfoldermessage.MeMailFolderMessageClient
	MeMailFolderMessageAttachment                                              *memailfoldermessageattachment.MeMailFolderMessageAttachmentClient
	MeMailFolderMessageExtension                                               *memailfoldermessageextension.MeMailFolderMessageExtensionClient
	MeMailFolderMessageMention                                                 *memailfoldermessagemention.MeMailFolderMessageMentionClient
	MeMailFolderMessageRule                                                    *memailfoldermessagerule.MeMailFolderMessageRuleClient
	MeMailFolderUserConfiguration                                              *memailfolderuserconfiguration.MeMailFolderUserConfigurationClient
	MeMailboxSetting                                                           *memailboxsetting.MeMailboxSettingClient
	MeManagedAppRegistration                                                   *memanagedappregistration.MeManagedAppRegistrationClient
	MeManagedDevice                                                            *memanageddevice.MeManagedDeviceClient
	MeManagedDeviceAssignmentFilterEvaluationStatusDetail                      *memanageddeviceassignmentfilterevaluationstatusdetail.MeManagedDeviceAssignmentFilterEvaluationStatusDetailClient
	MeManagedDeviceDetectedApp                                                 *memanageddevicedetectedapp.MeManagedDeviceDetectedAppClient
	MeManagedDeviceDeviceCategory                                              *memanageddevicedevicecategory.MeManagedDeviceDeviceCategoryClient
	MeManagedDeviceDeviceCompliancePolicyState                                 *memanageddevicedevicecompliancepolicystate.MeManagedDeviceDeviceCompliancePolicyStateClient
	MeManagedDeviceDeviceConfigurationState                                    *memanageddevicedeviceconfigurationstate.MeManagedDeviceDeviceConfigurationStateClient
	MeManagedDeviceDeviceHealthScriptState                                     *memanageddevicedevicehealthscriptstate.MeManagedDeviceDeviceHealthScriptStateClient
	MeManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId *memanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid.MeManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient
	MeManagedDeviceLogCollectionRequest                                        *memanageddevicelogcollectionrequest.MeManagedDeviceLogCollectionRequestClient
	MeManagedDeviceManagedDeviceMobileAppConfigurationState                    *memanageddevicemanageddevicemobileappconfigurationstate.MeManagedDeviceManagedDeviceMobileAppConfigurationStateClient
	MeManagedDeviceSecurityBaselineState                                       *memanageddevicesecuritybaselinestate.MeManagedDeviceSecurityBaselineStateClient
	MeManagedDeviceSecurityBaselineStateSettingState                           *memanageddevicesecuritybaselinestatesettingstate.MeManagedDeviceSecurityBaselineStateSettingStateClient
	MeManagedDeviceUser                                                        *memanageddeviceuser.MeManagedDeviceUserClient
	MeManagedDeviceWindowsProtectionState                                      *memanageddevicewindowsprotectionstate.MeManagedDeviceWindowsProtectionStateClient
	MeManagedDeviceWindowsProtectionStateDetectedMalwareState                  *memanageddevicewindowsprotectionstatedetectedmalwarestate.MeManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	MeManager                                                                  *memanager.MeManagerClient
	MeMemberOf                                                                 *mememberof.MeMemberOfClient
	MeMessage                                                                  *memessage.MeMessageClient
	MeMessageAttachment                                                        *memessageattachment.MeMessageAttachmentClient
	MeMessageExtension                                                         *memessageextension.MeMessageExtensionClient
	MeMessageMention                                                           *memessagemention.MeMessageMentionClient
	MeMobileAppIntentAndState                                                  *memobileappintentandstate.MeMobileAppIntentAndStateClient
	MeMobileAppTroubleshootingEvent                                            *memobileapptroubleshootingevent.MeMobileAppTroubleshootingEventClient
	MeMobileAppTroubleshootingEventAppLogCollectionRequest                     *memobileapptroubleshootingeventapplogcollectionrequest.MeMobileAppTroubleshootingEventAppLogCollectionRequestClient
	MeNotification                                                             *menotification.MeNotificationClient
	MeOauth2PermissionGrant                                                    *meoauth2permissiongrant.MeOauth2PermissionGrantClient
	MeOnenote                                                                  *meonenote.MeOnenoteClient
	MeOnenoteNotebook                                                          *meonenotenotebook.MeOnenoteNotebookClient
	MeOnenoteNotebookSection                                                   *meonenotenotebooksection.MeOnenoteNotebookSectionClient
	MeOnenoteNotebookSectionGroup                                              *meonenotenotebooksectiongroup.MeOnenoteNotebookSectionGroupClient
	MeOnenoteNotebookSectionGroupParentNotebook                                *meonenotenotebooksectiongroupparentnotebook.MeOnenoteNotebookSectionGroupParentNotebookClient
	MeOnenoteNotebookSectionGroupParentSectionGroup                            *meonenotenotebooksectiongroupparentsectiongroup.MeOnenoteNotebookSectionGroupParentSectionGroupClient
	MeOnenoteNotebookSectionGroupSection                                       *meonenotenotebooksectiongroupsection.MeOnenoteNotebookSectionGroupSectionClient
	MeOnenoteNotebookSectionGroupSectionGroup                                  *meonenotenotebooksectiongroupsectiongroup.MeOnenoteNotebookSectionGroupSectionGroupClient
	MeOnenoteNotebookSectionGroupSectionPage                                   *meonenotenotebooksectiongroupsectionpage.MeOnenoteNotebookSectionGroupSectionPageClient
	MeOnenoteNotebookSectionGroupSectionPageContent                            *meonenotenotebooksectiongroupsectionpagecontent.MeOnenoteNotebookSectionGroupSectionPageContentClient
	MeOnenoteNotebookSectionGroupSectionPageParentNotebook                     *meonenotenotebooksectiongroupsectionpageparentnotebook.MeOnenoteNotebookSectionGroupSectionPageParentNotebookClient
	MeOnenoteNotebookSectionGroupSectionPageParentSection                      *meonenotenotebooksectiongroupsectionpageparentsection.MeOnenoteNotebookSectionGroupSectionPageParentSectionClient
	MeOnenoteNotebookSectionGroupSectionParentNotebook                         *meonenotenotebooksectiongroupsectionparentnotebook.MeOnenoteNotebookSectionGroupSectionParentNotebookClient
	MeOnenoteNotebookSectionGroupSectionParentSectionGroup                     *meonenotenotebooksectiongroupsectionparentsectiongroup.MeOnenoteNotebookSectionGroupSectionParentSectionGroupClient
	MeOnenoteNotebookSectionPage                                               *meonenotenotebooksectionpage.MeOnenoteNotebookSectionPageClient
	MeOnenoteNotebookSectionPageContent                                        *meonenotenotebooksectionpagecontent.MeOnenoteNotebookSectionPageContentClient
	MeOnenoteNotebookSectionPageParentNotebook                                 *meonenotenotebooksectionpageparentnotebook.MeOnenoteNotebookSectionPageParentNotebookClient
	MeOnenoteNotebookSectionPageParentSection                                  *meonenotenotebooksectionpageparentsection.MeOnenoteNotebookSectionPageParentSectionClient
	MeOnenoteNotebookSectionParentNotebook                                     *meonenotenotebooksectionparentnotebook.MeOnenoteNotebookSectionParentNotebookClient
	MeOnenoteNotebookSectionParentSectionGroup                                 *meonenotenotebooksectionparentsectiongroup.MeOnenoteNotebookSectionParentSectionGroupClient
	MeOnenoteOperation                                                         *meonenoteoperation.MeOnenoteOperationClient
	MeOnenotePage                                                              *meonenotepage.MeOnenotePageClient
	MeOnenotePageContent                                                       *meonenotepagecontent.MeOnenotePageContentClient
	MeOnenotePageParentNotebook                                                *meonenotepageparentnotebook.MeOnenotePageParentNotebookClient
	MeOnenotePageParentSection                                                 *meonenotepageparentsection.MeOnenotePageParentSectionClient
	MeOnenoteResource                                                          *meonenoteresource.MeOnenoteResourceClient
	MeOnenoteResourceContent                                                   *meonenoteresourcecontent.MeOnenoteResourceContentClient
	MeOnenoteSection                                                           *meonenotesection.MeOnenoteSectionClient
	MeOnenoteSectionGroup                                                      *meonenotesectiongroup.MeOnenoteSectionGroupClient
	MeOnenoteSectionGroupParentNotebook                                        *meonenotesectiongroupparentnotebook.MeOnenoteSectionGroupParentNotebookClient
	MeOnenoteSectionGroupParentSectionGroup                                    *meonenotesectiongroupparentsectiongroup.MeOnenoteSectionGroupParentSectionGroupClient
	MeOnenoteSectionGroupSection                                               *meonenotesectiongroupsection.MeOnenoteSectionGroupSectionClient
	MeOnenoteSectionGroupSectionGroup                                          *meonenotesectiongroupsectiongroup.MeOnenoteSectionGroupSectionGroupClient
	MeOnenoteSectionGroupSectionPage                                           *meonenotesectiongroupsectionpage.MeOnenoteSectionGroupSectionPageClient
	MeOnenoteSectionGroupSectionPageContent                                    *meonenotesectiongroupsectionpagecontent.MeOnenoteSectionGroupSectionPageContentClient
	MeOnenoteSectionGroupSectionPageParentNotebook                             *meonenotesectiongroupsectionpageparentnotebook.MeOnenoteSectionGroupSectionPageParentNotebookClient
	MeOnenoteSectionGroupSectionPageParentSection                              *meonenotesectiongroupsectionpageparentsection.MeOnenoteSectionGroupSectionPageParentSectionClient
	MeOnenoteSectionGroupSectionParentNotebook                                 *meonenotesectiongroupsectionparentnotebook.MeOnenoteSectionGroupSectionParentNotebookClient
	MeOnenoteSectionGroupSectionParentSectionGroup                             *meonenotesectiongroupsectionparentsectiongroup.MeOnenoteSectionGroupSectionParentSectionGroupClient
	MeOnenoteSectionPage                                                       *meonenotesectionpage.MeOnenoteSectionPageClient
	MeOnenoteSectionPageContent                                                *meonenotesectionpagecontent.MeOnenoteSectionPageContentClient
	MeOnenoteSectionPageParentNotebook                                         *meonenotesectionpageparentnotebook.MeOnenoteSectionPageParentNotebookClient
	MeOnenoteSectionPageParentSection                                          *meonenotesectionpageparentsection.MeOnenoteSectionPageParentSectionClient
	MeOnenoteSectionParentNotebook                                             *meonenotesectionparentnotebook.MeOnenoteSectionParentNotebookClient
	MeOnenoteSectionParentSectionGroup                                         *meonenotesectionparentsectiongroup.MeOnenoteSectionParentSectionGroupClient
	MeOnlineMeeting                                                            *meonlinemeeting.MeOnlineMeetingClient
	MeOnlineMeetingAlternativeRecording                                        *meonlinemeetingalternativerecording.MeOnlineMeetingAlternativeRecordingClient
	MeOnlineMeetingAttendanceReport                                            *meonlinemeetingattendancereport.MeOnlineMeetingAttendanceReportClient
	MeOnlineMeetingAttendanceReportAttendanceRecord                            *meonlinemeetingattendancereportattendancerecord.MeOnlineMeetingAttendanceReportAttendanceRecordClient
	MeOnlineMeetingAttendeeReport                                              *meonlinemeetingattendeereport.MeOnlineMeetingAttendeeReportClient
	MeOnlineMeetingBroadcastRecording                                          *meonlinemeetingbroadcastrecording.MeOnlineMeetingBroadcastRecordingClient
	MeOnlineMeetingMeetingAttendanceReport                                     *meonlinemeetingmeetingattendancereport.MeOnlineMeetingMeetingAttendanceReportClient
	MeOnlineMeetingMeetingAttendanceReportAttendanceRecord                     *meonlinemeetingmeetingattendancereportattendancerecord.MeOnlineMeetingMeetingAttendanceReportAttendanceRecordClient
	MeOnlineMeetingRecording                                                   *meonlinemeetingrecording.MeOnlineMeetingRecordingClient
	MeOnlineMeetingRecordingContent                                            *meonlinemeetingrecordingcontent.MeOnlineMeetingRecordingContentClient
	MeOnlineMeetingRegistration                                                *meonlinemeetingregistration.MeOnlineMeetingRegistrationClient
	MeOnlineMeetingRegistrationCustomQuestion                                  *meonlinemeetingregistrationcustomquestion.MeOnlineMeetingRegistrationCustomQuestionClient
	MeOnlineMeetingRegistrationRegistrant                                      *meonlinemeetingregistrationregistrant.MeOnlineMeetingRegistrationRegistrantClient
	MeOnlineMeetingTranscript                                                  *meonlinemeetingtranscript.MeOnlineMeetingTranscriptClient
	MeOnlineMeetingTranscriptContent                                           *meonlinemeetingtranscriptcontent.MeOnlineMeetingTranscriptContentClient
	MeOnlineMeetingTranscriptMetadataContent                                   *meonlinemeetingtranscriptmetadatacontent.MeOnlineMeetingTranscriptMetadataContentClient
	MeOutlook                                                                  *meoutlook.MeOutlookClient
	MeOutlookMasterCategory                                                    *meoutlookmastercategory.MeOutlookMasterCategoryClient
	MeOutlookTask                                                              *meoutlooktask.MeOutlookTaskClient
	MeOutlookTaskAttachment                                                    *meoutlooktaskattachment.MeOutlookTaskAttachmentClient
	MeOutlookTaskFolder                                                        *meoutlooktaskfolder.MeOutlookTaskFolderClient
	MeOutlookTaskFolderTask                                                    *meoutlooktaskfoldertask.MeOutlookTaskFolderTaskClient
	MeOutlookTaskFolderTaskAttachment                                          *meoutlooktaskfoldertaskattachment.MeOutlookTaskFolderTaskAttachmentClient
	MeOutlookTaskGroup                                                         *meoutlooktaskgroup.MeOutlookTaskGroupClient
	MeOutlookTaskGroupTaskFolder                                               *meoutlooktaskgrouptaskfolder.MeOutlookTaskGroupTaskFolderClient
	MeOutlookTaskGroupTaskFolderTask                                           *meoutlooktaskgrouptaskfoldertask.MeOutlookTaskGroupTaskFolderTaskClient
	MeOutlookTaskGroupTaskFolderTaskAttachment                                 *meoutlooktaskgrouptaskfoldertaskattachment.MeOutlookTaskGroupTaskFolderTaskAttachmentClient
	MeOwnedDevice                                                              *meowneddevice.MeOwnedDeviceClient
	MeOwnedObject                                                              *meownedobject.MeOwnedObjectClient
	MePendingAccessReviewInstance                                              *mependingaccessreviewinstance.MePendingAccessReviewInstanceClient
	MePendingAccessReviewInstanceContactedReviewer                             *mependingaccessreviewinstancecontactedreviewer.MePendingAccessReviewInstanceContactedReviewerClient
	MePendingAccessReviewInstanceDecision                                      *mependingaccessreviewinstancedecision.MePendingAccessReviewInstanceDecisionClient
	MePendingAccessReviewInstanceDecisionInsight                               *mependingaccessreviewinstancedecisioninsight.MePendingAccessReviewInstanceDecisionInsightClient
	MePendingAccessReviewInstanceDecisionInstance                              *mependingaccessreviewinstancedecisioninstance.MePendingAccessReviewInstanceDecisionInstanceClient
	MePendingAccessReviewInstanceDecisionInstanceContactedReviewer             *mependingaccessreviewinstancedecisioninstancecontactedreviewer.MePendingAccessReviewInstanceDecisionInstanceContactedReviewerClient
	MePendingAccessReviewInstanceDecisionInstanceDefinition                    *mependingaccessreviewinstancedecisioninstancedefinition.MePendingAccessReviewInstanceDecisionInstanceDefinitionClient
	MePendingAccessReviewInstanceDecisionInstanceStage                         *mependingaccessreviewinstancedecisioninstancestage.MePendingAccessReviewInstanceDecisionInstanceStageClient
	MePendingAccessReviewInstanceDecisionInstanceStageDecision                 *mependingaccessreviewinstancedecisioninstancestagedecision.MePendingAccessReviewInstanceDecisionInstanceStageDecisionClient
	MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsight          *mependingaccessreviewinstancedecisioninstancestagedecisioninsight.MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient
	MePendingAccessReviewInstanceDefinition                                    *mependingaccessreviewinstancedefinition.MePendingAccessReviewInstanceDefinitionClient
	MePendingAccessReviewInstanceStage                                         *mependingaccessreviewinstancestage.MePendingAccessReviewInstanceStageClient
	MePendingAccessReviewInstanceStageDecision                                 *mependingaccessreviewinstancestagedecision.MePendingAccessReviewInstanceStageDecisionClient
	MePendingAccessReviewInstanceStageDecisionInsight                          *mependingaccessreviewinstancestagedecisioninsight.MePendingAccessReviewInstanceStageDecisionInsightClient
	MePendingAccessReviewInstanceStageDecisionInstance                         *mependingaccessreviewinstancestagedecisioninstance.MePendingAccessReviewInstanceStageDecisionInstanceClient
	MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewer        *mependingaccessreviewinstancestagedecisioninstancecontactedreviewer.MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient
	MePendingAccessReviewInstanceStageDecisionInstanceDecision                 *mependingaccessreviewinstancestagedecisioninstancedecision.MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient
	MePendingAccessReviewInstanceStageDecisionInstanceDecisionInsight          *mependingaccessreviewinstancestagedecisioninstancedecisioninsight.MePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient
	MePendingAccessReviewInstanceStageDecisionInstanceDefinition               *mependingaccessreviewinstancestagedecisioninstancedefinition.MePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient
	MePeople                                                                   *mepeople.MePeopleClient
	MePermissionGrant                                                          *mepermissiongrant.MePermissionGrantClient
	MePhoto                                                                    *mephoto.MePhotoClient
	MePlanner                                                                  *meplanner.MePlannerClient
	MePlannerAll                                                               *meplannerall.MePlannerAllClient
	MePlannerFavoritePlan                                                      *meplannerfavoriteplan.MePlannerFavoritePlanClient
	MePlannerPlan                                                              *meplannerplan.MePlannerPlanClient
	MePlannerPlanBucket                                                        *meplannerplanbucket.MePlannerPlanBucketClient
	MePlannerPlanBucketTask                                                    *meplannerplanbuckettask.MePlannerPlanBucketTaskClient
	MePlannerPlanBucketTaskAssignedToTaskBoardFormat                           *meplannerplanbuckettaskassignedtotaskboardformat.MePlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	MePlannerPlanBucketTaskBucketTaskBoardFormat                               *meplannerplanbuckettaskbuckettaskboardformat.MePlannerPlanBucketTaskBucketTaskBoardFormatClient
	MePlannerPlanBucketTaskDetail                                              *meplannerplanbuckettaskdetail.MePlannerPlanBucketTaskDetailClient
	MePlannerPlanBucketTaskProgressTaskBoardFormat                             *meplannerplanbuckettaskprogresstaskboardformat.MePlannerPlanBucketTaskProgressTaskBoardFormatClient
	MePlannerPlanDetail                                                        *meplannerplandetail.MePlannerPlanDetailClient
	MePlannerPlanTask                                                          *meplannerplantask.MePlannerPlanTaskClient
	MePlannerPlanTaskAssignedToTaskBoardFormat                                 *meplannerplantaskassignedtotaskboardformat.MePlannerPlanTaskAssignedToTaskBoardFormatClient
	MePlannerPlanTaskBucketTaskBoardFormat                                     *meplannerplantaskbuckettaskboardformat.MePlannerPlanTaskBucketTaskBoardFormatClient
	MePlannerPlanTaskDetail                                                    *meplannerplantaskdetail.MePlannerPlanTaskDetailClient
	MePlannerPlanTaskProgressTaskBoardFormat                                   *meplannerplantaskprogresstaskboardformat.MePlannerPlanTaskProgressTaskBoardFormatClient
	MePlannerRecentPlan                                                        *meplannerrecentplan.MePlannerRecentPlanClient
	MePlannerRosterPlan                                                        *meplannerrosterplan.MePlannerRosterPlanClient
	MePlannerTask                                                              *meplannertask.MePlannerTaskClient
	MePlannerTaskAssignedToTaskBoardFormat                                     *meplannertaskassignedtotaskboardformat.MePlannerTaskAssignedToTaskBoardFormatClient
	MePlannerTaskBucketTaskBoardFormat                                         *meplannertaskbuckettaskboardformat.MePlannerTaskBucketTaskBoardFormatClient
	MePlannerTaskDetail                                                        *meplannertaskdetail.MePlannerTaskDetailClient
	MePlannerTaskProgressTaskBoardFormat                                       *meplannertaskprogresstaskboardformat.MePlannerTaskProgressTaskBoardFormatClient
	MePresence                                                                 *mepresence.MePresenceClient
	MeProfile                                                                  *meprofile.MeProfileClient
	MeProfileAccount                                                           *meprofileaccount.MeProfileAccountClient
	MeProfileAddress                                                           *meprofileaddress.MeProfileAddressClient
	MeProfileAnniversary                                                       *meprofileanniversary.MeProfileAnniversaryClient
	MeProfileAward                                                             *meprofileaward.MeProfileAwardClient
	MeProfileCertification                                                     *meprofilecertification.MeProfileCertificationClient
	MeProfileEducationalActivity                                               *meprofileeducationalactivity.MeProfileEducationalActivityClient
	MeProfileEmail                                                             *meprofileemail.MeProfileEmailClient
	MeProfileInterest                                                          *meprofileinterest.MeProfileInterestClient
	MeProfileLanguage                                                          *meprofilelanguage.MeProfileLanguageClient
	MeProfileName                                                              *meprofilename.MeProfileNameClient
	MeProfileNote                                                              *meprofilenote.MeProfileNoteClient
	MeProfilePatent                                                            *meprofilepatent.MeProfilePatentClient
	MeProfilePhone                                                             *meprofilephone.MeProfilePhoneClient
	MeProfilePosition                                                          *meprofileposition.MeProfilePositionClient
	MeProfileProject                                                           *meprofileproject.MeProfileProjectClient
	MeProfilePublication                                                       *meprofilepublication.MeProfilePublicationClient
	MeProfileSkill                                                             *meprofileskill.MeProfileSkillClient
	MeProfileWebAccount                                                        *meprofilewebaccount.MeProfileWebAccountClient
	MeProfileWebsite                                                           *meprofilewebsite.MeProfileWebsiteClient
	MeRegisteredDevice                                                         *meregistereddevice.MeRegisteredDeviceClient
	MeScopedRoleMemberOf                                                       *mescopedrolememberof.MeScopedRoleMemberOfClient
	MeSecurity                                                                 *mesecurity.MeSecurityClient
	MeSecurityInformationProtection                                            *mesecurityinformationprotection.MeSecurityInformationProtectionClient
	MeSecurityInformationProtectionLabelPolicySetting                          *mesecurityinformationprotectionlabelpolicysetting.MeSecurityInformationProtectionLabelPolicySettingClient
	MeSecurityInformationProtectionSensitivityLabel                            *mesecurityinformationprotectionsensitivitylabel.MeSecurityInformationProtectionSensitivityLabelClient
	MeSecurityInformationProtectionSensitivityLabelParent                      *mesecurityinformationprotectionsensitivitylabelparent.MeSecurityInformationProtectionSensitivityLabelParentClient
	MeSetting                                                                  *mesetting.MeSettingClient
	MeSettingContactMergeSuggestion                                            *mesettingcontactmergesuggestion.MeSettingContactMergeSuggestionClient
	MeSettingItemInsight                                                       *mesettingiteminsight.MeSettingItemInsightClient
	MeSettingRegionalAndLanguageSetting                                        *mesettingregionalandlanguagesetting.MeSettingRegionalAndLanguageSettingClient
	MeSettingShiftPreference                                                   *mesettingshiftpreference.MeSettingShiftPreferenceClient
	MeSponsor                                                                  *mesponsor.MeSponsorClient
	MeTeamwork                                                                 *meteamwork.MeTeamworkClient
	MeTeamworkAssociatedTeam                                                   *meteamworkassociatedteam.MeTeamworkAssociatedTeamClient
	MeTeamworkAssociatedTeamTeam                                               *meteamworkassociatedteamteam.MeTeamworkAssociatedTeamTeamClient
	MeTeamworkInstalledApp                                                     *meteamworkinstalledapp.MeTeamworkInstalledAppClient
	MeTeamworkInstalledAppChat                                                 *meteamworkinstalledappchat.MeTeamworkInstalledAppChatClient
	MeTeamworkInstalledAppTeamsApp                                             *meteamworkinstalledappteamsapp.MeTeamworkInstalledAppTeamsAppClient
	MeTeamworkInstalledAppTeamsAppDefinition                                   *meteamworkinstalledappteamsappdefinition.MeTeamworkInstalledAppTeamsAppDefinitionClient
	MeTodo                                                                     *metodo.MeTodoClient
	MeTodoList                                                                 *metodolist.MeTodoListClient
	MeTodoListExtension                                                        *metodolistextension.MeTodoListExtensionClient
	MeTodoListTask                                                             *metodolisttask.MeTodoListTaskClient
	MeTodoListTaskAttachment                                                   *metodolisttaskattachment.MeTodoListTaskAttachmentClient
	MeTodoListTaskAttachmentSession                                            *metodolisttaskattachmentsession.MeTodoListTaskAttachmentSessionClient
	MeTodoListTaskAttachmentSessionContent                                     *metodolisttaskattachmentsessioncontent.MeTodoListTaskAttachmentSessionContentClient
	MeTodoListTaskChecklistItem                                                *metodolisttaskchecklistitem.MeTodoListTaskChecklistItemClient
	MeTodoListTaskExtension                                                    *metodolisttaskextension.MeTodoListTaskExtensionClient
	MeTodoListTaskLinkedResource                                               *metodolisttasklinkedresource.MeTodoListTaskLinkedResourceClient
	MeTransitiveMemberOf                                                       *metransitivememberof.MeTransitiveMemberOfClient
	MeTransitiveReport                                                         *metransitivereport.MeTransitiveReportClient
	MeUsageRight                                                               *meusageright.MeUsageRightClient
	MeWindowsInformationProtectionDeviceRegistration                           *mewindowsinformationprotectiondeviceregistration.MeWindowsInformationProtectionDeviceRegistrationClient
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

	meAnalyticActivityStatisticClient, err := meanalyticactivitystatistic.NewMeAnalyticActivityStatisticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAnalyticActivityStatistic client: %+v", err)
	}
	configureFunc(meAnalyticActivityStatisticClient.Client)

	meAnalyticClient, err := meanalytic.NewMeAnalyticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAnalytic client: %+v", err)
	}
	configureFunc(meAnalyticClient.Client)

	meAppConsentRequestsForApprovalClient, err := meappconsentrequestsforapproval.NewMeAppConsentRequestsForApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAppConsentRequestsForApproval client: %+v", err)
	}
	configureFunc(meAppConsentRequestsForApprovalClient.Client)

	meAppConsentRequestsForApprovalUserConsentRequestApprovalClient, err := meappconsentrequestsforapprovaluserconsentrequestapproval.NewMeAppConsentRequestsForApprovalUserConsentRequestApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAppConsentRequestsForApprovalUserConsentRequestApproval client: %+v", err)
	}
	configureFunc(meAppConsentRequestsForApprovalUserConsentRequestApprovalClient.Client)

	meAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient, err := meappconsentrequestsforapprovaluserconsentrequestapprovalstep.NewMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAppConsentRequestsForApprovalUserConsentRequestApprovalStep client: %+v", err)
	}
	configureFunc(meAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient.Client)

	meAppConsentRequestsForApprovalUserConsentRequestClient, err := meappconsentrequestsforapprovaluserconsentrequest.NewMeAppConsentRequestsForApprovalUserConsentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAppConsentRequestsForApprovalUserConsentRequest client: %+v", err)
	}
	configureFunc(meAppConsentRequestsForApprovalUserConsentRequestClient.Client)

	meAppRoleAssignedResourceClient, err := meapproleassignedresource.NewMeAppRoleAssignedResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAppRoleAssignedResource client: %+v", err)
	}
	configureFunc(meAppRoleAssignedResourceClient.Client)

	meAppRoleAssignmentClient, err := meapproleassignment.NewMeAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAppRoleAssignment client: %+v", err)
	}
	configureFunc(meAppRoleAssignmentClient.Client)

	meApprovalClient, err := meapproval.NewMeApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeApproval client: %+v", err)
	}
	configureFunc(meApprovalClient.Client)

	meApprovalStepClient, err := meapprovalstep.NewMeApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeApprovalStep client: %+v", err)
	}
	configureFunc(meApprovalStepClient.Client)

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

	meAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient, err := meauthenticationpasswordlessmicrosoftauthenticatormethod.NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationPasswordlessMicrosoftAuthenticatorMethod client: %+v", err)
	}
	configureFunc(meAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient.Client)

	meAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient, err := meauthenticationpasswordlessmicrosoftauthenticatormethoddevice.NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice client: %+v", err)
	}
	configureFunc(meAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient.Client)

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

	meCalendarCalendarViewExceptionOccurrenceAttachmentClient, err := mecalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceAttachmentClient.Client)

	meCalendarCalendarViewExceptionOccurrenceCalendarClient, err := mecalendarcalendarviewexceptionoccurrencecalendar.NewMeCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceCalendarClient.Client)

	meCalendarCalendarViewExceptionOccurrenceClient, err := mecalendarcalendarviewexceptionoccurrence.NewMeCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceClient.Client)

	meCalendarCalendarViewExceptionOccurrenceExtensionClient, err := mecalendarcalendarviewexceptionoccurrenceextension.NewMeCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceExtensionClient.Client)

	meCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := mecalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	meCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, err := mecalendarcalendarviewexceptionoccurrenceinstancecalendar.NewMeCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	meCalendarCalendarViewExceptionOccurrenceInstanceClient, err := mecalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceInstanceClient.Client)

	meCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient, err := mecalendarcalendarviewexceptionoccurrenceinstanceextension.NewMeCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

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

	meCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := mecalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	meCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, err := mecalendarcalendarviewinstanceexceptionoccurrencecalendar.NewMeCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	meCalendarCalendarViewInstanceExceptionOccurrenceClient, err := mecalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceExceptionOccurrenceClient.Client)

	meCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, err := mecalendarcalendarviewinstanceexceptionoccurrenceextension.NewMeCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

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

	meCalendarEventExceptionOccurrenceAttachmentClient, err := mecalendareventexceptionoccurrenceattachment.NewMeCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceAttachmentClient.Client)

	meCalendarEventExceptionOccurrenceCalendarClient, err := mecalendareventexceptionoccurrencecalendar.NewMeCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceCalendarClient.Client)

	meCalendarEventExceptionOccurrenceClient, err := mecalendareventexceptionoccurrence.NewMeCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceClient.Client)

	meCalendarEventExceptionOccurrenceExtensionClient, err := mecalendareventexceptionoccurrenceextension.NewMeCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceExtensionClient.Client)

	meCalendarEventExceptionOccurrenceInstanceAttachmentClient, err := mecalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceInstanceAttachmentClient.Client)

	meCalendarEventExceptionOccurrenceInstanceCalendarClient, err := mecalendareventexceptionoccurrenceinstancecalendar.NewMeCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceInstanceCalendarClient.Client)

	meCalendarEventExceptionOccurrenceInstanceClient, err := mecalendareventexceptionoccurrenceinstance.NewMeCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceInstanceClient.Client)

	meCalendarEventExceptionOccurrenceInstanceExtensionClient, err := mecalendareventexceptionoccurrenceinstanceextension.NewMeCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarEventExceptionOccurrenceInstanceExtensionClient.Client)

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

	meCalendarEventInstanceExceptionOccurrenceAttachmentClient, err := mecalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceExceptionOccurrenceAttachmentClient.Client)

	meCalendarEventInstanceExceptionOccurrenceCalendarClient, err := mecalendareventinstanceexceptionoccurrencecalendar.NewMeCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceExceptionOccurrenceCalendarClient.Client)

	meCalendarEventInstanceExceptionOccurrenceClient, err := mecalendareventinstanceexceptionoccurrence.NewMeCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceExceptionOccurrenceClient.Client)

	meCalendarEventInstanceExceptionOccurrenceExtensionClient, err := mecalendareventinstanceexceptionoccurrenceextension.NewMeCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarEventInstanceExceptionOccurrenceExtensionClient.Client)

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

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient.Client)

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrencecalendar.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient.Client)

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceClient.Client)

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrenceextension.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient.Client)

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrenceinstance.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient.Client)

	meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient, err := mecalendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension.NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

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

	meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, err := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient, err := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrence.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient.Client)

	meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, err := mecalendargroupcalendarcalendarviewinstanceexceptionoccurrenceextension.NewMeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

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

	meCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient, err := mecalendargroupcalendareventexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient.Client)

	meCalendarGroupCalendarEventExceptionOccurrenceCalendarClient, err := mecalendargroupcalendareventexceptionoccurrencecalendar.NewMeCalendarGroupCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceCalendarClient.Client)

	meCalendarGroupCalendarEventExceptionOccurrenceClient, err := mecalendargroupcalendareventexceptionoccurrence.NewMeCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceClient.Client)

	meCalendarGroupCalendarEventExceptionOccurrenceExtensionClient, err := mecalendargroupcalendareventexceptionoccurrenceextension.NewMeCalendarGroupCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceExtensionClient.Client)

	meCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient, err := mecalendargroupcalendareventexceptionoccurrenceinstanceattachment.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient.Client)

	meCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient, err := mecalendargroupcalendareventexceptionoccurrenceinstancecalendar.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient.Client)

	meCalendarGroupCalendarEventExceptionOccurrenceInstanceClient, err := mecalendargroupcalendareventexceptionoccurrenceinstance.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceInstanceClient.Client)

	meCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient, err := mecalendargroupcalendareventexceptionoccurrenceinstanceextension.NewMeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient.Client)

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

	meCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient, err := mecalendargroupcalendareventinstanceexceptionoccurrenceattachment.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient.Client)

	meCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient, err := mecalendargroupcalendareventinstanceexceptionoccurrencecalendar.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient.Client)

	meCalendarGroupCalendarEventInstanceExceptionOccurrenceClient, err := mecalendargroupcalendareventinstanceexceptionoccurrence.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceExceptionOccurrenceClient.Client)

	meCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient, err := mecalendargroupcalendareventinstanceexceptionoccurrenceextension.NewMeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient.Client)

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

	meCalendarViewExceptionOccurrenceAttachmentClient, err := mecalendarviewexceptionoccurrenceattachment.NewMeCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceAttachmentClient.Client)

	meCalendarViewExceptionOccurrenceCalendarClient, err := mecalendarviewexceptionoccurrencecalendar.NewMeCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceCalendarClient.Client)

	meCalendarViewExceptionOccurrenceClient, err := mecalendarviewexceptionoccurrence.NewMeCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceClient.Client)

	meCalendarViewExceptionOccurrenceExtensionClient, err := mecalendarviewexceptionoccurrenceextension.NewMeCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceExtensionClient.Client)

	meCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := mecalendarviewexceptionoccurrenceinstanceattachment.NewMeCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	meCalendarViewExceptionOccurrenceInstanceCalendarClient, err := mecalendarviewexceptionoccurrenceinstancecalendar.NewMeCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	meCalendarViewExceptionOccurrenceInstanceClient, err := mecalendarviewexceptionoccurrenceinstance.NewMeCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceInstanceClient.Client)

	meCalendarViewExceptionOccurrenceInstanceExtensionClient, err := mecalendarviewexceptionoccurrenceinstanceextension.NewMeCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(meCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

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

	meCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := mecalendarviewinstanceexceptionoccurrenceattachment.NewMeCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	meCalendarViewInstanceExceptionOccurrenceCalendarClient, err := mecalendarviewinstanceexceptionoccurrencecalendar.NewMeCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	meCalendarViewInstanceExceptionOccurrenceClient, err := mecalendarviewinstanceexceptionoccurrence.NewMeCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceExceptionOccurrenceClient.Client)

	meCalendarViewInstanceExceptionOccurrenceExtensionClient, err := mecalendarviewinstanceexceptionoccurrenceextension.NewMeCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

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

	meChatOperationClient, err := mechatoperation.NewMeChatOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeChatOperation client: %+v", err)
	}
	configureFunc(meChatOperationClient.Client)

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

	meCloudPCClient, err := mecloudpc.NewMeCloudPCClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeCloudPC client: %+v", err)
	}
	configureFunc(meCloudPCClient.Client)

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

	meDeviceClient, err := medevice.NewMeDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDevice client: %+v", err)
	}
	configureFunc(meDeviceClient.Client)

	meDeviceCommandClient, err := medevicecommand.NewMeDeviceCommandClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceCommand client: %+v", err)
	}
	configureFunc(meDeviceCommandClient.Client)

	meDeviceCommandResponsepayloadClient, err := medevicecommandresponsepayload.NewMeDeviceCommandResponsepayloadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceCommandResponsepayload client: %+v", err)
	}
	configureFunc(meDeviceCommandResponsepayloadClient.Client)

	meDeviceEnrollmentConfigurationAssignmentClient, err := medeviceenrollmentconfigurationassignment.NewMeDeviceEnrollmentConfigurationAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceEnrollmentConfigurationAssignment client: %+v", err)
	}
	configureFunc(meDeviceEnrollmentConfigurationAssignmentClient.Client)

	meDeviceEnrollmentConfigurationClient, err := medeviceenrollmentconfiguration.NewMeDeviceEnrollmentConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceEnrollmentConfiguration client: %+v", err)
	}
	configureFunc(meDeviceEnrollmentConfigurationClient.Client)

	meDeviceExtensionClient, err := medeviceextension.NewMeDeviceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceExtension client: %+v", err)
	}
	configureFunc(meDeviceExtensionClient.Client)

	meDeviceManagementTroubleshootingEventClient, err := medevicemanagementtroubleshootingevent.NewMeDeviceManagementTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceManagementTroubleshootingEvent client: %+v", err)
	}
	configureFunc(meDeviceManagementTroubleshootingEventClient.Client)

	meDeviceMemberOfClient, err := medevicememberof.NewMeDeviceMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceMemberOf client: %+v", err)
	}
	configureFunc(meDeviceMemberOfClient.Client)

	meDeviceRegisteredOwnerClient, err := medeviceregisteredowner.NewMeDeviceRegisteredOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceRegisteredOwner client: %+v", err)
	}
	configureFunc(meDeviceRegisteredOwnerClient.Client)

	meDeviceRegisteredUserClient, err := medeviceregistereduser.NewMeDeviceRegisteredUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceRegisteredUser client: %+v", err)
	}
	configureFunc(meDeviceRegisteredUserClient.Client)

	meDeviceTransitiveMemberOfClient, err := medevicetransitivememberof.NewMeDeviceTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceTransitiveMemberOf client: %+v", err)
	}
	configureFunc(meDeviceTransitiveMemberOfClient.Client)

	meDeviceUsageRightClient, err := medeviceusageright.NewMeDeviceUsageRightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeDeviceUsageRight client: %+v", err)
	}
	configureFunc(meDeviceUsageRightClient.Client)

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

	meEventExceptionOccurrenceAttachmentClient, err := meeventexceptionoccurrenceattachment.NewMeEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceAttachmentClient.Client)

	meEventExceptionOccurrenceCalendarClient, err := meeventexceptionoccurrencecalendar.NewMeEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceCalendarClient.Client)

	meEventExceptionOccurrenceClient, err := meeventexceptionoccurrence.NewMeEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceClient.Client)

	meEventExceptionOccurrenceExtensionClient, err := meeventexceptionoccurrenceextension.NewMeEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceExtensionClient.Client)

	meEventExceptionOccurrenceInstanceAttachmentClient, err := meeventexceptionoccurrenceinstanceattachment.NewMeEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceInstanceAttachmentClient.Client)

	meEventExceptionOccurrenceInstanceCalendarClient, err := meeventexceptionoccurrenceinstancecalendar.NewMeEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceInstanceCalendarClient.Client)

	meEventExceptionOccurrenceInstanceClient, err := meeventexceptionoccurrenceinstance.NewMeEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceInstanceClient.Client)

	meEventExceptionOccurrenceInstanceExtensionClient, err := meeventexceptionoccurrenceinstanceextension.NewMeEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(meEventExceptionOccurrenceInstanceExtensionClient.Client)

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

	meEventInstanceExceptionOccurrenceAttachmentClient, err := meeventinstanceexceptionoccurrenceattachment.NewMeEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(meEventInstanceExceptionOccurrenceAttachmentClient.Client)

	meEventInstanceExceptionOccurrenceCalendarClient, err := meeventinstanceexceptionoccurrencecalendar.NewMeEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(meEventInstanceExceptionOccurrenceCalendarClient.Client)

	meEventInstanceExceptionOccurrenceClient, err := meeventinstanceexceptionoccurrence.NewMeEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(meEventInstanceExceptionOccurrenceClient.Client)

	meEventInstanceExceptionOccurrenceExtensionClient, err := meeventinstanceexceptionoccurrenceextension.NewMeEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(meEventInstanceExceptionOccurrenceExtensionClient.Client)

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

	meInformationProtectionBitlockerClient, err := meinformationprotectionbitlocker.NewMeInformationProtectionBitlockerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionBitlocker client: %+v", err)
	}
	configureFunc(meInformationProtectionBitlockerClient.Client)

	meInformationProtectionBitlockerRecoveryKeyClient, err := meinformationprotectionbitlockerrecoverykey.NewMeInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionBitlockerRecoveryKey client: %+v", err)
	}
	configureFunc(meInformationProtectionBitlockerRecoveryKeyClient.Client)

	meInformationProtectionClient, err := meinformationprotection.NewMeInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtection client: %+v", err)
	}
	configureFunc(meInformationProtectionClient.Client)

	meInformationProtectionDataLossPreventionPolicyClient, err := meinformationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionDataLossPreventionPolicy client: %+v", err)
	}
	configureFunc(meInformationProtectionDataLossPreventionPolicyClient.Client)

	meInformationProtectionPolicyClient, err := meinformationprotectionpolicy.NewMeInformationProtectionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionPolicy client: %+v", err)
	}
	configureFunc(meInformationProtectionPolicyClient.Client)

	meInformationProtectionPolicyLabelClient, err := meinformationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionPolicyLabel client: %+v", err)
	}
	configureFunc(meInformationProtectionPolicyLabelClient.Client)

	meInformationProtectionSensitivityLabelClient, err := meinformationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(meInformationProtectionSensitivityLabelClient.Client)

	meInformationProtectionSensitivityLabelSublabelClient, err := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelSublabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionSensitivityLabelSublabel client: %+v", err)
	}
	configureFunc(meInformationProtectionSensitivityLabelSublabelClient.Client)

	meInformationProtectionSensitivityPolicySettingClient, err := meinformationprotectionsensitivitypolicysetting.NewMeInformationProtectionSensitivityPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionSensitivityPolicySetting client: %+v", err)
	}
	configureFunc(meInformationProtectionSensitivityPolicySettingClient.Client)

	meInformationProtectionThreatAssessmentRequestClient, err := meinformationprotectionthreatassessmentrequest.NewMeInformationProtectionThreatAssessmentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionThreatAssessmentRequest client: %+v", err)
	}
	configureFunc(meInformationProtectionThreatAssessmentRequestClient.Client)

	meInformationProtectionThreatAssessmentRequestResultClient, err := meinformationprotectionthreatassessmentrequestresult.NewMeInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeInformationProtectionThreatAssessmentRequestResult client: %+v", err)
	}
	configureFunc(meInformationProtectionThreatAssessmentRequestResultClient.Client)

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

	meJoinedGroupClient, err := mejoinedgroup.NewMeJoinedGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedGroup client: %+v", err)
	}
	configureFunc(meJoinedGroupClient.Client)

	meJoinedTeamClient, err := mejoinedteam.NewMeJoinedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeJoinedTeam client: %+v", err)
	}
	configureFunc(meJoinedTeamClient.Client)

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

	meMailFolderChildFolderMessageMentionClient, err := memailfolderchildfoldermessagemention.NewMeMailFolderChildFolderMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolderMessageMention client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderMessageMentionClient.Client)

	meMailFolderChildFolderMessageRuleClient, err := memailfolderchildfoldermessagerule.NewMeMailFolderChildFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolderMessageRule client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderMessageRuleClient.Client)

	meMailFolderChildFolderUserConfigurationClient, err := memailfolderchildfolderuserconfiguration.NewMeMailFolderChildFolderUserConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderChildFolderUserConfiguration client: %+v", err)
	}
	configureFunc(meMailFolderChildFolderUserConfigurationClient.Client)

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

	meMailFolderMessageMentionClient, err := memailfoldermessagemention.NewMeMailFolderMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderMessageMention client: %+v", err)
	}
	configureFunc(meMailFolderMessageMentionClient.Client)

	meMailFolderMessageRuleClient, err := memailfoldermessagerule.NewMeMailFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderMessageRule client: %+v", err)
	}
	configureFunc(meMailFolderMessageRuleClient.Client)

	meMailFolderUserConfigurationClient, err := memailfolderuserconfiguration.NewMeMailFolderUserConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMailFolderUserConfiguration client: %+v", err)
	}
	configureFunc(meMailFolderUserConfigurationClient.Client)

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

	meManagedDeviceAssignmentFilterEvaluationStatusDetailClient, err := memanageddeviceassignmentfilterevaluationstatusdetail.NewMeManagedDeviceAssignmentFilterEvaluationStatusDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceAssignmentFilterEvaluationStatusDetail client: %+v", err)
	}
	configureFunc(meManagedDeviceAssignmentFilterEvaluationStatusDetailClient.Client)

	meManagedDeviceClient, err := memanageddevice.NewMeManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDevice client: %+v", err)
	}
	configureFunc(meManagedDeviceClient.Client)

	meManagedDeviceDetectedAppClient, err := memanageddevicedetectedapp.NewMeManagedDeviceDetectedAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceDetectedApp client: %+v", err)
	}
	configureFunc(meManagedDeviceDetectedAppClient.Client)

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

	meManagedDeviceDeviceHealthScriptStateClient, err := memanageddevicedevicehealthscriptstate.NewMeManagedDeviceDeviceHealthScriptStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceDeviceHealthScriptState client: %+v", err)
	}
	configureFunc(meManagedDeviceDeviceHealthScriptStateClient.Client)

	meManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient, err := memanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid.NewMeManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId client: %+v", err)
	}
	configureFunc(meManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient.Client)

	meManagedDeviceLogCollectionRequestClient, err := memanageddevicelogcollectionrequest.NewMeManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(meManagedDeviceLogCollectionRequestClient.Client)

	meManagedDeviceManagedDeviceMobileAppConfigurationStateClient, err := memanageddevicemanageddevicemobileappconfigurationstate.NewMeManagedDeviceManagedDeviceMobileAppConfigurationStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceManagedDeviceMobileAppConfigurationState client: %+v", err)
	}
	configureFunc(meManagedDeviceManagedDeviceMobileAppConfigurationStateClient.Client)

	meManagedDeviceSecurityBaselineStateClient, err := memanageddevicesecuritybaselinestate.NewMeManagedDeviceSecurityBaselineStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceSecurityBaselineState client: %+v", err)
	}
	configureFunc(meManagedDeviceSecurityBaselineStateClient.Client)

	meManagedDeviceSecurityBaselineStateSettingStateClient, err := memanageddevicesecuritybaselinestatesettingstate.NewMeManagedDeviceSecurityBaselineStateSettingStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeManagedDeviceSecurityBaselineStateSettingState client: %+v", err)
	}
	configureFunc(meManagedDeviceSecurityBaselineStateSettingStateClient.Client)

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

	meMessageMentionClient, err := memessagemention.NewMeMessageMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMessageMention client: %+v", err)
	}
	configureFunc(meMessageMentionClient.Client)

	meMobileAppIntentAndStateClient, err := memobileappintentandstate.NewMeMobileAppIntentAndStateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMobileAppIntentAndState client: %+v", err)
	}
	configureFunc(meMobileAppIntentAndStateClient.Client)

	meMobileAppTroubleshootingEventAppLogCollectionRequestClient, err := memobileapptroubleshootingeventapplogcollectionrequest.NewMeMobileAppTroubleshootingEventAppLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMobileAppTroubleshootingEventAppLogCollectionRequest client: %+v", err)
	}
	configureFunc(meMobileAppTroubleshootingEventAppLogCollectionRequestClient.Client)

	meMobileAppTroubleshootingEventClient, err := memobileapptroubleshootingevent.NewMeMobileAppTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeMobileAppTroubleshootingEvent client: %+v", err)
	}
	configureFunc(meMobileAppTroubleshootingEventClient.Client)

	meNotificationClient, err := menotification.NewMeNotificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeNotification client: %+v", err)
	}
	configureFunc(meNotificationClient.Client)

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

	meOnlineMeetingAlternativeRecordingClient, err := meonlinemeetingalternativerecording.NewMeOnlineMeetingAlternativeRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingAlternativeRecording client: %+v", err)
	}
	configureFunc(meOnlineMeetingAlternativeRecordingClient.Client)

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

	meOnlineMeetingBroadcastRecordingClient, err := meonlinemeetingbroadcastrecording.NewMeOnlineMeetingBroadcastRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingBroadcastRecording client: %+v", err)
	}
	configureFunc(meOnlineMeetingBroadcastRecordingClient.Client)

	meOnlineMeetingClient, err := meonlinemeeting.NewMeOnlineMeetingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeeting client: %+v", err)
	}
	configureFunc(meOnlineMeetingClient.Client)

	meOnlineMeetingMeetingAttendanceReportAttendanceRecordClient, err := meonlinemeetingmeetingattendancereportattendancerecord.NewMeOnlineMeetingMeetingAttendanceReportAttendanceRecordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingMeetingAttendanceReportAttendanceRecord client: %+v", err)
	}
	configureFunc(meOnlineMeetingMeetingAttendanceReportAttendanceRecordClient.Client)

	meOnlineMeetingMeetingAttendanceReportClient, err := meonlinemeetingmeetingattendancereport.NewMeOnlineMeetingMeetingAttendanceReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingMeetingAttendanceReport client: %+v", err)
	}
	configureFunc(meOnlineMeetingMeetingAttendanceReportClient.Client)

	meOnlineMeetingRecordingClient, err := meonlinemeetingrecording.NewMeOnlineMeetingRecordingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingRecording client: %+v", err)
	}
	configureFunc(meOnlineMeetingRecordingClient.Client)

	meOnlineMeetingRecordingContentClient, err := meonlinemeetingrecordingcontent.NewMeOnlineMeetingRecordingContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingRecordingContent client: %+v", err)
	}
	configureFunc(meOnlineMeetingRecordingContentClient.Client)

	meOnlineMeetingRegistrationClient, err := meonlinemeetingregistration.NewMeOnlineMeetingRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingRegistration client: %+v", err)
	}
	configureFunc(meOnlineMeetingRegistrationClient.Client)

	meOnlineMeetingRegistrationCustomQuestionClient, err := meonlinemeetingregistrationcustomquestion.NewMeOnlineMeetingRegistrationCustomQuestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingRegistrationCustomQuestion client: %+v", err)
	}
	configureFunc(meOnlineMeetingRegistrationCustomQuestionClient.Client)

	meOnlineMeetingRegistrationRegistrantClient, err := meonlinemeetingregistrationregistrant.NewMeOnlineMeetingRegistrationRegistrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingRegistrationRegistrant client: %+v", err)
	}
	configureFunc(meOnlineMeetingRegistrationRegistrantClient.Client)

	meOnlineMeetingTranscriptClient, err := meonlinemeetingtranscript.NewMeOnlineMeetingTranscriptClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingTranscript client: %+v", err)
	}
	configureFunc(meOnlineMeetingTranscriptClient.Client)

	meOnlineMeetingTranscriptContentClient, err := meonlinemeetingtranscriptcontent.NewMeOnlineMeetingTranscriptContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingTranscriptContent client: %+v", err)
	}
	configureFunc(meOnlineMeetingTranscriptContentClient.Client)

	meOnlineMeetingTranscriptMetadataContentClient, err := meonlinemeetingtranscriptmetadatacontent.NewMeOnlineMeetingTranscriptMetadataContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOnlineMeetingTranscriptMetadataContent client: %+v", err)
	}
	configureFunc(meOnlineMeetingTranscriptMetadataContentClient.Client)

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

	meOutlookTaskAttachmentClient, err := meoutlooktaskattachment.NewMeOutlookTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskAttachment client: %+v", err)
	}
	configureFunc(meOutlookTaskAttachmentClient.Client)

	meOutlookTaskClient, err := meoutlooktask.NewMeOutlookTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTask client: %+v", err)
	}
	configureFunc(meOutlookTaskClient.Client)

	meOutlookTaskFolderClient, err := meoutlooktaskfolder.NewMeOutlookTaskFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskFolder client: %+v", err)
	}
	configureFunc(meOutlookTaskFolderClient.Client)

	meOutlookTaskFolderTaskAttachmentClient, err := meoutlooktaskfoldertaskattachment.NewMeOutlookTaskFolderTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskFolderTaskAttachment client: %+v", err)
	}
	configureFunc(meOutlookTaskFolderTaskAttachmentClient.Client)

	meOutlookTaskFolderTaskClient, err := meoutlooktaskfoldertask.NewMeOutlookTaskFolderTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskFolderTask client: %+v", err)
	}
	configureFunc(meOutlookTaskFolderTaskClient.Client)

	meOutlookTaskGroupClient, err := meoutlooktaskgroup.NewMeOutlookTaskGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskGroup client: %+v", err)
	}
	configureFunc(meOutlookTaskGroupClient.Client)

	meOutlookTaskGroupTaskFolderClient, err := meoutlooktaskgrouptaskfolder.NewMeOutlookTaskGroupTaskFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskGroupTaskFolder client: %+v", err)
	}
	configureFunc(meOutlookTaskGroupTaskFolderClient.Client)

	meOutlookTaskGroupTaskFolderTaskAttachmentClient, err := meoutlooktaskgrouptaskfoldertaskattachment.NewMeOutlookTaskGroupTaskFolderTaskAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskGroupTaskFolderTaskAttachment client: %+v", err)
	}
	configureFunc(meOutlookTaskGroupTaskFolderTaskAttachmentClient.Client)

	meOutlookTaskGroupTaskFolderTaskClient, err := meoutlooktaskgrouptaskfoldertask.NewMeOutlookTaskGroupTaskFolderTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeOutlookTaskGroupTaskFolderTask client: %+v", err)
	}
	configureFunc(meOutlookTaskGroupTaskFolderTaskClient.Client)

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

	mePendingAccessReviewInstanceClient, err := mependingaccessreviewinstance.NewMePendingAccessReviewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstance client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceClient.Client)

	mePendingAccessReviewInstanceContactedReviewerClient, err := mependingaccessreviewinstancecontactedreviewer.NewMePendingAccessReviewInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceContactedReviewerClient.Client)

	mePendingAccessReviewInstanceDecisionClient, err := mependingaccessreviewinstancedecision.NewMePendingAccessReviewInstanceDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecision client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionClient.Client)

	mePendingAccessReviewInstanceDecisionInsightClient, err := mependingaccessreviewinstancedecisioninsight.NewMePendingAccessReviewInstanceDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecisionInsight client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionInsightClient.Client)

	mePendingAccessReviewInstanceDecisionInstanceClient, err := mependingaccessreviewinstancedecisioninstance.NewMePendingAccessReviewInstanceDecisionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecisionInstance client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionInstanceClient.Client)

	mePendingAccessReviewInstanceDecisionInstanceContactedReviewerClient, err := mependingaccessreviewinstancedecisioninstancecontactedreviewer.NewMePendingAccessReviewInstanceDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecisionInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionInstanceContactedReviewerClient.Client)

	mePendingAccessReviewInstanceDecisionInstanceDefinitionClient, err := mependingaccessreviewinstancedecisioninstancedefinition.NewMePendingAccessReviewInstanceDecisionInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecisionInstanceDefinition client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionInstanceDefinitionClient.Client)

	mePendingAccessReviewInstanceDecisionInstanceStageClient, err := mependingaccessreviewinstancedecisioninstancestage.NewMePendingAccessReviewInstanceDecisionInstanceStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecisionInstanceStage client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionInstanceStageClient.Client)

	mePendingAccessReviewInstanceDecisionInstanceStageDecisionClient, err := mependingaccessreviewinstancedecisioninstancestagedecision.NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecisionInstanceStageDecision client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionInstanceStageDecisionClient.Client)

	mePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient, err := mependingaccessreviewinstancedecisioninstancestagedecisioninsight.NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsight client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient.Client)

	mePendingAccessReviewInstanceDefinitionClient, err := mependingaccessreviewinstancedefinition.NewMePendingAccessReviewInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceDefinition client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceDefinitionClient.Client)

	mePendingAccessReviewInstanceStageClient, err := mependingaccessreviewinstancestage.NewMePendingAccessReviewInstanceStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStage client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageClient.Client)

	mePendingAccessReviewInstanceStageDecisionClient, err := mependingaccessreviewinstancestagedecision.NewMePendingAccessReviewInstanceStageDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStageDecision client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageDecisionClient.Client)

	mePendingAccessReviewInstanceStageDecisionInsightClient, err := mependingaccessreviewinstancestagedecisioninsight.NewMePendingAccessReviewInstanceStageDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStageDecisionInsight client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageDecisionInsightClient.Client)

	mePendingAccessReviewInstanceStageDecisionInstanceClient, err := mependingaccessreviewinstancestagedecisioninstance.NewMePendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStageDecisionInstance client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageDecisionInstanceClient.Client)

	mePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient, err := mependingaccessreviewinstancestagedecisioninstancecontactedreviewer.NewMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient.Client)

	mePendingAccessReviewInstanceStageDecisionInstanceDecisionClient, err := mependingaccessreviewinstancestagedecisioninstancedecision.NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStageDecisionInstanceDecision client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageDecisionInstanceDecisionClient.Client)

	mePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient, err := mependingaccessreviewinstancestagedecisioninstancedecisioninsight.NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStageDecisionInstanceDecisionInsight client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient.Client)

	mePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient, err := mependingaccessreviewinstancestagedecisioninstancedefinition.NewMePendingAccessReviewInstanceStageDecisionInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePendingAccessReviewInstanceStageDecisionInstanceDefinition client: %+v", err)
	}
	configureFunc(mePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient.Client)

	mePeopleClient, err := mepeople.NewMePeopleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePeople client: %+v", err)
	}
	configureFunc(mePeopleClient.Client)

	mePermissionGrantClient, err := mepermissiongrant.NewMePermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePermissionGrant client: %+v", err)
	}
	configureFunc(mePermissionGrantClient.Client)

	mePhotoClient, err := mephoto.NewMePhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePhoto client: %+v", err)
	}
	configureFunc(mePhotoClient.Client)

	mePlannerAllClient, err := meplannerall.NewMePlannerAllClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerAll client: %+v", err)
	}
	configureFunc(mePlannerAllClient.Client)

	mePlannerClient, err := meplanner.NewMePlannerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlanner client: %+v", err)
	}
	configureFunc(mePlannerClient.Client)

	mePlannerFavoritePlanClient, err := meplannerfavoriteplan.NewMePlannerFavoritePlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerFavoritePlan client: %+v", err)
	}
	configureFunc(mePlannerFavoritePlanClient.Client)

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

	mePlannerRecentPlanClient, err := meplannerrecentplan.NewMePlannerRecentPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerRecentPlan client: %+v", err)
	}
	configureFunc(mePlannerRecentPlanClient.Client)

	mePlannerRosterPlanClient, err := meplannerrosterplan.NewMePlannerRosterPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MePlannerRosterPlan client: %+v", err)
	}
	configureFunc(mePlannerRosterPlanClient.Client)

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

	meProfileAccountClient, err := meprofileaccount.NewMeProfileAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileAccount client: %+v", err)
	}
	configureFunc(meProfileAccountClient.Client)

	meProfileAddressClient, err := meprofileaddress.NewMeProfileAddressClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileAddress client: %+v", err)
	}
	configureFunc(meProfileAddressClient.Client)

	meProfileAnniversaryClient, err := meprofileanniversary.NewMeProfileAnniversaryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileAnniversary client: %+v", err)
	}
	configureFunc(meProfileAnniversaryClient.Client)

	meProfileAwardClient, err := meprofileaward.NewMeProfileAwardClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileAward client: %+v", err)
	}
	configureFunc(meProfileAwardClient.Client)

	meProfileCertificationClient, err := meprofilecertification.NewMeProfileCertificationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileCertification client: %+v", err)
	}
	configureFunc(meProfileCertificationClient.Client)

	meProfileClient, err := meprofile.NewMeProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfile client: %+v", err)
	}
	configureFunc(meProfileClient.Client)

	meProfileEducationalActivityClient, err := meprofileeducationalactivity.NewMeProfileEducationalActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileEducationalActivity client: %+v", err)
	}
	configureFunc(meProfileEducationalActivityClient.Client)

	meProfileEmailClient, err := meprofileemail.NewMeProfileEmailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileEmail client: %+v", err)
	}
	configureFunc(meProfileEmailClient.Client)

	meProfileInterestClient, err := meprofileinterest.NewMeProfileInterestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileInterest client: %+v", err)
	}
	configureFunc(meProfileInterestClient.Client)

	meProfileLanguageClient, err := meprofilelanguage.NewMeProfileLanguageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileLanguage client: %+v", err)
	}
	configureFunc(meProfileLanguageClient.Client)

	meProfileNameClient, err := meprofilename.NewMeProfileNameClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileName client: %+v", err)
	}
	configureFunc(meProfileNameClient.Client)

	meProfileNoteClient, err := meprofilenote.NewMeProfileNoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileNote client: %+v", err)
	}
	configureFunc(meProfileNoteClient.Client)

	meProfilePatentClient, err := meprofilepatent.NewMeProfilePatentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfilePatent client: %+v", err)
	}
	configureFunc(meProfilePatentClient.Client)

	meProfilePhoneClient, err := meprofilephone.NewMeProfilePhoneClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfilePhone client: %+v", err)
	}
	configureFunc(meProfilePhoneClient.Client)

	meProfilePositionClient, err := meprofileposition.NewMeProfilePositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfilePosition client: %+v", err)
	}
	configureFunc(meProfilePositionClient.Client)

	meProfileProjectClient, err := meprofileproject.NewMeProfileProjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileProject client: %+v", err)
	}
	configureFunc(meProfileProjectClient.Client)

	meProfilePublicationClient, err := meprofilepublication.NewMeProfilePublicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfilePublication client: %+v", err)
	}
	configureFunc(meProfilePublicationClient.Client)

	meProfileSkillClient, err := meprofileskill.NewMeProfileSkillClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileSkill client: %+v", err)
	}
	configureFunc(meProfileSkillClient.Client)

	meProfileWebAccountClient, err := meprofilewebaccount.NewMeProfileWebAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileWebAccount client: %+v", err)
	}
	configureFunc(meProfileWebAccountClient.Client)

	meProfileWebsiteClient, err := meprofilewebsite.NewMeProfileWebsiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeProfileWebsite client: %+v", err)
	}
	configureFunc(meProfileWebsiteClient.Client)

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

	meSecurityClient, err := mesecurity.NewMeSecurityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSecurity client: %+v", err)
	}
	configureFunc(meSecurityClient.Client)

	meSecurityInformationProtectionClient, err := mesecurityinformationprotection.NewMeSecurityInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSecurityInformationProtection client: %+v", err)
	}
	configureFunc(meSecurityInformationProtectionClient.Client)

	meSecurityInformationProtectionLabelPolicySettingClient, err := mesecurityinformationprotectionlabelpolicysetting.NewMeSecurityInformationProtectionLabelPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSecurityInformationProtectionLabelPolicySetting client: %+v", err)
	}
	configureFunc(meSecurityInformationProtectionLabelPolicySettingClient.Client)

	meSecurityInformationProtectionSensitivityLabelClient, err := mesecurityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSecurityInformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(meSecurityInformationProtectionSensitivityLabelClient.Client)

	meSecurityInformationProtectionSensitivityLabelParentClient, err := mesecurityinformationprotectionsensitivitylabelparent.NewMeSecurityInformationProtectionSensitivityLabelParentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSecurityInformationProtectionSensitivityLabelParent client: %+v", err)
	}
	configureFunc(meSecurityInformationProtectionSensitivityLabelParentClient.Client)

	meSettingClient, err := mesetting.NewMeSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSetting client: %+v", err)
	}
	configureFunc(meSettingClient.Client)

	meSettingContactMergeSuggestionClient, err := mesettingcontactmergesuggestion.NewMeSettingContactMergeSuggestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSettingContactMergeSuggestion client: %+v", err)
	}
	configureFunc(meSettingContactMergeSuggestionClient.Client)

	meSettingItemInsightClient, err := mesettingiteminsight.NewMeSettingItemInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSettingItemInsight client: %+v", err)
	}
	configureFunc(meSettingItemInsightClient.Client)

	meSettingRegionalAndLanguageSettingClient, err := mesettingregionalandlanguagesetting.NewMeSettingRegionalAndLanguageSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSettingRegionalAndLanguageSetting client: %+v", err)
	}
	configureFunc(meSettingRegionalAndLanguageSettingClient.Client)

	meSettingShiftPreferenceClient, err := mesettingshiftpreference.NewMeSettingShiftPreferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSettingShiftPreference client: %+v", err)
	}
	configureFunc(meSettingShiftPreferenceClient.Client)

	meSponsorClient, err := mesponsor.NewMeSponsorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeSponsor client: %+v", err)
	}
	configureFunc(meSponsorClient.Client)

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

	meTransitiveReportClient, err := metransitivereport.NewMeTransitiveReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeTransitiveReport client: %+v", err)
	}
	configureFunc(meTransitiveReportClient.Client)

	meUsageRightClient, err := meusageright.NewMeUsageRightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeUsageRight client: %+v", err)
	}
	configureFunc(meUsageRightClient.Client)

	meWindowsInformationProtectionDeviceRegistrationClient, err := mewindowsinformationprotectiondeviceregistration.NewMeWindowsInformationProtectionDeviceRegistrationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MeWindowsInformationProtectionDeviceRegistration client: %+v", err)
	}
	configureFunc(meWindowsInformationProtectionDeviceRegistrationClient.Client)

	return &Client{
		Me:                              meClient,
		MeActivity:                      meActivityClient,
		MeActivityHistoryItem:           meActivityHistoryItemClient,
		MeActivityHistoryItemActivity:   meActivityHistoryItemActivityClient,
		MeAgreementAcceptance:           meAgreementAcceptanceClient,
		MeAnalytic:                      meAnalyticClient,
		MeAnalyticActivityStatistic:     meAnalyticActivityStatisticClient,
		MeAppConsentRequestsForApproval: meAppConsentRequestsForApprovalClient,
		MeAppConsentRequestsForApprovalUserConsentRequest:              meAppConsentRequestsForApprovalUserConsentRequestClient,
		MeAppConsentRequestsForApprovalUserConsentRequestApproval:      meAppConsentRequestsForApprovalUserConsentRequestApprovalClient,
		MeAppConsentRequestsForApprovalUserConsentRequestApprovalStep:  meAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient,
		MeAppRoleAssignedResource:                                      meAppRoleAssignedResourceClient,
		MeAppRoleAssignment:                                            meAppRoleAssignmentClient,
		MeApproval:                                                     meApprovalClient,
		MeApprovalStep:                                                 meApprovalStepClient,
		MeAuthentication:                                               meAuthenticationClient,
		MeAuthenticationEmailMethod:                                    meAuthenticationEmailMethodClient,
		MeAuthenticationFido2Method:                                    meAuthenticationFido2MethodClient,
		MeAuthenticationMethod:                                         meAuthenticationMethodClient,
		MeAuthenticationMicrosoftAuthenticatorMethod:                   meAuthenticationMicrosoftAuthenticatorMethodClient,
		MeAuthenticationMicrosoftAuthenticatorMethodDevice:             meAuthenticationMicrosoftAuthenticatorMethodDeviceClient,
		MeAuthenticationOperation:                                      meAuthenticationOperationClient,
		MeAuthenticationPasswordMethod:                                 meAuthenticationPasswordMethodClient,
		MeAuthenticationPasswordlessMicrosoftAuthenticatorMethod:       meAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient,
		MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDevice: meAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient,
		MeAuthenticationPhoneMethod:                                    meAuthenticationPhoneMethodClient,
		MeAuthenticationSoftwareOathMethod:                             meAuthenticationSoftwareOathMethodClient,
		MeAuthenticationTemporaryAccessPassMethod:                      meAuthenticationTemporaryAccessPassMethodClient,
		MeAuthenticationWindowsHelloForBusinessMethod:                  meAuthenticationWindowsHelloForBusinessMethodClient,
		MeAuthenticationWindowsHelloForBusinessMethodDevice:            meAuthenticationWindowsHelloForBusinessMethodDeviceClient,
		MeCalendar:                                                               meCalendarClient,
		MeCalendarCalendarPermission:                                             meCalendarCalendarPermissionClient,
		MeCalendarCalendarView:                                                   meCalendarCalendarViewClient,
		MeCalendarCalendarViewAttachment:                                         meCalendarCalendarViewAttachmentClient,
		MeCalendarCalendarViewCalendar:                                           meCalendarCalendarViewCalendarClient,
		MeCalendarCalendarViewExceptionOccurrence:                                meCalendarCalendarViewExceptionOccurrenceClient,
		MeCalendarCalendarViewExceptionOccurrenceAttachment:                      meCalendarCalendarViewExceptionOccurrenceAttachmentClient,
		MeCalendarCalendarViewExceptionOccurrenceCalendar:                        meCalendarCalendarViewExceptionOccurrenceCalendarClient,
		MeCalendarCalendarViewExceptionOccurrenceExtension:                       meCalendarCalendarViewExceptionOccurrenceExtensionClient,
		MeCalendarCalendarViewExceptionOccurrenceInstance:                        meCalendarCalendarViewExceptionOccurrenceInstanceClient,
		MeCalendarCalendarViewExceptionOccurrenceInstanceAttachment:              meCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		MeCalendarCalendarViewExceptionOccurrenceInstanceCalendar:                meCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		MeCalendarCalendarViewExceptionOccurrenceInstanceExtension:               meCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		MeCalendarCalendarViewExtension:                                          meCalendarCalendarViewExtensionClient,
		MeCalendarCalendarViewInstance:                                           meCalendarCalendarViewInstanceClient,
		MeCalendarCalendarViewInstanceAttachment:                                 meCalendarCalendarViewInstanceAttachmentClient,
		MeCalendarCalendarViewInstanceCalendar:                                   meCalendarCalendarViewInstanceCalendarClient,
		MeCalendarCalendarViewInstanceExceptionOccurrence:                        meCalendarCalendarViewInstanceExceptionOccurrenceClient,
		MeCalendarCalendarViewInstanceExceptionOccurrenceAttachment:              meCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		MeCalendarCalendarViewInstanceExceptionOccurrenceCalendar:                meCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		MeCalendarCalendarViewInstanceExceptionOccurrenceExtension:               meCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		MeCalendarCalendarViewInstanceExtension:                                  meCalendarCalendarViewInstanceExtensionClient,
		MeCalendarEvent:                                                          meCalendarEventClient,
		MeCalendarEventAttachment:                                                meCalendarEventAttachmentClient,
		MeCalendarEventCalendar:                                                  meCalendarEventCalendarClient,
		MeCalendarEventExceptionOccurrence:                                       meCalendarEventExceptionOccurrenceClient,
		MeCalendarEventExceptionOccurrenceAttachment:                             meCalendarEventExceptionOccurrenceAttachmentClient,
		MeCalendarEventExceptionOccurrenceCalendar:                               meCalendarEventExceptionOccurrenceCalendarClient,
		MeCalendarEventExceptionOccurrenceExtension:                              meCalendarEventExceptionOccurrenceExtensionClient,
		MeCalendarEventExceptionOccurrenceInstance:                               meCalendarEventExceptionOccurrenceInstanceClient,
		MeCalendarEventExceptionOccurrenceInstanceAttachment:                     meCalendarEventExceptionOccurrenceInstanceAttachmentClient,
		MeCalendarEventExceptionOccurrenceInstanceCalendar:                       meCalendarEventExceptionOccurrenceInstanceCalendarClient,
		MeCalendarEventExceptionOccurrenceInstanceExtension:                      meCalendarEventExceptionOccurrenceInstanceExtensionClient,
		MeCalendarEventExtension:                                                 meCalendarEventExtensionClient,
		MeCalendarEventInstance:                                                  meCalendarEventInstanceClient,
		MeCalendarEventInstanceAttachment:                                        meCalendarEventInstanceAttachmentClient,
		MeCalendarEventInstanceCalendar:                                          meCalendarEventInstanceCalendarClient,
		MeCalendarEventInstanceExceptionOccurrence:                               meCalendarEventInstanceExceptionOccurrenceClient,
		MeCalendarEventInstanceExceptionOccurrenceAttachment:                     meCalendarEventInstanceExceptionOccurrenceAttachmentClient,
		MeCalendarEventInstanceExceptionOccurrenceCalendar:                       meCalendarEventInstanceExceptionOccurrenceCalendarClient,
		MeCalendarEventInstanceExceptionOccurrenceExtension:                      meCalendarEventInstanceExceptionOccurrenceExtensionClient,
		MeCalendarEventInstanceExtension:                                         meCalendarEventInstanceExtensionClient,
		MeCalendarGroup:                                                          meCalendarGroupClient,
		MeCalendarGroupCalendar:                                                  meCalendarGroupCalendarClient,
		MeCalendarGroupCalendarCalendarPermission:                                meCalendarGroupCalendarCalendarPermissionClient,
		MeCalendarGroupCalendarCalendarView:                                      meCalendarGroupCalendarCalendarViewClient,
		MeCalendarGroupCalendarCalendarViewAttachment:                            meCalendarGroupCalendarCalendarViewAttachmentClient,
		MeCalendarGroupCalendarCalendarViewCalendar:                              meCalendarGroupCalendarCalendarViewCalendarClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrence:                   meCalendarGroupCalendarCalendarViewExceptionOccurrenceClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachment:         meCalendarGroupCalendarCalendarViewExceptionOccurrenceAttachmentClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendar:           meCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrenceExtension:          meCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstance:           meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment: meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar:   meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtension:  meCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		MeCalendarGroupCalendarCalendarViewExtension:                             meCalendarGroupCalendarCalendarViewExtensionClient,
		MeCalendarGroupCalendarCalendarViewInstance:                              meCalendarGroupCalendarCalendarViewInstanceClient,
		MeCalendarGroupCalendarCalendarViewInstanceAttachment:                    meCalendarGroupCalendarCalendarViewInstanceAttachmentClient,
		MeCalendarGroupCalendarCalendarViewInstanceCalendar:                      meCalendarGroupCalendarCalendarViewInstanceCalendarClient,
		MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrence:           meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient,
		MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment: meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar:   meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		MeCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtension:  meCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		MeCalendarGroupCalendarCalendarViewInstanceExtension:                     meCalendarGroupCalendarCalendarViewInstanceExtensionClient,
		MeCalendarGroupCalendarEvent:                                             meCalendarGroupCalendarEventClient,
		MeCalendarGroupCalendarEventAttachment:                                   meCalendarGroupCalendarEventAttachmentClient,
		MeCalendarGroupCalendarEventCalendar:                                     meCalendarGroupCalendarEventCalendarClient,
		MeCalendarGroupCalendarEventExceptionOccurrence:                          meCalendarGroupCalendarEventExceptionOccurrenceClient,
		MeCalendarGroupCalendarEventExceptionOccurrenceAttachment:                meCalendarGroupCalendarEventExceptionOccurrenceAttachmentClient,
		MeCalendarGroupCalendarEventExceptionOccurrenceCalendar:                  meCalendarGroupCalendarEventExceptionOccurrenceCalendarClient,
		MeCalendarGroupCalendarEventExceptionOccurrenceExtension:                 meCalendarGroupCalendarEventExceptionOccurrenceExtensionClient,
		MeCalendarGroupCalendarEventExceptionOccurrenceInstance:                  meCalendarGroupCalendarEventExceptionOccurrenceInstanceClient,
		MeCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachment:        meCalendarGroupCalendarEventExceptionOccurrenceInstanceAttachmentClient,
		MeCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendar:          meCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient,
		MeCalendarGroupCalendarEventExceptionOccurrenceInstanceExtension:         meCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient,
		MeCalendarGroupCalendarEventExtension:                                    meCalendarGroupCalendarEventExtensionClient,
		MeCalendarGroupCalendarEventInstance:                                     meCalendarGroupCalendarEventInstanceClient,
		MeCalendarGroupCalendarEventInstanceAttachment:                           meCalendarGroupCalendarEventInstanceAttachmentClient,
		MeCalendarGroupCalendarEventInstanceCalendar:                             meCalendarGroupCalendarEventInstanceCalendarClient,
		MeCalendarGroupCalendarEventInstanceExceptionOccurrence:                  meCalendarGroupCalendarEventInstanceExceptionOccurrenceClient,
		MeCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachment:        meCalendarGroupCalendarEventInstanceExceptionOccurrenceAttachmentClient,
		MeCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendar:          meCalendarGroupCalendarEventInstanceExceptionOccurrenceCalendarClient,
		MeCalendarGroupCalendarEventInstanceExceptionOccurrenceExtension:         meCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient,
		MeCalendarGroupCalendarEventInstanceExtension:                            meCalendarGroupCalendarEventInstanceExtensionClient,
		MeCalendarView:                                      meCalendarViewClient,
		MeCalendarViewAttachment:                            meCalendarViewAttachmentClient,
		MeCalendarViewCalendar:                              meCalendarViewCalendarClient,
		MeCalendarViewExceptionOccurrence:                   meCalendarViewExceptionOccurrenceClient,
		MeCalendarViewExceptionOccurrenceAttachment:         meCalendarViewExceptionOccurrenceAttachmentClient,
		MeCalendarViewExceptionOccurrenceCalendar:           meCalendarViewExceptionOccurrenceCalendarClient,
		MeCalendarViewExceptionOccurrenceExtension:          meCalendarViewExceptionOccurrenceExtensionClient,
		MeCalendarViewExceptionOccurrenceInstance:           meCalendarViewExceptionOccurrenceInstanceClient,
		MeCalendarViewExceptionOccurrenceInstanceAttachment: meCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		MeCalendarViewExceptionOccurrenceInstanceCalendar:   meCalendarViewExceptionOccurrenceInstanceCalendarClient,
		MeCalendarViewExceptionOccurrenceInstanceExtension:  meCalendarViewExceptionOccurrenceInstanceExtensionClient,
		MeCalendarViewExtension:                             meCalendarViewExtensionClient,
		MeCalendarViewInstance:                              meCalendarViewInstanceClient,
		MeCalendarViewInstanceAttachment:                    meCalendarViewInstanceAttachmentClient,
		MeCalendarViewInstanceCalendar:                      meCalendarViewInstanceCalendarClient,
		MeCalendarViewInstanceExceptionOccurrence:           meCalendarViewInstanceExceptionOccurrenceClient,
		MeCalendarViewInstanceExceptionOccurrenceAttachment: meCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		MeCalendarViewInstanceExceptionOccurrenceCalendar:   meCalendarViewInstanceExceptionOccurrenceCalendarClient,
		MeCalendarViewInstanceExceptionOccurrenceExtension:  meCalendarViewInstanceExceptionOccurrenceExtensionClient,
		MeCalendarViewInstanceExtension:                     meCalendarViewInstanceExtensionClient,
		MeChat:                                              meChatClient,
		MeChatInstalledApp:                                  meChatInstalledAppClient,
		MeChatInstalledAppTeamsApp:                          meChatInstalledAppTeamsAppClient,
		MeChatInstalledAppTeamsAppDefinition:                meChatInstalledAppTeamsAppDefinitionClient,
		MeChatLastMessagePreview:                            meChatLastMessagePreviewClient,
		MeChatMember:                                        meChatMemberClient,
		MeChatMessage:                                       meChatMessageClient,
		MeChatMessageHostedContent:                          meChatMessageHostedContentClient,
		MeChatMessageReply:                                  meChatMessageReplyClient,
		MeChatMessageReplyHostedContent:                     meChatMessageReplyHostedContentClient,
		MeChatOperation:                                     meChatOperationClient,
		MeChatPermissionGrant:                               meChatPermissionGrantClient,
		MeChatPinnedMessage:                                 meChatPinnedMessageClient,
		MeChatPinnedMessageMessage:                          meChatPinnedMessageMessageClient,
		MeChatTab:                                           meChatTabClient,
		MeChatTabTeamsApp:                                   meChatTabTeamsAppClient,
		MeCloudPC:                                           meCloudPCClient,
		MeContact:                                           meContactClient,
		MeContactExtension:                                  meContactExtensionClient,
		MeContactFolder:                                     meContactFolderClient,
		MeContactFolderChildFolder:                          meContactFolderChildFolderClient,
		MeContactFolderChildFolderContact:                   meContactFolderChildFolderContactClient,
		MeContactFolderChildFolderContactExtension:          meContactFolderChildFolderContactExtensionClient,
		MeContactFolderChildFolderContactPhoto:              meContactFolderChildFolderContactPhotoClient,
		MeContactFolderContact:                              meContactFolderContactClient,
		MeContactFolderContactExtension:                     meContactFolderContactExtensionClient,
		MeContactFolderContactPhoto:                         meContactFolderContactPhotoClient,
		MeContactPhoto:                                      meContactPhotoClient,
		MeCreatedObject:                                     meCreatedObjectClient,
		MeDevice:                                            meDeviceClient,
		MeDeviceCommand:                                     meDeviceCommandClient,
		MeDeviceCommandResponsepayload:                      meDeviceCommandResponsepayloadClient,
		MeDeviceEnrollmentConfiguration:                     meDeviceEnrollmentConfigurationClient,
		MeDeviceEnrollmentConfigurationAssignment:           meDeviceEnrollmentConfigurationAssignmentClient,
		MeDeviceExtension:                                   meDeviceExtensionClient,
		MeDeviceManagementTroubleshootingEvent:              meDeviceManagementTroubleshootingEventClient,
		MeDeviceMemberOf:                                    meDeviceMemberOfClient,
		MeDeviceRegisteredOwner:                             meDeviceRegisteredOwnerClient,
		MeDeviceRegisteredUser:                              meDeviceRegisteredUserClient,
		MeDeviceTransitiveMemberOf:                          meDeviceTransitiveMemberOfClient,
		MeDeviceUsageRight:                                  meDeviceUsageRightClient,
		MeDirectReport:                                      meDirectReportClient,
		MeDrive:                                             meDriveClient,
		MeEmployeeExperience:                                meEmployeeExperienceClient,
		MeEmployeeExperienceLearningCourseActivity:          meEmployeeExperienceLearningCourseActivityClient,
		MeEvent:                                              meEventClient,
		MeEventAttachment:                                    meEventAttachmentClient,
		MeEventCalendar:                                      meEventCalendarClient,
		MeEventExceptionOccurrence:                           meEventExceptionOccurrenceClient,
		MeEventExceptionOccurrenceAttachment:                 meEventExceptionOccurrenceAttachmentClient,
		MeEventExceptionOccurrenceCalendar:                   meEventExceptionOccurrenceCalendarClient,
		MeEventExceptionOccurrenceExtension:                  meEventExceptionOccurrenceExtensionClient,
		MeEventExceptionOccurrenceInstance:                   meEventExceptionOccurrenceInstanceClient,
		MeEventExceptionOccurrenceInstanceAttachment:         meEventExceptionOccurrenceInstanceAttachmentClient,
		MeEventExceptionOccurrenceInstanceCalendar:           meEventExceptionOccurrenceInstanceCalendarClient,
		MeEventExceptionOccurrenceInstanceExtension:          meEventExceptionOccurrenceInstanceExtensionClient,
		MeEventExtension:                                     meEventExtensionClient,
		MeEventInstance:                                      meEventInstanceClient,
		MeEventInstanceAttachment:                            meEventInstanceAttachmentClient,
		MeEventInstanceCalendar:                              meEventInstanceCalendarClient,
		MeEventInstanceExceptionOccurrence:                   meEventInstanceExceptionOccurrenceClient,
		MeEventInstanceExceptionOccurrenceAttachment:         meEventInstanceExceptionOccurrenceAttachmentClient,
		MeEventInstanceExceptionOccurrenceCalendar:           meEventInstanceExceptionOccurrenceCalendarClient,
		MeEventInstanceExceptionOccurrenceExtension:          meEventInstanceExceptionOccurrenceExtensionClient,
		MeEventInstanceExtension:                             meEventInstanceExtensionClient,
		MeExtension:                                          meExtensionClient,
		MeFollowedSite:                                       meFollowedSiteClient,
		MeInferenceClassification:                            meInferenceClassificationClient,
		MeInferenceClassificationOverride:                    meInferenceClassificationOverrideClient,
		MeInformationProtection:                              meInformationProtectionClient,
		MeInformationProtectionBitlocker:                     meInformationProtectionBitlockerClient,
		MeInformationProtectionBitlockerRecoveryKey:          meInformationProtectionBitlockerRecoveryKeyClient,
		MeInformationProtectionDataLossPreventionPolicy:      meInformationProtectionDataLossPreventionPolicyClient,
		MeInformationProtectionPolicy:                        meInformationProtectionPolicyClient,
		MeInformationProtectionPolicyLabel:                   meInformationProtectionPolicyLabelClient,
		MeInformationProtectionSensitivityLabel:              meInformationProtectionSensitivityLabelClient,
		MeInformationProtectionSensitivityLabelSublabel:      meInformationProtectionSensitivityLabelSublabelClient,
		MeInformationProtectionSensitivityPolicySetting:      meInformationProtectionSensitivityPolicySettingClient,
		MeInformationProtectionThreatAssessmentRequest:       meInformationProtectionThreatAssessmentRequestClient,
		MeInformationProtectionThreatAssessmentRequestResult: meInformationProtectionThreatAssessmentRequestResultClient,
		MeInsight:                                             meInsightClient,
		MeInsightShared:                                       meInsightSharedClient,
		MeInsightSharedLastSharedMethod:                       meInsightSharedLastSharedMethodClient,
		MeInsightSharedResource:                               meInsightSharedResourceClient,
		MeInsightTrending:                                     meInsightTrendingClient,
		MeInsightTrendingResource:                             meInsightTrendingResourceClient,
		MeInsightUsed:                                         meInsightUsedClient,
		MeInsightUsedResource:                                 meInsightUsedResourceClient,
		MeJoinedGroup:                                         meJoinedGroupClient,
		MeJoinedTeam:                                          meJoinedTeamClient,
		MeLicenseDetail:                                       meLicenseDetailClient,
		MeMailFolder:                                          meMailFolderClient,
		MeMailFolderChildFolder:                               meMailFolderChildFolderClient,
		MeMailFolderChildFolderMessage:                        meMailFolderChildFolderMessageClient,
		MeMailFolderChildFolderMessageAttachment:              meMailFolderChildFolderMessageAttachmentClient,
		MeMailFolderChildFolderMessageExtension:               meMailFolderChildFolderMessageExtensionClient,
		MeMailFolderChildFolderMessageMention:                 meMailFolderChildFolderMessageMentionClient,
		MeMailFolderChildFolderMessageRule:                    meMailFolderChildFolderMessageRuleClient,
		MeMailFolderChildFolderUserConfiguration:              meMailFolderChildFolderUserConfigurationClient,
		MeMailFolderMessage:                                   meMailFolderMessageClient,
		MeMailFolderMessageAttachment:                         meMailFolderMessageAttachmentClient,
		MeMailFolderMessageExtension:                          meMailFolderMessageExtensionClient,
		MeMailFolderMessageMention:                            meMailFolderMessageMentionClient,
		MeMailFolderMessageRule:                               meMailFolderMessageRuleClient,
		MeMailFolderUserConfiguration:                         meMailFolderUserConfigurationClient,
		MeMailboxSetting:                                      meMailboxSettingClient,
		MeManagedAppRegistration:                              meManagedAppRegistrationClient,
		MeManagedDevice:                                       meManagedDeviceClient,
		MeManagedDeviceAssignmentFilterEvaluationStatusDetail: meManagedDeviceAssignmentFilterEvaluationStatusDetailClient,
		MeManagedDeviceDetectedApp:                            meManagedDeviceDetectedAppClient,
		MeManagedDeviceDeviceCategory:                         meManagedDeviceDeviceCategoryClient,
		MeManagedDeviceDeviceCompliancePolicyState:            meManagedDeviceDeviceCompliancePolicyStateClient,
		MeManagedDeviceDeviceConfigurationState:               meManagedDeviceDeviceConfigurationStateClient,
		MeManagedDeviceDeviceHealthScriptState:                meManagedDeviceDeviceHealthScriptStateClient,
		MeManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId: meManagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient,
		MeManagedDeviceLogCollectionRequest:                                        meManagedDeviceLogCollectionRequestClient,
		MeManagedDeviceManagedDeviceMobileAppConfigurationState:                    meManagedDeviceManagedDeviceMobileAppConfigurationStateClient,
		MeManagedDeviceSecurityBaselineState:                                       meManagedDeviceSecurityBaselineStateClient,
		MeManagedDeviceSecurityBaselineStateSettingState:                           meManagedDeviceSecurityBaselineStateSettingStateClient,
		MeManagedDeviceUser:                                       meManagedDeviceUserClient,
		MeManagedDeviceWindowsProtectionState:                     meManagedDeviceWindowsProtectionStateClient,
		MeManagedDeviceWindowsProtectionStateDetectedMalwareState: meManagedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		MeManager:                       meManagerClient,
		MeMemberOf:                      meMemberOfClient,
		MeMessage:                       meMessageClient,
		MeMessageAttachment:             meMessageAttachmentClient,
		MeMessageExtension:              meMessageExtensionClient,
		MeMessageMention:                meMessageMentionClient,
		MeMobileAppIntentAndState:       meMobileAppIntentAndStateClient,
		MeMobileAppTroubleshootingEvent: meMobileAppTroubleshootingEventClient,
		MeMobileAppTroubleshootingEventAppLogCollectionRequest: meMobileAppTroubleshootingEventAppLogCollectionRequestClient,
		MeNotification:                                                      meNotificationClient,
		MeOauth2PermissionGrant:                                             meOauth2PermissionGrantClient,
		MeOnenote:                                                           meOnenoteClient,
		MeOnenoteNotebook:                                                   meOnenoteNotebookClient,
		MeOnenoteNotebookSection:                                            meOnenoteNotebookSectionClient,
		MeOnenoteNotebookSectionGroup:                                       meOnenoteNotebookSectionGroupClient,
		MeOnenoteNotebookSectionGroupParentNotebook:                         meOnenoteNotebookSectionGroupParentNotebookClient,
		MeOnenoteNotebookSectionGroupParentSectionGroup:                     meOnenoteNotebookSectionGroupParentSectionGroupClient,
		MeOnenoteNotebookSectionGroupSection:                                meOnenoteNotebookSectionGroupSectionClient,
		MeOnenoteNotebookSectionGroupSectionGroup:                           meOnenoteNotebookSectionGroupSectionGroupClient,
		MeOnenoteNotebookSectionGroupSectionPage:                            meOnenoteNotebookSectionGroupSectionPageClient,
		MeOnenoteNotebookSectionGroupSectionPageContent:                     meOnenoteNotebookSectionGroupSectionPageContentClient,
		MeOnenoteNotebookSectionGroupSectionPageParentNotebook:              meOnenoteNotebookSectionGroupSectionPageParentNotebookClient,
		MeOnenoteNotebookSectionGroupSectionPageParentSection:               meOnenoteNotebookSectionGroupSectionPageParentSectionClient,
		MeOnenoteNotebookSectionGroupSectionParentNotebook:                  meOnenoteNotebookSectionGroupSectionParentNotebookClient,
		MeOnenoteNotebookSectionGroupSectionParentSectionGroup:              meOnenoteNotebookSectionGroupSectionParentSectionGroupClient,
		MeOnenoteNotebookSectionPage:                                        meOnenoteNotebookSectionPageClient,
		MeOnenoteNotebookSectionPageContent:                                 meOnenoteNotebookSectionPageContentClient,
		MeOnenoteNotebookSectionPageParentNotebook:                          meOnenoteNotebookSectionPageParentNotebookClient,
		MeOnenoteNotebookSectionPageParentSection:                           meOnenoteNotebookSectionPageParentSectionClient,
		MeOnenoteNotebookSectionParentNotebook:                              meOnenoteNotebookSectionParentNotebookClient,
		MeOnenoteNotebookSectionParentSectionGroup:                          meOnenoteNotebookSectionParentSectionGroupClient,
		MeOnenoteOperation:                                                  meOnenoteOperationClient,
		MeOnenotePage:                                                       meOnenotePageClient,
		MeOnenotePageContent:                                                meOnenotePageContentClient,
		MeOnenotePageParentNotebook:                                         meOnenotePageParentNotebookClient,
		MeOnenotePageParentSection:                                          meOnenotePageParentSectionClient,
		MeOnenoteResource:                                                   meOnenoteResourceClient,
		MeOnenoteResourceContent:                                            meOnenoteResourceContentClient,
		MeOnenoteSection:                                                    meOnenoteSectionClient,
		MeOnenoteSectionGroup:                                               meOnenoteSectionGroupClient,
		MeOnenoteSectionGroupParentNotebook:                                 meOnenoteSectionGroupParentNotebookClient,
		MeOnenoteSectionGroupParentSectionGroup:                             meOnenoteSectionGroupParentSectionGroupClient,
		MeOnenoteSectionGroupSection:                                        meOnenoteSectionGroupSectionClient,
		MeOnenoteSectionGroupSectionGroup:                                   meOnenoteSectionGroupSectionGroupClient,
		MeOnenoteSectionGroupSectionPage:                                    meOnenoteSectionGroupSectionPageClient,
		MeOnenoteSectionGroupSectionPageContent:                             meOnenoteSectionGroupSectionPageContentClient,
		MeOnenoteSectionGroupSectionPageParentNotebook:                      meOnenoteSectionGroupSectionPageParentNotebookClient,
		MeOnenoteSectionGroupSectionPageParentSection:                       meOnenoteSectionGroupSectionPageParentSectionClient,
		MeOnenoteSectionGroupSectionParentNotebook:                          meOnenoteSectionGroupSectionParentNotebookClient,
		MeOnenoteSectionGroupSectionParentSectionGroup:                      meOnenoteSectionGroupSectionParentSectionGroupClient,
		MeOnenoteSectionPage:                                                meOnenoteSectionPageClient,
		MeOnenoteSectionPageContent:                                         meOnenoteSectionPageContentClient,
		MeOnenoteSectionPageParentNotebook:                                  meOnenoteSectionPageParentNotebookClient,
		MeOnenoteSectionPageParentSection:                                   meOnenoteSectionPageParentSectionClient,
		MeOnenoteSectionParentNotebook:                                      meOnenoteSectionParentNotebookClient,
		MeOnenoteSectionParentSectionGroup:                                  meOnenoteSectionParentSectionGroupClient,
		MeOnlineMeeting:                                                     meOnlineMeetingClient,
		MeOnlineMeetingAlternativeRecording:                                 meOnlineMeetingAlternativeRecordingClient,
		MeOnlineMeetingAttendanceReport:                                     meOnlineMeetingAttendanceReportClient,
		MeOnlineMeetingAttendanceReportAttendanceRecord:                     meOnlineMeetingAttendanceReportAttendanceRecordClient,
		MeOnlineMeetingAttendeeReport:                                       meOnlineMeetingAttendeeReportClient,
		MeOnlineMeetingBroadcastRecording:                                   meOnlineMeetingBroadcastRecordingClient,
		MeOnlineMeetingMeetingAttendanceReport:                              meOnlineMeetingMeetingAttendanceReportClient,
		MeOnlineMeetingMeetingAttendanceReportAttendanceRecord:              meOnlineMeetingMeetingAttendanceReportAttendanceRecordClient,
		MeOnlineMeetingRecording:                                            meOnlineMeetingRecordingClient,
		MeOnlineMeetingRecordingContent:                                     meOnlineMeetingRecordingContentClient,
		MeOnlineMeetingRegistration:                                         meOnlineMeetingRegistrationClient,
		MeOnlineMeetingRegistrationCustomQuestion:                           meOnlineMeetingRegistrationCustomQuestionClient,
		MeOnlineMeetingRegistrationRegistrant:                               meOnlineMeetingRegistrationRegistrantClient,
		MeOnlineMeetingTranscript:                                           meOnlineMeetingTranscriptClient,
		MeOnlineMeetingTranscriptContent:                                    meOnlineMeetingTranscriptContentClient,
		MeOnlineMeetingTranscriptMetadataContent:                            meOnlineMeetingTranscriptMetadataContentClient,
		MeOutlook:                                                           meOutlookClient,
		MeOutlookMasterCategory:                                             meOutlookMasterCategoryClient,
		MeOutlookTask:                                                       meOutlookTaskClient,
		MeOutlookTaskAttachment:                                             meOutlookTaskAttachmentClient,
		MeOutlookTaskFolder:                                                 meOutlookTaskFolderClient,
		MeOutlookTaskFolderTask:                                             meOutlookTaskFolderTaskClient,
		MeOutlookTaskFolderTaskAttachment:                                   meOutlookTaskFolderTaskAttachmentClient,
		MeOutlookTaskGroup:                                                  meOutlookTaskGroupClient,
		MeOutlookTaskGroupTaskFolder:                                        meOutlookTaskGroupTaskFolderClient,
		MeOutlookTaskGroupTaskFolderTask:                                    meOutlookTaskGroupTaskFolderTaskClient,
		MeOutlookTaskGroupTaskFolderTaskAttachment:                          meOutlookTaskGroupTaskFolderTaskAttachmentClient,
		MeOwnedDevice:                                                       meOwnedDeviceClient,
		MeOwnedObject:                                                       meOwnedObjectClient,
		MePendingAccessReviewInstance:                                       mePendingAccessReviewInstanceClient,
		MePendingAccessReviewInstanceContactedReviewer:                      mePendingAccessReviewInstanceContactedReviewerClient,
		MePendingAccessReviewInstanceDecision:                               mePendingAccessReviewInstanceDecisionClient,
		MePendingAccessReviewInstanceDecisionInsight:                        mePendingAccessReviewInstanceDecisionInsightClient,
		MePendingAccessReviewInstanceDecisionInstance:                       mePendingAccessReviewInstanceDecisionInstanceClient,
		MePendingAccessReviewInstanceDecisionInstanceContactedReviewer:      mePendingAccessReviewInstanceDecisionInstanceContactedReviewerClient,
		MePendingAccessReviewInstanceDecisionInstanceDefinition:             mePendingAccessReviewInstanceDecisionInstanceDefinitionClient,
		MePendingAccessReviewInstanceDecisionInstanceStage:                  mePendingAccessReviewInstanceDecisionInstanceStageClient,
		MePendingAccessReviewInstanceDecisionInstanceStageDecision:          mePendingAccessReviewInstanceDecisionInstanceStageDecisionClient,
		MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsight:   mePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient,
		MePendingAccessReviewInstanceDefinition:                             mePendingAccessReviewInstanceDefinitionClient,
		MePendingAccessReviewInstanceStage:                                  mePendingAccessReviewInstanceStageClient,
		MePendingAccessReviewInstanceStageDecision:                          mePendingAccessReviewInstanceStageDecisionClient,
		MePendingAccessReviewInstanceStageDecisionInsight:                   mePendingAccessReviewInstanceStageDecisionInsightClient,
		MePendingAccessReviewInstanceStageDecisionInstance:                  mePendingAccessReviewInstanceStageDecisionInstanceClient,
		MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewer: mePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient,
		MePendingAccessReviewInstanceStageDecisionInstanceDecision:          mePendingAccessReviewInstanceStageDecisionInstanceDecisionClient,
		MePendingAccessReviewInstanceStageDecisionInstanceDecisionInsight:   mePendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient,
		MePendingAccessReviewInstanceStageDecisionInstanceDefinition:        mePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient,
		MePeople:                mePeopleClient,
		MePermissionGrant:       mePermissionGrantClient,
		MePhoto:                 mePhotoClient,
		MePlanner:               mePlannerClient,
		MePlannerAll:            mePlannerAllClient,
		MePlannerFavoritePlan:   mePlannerFavoritePlanClient,
		MePlannerPlan:           mePlannerPlanClient,
		MePlannerPlanBucket:     mePlannerPlanBucketClient,
		MePlannerPlanBucketTask: mePlannerPlanBucketTaskClient,
		MePlannerPlanBucketTaskAssignedToTaskBoardFormat:  mePlannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		MePlannerPlanBucketTaskBucketTaskBoardFormat:      mePlannerPlanBucketTaskBucketTaskBoardFormatClient,
		MePlannerPlanBucketTaskDetail:                     mePlannerPlanBucketTaskDetailClient,
		MePlannerPlanBucketTaskProgressTaskBoardFormat:    mePlannerPlanBucketTaskProgressTaskBoardFormatClient,
		MePlannerPlanDetail:                               mePlannerPlanDetailClient,
		MePlannerPlanTask:                                 mePlannerPlanTaskClient,
		MePlannerPlanTaskAssignedToTaskBoardFormat:        mePlannerPlanTaskAssignedToTaskBoardFormatClient,
		MePlannerPlanTaskBucketTaskBoardFormat:            mePlannerPlanTaskBucketTaskBoardFormatClient,
		MePlannerPlanTaskDetail:                           mePlannerPlanTaskDetailClient,
		MePlannerPlanTaskProgressTaskBoardFormat:          mePlannerPlanTaskProgressTaskBoardFormatClient,
		MePlannerRecentPlan:                               mePlannerRecentPlanClient,
		MePlannerRosterPlan:                               mePlannerRosterPlanClient,
		MePlannerTask:                                     mePlannerTaskClient,
		MePlannerTaskAssignedToTaskBoardFormat:            mePlannerTaskAssignedToTaskBoardFormatClient,
		MePlannerTaskBucketTaskBoardFormat:                mePlannerTaskBucketTaskBoardFormatClient,
		MePlannerTaskDetail:                               mePlannerTaskDetailClient,
		MePlannerTaskProgressTaskBoardFormat:              mePlannerTaskProgressTaskBoardFormatClient,
		MePresence:                                        mePresenceClient,
		MeProfile:                                         meProfileClient,
		MeProfileAccount:                                  meProfileAccountClient,
		MeProfileAddress:                                  meProfileAddressClient,
		MeProfileAnniversary:                              meProfileAnniversaryClient,
		MeProfileAward:                                    meProfileAwardClient,
		MeProfileCertification:                            meProfileCertificationClient,
		MeProfileEducationalActivity:                      meProfileEducationalActivityClient,
		MeProfileEmail:                                    meProfileEmailClient,
		MeProfileInterest:                                 meProfileInterestClient,
		MeProfileLanguage:                                 meProfileLanguageClient,
		MeProfileName:                                     meProfileNameClient,
		MeProfileNote:                                     meProfileNoteClient,
		MeProfilePatent:                                   meProfilePatentClient,
		MeProfilePhone:                                    meProfilePhoneClient,
		MeProfilePosition:                                 meProfilePositionClient,
		MeProfileProject:                                  meProfileProjectClient,
		MeProfilePublication:                              meProfilePublicationClient,
		MeProfileSkill:                                    meProfileSkillClient,
		MeProfileWebAccount:                               meProfileWebAccountClient,
		MeProfileWebsite:                                  meProfileWebsiteClient,
		MeRegisteredDevice:                                meRegisteredDeviceClient,
		MeScopedRoleMemberOf:                              meScopedRoleMemberOfClient,
		MeSecurity:                                        meSecurityClient,
		MeSecurityInformationProtection:                   meSecurityInformationProtectionClient,
		MeSecurityInformationProtectionLabelPolicySetting: meSecurityInformationProtectionLabelPolicySettingClient,
		MeSecurityInformationProtectionSensitivityLabel:   meSecurityInformationProtectionSensitivityLabelClient,
		MeSecurityInformationProtectionSensitivityLabelParent: meSecurityInformationProtectionSensitivityLabelParentClient,
		MeSetting:                                        meSettingClient,
		MeSettingContactMergeSuggestion:                  meSettingContactMergeSuggestionClient,
		MeSettingItemInsight:                             meSettingItemInsightClient,
		MeSettingRegionalAndLanguageSetting:              meSettingRegionalAndLanguageSettingClient,
		MeSettingShiftPreference:                         meSettingShiftPreferenceClient,
		MeSponsor:                                        meSponsorClient,
		MeTeamwork:                                       meTeamworkClient,
		MeTeamworkAssociatedTeam:                         meTeamworkAssociatedTeamClient,
		MeTeamworkAssociatedTeamTeam:                     meTeamworkAssociatedTeamTeamClient,
		MeTeamworkInstalledApp:                           meTeamworkInstalledAppClient,
		MeTeamworkInstalledAppChat:                       meTeamworkInstalledAppChatClient,
		MeTeamworkInstalledAppTeamsApp:                   meTeamworkInstalledAppTeamsAppClient,
		MeTeamworkInstalledAppTeamsAppDefinition:         meTeamworkInstalledAppTeamsAppDefinitionClient,
		MeTodo:                                           meTodoClient,
		MeTodoList:                                       meTodoListClient,
		MeTodoListExtension:                              meTodoListExtensionClient,
		MeTodoListTask:                                   meTodoListTaskClient,
		MeTodoListTaskAttachment:                         meTodoListTaskAttachmentClient,
		MeTodoListTaskAttachmentSession:                  meTodoListTaskAttachmentSessionClient,
		MeTodoListTaskAttachmentSessionContent:           meTodoListTaskAttachmentSessionContentClient,
		MeTodoListTaskChecklistItem:                      meTodoListTaskChecklistItemClient,
		MeTodoListTaskExtension:                          meTodoListTaskExtensionClient,
		MeTodoListTaskLinkedResource:                     meTodoListTaskLinkedResourceClient,
		MeTransitiveMemberOf:                             meTransitiveMemberOfClient,
		MeTransitiveReport:                               meTransitiveReportClient,
		MeUsageRight:                                     meUsageRightClient,
		MeWindowsInformationProtectionDeviceRegistration: meWindowsInformationProtectionDeviceRegistrationClient,
	}, nil
}
