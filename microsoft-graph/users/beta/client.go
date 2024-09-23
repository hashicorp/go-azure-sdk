package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/activity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/activityhistoryitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/activityhistoryitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/agreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/analytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/analyticsactivitystatistic"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/cloudclipboard"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/cloudclipboarditem"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivebundle"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivebundlecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivebundlecontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivefollowing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivefollowingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivefollowingcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalyticsitemactivitystatactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemchildcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemdriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveitemversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemdriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivelistsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driveroot"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalyticsitemactivitystatactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootchildcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemdriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/driverootversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivespecial"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivespecialcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/drivespecialcontentstream"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/invitedby"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/managedapplogcollectionrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/managedappregistration"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddeviceassignmentfilterevaluationstatusdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedetectedapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicecategory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicecompliancepolicystate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedeviceconfigurationstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicehealthscriptstate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotenotebooksectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenoteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotepagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotepageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotepageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenoteresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenoteresourcecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onenotesectionparentsectiongroup"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedecisioninstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstancedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/pendingaccessreviewinstancestagedecisioninstancedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/person"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/photo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/planner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerall"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannerfavoriteplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/plannermydaytask"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingstorage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingstoragequota"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingstoragequotaservice"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingwindow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/settingwindowinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/solution"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/solutionworkingtimeschedule"
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
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Activity                                                                 *activity.ActivityClient
	ActivityHistoryItem                                                      *activityhistoryitem.ActivityHistoryItemClient
	ActivityHistoryItemActivity                                              *activityhistoryitemactivity.ActivityHistoryItemActivityClient
	AgreementAcceptance                                                      *agreementacceptance.AgreementAcceptanceClient
	Analytics                                                                *analytics.AnalyticsClient
	AnalyticsActivityStatistic                                               *analyticsactivitystatistic.AnalyticsActivityStatisticClient
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
	CloudClipboard                                                           *cloudclipboard.CloudClipboardClient
	CloudClipboardItem                                                       *cloudclipboarditem.CloudClipboardItemClient
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
	DriveActivity                                                            *driveactivity.DriveActivityClient
	DriveActivityDriveItem                                                   *driveactivitydriveitem.DriveActivityDriveItemClient
	DriveActivityDriveItemContent                                            *driveactivitydriveitemcontent.DriveActivityDriveItemContentClient
	DriveActivityDriveItemContentStream                                      *driveactivitydriveitemcontentstream.DriveActivityDriveItemContentStreamClient
	DriveActivityListItem                                                    *driveactivitylistitem.DriveActivityListItemClient
	DriveBundle                                                              *drivebundle.DriveBundleClient
	DriveBundleContent                                                       *drivebundlecontent.DriveBundleContentClient
	DriveBundleContentStream                                                 *drivebundlecontentstream.DriveBundleContentStreamClient
	DriveCreatedByUser                                                       *drivecreatedbyuser.DriveCreatedByUserClient
	DriveCreatedByUserMailboxSetting                                         *drivecreatedbyusermailboxsetting.DriveCreatedByUserMailboxSettingClient
	DriveCreatedByUserServiceProvisioningError                               *drivecreatedbyuserserviceprovisioningerror.DriveCreatedByUserServiceProvisioningErrorClient
	DriveFollowing                                                           *drivefollowing.DriveFollowingClient
	DriveFollowingContent                                                    *drivefollowingcontent.DriveFollowingContentClient
	DriveFollowingContentStream                                              *drivefollowingcontentstream.DriveFollowingContentStreamClient
	DriveItem                                                                *driveitem.DriveItemClient
	DriveItemActivity                                                        *driveitemactivity.DriveItemActivityClient
	DriveItemAnalytics                                                       *driveitemanalytics.DriveItemAnalyticsClient
	DriveItemAnalyticsAllTime                                                *driveitemanalyticsalltime.DriveItemAnalyticsAllTimeClient
	DriveItemAnalyticsItemActivityStat                                       *driveitemanalyticsitemactivitystat.DriveItemAnalyticsItemActivityStatClient
	DriveItemAnalyticsItemActivityStatActivity                               *driveitemanalyticsitemactivitystatactivity.DriveItemAnalyticsItemActivityStatActivityClient
	DriveItemAnalyticsItemActivityStatActivityDriveItem                      *driveitemanalyticsitemactivitystatactivitydriveitem.DriveItemAnalyticsItemActivityStatActivityDriveItemClient
	DriveItemAnalyticsItemActivityStatActivityDriveItemContent               *driveitemanalyticsitemactivitystatactivitydriveitemcontent.DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient
	DriveItemAnalyticsItemActivityStatActivityDriveItemContentStream         *driveitemanalyticsitemactivitystatactivitydriveitemcontentstream.DriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient
	DriveItemAnalyticsLastSevenDay                                           *driveitemanalyticslastsevenday.DriveItemAnalyticsLastSevenDayClient
	DriveItemChild                                                           *driveitemchild.DriveItemChildClient
	DriveItemChildContent                                                    *driveitemchildcontent.DriveItemChildContentClient
	DriveItemChildContentStream                                              *driveitemchildcontentstream.DriveItemChildContentStreamClient
	DriveItemContent                                                         *driveitemcontent.DriveItemContentClient
	DriveItemContentStream                                                   *driveitemcontentstream.DriveItemContentStreamClient
	DriveItemCreatedByUser                                                   *driveitemcreatedbyuser.DriveItemCreatedByUserClient
	DriveItemCreatedByUserMailboxSetting                                     *driveitemcreatedbyusermailboxsetting.DriveItemCreatedByUserMailboxSettingClient
	DriveItemCreatedByUserServiceProvisioningError                           *driveitemcreatedbyuserserviceprovisioningerror.DriveItemCreatedByUserServiceProvisioningErrorClient
	DriveItemLastModifiedByUser                                              *driveitemlastmodifiedbyuser.DriveItemLastModifiedByUserClient
	DriveItemLastModifiedByUserMailboxSetting                                *driveitemlastmodifiedbyusermailboxsetting.DriveItemLastModifiedByUserMailboxSettingClient
	DriveItemLastModifiedByUserServiceProvisioningError                      *driveitemlastmodifiedbyuserserviceprovisioningerror.DriveItemLastModifiedByUserServiceProvisioningErrorClient
	DriveItemListItem                                                        *driveitemlistitem.DriveItemListItemClient
	DriveItemListItemActivity                                                *driveitemlistitemactivity.DriveItemListItemActivityClient
	DriveItemListItemActivityDriveItem                                       *driveitemlistitemactivitydriveitem.DriveItemListItemActivityDriveItemClient
	DriveItemListItemActivityDriveItemContent                                *driveitemlistitemactivitydriveitemcontent.DriveItemListItemActivityDriveItemContentClient
	DriveItemListItemActivityDriveItemContentStream                          *driveitemlistitemactivitydriveitemcontentstream.DriveItemListItemActivityDriveItemContentStreamClient
	DriveItemListItemActivityListItem                                        *driveitemlistitemactivitylistitem.DriveItemListItemActivityListItemClient
	DriveItemListItemAnalytics                                               *driveitemlistitemanalytics.DriveItemListItemAnalyticsClient
	DriveItemListItemCreatedByUser                                           *driveitemlistitemcreatedbyuser.DriveItemListItemCreatedByUserClient
	DriveItemListItemCreatedByUserMailboxSetting                             *driveitemlistitemcreatedbyusermailboxsetting.DriveItemListItemCreatedByUserMailboxSettingClient
	DriveItemListItemCreatedByUserServiceProvisioningError                   *driveitemlistitemcreatedbyuserserviceprovisioningerror.DriveItemListItemCreatedByUserServiceProvisioningErrorClient
	DriveItemListItemDocumentSetVersion                                      *driveitemlistitemdocumentsetversion.DriveItemListItemDocumentSetVersionClient
	DriveItemListItemDocumentSetVersionField                                 *driveitemlistitemdocumentsetversionfield.DriveItemListItemDocumentSetVersionFieldClient
	DriveItemListItemDriveItem                                               *driveitemlistitemdriveitem.DriveItemListItemDriveItemClient
	DriveItemListItemDriveItemContent                                        *driveitemlistitemdriveitemcontent.DriveItemListItemDriveItemContentClient
	DriveItemListItemDriveItemContentStream                                  *driveitemlistitemdriveitemcontentstream.DriveItemListItemDriveItemContentStreamClient
	DriveItemListItemField                                                   *driveitemlistitemfield.DriveItemListItemFieldClient
	DriveItemListItemLastModifiedByUser                                      *driveitemlistitemlastmodifiedbyuser.DriveItemListItemLastModifiedByUserClient
	DriveItemListItemLastModifiedByUserMailboxSetting                        *driveitemlistitemlastmodifiedbyusermailboxsetting.DriveItemListItemLastModifiedByUserMailboxSettingClient
	DriveItemListItemLastModifiedByUserServiceProvisioningError              *driveitemlistitemlastmodifiedbyuserserviceprovisioningerror.DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveItemListItemPermission                                              *driveitemlistitempermission.DriveItemListItemPermissionClient
	DriveItemListItemVersion                                                 *driveitemlistitemversion.DriveItemListItemVersionClient
	DriveItemListItemVersionField                                            *driveitemlistitemversionfield.DriveItemListItemVersionFieldClient
	DriveItemPermission                                                      *driveitempermission.DriveItemPermissionClient
	DriveItemRetentionLabel                                                  *driveitemretentionlabel.DriveItemRetentionLabelClient
	DriveItemSubscription                                                    *driveitemsubscription.DriveItemSubscriptionClient
	DriveItemThumbnail                                                       *driveitemthumbnail.DriveItemThumbnailClient
	DriveItemVersion                                                         *driveitemversion.DriveItemVersionClient
	DriveItemVersionContent                                                  *driveitemversioncontent.DriveItemVersionContentClient
	DriveLastModifiedByUser                                                  *drivelastmodifiedbyuser.DriveLastModifiedByUserClient
	DriveLastModifiedByUserMailboxSetting                                    *drivelastmodifiedbyusermailboxsetting.DriveLastModifiedByUserMailboxSettingClient
	DriveLastModifiedByUserServiceProvisioningError                          *drivelastmodifiedbyuserserviceprovisioningerror.DriveLastModifiedByUserServiceProvisioningErrorClient
	DriveList                                                                *drivelist.DriveListClient
	DriveListActivity                                                        *drivelistactivity.DriveListActivityClient
	DriveListColumn                                                          *drivelistcolumn.DriveListColumnClient
	DriveListColumnSourceColumn                                              *drivelistcolumnsourcecolumn.DriveListColumnSourceColumnClient
	DriveListContentType                                                     *drivelistcontenttype.DriveListContentTypeClient
	DriveListContentTypeBase                                                 *drivelistcontenttypebase.DriveListContentTypeBaseClient
	DriveListContentTypeBaseType                                             *drivelistcontenttypebasetype.DriveListContentTypeBaseTypeClient
	DriveListContentTypeColumn                                               *drivelistcontenttypecolumn.DriveListContentTypeColumnClient
	DriveListContentTypeColumnLink                                           *drivelistcontenttypecolumnlink.DriveListContentTypeColumnLinkClient
	DriveListContentTypeColumnPosition                                       *drivelistcontenttypecolumnposition.DriveListContentTypeColumnPositionClient
	DriveListContentTypeColumnSourceColumn                                   *drivelistcontenttypecolumnsourcecolumn.DriveListContentTypeColumnSourceColumnClient
	DriveListCreatedByUser                                                   *drivelistcreatedbyuser.DriveListCreatedByUserClient
	DriveListCreatedByUserMailboxSetting                                     *drivelistcreatedbyusermailboxsetting.DriveListCreatedByUserMailboxSettingClient
	DriveListCreatedByUserServiceProvisioningError                           *drivelistcreatedbyuserserviceprovisioningerror.DriveListCreatedByUserServiceProvisioningErrorClient
	DriveListDrive                                                           *drivelistdrive.DriveListDriveClient
	DriveListItem                                                            *drivelistitem.DriveListItemClient
	DriveListItemActivity                                                    *drivelistitemactivity.DriveListItemActivityClient
	DriveListItemActivityDriveItem                                           *drivelistitemactivitydriveitem.DriveListItemActivityDriveItemClient
	DriveListItemActivityDriveItemContent                                    *drivelistitemactivitydriveitemcontent.DriveListItemActivityDriveItemContentClient
	DriveListItemActivityDriveItemContentStream                              *drivelistitemactivitydriveitemcontentstream.DriveListItemActivityDriveItemContentStreamClient
	DriveListItemActivityListItem                                            *drivelistitemactivitylistitem.DriveListItemActivityListItemClient
	DriveListItemAnalytics                                                   *drivelistitemanalytics.DriveListItemAnalyticsClient
	DriveListItemCreatedByUser                                               *drivelistitemcreatedbyuser.DriveListItemCreatedByUserClient
	DriveListItemCreatedByUserMailboxSetting                                 *drivelistitemcreatedbyusermailboxsetting.DriveListItemCreatedByUserMailboxSettingClient
	DriveListItemCreatedByUserServiceProvisioningError                       *drivelistitemcreatedbyuserserviceprovisioningerror.DriveListItemCreatedByUserServiceProvisioningErrorClient
	DriveListItemDocumentSetVersion                                          *drivelistitemdocumentsetversion.DriveListItemDocumentSetVersionClient
	DriveListItemDocumentSetVersionField                                     *drivelistitemdocumentsetversionfield.DriveListItemDocumentSetVersionFieldClient
	DriveListItemDriveItem                                                   *drivelistitemdriveitem.DriveListItemDriveItemClient
	DriveListItemDriveItemContent                                            *drivelistitemdriveitemcontent.DriveListItemDriveItemContentClient
	DriveListItemDriveItemContentStream                                      *drivelistitemdriveitemcontentstream.DriveListItemDriveItemContentStreamClient
	DriveListItemField                                                       *drivelistitemfield.DriveListItemFieldClient
	DriveListItemLastModifiedByUser                                          *drivelistitemlastmodifiedbyuser.DriveListItemLastModifiedByUserClient
	DriveListItemLastModifiedByUserMailboxSetting                            *drivelistitemlastmodifiedbyusermailboxsetting.DriveListItemLastModifiedByUserMailboxSettingClient
	DriveListItemLastModifiedByUserServiceProvisioningError                  *drivelistitemlastmodifiedbyuserserviceprovisioningerror.DriveListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveListItemPermission                                                  *drivelistitempermission.DriveListItemPermissionClient
	DriveListItemVersion                                                     *drivelistitemversion.DriveListItemVersionClient
	DriveListItemVersionField                                                *drivelistitemversionfield.DriveListItemVersionFieldClient
	DriveListLastModifiedByUser                                              *drivelistlastmodifiedbyuser.DriveListLastModifiedByUserClient
	DriveListLastModifiedByUserMailboxSetting                                *drivelistlastmodifiedbyusermailboxsetting.DriveListLastModifiedByUserMailboxSettingClient
	DriveListLastModifiedByUserServiceProvisioningError                      *drivelistlastmodifiedbyuserserviceprovisioningerror.DriveListLastModifiedByUserServiceProvisioningErrorClient
	DriveListOperation                                                       *drivelistoperation.DriveListOperationClient
	DriveListPermission                                                      *drivelistpermission.DriveListPermissionClient
	DriveListSubscription                                                    *drivelistsubscription.DriveListSubscriptionClient
	DriveRoot                                                                *driveroot.DriveRootClient
	DriveRootActivity                                                        *driverootactivity.DriveRootActivityClient
	DriveRootAnalytics                                                       *driverootanalytics.DriveRootAnalyticsClient
	DriveRootAnalyticsAllTime                                                *driverootanalyticsalltime.DriveRootAnalyticsAllTimeClient
	DriveRootAnalyticsItemActivityStat                                       *driverootanalyticsitemactivitystat.DriveRootAnalyticsItemActivityStatClient
	DriveRootAnalyticsItemActivityStatActivity                               *driverootanalyticsitemactivitystatactivity.DriveRootAnalyticsItemActivityStatActivityClient
	DriveRootAnalyticsItemActivityStatActivityDriveItem                      *driverootanalyticsitemactivitystatactivitydriveitem.DriveRootAnalyticsItemActivityStatActivityDriveItemClient
	DriveRootAnalyticsItemActivityStatActivityDriveItemContent               *driverootanalyticsitemactivitystatactivitydriveitemcontent.DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient
	DriveRootAnalyticsItemActivityStatActivityDriveItemContentStream         *driverootanalyticsitemactivitystatactivitydriveitemcontentstream.DriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClient
	DriveRootAnalyticsLastSevenDay                                           *driverootanalyticslastsevenday.DriveRootAnalyticsLastSevenDayClient
	DriveRootChild                                                           *driverootchild.DriveRootChildClient
	DriveRootChildContent                                                    *driverootchildcontent.DriveRootChildContentClient
	DriveRootChildContentStream                                              *driverootchildcontentstream.DriveRootChildContentStreamClient
	DriveRootContent                                                         *driverootcontent.DriveRootContentClient
	DriveRootContentStream                                                   *driverootcontentstream.DriveRootContentStreamClient
	DriveRootCreatedByUser                                                   *driverootcreatedbyuser.DriveRootCreatedByUserClient
	DriveRootCreatedByUserMailboxSetting                                     *driverootcreatedbyusermailboxsetting.DriveRootCreatedByUserMailboxSettingClient
	DriveRootCreatedByUserServiceProvisioningError                           *driverootcreatedbyuserserviceprovisioningerror.DriveRootCreatedByUserServiceProvisioningErrorClient
	DriveRootLastModifiedByUser                                              *driverootlastmodifiedbyuser.DriveRootLastModifiedByUserClient
	DriveRootLastModifiedByUserMailboxSetting                                *driverootlastmodifiedbyusermailboxsetting.DriveRootLastModifiedByUserMailboxSettingClient
	DriveRootLastModifiedByUserServiceProvisioningError                      *driverootlastmodifiedbyuserserviceprovisioningerror.DriveRootLastModifiedByUserServiceProvisioningErrorClient
	DriveRootListItem                                                        *driverootlistitem.DriveRootListItemClient
	DriveRootListItemActivity                                                *driverootlistitemactivity.DriveRootListItemActivityClient
	DriveRootListItemActivityDriveItem                                       *driverootlistitemactivitydriveitem.DriveRootListItemActivityDriveItemClient
	DriveRootListItemActivityDriveItemContent                                *driverootlistitemactivitydriveitemcontent.DriveRootListItemActivityDriveItemContentClient
	DriveRootListItemActivityDriveItemContentStream                          *driverootlistitemactivitydriveitemcontentstream.DriveRootListItemActivityDriveItemContentStreamClient
	DriveRootListItemActivityListItem                                        *driverootlistitemactivitylistitem.DriveRootListItemActivityListItemClient
	DriveRootListItemAnalytics                                               *driverootlistitemanalytics.DriveRootListItemAnalyticsClient
	DriveRootListItemCreatedByUser                                           *driverootlistitemcreatedbyuser.DriveRootListItemCreatedByUserClient
	DriveRootListItemCreatedByUserMailboxSetting                             *driverootlistitemcreatedbyusermailboxsetting.DriveRootListItemCreatedByUserMailboxSettingClient
	DriveRootListItemCreatedByUserServiceProvisioningError                   *driverootlistitemcreatedbyuserserviceprovisioningerror.DriveRootListItemCreatedByUserServiceProvisioningErrorClient
	DriveRootListItemDocumentSetVersion                                      *driverootlistitemdocumentsetversion.DriveRootListItemDocumentSetVersionClient
	DriveRootListItemDocumentSetVersionField                                 *driverootlistitemdocumentsetversionfield.DriveRootListItemDocumentSetVersionFieldClient
	DriveRootListItemDriveItem                                               *driverootlistitemdriveitem.DriveRootListItemDriveItemClient
	DriveRootListItemDriveItemContent                                        *driverootlistitemdriveitemcontent.DriveRootListItemDriveItemContentClient
	DriveRootListItemDriveItemContentStream                                  *driverootlistitemdriveitemcontentstream.DriveRootListItemDriveItemContentStreamClient
	DriveRootListItemField                                                   *driverootlistitemfield.DriveRootListItemFieldClient
	DriveRootListItemLastModifiedByUser                                      *driverootlistitemlastmodifiedbyuser.DriveRootListItemLastModifiedByUserClient
	DriveRootListItemLastModifiedByUserMailboxSetting                        *driverootlistitemlastmodifiedbyusermailboxsetting.DriveRootListItemLastModifiedByUserMailboxSettingClient
	DriveRootListItemLastModifiedByUserServiceProvisioningError              *driverootlistitemlastmodifiedbyuserserviceprovisioningerror.DriveRootListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveRootListItemPermission                                              *driverootlistitempermission.DriveRootListItemPermissionClient
	DriveRootListItemVersion                                                 *driverootlistitemversion.DriveRootListItemVersionClient
	DriveRootListItemVersionField                                            *driverootlistitemversionfield.DriveRootListItemVersionFieldClient
	DriveRootPermission                                                      *driverootpermission.DriveRootPermissionClient
	DriveRootRetentionLabel                                                  *driverootretentionlabel.DriveRootRetentionLabelClient
	DriveRootSubscription                                                    *driverootsubscription.DriveRootSubscriptionClient
	DriveRootThumbnail                                                       *driverootthumbnail.DriveRootThumbnailClient
	DriveRootVersion                                                         *driverootversion.DriveRootVersionClient
	DriveRootVersionContent                                                  *driverootversioncontent.DriveRootVersionContentClient
	DriveSpecial                                                             *drivespecial.DriveSpecialClient
	DriveSpecialContent                                                      *drivespecialcontent.DriveSpecialContentClient
	DriveSpecialContentStream                                                *drivespecialcontentstream.DriveSpecialContentStreamClient
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
	InvitedBy                                                                *invitedby.InvitedByClient
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
	ManagedAppLogCollectionRequest                                           *managedapplogcollectionrequest.ManagedAppLogCollectionRequestClient
	ManagedAppRegistration                                                   *managedappregistration.ManagedAppRegistrationClient
	ManagedDevice                                                            *manageddevice.ManagedDeviceClient
	ManagedDeviceAssignmentFilterEvaluationStatusDetail                      *manageddeviceassignmentfilterevaluationstatusdetail.ManagedDeviceAssignmentFilterEvaluationStatusDetailClient
	ManagedDeviceDetectedApp                                                 *manageddevicedetectedapp.ManagedDeviceDetectedAppClient
	ManagedDeviceDeviceCategory                                              *manageddevicedevicecategory.ManagedDeviceDeviceCategoryClient
	ManagedDeviceDeviceCompliancePolicyState                                 *manageddevicedevicecompliancepolicystate.ManagedDeviceDeviceCompliancePolicyStateClient
	ManagedDeviceDeviceConfigurationState                                    *manageddevicedeviceconfigurationstate.ManagedDeviceDeviceConfigurationStateClient
	ManagedDeviceDeviceHealthScriptState                                     *manageddevicedevicehealthscriptstate.ManagedDeviceDeviceHealthScriptStateClient
	ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId *manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid.ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient
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
	OAuth2PermissionGrant                                                    *oauth2permissiongrant.OAuth2PermissionGrantClient
	Onenote                                                                  *onenote.OnenoteClient
	OnenoteNotebook                                                          *onenotenotebook.OnenoteNotebookClient
	OnenoteNotebookSection                                                   *onenotenotebooksection.OnenoteNotebookSectionClient
	OnenoteNotebookSectionGroup                                              *onenotenotebooksectiongroup.OnenoteNotebookSectionGroupClient
	OnenoteNotebookSectionGroupParentNotebook                                *onenotenotebooksectiongroupparentnotebook.OnenoteNotebookSectionGroupParentNotebookClient
	OnenoteNotebookSectionGroupParentSectionGroup                            *onenotenotebooksectiongroupparentsectiongroup.OnenoteNotebookSectionGroupParentSectionGroupClient
	OnenoteNotebookSectionGroupSection                                       *onenotenotebooksectiongroupsection.OnenoteNotebookSectionGroupSectionClient
	OnenoteNotebookSectionGroupSectionGroup                                  *onenotenotebooksectiongroupsectiongroup.OnenoteNotebookSectionGroupSectionGroupClient
	OnenoteNotebookSectionGroupSectionPage                                   *onenotenotebooksectiongroupsectionpage.OnenoteNotebookSectionGroupSectionPageClient
	OnenoteNotebookSectionGroupSectionPageContent                            *onenotenotebooksectiongroupsectionpagecontent.OnenoteNotebookSectionGroupSectionPageContentClient
	OnenoteNotebookSectionGroupSectionPageParentNotebook                     *onenotenotebooksectiongroupsectionpageparentnotebook.OnenoteNotebookSectionGroupSectionPageParentNotebookClient
	OnenoteNotebookSectionGroupSectionPageParentSection                      *onenotenotebooksectiongroupsectionpageparentsection.OnenoteNotebookSectionGroupSectionPageParentSectionClient
	OnenoteNotebookSectionGroupSectionParentNotebook                         *onenotenotebooksectiongroupsectionparentnotebook.OnenoteNotebookSectionGroupSectionParentNotebookClient
	OnenoteNotebookSectionGroupSectionParentSectionGroup                     *onenotenotebooksectiongroupsectionparentsectiongroup.OnenoteNotebookSectionGroupSectionParentSectionGroupClient
	OnenoteNotebookSectionPage                                               *onenotenotebooksectionpage.OnenoteNotebookSectionPageClient
	OnenoteNotebookSectionPageContent                                        *onenotenotebooksectionpagecontent.OnenoteNotebookSectionPageContentClient
	OnenoteNotebookSectionPageParentNotebook                                 *onenotenotebooksectionpageparentnotebook.OnenoteNotebookSectionPageParentNotebookClient
	OnenoteNotebookSectionPageParentSection                                  *onenotenotebooksectionpageparentsection.OnenoteNotebookSectionPageParentSectionClient
	OnenoteNotebookSectionParentNotebook                                     *onenotenotebooksectionparentnotebook.OnenoteNotebookSectionParentNotebookClient
	OnenoteNotebookSectionParentSectionGroup                                 *onenotenotebooksectionparentsectiongroup.OnenoteNotebookSectionParentSectionGroupClient
	OnenoteOperation                                                         *onenoteoperation.OnenoteOperationClient
	OnenotePage                                                              *onenotepage.OnenotePageClient
	OnenotePageContent                                                       *onenotepagecontent.OnenotePageContentClient
	OnenotePageParentNotebook                                                *onenotepageparentnotebook.OnenotePageParentNotebookClient
	OnenotePageParentSection                                                 *onenotepageparentsection.OnenotePageParentSectionClient
	OnenoteResource                                                          *onenoteresource.OnenoteResourceClient
	OnenoteResourceContent                                                   *onenoteresourcecontent.OnenoteResourceContentClient
	OnenoteSection                                                           *onenotesection.OnenoteSectionClient
	OnenoteSectionGroup                                                      *onenotesectiongroup.OnenoteSectionGroupClient
	OnenoteSectionGroupParentNotebook                                        *onenotesectiongroupparentnotebook.OnenoteSectionGroupParentNotebookClient
	OnenoteSectionGroupParentSectionGroup                                    *onenotesectiongroupparentsectiongroup.OnenoteSectionGroupParentSectionGroupClient
	OnenoteSectionGroupSection                                               *onenotesectiongroupsection.OnenoteSectionGroupSectionClient
	OnenoteSectionGroupSectionGroup                                          *onenotesectiongroupsectiongroup.OnenoteSectionGroupSectionGroupClient
	OnenoteSectionGroupSectionPage                                           *onenotesectiongroupsectionpage.OnenoteSectionGroupSectionPageClient
	OnenoteSectionGroupSectionPageContent                                    *onenotesectiongroupsectionpagecontent.OnenoteSectionGroupSectionPageContentClient
	OnenoteSectionGroupSectionPageParentNotebook                             *onenotesectiongroupsectionpageparentnotebook.OnenoteSectionGroupSectionPageParentNotebookClient
	OnenoteSectionGroupSectionPageParentSection                              *onenotesectiongroupsectionpageparentsection.OnenoteSectionGroupSectionPageParentSectionClient
	OnenoteSectionGroupSectionParentNotebook                                 *onenotesectiongroupsectionparentnotebook.OnenoteSectionGroupSectionParentNotebookClient
	OnenoteSectionGroupSectionParentSectionGroup                             *onenotesectiongroupsectionparentsectiongroup.OnenoteSectionGroupSectionParentSectionGroupClient
	OnenoteSectionPage                                                       *onenotesectionpage.OnenoteSectionPageClient
	OnenoteSectionPageContent                                                *onenotesectionpagecontent.OnenoteSectionPageContentClient
	OnenoteSectionPageParentNotebook                                         *onenotesectionpageparentnotebook.OnenoteSectionPageParentNotebookClient
	OnenoteSectionPageParentSection                                          *onenotesectionpageparentsection.OnenoteSectionPageParentSectionClient
	OnenoteSectionParentNotebook                                             *onenotesectionparentnotebook.OnenoteSectionParentNotebookClient
	OnenoteSectionParentSectionGroup                                         *onenotesectionparentsectiongroup.OnenoteSectionParentSectionGroupClient
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
	PendingAccessReviewInstanceDecisionInstanceStageDecision                 *pendingaccessreviewinstancedecisioninstancestagedecision.PendingAccessReviewInstanceDecisionInstanceStageDecisionClient
	PendingAccessReviewInstanceDefinition                                    *pendingaccessreviewinstancedefinition.PendingAccessReviewInstanceDefinitionClient
	PendingAccessReviewInstanceStage                                         *pendingaccessreviewinstancestage.PendingAccessReviewInstanceStageClient
	PendingAccessReviewInstanceStageDecision                                 *pendingaccessreviewinstancestagedecision.PendingAccessReviewInstanceStageDecisionClient
	PendingAccessReviewInstanceStageDecisionInsight                          *pendingaccessreviewinstancestagedecisioninsight.PendingAccessReviewInstanceStageDecisionInsightClient
	PendingAccessReviewInstanceStageDecisionInstance                         *pendingaccessreviewinstancestagedecisioninstance.PendingAccessReviewInstanceStageDecisionInstanceClient
	PendingAccessReviewInstanceStageDecisionInstanceContactedReviewer        *pendingaccessreviewinstancestagedecisioninstancecontactedreviewer.PendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient
	PendingAccessReviewInstanceStageDecisionInstanceDecision                 *pendingaccessreviewinstancestagedecisioninstancedecision.PendingAccessReviewInstanceStageDecisionInstanceDecisionClient
	PendingAccessReviewInstanceStageDecisionInstanceDefinition               *pendingaccessreviewinstancestagedecisioninstancedefinition.PendingAccessReviewInstanceStageDecisionInstanceDefinitionClient
	PermissionGrant                                                          *permissiongrant.PermissionGrantClient
	Person                                                                   *person.PersonClient
	Photo                                                                    *photo.PhotoClient
	Planner                                                                  *planner.PlannerClient
	PlannerAll                                                               *plannerall.PlannerAllClient
	PlannerFavoritePlan                                                      *plannerfavoriteplan.PlannerFavoritePlanClient
	PlannerMyDayTask                                                         *plannermydaytask.PlannerMyDayTaskClient
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
	SettingStorage                                                           *settingstorage.SettingStorageClient
	SettingStorageQuota                                                      *settingstoragequota.SettingStorageQuotaClient
	SettingStorageQuotaService                                               *settingstoragequotaservice.SettingStorageQuotaServiceClient
	SettingWindow                                                            *settingwindow.SettingWindowClient
	SettingWindowInstance                                                    *settingwindowinstance.SettingWindowInstanceClient
	Solution                                                                 *solution.SolutionClient
	SolutionWorkingTimeSchedule                                              *solutionworkingtimeschedule.SolutionWorkingTimeScheduleClient
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

	analyticsActivityStatisticClient, err := analyticsactivitystatistic.NewAnalyticsActivityStatisticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AnalyticsActivityStatistic client: %+v", err)
	}
	configureFunc(analyticsActivityStatisticClient.Client)

	analyticsClient, err := analytics.NewAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Analytics client: %+v", err)
	}
	configureFunc(analyticsClient.Client)

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

	driveActivityClient, err := driveactivity.NewDriveActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveActivity client: %+v", err)
	}
	configureFunc(driveActivityClient.Client)

	driveActivityDriveItemClient, err := driveactivitydriveitem.NewDriveActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveActivityDriveItem client: %+v", err)
	}
	configureFunc(driveActivityDriveItemClient.Client)

	driveActivityDriveItemContentClient, err := driveactivitydriveitemcontent.NewDriveActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveActivityDriveItemContent client: %+v", err)
	}
	configureFunc(driveActivityDriveItemContentClient.Client)

	driveActivityDriveItemContentStreamClient, err := driveactivitydriveitemcontentstream.NewDriveActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveActivityDriveItemContentStreamClient.Client)

	driveActivityListItemClient, err := driveactivitylistitem.NewDriveActivityListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveActivityListItem client: %+v", err)
	}
	configureFunc(driveActivityListItemClient.Client)

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

	driveBundleContentStreamClient, err := drivebundlecontentstream.NewDriveBundleContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveBundleContentStream client: %+v", err)
	}
	configureFunc(driveBundleContentStreamClient.Client)

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

	driveFollowingContentStreamClient, err := drivefollowingcontentstream.NewDriveFollowingContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveFollowingContentStream client: %+v", err)
	}
	configureFunc(driveFollowingContentStreamClient.Client)

	driveItemActivityClient, err := driveitemactivity.NewDriveItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemActivity client: %+v", err)
	}
	configureFunc(driveItemActivityClient.Client)

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

	driveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient, err := driveitemanalyticsitemactivitystatactivitydriveitemcontentstream.NewDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemAnalyticsItemActivityStatActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient.Client)

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

	driveItemChildContentStreamClient, err := driveitemchildcontentstream.NewDriveItemChildContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemChildContentStream client: %+v", err)
	}
	configureFunc(driveItemChildContentStreamClient.Client)

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

	driveItemContentStreamClient, err := driveitemcontentstream.NewDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemContentStream client: %+v", err)
	}
	configureFunc(driveItemContentStreamClient.Client)

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

	driveItemListItemActivityClient, err := driveitemlistitemactivity.NewDriveItemListItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemActivity client: %+v", err)
	}
	configureFunc(driveItemListItemActivityClient.Client)

	driveItemListItemActivityDriveItemClient, err := driveitemlistitemactivitydriveitem.NewDriveItemListItemActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemActivityDriveItem client: %+v", err)
	}
	configureFunc(driveItemListItemActivityDriveItemClient.Client)

	driveItemListItemActivityDriveItemContentClient, err := driveitemlistitemactivitydriveitemcontent.NewDriveItemListItemActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemActivityDriveItemContent client: %+v", err)
	}
	configureFunc(driveItemListItemActivityDriveItemContentClient.Client)

	driveItemListItemActivityDriveItemContentStreamClient, err := driveitemlistitemactivitydriveitemcontentstream.NewDriveItemListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveItemListItemActivityDriveItemContentStreamClient.Client)

	driveItemListItemActivityListItemClient, err := driveitemlistitemactivitylistitem.NewDriveItemListItemActivityListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemActivityListItem client: %+v", err)
	}
	configureFunc(driveItemListItemActivityListItemClient.Client)

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

	driveItemListItemDriveItemContentStreamClient, err := driveitemlistitemdriveitemcontentstream.NewDriveItemListItemDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveItemListItemDriveItemContentStreamClient.Client)

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

	driveItemListItemPermissionClient, err := driveitemlistitempermission.NewDriveItemListItemPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveItemListItemPermission client: %+v", err)
	}
	configureFunc(driveItemListItemPermissionClient.Client)

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

	driveListActivityClient, err := drivelistactivity.NewDriveListActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListActivity client: %+v", err)
	}
	configureFunc(driveListActivityClient.Client)

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

	driveListItemActivityClient, err := drivelistitemactivity.NewDriveListItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemActivity client: %+v", err)
	}
	configureFunc(driveListItemActivityClient.Client)

	driveListItemActivityDriveItemClient, err := drivelistitemactivitydriveitem.NewDriveListItemActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemActivityDriveItem client: %+v", err)
	}
	configureFunc(driveListItemActivityDriveItemClient.Client)

	driveListItemActivityDriveItemContentClient, err := drivelistitemactivitydriveitemcontent.NewDriveListItemActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemActivityDriveItemContent client: %+v", err)
	}
	configureFunc(driveListItemActivityDriveItemContentClient.Client)

	driveListItemActivityDriveItemContentStreamClient, err := drivelistitemactivitydriveitemcontentstream.NewDriveListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveListItemActivityDriveItemContentStreamClient.Client)

	driveListItemActivityListItemClient, err := drivelistitemactivitylistitem.NewDriveListItemActivityListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemActivityListItem client: %+v", err)
	}
	configureFunc(driveListItemActivityListItemClient.Client)

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

	driveListItemDriveItemContentStreamClient, err := drivelistitemdriveitemcontentstream.NewDriveListItemDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveListItemDriveItemContentStreamClient.Client)

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

	driveListItemPermissionClient, err := drivelistitempermission.NewDriveListItemPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListItemPermission client: %+v", err)
	}
	configureFunc(driveListItemPermissionClient.Client)

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

	driveListPermissionClient, err := drivelistpermission.NewDriveListPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListPermission client: %+v", err)
	}
	configureFunc(driveListPermissionClient.Client)

	driveListSubscriptionClient, err := drivelistsubscription.NewDriveListSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveListSubscription client: %+v", err)
	}
	configureFunc(driveListSubscriptionClient.Client)

	driveRootActivityClient, err := driverootactivity.NewDriveRootActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootActivity client: %+v", err)
	}
	configureFunc(driveRootActivityClient.Client)

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

	driveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClient, err := driverootanalyticsitemactivitystatactivitydriveitemcontentstream.NewDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootAnalyticsItemActivityStatActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClient.Client)

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

	driveRootChildContentStreamClient, err := driverootchildcontentstream.NewDriveRootChildContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootChildContentStream client: %+v", err)
	}
	configureFunc(driveRootChildContentStreamClient.Client)

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

	driveRootContentStreamClient, err := driverootcontentstream.NewDriveRootContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootContentStream client: %+v", err)
	}
	configureFunc(driveRootContentStreamClient.Client)

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

	driveRootListItemActivityClient, err := driverootlistitemactivity.NewDriveRootListItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemActivity client: %+v", err)
	}
	configureFunc(driveRootListItemActivityClient.Client)

	driveRootListItemActivityDriveItemClient, err := driverootlistitemactivitydriveitem.NewDriveRootListItemActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemActivityDriveItem client: %+v", err)
	}
	configureFunc(driveRootListItemActivityDriveItemClient.Client)

	driveRootListItemActivityDriveItemContentClient, err := driverootlistitemactivitydriveitemcontent.NewDriveRootListItemActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemActivityDriveItemContent client: %+v", err)
	}
	configureFunc(driveRootListItemActivityDriveItemContentClient.Client)

	driveRootListItemActivityDriveItemContentStreamClient, err := driverootlistitemactivitydriveitemcontentstream.NewDriveRootListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveRootListItemActivityDriveItemContentStreamClient.Client)

	driveRootListItemActivityListItemClient, err := driverootlistitemactivitylistitem.NewDriveRootListItemActivityListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemActivityListItem client: %+v", err)
	}
	configureFunc(driveRootListItemActivityListItemClient.Client)

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

	driveRootListItemDriveItemContentStreamClient, err := driverootlistitemdriveitemcontentstream.NewDriveRootListItemDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemDriveItemContentStream client: %+v", err)
	}
	configureFunc(driveRootListItemDriveItemContentStreamClient.Client)

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

	driveRootListItemPermissionClient, err := driverootlistitempermission.NewDriveRootListItemPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveRootListItemPermission client: %+v", err)
	}
	configureFunc(driveRootListItemPermissionClient.Client)

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

	driveSpecialContentStreamClient, err := drivespecialcontentstream.NewDriveSpecialContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DriveSpecialContentStream client: %+v", err)
	}
	configureFunc(driveSpecialContentStreamClient.Client)

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

	invitedByClient, err := invitedby.NewInvitedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvitedBy client: %+v", err)
	}
	configureFunc(invitedByClient.Client)

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

	managedAppLogCollectionRequestClient, err := managedapplogcollectionrequest.NewManagedAppLogCollectionRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedAppLogCollectionRequest client: %+v", err)
	}
	configureFunc(managedAppLogCollectionRequestClient.Client)

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

	managedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient, err := manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid.NewManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId client: %+v", err)
	}
	configureFunc(managedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient.Client)

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

	oAuth2PermissionGrantClient, err := oauth2permissiongrant.NewOAuth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OAuth2PermissionGrant client: %+v", err)
	}
	configureFunc(oAuth2PermissionGrantClient.Client)

	onenoteClient, err := onenote.NewOnenoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Onenote client: %+v", err)
	}
	configureFunc(onenoteClient.Client)

	onenoteNotebookClient, err := onenotenotebook.NewOnenoteNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebook client: %+v", err)
	}
	configureFunc(onenoteNotebookClient.Client)

	onenoteNotebookSectionClient, err := onenotenotebooksection.NewOnenoteNotebookSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSection client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionClient.Client)

	onenoteNotebookSectionGroupClient, err := onenotenotebooksectiongroup.NewOnenoteNotebookSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroup client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupClient.Client)

	onenoteNotebookSectionGroupParentNotebookClient, err := onenotenotebooksectiongroupparentnotebook.NewOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupParentNotebookClient.Client)

	onenoteNotebookSectionGroupParentSectionGroupClient, err := onenotenotebooksectiongroupparentsectiongroup.NewOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupParentSectionGroupClient.Client)

	onenoteNotebookSectionGroupSectionClient, err := onenotenotebooksectiongroupsection.NewOnenoteNotebookSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSection client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionClient.Client)

	onenoteNotebookSectionGroupSectionGroupClient, err := onenotenotebooksectiongroupsectiongroup.NewOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionGroupClient.Client)

	onenoteNotebookSectionGroupSectionPageClient, err := onenotenotebooksectiongroupsectionpage.NewOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionPageClient.Client)

	onenoteNotebookSectionGroupSectionPageContentClient, err := onenotenotebooksectiongroupsectionpagecontent.NewOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionPageContentClient.Client)

	onenoteNotebookSectionGroupSectionPageParentNotebookClient, err := onenotenotebooksectiongroupsectionpageparentnotebook.NewOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionPageParentNotebookClient.Client)

	onenoteNotebookSectionGroupSectionPageParentSectionClient, err := onenotenotebooksectiongroupsectionpageparentsection.NewOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionPageParentSectionClient.Client)

	onenoteNotebookSectionGroupSectionParentNotebookClient, err := onenotenotebooksectiongroupsectionparentnotebook.NewOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionParentNotebookClient.Client)

	onenoteNotebookSectionGroupSectionParentSectionGroupClient, err := onenotenotebooksectiongroupsectionparentsectiongroup.NewOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionGroupSectionParentSectionGroupClient.Client)

	onenoteNotebookSectionPageClient, err := onenotenotebooksectionpage.NewOnenoteNotebookSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionPage client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionPageClient.Client)

	onenoteNotebookSectionPageContentClient, err := onenotenotebooksectionpagecontent.NewOnenoteNotebookSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionPageContent client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionPageContentClient.Client)

	onenoteNotebookSectionPageParentNotebookClient, err := onenotenotebooksectionpageparentnotebook.NewOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionPageParentNotebookClient.Client)

	onenoteNotebookSectionPageParentSectionClient, err := onenotenotebooksectionpageparentsection.NewOnenoteNotebookSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionPageParentSection client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionPageParentSectionClient.Client)

	onenoteNotebookSectionParentNotebookClient, err := onenotenotebooksectionparentnotebook.NewOnenoteNotebookSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionParentNotebook client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionParentNotebookClient.Client)

	onenoteNotebookSectionParentSectionGroupClient, err := onenotenotebooksectionparentsectiongroup.NewOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteNotebookSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(onenoteNotebookSectionParentSectionGroupClient.Client)

	onenoteOperationClient, err := onenoteoperation.NewOnenoteOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteOperation client: %+v", err)
	}
	configureFunc(onenoteOperationClient.Client)

	onenotePageClient, err := onenotepage.NewOnenotePageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenotePage client: %+v", err)
	}
	configureFunc(onenotePageClient.Client)

	onenotePageContentClient, err := onenotepagecontent.NewOnenotePageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenotePageContent client: %+v", err)
	}
	configureFunc(onenotePageContentClient.Client)

	onenotePageParentNotebookClient, err := onenotepageparentnotebook.NewOnenotePageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenotePageParentNotebook client: %+v", err)
	}
	configureFunc(onenotePageParentNotebookClient.Client)

	onenotePageParentSectionClient, err := onenotepageparentsection.NewOnenotePageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenotePageParentSection client: %+v", err)
	}
	configureFunc(onenotePageParentSectionClient.Client)

	onenoteResourceClient, err := onenoteresource.NewOnenoteResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteResource client: %+v", err)
	}
	configureFunc(onenoteResourceClient.Client)

	onenoteResourceContentClient, err := onenoteresourcecontent.NewOnenoteResourceContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteResourceContent client: %+v", err)
	}
	configureFunc(onenoteResourceContentClient.Client)

	onenoteSectionClient, err := onenotesection.NewOnenoteSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSection client: %+v", err)
	}
	configureFunc(onenoteSectionClient.Client)

	onenoteSectionGroupClient, err := onenotesectiongroup.NewOnenoteSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroup client: %+v", err)
	}
	configureFunc(onenoteSectionGroupClient.Client)

	onenoteSectionGroupParentNotebookClient, err := onenotesectiongroupparentnotebook.NewOnenoteSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(onenoteSectionGroupParentNotebookClient.Client)

	onenoteSectionGroupParentSectionGroupClient, err := onenotesectiongroupparentsectiongroup.NewOnenoteSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(onenoteSectionGroupParentSectionGroupClient.Client)

	onenoteSectionGroupSectionClient, err := onenotesectiongroupsection.NewOnenoteSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSection client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionClient.Client)

	onenoteSectionGroupSectionGroupClient, err := onenotesectiongroupsectiongroup.NewOnenoteSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionGroupClient.Client)

	onenoteSectionGroupSectionPageClient, err := onenotesectiongroupsectionpage.NewOnenoteSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionPageClient.Client)

	onenoteSectionGroupSectionPageContentClient, err := onenotesectiongroupsectionpagecontent.NewOnenoteSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionPageContentClient.Client)

	onenoteSectionGroupSectionPageParentNotebookClient, err := onenotesectiongroupsectionpageparentnotebook.NewOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionPageParentNotebookClient.Client)

	onenoteSectionGroupSectionPageParentSectionClient, err := onenotesectiongroupsectionpageparentsection.NewOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionPageParentSectionClient.Client)

	onenoteSectionGroupSectionParentNotebookClient, err := onenotesectiongroupsectionparentnotebook.NewOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionParentNotebookClient.Client)

	onenoteSectionGroupSectionParentSectionGroupClient, err := onenotesectiongroupsectionparentsectiongroup.NewOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(onenoteSectionGroupSectionParentSectionGroupClient.Client)

	onenoteSectionPageClient, err := onenotesectionpage.NewOnenoteSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionPage client: %+v", err)
	}
	configureFunc(onenoteSectionPageClient.Client)

	onenoteSectionPageContentClient, err := onenotesectionpagecontent.NewOnenoteSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionPageContent client: %+v", err)
	}
	configureFunc(onenoteSectionPageContentClient.Client)

	onenoteSectionPageParentNotebookClient, err := onenotesectionpageparentnotebook.NewOnenoteSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(onenoteSectionPageParentNotebookClient.Client)

	onenoteSectionPageParentSectionClient, err := onenotesectionpageparentsection.NewOnenoteSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionPageParentSection client: %+v", err)
	}
	configureFunc(onenoteSectionPageParentSectionClient.Client)

	onenoteSectionParentNotebookClient, err := onenotesectionparentnotebook.NewOnenoteSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionParentNotebook client: %+v", err)
	}
	configureFunc(onenoteSectionParentNotebookClient.Client)

	onenoteSectionParentSectionGroupClient, err := onenotesectionparentsectiongroup.NewOnenoteSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnenoteSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(onenoteSectionParentSectionGroupClient.Client)

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

	pendingAccessReviewInstanceDecisionInstanceStageDecisionClient, err := pendingaccessreviewinstancedecisioninstancestagedecision.NewPendingAccessReviewInstanceDecisionInstanceStageDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceDecisionInstanceStageDecision client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceDecisionInstanceStageDecisionClient.Client)

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

	pendingAccessReviewInstanceStageDecisionInstanceDecisionClient, err := pendingaccessreviewinstancestagedecisioninstancedecision.NewPendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStageDecisionInstanceDecision client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageDecisionInstanceDecisionClient.Client)

	pendingAccessReviewInstanceStageDecisionInstanceDefinitionClient, err := pendingaccessreviewinstancestagedecisioninstancedefinition.NewPendingAccessReviewInstanceStageDecisionInstanceDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PendingAccessReviewInstanceStageDecisionInstanceDefinition client: %+v", err)
	}
	configureFunc(pendingAccessReviewInstanceStageDecisionInstanceDefinitionClient.Client)

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

	plannerMyDayTaskClient, err := plannermydaytask.NewPlannerMyDayTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlannerMyDayTask client: %+v", err)
	}
	configureFunc(plannerMyDayTaskClient.Client)

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

	solutionClient, err := solution.NewSolutionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Solution client: %+v", err)
	}
	configureFunc(solutionClient.Client)

	solutionWorkingTimeScheduleClient, err := solutionworkingtimeschedule.NewSolutionWorkingTimeScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SolutionWorkingTimeSchedule client: %+v", err)
	}
	configureFunc(solutionWorkingTimeScheduleClient.Client)

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
		Analytics:                     analyticsClient,
		AnalyticsActivityStatistic:    analyticsActivityStatisticClient,
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
		ChatOperation:                              chatOperationClient,
		ChatPermissionGrant:                        chatPermissionGrantClient,
		ChatPinnedMessage:                          chatPinnedMessageClient,
		ChatPinnedMessageMessage:                   chatPinnedMessageMessageClient,
		ChatTab:                                    chatTabClient,
		ChatTabTeamsApp:                            chatTabTeamsAppClient,
		CloudClipboard:                             cloudClipboardClient,
		CloudClipboardItem:                         cloudClipboardItemClient,
		CloudPC:                                    cloudPCClient,
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
		Device:                                     deviceClient,
		DeviceCommand:                              deviceCommandClient,
		DeviceCommandResponsepayload:               deviceCommandResponsepayloadClient,
		DeviceEnrollmentConfiguration:              deviceEnrollmentConfigurationClient,
		DeviceEnrollmentConfigurationAssignment:    deviceEnrollmentConfigurationAssignmentClient,
		DeviceExtension:                            deviceExtensionClient,
		DeviceManagementTroubleshootingEvent:       deviceManagementTroubleshootingEventClient,
		DeviceMemberOf:                             deviceMemberOfClient,
		DeviceRegisteredOwner:                      deviceRegisteredOwnerClient,
		DeviceRegisteredUser:                       deviceRegisteredUserClient,
		DeviceTransitiveMemberOf:                   deviceTransitiveMemberOfClient,
		DeviceUsageRight:                           deviceUsageRightClient,
		DirectReport:                               directReportClient,
		Drive:                                      driveClient,
		DriveActivity:                              driveActivityClient,
		DriveActivityDriveItem:                     driveActivityDriveItemClient,
		DriveActivityDriveItemContent:              driveActivityDriveItemContentClient,
		DriveActivityDriveItemContentStream:        driveActivityDriveItemContentStreamClient,
		DriveActivityListItem:                      driveActivityListItemClient,
		DriveBundle:                                driveBundleClient,
		DriveBundleContent:                         driveBundleContentClient,
		DriveBundleContentStream:                   driveBundleContentStreamClient,
		DriveCreatedByUser:                         driveCreatedByUserClient,
		DriveCreatedByUserMailboxSetting:           driveCreatedByUserMailboxSettingClient,
		DriveCreatedByUserServiceProvisioningError: driveCreatedByUserServiceProvisioningErrorClient,
		DriveFollowing:                             driveFollowingClient,
		DriveFollowingContent:                      driveFollowingContentClient,
		DriveFollowingContentStream:                driveFollowingContentStreamClient,
		DriveItem:                                  driveItemClient,
		DriveItemActivity:                          driveItemActivityClient,
		DriveItemAnalytics:                         driveItemAnalyticsClient,
		DriveItemAnalyticsAllTime:                  driveItemAnalyticsAllTimeClient,
		DriveItemAnalyticsItemActivityStat:         driveItemAnalyticsItemActivityStatClient,
		DriveItemAnalyticsItemActivityStatActivity: driveItemAnalyticsItemActivityStatActivityClient,
		DriveItemAnalyticsItemActivityStatActivityDriveItem:              driveItemAnalyticsItemActivityStatActivityDriveItemClient,
		DriveItemAnalyticsItemActivityStatActivityDriveItemContent:       driveItemAnalyticsItemActivityStatActivityDriveItemContentClient,
		DriveItemAnalyticsItemActivityStatActivityDriveItemContentStream: driveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient,
		DriveItemAnalyticsLastSevenDay:                                   driveItemAnalyticsLastSevenDayClient,
		DriveItemChild:                                                   driveItemChildClient,
		DriveItemChildContent:                                            driveItemChildContentClient,
		DriveItemChildContentStream:                                      driveItemChildContentStreamClient,
		DriveItemContent:                                                 driveItemContentClient,
		DriveItemContentStream:                                           driveItemContentStreamClient,
		DriveItemCreatedByUser:                                           driveItemCreatedByUserClient,
		DriveItemCreatedByUserMailboxSetting:                             driveItemCreatedByUserMailboxSettingClient,
		DriveItemCreatedByUserServiceProvisioningError:                   driveItemCreatedByUserServiceProvisioningErrorClient,
		DriveItemLastModifiedByUser:                                      driveItemLastModifiedByUserClient,
		DriveItemLastModifiedByUserMailboxSetting:                        driveItemLastModifiedByUserMailboxSettingClient,
		DriveItemLastModifiedByUserServiceProvisioningError:              driveItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveItemListItem:                                                driveItemListItemClient,
		DriveItemListItemActivity:                                        driveItemListItemActivityClient,
		DriveItemListItemActivityDriveItem:                               driveItemListItemActivityDriveItemClient,
		DriveItemListItemActivityDriveItemContent:                        driveItemListItemActivityDriveItemContentClient,
		DriveItemListItemActivityDriveItemContentStream:                  driveItemListItemActivityDriveItemContentStreamClient,
		DriveItemListItemActivityListItem:                                driveItemListItemActivityListItemClient,
		DriveItemListItemAnalytics:                                       driveItemListItemAnalyticsClient,
		DriveItemListItemCreatedByUser:                                   driveItemListItemCreatedByUserClient,
		DriveItemListItemCreatedByUserMailboxSetting:                     driveItemListItemCreatedByUserMailboxSettingClient,
		DriveItemListItemCreatedByUserServiceProvisioningError:           driveItemListItemCreatedByUserServiceProvisioningErrorClient,
		DriveItemListItemDocumentSetVersion:                              driveItemListItemDocumentSetVersionClient,
		DriveItemListItemDocumentSetVersionField:                         driveItemListItemDocumentSetVersionFieldClient,
		DriveItemListItemDriveItem:                                       driveItemListItemDriveItemClient,
		DriveItemListItemDriveItemContent:                                driveItemListItemDriveItemContentClient,
		DriveItemListItemDriveItemContentStream:                          driveItemListItemDriveItemContentStreamClient,
		DriveItemListItemField:                                           driveItemListItemFieldClient,
		DriveItemListItemLastModifiedByUser:                              driveItemListItemLastModifiedByUserClient,
		DriveItemListItemLastModifiedByUserMailboxSetting:                driveItemListItemLastModifiedByUserMailboxSettingClient,
		DriveItemListItemLastModifiedByUserServiceProvisioningError:      driveItemListItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveItemListItemPermission:                                      driveItemListItemPermissionClient,
		DriveItemListItemVersion:                                         driveItemListItemVersionClient,
		DriveItemListItemVersionField:                                    driveItemListItemVersionFieldClient,
		DriveItemPermission:                                              driveItemPermissionClient,
		DriveItemRetentionLabel:                                          driveItemRetentionLabelClient,
		DriveItemSubscription:                                            driveItemSubscriptionClient,
		DriveItemThumbnail:                                               driveItemThumbnailClient,
		DriveItemVersion:                                                 driveItemVersionClient,
		DriveItemVersionContent:                                          driveItemVersionContentClient,
		DriveLastModifiedByUser:                                          driveLastModifiedByUserClient,
		DriveLastModifiedByUserMailboxSetting:                            driveLastModifiedByUserMailboxSettingClient,
		DriveLastModifiedByUserServiceProvisioningError:                  driveLastModifiedByUserServiceProvisioningErrorClient,
		DriveList:                                                        driveListClient,
		DriveListActivity:                                                driveListActivityClient,
		DriveListColumn:                                                  driveListColumnClient,
		DriveListColumnSourceColumn:                                      driveListColumnSourceColumnClient,
		DriveListContentType:                                             driveListContentTypeClient,
		DriveListContentTypeBase:                                         driveListContentTypeBaseClient,
		DriveListContentTypeBaseType:                                     driveListContentTypeBaseTypeClient,
		DriveListContentTypeColumn:                                       driveListContentTypeColumnClient,
		DriveListContentTypeColumnLink:                                   driveListContentTypeColumnLinkClient,
		DriveListContentTypeColumnPosition:                               driveListContentTypeColumnPositionClient,
		DriveListContentTypeColumnSourceColumn:                           driveListContentTypeColumnSourceColumnClient,
		DriveListCreatedByUser:                                           driveListCreatedByUserClient,
		DriveListCreatedByUserMailboxSetting:                             driveListCreatedByUserMailboxSettingClient,
		DriveListCreatedByUserServiceProvisioningError:                   driveListCreatedByUserServiceProvisioningErrorClient,
		DriveListDrive:                                                   driveListDriveClient,
		DriveListItem:                                                    driveListItemClient,
		DriveListItemActivity:                                            driveListItemActivityClient,
		DriveListItemActivityDriveItem:                                   driveListItemActivityDriveItemClient,
		DriveListItemActivityDriveItemContent:                            driveListItemActivityDriveItemContentClient,
		DriveListItemActivityDriveItemContentStream:                      driveListItemActivityDriveItemContentStreamClient,
		DriveListItemActivityListItem:                                    driveListItemActivityListItemClient,
		DriveListItemAnalytics:                                           driveListItemAnalyticsClient,
		DriveListItemCreatedByUser:                                       driveListItemCreatedByUserClient,
		DriveListItemCreatedByUserMailboxSetting:                         driveListItemCreatedByUserMailboxSettingClient,
		DriveListItemCreatedByUserServiceProvisioningError:               driveListItemCreatedByUserServiceProvisioningErrorClient,
		DriveListItemDocumentSetVersion:                                  driveListItemDocumentSetVersionClient,
		DriveListItemDocumentSetVersionField:                             driveListItemDocumentSetVersionFieldClient,
		DriveListItemDriveItem:                                           driveListItemDriveItemClient,
		DriveListItemDriveItemContent:                                    driveListItemDriveItemContentClient,
		DriveListItemDriveItemContentStream:                              driveListItemDriveItemContentStreamClient,
		DriveListItemField:                                               driveListItemFieldClient,
		DriveListItemLastModifiedByUser:                                  driveListItemLastModifiedByUserClient,
		DriveListItemLastModifiedByUserMailboxSetting:                    driveListItemLastModifiedByUserMailboxSettingClient,
		DriveListItemLastModifiedByUserServiceProvisioningError:          driveListItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveListItemPermission:                                          driveListItemPermissionClient,
		DriveListItemVersion:                                             driveListItemVersionClient,
		DriveListItemVersionField:                                        driveListItemVersionFieldClient,
		DriveListLastModifiedByUser:                                      driveListLastModifiedByUserClient,
		DriveListLastModifiedByUserMailboxSetting:                        driveListLastModifiedByUserMailboxSettingClient,
		DriveListLastModifiedByUserServiceProvisioningError:              driveListLastModifiedByUserServiceProvisioningErrorClient,
		DriveListOperation:                                               driveListOperationClient,
		DriveListPermission:                                              driveListPermissionClient,
		DriveListSubscription:                                            driveListSubscriptionClient,
		DriveRoot:                                                        driveRootClient,
		DriveRootActivity:                                                driveRootActivityClient,
		DriveRootAnalytics:                                               driveRootAnalyticsClient,
		DriveRootAnalyticsAllTime:                                        driveRootAnalyticsAllTimeClient,
		DriveRootAnalyticsItemActivityStat:                               driveRootAnalyticsItemActivityStatClient,
		DriveRootAnalyticsItemActivityStatActivity:                       driveRootAnalyticsItemActivityStatActivityClient,
		DriveRootAnalyticsItemActivityStatActivityDriveItem:              driveRootAnalyticsItemActivityStatActivityDriveItemClient,
		DriveRootAnalyticsItemActivityStatActivityDriveItemContent:       driveRootAnalyticsItemActivityStatActivityDriveItemContentClient,
		DriveRootAnalyticsItemActivityStatActivityDriveItemContentStream: driveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClient,
		DriveRootAnalyticsLastSevenDay:                                   driveRootAnalyticsLastSevenDayClient,
		DriveRootChild:                                                   driveRootChildClient,
		DriveRootChildContent:                                            driveRootChildContentClient,
		DriveRootChildContentStream:                                      driveRootChildContentStreamClient,
		DriveRootContent:                                                 driveRootContentClient,
		DriveRootContentStream:                                           driveRootContentStreamClient,
		DriveRootCreatedByUser:                                           driveRootCreatedByUserClient,
		DriveRootCreatedByUserMailboxSetting:                             driveRootCreatedByUserMailboxSettingClient,
		DriveRootCreatedByUserServiceProvisioningError:                   driveRootCreatedByUserServiceProvisioningErrorClient,
		DriveRootLastModifiedByUser:                                      driveRootLastModifiedByUserClient,
		DriveRootLastModifiedByUserMailboxSetting:                        driveRootLastModifiedByUserMailboxSettingClient,
		DriveRootLastModifiedByUserServiceProvisioningError:              driveRootLastModifiedByUserServiceProvisioningErrorClient,
		DriveRootListItem:                                                driveRootListItemClient,
		DriveRootListItemActivity:                                        driveRootListItemActivityClient,
		DriveRootListItemActivityDriveItem:                               driveRootListItemActivityDriveItemClient,
		DriveRootListItemActivityDriveItemContent:                        driveRootListItemActivityDriveItemContentClient,
		DriveRootListItemActivityDriveItemContentStream:                  driveRootListItemActivityDriveItemContentStreamClient,
		DriveRootListItemActivityListItem:                                driveRootListItemActivityListItemClient,
		DriveRootListItemAnalytics:                                       driveRootListItemAnalyticsClient,
		DriveRootListItemCreatedByUser:                                   driveRootListItemCreatedByUserClient,
		DriveRootListItemCreatedByUserMailboxSetting:                     driveRootListItemCreatedByUserMailboxSettingClient,
		DriveRootListItemCreatedByUserServiceProvisioningError:           driveRootListItemCreatedByUserServiceProvisioningErrorClient,
		DriveRootListItemDocumentSetVersion:                              driveRootListItemDocumentSetVersionClient,
		DriveRootListItemDocumentSetVersionField:                         driveRootListItemDocumentSetVersionFieldClient,
		DriveRootListItemDriveItem:                                       driveRootListItemDriveItemClient,
		DriveRootListItemDriveItemContent:                                driveRootListItemDriveItemContentClient,
		DriveRootListItemDriveItemContentStream:                          driveRootListItemDriveItemContentStreamClient,
		DriveRootListItemField:                                           driveRootListItemFieldClient,
		DriveRootListItemLastModifiedByUser:                              driveRootListItemLastModifiedByUserClient,
		DriveRootListItemLastModifiedByUserMailboxSetting:                driveRootListItemLastModifiedByUserMailboxSettingClient,
		DriveRootListItemLastModifiedByUserServiceProvisioningError:      driveRootListItemLastModifiedByUserServiceProvisioningErrorClient,
		DriveRootListItemPermission:                                      driveRootListItemPermissionClient,
		DriveRootListItemVersion:                                         driveRootListItemVersionClient,
		DriveRootListItemVersionField:                                    driveRootListItemVersionFieldClient,
		DriveRootPermission:                                              driveRootPermissionClient,
		DriveRootRetentionLabel:                                          driveRootRetentionLabelClient,
		DriveRootSubscription:                                            driveRootSubscriptionClient,
		DriveRootThumbnail:                                               driveRootThumbnailClient,
		DriveRootVersion:                                                 driveRootVersionClient,
		DriveRootVersionContent:                                          driveRootVersionContentClient,
		DriveSpecial:                                                     driveSpecialClient,
		DriveSpecialContent:                                              driveSpecialContentClient,
		DriveSpecialContentStream:                                        driveSpecialContentStreamClient,
		EmployeeExperience:                                               employeeExperienceClient,
		EmployeeExperienceLearningCourseActivity:                         employeeExperienceLearningCourseActivityClient,
		Event:                                                            eventClient,
		EventAttachment:                                                  eventAttachmentClient,
		EventCalendar:                                                    eventCalendarClient,
		EventExceptionOccurrence:                                         eventExceptionOccurrenceClient,
		EventExceptionOccurrenceAttachment:                               eventExceptionOccurrenceAttachmentClient,
		EventExceptionOccurrenceCalendar:                                 eventExceptionOccurrenceCalendarClient,
		EventExceptionOccurrenceExtension:                                eventExceptionOccurrenceExtensionClient,
		EventExceptionOccurrenceInstance:                                 eventExceptionOccurrenceInstanceClient,
		EventExceptionOccurrenceInstanceAttachment:                       eventExceptionOccurrenceInstanceAttachmentClient,
		EventExceptionOccurrenceInstanceCalendar:                         eventExceptionOccurrenceInstanceCalendarClient,
		EventExceptionOccurrenceInstanceExtension:                        eventExceptionOccurrenceInstanceExtensionClient,
		EventExtension:                                                   eventExtensionClient,
		EventInstance:                                                    eventInstanceClient,
		EventInstanceAttachment:                                          eventInstanceAttachmentClient,
		EventInstanceCalendar:                                            eventInstanceCalendarClient,
		EventInstanceExceptionOccurrence:                                 eventInstanceExceptionOccurrenceClient,
		EventInstanceExceptionOccurrenceAttachment:                       eventInstanceExceptionOccurrenceAttachmentClient,
		EventInstanceExceptionOccurrenceCalendar:                         eventInstanceExceptionOccurrenceCalendarClient,
		EventInstanceExceptionOccurrenceExtension:                        eventInstanceExceptionOccurrenceExtensionClient,
		EventInstanceExtension:                                           eventInstanceExtensionClient,
		Extension:                                                        extensionClient,
		FollowedSite:                                                     followedSiteClient,
		InferenceClassification:                                          inferenceClassificationClient,
		InferenceClassificationOverride:                                  inferenceClassificationOverrideClient,
		InformationProtection:                                            informationProtectionClient,
		InformationProtectionBitlocker:                                   informationProtectionBitlockerClient,
		InformationProtectionBitlockerRecoveryKey:                        informationProtectionBitlockerRecoveryKeyClient,
		InformationProtectionDataLossPreventionPolicy:                    informationProtectionDataLossPreventionPolicyClient,
		InformationProtectionPolicy:                                      informationProtectionPolicyClient,
		InformationProtectionPolicyLabel:                                 informationProtectionPolicyLabelClient,
		InformationProtectionSensitivityLabel:                            informationProtectionSensitivityLabelClient,
		InformationProtectionSensitivityLabelSublabel:                    informationProtectionSensitivityLabelSublabelClient,
		InformationProtectionSensitivityPolicySetting:                    informationProtectionSensitivityPolicySettingClient,
		InformationProtectionThreatAssessmentRequest:                     informationProtectionThreatAssessmentRequestClient,
		InformationProtectionThreatAssessmentRequestResult:               informationProtectionThreatAssessmentRequestResultClient,
		Insight:                                insightClient,
		InsightShared:                          insightSharedClient,
		InsightSharedLastSharedMethod:          insightSharedLastSharedMethodClient,
		InsightSharedResource:                  insightSharedResourceClient,
		InsightTrending:                        insightTrendingClient,
		InsightTrendingResource:                insightTrendingResourceClient,
		InsightUsed:                            insightUsedClient,
		InsightUsedResource:                    insightUsedResourceClient,
		InvitedBy:                              invitedByClient,
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
		ManagedAppLogCollectionRequest:         managedAppLogCollectionRequestClient,
		ManagedAppRegistration:                 managedAppRegistrationClient,
		ManagedDevice:                          managedDeviceClient,
		ManagedDeviceAssignmentFilterEvaluationStatusDetail:                      managedDeviceAssignmentFilterEvaluationStatusDetailClient,
		ManagedDeviceDetectedApp:                                                 managedDeviceDetectedAppClient,
		ManagedDeviceDeviceCategory:                                              managedDeviceDeviceCategoryClient,
		ManagedDeviceDeviceCompliancePolicyState:                                 managedDeviceDeviceCompliancePolicyStateClient,
		ManagedDeviceDeviceConfigurationState:                                    managedDeviceDeviceConfigurationStateClient,
		ManagedDeviceDeviceHealthScriptState:                                     managedDeviceDeviceHealthScriptStateClient,
		ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceId: managedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient,
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
		Notification:                notificationClient,
		OAuth2PermissionGrant:       oAuth2PermissionGrantClient,
		Onenote:                     onenoteClient,
		OnenoteNotebook:             onenoteNotebookClient,
		OnenoteNotebookSection:      onenoteNotebookSectionClient,
		OnenoteNotebookSectionGroup: onenoteNotebookSectionGroupClient,
		OnenoteNotebookSectionGroupParentNotebook:                         onenoteNotebookSectionGroupParentNotebookClient,
		OnenoteNotebookSectionGroupParentSectionGroup:                     onenoteNotebookSectionGroupParentSectionGroupClient,
		OnenoteNotebookSectionGroupSection:                                onenoteNotebookSectionGroupSectionClient,
		OnenoteNotebookSectionGroupSectionGroup:                           onenoteNotebookSectionGroupSectionGroupClient,
		OnenoteNotebookSectionGroupSectionPage:                            onenoteNotebookSectionGroupSectionPageClient,
		OnenoteNotebookSectionGroupSectionPageContent:                     onenoteNotebookSectionGroupSectionPageContentClient,
		OnenoteNotebookSectionGroupSectionPageParentNotebook:              onenoteNotebookSectionGroupSectionPageParentNotebookClient,
		OnenoteNotebookSectionGroupSectionPageParentSection:               onenoteNotebookSectionGroupSectionPageParentSectionClient,
		OnenoteNotebookSectionGroupSectionParentNotebook:                  onenoteNotebookSectionGroupSectionParentNotebookClient,
		OnenoteNotebookSectionGroupSectionParentSectionGroup:              onenoteNotebookSectionGroupSectionParentSectionGroupClient,
		OnenoteNotebookSectionPage:                                        onenoteNotebookSectionPageClient,
		OnenoteNotebookSectionPageContent:                                 onenoteNotebookSectionPageContentClient,
		OnenoteNotebookSectionPageParentNotebook:                          onenoteNotebookSectionPageParentNotebookClient,
		OnenoteNotebookSectionPageParentSection:                           onenoteNotebookSectionPageParentSectionClient,
		OnenoteNotebookSectionParentNotebook:                              onenoteNotebookSectionParentNotebookClient,
		OnenoteNotebookSectionParentSectionGroup:                          onenoteNotebookSectionParentSectionGroupClient,
		OnenoteOperation:                                                  onenoteOperationClient,
		OnenotePage:                                                       onenotePageClient,
		OnenotePageContent:                                                onenotePageContentClient,
		OnenotePageParentNotebook:                                         onenotePageParentNotebookClient,
		OnenotePageParentSection:                                          onenotePageParentSectionClient,
		OnenoteResource:                                                   onenoteResourceClient,
		OnenoteResourceContent:                                            onenoteResourceContentClient,
		OnenoteSection:                                                    onenoteSectionClient,
		OnenoteSectionGroup:                                               onenoteSectionGroupClient,
		OnenoteSectionGroupParentNotebook:                                 onenoteSectionGroupParentNotebookClient,
		OnenoteSectionGroupParentSectionGroup:                             onenoteSectionGroupParentSectionGroupClient,
		OnenoteSectionGroupSection:                                        onenoteSectionGroupSectionClient,
		OnenoteSectionGroupSectionGroup:                                   onenoteSectionGroupSectionGroupClient,
		OnenoteSectionGroupSectionPage:                                    onenoteSectionGroupSectionPageClient,
		OnenoteSectionGroupSectionPageContent:                             onenoteSectionGroupSectionPageContentClient,
		OnenoteSectionGroupSectionPageParentNotebook:                      onenoteSectionGroupSectionPageParentNotebookClient,
		OnenoteSectionGroupSectionPageParentSection:                       onenoteSectionGroupSectionPageParentSectionClient,
		OnenoteSectionGroupSectionParentNotebook:                          onenoteSectionGroupSectionParentNotebookClient,
		OnenoteSectionGroupSectionParentSectionGroup:                      onenoteSectionGroupSectionParentSectionGroupClient,
		OnenoteSectionPage:                                                onenoteSectionPageClient,
		OnenoteSectionPageContent:                                         onenoteSectionPageContentClient,
		OnenoteSectionPageParentNotebook:                                  onenoteSectionPageParentNotebookClient,
		OnenoteSectionPageParentSection:                                   onenoteSectionPageParentSectionClient,
		OnenoteSectionParentNotebook:                                      onenoteSectionParentNotebookClient,
		OnenoteSectionParentSectionGroup:                                  onenoteSectionParentSectionGroupClient,
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
		PendingAccessReviewInstanceDecisionInstanceStageDecision:          pendingAccessReviewInstanceDecisionInstanceStageDecisionClient,
		PendingAccessReviewInstanceDefinition:                             pendingAccessReviewInstanceDefinitionClient,
		PendingAccessReviewInstanceStage:                                  pendingAccessReviewInstanceStageClient,
		PendingAccessReviewInstanceStageDecision:                          pendingAccessReviewInstanceStageDecisionClient,
		PendingAccessReviewInstanceStageDecisionInsight:                   pendingAccessReviewInstanceStageDecisionInsightClient,
		PendingAccessReviewInstanceStageDecisionInstance:                  pendingAccessReviewInstanceStageDecisionInstanceClient,
		PendingAccessReviewInstanceStageDecisionInstanceContactedReviewer: pendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient,
		PendingAccessReviewInstanceStageDecisionInstanceDecision:          pendingAccessReviewInstanceStageDecisionInstanceDecisionClient,
		PendingAccessReviewInstanceStageDecisionInstanceDefinition:        pendingAccessReviewInstanceStageDecisionInstanceDefinitionClient,
		PermissionGrant:                                                   permissionGrantClient,
		Person:                                                            personClient,
		Photo:                                                             photoClient,
		Planner:                                                           plannerClient,
		PlannerAll:                                                        plannerAllClient,
		PlannerFavoritePlan:                                               plannerFavoritePlanClient,
		PlannerMyDayTask:                                                  plannerMyDayTaskClient,
		PlannerPlan:                                                       plannerPlanClient,
		PlannerPlanBucket:                                                 plannerPlanBucketClient,
		PlannerPlanBucketTask:                                             plannerPlanBucketTaskClient,
		PlannerPlanBucketTaskAssignedToTaskBoardFormat:                    plannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		PlannerPlanBucketTaskBucketTaskBoardFormat:                        plannerPlanBucketTaskBucketTaskBoardFormatClient,
		PlannerPlanBucketTaskDetail:                                       plannerPlanBucketTaskDetailClient,
		PlannerPlanBucketTaskProgressTaskBoardFormat:                      plannerPlanBucketTaskProgressTaskBoardFormatClient,
		PlannerPlanDetail:                                                 plannerPlanDetailClient,
		PlannerPlanTask:                                                   plannerPlanTaskClient,
		PlannerPlanTaskAssignedToTaskBoardFormat:                          plannerPlanTaskAssignedToTaskBoardFormatClient,
		PlannerPlanTaskBucketTaskBoardFormat:                              plannerPlanTaskBucketTaskBoardFormatClient,
		PlannerPlanTaskDetail:                                             plannerPlanTaskDetailClient,
		PlannerPlanTaskProgressTaskBoardFormat:                            plannerPlanTaskProgressTaskBoardFormatClient,
		PlannerRecentPlan:                                                 plannerRecentPlanClient,
		PlannerRosterPlan:                                                 plannerRosterPlanClient,
		PlannerTask:                                                       plannerTaskClient,
		PlannerTaskAssignedToTaskBoardFormat:                              plannerTaskAssignedToTaskBoardFormatClient,
		PlannerTaskBucketTaskBoardFormat:                                  plannerTaskBucketTaskBoardFormatClient,
		PlannerTaskDetail:                                                 plannerTaskDetailClient,
		PlannerTaskProgressTaskBoardFormat:                                plannerTaskProgressTaskBoardFormatClient,
		Presence:                                                          presenceClient,
		Profile:                                                           profileClient,
		ProfileAccount:                                                    profileAccountClient,
		ProfileAddress:                                                    profileAddressClient,
		ProfileAnniversary:                                                profileAnniversaryClient,
		ProfileAward:                                                      profileAwardClient,
		ProfileCertification:                                              profileCertificationClient,
		ProfileEducationalActivity:                                        profileEducationalActivityClient,
		ProfileEmail:                                                      profileEmailClient,
		ProfileInterest:                                                   profileInterestClient,
		ProfileLanguage:                                                   profileLanguageClient,
		ProfileName:                                                       profileNameClient,
		ProfileNote:                                                       profileNoteClient,
		ProfilePatent:                                                     profilePatentClient,
		ProfilePhone:                                                      profilePhoneClient,
		ProfilePosition:                                                   profilePositionClient,
		ProfileProject:                                                    profileProjectClient,
		ProfilePublication:                                                profilePublicationClient,
		ProfileSkill:                                                      profileSkillClient,
		ProfileWebAccount:                                                 profileWebAccountClient,
		ProfileWebsite:                                                    profileWebsiteClient,
		RegisteredDevice:                                                  registeredDeviceClient,
		ScopedRoleMemberOf:                                                scopedRoleMemberOfClient,
		Security:                                                          securityClient,
		SecurityInformationProtection:                                     securityInformationProtectionClient,
		SecurityInformationProtectionLabelPolicySetting:                   securityInformationProtectionLabelPolicySettingClient,
		SecurityInformationProtectionSensitivityLabel:                     securityInformationProtectionSensitivityLabelClient,
		SecurityInformationProtectionSensitivityLabelParent:               securityInformationProtectionSensitivityLabelParentClient,
		ServiceProvisioningError:                                          serviceProvisioningErrorClient,
		Setting:                                                           settingClient,
		SettingContactMergeSuggestion:                                     settingContactMergeSuggestionClient,
		SettingItemInsight:                                                settingItemInsightClient,
		SettingRegionalAndLanguageSetting:                                 settingRegionalAndLanguageSettingClient,
		SettingShiftPreference:                                            settingShiftPreferenceClient,
		SettingStorage:                                                    settingStorageClient,
		SettingStorageQuota:                                               settingStorageQuotaClient,
		SettingStorageQuotaService:                                        settingStorageQuotaServiceClient,
		SettingWindow:                                                     settingWindowClient,
		SettingWindowInstance:                                             settingWindowInstanceClient,
		Solution:                                                          solutionClient,
		SolutionWorkingTimeSchedule:                                       solutionWorkingTimeScheduleClient,
		Sponsor:                                                           sponsorClient,
		Teamwork:                                                          teamworkClient,
		TeamworkAssociatedTeam:                                            teamworkAssociatedTeamClient,
		TeamworkAssociatedTeamTeam:                                        teamworkAssociatedTeamTeamClient,
		TeamworkInstalledApp:                                              teamworkInstalledAppClient,
		TeamworkInstalledAppChat:                                          teamworkInstalledAppChatClient,
		TeamworkInstalledAppTeamsApp:                                      teamworkInstalledAppTeamsAppClient,
		TeamworkInstalledAppTeamsAppDefinition:                            teamworkInstalledAppTeamsAppDefinitionClient,
		Todo:                                                              todoClient,
		TodoList:                                                          todoListClient,
		TodoListExtension:                                                 todoListExtensionClient,
		TodoListTask:                                                      todoListTaskClient,
		TodoListTaskAttachment:                                            todoListTaskAttachmentClient,
		TodoListTaskAttachmentSession:                                     todoListTaskAttachmentSessionClient,
		TodoListTaskAttachmentSessionContent:                              todoListTaskAttachmentSessionContentClient,
		TodoListTaskChecklistItem:                                         todoListTaskChecklistItemClient,
		TodoListTaskExtension:                                             todoListTaskExtensionClient,
		TodoListTaskLinkedResource:                                        todoListTaskLinkedResourceClient,
		TransitiveMemberOf:                                                transitiveMemberOfClient,
		TransitiveReport:                                                  transitiveReportClient,
		UsageRight:                                                        usageRightClient,
		User:                                                              userClient,
		VirtualEvent:                                                      virtualEventClient,
		VirtualEventWebinar:                                               virtualEventWebinarClient,
		WindowsInformationProtectionDeviceRegistration:                    windowsInformationProtectionDeviceRegistrationClient,
	}, nil
}
