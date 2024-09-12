package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/activity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/activityhistoryitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/activityhistoryitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/agreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/approleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authentication"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationemailmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationfido2method"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationmicrosoftauthenticatormethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationmicrosoftauthenticatormethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationpasswordmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationphonemethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationsoftwareoathmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationtemporaryaccesspassmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationwindowshelloforbusinessmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationwindowshelloforbusinessmethoddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendargroupcalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/calendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatlastmessagepreview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatpermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatpinnedmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chatpinnedmessagemessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chattab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/chattabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/cloudclipboard"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/cloudclipboarditem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfolderchildfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfolderchildfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfolderchildfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfoldercontact"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfoldercontactextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactfoldercontactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/contactphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/createdobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/devicemanagementtroubleshootingevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/directreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivebundle"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivebundlecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivefollowing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivefollowingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveitemversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivelistsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driveroot"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/driverootversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivespecial"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/drivespecialcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/employeeexperience"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/employeeexperiencelearningcourseactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/event"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/eventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/eventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/eventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/eventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/eventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/eventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/eventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/extension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/followedsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/inferenceclassification"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/inferenceclassificationoverride"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insightshared"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insightsharedlastsharedmethod"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insightsharedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insighttrending"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insighttrendingresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insightused"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/insightusedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamallchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamchanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamgroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamincomingchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteaminstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteaminstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteaminstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteammember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteampermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamprimarychanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduleoffershiftrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduleopenshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduleopenshiftchangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduleschedulinggroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduleshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduleswapshiftschangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduletimeoffreason"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduletimeoffrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamscheduletimesoff"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamtag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamtagmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/joinedteamtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/licensedetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfolderchildfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/mailfoldermessagerule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/managedappregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddevicelogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddeviceuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddevicewindowsprotectionstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manageddevicewindowsprotectionstatedetectedmalwarestate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/memberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/message"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/messageattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/messageextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/oauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeeting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingattendancereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingattendancereportattendancerecord"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingattendeereport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingrecording"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingrecordingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingtranscript"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingtranscriptcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/onlinemeetingtranscriptmetadatacontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/outlook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/outlookmastercategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/owneddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/ownedobject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/person"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/photo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/planner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannertask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannertaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannertaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannertaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/plannertaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/presence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/registereddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/scopedrolememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/serviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/setting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/settingiteminsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/settingshiftpreference"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/settingstorage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/settingstoragequota"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/settingstoragequotaservice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/settingwindow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/settingwindowinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/sponsor"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamwork"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamworkassociatedteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamworkassociatedteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamworkinstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamworkinstalledappchat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamworkinstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/teamworkinstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolistextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttaskattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttaskattachmentsession"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttaskattachmentsessioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttaskchecklistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttaskextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/todolisttasklinkedresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/transitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Activity                                                    *activity.ActivityClient
	ActivityHistoryItem                                         *activityhistoryitem.ActivityHistoryItemClient
	ActivityHistoryItemActivity                                 *activityhistoryitemactivity.ActivityHistoryItemActivityClient
	AgreementAcceptance                                         *agreementacceptance.AgreementAcceptanceClient
	AppRoleAssignment                                           *approleassignment.AppRoleAssignmentClient
	Authentication                                              *authentication.AuthenticationClient
	AuthenticationEmailMethod                                   *authenticationemailmethod.AuthenticationEmailMethodClient
	AuthenticationFido2Method                                   *authenticationfido2method.AuthenticationFido2MethodClient
	AuthenticationMethod                                        *authenticationmethod.AuthenticationMethodClient
	AuthenticationMicrosoftAuthenticatorMethod                  *authenticationmicrosoftauthenticatormethod.AuthenticationMicrosoftAuthenticatorMethodClient
	AuthenticationMicrosoftAuthenticatorMethodDevice            *authenticationmicrosoftauthenticatormethoddevice.AuthenticationMicrosoftAuthenticatorMethodDeviceClient
	AuthenticationOperation                                     *authenticationoperation.AuthenticationOperationClient
	AuthenticationPasswordMethod                                *authenticationpasswordmethod.AuthenticationPasswordMethodClient
	AuthenticationPhoneMethod                                   *authenticationphonemethod.AuthenticationPhoneMethodClient
	AuthenticationSoftwareOathMethod                            *authenticationsoftwareoathmethod.AuthenticationSoftwareOathMethodClient
	AuthenticationTemporaryAccessPassMethod                     *authenticationtemporaryaccesspassmethod.AuthenticationTemporaryAccessPassMethodClient
	AuthenticationWindowsHelloForBusinessMethod                 *authenticationwindowshelloforbusinessmethod.AuthenticationWindowsHelloForBusinessMethodClient
	AuthenticationWindowsHelloForBusinessMethodDevice           *authenticationwindowshelloforbusinessmethoddevice.AuthenticationWindowsHelloForBusinessMethodDeviceClient
	Calendar                                                    *calendar.CalendarClient
	CalendarCalendarPermission                                  *calendarcalendarpermission.CalendarCalendarPermissionClient
	CalendarCalendarView                                        *calendarcalendarview.CalendarCalendarViewClient
	CalendarCalendarViewAttachment                              *calendarcalendarviewattachment.CalendarCalendarViewAttachmentClient
	CalendarCalendarViewCalendar                                *calendarcalendarviewcalendar.CalendarCalendarViewCalendarClient
	CalendarCalendarViewExtension                               *calendarcalendarviewextension.CalendarCalendarViewExtensionClient
	CalendarCalendarViewInstance                                *calendarcalendarviewinstance.CalendarCalendarViewInstanceClient
	CalendarCalendarViewInstanceAttachment                      *calendarcalendarviewinstanceattachment.CalendarCalendarViewInstanceAttachmentClient
	CalendarCalendarViewInstanceCalendar                        *calendarcalendarviewinstancecalendar.CalendarCalendarViewInstanceCalendarClient
	CalendarCalendarViewInstanceExtension                       *calendarcalendarviewinstanceextension.CalendarCalendarViewInstanceExtensionClient
	CalendarEvent                                               *calendarevent.CalendarEventClient
	CalendarEventAttachment                                     *calendareventattachment.CalendarEventAttachmentClient
	CalendarEventCalendar                                       *calendareventcalendar.CalendarEventCalendarClient
	CalendarEventExtension                                      *calendareventextension.CalendarEventExtensionClient
	CalendarEventInstance                                       *calendareventinstance.CalendarEventInstanceClient
	CalendarEventInstanceAttachment                             *calendareventinstanceattachment.CalendarEventInstanceAttachmentClient
	CalendarEventInstanceCalendar                               *calendareventinstancecalendar.CalendarEventInstanceCalendarClient
	CalendarEventInstanceExtension                              *calendareventinstanceextension.CalendarEventInstanceExtensionClient
	CalendarGroup                                               *calendargroup.CalendarGroupClient
	CalendarGroupCalendar                                       *calendargroupcalendar.CalendarGroupCalendarClient
	CalendarGroupCalendarCalendarPermission                     *calendargroupcalendarcalendarpermission.CalendarGroupCalendarCalendarPermissionClient
	CalendarGroupCalendarCalendarView                           *calendargroupcalendarcalendarview.CalendarGroupCalendarCalendarViewClient
	CalendarGroupCalendarCalendarViewAttachment                 *calendargroupcalendarcalendarviewattachment.CalendarGroupCalendarCalendarViewAttachmentClient
	CalendarGroupCalendarCalendarViewCalendar                   *calendargroupcalendarcalendarviewcalendar.CalendarGroupCalendarCalendarViewCalendarClient
	CalendarGroupCalendarCalendarViewExtension                  *calendargroupcalendarcalendarviewextension.CalendarGroupCalendarCalendarViewExtensionClient
	CalendarGroupCalendarCalendarViewInstance                   *calendargroupcalendarcalendarviewinstance.CalendarGroupCalendarCalendarViewInstanceClient
	CalendarGroupCalendarCalendarViewInstanceAttachment         *calendargroupcalendarcalendarviewinstanceattachment.CalendarGroupCalendarCalendarViewInstanceAttachmentClient
	CalendarGroupCalendarCalendarViewInstanceCalendar           *calendargroupcalendarcalendarviewinstancecalendar.CalendarGroupCalendarCalendarViewInstanceCalendarClient
	CalendarGroupCalendarCalendarViewInstanceExtension          *calendargroupcalendarcalendarviewinstanceextension.CalendarGroupCalendarCalendarViewInstanceExtensionClient
	CalendarGroupCalendarEvent                                  *calendargroupcalendarevent.CalendarGroupCalendarEventClient
	CalendarGroupCalendarEventAttachment                        *calendargroupcalendareventattachment.CalendarGroupCalendarEventAttachmentClient
	CalendarGroupCalendarEventCalendar                          *calendargroupcalendareventcalendar.CalendarGroupCalendarEventCalendarClient
	CalendarGroupCalendarEventExtension                         *calendargroupcalendareventextension.CalendarGroupCalendarEventExtensionClient
	CalendarGroupCalendarEventInstance                          *calendargroupcalendareventinstance.CalendarGroupCalendarEventInstanceClient
	CalendarGroupCalendarEventInstanceAttachment                *calendargroupcalendareventinstanceattachment.CalendarGroupCalendarEventInstanceAttachmentClient
	CalendarGroupCalendarEventInstanceCalendar                  *calendargroupcalendareventinstancecalendar.CalendarGroupCalendarEventInstanceCalendarClient
	CalendarGroupCalendarEventInstanceExtension                 *calendargroupcalendareventinstanceextension.CalendarGroupCalendarEventInstanceExtensionClient
	CalendarView                                                *calendarview.CalendarViewClient
	CalendarViewAttachment                                      *calendarviewattachment.CalendarViewAttachmentClient
	CalendarViewCalendar                                        *calendarviewcalendar.CalendarViewCalendarClient
	CalendarViewExtension                                       *calendarviewextension.CalendarViewExtensionClient
	CalendarViewInstance                                        *calendarviewinstance.CalendarViewInstanceClient
	CalendarViewInstanceAttachment                              *calendarviewinstanceattachment.CalendarViewInstanceAttachmentClient
	CalendarViewInstanceCalendar                                *calendarviewinstancecalendar.CalendarViewInstanceCalendarClient
	CalendarViewInstanceExtension                               *calendarviewinstanceextension.CalendarViewInstanceExtensionClient
	Chat                                                        *chat.ChatClient
	ChatInstalledApp                                            *chatinstalledapp.ChatInstalledAppClient
	ChatInstalledAppTeamsApp                                    *chatinstalledappteamsapp.ChatInstalledAppTeamsAppClient
	ChatInstalledAppTeamsAppDefinition                          *chatinstalledappteamsappdefinition.ChatInstalledAppTeamsAppDefinitionClient
	ChatLastMessagePreview                                      *chatlastmessagepreview.ChatLastMessagePreviewClient
	ChatMember                                                  *chatmember.ChatMemberClient
	ChatMessage                                                 *chatmessage.ChatMessageClient
	ChatMessageHostedContent                                    *chatmessagehostedcontent.ChatMessageHostedContentClient
	ChatMessageReply                                            *chatmessagereply.ChatMessageReplyClient
	ChatMessageReplyHostedContent                               *chatmessagereplyhostedcontent.ChatMessageReplyHostedContentClient
	ChatPermissionGrant                                         *chatpermissiongrant.ChatPermissionGrantClient
	ChatPinnedMessage                                           *chatpinnedmessage.ChatPinnedMessageClient
	ChatPinnedMessageMessage                                    *chatpinnedmessagemessage.ChatPinnedMessageMessageClient
	ChatTab                                                     *chattab.ChatTabClient
	ChatTabTeamsApp                                             *chattabteamsapp.ChatTabTeamsAppClient
	CloudClipboard                                              *cloudclipboard.CloudClipboardClient
	CloudClipboardItem                                          *cloudclipboarditem.CloudClipboardItemClient
	Contact                                                     *contact.ContactClient
	ContactExtension                                            *contactextension.ContactExtensionClient
	ContactFolder                                               *contactfolder.ContactFolderClient
	ContactFolderChildFolder                                    *contactfolderchildfolder.ContactFolderChildFolderClient
	ContactFolderChildFolderContact                             *contactfolderchildfoldercontact.ContactFolderChildFolderContactClient
	ContactFolderChildFolderContactExtension                    *contactfolderchildfoldercontactextension.ContactFolderChildFolderContactExtensionClient
	ContactFolderChildFolderContactPhoto                        *contactfolderchildfoldercontactphoto.ContactFolderChildFolderContactPhotoClient
	ContactFolderContact                                        *contactfoldercontact.ContactFolderContactClient
	ContactFolderContactExtension                               *contactfoldercontactextension.ContactFolderContactExtensionClient
	ContactFolderContactPhoto                                   *contactfoldercontactphoto.ContactFolderContactPhotoClient
	ContactPhoto                                                *contactphoto.ContactPhotoClient
	CreatedObject                                               *createdobject.CreatedObjectClient
	DeviceManagementTroubleshootingEvent                        *devicemanagementtroubleshootingevent.DeviceManagementTroubleshootingEventClient
	DirectReport                                                *directreport.DirectReportClient
	Drive                                                       *drive.DriveClient
	DriveBundle                                                 *drivebundle.DriveBundleClient
	DriveBundleContent                                          *drivebundlecontent.DriveBundleContentClient
	DriveCreatedByUser                                          *drivecreatedbyuser.DriveCreatedByUserClient
	DriveCreatedByUserMailboxSetting                            *drivecreatedbyusermailboxsetting.DriveCreatedByUserMailboxSettingClient
	DriveCreatedByUserServiceProvisioningError                  *drivecreatedbyuserserviceprovisioningerror.DriveCreatedByUserServiceProvisioningErrorClient
	DriveFollowing                                              *drivefollowing.DriveFollowingClient
	DriveFollowingContent                                       *drivefollowingcontent.DriveFollowingContentClient
	DriveItem                                                   *driveitem.DriveItemClient
	DriveItemAnalytics                                          *driveitemanalytics.DriveItemAnalyticsClient
	DriveItemAnalyticsAllTime                                   *driveitemanalyticsalltime.DriveItemAnalyticsAllTimeClient
	DriveItemAnalyticsItemActivityStat                          *driveitemanalyticsitemactivitystat.DriveItemAnalyticsItemActivityStatClient
	DriveItemAnalyticsItemActivityStatActivity                  *driveitemanalyticsitemactivitystatactivity.DriveItemAnalyticsItemActivityStatActivityClient
	DriveItemAnalyticsItemActivityStatActivityDriveItem         *driveitemanalyticsitemactivitystatactivitydriveitem.DriveItemAnalyticsItemActivityStatActivityDriveItemClient
	DriveItemAnalyticsItemActivityStatActivityDriveItemContent  *driveitemanalyticsitemactivitystatactivitydriveitemcontent.DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient
	DriveItemAnalyticsLastSevenDay                              *driveitemanalyticslastsevenday.DriveItemAnalyticsLastSevenDayClient
	DriveItemChild                                              *driveitemchild.DriveItemChildClient
	DriveItemChildContent                                       *driveitemchildcontent.DriveItemChildContentClient
	DriveItemContent                                            *driveitemcontent.DriveItemContentClient
	DriveItemCreatedByUser                                      *driveitemcreatedbyuser.DriveItemCreatedByUserClient
	DriveItemCreatedByUserMailboxSetting                        *driveitemcreatedbyusermailboxsetting.DriveItemCreatedByUserMailboxSettingClient
	DriveItemCreatedByUserServiceProvisioningError              *driveitemcreatedbyuserserviceprovisioningerror.DriveItemCreatedByUserServiceProvisioningErrorClient
	DriveItemLastModifiedByUser                                 *driveitemlastmodifiedbyuser.DriveItemLastModifiedByUserClient
	DriveItemLastModifiedByUserMailboxSetting                   *driveitemlastmodifiedbyusermailboxsetting.DriveItemLastModifiedByUserMailboxSettingClient
	DriveItemLastModifiedByUserServiceProvisioningError         *driveitemlastmodifiedbyuserserviceprovisioningerror.DriveItemLastModifiedByUserServiceProvisioningErrorClient
	DriveItemListItem                                           *driveitemlistitem.DriveItemListItemClient
	DriveItemListItemAnalytics                                  *driveitemlistitemanalytics.DriveItemListItemAnalyticsClient
	DriveItemListItemCreatedByUser                              *driveitemlistitemcreatedbyuser.DriveItemListItemCreatedByUserClient
	DriveItemListItemCreatedByUserMailboxSetting                *driveitemlistitemcreatedbyusermailboxsetting.DriveItemListItemCreatedByUserMailboxSettingClient
	DriveItemListItemCreatedByUserServiceProvisioningError      *driveitemlistitemcreatedbyuserserviceprovisioningerror.DriveItemListItemCreatedByUserServiceProvisioningErrorClient
	DriveItemListItemDocumentSetVersion                         *driveitemlistitemdocumentsetversion.DriveItemListItemDocumentSetVersionClient
	DriveItemListItemDocumentSetVersionField                    *driveitemlistitemdocumentsetversionfield.DriveItemListItemDocumentSetVersionFieldClient
	DriveItemListItemDriveItem                                  *driveitemlistitemdriveitem.DriveItemListItemDriveItemClient
	DriveItemListItemDriveItemContent                           *driveitemlistitemdriveitemcontent.DriveItemListItemDriveItemContentClient
	DriveItemListItemField                                      *driveitemlistitemfield.DriveItemListItemFieldClient
	DriveItemListItemLastModifiedByUser                         *driveitemlistitemlastmodifiedbyuser.DriveItemListItemLastModifiedByUserClient
	DriveItemListItemLastModifiedByUserMailboxSetting           *driveitemlistitemlastmodifiedbyusermailboxsetting.DriveItemListItemLastModifiedByUserMailboxSettingClient
	DriveItemListItemLastModifiedByUserServiceProvisioningError *driveitemlistitemlastmodifiedbyuserserviceprovisioningerror.DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveItemListItemVersion                                    *driveitemlistitemversion.DriveItemListItemVersionClient
	DriveItemListItemVersionField                               *driveitemlistitemversionfield.DriveItemListItemVersionFieldClient
	DriveItemPermission                                         *driveitempermission.DriveItemPermissionClient
	DriveItemRetentionLabel                                     *driveitemretentionlabel.DriveItemRetentionLabelClient
	DriveItemSubscription                                       *driveitemsubscription.DriveItemSubscriptionClient
	DriveItemThumbnail                                          *driveitemthumbnail.DriveItemThumbnailClient
	DriveItemVersion                                            *driveitemversion.DriveItemVersionClient
	DriveItemVersionContent                                     *driveitemversioncontent.DriveItemVersionContentClient
	DriveLastModifiedByUser                                     *drivelastmodifiedbyuser.DriveLastModifiedByUserClient
	DriveLastModifiedByUserMailboxSetting                       *drivelastmodifiedbyusermailboxsetting.DriveLastModifiedByUserMailboxSettingClient
	DriveLastModifiedByUserServiceProvisioningError             *drivelastmodifiedbyuserserviceprovisioningerror.DriveLastModifiedByUserServiceProvisioningErrorClient
	DriveList                                                   *drivelist.DriveListClient
	DriveListColumn                                             *drivelistcolumn.DriveListColumnClient
	DriveListColumnSourceColumn                                 *drivelistcolumnsourcecolumn.DriveListColumnSourceColumnClient
	DriveListContentType                                        *drivelistcontenttype.DriveListContentTypeClient
	DriveListContentTypeBase                                    *drivelistcontenttypebase.DriveListContentTypeBaseClient
	DriveListContentTypeBaseType                                *drivelistcontenttypebasetype.DriveListContentTypeBaseTypeClient
	DriveListContentTypeColumn                                  *drivelistcontenttypecolumn.DriveListContentTypeColumnClient
	DriveListContentTypeColumnLink                              *drivelistcontenttypecolumnlink.DriveListContentTypeColumnLinkClient
	DriveListContentTypeColumnPosition                          *drivelistcontenttypecolumnposition.DriveListContentTypeColumnPositionClient
	DriveListContentTypeColumnSourceColumn                      *drivelistcontenttypecolumnsourcecolumn.DriveListContentTypeColumnSourceColumnClient
	DriveListCreatedByUser                                      *drivelistcreatedbyuser.DriveListCreatedByUserClient
	DriveListCreatedByUserMailboxSetting                        *drivelistcreatedbyusermailboxsetting.DriveListCreatedByUserMailboxSettingClient
	DriveListCreatedByUserServiceProvisioningError              *drivelistcreatedbyuserserviceprovisioningerror.DriveListCreatedByUserServiceProvisioningErrorClient
	DriveListDrive                                              *drivelistdrive.DriveListDriveClient
	DriveListItem                                               *drivelistitem.DriveListItemClient
	DriveListItemAnalytics                                      *drivelistitemanalytics.DriveListItemAnalyticsClient
	DriveListItemCreatedByUser                                  *drivelistitemcreatedbyuser.DriveListItemCreatedByUserClient
	DriveListItemCreatedByUserMailboxSetting                    *drivelistitemcreatedbyusermailboxsetting.DriveListItemCreatedByUserMailboxSettingClient
	DriveListItemCreatedByUserServiceProvisioningError          *drivelistitemcreatedbyuserserviceprovisioningerror.DriveListItemCreatedByUserServiceProvisioningErrorClient
	DriveListItemDocumentSetVersion                             *drivelistitemdocumentsetversion.DriveListItemDocumentSetVersionClient
	DriveListItemDocumentSetVersionField                        *drivelistitemdocumentsetversionfield.DriveListItemDocumentSetVersionFieldClient
	DriveListItemDriveItem                                      *drivelistitemdriveitem.DriveListItemDriveItemClient
	DriveListItemDriveItemContent                               *drivelistitemdriveitemcontent.DriveListItemDriveItemContentClient
	DriveListItemField                                          *drivelistitemfield.DriveListItemFieldClient
	DriveListItemLastModifiedByUser                             *drivelistitemlastmodifiedbyuser.DriveListItemLastModifiedByUserClient
	DriveListItemLastModifiedByUserMailboxSetting               *drivelistitemlastmodifiedbyusermailboxsetting.DriveListItemLastModifiedByUserMailboxSettingClient
	DriveListItemLastModifiedByUserServiceProvisioningError     *drivelistitemlastmodifiedbyuserserviceprovisioningerror.DriveListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveListItemVersion                                        *drivelistitemversion.DriveListItemVersionClient
	DriveListItemVersionField                                   *drivelistitemversionfield.DriveListItemVersionFieldClient
	DriveListLastModifiedByUser                                 *drivelistlastmodifiedbyuser.DriveListLastModifiedByUserClient
	DriveListLastModifiedByUserMailboxSetting                   *drivelistlastmodifiedbyusermailboxsetting.DriveListLastModifiedByUserMailboxSettingClient
	DriveListLastModifiedByUserServiceProvisioningError         *drivelistlastmodifiedbyuserserviceprovisioningerror.DriveListLastModifiedByUserServiceProvisioningErrorClient
	DriveListOperation                                          *drivelistoperation.DriveListOperationClient
	DriveListSubscription                                       *drivelistsubscription.DriveListSubscriptionClient
	DriveRoot                                                   *driveroot.DriveRootClient
	DriveRootAnalytics                                          *driverootanalytics.DriveRootAnalyticsClient
	DriveRootAnalyticsAllTime                                   *driverootanalyticsalltime.DriveRootAnalyticsAllTimeClient
	DriveRootAnalyticsItemActivityStat                          *driverootanalyticsitemactivitystat.DriveRootAnalyticsItemActivityStatClient
	DriveRootAnalyticsItemActivityStatActivity                  *driverootanalyticsitemactivitystatactivity.DriveRootAnalyticsItemActivityStatActivityClient
	DriveRootAnalyticsItemActivityStatActivityDriveItem         *driverootanalyticsitemactivitystatactivitydriveitem.DriveRootAnalyticsItemActivityStatActivityDriveItemClient
	DriveRootAnalyticsItemActivityStatActivityDriveItemContent  *driverootanalyticsitemactivitystatactivitydriveitemcontent.DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient
	DriveRootAnalyticsLastSevenDay                              *driverootanalyticslastsevenday.DriveRootAnalyticsLastSevenDayClient
	DriveRootChild                                              *driverootchild.DriveRootChildClient
	DriveRootChildContent                                       *driverootchildcontent.DriveRootChildContentClient
	DriveRootContent                                            *driverootcontent.DriveRootContentClient
	DriveRootCreatedByUser                                      *driverootcreatedbyuser.DriveRootCreatedByUserClient
	DriveRootCreatedByUserMailboxSetting                        *driverootcreatedbyusermailboxsetting.DriveRootCreatedByUserMailboxSettingClient
	DriveRootCreatedByUserServiceProvisioningError              *driverootcreatedbyuserserviceprovisioningerror.DriveRootCreatedByUserServiceProvisioningErrorClient
	DriveRootLastModifiedByUser                                 *driverootlastmodifiedbyuser.DriveRootLastModifiedByUserClient
	DriveRootLastModifiedByUserMailboxSetting                   *driverootlastmodifiedbyusermailboxsetting.DriveRootLastModifiedByUserMailboxSettingClient
	DriveRootLastModifiedByUserServiceProvisioningError         *driverootlastmodifiedbyuserserviceprovisioningerror.DriveRootLastModifiedByUserServiceProvisioningErrorClient
	DriveRootListItem                                           *driverootlistitem.DriveRootListItemClient
	DriveRootListItemAnalytics                                  *driverootlistitemanalytics.DriveRootListItemAnalyticsClient
	DriveRootListItemCreatedByUser                              *driverootlistitemcreatedbyuser.DriveRootListItemCreatedByUserClient
	DriveRootListItemCreatedByUserMailboxSetting                *driverootlistitemcreatedbyusermailboxsetting.DriveRootListItemCreatedByUserMailboxSettingClient
	DriveRootListItemCreatedByUserServiceProvisioningError      *driverootlistitemcreatedbyuserserviceprovisioningerror.DriveRootListItemCreatedByUserServiceProvisioningErrorClient
	DriveRootListItemDocumentSetVersion                         *driverootlistitemdocumentsetversion.DriveRootListItemDocumentSetVersionClient
	DriveRootListItemDocumentSetVersionField                    *driverootlistitemdocumentsetversionfield.DriveRootListItemDocumentSetVersionFieldClient
	DriveRootListItemDriveItem                                  *driverootlistitemdriveitem.DriveRootListItemDriveItemClient
	DriveRootListItemDriveItemContent                           *driverootlistitemdriveitemcontent.DriveRootListItemDriveItemContentClient
	DriveRootListItemField                                      *driverootlistitemfield.DriveRootListItemFieldClient
	DriveRootListItemLastModifiedByUser                         *driverootlistitemlastmodifiedbyuser.DriveRootListItemLastModifiedByUserClient
	DriveRootListItemLastModifiedByUserMailboxSetting           *driverootlistitemlastmodifiedbyusermailboxsetting.DriveRootListItemLastModifiedByUserMailboxSettingClient
	DriveRootListItemLastModifiedByUserServiceProvisioningError *driverootlistitemlastmodifiedbyuserserviceprovisioningerror.DriveRootListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveRootListItemVersion                                    *driverootlistitemversion.DriveRootListItemVersionClient
	DriveRootListItemVersionField                               *driverootlistitemversionfield.DriveRootListItemVersionFieldClient
	DriveRootPermission                                         *driverootpermission.DriveRootPermissionClient
	DriveRootRetentionLabel                                     *driverootretentionlabel.DriveRootRetentionLabelClient
	DriveRootSubscription                                       *driverootsubscription.DriveRootSubscriptionClient
	DriveRootThumbnail                                          *driverootthumbnail.DriveRootThumbnailClient
	DriveRootVersion                                            *driverootversion.DriveRootVersionClient
	DriveRootVersionContent                                     *driverootversioncontent.DriveRootVersionContentClient
	DriveSpecial                                                *drivespecial.DriveSpecialClient
	DriveSpecialContent                                         *drivespecialcontent.DriveSpecialContentClient
	EmployeeExperience                                          *employeeexperience.EmployeeExperienceClient
	EmployeeExperienceLearningCourseActivity                    *employeeexperiencelearningcourseactivity.EmployeeExperienceLearningCourseActivityClient
	Event                                                       *event.EventClient
	EventAttachment                                             *eventattachment.EventAttachmentClient
	EventCalendar                                               *eventcalendar.EventCalendarClient
	EventExtension                                              *eventextension.EventExtensionClient
	EventInstance                                               *eventinstance.EventInstanceClient
	EventInstanceAttachment                                     *eventinstanceattachment.EventInstanceAttachmentClient
	EventInstanceCalendar                                       *eventinstancecalendar.EventInstanceCalendarClient
	EventInstanceExtension                                      *eventinstanceextension.EventInstanceExtensionClient
	Extension                                                   *extension.ExtensionClient
	FollowedSite                                                *followedsite.FollowedSiteClient
	InferenceClassification                                     *inferenceclassification.InferenceClassificationClient
	InferenceClassificationOverride                             *inferenceclassificationoverride.InferenceClassificationOverrideClient
	Insight                                                     *insight.InsightClient
	InsightShared                                               *insightshared.InsightSharedClient
	InsightSharedLastSharedMethod                               *insightsharedlastsharedmethod.InsightSharedLastSharedMethodClient
	InsightSharedResource                                       *insightsharedresource.InsightSharedResourceClient
	InsightTrending                                             *insighttrending.InsightTrendingClient
	InsightTrendingResource                                     *insighttrendingresource.InsightTrendingResourceClient
	InsightUsed                                                 *insightused.InsightUsedClient
	InsightUsedResource                                         *insightusedresource.InsightUsedResourceClient
	JoinedTeam                                                  *joinedteam.JoinedTeamClient
	JoinedTeamAllChannel                                        *joinedteamallchannel.JoinedTeamAllChannelClient
	JoinedTeamChannel                                           *joinedteamchannel.JoinedTeamChannelClient
	JoinedTeamChannelFilesFolder                                *joinedteamchannelfilesfolder.JoinedTeamChannelFilesFolderClient
	JoinedTeamChannelFilesFolderContent                         *joinedteamchannelfilesfoldercontent.JoinedTeamChannelFilesFolderContentClient
	JoinedTeamChannelMember                                     *joinedteamchannelmember.JoinedTeamChannelMemberClient
	JoinedTeamChannelMessage                                    *joinedteamchannelmessage.JoinedTeamChannelMessageClient
	JoinedTeamChannelMessageHostedContent                       *joinedteamchannelmessagehostedcontent.JoinedTeamChannelMessageHostedContentClient
	JoinedTeamChannelMessageReply                               *joinedteamchannelmessagereply.JoinedTeamChannelMessageReplyClient
	JoinedTeamChannelMessageReplyHostedContent                  *joinedteamchannelmessagereplyhostedcontent.JoinedTeamChannelMessageReplyHostedContentClient
	JoinedTeamChannelSharedWithTeam                             *joinedteamchannelsharedwithteam.JoinedTeamChannelSharedWithTeamClient
	JoinedTeamChannelSharedWithTeamAllowedMember                *joinedteamchannelsharedwithteamallowedmember.JoinedTeamChannelSharedWithTeamAllowedMemberClient
	JoinedTeamChannelSharedWithTeamTeam                         *joinedteamchannelsharedwithteamteam.JoinedTeamChannelSharedWithTeamTeamClient
	JoinedTeamChannelTab                                        *joinedteamchanneltab.JoinedTeamChannelTabClient
	JoinedTeamChannelTabTeamsApp                                *joinedteamchanneltabteamsapp.JoinedTeamChannelTabTeamsAppClient
	JoinedTeamGroup                                             *joinedteamgroup.JoinedTeamGroupClient
	JoinedTeamGroupServiceProvisioningError                     *joinedteamgroupserviceprovisioningerror.JoinedTeamGroupServiceProvisioningErrorClient
	JoinedTeamIncomingChannel                                   *joinedteamincomingchannel.JoinedTeamIncomingChannelClient
	JoinedTeamInstalledApp                                      *joinedteaminstalledapp.JoinedTeamInstalledAppClient
	JoinedTeamInstalledAppTeamsApp                              *joinedteaminstalledappteamsapp.JoinedTeamInstalledAppTeamsAppClient
	JoinedTeamInstalledAppTeamsAppDefinition                    *joinedteaminstalledappteamsappdefinition.JoinedTeamInstalledAppTeamsAppDefinitionClient
	JoinedTeamMember                                            *joinedteammember.JoinedTeamMemberClient
	JoinedTeamOperation                                         *joinedteamoperation.JoinedTeamOperationClient
	JoinedTeamPermissionGrant                                   *joinedteampermissiongrant.JoinedTeamPermissionGrantClient
	JoinedTeamPhoto                                             *joinedteamphoto.JoinedTeamPhotoClient
	JoinedTeamPrimaryChannel                                    *joinedteamprimarychannel.JoinedTeamPrimaryChannelClient
	JoinedTeamPrimaryChannelFilesFolder                         *joinedteamprimarychannelfilesfolder.JoinedTeamPrimaryChannelFilesFolderClient
	JoinedTeamPrimaryChannelFilesFolderContent                  *joinedteamprimarychannelfilesfoldercontent.JoinedTeamPrimaryChannelFilesFolderContentClient
	JoinedTeamPrimaryChannelMember                              *joinedteamprimarychannelmember.JoinedTeamPrimaryChannelMemberClient
	JoinedTeamPrimaryChannelMessage                             *joinedteamprimarychannelmessage.JoinedTeamPrimaryChannelMessageClient
	JoinedTeamPrimaryChannelMessageHostedContent                *joinedteamprimarychannelmessagehostedcontent.JoinedTeamPrimaryChannelMessageHostedContentClient
	JoinedTeamPrimaryChannelMessageReply                        *joinedteamprimarychannelmessagereply.JoinedTeamPrimaryChannelMessageReplyClient
	JoinedTeamPrimaryChannelMessageReplyHostedContent           *joinedteamprimarychannelmessagereplyhostedcontent.JoinedTeamPrimaryChannelMessageReplyHostedContentClient
	JoinedTeamPrimaryChannelSharedWithTeam                      *joinedteamprimarychannelsharedwithteam.JoinedTeamPrimaryChannelSharedWithTeamClient
	JoinedTeamPrimaryChannelSharedWithTeamAllowedMember         *joinedteamprimarychannelsharedwithteamallowedmember.JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient
	JoinedTeamPrimaryChannelSharedWithTeamTeam                  *joinedteamprimarychannelsharedwithteamteam.JoinedTeamPrimaryChannelSharedWithTeamTeamClient
	JoinedTeamPrimaryChannelTab                                 *joinedteamprimarychanneltab.JoinedTeamPrimaryChannelTabClient
	JoinedTeamPrimaryChannelTabTeamsApp                         *joinedteamprimarychanneltabteamsapp.JoinedTeamPrimaryChannelTabTeamsAppClient
	JoinedTeamSchedule                                          *joinedteamschedule.JoinedTeamScheduleClient
	JoinedTeamScheduleOfferShiftRequest                         *joinedteamscheduleoffershiftrequest.JoinedTeamScheduleOfferShiftRequestClient
	JoinedTeamScheduleOpenShift                                 *joinedteamscheduleopenshift.JoinedTeamScheduleOpenShiftClient
	JoinedTeamScheduleOpenShiftChangeRequest                    *joinedteamscheduleopenshiftchangerequest.JoinedTeamScheduleOpenShiftChangeRequestClient
	JoinedTeamScheduleSchedulingGroup                           *joinedteamscheduleschedulinggroup.JoinedTeamScheduleSchedulingGroupClient
	JoinedTeamScheduleShift                                     *joinedteamscheduleshift.JoinedTeamScheduleShiftClient
	JoinedTeamScheduleSwapShiftsChangeRequest                   *joinedteamscheduleswapshiftschangerequest.JoinedTeamScheduleSwapShiftsChangeRequestClient
	JoinedTeamScheduleTimeOffReason                             *joinedteamscheduletimeoffreason.JoinedTeamScheduleTimeOffReasonClient
	JoinedTeamScheduleTimeOffRequest                            *joinedteamscheduletimeoffrequest.JoinedTeamScheduleTimeOffRequestClient
	JoinedTeamScheduleTimesOff                                  *joinedteamscheduletimesoff.JoinedTeamScheduleTimesOffClient
	JoinedTeamTag                                               *joinedteamtag.JoinedTeamTagClient
	JoinedTeamTagMember                                         *joinedteamtagmember.JoinedTeamTagMemberClient
	JoinedTeamTemplate                                          *joinedteamtemplate.JoinedTeamTemplateClient
	LicenseDetail                                               *licensedetail.LicenseDetailClient
	MailFolder                                                  *mailfolder.MailFolderClient
	MailFolderChildFolder                                       *mailfolderchildfolder.MailFolderChildFolderClient
	MailFolderChildFolderMessage                                *mailfolderchildfoldermessage.MailFolderChildFolderMessageClient
	MailFolderChildFolderMessageAttachment                      *mailfolderchildfoldermessageattachment.MailFolderChildFolderMessageAttachmentClient
	MailFolderChildFolderMessageExtension                       *mailfolderchildfoldermessageextension.MailFolderChildFolderMessageExtensionClient
	MailFolderChildFolderMessageRule                            *mailfolderchildfoldermessagerule.MailFolderChildFolderMessageRuleClient
	MailFolderMessage                                           *mailfoldermessage.MailFolderMessageClient
	MailFolderMessageAttachment                                 *mailfoldermessageattachment.MailFolderMessageAttachmentClient
	MailFolderMessageExtension                                  *mailfoldermessageextension.MailFolderMessageExtensionClient
	MailFolderMessageRule                                       *mailfoldermessagerule.MailFolderMessageRuleClient
	MailboxSetting                                              *mailboxsetting.MailboxSettingClient
	ManagedAppRegistration                                      *managedappregistration.ManagedAppRegistrationClient
	ManagedDevice                                               *manageddevice.ManagedDeviceClient
	ManagedDeviceDeviceCategory                                 *manageddevicedevicecategory.ManagedDeviceDeviceCategoryClient
	ManagedDeviceDeviceCompliancePolicyState                    *manageddevicedevicecompliancepolicystate.ManagedDeviceDeviceCompliancePolicyStateClient
	ManagedDeviceDeviceConfigurationState                       *manageddevicedeviceconfigurationstate.ManagedDeviceDeviceConfigurationStateClient
	ManagedDeviceLogCollectionRequest                           *manageddevicelogcollectionrequest.ManagedDeviceLogCollectionRequestClient
	ManagedDeviceUser                                           *manageddeviceuser.ManagedDeviceUserClient
	ManagedDeviceWindowsProtectionState                         *manageddevicewindowsprotectionstate.ManagedDeviceWindowsProtectionStateClient
	ManagedDeviceWindowsProtectionStateDetectedMalwareState     *manageddevicewindowsprotectionstatedetectedmalwarestate.ManagedDeviceWindowsProtectionStateDetectedMalwareStateClient
	Manager                                                     *manager.ManagerClient
	MemberOf                                                    *memberof.MemberOfClient
	Message                                                     *message.MessageClient
	MessageAttachment                                           *messageattachment.MessageAttachmentClient
	MessageExtension                                            *messageextension.MessageExtensionClient
	OAuth2PermissionGrant                                       *oauth2permissiongrant.OAuth2PermissionGrantClient
	OnlineMeeting                                               *onlinemeeting.OnlineMeetingClient
	OnlineMeetingAttendanceReport                               *onlinemeetingattendancereport.OnlineMeetingAttendanceReportClient
	OnlineMeetingAttendanceReportAttendanceRecord               *onlinemeetingattendancereportattendancerecord.OnlineMeetingAttendanceReportAttendanceRecordClient
	OnlineMeetingAttendeeReport                                 *onlinemeetingattendeereport.OnlineMeetingAttendeeReportClient
	OnlineMeetingRecording                                      *onlinemeetingrecording.OnlineMeetingRecordingClient
	OnlineMeetingRecordingContent                               *onlinemeetingrecordingcontent.OnlineMeetingRecordingContentClient
	OnlineMeetingTranscript                                     *onlinemeetingtranscript.OnlineMeetingTranscriptClient
	OnlineMeetingTranscriptContent                              *onlinemeetingtranscriptcontent.OnlineMeetingTranscriptContentClient
	OnlineMeetingTranscriptMetadataContent                      *onlinemeetingtranscriptmetadatacontent.OnlineMeetingTranscriptMetadataContentClient
	Outlook                                                     *outlook.OutlookClient
	OutlookMasterCategory                                       *outlookmastercategory.OutlookMasterCategoryClient
	OwnedDevice                                                 *owneddevice.OwnedDeviceClient
	OwnedObject                                                 *ownedobject.OwnedObjectClient
	PermissionGrant                                             *permissiongrant.PermissionGrantClient
	Person                                                      *person.PersonClient
	Photo                                                       *photo.PhotoClient
	Planner                                                     *planner.PlannerClient
	PlannerPlan                                                 *plannerplan.PlannerPlanClient
	PlannerPlanBucket                                           *plannerplanbucket.PlannerPlanBucketClient
	PlannerPlanBucketTask                                       *plannerplanbuckettask.PlannerPlanBucketTaskClient
	PlannerPlanBucketTaskAssignedToTaskBoardFormat              *plannerplanbuckettaskassignedtotaskboardformat.PlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	PlannerPlanBucketTaskBucketTaskBoardFormat                  *plannerplanbuckettaskbuckettaskboardformat.PlannerPlanBucketTaskBucketTaskBoardFormatClient
	PlannerPlanBucketTaskDetail                                 *plannerplanbuckettaskdetail.PlannerPlanBucketTaskDetailClient
	PlannerPlanBucketTaskProgressTaskBoardFormat                *plannerplanbuckettaskprogresstaskboardformat.PlannerPlanBucketTaskProgressTaskBoardFormatClient
	PlannerPlanDetail                                           *plannerplandetail.PlannerPlanDetailClient
	PlannerPlanTask                                             *plannerplantask.PlannerPlanTaskClient
	PlannerPlanTaskAssignedToTaskBoardFormat                    *plannerplantaskassignedtotaskboardformat.PlannerPlanTaskAssignedToTaskBoardFormatClient
	PlannerPlanTaskBucketTaskBoardFormat                        *plannerplantaskbuckettaskboardformat.PlannerPlanTaskBucketTaskBoardFormatClient
	PlannerPlanTaskDetail                                       *plannerplantaskdetail.PlannerPlanTaskDetailClient
	PlannerPlanTaskProgressTaskBoardFormat                      *plannerplantaskprogresstaskboardformat.PlannerPlanTaskProgressTaskBoardFormatClient
	PlannerTask                                                 *plannertask.PlannerTaskClient
	PlannerTaskAssignedToTaskBoardFormat                        *plannertaskassignedtotaskboardformat.PlannerTaskAssignedToTaskBoardFormatClient
	PlannerTaskBucketTaskBoardFormat                            *plannertaskbuckettaskboardformat.PlannerTaskBucketTaskBoardFormatClient
	PlannerTaskDetail                                           *plannertaskdetail.PlannerTaskDetailClient
	PlannerTaskProgressTaskBoardFormat                          *plannertaskprogresstaskboardformat.PlannerTaskProgressTaskBoardFormatClient
	Presence                                                    *presence.PresenceClient
	RegisteredDevice                                            *registereddevice.RegisteredDeviceClient
	ScopedRoleMemberOf                                          *scopedrolememberof.ScopedRoleMemberOfClient
	ServiceProvisioningError                                    *serviceprovisioningerror.ServiceProvisioningErrorClient
	Setting                                                     *setting.SettingClient
	SettingItemInsight                                          *settingiteminsight.SettingItemInsightClient
	SettingShiftPreference                                      *settingshiftpreference.SettingShiftPreferenceClient
	SettingStorage                                              *settingstorage.SettingStorageClient
	SettingStorageQuota                                         *settingstoragequota.SettingStorageQuotaClient
	SettingStorageQuotaService                                  *settingstoragequotaservice.SettingStorageQuotaServiceClient
	SettingWindow                                               *settingwindow.SettingWindowClient
	SettingWindowInstance                                       *settingwindowinstance.SettingWindowInstanceClient
	Sponsor                                                     *sponsor.SponsorClient
	Teamwork                                                    *teamwork.TeamworkClient
	TeamworkAssociatedTeam                                      *teamworkassociatedteam.TeamworkAssociatedTeamClient
	TeamworkAssociatedTeamTeam                                  *teamworkassociatedteamteam.TeamworkAssociatedTeamTeamClient
	TeamworkInstalledApp                                        *teamworkinstalledapp.TeamworkInstalledAppClient
	TeamworkInstalledAppChat                                    *teamworkinstalledappchat.TeamworkInstalledAppChatClient
	TeamworkInstalledAppTeamsApp                                *teamworkinstalledappteamsapp.TeamworkInstalledAppTeamsAppClient
	TeamworkInstalledAppTeamsAppDefinition                      *teamworkinstalledappteamsappdefinition.TeamworkInstalledAppTeamsAppDefinitionClient
	Todo                                                        *todo.TodoClient
	TodoList                                                    *todolist.TodoListClient
	TodoListExtension                                           *todolistextension.TodoListExtensionClient
	TodoListTask                                                *todolisttask.TodoListTaskClient
	TodoListTaskAttachment                                      *todolisttaskattachment.TodoListTaskAttachmentClient
	TodoListTaskAttachmentSession                               *todolisttaskattachmentsession.TodoListTaskAttachmentSessionClient
	TodoListTaskAttachmentSessionContent                        *todolisttaskattachmentsessioncontent.TodoListTaskAttachmentSessionContentClient
	TodoListTaskChecklistItem                                   *todolisttaskchecklistitem.TodoListTaskChecklistItemClient
	TodoListTaskExtension                                       *todolisttaskextension.TodoListTaskExtensionClient
	TodoListTaskLinkedResource                                  *todolisttasklinkedresource.TodoListTaskLinkedResourceClient
	TransitiveMemberOf                                          *transitivememberof.TransitiveMemberOfClient
	User                                                        *user.UserClient
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

	appRoleAssignmentClient, err := approleassignment.NewAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppRoleAssignment client: %+v", err)
	}
	configureFunc(appRoleAssignmentClient.Client)

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

	authenticationPhoneMethodClient, err := authenticationphonemethod.NewAuthenticationPhoneMethodClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AuthenticationPhoneMethod client: %+v", err)
	}
	configureFunc(authenticationPhoneMethodClient.Client)

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

	cloudClipboardClient, err := cloudclipboard.NewCloudClipboardClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudClipboard client: %+v", err)
	}
	configureFunc(cloudClipboardClient.Client)

	cloudClipboardItemClient, err := cloudclipboarditem.NewCloudClipboardItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudClipboardItem client: %+v", err)
	}
	configureFunc(cloudClipboardItemClient.Client)

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

	deviceManagementTroubleshootingEventClient, err := devicemanagementtroubleshootingevent.NewDeviceManagementTroubleshootingEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementTroubleshootingEvent client: %+v", err)
	}
	configureFunc(deviceManagementTroubleshootingEventClient.Client)

	directReportClient, err := directreport.NewDirectReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectReport client: %+v", err)
	}
	configureFunc(directReportClient.Client)

	driveBundleClient, err := drivebundle.NewDriveBundleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveBundle client: %+v", err)
	}
	configureFunc(driveBundleClient.Client)

	driveBundleContentClient, err := drivebundlecontent.NewDriveBundleContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveBundleContent client: %+v", err)
	}
	configureFunc(driveBundleContentClient.Client)

	driveClient, err := drive.NewDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Drive client: %+v", err)
	}
	configureFunc(driveClient.Client)

	driveCreatedByUserClient, err := drivecreatedbyuser.NewDriveCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveCreatedByUser client: %+v", err)
	}
	configureFunc(driveCreatedByUserClient.Client)

	driveCreatedByUserMailboxSettingClient, err := drivecreatedbyusermailboxsetting.NewDriveCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveCreatedByUserMailboxSettingClient.Client)

	driveCreatedByUserServiceProvisioningErrorClient, err := drivecreatedbyuserserviceprovisioningerror.NewDriveCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveCreatedByUserServiceProvisioningErrorClient.Client)

	driveFollowingClient, err := drivefollowing.NewDriveFollowingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveFollowing client: %+v", err)
	}
	configureFunc(driveFollowingClient.Client)

	driveFollowingContentClient, err := drivefollowingcontent.NewDriveFollowingContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveFollowingContent client: %+v", err)
	}
	configureFunc(driveFollowingContentClient.Client)

	driveItemAnalyticsAllTimeClient, err := driveitemanalyticsalltime.NewDriveItemAnalyticsAllTimeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalyticsAllTime client: %+v", err)
	}
	configureFunc(driveItemAnalyticsAllTimeClient.Client)

	driveItemAnalyticsClient, err := driveitemanalytics.NewDriveItemAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalytics client: %+v", err)
	}
	configureFunc(driveItemAnalyticsClient.Client)

	driveItemAnalyticsItemActivityStatActivityClient, err := driveitemanalyticsitemactivitystatactivity.NewDriveItemAnalyticsItemActivityStatActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalyticsItemActivityStatActivity client: %+v", err)
	}
	configureFunc(driveItemAnalyticsItemActivityStatActivityClient.Client)

	driveItemAnalyticsItemActivityStatActivityDriveItemClient, err := driveitemanalyticsitemactivitystatactivitydriveitem.NewDriveItemAnalyticsItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalyticsItemActivityStatActivityDriveItem client: %+v", err)
	}
	configureFunc(driveItemAnalyticsItemActivityStatActivityDriveItemClient.Client)

	driveItemAnalyticsItemActivityStatActivityDriveItemContentClient, err := driveitemanalyticsitemactivitystatactivitydriveitemcontent.NewDriveItemAnalyticsItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalyticsItemActivityStatActivityDriveItemContent client: %+v", err)
	}
	configureFunc(driveItemAnalyticsItemActivityStatActivityDriveItemContentClient.Client)

	driveItemAnalyticsItemActivityStatClient, err := driveitemanalyticsitemactivitystat.NewDriveItemAnalyticsItemActivityStatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalyticsItemActivityStat client: %+v", err)
	}
	configureFunc(driveItemAnalyticsItemActivityStatClient.Client)

	driveItemAnalyticsLastSevenDayClient, err := driveitemanalyticslastsevenday.NewDriveItemAnalyticsLastSevenDayClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalyticsLastSevenDay client: %+v", err)
	}
	configureFunc(driveItemAnalyticsLastSevenDayClient.Client)

	driveItemChildClient, err := driveitemchild.NewDriveItemChildClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemChild client: %+v", err)
	}
	configureFunc(driveItemChildClient.Client)

	driveItemChildContentClient, err := driveitemchildcontent.NewDriveItemChildContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemChildContent client: %+v", err)
	}
	configureFunc(driveItemChildContentClient.Client)

	driveItemClient, err := driveitem.NewDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItem client: %+v", err)
	}
	configureFunc(driveItemClient.Client)

	driveItemContentClient, err := driveitemcontent.NewDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemContent client: %+v", err)
	}
	configureFunc(driveItemContentClient.Client)

	driveItemCreatedByUserClient, err := driveitemcreatedbyuser.NewDriveItemCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemCreatedByUser client: %+v", err)
	}
	configureFunc(driveItemCreatedByUserClient.Client)

	driveItemCreatedByUserMailboxSettingClient, err := driveitemcreatedbyusermailboxsetting.NewDriveItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveItemCreatedByUserMailboxSettingClient.Client)

	driveItemCreatedByUserServiceProvisioningErrorClient, err := driveitemcreatedbyuserserviceprovisioningerror.NewDriveItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveItemCreatedByUserServiceProvisioningErrorClient.Client)

	driveItemLastModifiedByUserClient, err := driveitemlastmodifiedbyuser.NewDriveItemLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemLastModifiedByUser client: %+v", err)
	}
	configureFunc(driveItemLastModifiedByUserClient.Client)

	driveItemLastModifiedByUserMailboxSettingClient, err := driveitemlastmodifiedbyusermailboxsetting.NewDriveItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveItemLastModifiedByUserMailboxSettingClient.Client)

	driveItemLastModifiedByUserServiceProvisioningErrorClient, err := driveitemlastmodifiedbyuserserviceprovisioningerror.NewDriveItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveItemLastModifiedByUserServiceProvisioningErrorClient.Client)

	driveItemListItemAnalyticsClient, err := driveitemlistitemanalytics.NewDriveItemListItemAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemAnalytics client: %+v", err)
	}
	configureFunc(driveItemListItemAnalyticsClient.Client)

	driveItemListItemClient, err := driveitemlistitem.NewDriveItemListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItem client: %+v", err)
	}
	configureFunc(driveItemListItemClient.Client)

	driveItemListItemCreatedByUserClient, err := driveitemlistitemcreatedbyuser.NewDriveItemListItemCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemCreatedByUser client: %+v", err)
	}
	configureFunc(driveItemListItemCreatedByUserClient.Client)

	driveItemListItemCreatedByUserMailboxSettingClient, err := driveitemlistitemcreatedbyusermailboxsetting.NewDriveItemListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveItemListItemCreatedByUserMailboxSettingClient.Client)

	driveItemListItemCreatedByUserServiceProvisioningErrorClient, err := driveitemlistitemcreatedbyuserserviceprovisioningerror.NewDriveItemListItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveItemListItemCreatedByUserServiceProvisioningErrorClient.Client)

	driveItemListItemDocumentSetVersionClient, err := driveitemlistitemdocumentsetversion.NewDriveItemListItemDocumentSetVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemDocumentSetVersion client: %+v", err)
	}
	configureFunc(driveItemListItemDocumentSetVersionClient.Client)

	driveItemListItemDocumentSetVersionFieldClient, err := driveitemlistitemdocumentsetversionfield.NewDriveItemListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemDocumentSetVersionField client: %+v", err)
	}
	configureFunc(driveItemListItemDocumentSetVersionFieldClient.Client)

	driveItemListItemDriveItemClient, err := driveitemlistitemdriveitem.NewDriveItemListItemDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemDriveItem client: %+v", err)
	}
	configureFunc(driveItemListItemDriveItemClient.Client)

	driveItemListItemDriveItemContentClient, err := driveitemlistitemdriveitemcontent.NewDriveItemListItemDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemDriveItemContent client: %+v", err)
	}
	configureFunc(driveItemListItemDriveItemContentClient.Client)

	driveItemListItemFieldClient, err := driveitemlistitemfield.NewDriveItemListItemFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemField client: %+v", err)
	}
	configureFunc(driveItemListItemFieldClient.Client)

	driveItemListItemLastModifiedByUserClient, err := driveitemlistitemlastmodifiedbyuser.NewDriveItemListItemLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemLastModifiedByUser client: %+v", err)
	}
	configureFunc(driveItemListItemLastModifiedByUserClient.Client)

	driveItemListItemLastModifiedByUserMailboxSettingClient, err := driveitemlistitemlastmodifiedbyusermailboxsetting.NewDriveItemListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveItemListItemLastModifiedByUserMailboxSettingClient.Client)

	driveItemListItemLastModifiedByUserServiceProvisioningErrorClient, err := driveitemlistitemlastmodifiedbyuserserviceprovisioningerror.NewDriveItemListItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveItemListItemLastModifiedByUserServiceProvisioningErrorClient.Client)

	driveItemListItemVersionClient, err := driveitemlistitemversion.NewDriveItemListItemVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemVersion client: %+v", err)
	}
	configureFunc(driveItemListItemVersionClient.Client)

	driveItemListItemVersionFieldClient, err := driveitemlistitemversionfield.NewDriveItemListItemVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemVersionField client: %+v", err)
	}
	configureFunc(driveItemListItemVersionFieldClient.Client)

	driveItemPermissionClient, err := driveitempermission.NewDriveItemPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemPermission client: %+v", err)
	}
	configureFunc(driveItemPermissionClient.Client)

	driveItemRetentionLabelClient, err := driveitemretentionlabel.NewDriveItemRetentionLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemRetentionLabel client: %+v", err)
	}
	configureFunc(driveItemRetentionLabelClient.Client)

	driveItemSubscriptionClient, err := driveitemsubscription.NewDriveItemSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemSubscription client: %+v", err)
	}
	configureFunc(driveItemSubscriptionClient.Client)

	driveItemThumbnailClient, err := driveitemthumbnail.NewDriveItemThumbnailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemThumbnail client: %+v", err)
	}
	configureFunc(driveItemThumbnailClient.Client)

	driveItemVersionClient, err := driveitemversion.NewDriveItemVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemVersion client: %+v", err)
	}
	configureFunc(driveItemVersionClient.Client)

	driveItemVersionContentClient, err := driveitemversioncontent.NewDriveItemVersionContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemVersionContent client: %+v", err)
	}
	configureFunc(driveItemVersionContentClient.Client)

	driveLastModifiedByUserClient, err := drivelastmodifiedbyuser.NewDriveLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveLastModifiedByUser client: %+v", err)
	}
	configureFunc(driveLastModifiedByUserClient.Client)

	driveLastModifiedByUserMailboxSettingClient, err := drivelastmodifiedbyusermailboxsetting.NewDriveLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveLastModifiedByUserMailboxSettingClient.Client)

	driveLastModifiedByUserServiceProvisioningErrorClient, err := drivelastmodifiedbyuserserviceprovisioningerror.NewDriveLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveLastModifiedByUserServiceProvisioningErrorClient.Client)

	driveListClient, err := drivelist.NewDriveListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveList client: %+v", err)
	}
	configureFunc(driveListClient.Client)

	driveListColumnClient, err := drivelistcolumn.NewDriveListColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListColumn client: %+v", err)
	}
	configureFunc(driveListColumnClient.Client)

	driveListColumnSourceColumnClient, err := drivelistcolumnsourcecolumn.NewDriveListColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListColumnSourceColumn client: %+v", err)
	}
	configureFunc(driveListColumnSourceColumnClient.Client)

	driveListContentTypeBaseClient, err := drivelistcontenttypebase.NewDriveListContentTypeBaseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListContentTypeBase client: %+v", err)
	}
	configureFunc(driveListContentTypeBaseClient.Client)

	driveListContentTypeBaseTypeClient, err := drivelistcontenttypebasetype.NewDriveListContentTypeBaseTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListContentTypeBaseType client: %+v", err)
	}
	configureFunc(driveListContentTypeBaseTypeClient.Client)

	driveListContentTypeClient, err := drivelistcontenttype.NewDriveListContentTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListContentType client: %+v", err)
	}
	configureFunc(driveListContentTypeClient.Client)

	driveListContentTypeColumnClient, err := drivelistcontenttypecolumn.NewDriveListContentTypeColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListContentTypeColumn client: %+v", err)
	}
	configureFunc(driveListContentTypeColumnClient.Client)

	driveListContentTypeColumnLinkClient, err := drivelistcontenttypecolumnlink.NewDriveListContentTypeColumnLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListContentTypeColumnLink client: %+v", err)
	}
	configureFunc(driveListContentTypeColumnLinkClient.Client)

	driveListContentTypeColumnPositionClient, err := drivelistcontenttypecolumnposition.NewDriveListContentTypeColumnPositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListContentTypeColumnPosition client: %+v", err)
	}
	configureFunc(driveListContentTypeColumnPositionClient.Client)

	driveListContentTypeColumnSourceColumnClient, err := drivelistcontenttypecolumnsourcecolumn.NewDriveListContentTypeColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListContentTypeColumnSourceColumn client: %+v", err)
	}
	configureFunc(driveListContentTypeColumnSourceColumnClient.Client)

	driveListCreatedByUserClient, err := drivelistcreatedbyuser.NewDriveListCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListCreatedByUser client: %+v", err)
	}
	configureFunc(driveListCreatedByUserClient.Client)

	driveListCreatedByUserMailboxSettingClient, err := drivelistcreatedbyusermailboxsetting.NewDriveListCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveListCreatedByUserMailboxSettingClient.Client)

	driveListCreatedByUserServiceProvisioningErrorClient, err := drivelistcreatedbyuserserviceprovisioningerror.NewDriveListCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveListCreatedByUserServiceProvisioningErrorClient.Client)

	driveListDriveClient, err := drivelistdrive.NewDriveListDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListDrive client: %+v", err)
	}
	configureFunc(driveListDriveClient.Client)

	driveListItemAnalyticsClient, err := drivelistitemanalytics.NewDriveListItemAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemAnalytics client: %+v", err)
	}
	configureFunc(driveListItemAnalyticsClient.Client)

	driveListItemClient, err := drivelistitem.NewDriveListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItem client: %+v", err)
	}
	configureFunc(driveListItemClient.Client)

	driveListItemCreatedByUserClient, err := drivelistitemcreatedbyuser.NewDriveListItemCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemCreatedByUser client: %+v", err)
	}
	configureFunc(driveListItemCreatedByUserClient.Client)

	driveListItemCreatedByUserMailboxSettingClient, err := drivelistitemcreatedbyusermailboxsetting.NewDriveListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveListItemCreatedByUserMailboxSettingClient.Client)

	driveListItemCreatedByUserServiceProvisioningErrorClient, err := drivelistitemcreatedbyuserserviceprovisioningerror.NewDriveListItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveListItemCreatedByUserServiceProvisioningErrorClient.Client)

	driveListItemDocumentSetVersionClient, err := drivelistitemdocumentsetversion.NewDriveListItemDocumentSetVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemDocumentSetVersion client: %+v", err)
	}
	configureFunc(driveListItemDocumentSetVersionClient.Client)

	driveListItemDocumentSetVersionFieldClient, err := drivelistitemdocumentsetversionfield.NewDriveListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemDocumentSetVersionField client: %+v", err)
	}
	configureFunc(driveListItemDocumentSetVersionFieldClient.Client)

	driveListItemDriveItemClient, err := drivelistitemdriveitem.NewDriveListItemDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemDriveItem client: %+v", err)
	}
	configureFunc(driveListItemDriveItemClient.Client)

	driveListItemDriveItemContentClient, err := drivelistitemdriveitemcontent.NewDriveListItemDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemDriveItemContent client: %+v", err)
	}
	configureFunc(driveListItemDriveItemContentClient.Client)

	driveListItemFieldClient, err := drivelistitemfield.NewDriveListItemFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemField client: %+v", err)
	}
	configureFunc(driveListItemFieldClient.Client)

	driveListItemLastModifiedByUserClient, err := drivelistitemlastmodifiedbyuser.NewDriveListItemLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemLastModifiedByUser client: %+v", err)
	}
	configureFunc(driveListItemLastModifiedByUserClient.Client)

	driveListItemLastModifiedByUserMailboxSettingClient, err := drivelistitemlastmodifiedbyusermailboxsetting.NewDriveListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveListItemLastModifiedByUserMailboxSettingClient.Client)

	driveListItemLastModifiedByUserServiceProvisioningErrorClient, err := drivelistitemlastmodifiedbyuserserviceprovisioningerror.NewDriveListItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveListItemLastModifiedByUserServiceProvisioningErrorClient.Client)

	driveListItemVersionClient, err := drivelistitemversion.NewDriveListItemVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemVersion client: %+v", err)
	}
	configureFunc(driveListItemVersionClient.Client)

	driveListItemVersionFieldClient, err := drivelistitemversionfield.NewDriveListItemVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemVersionField client: %+v", err)
	}
	configureFunc(driveListItemVersionFieldClient.Client)

	driveListLastModifiedByUserClient, err := drivelistlastmodifiedbyuser.NewDriveListLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListLastModifiedByUser client: %+v", err)
	}
	configureFunc(driveListLastModifiedByUserClient.Client)

	driveListLastModifiedByUserMailboxSettingClient, err := drivelistlastmodifiedbyusermailboxsetting.NewDriveListLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveListLastModifiedByUserMailboxSettingClient.Client)

	driveListLastModifiedByUserServiceProvisioningErrorClient, err := drivelistlastmodifiedbyuserserviceprovisioningerror.NewDriveListLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveListLastModifiedByUserServiceProvisioningErrorClient.Client)

	driveListOperationClient, err := drivelistoperation.NewDriveListOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListOperation client: %+v", err)
	}
	configureFunc(driveListOperationClient.Client)

	driveListSubscriptionClient, err := drivelistsubscription.NewDriveListSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListSubscription client: %+v", err)
	}
	configureFunc(driveListSubscriptionClient.Client)

	driveRootAnalyticsAllTimeClient, err := driverootanalyticsalltime.NewDriveRootAnalyticsAllTimeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalyticsAllTime client: %+v", err)
	}
	configureFunc(driveRootAnalyticsAllTimeClient.Client)

	driveRootAnalyticsClient, err := driverootanalytics.NewDriveRootAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalytics client: %+v", err)
	}
	configureFunc(driveRootAnalyticsClient.Client)

	driveRootAnalyticsItemActivityStatActivityClient, err := driverootanalyticsitemactivitystatactivity.NewDriveRootAnalyticsItemActivityStatActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalyticsItemActivityStatActivity client: %+v", err)
	}
	configureFunc(driveRootAnalyticsItemActivityStatActivityClient.Client)

	driveRootAnalyticsItemActivityStatActivityDriveItemClient, err := driverootanalyticsitemactivitystatactivitydriveitem.NewDriveRootAnalyticsItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalyticsItemActivityStatActivityDriveItem client: %+v", err)
	}
	configureFunc(driveRootAnalyticsItemActivityStatActivityDriveItemClient.Client)

	driveRootAnalyticsItemActivityStatActivityDriveItemContentClient, err := driverootanalyticsitemactivitystatactivitydriveitemcontent.NewDriveRootAnalyticsItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalyticsItemActivityStatActivityDriveItemContent client: %+v", err)
	}
	configureFunc(driveRootAnalyticsItemActivityStatActivityDriveItemContentClient.Client)

	driveRootAnalyticsItemActivityStatClient, err := driverootanalyticsitemactivitystat.NewDriveRootAnalyticsItemActivityStatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalyticsItemActivityStat client: %+v", err)
	}
	configureFunc(driveRootAnalyticsItemActivityStatClient.Client)

	driveRootAnalyticsLastSevenDayClient, err := driverootanalyticslastsevenday.NewDriveRootAnalyticsLastSevenDayClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalyticsLastSevenDay client: %+v", err)
	}
	configureFunc(driveRootAnalyticsLastSevenDayClient.Client)

	driveRootChildClient, err := driverootchild.NewDriveRootChildClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootChild client: %+v", err)
	}
	configureFunc(driveRootChildClient.Client)

	driveRootChildContentClient, err := driverootchildcontent.NewDriveRootChildContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootChildContent client: %+v", err)
	}
	configureFunc(driveRootChildContentClient.Client)

	driveRootClient, err := driveroot.NewDriveRootClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRoot client: %+v", err)
	}
	configureFunc(driveRootClient.Client)

	driveRootContentClient, err := driverootcontent.NewDriveRootContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootContent client: %+v", err)
	}
	configureFunc(driveRootContentClient.Client)

	driveRootCreatedByUserClient, err := driverootcreatedbyuser.NewDriveRootCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootCreatedByUser client: %+v", err)
	}
	configureFunc(driveRootCreatedByUserClient.Client)

	driveRootCreatedByUserMailboxSettingClient, err := driverootcreatedbyusermailboxsetting.NewDriveRootCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveRootCreatedByUserMailboxSettingClient.Client)

	driveRootCreatedByUserServiceProvisioningErrorClient, err := driverootcreatedbyuserserviceprovisioningerror.NewDriveRootCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveRootCreatedByUserServiceProvisioningErrorClient.Client)

	driveRootLastModifiedByUserClient, err := driverootlastmodifiedbyuser.NewDriveRootLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootLastModifiedByUser client: %+v", err)
	}
	configureFunc(driveRootLastModifiedByUserClient.Client)

	driveRootLastModifiedByUserMailboxSettingClient, err := driverootlastmodifiedbyusermailboxsetting.NewDriveRootLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveRootLastModifiedByUserMailboxSettingClient.Client)

	driveRootLastModifiedByUserServiceProvisioningErrorClient, err := driverootlastmodifiedbyuserserviceprovisioningerror.NewDriveRootLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveRootLastModifiedByUserServiceProvisioningErrorClient.Client)

	driveRootListItemAnalyticsClient, err := driverootlistitemanalytics.NewDriveRootListItemAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemAnalytics client: %+v", err)
	}
	configureFunc(driveRootListItemAnalyticsClient.Client)

	driveRootListItemClient, err := driverootlistitem.NewDriveRootListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItem client: %+v", err)
	}
	configureFunc(driveRootListItemClient.Client)

	driveRootListItemCreatedByUserClient, err := driverootlistitemcreatedbyuser.NewDriveRootListItemCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemCreatedByUser client: %+v", err)
	}
	configureFunc(driveRootListItemCreatedByUserClient.Client)

	driveRootListItemCreatedByUserMailboxSettingClient, err := driverootlistitemcreatedbyusermailboxsetting.NewDriveRootListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveRootListItemCreatedByUserMailboxSettingClient.Client)

	driveRootListItemCreatedByUserServiceProvisioningErrorClient, err := driverootlistitemcreatedbyuserserviceprovisioningerror.NewDriveRootListItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveRootListItemCreatedByUserServiceProvisioningErrorClient.Client)

	driveRootListItemDocumentSetVersionClient, err := driverootlistitemdocumentsetversion.NewDriveRootListItemDocumentSetVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemDocumentSetVersion client: %+v", err)
	}
	configureFunc(driveRootListItemDocumentSetVersionClient.Client)

	driveRootListItemDocumentSetVersionFieldClient, err := driverootlistitemdocumentsetversionfield.NewDriveRootListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemDocumentSetVersionField client: %+v", err)
	}
	configureFunc(driveRootListItemDocumentSetVersionFieldClient.Client)

	driveRootListItemDriveItemClient, err := driverootlistitemdriveitem.NewDriveRootListItemDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemDriveItem client: %+v", err)
	}
	configureFunc(driveRootListItemDriveItemClient.Client)

	driveRootListItemDriveItemContentClient, err := driverootlistitemdriveitemcontent.NewDriveRootListItemDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemDriveItemContent client: %+v", err)
	}
	configureFunc(driveRootListItemDriveItemContentClient.Client)

	driveRootListItemFieldClient, err := driverootlistitemfield.NewDriveRootListItemFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemField client: %+v", err)
	}
	configureFunc(driveRootListItemFieldClient.Client)

	driveRootListItemLastModifiedByUserClient, err := driverootlistitemlastmodifiedbyuser.NewDriveRootListItemLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemLastModifiedByUser client: %+v", err)
	}
	configureFunc(driveRootListItemLastModifiedByUserClient.Client)

	driveRootListItemLastModifiedByUserMailboxSettingClient, err := driverootlistitemlastmodifiedbyusermailboxsetting.NewDriveRootListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(driveRootListItemLastModifiedByUserMailboxSettingClient.Client)

	driveRootListItemLastModifiedByUserServiceProvisioningErrorClient, err := driverootlistitemlastmodifiedbyuserserviceprovisioningerror.NewDriveRootListItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(driveRootListItemLastModifiedByUserServiceProvisioningErrorClient.Client)

	driveRootListItemVersionClient, err := driverootlistitemversion.NewDriveRootListItemVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemVersion client: %+v", err)
	}
	configureFunc(driveRootListItemVersionClient.Client)

	driveRootListItemVersionFieldClient, err := driverootlistitemversionfield.NewDriveRootListItemVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemVersionField client: %+v", err)
	}
	configureFunc(driveRootListItemVersionFieldClient.Client)

	driveRootPermissionClient, err := driverootpermission.NewDriveRootPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootPermission client: %+v", err)
	}
	configureFunc(driveRootPermissionClient.Client)

	driveRootRetentionLabelClient, err := driverootretentionlabel.NewDriveRootRetentionLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootRetentionLabel client: %+v", err)
	}
	configureFunc(driveRootRetentionLabelClient.Client)

	driveRootSubscriptionClient, err := driverootsubscription.NewDriveRootSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootSubscription client: %+v", err)
	}
	configureFunc(driveRootSubscriptionClient.Client)

	driveRootThumbnailClient, err := driverootthumbnail.NewDriveRootThumbnailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootThumbnail client: %+v", err)
	}
	configureFunc(driveRootThumbnailClient.Client)

	driveRootVersionClient, err := driverootversion.NewDriveRootVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootVersion client: %+v", err)
	}
	configureFunc(driveRootVersionClient.Client)

	driveRootVersionContentClient, err := driverootversioncontent.NewDriveRootVersionContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootVersionContent client: %+v", err)
	}
	configureFunc(driveRootVersionContentClient.Client)

	driveSpecialClient, err := drivespecial.NewDriveSpecialClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveSpecial client: %+v", err)
	}
	configureFunc(driveSpecialClient.Client)

	driveSpecialContentClient, err := drivespecialcontent.NewDriveSpecialContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveSpecialContent client: %+v", err)
	}
	configureFunc(driveSpecialContentClient.Client)

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

	joinedTeamAllChannelClient, err := joinedteamallchannel.NewJoinedTeamAllChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamAllChannel client: %+v", err)
	}
	configureFunc(joinedTeamAllChannelClient.Client)

	joinedTeamChannelClient, err := joinedteamchannel.NewJoinedTeamChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannel client: %+v", err)
	}
	configureFunc(joinedTeamChannelClient.Client)

	joinedTeamChannelFilesFolderClient, err := joinedteamchannelfilesfolder.NewJoinedTeamChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelFilesFolder client: %+v", err)
	}
	configureFunc(joinedTeamChannelFilesFolderClient.Client)

	joinedTeamChannelFilesFolderContentClient, err := joinedteamchannelfilesfoldercontent.NewJoinedTeamChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(joinedTeamChannelFilesFolderContentClient.Client)

	joinedTeamChannelMemberClient, err := joinedteamchannelmember.NewJoinedTeamChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelMember client: %+v", err)
	}
	configureFunc(joinedTeamChannelMemberClient.Client)

	joinedTeamChannelMessageClient, err := joinedteamchannelmessage.NewJoinedTeamChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelMessage client: %+v", err)
	}
	configureFunc(joinedTeamChannelMessageClient.Client)

	joinedTeamChannelMessageHostedContentClient, err := joinedteamchannelmessagehostedcontent.NewJoinedTeamChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(joinedTeamChannelMessageHostedContentClient.Client)

	joinedTeamChannelMessageReplyClient, err := joinedteamchannelmessagereply.NewJoinedTeamChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelMessageReply client: %+v", err)
	}
	configureFunc(joinedTeamChannelMessageReplyClient.Client)

	joinedTeamChannelMessageReplyHostedContentClient, err := joinedteamchannelmessagereplyhostedcontent.NewJoinedTeamChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(joinedTeamChannelMessageReplyHostedContentClient.Client)

	joinedTeamChannelSharedWithTeamAllowedMemberClient, err := joinedteamchannelsharedwithteamallowedmember.NewJoinedTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(joinedTeamChannelSharedWithTeamAllowedMemberClient.Client)

	joinedTeamChannelSharedWithTeamClient, err := joinedteamchannelsharedwithteam.NewJoinedTeamChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(joinedTeamChannelSharedWithTeamClient.Client)

	joinedTeamChannelSharedWithTeamTeamClient, err := joinedteamchannelsharedwithteamteam.NewJoinedTeamChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(joinedTeamChannelSharedWithTeamTeamClient.Client)

	joinedTeamChannelTabClient, err := joinedteamchanneltab.NewJoinedTeamChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelTab client: %+v", err)
	}
	configureFunc(joinedTeamChannelTabClient.Client)

	joinedTeamChannelTabTeamsAppClient, err := joinedteamchanneltabteamsapp.NewJoinedTeamChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(joinedTeamChannelTabTeamsAppClient.Client)

	joinedTeamClient, err := joinedteam.NewJoinedTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeam client: %+v", err)
	}
	configureFunc(joinedTeamClient.Client)

	joinedTeamGroupClient, err := joinedteamgroup.NewJoinedTeamGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamGroup client: %+v", err)
	}
	configureFunc(joinedTeamGroupClient.Client)

	joinedTeamGroupServiceProvisioningErrorClient, err := joinedteamgroupserviceprovisioningerror.NewJoinedTeamGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(joinedTeamGroupServiceProvisioningErrorClient.Client)

	joinedTeamIncomingChannelClient, err := joinedteamincomingchannel.NewJoinedTeamIncomingChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamIncomingChannel client: %+v", err)
	}
	configureFunc(joinedTeamIncomingChannelClient.Client)

	joinedTeamInstalledAppClient, err := joinedteaminstalledapp.NewJoinedTeamInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamInstalledApp client: %+v", err)
	}
	configureFunc(joinedTeamInstalledAppClient.Client)

	joinedTeamInstalledAppTeamsAppClient, err := joinedteaminstalledappteamsapp.NewJoinedTeamInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(joinedTeamInstalledAppTeamsAppClient.Client)

	joinedTeamInstalledAppTeamsAppDefinitionClient, err := joinedteaminstalledappteamsappdefinition.NewJoinedTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(joinedTeamInstalledAppTeamsAppDefinitionClient.Client)

	joinedTeamMemberClient, err := joinedteammember.NewJoinedTeamMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamMember client: %+v", err)
	}
	configureFunc(joinedTeamMemberClient.Client)

	joinedTeamOperationClient, err := joinedteamoperation.NewJoinedTeamOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamOperation client: %+v", err)
	}
	configureFunc(joinedTeamOperationClient.Client)

	joinedTeamPermissionGrantClient, err := joinedteampermissiongrant.NewJoinedTeamPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPermissionGrant client: %+v", err)
	}
	configureFunc(joinedTeamPermissionGrantClient.Client)

	joinedTeamPhotoClient, err := joinedteamphoto.NewJoinedTeamPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPhoto client: %+v", err)
	}
	configureFunc(joinedTeamPhotoClient.Client)

	joinedTeamPrimaryChannelClient, err := joinedteamprimarychannel.NewJoinedTeamPrimaryChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannel client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelClient.Client)

	joinedTeamPrimaryChannelFilesFolderClient, err := joinedteamprimarychannelfilesfolder.NewJoinedTeamPrimaryChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelFilesFolder client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelFilesFolderClient.Client)

	joinedTeamPrimaryChannelFilesFolderContentClient, err := joinedteamprimarychannelfilesfoldercontent.NewJoinedTeamPrimaryChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelFilesFolderContentClient.Client)

	joinedTeamPrimaryChannelMemberClient, err := joinedteamprimarychannelmember.NewJoinedTeamPrimaryChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelMember client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelMemberClient.Client)

	joinedTeamPrimaryChannelMessageClient, err := joinedteamprimarychannelmessage.NewJoinedTeamPrimaryChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelMessage client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelMessageClient.Client)

	joinedTeamPrimaryChannelMessageHostedContentClient, err := joinedteamprimarychannelmessagehostedcontent.NewJoinedTeamPrimaryChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelMessageHostedContentClient.Client)

	joinedTeamPrimaryChannelMessageReplyClient, err := joinedteamprimarychannelmessagereply.NewJoinedTeamPrimaryChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelMessageReply client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelMessageReplyClient.Client)

	joinedTeamPrimaryChannelMessageReplyHostedContentClient, err := joinedteamprimarychannelmessagereplyhostedcontent.NewJoinedTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelMessageReplyHostedContentClient.Client)

	joinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient, err := joinedteamprimarychannelsharedwithteamallowedmember.NewJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient.Client)

	joinedTeamPrimaryChannelSharedWithTeamClient, err := joinedteamprimarychannelsharedwithteam.NewJoinedTeamPrimaryChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelSharedWithTeamClient.Client)

	joinedTeamPrimaryChannelSharedWithTeamTeamClient, err := joinedteamprimarychannelsharedwithteamteam.NewJoinedTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelSharedWithTeamTeamClient.Client)

	joinedTeamPrimaryChannelTabClient, err := joinedteamprimarychanneltab.NewJoinedTeamPrimaryChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelTab client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelTabClient.Client)

	joinedTeamPrimaryChannelTabTeamsAppClient, err := joinedteamprimarychanneltabteamsapp.NewJoinedTeamPrimaryChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamPrimaryChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(joinedTeamPrimaryChannelTabTeamsAppClient.Client)

	joinedTeamScheduleClient, err := joinedteamschedule.NewJoinedTeamScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamSchedule client: %+v", err)
	}
	configureFunc(joinedTeamScheduleClient.Client)

	joinedTeamScheduleOfferShiftRequestClient, err := joinedteamscheduleoffershiftrequest.NewJoinedTeamScheduleOfferShiftRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleOfferShiftRequest client: %+v", err)
	}
	configureFunc(joinedTeamScheduleOfferShiftRequestClient.Client)

	joinedTeamScheduleOpenShiftChangeRequestClient, err := joinedteamscheduleopenshiftchangerequest.NewJoinedTeamScheduleOpenShiftChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleOpenShiftChangeRequest client: %+v", err)
	}
	configureFunc(joinedTeamScheduleOpenShiftChangeRequestClient.Client)

	joinedTeamScheduleOpenShiftClient, err := joinedteamscheduleopenshift.NewJoinedTeamScheduleOpenShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleOpenShift client: %+v", err)
	}
	configureFunc(joinedTeamScheduleOpenShiftClient.Client)

	joinedTeamScheduleSchedulingGroupClient, err := joinedteamscheduleschedulinggroup.NewJoinedTeamScheduleSchedulingGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleSchedulingGroup client: %+v", err)
	}
	configureFunc(joinedTeamScheduleSchedulingGroupClient.Client)

	joinedTeamScheduleShiftClient, err := joinedteamscheduleshift.NewJoinedTeamScheduleShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleShift client: %+v", err)
	}
	configureFunc(joinedTeamScheduleShiftClient.Client)

	joinedTeamScheduleSwapShiftsChangeRequestClient, err := joinedteamscheduleswapshiftschangerequest.NewJoinedTeamScheduleSwapShiftsChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleSwapShiftsChangeRequest client: %+v", err)
	}
	configureFunc(joinedTeamScheduleSwapShiftsChangeRequestClient.Client)

	joinedTeamScheduleTimeOffReasonClient, err := joinedteamscheduletimeoffreason.NewJoinedTeamScheduleTimeOffReasonClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleTimeOffReason client: %+v", err)
	}
	configureFunc(joinedTeamScheduleTimeOffReasonClient.Client)

	joinedTeamScheduleTimeOffRequestClient, err := joinedteamscheduletimeoffrequest.NewJoinedTeamScheduleTimeOffRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleTimeOffRequest client: %+v", err)
	}
	configureFunc(joinedTeamScheduleTimeOffRequestClient.Client)

	joinedTeamScheduleTimesOffClient, err := joinedteamscheduletimesoff.NewJoinedTeamScheduleTimesOffClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamScheduleTimesOff client: %+v", err)
	}
	configureFunc(joinedTeamScheduleTimesOffClient.Client)

	joinedTeamTagClient, err := joinedteamtag.NewJoinedTeamTagClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamTag client: %+v", err)
	}
	configureFunc(joinedTeamTagClient.Client)

	joinedTeamTagMemberClient, err := joinedteamtagmember.NewJoinedTeamTagMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamTagMember client: %+v", err)
	}
	configureFunc(joinedTeamTagMemberClient.Client)

	joinedTeamTemplateClient, err := joinedteamtemplate.NewJoinedTeamTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JoinedTeamTemplate client: %+v", err)
	}
	configureFunc(joinedTeamTemplateClient.Client)

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

	mailFolderChildFolderMessageRuleClient, err := mailfolderchildfoldermessagerule.NewMailFolderChildFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderChildFolderMessageRule client: %+v", err)
	}
	configureFunc(mailFolderChildFolderMessageRuleClient.Client)

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

	mailFolderMessageRuleClient, err := mailfoldermessagerule.NewMailFolderMessageRuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MailFolderMessageRule client: %+v", err)
	}
	configureFunc(mailFolderMessageRuleClient.Client)

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

	managedDeviceClient, err := manageddevice.NewManagedDeviceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDevice client: %+v", err)
	}
	configureFunc(managedDeviceClient.Client)

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

	managedDeviceLogCollectionRequestClient, err := manageddevicelogcollectionrequest.NewManagedDeviceLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceLogCollectionRequest client: %+v", err)
	}
	configureFunc(managedDeviceLogCollectionRequestClient.Client)

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

	oAuth2PermissionGrantClient, err := oauth2permissiongrant.NewOAuth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OAuth2PermissionGrant client: %+v", err)
	}
	configureFunc(oAuth2PermissionGrantClient.Client)

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

	onlineMeetingClient, err := onlinemeeting.NewOnlineMeetingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineMeeting client: %+v", err)
	}
	configureFunc(onlineMeetingClient.Client)

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

	permissionGrantClient, err := permissiongrant.NewPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PermissionGrant client: %+v", err)
	}
	configureFunc(permissionGrantClient.Client)

	personClient, err := person.NewPersonClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Person client: %+v", err)
	}
	configureFunc(personClient.Client)

	photoClient, err := photo.NewPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Photo client: %+v", err)
	}
	configureFunc(photoClient.Client)

	plannerClient, err := planner.NewPlannerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Planner client: %+v", err)
	}
	configureFunc(plannerClient.Client)

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

	settingItemInsightClient, err := settingiteminsight.NewSettingItemInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingItemInsight client: %+v", err)
	}
	configureFunc(settingItemInsightClient.Client)

	settingShiftPreferenceClient, err := settingshiftpreference.NewSettingShiftPreferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingShiftPreference client: %+v", err)
	}
	configureFunc(settingShiftPreferenceClient.Client)

	settingStorageClient, err := settingstorage.NewSettingStorageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingStorage client: %+v", err)
	}
	configureFunc(settingStorageClient.Client)

	settingStorageQuotaClient, err := settingstoragequota.NewSettingStorageQuotaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingStorageQuota client: %+v", err)
	}
	configureFunc(settingStorageQuotaClient.Client)

	settingStorageQuotaServiceClient, err := settingstoragequotaservice.NewSettingStorageQuotaServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingStorageQuotaService client: %+v", err)
	}
	configureFunc(settingStorageQuotaServiceClient.Client)

	settingWindowClient, err := settingwindow.NewSettingWindowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingWindow client: %+v", err)
	}
	configureFunc(settingWindowClient.Client)

	settingWindowInstanceClient, err := settingwindowinstance.NewSettingWindowInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SettingWindowInstance client: %+v", err)
	}
	configureFunc(settingWindowInstanceClient.Client)

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

	userClient, err := user.NewUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building User client: %+v", err)
	}
	configureFunc(userClient.Client)

	return &Client{
		Activity:                                          activityClient,
		ActivityHistoryItem:                               activityHistoryItemClient,
		ActivityHistoryItemActivity:                       activityHistoryItemActivityClient,
		AgreementAcceptance:                               agreementAcceptanceClient,
		AppRoleAssignment:                                 appRoleAssignmentClient,
		Authentication:                                    authenticationClient,
		AuthenticationEmailMethod:                         authenticationEmailMethodClient,
		AuthenticationFido2Method:                         authenticationFido2MethodClient,
		AuthenticationMethod:                              authenticationMethodClient,
		AuthenticationMicrosoftAuthenticatorMethod:        authenticationMicrosoftAuthenticatorMethodClient,
		AuthenticationMicrosoftAuthenticatorMethodDevice:  authenticationMicrosoftAuthenticatorMethodDeviceClient,
		AuthenticationOperation:                           authenticationOperationClient,
		AuthenticationPasswordMethod:                      authenticationPasswordMethodClient,
		AuthenticationPhoneMethod:                         authenticationPhoneMethodClient,
		AuthenticationSoftwareOathMethod:                  authenticationSoftwareOathMethodClient,
		AuthenticationTemporaryAccessPassMethod:           authenticationTemporaryAccessPassMethodClient,
		AuthenticationWindowsHelloForBusinessMethod:       authenticationWindowsHelloForBusinessMethodClient,
		AuthenticationWindowsHelloForBusinessMethodDevice: authenticationWindowsHelloForBusinessMethodDeviceClient,
		Calendar:                                            calendarClient,
		CalendarCalendarPermission:                          calendarCalendarPermissionClient,
		CalendarCalendarView:                                calendarCalendarViewClient,
		CalendarCalendarViewAttachment:                      calendarCalendarViewAttachmentClient,
		CalendarCalendarViewCalendar:                        calendarCalendarViewCalendarClient,
		CalendarCalendarViewExtension:                       calendarCalendarViewExtensionClient,
		CalendarCalendarViewInstance:                        calendarCalendarViewInstanceClient,
		CalendarCalendarViewInstanceAttachment:              calendarCalendarViewInstanceAttachmentClient,
		CalendarCalendarViewInstanceCalendar:                calendarCalendarViewInstanceCalendarClient,
		CalendarCalendarViewInstanceExtension:               calendarCalendarViewInstanceExtensionClient,
		CalendarEvent:                                       calendarEventClient,
		CalendarEventAttachment:                             calendarEventAttachmentClient,
		CalendarEventCalendar:                               calendarEventCalendarClient,
		CalendarEventExtension:                              calendarEventExtensionClient,
		CalendarEventInstance:                               calendarEventInstanceClient,
		CalendarEventInstanceAttachment:                     calendarEventInstanceAttachmentClient,
		CalendarEventInstanceCalendar:                       calendarEventInstanceCalendarClient,
		CalendarEventInstanceExtension:                      calendarEventInstanceExtensionClient,
		CalendarGroup:                                       calendarGroupClient,
		CalendarGroupCalendar:                               calendarGroupCalendarClient,
		CalendarGroupCalendarCalendarPermission:             calendarGroupCalendarCalendarPermissionClient,
		CalendarGroupCalendarCalendarView:                   calendarGroupCalendarCalendarViewClient,
		CalendarGroupCalendarCalendarViewAttachment:         calendarGroupCalendarCalendarViewAttachmentClient,
		CalendarGroupCalendarCalendarViewCalendar:           calendarGroupCalendarCalendarViewCalendarClient,
		CalendarGroupCalendarCalendarViewExtension:          calendarGroupCalendarCalendarViewExtensionClient,
		CalendarGroupCalendarCalendarViewInstance:           calendarGroupCalendarCalendarViewInstanceClient,
		CalendarGroupCalendarCalendarViewInstanceAttachment: calendarGroupCalendarCalendarViewInstanceAttachmentClient,
		CalendarGroupCalendarCalendarViewInstanceCalendar:   calendarGroupCalendarCalendarViewInstanceCalendarClient,
		CalendarGroupCalendarCalendarViewInstanceExtension:  calendarGroupCalendarCalendarViewInstanceExtensionClient,
		CalendarGroupCalendarEvent:                          calendarGroupCalendarEventClient,
		CalendarGroupCalendarEventAttachment:                calendarGroupCalendarEventAttachmentClient,
		CalendarGroupCalendarEventCalendar:                  calendarGroupCalendarEventCalendarClient,
		CalendarGroupCalendarEventExtension:                 calendarGroupCalendarEventExtensionClient,
		CalendarGroupCalendarEventInstance:                  calendarGroupCalendarEventInstanceClient,
		CalendarGroupCalendarEventInstanceAttachment:        calendarGroupCalendarEventInstanceAttachmentClient,
		CalendarGroupCalendarEventInstanceCalendar:          calendarGroupCalendarEventInstanceCalendarClient,
		CalendarGroupCalendarEventInstanceExtension:         calendarGroupCalendarEventInstanceExtensionClient,
		CalendarView:                               calendarViewClient,
		CalendarViewAttachment:                     calendarViewAttachmentClient,
		CalendarViewCalendar:                       calendarViewCalendarClient,
		CalendarViewExtension:                      calendarViewExtensionClient,
		CalendarViewInstance:                       calendarViewInstanceClient,
		CalendarViewInstanceAttachment:             calendarViewInstanceAttachmentClient,
		CalendarViewInstanceCalendar:               calendarViewInstanceCalendarClient,
		CalendarViewInstanceExtension:              calendarViewInstanceExtensionClient,
		Chat:                                       chatClient,
		ChatInstalledApp:                           chatInstalledAppClient,
		ChatInstalledAppTeamsApp:                   chatInstalledAppTeamsAppClient,
		ChatInstalledAppTeamsAppDefinition:         chatInstalledAppTeamsAppDefinitionClient,
		ChatLastMessagePreview:                     chatLastMessagePreviewClient,
		ChatMember:                                 chatMemberClient,
		ChatMessage:                                chatMessageClient,
		ChatMessageHostedContent:                   chatMessageHostedContentClient,
		ChatMessageReply:                           chatMessageReplyClient,
		ChatMessageReplyHostedContent:              chatMessageReplyHostedContentClient,
		ChatPermissionGrant:                        chatPermissionGrantClient,
		ChatPinnedMessage:                          chatPinnedMessageClient,
		ChatPinnedMessageMessage:                   chatPinnedMessageMessageClient,
		ChatTab:                                    chatTabClient,
		ChatTabTeamsApp:                            chatTabTeamsAppClient,
		CloudClipboard:                             cloudClipboardClient,
		CloudClipboardItem:                         cloudClipboardItemClient,
		Contact:                                    contactClient,
		ContactExtension:                           contactExtensionClient,
		ContactFolder:                              contactFolderClient,
		ContactFolderChildFolder:                   contactFolderChildFolderClient,
		ContactFolderChildFolderContact:            contactFolderChildFolderContactClient,
		ContactFolderChildFolderContactExtension:   contactFolderChildFolderContactExtensionClient,
		ContactFolderChildFolderContactPhoto:       contactFolderChildFolderContactPhotoClient,
		ContactFolderContact:                       contactFolderContactClient,
		ContactFolderContactExtension:              contactFolderContactExtensionClient,
		ContactFolderContactPhoto:                  contactFolderContactPhotoClient,
		ContactPhoto:                               contactPhotoClient,
		CreatedObject:                              createdObjectClient,
		DeviceManagementTroubleshootingEvent:       deviceManagementTroubleshootingEventClient,
		DirectReport:                               directReportClient,
		Drive:                                      driveClient,
		DriveBundle:                                driveBundleClient,
		DriveBundleContent:                         driveBundleContentClient,
		DriveCreatedByUser:                         driveCreatedByUserClient,
		DriveCreatedByUserMailboxSetting:           driveCreatedByUserMailboxSettingClient,
		DriveCreatedByUserServiceProvisioningError: driveCreatedByUserServiceProvisioningErrorClient,
		DriveFollowing:                             driveFollowingClient,
		DriveFollowingContent:                      driveFollowingContentClient,
		DriveItem:                                  driveItemClient,
		DriveItemAnalytics:                         driveItemAnalyticsClient,
		DriveItemAnalyticsAllTime:                  driveItemAnalyticsAllTimeClient,
		DriveItemAnalyticsItemActivityStat:         driveItemAnalyticsItemActivityStatClient,
		DriveItemAnalyticsItemActivityStatActivity: driveItemAnalyticsItemActivityStatActivityClient,
		DriveItemAnalyticsItemActivityStatActivityDriveItem:         driveItemAnalyticsItemActivityStatActivityDriveItemClient,
		DriveItemAnalyticsItemActivityStatActivityDriveItemContent:  driveItemAnalyticsItemActivityStatActivityDriveItemContentClient,
		DriveItemAnalyticsLastSevenDay:                              driveItemAnalyticsLastSevenDayClient,
		DriveItemChild:                                              driveItemChildClient,
		DriveItemChildContent:                                       driveItemChildContentClient,
		DriveItemContent:                                            driveItemContentClient,
		DriveItemCreatedByUser:                                      driveItemCreatedByUserClient,
		DriveItemCreatedByUserMailboxSetting:                        driveItemCreatedByUserMailboxSettingClient,
		DriveItemCreatedByUserServiceProvisioningError:              driveItemCreatedByUserServiceProvisioningErrorClient,
		DriveItemLastModifiedByUser:                                 driveItemLastModifiedByUserClient,
		DriveItemLastModifiedByUserMailboxSetting:                   driveItemLastModifiedByUserMailboxSettingClient,
		DriveItemLastModifiedByUserServiceProvisioningError:         driveItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveItemListItem:                                           driveItemListItemClient,
		DriveItemListItemAnalytics:                                  driveItemListItemAnalyticsClient,
		DriveItemListItemCreatedByUser:                              driveItemListItemCreatedByUserClient,
		DriveItemListItemCreatedByUserMailboxSetting:                driveItemListItemCreatedByUserMailboxSettingClient,
		DriveItemListItemCreatedByUserServiceProvisioningError:      driveItemListItemCreatedByUserServiceProvisioningErrorClient,
		DriveItemListItemDocumentSetVersion:                         driveItemListItemDocumentSetVersionClient,
		DriveItemListItemDocumentSetVersionField:                    driveItemListItemDocumentSetVersionFieldClient,
		DriveItemListItemDriveItem:                                  driveItemListItemDriveItemClient,
		DriveItemListItemDriveItemContent:                           driveItemListItemDriveItemContentClient,
		DriveItemListItemField:                                      driveItemListItemFieldClient,
		DriveItemListItemLastModifiedByUser:                         driveItemListItemLastModifiedByUserClient,
		DriveItemListItemLastModifiedByUserMailboxSetting:           driveItemListItemLastModifiedByUserMailboxSettingClient,
		DriveItemListItemLastModifiedByUserServiceProvisioningError: driveItemListItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveItemListItemVersion:                                    driveItemListItemVersionClient,
		DriveItemListItemVersionField:                               driveItemListItemVersionFieldClient,
		DriveItemPermission:                                         driveItemPermissionClient,
		DriveItemRetentionLabel:                                     driveItemRetentionLabelClient,
		DriveItemSubscription:                                       driveItemSubscriptionClient,
		DriveItemThumbnail:                                          driveItemThumbnailClient,
		DriveItemVersion:                                            driveItemVersionClient,
		DriveItemVersionContent:                                     driveItemVersionContentClient,
		DriveLastModifiedByUser:                                     driveLastModifiedByUserClient,
		DriveLastModifiedByUserMailboxSetting:                       driveLastModifiedByUserMailboxSettingClient,
		DriveLastModifiedByUserServiceProvisioningError:             driveLastModifiedByUserServiceProvisioningErrorClient,
		DriveList:                                                   driveListClient,
		DriveListColumn:                                             driveListColumnClient,
		DriveListColumnSourceColumn:                                 driveListColumnSourceColumnClient,
		DriveListContentType:                                        driveListContentTypeClient,
		DriveListContentTypeBase:                                    driveListContentTypeBaseClient,
		DriveListContentTypeBaseType:                                driveListContentTypeBaseTypeClient,
		DriveListContentTypeColumn:                                  driveListContentTypeColumnClient,
		DriveListContentTypeColumnLink:                              driveListContentTypeColumnLinkClient,
		DriveListContentTypeColumnPosition:                          driveListContentTypeColumnPositionClient,
		DriveListContentTypeColumnSourceColumn:                      driveListContentTypeColumnSourceColumnClient,
		DriveListCreatedByUser:                                      driveListCreatedByUserClient,
		DriveListCreatedByUserMailboxSetting:                        driveListCreatedByUserMailboxSettingClient,
		DriveListCreatedByUserServiceProvisioningError:              driveListCreatedByUserServiceProvisioningErrorClient,
		DriveListDrive:                                              driveListDriveClient,
		DriveListItem:                                               driveListItemClient,
		DriveListItemAnalytics:                                      driveListItemAnalyticsClient,
		DriveListItemCreatedByUser:                                  driveListItemCreatedByUserClient,
		DriveListItemCreatedByUserMailboxSetting:                    driveListItemCreatedByUserMailboxSettingClient,
		DriveListItemCreatedByUserServiceProvisioningError:          driveListItemCreatedByUserServiceProvisioningErrorClient,
		DriveListItemDocumentSetVersion:                             driveListItemDocumentSetVersionClient,
		DriveListItemDocumentSetVersionField:                        driveListItemDocumentSetVersionFieldClient,
		DriveListItemDriveItem:                                      driveListItemDriveItemClient,
		DriveListItemDriveItemContent:                               driveListItemDriveItemContentClient,
		DriveListItemField:                                          driveListItemFieldClient,
		DriveListItemLastModifiedByUser:                             driveListItemLastModifiedByUserClient,
		DriveListItemLastModifiedByUserMailboxSetting:               driveListItemLastModifiedByUserMailboxSettingClient,
		DriveListItemLastModifiedByUserServiceProvisioningError:     driveListItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveListItemVersion:                                        driveListItemVersionClient,
		DriveListItemVersionField:                                   driveListItemVersionFieldClient,
		DriveListLastModifiedByUser:                                 driveListLastModifiedByUserClient,
		DriveListLastModifiedByUserMailboxSetting:                   driveListLastModifiedByUserMailboxSettingClient,
		DriveListLastModifiedByUserServiceProvisioningError:         driveListLastModifiedByUserServiceProvisioningErrorClient,
		DriveListOperation:                                          driveListOperationClient,
		DriveListSubscription:                                       driveListSubscriptionClient,
		DriveRoot:                                                   driveRootClient,
		DriveRootAnalytics:                                          driveRootAnalyticsClient,
		DriveRootAnalyticsAllTime:                                   driveRootAnalyticsAllTimeClient,
		DriveRootAnalyticsItemActivityStat:                          driveRootAnalyticsItemActivityStatClient,
		DriveRootAnalyticsItemActivityStatActivity:                  driveRootAnalyticsItemActivityStatActivityClient,
		DriveRootAnalyticsItemActivityStatActivityDriveItem:         driveRootAnalyticsItemActivityStatActivityDriveItemClient,
		DriveRootAnalyticsItemActivityStatActivityDriveItemContent:  driveRootAnalyticsItemActivityStatActivityDriveItemContentClient,
		DriveRootAnalyticsLastSevenDay:                              driveRootAnalyticsLastSevenDayClient,
		DriveRootChild:                                              driveRootChildClient,
		DriveRootChildContent:                                       driveRootChildContentClient,
		DriveRootContent:                                            driveRootContentClient,
		DriveRootCreatedByUser:                                      driveRootCreatedByUserClient,
		DriveRootCreatedByUserMailboxSetting:                        driveRootCreatedByUserMailboxSettingClient,
		DriveRootCreatedByUserServiceProvisioningError:              driveRootCreatedByUserServiceProvisioningErrorClient,
		DriveRootLastModifiedByUser:                                 driveRootLastModifiedByUserClient,
		DriveRootLastModifiedByUserMailboxSetting:                   driveRootLastModifiedByUserMailboxSettingClient,
		DriveRootLastModifiedByUserServiceProvisioningError:         driveRootLastModifiedByUserServiceProvisioningErrorClient,
		DriveRootListItem:                                           driveRootListItemClient,
		DriveRootListItemAnalytics:                                  driveRootListItemAnalyticsClient,
		DriveRootListItemCreatedByUser:                              driveRootListItemCreatedByUserClient,
		DriveRootListItemCreatedByUserMailboxSetting:                driveRootListItemCreatedByUserMailboxSettingClient,
		DriveRootListItemCreatedByUserServiceProvisioningError:      driveRootListItemCreatedByUserServiceProvisioningErrorClient,
		DriveRootListItemDocumentSetVersion:                         driveRootListItemDocumentSetVersionClient,
		DriveRootListItemDocumentSetVersionField:                    driveRootListItemDocumentSetVersionFieldClient,
		DriveRootListItemDriveItem:                                  driveRootListItemDriveItemClient,
		DriveRootListItemDriveItemContent:                           driveRootListItemDriveItemContentClient,
		DriveRootListItemField:                                      driveRootListItemFieldClient,
		DriveRootListItemLastModifiedByUser:                         driveRootListItemLastModifiedByUserClient,
		DriveRootListItemLastModifiedByUserMailboxSetting:           driveRootListItemLastModifiedByUserMailboxSettingClient,
		DriveRootListItemLastModifiedByUserServiceProvisioningError: driveRootListItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveRootListItemVersion:                                    driveRootListItemVersionClient,
		DriveRootListItemVersionField:                               driveRootListItemVersionFieldClient,
		DriveRootPermission:                                         driveRootPermissionClient,
		DriveRootRetentionLabel:                                     driveRootRetentionLabelClient,
		DriveRootSubscription:                                       driveRootSubscriptionClient,
		DriveRootThumbnail:                                          driveRootThumbnailClient,
		DriveRootVersion:                                            driveRootVersionClient,
		DriveRootVersionContent:                                     driveRootVersionContentClient,
		DriveSpecial:                                                driveSpecialClient,
		DriveSpecialContent:                                         driveSpecialContentClient,
		EmployeeExperience:                                          employeeExperienceClient,
		EmployeeExperienceLearningCourseActivity:                    employeeExperienceLearningCourseActivityClient,
		Event:                                                       eventClient,
		EventAttachment:                                             eventAttachmentClient,
		EventCalendar:                                               eventCalendarClient,
		EventExtension:                                              eventExtensionClient,
		EventInstance:                                               eventInstanceClient,
		EventInstanceAttachment:                                     eventInstanceAttachmentClient,
		EventInstanceCalendar:                                       eventInstanceCalendarClient,
		EventInstanceExtension:                                      eventInstanceExtensionClient,
		Extension:                                                   extensionClient,
		FollowedSite:                                                followedSiteClient,
		InferenceClassification:                                     inferenceClassificationClient,
		InferenceClassificationOverride:                             inferenceClassificationOverrideClient,
		Insight:                                                     insightClient,
		InsightShared:                                               insightSharedClient,
		InsightSharedLastSharedMethod:                               insightSharedLastSharedMethodClient,
		InsightSharedResource:                                       insightSharedResourceClient,
		InsightTrending:                                             insightTrendingClient,
		InsightTrendingResource:                                     insightTrendingResourceClient,
		InsightUsed:                                                 insightUsedClient,
		InsightUsedResource:                                         insightUsedResourceClient,
		JoinedTeam:                                                  joinedTeamClient,
		JoinedTeamAllChannel:                                        joinedTeamAllChannelClient,
		JoinedTeamChannel:                                           joinedTeamChannelClient,
		JoinedTeamChannelFilesFolder:                                joinedTeamChannelFilesFolderClient,
		JoinedTeamChannelFilesFolderContent:                         joinedTeamChannelFilesFolderContentClient,
		JoinedTeamChannelMember:                                     joinedTeamChannelMemberClient,
		JoinedTeamChannelMessage:                                    joinedTeamChannelMessageClient,
		JoinedTeamChannelMessageHostedContent:                       joinedTeamChannelMessageHostedContentClient,
		JoinedTeamChannelMessageReply:                               joinedTeamChannelMessageReplyClient,
		JoinedTeamChannelMessageReplyHostedContent:                  joinedTeamChannelMessageReplyHostedContentClient,
		JoinedTeamChannelSharedWithTeam:                             joinedTeamChannelSharedWithTeamClient,
		JoinedTeamChannelSharedWithTeamAllowedMember:                joinedTeamChannelSharedWithTeamAllowedMemberClient,
		JoinedTeamChannelSharedWithTeamTeam:                         joinedTeamChannelSharedWithTeamTeamClient,
		JoinedTeamChannelTab:                                        joinedTeamChannelTabClient,
		JoinedTeamChannelTabTeamsApp:                                joinedTeamChannelTabTeamsAppClient,
		JoinedTeamGroup:                                             joinedTeamGroupClient,
		JoinedTeamGroupServiceProvisioningError:                     joinedTeamGroupServiceProvisioningErrorClient,
		JoinedTeamIncomingChannel:                                   joinedTeamIncomingChannelClient,
		JoinedTeamInstalledApp:                                      joinedTeamInstalledAppClient,
		JoinedTeamInstalledAppTeamsApp:                              joinedTeamInstalledAppTeamsAppClient,
		JoinedTeamInstalledAppTeamsAppDefinition:                    joinedTeamInstalledAppTeamsAppDefinitionClient,
		JoinedTeamMember:                                            joinedTeamMemberClient,
		JoinedTeamOperation:                                         joinedTeamOperationClient,
		JoinedTeamPermissionGrant:                                   joinedTeamPermissionGrantClient,
		JoinedTeamPhoto:                                             joinedTeamPhotoClient,
		JoinedTeamPrimaryChannel:                                    joinedTeamPrimaryChannelClient,
		JoinedTeamPrimaryChannelFilesFolder:                         joinedTeamPrimaryChannelFilesFolderClient,
		JoinedTeamPrimaryChannelFilesFolderContent:                  joinedTeamPrimaryChannelFilesFolderContentClient,
		JoinedTeamPrimaryChannelMember:                              joinedTeamPrimaryChannelMemberClient,
		JoinedTeamPrimaryChannelMessage:                             joinedTeamPrimaryChannelMessageClient,
		JoinedTeamPrimaryChannelMessageHostedContent:                joinedTeamPrimaryChannelMessageHostedContentClient,
		JoinedTeamPrimaryChannelMessageReply:                        joinedTeamPrimaryChannelMessageReplyClient,
		JoinedTeamPrimaryChannelMessageReplyHostedContent:           joinedTeamPrimaryChannelMessageReplyHostedContentClient,
		JoinedTeamPrimaryChannelSharedWithTeam:                      joinedTeamPrimaryChannelSharedWithTeamClient,
		JoinedTeamPrimaryChannelSharedWithTeamAllowedMember:         joinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient,
		JoinedTeamPrimaryChannelSharedWithTeamTeam:                  joinedTeamPrimaryChannelSharedWithTeamTeamClient,
		JoinedTeamPrimaryChannelTab:                                 joinedTeamPrimaryChannelTabClient,
		JoinedTeamPrimaryChannelTabTeamsApp:                         joinedTeamPrimaryChannelTabTeamsAppClient,
		JoinedTeamSchedule:                                          joinedTeamScheduleClient,
		JoinedTeamScheduleOfferShiftRequest:                         joinedTeamScheduleOfferShiftRequestClient,
		JoinedTeamScheduleOpenShift:                                 joinedTeamScheduleOpenShiftClient,
		JoinedTeamScheduleOpenShiftChangeRequest:                    joinedTeamScheduleOpenShiftChangeRequestClient,
		JoinedTeamScheduleSchedulingGroup:                           joinedTeamScheduleSchedulingGroupClient,
		JoinedTeamScheduleShift:                                     joinedTeamScheduleShiftClient,
		JoinedTeamScheduleSwapShiftsChangeRequest:                   joinedTeamScheduleSwapShiftsChangeRequestClient,
		JoinedTeamScheduleTimeOffReason:                             joinedTeamScheduleTimeOffReasonClient,
		JoinedTeamScheduleTimeOffRequest:                            joinedTeamScheduleTimeOffRequestClient,
		JoinedTeamScheduleTimesOff:                                  joinedTeamScheduleTimesOffClient,
		JoinedTeamTag:                                               joinedTeamTagClient,
		JoinedTeamTagMember:                                         joinedTeamTagMemberClient,
		JoinedTeamTemplate:                                          joinedTeamTemplateClient,
		LicenseDetail:                                               licenseDetailClient,
		MailFolder:                                                  mailFolderClient,
		MailFolderChildFolder:                                       mailFolderChildFolderClient,
		MailFolderChildFolderMessage:                                mailFolderChildFolderMessageClient,
		MailFolderChildFolderMessageAttachment:                      mailFolderChildFolderMessageAttachmentClient,
		MailFolderChildFolderMessageExtension:                       mailFolderChildFolderMessageExtensionClient,
		MailFolderChildFolderMessageRule:                            mailFolderChildFolderMessageRuleClient,
		MailFolderMessage:                                           mailFolderMessageClient,
		MailFolderMessageAttachment:                                 mailFolderMessageAttachmentClient,
		MailFolderMessageExtension:                                  mailFolderMessageExtensionClient,
		MailFolderMessageRule:                                       mailFolderMessageRuleClient,
		MailboxSetting:                                              mailboxSettingClient,
		ManagedAppRegistration:                                      managedAppRegistrationClient,
		ManagedDevice:                                               managedDeviceClient,
		ManagedDeviceDeviceCategory:                                 managedDeviceDeviceCategoryClient,
		ManagedDeviceDeviceCompliancePolicyState:                    managedDeviceDeviceCompliancePolicyStateClient,
		ManagedDeviceDeviceConfigurationState:                       managedDeviceDeviceConfigurationStateClient,
		ManagedDeviceLogCollectionRequest:                           managedDeviceLogCollectionRequestClient,
		ManagedDeviceUser:                                           managedDeviceUserClient,
		ManagedDeviceWindowsProtectionState:                         managedDeviceWindowsProtectionStateClient,
		ManagedDeviceWindowsProtectionStateDetectedMalwareState:     managedDeviceWindowsProtectionStateDetectedMalwareStateClient,
		Manager:                       managerClient,
		MemberOf:                      memberOfClient,
		Message:                       messageClient,
		MessageAttachment:             messageAttachmentClient,
		MessageExtension:              messageExtensionClient,
		OAuth2PermissionGrant:         oAuth2PermissionGrantClient,
		OnlineMeeting:                 onlineMeetingClient,
		OnlineMeetingAttendanceReport: onlineMeetingAttendanceReportClient,
		OnlineMeetingAttendanceReportAttendanceRecord: onlineMeetingAttendanceReportAttendanceRecordClient,
		OnlineMeetingAttendeeReport:                   onlineMeetingAttendeeReportClient,
		OnlineMeetingRecording:                        onlineMeetingRecordingClient,
		OnlineMeetingRecordingContent:                 onlineMeetingRecordingContentClient,
		OnlineMeetingTranscript:                       onlineMeetingTranscriptClient,
		OnlineMeetingTranscriptContent:                onlineMeetingTranscriptContentClient,
		OnlineMeetingTranscriptMetadataContent:        onlineMeetingTranscriptMetadataContentClient,
		Outlook:                                       outlookClient,
		OutlookMasterCategory:                         outlookMasterCategoryClient,
		OwnedDevice:                                   ownedDeviceClient,
		OwnedObject:                                   ownedObjectClient,
		PermissionGrant:                               permissionGrantClient,
		Person:                                        personClient,
		Photo:                                         photoClient,
		Planner:                                       plannerClient,
		PlannerPlan:                                   plannerPlanClient,
		PlannerPlanBucket:                             plannerPlanBucketClient,
		PlannerPlanBucketTask:                         plannerPlanBucketTaskClient,
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
		PlannerTask:                                    plannerTaskClient,
		PlannerTaskAssignedToTaskBoardFormat:           plannerTaskAssignedToTaskBoardFormatClient,
		PlannerTaskBucketTaskBoardFormat:               plannerTaskBucketTaskBoardFormatClient,
		PlannerTaskDetail:                              plannerTaskDetailClient,
		PlannerTaskProgressTaskBoardFormat:             plannerTaskProgressTaskBoardFormatClient,
		Presence:                                       presenceClient,
		RegisteredDevice:                               registeredDeviceClient,
		ScopedRoleMemberOf:                             scopedRoleMemberOfClient,
		ServiceProvisioningError:                       serviceProvisioningErrorClient,
		Setting:                                        settingClient,
		SettingItemInsight:                             settingItemInsightClient,
		SettingShiftPreference:                         settingShiftPreferenceClient,
		SettingStorage:                                 settingStorageClient,
		SettingStorageQuota:                            settingStorageQuotaClient,
		SettingStorageQuotaService:                     settingStorageQuotaServiceClient,
		SettingWindow:                                  settingWindowClient,
		SettingWindowInstance:                          settingWindowInstanceClient,
		Sponsor:                                        sponsorClient,
		Teamwork:                                       teamworkClient,
		TeamworkAssociatedTeam:                         teamworkAssociatedTeamClient,
		TeamworkAssociatedTeamTeam:                     teamworkAssociatedTeamTeamClient,
		TeamworkInstalledApp:                           teamworkInstalledAppClient,
		TeamworkInstalledAppChat:                       teamworkInstalledAppChatClient,
		TeamworkInstalledAppTeamsApp:                   teamworkInstalledAppTeamsAppClient,
		TeamworkInstalledAppTeamsAppDefinition:         teamworkInstalledAppTeamsAppDefinitionClient,
		Todo:                                           todoClient,
		TodoList:                                       todoListClient,
		TodoListExtension:                              todoListExtensionClient,
		TodoListTask:                                   todoListTaskClient,
		TodoListTaskAttachment:                         todoListTaskAttachmentClient,
		TodoListTaskAttachmentSession:                  todoListTaskAttachmentSessionClient,
		TodoListTaskAttachmentSessionContent:           todoListTaskAttachmentSessionContentClient,
		TodoListTaskChecklistItem:                      todoListTaskChecklistItemClient,
		TodoListTaskExtension:                          todoListTaskExtensionClient,
		TodoListTaskLinkedResource:                     todoListTaskLinkedResourceClient,
		TransitiveMemberOf:                             transitiveMemberOfClient,
		User:                                           userClient,
	}, nil
}
