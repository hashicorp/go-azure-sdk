package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/acceptedsender"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/approleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/calendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthread"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpost"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostinreplyto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostinreplytoattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/conversationthreadpostinreplytoextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/createdonbehalfof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivebundle"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivebundlecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivefollowing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivefollowingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveitemversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivelistsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driveroot"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/driverootversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivespecial"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/drivespecialcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/event"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/eventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/extension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/group"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/grouplifecyclepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/member"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/memberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/memberswithlicenseerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/photo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/planner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/plannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/rejectedsender"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/serviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/setting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/site"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitedrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteexternalcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepagecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepagecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepagecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepagelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepagelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepagelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitepermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/team"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamallchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamchanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamgroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamincomingchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teaminstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teaminstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teaminstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teammember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teampermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamprimarychanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduleoffershiftrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduleopenshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduleopenshiftchangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduleschedulinggroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduleshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduleswapshiftschangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduletimeoffreason"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduletimeoffrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduletimesoff"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamtag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamtagmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/thread"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpost"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpostattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpostextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpostinreplyto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpostinreplytoattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/threadpostinreplytoextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/transitivemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/transitivememberof"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AcceptedSender                                              *acceptedsender.AcceptedSenderClient
	AppRoleAssignment                                           *approleassignment.AppRoleAssignmentClient
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
	CalendarView                                                *calendarview.CalendarViewClient
	CalendarViewAttachment                                      *calendarviewattachment.CalendarViewAttachmentClient
	CalendarViewCalendar                                        *calendarviewcalendar.CalendarViewCalendarClient
	CalendarViewExtension                                       *calendarviewextension.CalendarViewExtensionClient
	CalendarViewInstance                                        *calendarviewinstance.CalendarViewInstanceClient
	CalendarViewInstanceAttachment                              *calendarviewinstanceattachment.CalendarViewInstanceAttachmentClient
	CalendarViewInstanceCalendar                                *calendarviewinstancecalendar.CalendarViewInstanceCalendarClient
	CalendarViewInstanceExtension                               *calendarviewinstanceextension.CalendarViewInstanceExtensionClient
	Conversation                                                *conversation.ConversationClient
	ConversationThread                                          *conversationthread.ConversationThreadClient
	ConversationThreadPost                                      *conversationthreadpost.ConversationThreadPostClient
	ConversationThreadPostAttachment                            *conversationthreadpostattachment.ConversationThreadPostAttachmentClient
	ConversationThreadPostExtension                             *conversationthreadpostextension.ConversationThreadPostExtensionClient
	ConversationThreadPostInReplyTo                             *conversationthreadpostinreplyto.ConversationThreadPostInReplyToClient
	ConversationThreadPostInReplyToAttachment                   *conversationthreadpostinreplytoattachment.ConversationThreadPostInReplyToAttachmentClient
	ConversationThreadPostInReplyToExtension                    *conversationthreadpostinreplytoextension.ConversationThreadPostInReplyToExtensionClient
	CreatedOnBehalfOf                                           *createdonbehalfof.CreatedOnBehalfOfClient
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
	Event                                                       *event.EventClient
	EventAttachment                                             *eventattachment.EventAttachmentClient
	EventCalendar                                               *eventcalendar.EventCalendarClient
	EventExtension                                              *eventextension.EventExtensionClient
	EventInstance                                               *eventinstance.EventInstanceClient
	EventInstanceAttachment                                     *eventinstanceattachment.EventInstanceAttachmentClient
	EventInstanceCalendar                                       *eventinstancecalendar.EventInstanceCalendarClient
	EventInstanceExtension                                      *eventinstanceextension.EventInstanceExtensionClient
	Extension                                                   *extension.ExtensionClient
	Group                                                       *group.GroupClient
	GroupLifecyclePolicy                                        *grouplifecyclepolicy.GroupLifecyclePolicyClient
	Member                                                      *member.MemberClient
	MemberOf                                                    *memberof.MemberOfClient
	MembersWithLicenseError                                     *memberswithlicenseerror.MembersWithLicenseErrorClient
	Owner                                                       *owner.OwnerClient
	PermissionGrant                                             *permissiongrant.PermissionGrantClient
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
	RejectedSender                                              *rejectedsender.RejectedSenderClient
	ServiceProvisioningError                                    *serviceprovisioningerror.ServiceProvisioningErrorClient
	Setting                                                     *setting.SettingClient
	Site                                                        *site.SiteClient
	SiteAnalytics                                               *siteanalytics.SiteAnalyticsClient
	SiteAnalyticsAllTime                                        *siteanalyticsalltime.SiteAnalyticsAllTimeClient
	SiteAnalyticsItemActivityStat                               *siteanalyticsitemactivitystat.SiteAnalyticsItemActivityStatClient
	SiteAnalyticsItemActivityStatActivity                       *siteanalyticsitemactivitystatactivity.SiteAnalyticsItemActivityStatActivityClient
	SiteAnalyticsItemActivityStatActivityDriveItem              *siteanalyticsitemactivitystatactivitydriveitem.SiteAnalyticsItemActivityStatActivityDriveItemClient
	SiteAnalyticsItemActivityStatActivityDriveItemContent       *siteanalyticsitemactivitystatactivitydriveitemcontent.SiteAnalyticsItemActivityStatActivityDriveItemContentClient
	SiteAnalyticsLastSevenDay                                   *siteanalyticslastsevenday.SiteAnalyticsLastSevenDayClient
	SiteColumn                                                  *sitecolumn.SiteColumnClient
	SiteColumnSourceColumn                                      *sitecolumnsourcecolumn.SiteColumnSourceColumnClient
	SiteContentType                                             *sitecontenttype.SiteContentTypeClient
	SiteContentTypeBase                                         *sitecontenttypebase.SiteContentTypeBaseClient
	SiteContentTypeBaseType                                     *sitecontenttypebasetype.SiteContentTypeBaseTypeClient
	SiteContentTypeColumn                                       *sitecontenttypecolumn.SiteContentTypeColumnClient
	SiteContentTypeColumnLink                                   *sitecontenttypecolumnlink.SiteContentTypeColumnLinkClient
	SiteContentTypeColumnPosition                               *sitecontenttypecolumnposition.SiteContentTypeColumnPositionClient
	SiteContentTypeColumnSourceColumn                           *sitecontenttypecolumnsourcecolumn.SiteContentTypeColumnSourceColumnClient
	SiteCreatedByUser                                           *sitecreatedbyuser.SiteCreatedByUserClient
	SiteCreatedByUserMailboxSetting                             *sitecreatedbyusermailboxsetting.SiteCreatedByUserMailboxSettingClient
	SiteCreatedByUserServiceProvisioningError                   *sitecreatedbyuserserviceprovisioningerror.SiteCreatedByUserServiceProvisioningErrorClient
	SiteDrive                                                   *sitedrive.SiteDriveClient
	SiteExternalColumn                                          *siteexternalcolumn.SiteExternalColumnClient
	SiteItem                                                    *siteitem.SiteItemClient
	SiteLastModifiedByUser                                      *sitelastmodifiedbyuser.SiteLastModifiedByUserClient
	SiteLastModifiedByUserMailboxSetting                        *sitelastmodifiedbyusermailboxsetting.SiteLastModifiedByUserMailboxSettingClient
	SiteLastModifiedByUserServiceProvisioningError              *sitelastmodifiedbyuserserviceprovisioningerror.SiteLastModifiedByUserServiceProvisioningErrorClient
	SiteList                                                    *sitelist.SiteListClient
	SiteListColumn                                              *sitelistcolumn.SiteListColumnClient
	SiteListColumnSourceColumn                                  *sitelistcolumnsourcecolumn.SiteListColumnSourceColumnClient
	SiteListContentType                                         *sitelistcontenttype.SiteListContentTypeClient
	SiteListContentTypeBase                                     *sitelistcontenttypebase.SiteListContentTypeBaseClient
	SiteListContentTypeBaseType                                 *sitelistcontenttypebasetype.SiteListContentTypeBaseTypeClient
	SiteListContentTypeColumn                                   *sitelistcontenttypecolumn.SiteListContentTypeColumnClient
	SiteListContentTypeColumnLink                               *sitelistcontenttypecolumnlink.SiteListContentTypeColumnLinkClient
	SiteListContentTypeColumnPosition                           *sitelistcontenttypecolumnposition.SiteListContentTypeColumnPositionClient
	SiteListContentTypeColumnSourceColumn                       *sitelistcontenttypecolumnsourcecolumn.SiteListContentTypeColumnSourceColumnClient
	SiteListCreatedByUser                                       *sitelistcreatedbyuser.SiteListCreatedByUserClient
	SiteListCreatedByUserMailboxSetting                         *sitelistcreatedbyusermailboxsetting.SiteListCreatedByUserMailboxSettingClient
	SiteListCreatedByUserServiceProvisioningError               *sitelistcreatedbyuserserviceprovisioningerror.SiteListCreatedByUserServiceProvisioningErrorClient
	SiteListDrive                                               *sitelistdrive.SiteListDriveClient
	SiteListItem                                                *sitelistitem.SiteListItemClient
	SiteListItemAnalytics                                       *sitelistitemanalytics.SiteListItemAnalyticsClient
	SiteListItemCreatedByUser                                   *sitelistitemcreatedbyuser.SiteListItemCreatedByUserClient
	SiteListItemCreatedByUserMailboxSetting                     *sitelistitemcreatedbyusermailboxsetting.SiteListItemCreatedByUserMailboxSettingClient
	SiteListItemCreatedByUserServiceProvisioningError           *sitelistitemcreatedbyuserserviceprovisioningerror.SiteListItemCreatedByUserServiceProvisioningErrorClient
	SiteListItemDocumentSetVersion                              *sitelistitemdocumentsetversion.SiteListItemDocumentSetVersionClient
	SiteListItemDocumentSetVersionField                         *sitelistitemdocumentsetversionfield.SiteListItemDocumentSetVersionFieldClient
	SiteListItemDriveItem                                       *sitelistitemdriveitem.SiteListItemDriveItemClient
	SiteListItemDriveItemContent                                *sitelistitemdriveitemcontent.SiteListItemDriveItemContentClient
	SiteListItemField                                           *sitelistitemfield.SiteListItemFieldClient
	SiteListItemLastModifiedByUser                              *sitelistitemlastmodifiedbyuser.SiteListItemLastModifiedByUserClient
	SiteListItemLastModifiedByUserMailboxSetting                *sitelistitemlastmodifiedbyusermailboxsetting.SiteListItemLastModifiedByUserMailboxSettingClient
	SiteListItemLastModifiedByUserServiceProvisioningError      *sitelistitemlastmodifiedbyuserserviceprovisioningerror.SiteListItemLastModifiedByUserServiceProvisioningErrorClient
	SiteListItemVersion                                         *sitelistitemversion.SiteListItemVersionClient
	SiteListItemVersionField                                    *sitelistitemversionfield.SiteListItemVersionFieldClient
	SiteListLastModifiedByUser                                  *sitelistlastmodifiedbyuser.SiteListLastModifiedByUserClient
	SiteListLastModifiedByUserMailboxSetting                    *sitelistlastmodifiedbyusermailboxsetting.SiteListLastModifiedByUserMailboxSettingClient
	SiteListLastModifiedByUserServiceProvisioningError          *sitelistlastmodifiedbyuserserviceprovisioningerror.SiteListLastModifiedByUserServiceProvisioningErrorClient
	SiteListOperation                                           *sitelistoperation.SiteListOperationClient
	SiteListSubscription                                        *sitelistsubscription.SiteListSubscriptionClient
	SiteOperation                                               *siteoperation.SiteOperationClient
	SitePage                                                    *sitepage.SitePageClient
	SitePageCreatedByUser                                       *sitepagecreatedbyuser.SitePageCreatedByUserClient
	SitePageCreatedByUserMailboxSetting                         *sitepagecreatedbyusermailboxsetting.SitePageCreatedByUserMailboxSettingClient
	SitePageCreatedByUserServiceProvisioningError               *sitepagecreatedbyuserserviceprovisioningerror.SitePageCreatedByUserServiceProvisioningErrorClient
	SitePageLastModifiedByUser                                  *sitepagelastmodifiedbyuser.SitePageLastModifiedByUserClient
	SitePageLastModifiedByUserMailboxSetting                    *sitepagelastmodifiedbyusermailboxsetting.SitePageLastModifiedByUserMailboxSettingClient
	SitePageLastModifiedByUserServiceProvisioningError          *sitepagelastmodifiedbyuserserviceprovisioningerror.SitePageLastModifiedByUserServiceProvisioningErrorClient
	SitePermission                                              *sitepermission.SitePermissionClient
	Team                                                        *team.TeamClient
	TeamAllChannel                                              *teamallchannel.TeamAllChannelClient
	TeamChannel                                                 *teamchannel.TeamChannelClient
	TeamChannelFilesFolder                                      *teamchannelfilesfolder.TeamChannelFilesFolderClient
	TeamChannelFilesFolderContent                               *teamchannelfilesfoldercontent.TeamChannelFilesFolderContentClient
	TeamChannelMember                                           *teamchannelmember.TeamChannelMemberClient
	TeamChannelMessage                                          *teamchannelmessage.TeamChannelMessageClient
	TeamChannelMessageHostedContent                             *teamchannelmessagehostedcontent.TeamChannelMessageHostedContentClient
	TeamChannelMessageReply                                     *teamchannelmessagereply.TeamChannelMessageReplyClient
	TeamChannelMessageReplyHostedContent                        *teamchannelmessagereplyhostedcontent.TeamChannelMessageReplyHostedContentClient
	TeamChannelSharedWithTeam                                   *teamchannelsharedwithteam.TeamChannelSharedWithTeamClient
	TeamChannelSharedWithTeamAllowedMember                      *teamchannelsharedwithteamallowedmember.TeamChannelSharedWithTeamAllowedMemberClient
	TeamChannelSharedWithTeamTeam                               *teamchannelsharedwithteamteam.TeamChannelSharedWithTeamTeamClient
	TeamChannelTab                                              *teamchanneltab.TeamChannelTabClient
	TeamChannelTabTeamsApp                                      *teamchanneltabteamsapp.TeamChannelTabTeamsAppClient
	TeamGroup                                                   *teamgroup.TeamGroupClient
	TeamGroupServiceProvisioningError                           *teamgroupserviceprovisioningerror.TeamGroupServiceProvisioningErrorClient
	TeamIncomingChannel                                         *teamincomingchannel.TeamIncomingChannelClient
	TeamInstalledApp                                            *teaminstalledapp.TeamInstalledAppClient
	TeamInstalledAppTeamsApp                                    *teaminstalledappteamsapp.TeamInstalledAppTeamsAppClient
	TeamInstalledAppTeamsAppDefinition                          *teaminstalledappteamsappdefinition.TeamInstalledAppTeamsAppDefinitionClient
	TeamMember                                                  *teammember.TeamMemberClient
	TeamOperation                                               *teamoperation.TeamOperationClient
	TeamPermissionGrant                                         *teampermissiongrant.TeamPermissionGrantClient
	TeamPhoto                                                   *teamphoto.TeamPhotoClient
	TeamPrimaryChannel                                          *teamprimarychannel.TeamPrimaryChannelClient
	TeamPrimaryChannelFilesFolder                               *teamprimarychannelfilesfolder.TeamPrimaryChannelFilesFolderClient
	TeamPrimaryChannelFilesFolderContent                        *teamprimarychannelfilesfoldercontent.TeamPrimaryChannelFilesFolderContentClient
	TeamPrimaryChannelMember                                    *teamprimarychannelmember.TeamPrimaryChannelMemberClient
	TeamPrimaryChannelMessage                                   *teamprimarychannelmessage.TeamPrimaryChannelMessageClient
	TeamPrimaryChannelMessageHostedContent                      *teamprimarychannelmessagehostedcontent.TeamPrimaryChannelMessageHostedContentClient
	TeamPrimaryChannelMessageReply                              *teamprimarychannelmessagereply.TeamPrimaryChannelMessageReplyClient
	TeamPrimaryChannelMessageReplyHostedContent                 *teamprimarychannelmessagereplyhostedcontent.TeamPrimaryChannelMessageReplyHostedContentClient
	TeamPrimaryChannelSharedWithTeam                            *teamprimarychannelsharedwithteam.TeamPrimaryChannelSharedWithTeamClient
	TeamPrimaryChannelSharedWithTeamAllowedMember               *teamprimarychannelsharedwithteamallowedmember.TeamPrimaryChannelSharedWithTeamAllowedMemberClient
	TeamPrimaryChannelSharedWithTeamTeam                        *teamprimarychannelsharedwithteamteam.TeamPrimaryChannelSharedWithTeamTeamClient
	TeamPrimaryChannelTab                                       *teamprimarychanneltab.TeamPrimaryChannelTabClient
	TeamPrimaryChannelTabTeamsApp                               *teamprimarychanneltabteamsapp.TeamPrimaryChannelTabTeamsAppClient
	TeamSchedule                                                *teamschedule.TeamScheduleClient
	TeamScheduleOfferShiftRequest                               *teamscheduleoffershiftrequest.TeamScheduleOfferShiftRequestClient
	TeamScheduleOpenShift                                       *teamscheduleopenshift.TeamScheduleOpenShiftClient
	TeamScheduleOpenShiftChangeRequest                          *teamscheduleopenshiftchangerequest.TeamScheduleOpenShiftChangeRequestClient
	TeamScheduleSchedulingGroup                                 *teamscheduleschedulinggroup.TeamScheduleSchedulingGroupClient
	TeamScheduleShift                                           *teamscheduleshift.TeamScheduleShiftClient
	TeamScheduleSwapShiftsChangeRequest                         *teamscheduleswapshiftschangerequest.TeamScheduleSwapShiftsChangeRequestClient
	TeamScheduleTimeOffReason                                   *teamscheduletimeoffreason.TeamScheduleTimeOffReasonClient
	TeamScheduleTimeOffRequest                                  *teamscheduletimeoffrequest.TeamScheduleTimeOffRequestClient
	TeamScheduleTimesOff                                        *teamscheduletimesoff.TeamScheduleTimesOffClient
	TeamTag                                                     *teamtag.TeamTagClient
	TeamTagMember                                               *teamtagmember.TeamTagMemberClient
	TeamTemplate                                                *teamtemplate.TeamTemplateClient
	Thread                                                      *thread.ThreadClient
	ThreadPost                                                  *threadpost.ThreadPostClient
	ThreadPostAttachment                                        *threadpostattachment.ThreadPostAttachmentClient
	ThreadPostExtension                                         *threadpostextension.ThreadPostExtensionClient
	ThreadPostInReplyTo                                         *threadpostinreplyto.ThreadPostInReplyToClient
	ThreadPostInReplyToAttachment                               *threadpostinreplytoattachment.ThreadPostInReplyToAttachmentClient
	ThreadPostInReplyToExtension                                *threadpostinreplytoextension.ThreadPostInReplyToExtensionClient
	TransitiveMember                                            *transitivemember.TransitiveMemberClient
	TransitiveMemberOf                                          *transitivememberof.TransitiveMemberOfClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	acceptedSenderClient, err := acceptedsender.NewAcceptedSenderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AcceptedSender client: %+v", err)
	}
	configureFunc(acceptedSenderClient.Client)

	appRoleAssignmentClient, err := approleassignment.NewAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppRoleAssignment client: %+v", err)
	}
	configureFunc(appRoleAssignmentClient.Client)

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

	conversationClient, err := conversation.NewConversationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Conversation client: %+v", err)
	}
	configureFunc(conversationClient.Client)

	conversationThreadClient, err := conversationthread.NewConversationThreadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThread client: %+v", err)
	}
	configureFunc(conversationThreadClient.Client)

	conversationThreadPostAttachmentClient, err := conversationthreadpostattachment.NewConversationThreadPostAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPostAttachment client: %+v", err)
	}
	configureFunc(conversationThreadPostAttachmentClient.Client)

	conversationThreadPostClient, err := conversationthreadpost.NewConversationThreadPostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPost client: %+v", err)
	}
	configureFunc(conversationThreadPostClient.Client)

	conversationThreadPostExtensionClient, err := conversationthreadpostextension.NewConversationThreadPostExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPostExtension client: %+v", err)
	}
	configureFunc(conversationThreadPostExtensionClient.Client)

	conversationThreadPostInReplyToAttachmentClient, err := conversationthreadpostinreplytoattachment.NewConversationThreadPostInReplyToAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPostInReplyToAttachment client: %+v", err)
	}
	configureFunc(conversationThreadPostInReplyToAttachmentClient.Client)

	conversationThreadPostInReplyToClient, err := conversationthreadpostinreplyto.NewConversationThreadPostInReplyToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPostInReplyTo client: %+v", err)
	}
	configureFunc(conversationThreadPostInReplyToClient.Client)

	conversationThreadPostInReplyToExtensionClient, err := conversationthreadpostinreplytoextension.NewConversationThreadPostInReplyToExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPostInReplyToExtension client: %+v", err)
	}
	configureFunc(conversationThreadPostInReplyToExtensionClient.Client)

	createdOnBehalfOfClient, err := createdonbehalfof.NewCreatedOnBehalfOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreatedOnBehalfOf client: %+v", err)
	}
	configureFunc(createdOnBehalfOfClient.Client)

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

	groupClient, err := group.NewGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Group client: %+v", err)
	}
	configureFunc(groupClient.Client)

	groupLifecyclePolicyClient, err := grouplifecyclepolicy.NewGroupLifecyclePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupLifecyclePolicy client: %+v", err)
	}
	configureFunc(groupLifecyclePolicyClient.Client)

	memberClient, err := member.NewMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Member client: %+v", err)
	}
	configureFunc(memberClient.Client)

	memberOfClient, err := memberof.NewMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MemberOf client: %+v", err)
	}
	configureFunc(memberOfClient.Client)

	membersWithLicenseErrorClient, err := memberswithlicenseerror.NewMembersWithLicenseErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MembersWithLicenseError client: %+v", err)
	}
	configureFunc(membersWithLicenseErrorClient.Client)

	ownerClient, err := owner.NewOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Owner client: %+v", err)
	}
	configureFunc(ownerClient.Client)

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

	rejectedSenderClient, err := rejectedsender.NewRejectedSenderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RejectedSender client: %+v", err)
	}
	configureFunc(rejectedSenderClient.Client)

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

	siteAnalyticsAllTimeClient, err := siteanalyticsalltime.NewSiteAnalyticsAllTimeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticsAllTime client: %+v", err)
	}
	configureFunc(siteAnalyticsAllTimeClient.Client)

	siteAnalyticsClient, err := siteanalytics.NewSiteAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalytics client: %+v", err)
	}
	configureFunc(siteAnalyticsClient.Client)

	siteAnalyticsItemActivityStatActivityClient, err := siteanalyticsitemactivitystatactivity.NewSiteAnalyticsItemActivityStatActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticsItemActivityStatActivity client: %+v", err)
	}
	configureFunc(siteAnalyticsItemActivityStatActivityClient.Client)

	siteAnalyticsItemActivityStatActivityDriveItemClient, err := siteanalyticsitemactivitystatactivitydriveitem.NewSiteAnalyticsItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticsItemActivityStatActivityDriveItem client: %+v", err)
	}
	configureFunc(siteAnalyticsItemActivityStatActivityDriveItemClient.Client)

	siteAnalyticsItemActivityStatActivityDriveItemContentClient, err := siteanalyticsitemactivitystatactivitydriveitemcontent.NewSiteAnalyticsItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticsItemActivityStatActivityDriveItemContent client: %+v", err)
	}
	configureFunc(siteAnalyticsItemActivityStatActivityDriveItemContentClient.Client)

	siteAnalyticsItemActivityStatClient, err := siteanalyticsitemactivitystat.NewSiteAnalyticsItemActivityStatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticsItemActivityStat client: %+v", err)
	}
	configureFunc(siteAnalyticsItemActivityStatClient.Client)

	siteAnalyticsLastSevenDayClient, err := siteanalyticslastsevenday.NewSiteAnalyticsLastSevenDayClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticsLastSevenDay client: %+v", err)
	}
	configureFunc(siteAnalyticsLastSevenDayClient.Client)

	siteClient, err := site.NewSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Site client: %+v", err)
	}
	configureFunc(siteClient.Client)

	siteColumnClient, err := sitecolumn.NewSiteColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteColumn client: %+v", err)
	}
	configureFunc(siteColumnClient.Client)

	siteColumnSourceColumnClient, err := sitecolumnsourcecolumn.NewSiteColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteColumnSourceColumn client: %+v", err)
	}
	configureFunc(siteColumnSourceColumnClient.Client)

	siteContentTypeBaseClient, err := sitecontenttypebase.NewSiteContentTypeBaseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentTypeBase client: %+v", err)
	}
	configureFunc(siteContentTypeBaseClient.Client)

	siteContentTypeBaseTypeClient, err := sitecontenttypebasetype.NewSiteContentTypeBaseTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentTypeBaseType client: %+v", err)
	}
	configureFunc(siteContentTypeBaseTypeClient.Client)

	siteContentTypeClient, err := sitecontenttype.NewSiteContentTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentType client: %+v", err)
	}
	configureFunc(siteContentTypeClient.Client)

	siteContentTypeColumnClient, err := sitecontenttypecolumn.NewSiteContentTypeColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentTypeColumn client: %+v", err)
	}
	configureFunc(siteContentTypeColumnClient.Client)

	siteContentTypeColumnLinkClient, err := sitecontenttypecolumnlink.NewSiteContentTypeColumnLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentTypeColumnLink client: %+v", err)
	}
	configureFunc(siteContentTypeColumnLinkClient.Client)

	siteContentTypeColumnPositionClient, err := sitecontenttypecolumnposition.NewSiteContentTypeColumnPositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentTypeColumnPosition client: %+v", err)
	}
	configureFunc(siteContentTypeColumnPositionClient.Client)

	siteContentTypeColumnSourceColumnClient, err := sitecontenttypecolumnsourcecolumn.NewSiteContentTypeColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentTypeColumnSourceColumn client: %+v", err)
	}
	configureFunc(siteContentTypeColumnSourceColumnClient.Client)

	siteCreatedByUserClient, err := sitecreatedbyuser.NewSiteCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteCreatedByUser client: %+v", err)
	}
	configureFunc(siteCreatedByUserClient.Client)

	siteCreatedByUserMailboxSettingClient, err := sitecreatedbyusermailboxsetting.NewSiteCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteCreatedByUserMailboxSettingClient.Client)

	siteCreatedByUserServiceProvisioningErrorClient, err := sitecreatedbyuserserviceprovisioningerror.NewSiteCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteCreatedByUserServiceProvisioningErrorClient.Client)

	siteDriveClient, err := sitedrive.NewSiteDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteDrive client: %+v", err)
	}
	configureFunc(siteDriveClient.Client)

	siteExternalColumnClient, err := siteexternalcolumn.NewSiteExternalColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteExternalColumn client: %+v", err)
	}
	configureFunc(siteExternalColumnClient.Client)

	siteItemClient, err := siteitem.NewSiteItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteItem client: %+v", err)
	}
	configureFunc(siteItemClient.Client)

	siteLastModifiedByUserClient, err := sitelastmodifiedbyuser.NewSiteLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteLastModifiedByUser client: %+v", err)
	}
	configureFunc(siteLastModifiedByUserClient.Client)

	siteLastModifiedByUserMailboxSettingClient, err := sitelastmodifiedbyusermailboxsetting.NewSiteLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteLastModifiedByUserMailboxSettingClient.Client)

	siteLastModifiedByUserServiceProvisioningErrorClient, err := sitelastmodifiedbyuserserviceprovisioningerror.NewSiteLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteLastModifiedByUserServiceProvisioningErrorClient.Client)

	siteListClient, err := sitelist.NewSiteListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteList client: %+v", err)
	}
	configureFunc(siteListClient.Client)

	siteListColumnClient, err := sitelistcolumn.NewSiteListColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListColumn client: %+v", err)
	}
	configureFunc(siteListColumnClient.Client)

	siteListColumnSourceColumnClient, err := sitelistcolumnsourcecolumn.NewSiteListColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListColumnSourceColumn client: %+v", err)
	}
	configureFunc(siteListColumnSourceColumnClient.Client)

	siteListContentTypeBaseClient, err := sitelistcontenttypebase.NewSiteListContentTypeBaseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListContentTypeBase client: %+v", err)
	}
	configureFunc(siteListContentTypeBaseClient.Client)

	siteListContentTypeBaseTypeClient, err := sitelistcontenttypebasetype.NewSiteListContentTypeBaseTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListContentTypeBaseType client: %+v", err)
	}
	configureFunc(siteListContentTypeBaseTypeClient.Client)

	siteListContentTypeClient, err := sitelistcontenttype.NewSiteListContentTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListContentType client: %+v", err)
	}
	configureFunc(siteListContentTypeClient.Client)

	siteListContentTypeColumnClient, err := sitelistcontenttypecolumn.NewSiteListContentTypeColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListContentTypeColumn client: %+v", err)
	}
	configureFunc(siteListContentTypeColumnClient.Client)

	siteListContentTypeColumnLinkClient, err := sitelistcontenttypecolumnlink.NewSiteListContentTypeColumnLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListContentTypeColumnLink client: %+v", err)
	}
	configureFunc(siteListContentTypeColumnLinkClient.Client)

	siteListContentTypeColumnPositionClient, err := sitelistcontenttypecolumnposition.NewSiteListContentTypeColumnPositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListContentTypeColumnPosition client: %+v", err)
	}
	configureFunc(siteListContentTypeColumnPositionClient.Client)

	siteListContentTypeColumnSourceColumnClient, err := sitelistcontenttypecolumnsourcecolumn.NewSiteListContentTypeColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListContentTypeColumnSourceColumn client: %+v", err)
	}
	configureFunc(siteListContentTypeColumnSourceColumnClient.Client)

	siteListCreatedByUserClient, err := sitelistcreatedbyuser.NewSiteListCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListCreatedByUser client: %+v", err)
	}
	configureFunc(siteListCreatedByUserClient.Client)

	siteListCreatedByUserMailboxSettingClient, err := sitelistcreatedbyusermailboxsetting.NewSiteListCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteListCreatedByUserMailboxSettingClient.Client)

	siteListCreatedByUserServiceProvisioningErrorClient, err := sitelistcreatedbyuserserviceprovisioningerror.NewSiteListCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteListCreatedByUserServiceProvisioningErrorClient.Client)

	siteListDriveClient, err := sitelistdrive.NewSiteListDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListDrive client: %+v", err)
	}
	configureFunc(siteListDriveClient.Client)

	siteListItemAnalyticsClient, err := sitelistitemanalytics.NewSiteListItemAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemAnalytics client: %+v", err)
	}
	configureFunc(siteListItemAnalyticsClient.Client)

	siteListItemClient, err := sitelistitem.NewSiteListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItem client: %+v", err)
	}
	configureFunc(siteListItemClient.Client)

	siteListItemCreatedByUserClient, err := sitelistitemcreatedbyuser.NewSiteListItemCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemCreatedByUser client: %+v", err)
	}
	configureFunc(siteListItemCreatedByUserClient.Client)

	siteListItemCreatedByUserMailboxSettingClient, err := sitelistitemcreatedbyusermailboxsetting.NewSiteListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteListItemCreatedByUserMailboxSettingClient.Client)

	siteListItemCreatedByUserServiceProvisioningErrorClient, err := sitelistitemcreatedbyuserserviceprovisioningerror.NewSiteListItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteListItemCreatedByUserServiceProvisioningErrorClient.Client)

	siteListItemDocumentSetVersionClient, err := sitelistitemdocumentsetversion.NewSiteListItemDocumentSetVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemDocumentSetVersion client: %+v", err)
	}
	configureFunc(siteListItemDocumentSetVersionClient.Client)

	siteListItemDocumentSetVersionFieldClient, err := sitelistitemdocumentsetversionfield.NewSiteListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemDocumentSetVersionField client: %+v", err)
	}
	configureFunc(siteListItemDocumentSetVersionFieldClient.Client)

	siteListItemDriveItemClient, err := sitelistitemdriveitem.NewSiteListItemDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemDriveItem client: %+v", err)
	}
	configureFunc(siteListItemDriveItemClient.Client)

	siteListItemDriveItemContentClient, err := sitelistitemdriveitemcontent.NewSiteListItemDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemDriveItemContent client: %+v", err)
	}
	configureFunc(siteListItemDriveItemContentClient.Client)

	siteListItemFieldClient, err := sitelistitemfield.NewSiteListItemFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemField client: %+v", err)
	}
	configureFunc(siteListItemFieldClient.Client)

	siteListItemLastModifiedByUserClient, err := sitelistitemlastmodifiedbyuser.NewSiteListItemLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemLastModifiedByUser client: %+v", err)
	}
	configureFunc(siteListItemLastModifiedByUserClient.Client)

	siteListItemLastModifiedByUserMailboxSettingClient, err := sitelistitemlastmodifiedbyusermailboxsetting.NewSiteListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteListItemLastModifiedByUserMailboxSettingClient.Client)

	siteListItemLastModifiedByUserServiceProvisioningErrorClient, err := sitelistitemlastmodifiedbyuserserviceprovisioningerror.NewSiteListItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteListItemLastModifiedByUserServiceProvisioningErrorClient.Client)

	siteListItemVersionClient, err := sitelistitemversion.NewSiteListItemVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemVersion client: %+v", err)
	}
	configureFunc(siteListItemVersionClient.Client)

	siteListItemVersionFieldClient, err := sitelistitemversionfield.NewSiteListItemVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemVersionField client: %+v", err)
	}
	configureFunc(siteListItemVersionFieldClient.Client)

	siteListLastModifiedByUserClient, err := sitelistlastmodifiedbyuser.NewSiteListLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListLastModifiedByUser client: %+v", err)
	}
	configureFunc(siteListLastModifiedByUserClient.Client)

	siteListLastModifiedByUserMailboxSettingClient, err := sitelistlastmodifiedbyusermailboxsetting.NewSiteListLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteListLastModifiedByUserMailboxSettingClient.Client)

	siteListLastModifiedByUserServiceProvisioningErrorClient, err := sitelistlastmodifiedbyuserserviceprovisioningerror.NewSiteListLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteListLastModifiedByUserServiceProvisioningErrorClient.Client)

	siteListOperationClient, err := sitelistoperation.NewSiteListOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListOperation client: %+v", err)
	}
	configureFunc(siteListOperationClient.Client)

	siteListSubscriptionClient, err := sitelistsubscription.NewSiteListSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListSubscription client: %+v", err)
	}
	configureFunc(siteListSubscriptionClient.Client)

	siteOperationClient, err := siteoperation.NewSiteOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteOperation client: %+v", err)
	}
	configureFunc(siteOperationClient.Client)

	sitePageClient, err := sitepage.NewSitePageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePage client: %+v", err)
	}
	configureFunc(sitePageClient.Client)

	sitePageCreatedByUserClient, err := sitepagecreatedbyuser.NewSitePageCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePageCreatedByUser client: %+v", err)
	}
	configureFunc(sitePageCreatedByUserClient.Client)

	sitePageCreatedByUserMailboxSettingClient, err := sitepagecreatedbyusermailboxsetting.NewSitePageCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePageCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(sitePageCreatedByUserMailboxSettingClient.Client)

	sitePageCreatedByUserServiceProvisioningErrorClient, err := sitepagecreatedbyuserserviceprovisioningerror.NewSitePageCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePageCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(sitePageCreatedByUserServiceProvisioningErrorClient.Client)

	sitePageLastModifiedByUserClient, err := sitepagelastmodifiedbyuser.NewSitePageLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePageLastModifiedByUser client: %+v", err)
	}
	configureFunc(sitePageLastModifiedByUserClient.Client)

	sitePageLastModifiedByUserMailboxSettingClient, err := sitepagelastmodifiedbyusermailboxsetting.NewSitePageLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePageLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(sitePageLastModifiedByUserMailboxSettingClient.Client)

	sitePageLastModifiedByUserServiceProvisioningErrorClient, err := sitepagelastmodifiedbyuserserviceprovisioningerror.NewSitePageLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePageLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(sitePageLastModifiedByUserServiceProvisioningErrorClient.Client)

	sitePermissionClient, err := sitepermission.NewSitePermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SitePermission client: %+v", err)
	}
	configureFunc(sitePermissionClient.Client)

	teamAllChannelClient, err := teamallchannel.NewTeamAllChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamAllChannel client: %+v", err)
	}
	configureFunc(teamAllChannelClient.Client)

	teamChannelClient, err := teamchannel.NewTeamChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannel client: %+v", err)
	}
	configureFunc(teamChannelClient.Client)

	teamChannelFilesFolderClient, err := teamchannelfilesfolder.NewTeamChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelFilesFolder client: %+v", err)
	}
	configureFunc(teamChannelFilesFolderClient.Client)

	teamChannelFilesFolderContentClient, err := teamchannelfilesfoldercontent.NewTeamChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(teamChannelFilesFolderContentClient.Client)

	teamChannelMemberClient, err := teamchannelmember.NewTeamChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelMember client: %+v", err)
	}
	configureFunc(teamChannelMemberClient.Client)

	teamChannelMessageClient, err := teamchannelmessage.NewTeamChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelMessage client: %+v", err)
	}
	configureFunc(teamChannelMessageClient.Client)

	teamChannelMessageHostedContentClient, err := teamchannelmessagehostedcontent.NewTeamChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(teamChannelMessageHostedContentClient.Client)

	teamChannelMessageReplyClient, err := teamchannelmessagereply.NewTeamChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelMessageReply client: %+v", err)
	}
	configureFunc(teamChannelMessageReplyClient.Client)

	teamChannelMessageReplyHostedContentClient, err := teamchannelmessagereplyhostedcontent.NewTeamChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(teamChannelMessageReplyHostedContentClient.Client)

	teamChannelSharedWithTeamAllowedMemberClient, err := teamchannelsharedwithteamallowedmember.NewTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(teamChannelSharedWithTeamAllowedMemberClient.Client)

	teamChannelSharedWithTeamClient, err := teamchannelsharedwithteam.NewTeamChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(teamChannelSharedWithTeamClient.Client)

	teamChannelSharedWithTeamTeamClient, err := teamchannelsharedwithteamteam.NewTeamChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(teamChannelSharedWithTeamTeamClient.Client)

	teamChannelTabClient, err := teamchanneltab.NewTeamChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelTab client: %+v", err)
	}
	configureFunc(teamChannelTabClient.Client)

	teamChannelTabTeamsAppClient, err := teamchanneltabteamsapp.NewTeamChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(teamChannelTabTeamsAppClient.Client)

	teamClient, err := team.NewTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Team client: %+v", err)
	}
	configureFunc(teamClient.Client)

	teamGroupClient, err := teamgroup.NewTeamGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamGroup client: %+v", err)
	}
	configureFunc(teamGroupClient.Client)

	teamGroupServiceProvisioningErrorClient, err := teamgroupserviceprovisioningerror.NewTeamGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(teamGroupServiceProvisioningErrorClient.Client)

	teamIncomingChannelClient, err := teamincomingchannel.NewTeamIncomingChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamIncomingChannel client: %+v", err)
	}
	configureFunc(teamIncomingChannelClient.Client)

	teamInstalledAppClient, err := teaminstalledapp.NewTeamInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamInstalledApp client: %+v", err)
	}
	configureFunc(teamInstalledAppClient.Client)

	teamInstalledAppTeamsAppClient, err := teaminstalledappteamsapp.NewTeamInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(teamInstalledAppTeamsAppClient.Client)

	teamInstalledAppTeamsAppDefinitionClient, err := teaminstalledappteamsappdefinition.NewTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(teamInstalledAppTeamsAppDefinitionClient.Client)

	teamMemberClient, err := teammember.NewTeamMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamMember client: %+v", err)
	}
	configureFunc(teamMemberClient.Client)

	teamOperationClient, err := teamoperation.NewTeamOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamOperation client: %+v", err)
	}
	configureFunc(teamOperationClient.Client)

	teamPermissionGrantClient, err := teampermissiongrant.NewTeamPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPermissionGrant client: %+v", err)
	}
	configureFunc(teamPermissionGrantClient.Client)

	teamPhotoClient, err := teamphoto.NewTeamPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPhoto client: %+v", err)
	}
	configureFunc(teamPhotoClient.Client)

	teamPrimaryChannelClient, err := teamprimarychannel.NewTeamPrimaryChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannel client: %+v", err)
	}
	configureFunc(teamPrimaryChannelClient.Client)

	teamPrimaryChannelFilesFolderClient, err := teamprimarychannelfilesfolder.NewTeamPrimaryChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelFilesFolder client: %+v", err)
	}
	configureFunc(teamPrimaryChannelFilesFolderClient.Client)

	teamPrimaryChannelFilesFolderContentClient, err := teamprimarychannelfilesfoldercontent.NewTeamPrimaryChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(teamPrimaryChannelFilesFolderContentClient.Client)

	teamPrimaryChannelMemberClient, err := teamprimarychannelmember.NewTeamPrimaryChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelMember client: %+v", err)
	}
	configureFunc(teamPrimaryChannelMemberClient.Client)

	teamPrimaryChannelMessageClient, err := teamprimarychannelmessage.NewTeamPrimaryChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelMessage client: %+v", err)
	}
	configureFunc(teamPrimaryChannelMessageClient.Client)

	teamPrimaryChannelMessageHostedContentClient, err := teamprimarychannelmessagehostedcontent.NewTeamPrimaryChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(teamPrimaryChannelMessageHostedContentClient.Client)

	teamPrimaryChannelMessageReplyClient, err := teamprimarychannelmessagereply.NewTeamPrimaryChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelMessageReply client: %+v", err)
	}
	configureFunc(teamPrimaryChannelMessageReplyClient.Client)

	teamPrimaryChannelMessageReplyHostedContentClient, err := teamprimarychannelmessagereplyhostedcontent.NewTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(teamPrimaryChannelMessageReplyHostedContentClient.Client)

	teamPrimaryChannelSharedWithTeamAllowedMemberClient, err := teamprimarychannelsharedwithteamallowedmember.NewTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(teamPrimaryChannelSharedWithTeamAllowedMemberClient.Client)

	teamPrimaryChannelSharedWithTeamClient, err := teamprimarychannelsharedwithteam.NewTeamPrimaryChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(teamPrimaryChannelSharedWithTeamClient.Client)

	teamPrimaryChannelSharedWithTeamTeamClient, err := teamprimarychannelsharedwithteamteam.NewTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(teamPrimaryChannelSharedWithTeamTeamClient.Client)

	teamPrimaryChannelTabClient, err := teamprimarychanneltab.NewTeamPrimaryChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelTab client: %+v", err)
	}
	configureFunc(teamPrimaryChannelTabClient.Client)

	teamPrimaryChannelTabTeamsAppClient, err := teamprimarychanneltabteamsapp.NewTeamPrimaryChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(teamPrimaryChannelTabTeamsAppClient.Client)

	teamScheduleClient, err := teamschedule.NewTeamScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamSchedule client: %+v", err)
	}
	configureFunc(teamScheduleClient.Client)

	teamScheduleOfferShiftRequestClient, err := teamscheduleoffershiftrequest.NewTeamScheduleOfferShiftRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleOfferShiftRequest client: %+v", err)
	}
	configureFunc(teamScheduleOfferShiftRequestClient.Client)

	teamScheduleOpenShiftChangeRequestClient, err := teamscheduleopenshiftchangerequest.NewTeamScheduleOpenShiftChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleOpenShiftChangeRequest client: %+v", err)
	}
	configureFunc(teamScheduleOpenShiftChangeRequestClient.Client)

	teamScheduleOpenShiftClient, err := teamscheduleopenshift.NewTeamScheduleOpenShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleOpenShift client: %+v", err)
	}
	configureFunc(teamScheduleOpenShiftClient.Client)

	teamScheduleSchedulingGroupClient, err := teamscheduleschedulinggroup.NewTeamScheduleSchedulingGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleSchedulingGroup client: %+v", err)
	}
	configureFunc(teamScheduleSchedulingGroupClient.Client)

	teamScheduleShiftClient, err := teamscheduleshift.NewTeamScheduleShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleShift client: %+v", err)
	}
	configureFunc(teamScheduleShiftClient.Client)

	teamScheduleSwapShiftsChangeRequestClient, err := teamscheduleswapshiftschangerequest.NewTeamScheduleSwapShiftsChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleSwapShiftsChangeRequest client: %+v", err)
	}
	configureFunc(teamScheduleSwapShiftsChangeRequestClient.Client)

	teamScheduleTimeOffReasonClient, err := teamscheduletimeoffreason.NewTeamScheduleTimeOffReasonClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleTimeOffReason client: %+v", err)
	}
	configureFunc(teamScheduleTimeOffReasonClient.Client)

	teamScheduleTimeOffRequestClient, err := teamscheduletimeoffrequest.NewTeamScheduleTimeOffRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleTimeOffRequest client: %+v", err)
	}
	configureFunc(teamScheduleTimeOffRequestClient.Client)

	teamScheduleTimesOffClient, err := teamscheduletimesoff.NewTeamScheduleTimesOffClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleTimesOff client: %+v", err)
	}
	configureFunc(teamScheduleTimesOffClient.Client)

	teamTagClient, err := teamtag.NewTeamTagClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamTag client: %+v", err)
	}
	configureFunc(teamTagClient.Client)

	teamTagMemberClient, err := teamtagmember.NewTeamTagMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamTagMember client: %+v", err)
	}
	configureFunc(teamTagMemberClient.Client)

	teamTemplateClient, err := teamtemplate.NewTeamTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamTemplate client: %+v", err)
	}
	configureFunc(teamTemplateClient.Client)

	threadClient, err := thread.NewThreadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Thread client: %+v", err)
	}
	configureFunc(threadClient.Client)

	threadPostAttachmentClient, err := threadpostattachment.NewThreadPostAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPostAttachment client: %+v", err)
	}
	configureFunc(threadPostAttachmentClient.Client)

	threadPostClient, err := threadpost.NewThreadPostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPost client: %+v", err)
	}
	configureFunc(threadPostClient.Client)

	threadPostExtensionClient, err := threadpostextension.NewThreadPostExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPostExtension client: %+v", err)
	}
	configureFunc(threadPostExtensionClient.Client)

	threadPostInReplyToAttachmentClient, err := threadpostinreplytoattachment.NewThreadPostInReplyToAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPostInReplyToAttachment client: %+v", err)
	}
	configureFunc(threadPostInReplyToAttachmentClient.Client)

	threadPostInReplyToClient, err := threadpostinreplyto.NewThreadPostInReplyToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPostInReplyTo client: %+v", err)
	}
	configureFunc(threadPostInReplyToClient.Client)

	threadPostInReplyToExtensionClient, err := threadpostinreplytoextension.NewThreadPostInReplyToExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPostInReplyToExtension client: %+v", err)
	}
	configureFunc(threadPostInReplyToExtensionClient.Client)

	transitiveMemberClient, err := transitivemember.NewTransitiveMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TransitiveMember client: %+v", err)
	}
	configureFunc(transitiveMemberClient.Client)

	transitiveMemberOfClient, err := transitivememberof.NewTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TransitiveMemberOf client: %+v", err)
	}
	configureFunc(transitiveMemberOfClient.Client)

	return &Client{
		AcceptedSender:                                      acceptedSenderClient,
		AppRoleAssignment:                                   appRoleAssignmentClient,
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
		CalendarView:                                        calendarViewClient,
		CalendarViewAttachment:                              calendarViewAttachmentClient,
		CalendarViewCalendar:                                calendarViewCalendarClient,
		CalendarViewExtension:                               calendarViewExtensionClient,
		CalendarViewInstance:                                calendarViewInstanceClient,
		CalendarViewInstanceAttachment:                      calendarViewInstanceAttachmentClient,
		CalendarViewInstanceCalendar:                        calendarViewInstanceCalendarClient,
		CalendarViewInstanceExtension:                       calendarViewInstanceExtensionClient,
		Conversation:                                        conversationClient,
		ConversationThread:                                  conversationThreadClient,
		ConversationThreadPost:                              conversationThreadPostClient,
		ConversationThreadPostAttachment:                    conversationThreadPostAttachmentClient,
		ConversationThreadPostExtension:                     conversationThreadPostExtensionClient,
		ConversationThreadPostInReplyTo:                     conversationThreadPostInReplyToClient,
		ConversationThreadPostInReplyToAttachment:           conversationThreadPostInReplyToAttachmentClient,
		ConversationThreadPostInReplyToExtension:            conversationThreadPostInReplyToExtensionClient,
		CreatedOnBehalfOf:                                   createdOnBehalfOfClient,
		Drive:                                               driveClient,
		DriveBundle:                                         driveBundleClient,
		DriveBundleContent:                                  driveBundleContentClient,
		DriveCreatedByUser:                                  driveCreatedByUserClient,
		DriveCreatedByUserMailboxSetting:                    driveCreatedByUserMailboxSettingClient,
		DriveCreatedByUserServiceProvisioningError:          driveCreatedByUserServiceProvisioningErrorClient,
		DriveFollowing:                                      driveFollowingClient,
		DriveFollowingContent:                               driveFollowingContentClient,
		DriveItem:                                           driveItemClient,
		DriveItemAnalytics:                                  driveItemAnalyticsClient,
		DriveItemAnalyticsAllTime:                           driveItemAnalyticsAllTimeClient,
		DriveItemAnalyticsItemActivityStat:                  driveItemAnalyticsItemActivityStatClient,
		DriveItemAnalyticsItemActivityStatActivity:          driveItemAnalyticsItemActivityStatActivityClient,
		DriveItemAnalyticsItemActivityStatActivityDriveItem: driveItemAnalyticsItemActivityStatActivityDriveItemClient,
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
		Event:                                                       eventClient,
		EventAttachment:                                             eventAttachmentClient,
		EventCalendar:                                               eventCalendarClient,
		EventExtension:                                              eventExtensionClient,
		EventInstance:                                               eventInstanceClient,
		EventInstanceAttachment:                                     eventInstanceAttachmentClient,
		EventInstanceCalendar:                                       eventInstanceCalendarClient,
		EventInstanceExtension:                                      eventInstanceExtensionClient,
		Extension:                                                   extensionClient,
		Group:                                                       groupClient,
		GroupLifecyclePolicy:                                        groupLifecyclePolicyClient,
		Member:                                                      memberClient,
		MemberOf:                                                    memberOfClient,
		MembersWithLicenseError:                                     membersWithLicenseErrorClient,
		Owner:                                                       ownerClient,
		PermissionGrant:                                             permissionGrantClient,
		Photo:                                                       photoClient,
		Planner:                                                     plannerClient,
		PlannerPlan:                                                 plannerPlanClient,
		PlannerPlanBucket:                                           plannerPlanBucketClient,
		PlannerPlanBucketTask:                                       plannerPlanBucketTaskClient,
		PlannerPlanBucketTaskAssignedToTaskBoardFormat:              plannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		PlannerPlanBucketTaskBucketTaskBoardFormat:                  plannerPlanBucketTaskBucketTaskBoardFormatClient,
		PlannerPlanBucketTaskDetail:                                 plannerPlanBucketTaskDetailClient,
		PlannerPlanBucketTaskProgressTaskBoardFormat:                plannerPlanBucketTaskProgressTaskBoardFormatClient,
		PlannerPlanDetail:                                           plannerPlanDetailClient,
		PlannerPlanTask:                                             plannerPlanTaskClient,
		PlannerPlanTaskAssignedToTaskBoardFormat:                    plannerPlanTaskAssignedToTaskBoardFormatClient,
		PlannerPlanTaskBucketTaskBoardFormat:                        plannerPlanTaskBucketTaskBoardFormatClient,
		PlannerPlanTaskDetail:                                       plannerPlanTaskDetailClient,
		PlannerPlanTaskProgressTaskBoardFormat:                      plannerPlanTaskProgressTaskBoardFormatClient,
		RejectedSender:                                              rejectedSenderClient,
		ServiceProvisioningError:                                    serviceProvisioningErrorClient,
		Setting:                                                     settingClient,
		Site:                                                        siteClient,
		SiteAnalytics:                                               siteAnalyticsClient,
		SiteAnalyticsAllTime:                                        siteAnalyticsAllTimeClient,
		SiteAnalyticsItemActivityStat:                               siteAnalyticsItemActivityStatClient,
		SiteAnalyticsItemActivityStatActivity:                       siteAnalyticsItemActivityStatActivityClient,
		SiteAnalyticsItemActivityStatActivityDriveItem:              siteAnalyticsItemActivityStatActivityDriveItemClient,
		SiteAnalyticsItemActivityStatActivityDriveItemContent:       siteAnalyticsItemActivityStatActivityDriveItemContentClient,
		SiteAnalyticsLastSevenDay:                                   siteAnalyticsLastSevenDayClient,
		SiteColumn:                                                  siteColumnClient,
		SiteColumnSourceColumn:                                      siteColumnSourceColumnClient,
		SiteContentType:                                             siteContentTypeClient,
		SiteContentTypeBase:                                         siteContentTypeBaseClient,
		SiteContentTypeBaseType:                                     siteContentTypeBaseTypeClient,
		SiteContentTypeColumn:                                       siteContentTypeColumnClient,
		SiteContentTypeColumnLink:                                   siteContentTypeColumnLinkClient,
		SiteContentTypeColumnPosition:                               siteContentTypeColumnPositionClient,
		SiteContentTypeColumnSourceColumn:                           siteContentTypeColumnSourceColumnClient,
		SiteCreatedByUser:                                           siteCreatedByUserClient,
		SiteCreatedByUserMailboxSetting:                             siteCreatedByUserMailboxSettingClient,
		SiteCreatedByUserServiceProvisioningError:                   siteCreatedByUserServiceProvisioningErrorClient,
		SiteDrive:                            siteDriveClient,
		SiteExternalColumn:                   siteExternalColumnClient,
		SiteItem:                             siteItemClient,
		SiteLastModifiedByUser:               siteLastModifiedByUserClient,
		SiteLastModifiedByUserMailboxSetting: siteLastModifiedByUserMailboxSettingClient,
		SiteLastModifiedByUserServiceProvisioningError: siteLastModifiedByUserServiceProvisioningErrorClient,
		SiteList:                                               siteListClient,
		SiteListColumn:                                         siteListColumnClient,
		SiteListColumnSourceColumn:                             siteListColumnSourceColumnClient,
		SiteListContentType:                                    siteListContentTypeClient,
		SiteListContentTypeBase:                                siteListContentTypeBaseClient,
		SiteListContentTypeBaseType:                            siteListContentTypeBaseTypeClient,
		SiteListContentTypeColumn:                              siteListContentTypeColumnClient,
		SiteListContentTypeColumnLink:                          siteListContentTypeColumnLinkClient,
		SiteListContentTypeColumnPosition:                      siteListContentTypeColumnPositionClient,
		SiteListContentTypeColumnSourceColumn:                  siteListContentTypeColumnSourceColumnClient,
		SiteListCreatedByUser:                                  siteListCreatedByUserClient,
		SiteListCreatedByUserMailboxSetting:                    siteListCreatedByUserMailboxSettingClient,
		SiteListCreatedByUserServiceProvisioningError:          siteListCreatedByUserServiceProvisioningErrorClient,
		SiteListDrive:                                          siteListDriveClient,
		SiteListItem:                                           siteListItemClient,
		SiteListItemAnalytics:                                  siteListItemAnalyticsClient,
		SiteListItemCreatedByUser:                              siteListItemCreatedByUserClient,
		SiteListItemCreatedByUserMailboxSetting:                siteListItemCreatedByUserMailboxSettingClient,
		SiteListItemCreatedByUserServiceProvisioningError:      siteListItemCreatedByUserServiceProvisioningErrorClient,
		SiteListItemDocumentSetVersion:                         siteListItemDocumentSetVersionClient,
		SiteListItemDocumentSetVersionField:                    siteListItemDocumentSetVersionFieldClient,
		SiteListItemDriveItem:                                  siteListItemDriveItemClient,
		SiteListItemDriveItemContent:                           siteListItemDriveItemContentClient,
		SiteListItemField:                                      siteListItemFieldClient,
		SiteListItemLastModifiedByUser:                         siteListItemLastModifiedByUserClient,
		SiteListItemLastModifiedByUserMailboxSetting:           siteListItemLastModifiedByUserMailboxSettingClient,
		SiteListItemLastModifiedByUserServiceProvisioningError: siteListItemLastModifiedByUserServiceProvisioningErrorClient,
		SiteListItemVersion:                                    siteListItemVersionClient,
		SiteListItemVersionField:                               siteListItemVersionFieldClient,
		SiteListLastModifiedByUser:                             siteListLastModifiedByUserClient,
		SiteListLastModifiedByUserMailboxSetting:               siteListLastModifiedByUserMailboxSettingClient,
		SiteListLastModifiedByUserServiceProvisioningError:     siteListLastModifiedByUserServiceProvisioningErrorClient,
		SiteListOperation:                                      siteListOperationClient,
		SiteListSubscription:                                   siteListSubscriptionClient,
		SiteOperation:                                          siteOperationClient,
		SitePage:                                               sitePageClient,
		SitePageCreatedByUser:                                  sitePageCreatedByUserClient,
		SitePageCreatedByUserMailboxSetting:                    sitePageCreatedByUserMailboxSettingClient,
		SitePageCreatedByUserServiceProvisioningError:          sitePageCreatedByUserServiceProvisioningErrorClient,
		SitePageLastModifiedByUser:                             sitePageLastModifiedByUserClient,
		SitePageLastModifiedByUserMailboxSetting:               sitePageLastModifiedByUserMailboxSettingClient,
		SitePageLastModifiedByUserServiceProvisioningError:     sitePageLastModifiedByUserServiceProvisioningErrorClient,
		SitePermission:                                         sitePermissionClient,
		Team:                                                   teamClient,
		TeamAllChannel:                                         teamAllChannelClient,
		TeamChannel:                                            teamChannelClient,
		TeamChannelFilesFolder:                                 teamChannelFilesFolderClient,
		TeamChannelFilesFolderContent:                          teamChannelFilesFolderContentClient,
		TeamChannelMember:                                      teamChannelMemberClient,
		TeamChannelMessage:                                     teamChannelMessageClient,
		TeamChannelMessageHostedContent:                        teamChannelMessageHostedContentClient,
		TeamChannelMessageReply:                                teamChannelMessageReplyClient,
		TeamChannelMessageReplyHostedContent:                   teamChannelMessageReplyHostedContentClient,
		TeamChannelSharedWithTeam:                              teamChannelSharedWithTeamClient,
		TeamChannelSharedWithTeamAllowedMember:                 teamChannelSharedWithTeamAllowedMemberClient,
		TeamChannelSharedWithTeamTeam:                          teamChannelSharedWithTeamTeamClient,
		TeamChannelTab:                                         teamChannelTabClient,
		TeamChannelTabTeamsApp:                                 teamChannelTabTeamsAppClient,
		TeamGroup:                                              teamGroupClient,
		TeamGroupServiceProvisioningError:                      teamGroupServiceProvisioningErrorClient,
		TeamIncomingChannel:                                    teamIncomingChannelClient,
		TeamInstalledApp:                                       teamInstalledAppClient,
		TeamInstalledAppTeamsApp:                               teamInstalledAppTeamsAppClient,
		TeamInstalledAppTeamsAppDefinition:                     teamInstalledAppTeamsAppDefinitionClient,
		TeamMember:                                             teamMemberClient,
		TeamOperation:                                          teamOperationClient,
		TeamPermissionGrant:                                    teamPermissionGrantClient,
		TeamPhoto:                                              teamPhotoClient,
		TeamPrimaryChannel:                                     teamPrimaryChannelClient,
		TeamPrimaryChannelFilesFolder:                          teamPrimaryChannelFilesFolderClient,
		TeamPrimaryChannelFilesFolderContent:                   teamPrimaryChannelFilesFolderContentClient,
		TeamPrimaryChannelMember:                               teamPrimaryChannelMemberClient,
		TeamPrimaryChannelMessage:                              teamPrimaryChannelMessageClient,
		TeamPrimaryChannelMessageHostedContent:                 teamPrimaryChannelMessageHostedContentClient,
		TeamPrimaryChannelMessageReply:                         teamPrimaryChannelMessageReplyClient,
		TeamPrimaryChannelMessageReplyHostedContent:            teamPrimaryChannelMessageReplyHostedContentClient,
		TeamPrimaryChannelSharedWithTeam:                       teamPrimaryChannelSharedWithTeamClient,
		TeamPrimaryChannelSharedWithTeamAllowedMember:          teamPrimaryChannelSharedWithTeamAllowedMemberClient,
		TeamPrimaryChannelSharedWithTeamTeam:                   teamPrimaryChannelSharedWithTeamTeamClient,
		TeamPrimaryChannelTab:                                  teamPrimaryChannelTabClient,
		TeamPrimaryChannelTabTeamsApp:                          teamPrimaryChannelTabTeamsAppClient,
		TeamSchedule:                                           teamScheduleClient,
		TeamScheduleOfferShiftRequest:                          teamScheduleOfferShiftRequestClient,
		TeamScheduleOpenShift:                                  teamScheduleOpenShiftClient,
		TeamScheduleOpenShiftChangeRequest:                     teamScheduleOpenShiftChangeRequestClient,
		TeamScheduleSchedulingGroup:                            teamScheduleSchedulingGroupClient,
		TeamScheduleShift:                                      teamScheduleShiftClient,
		TeamScheduleSwapShiftsChangeRequest:                    teamScheduleSwapShiftsChangeRequestClient,
		TeamScheduleTimeOffReason:                              teamScheduleTimeOffReasonClient,
		TeamScheduleTimeOffRequest:                             teamScheduleTimeOffRequestClient,
		TeamScheduleTimesOff:                                   teamScheduleTimesOffClient,
		TeamTag:                                                teamTagClient,
		TeamTagMember:                                          teamTagMemberClient,
		TeamTemplate:                                           teamTemplateClient,
		Thread:                                                 threadClient,
		ThreadPost:                                             threadPostClient,
		ThreadPostAttachment:                                   threadPostAttachmentClient,
		ThreadPostExtension:                                    threadPostExtensionClient,
		ThreadPostInReplyTo:                                    threadPostInReplyToClient,
		ThreadPostInReplyToAttachment:                          threadPostInReplyToAttachmentClient,
		ThreadPostInReplyToExtension:                           threadPostInReplyToExtensionClient,
		TransitiveMember:                                       transitiveMemberClient,
		TransitiveMemberOf:                                     transitiveMemberOfClient,
	}, nil
}
