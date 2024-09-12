package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/acceptedsender"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/approleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/calendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthread"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpost"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplyto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplytoattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplytoextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostinreplytomention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/conversationthreadpostmention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/createdonbehalfof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivebundle"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivebundlecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivebundlecontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivefollowing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivefollowingcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivefollowingcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalyticsitemactivitystatactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemchildcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemdriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveitemversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemdriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivelistsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driveroot"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalyticsitemactivitystatactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootchild"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootchildcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootchildcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemdriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootlistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootretentionlabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootthumbnail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/driverootversioncontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivespecial"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivespecialcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/drivespecialcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/endpoint"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/event"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/eventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/extension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouplifecyclepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/member"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/memberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/memberswithlicenseerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/permissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/photo"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/planner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/plannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/rejectedsender"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/serviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/site"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticsalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticsitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticsitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticsitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticsitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticsitemactivitystatactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticslastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontentmodel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitedocumentprocessingjob"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitedrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteexternalcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionbitlocker"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionbitlockerrecoverykey"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectiondatalosspreventionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitylabelsublabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionsensitivitypolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionthreatassessmentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionthreatassessmentrequestresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemactivitydriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemanalytics"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdriveitemcontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitempermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepagecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepagecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepagecreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepagelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepagelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepagelastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitepermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebin"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebincreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebincreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebincreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siterecyclebinlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/team"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamallchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelfilesfoldercontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamchanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamgroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamincomingchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teaminstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teaminstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teaminstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teammember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamownermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamownerserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teampermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelfilesfoldercontentstream"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamprimarychanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduledaynote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduleoffershiftrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduleopenshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduleopenshiftchangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduleschedulinggroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduleshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduleshiftsroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduleswapshiftschangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduletimecard"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduletimeoffreason"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduletimeoffrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduletimesoff"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamtag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamtagmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamtemplatedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/thread"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpost"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplyto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplytoattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplytoextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostinreplytomention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/threadpostmention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/transitivemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/transitivememberof"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AcceptedSender                                                   *acceptedsender.AcceptedSenderClient
	AppRoleAssignment                                                *approleassignment.AppRoleAssignmentClient
	Calendar                                                         *calendar.CalendarClient
	CalendarCalendarPermission                                       *calendarcalendarpermission.CalendarCalendarPermissionClient
	CalendarCalendarView                                             *calendarcalendarview.CalendarCalendarViewClient
	CalendarCalendarViewAttachment                                   *calendarcalendarviewattachment.CalendarCalendarViewAttachmentClient
	CalendarCalendarViewCalendar                                     *calendarcalendarviewcalendar.CalendarCalendarViewCalendarClient
	CalendarCalendarViewExceptionOccurrence                          *calendarcalendarviewexceptionoccurrence.CalendarCalendarViewExceptionOccurrenceClient
	CalendarCalendarViewExceptionOccurrenceAttachment                *calendarcalendarviewexceptionoccurrenceattachment.CalendarCalendarViewExceptionOccurrenceAttachmentClient
	CalendarCalendarViewExceptionOccurrenceCalendar                  *calendarcalendarviewexceptionoccurrencecalendar.CalendarCalendarViewExceptionOccurrenceCalendarClient
	CalendarCalendarViewExceptionOccurrenceExtension                 *calendarcalendarviewexceptionoccurrenceextension.CalendarCalendarViewExceptionOccurrenceExtensionClient
	CalendarCalendarViewExceptionOccurrenceInstance                  *calendarcalendarviewexceptionoccurrenceinstance.CalendarCalendarViewExceptionOccurrenceInstanceClient
	CalendarCalendarViewExceptionOccurrenceInstanceAttachment        *calendarcalendarviewexceptionoccurrenceinstanceattachment.CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	CalendarCalendarViewExceptionOccurrenceInstanceCalendar          *calendarcalendarviewexceptionoccurrenceinstancecalendar.CalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	CalendarCalendarViewExceptionOccurrenceInstanceExtension         *calendarcalendarviewexceptionoccurrenceinstanceextension.CalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	CalendarCalendarViewExtension                                    *calendarcalendarviewextension.CalendarCalendarViewExtensionClient
	CalendarCalendarViewInstance                                     *calendarcalendarviewinstance.CalendarCalendarViewInstanceClient
	CalendarCalendarViewInstanceAttachment                           *calendarcalendarviewinstanceattachment.CalendarCalendarViewInstanceAttachmentClient
	CalendarCalendarViewInstanceCalendar                             *calendarcalendarviewinstancecalendar.CalendarCalendarViewInstanceCalendarClient
	CalendarCalendarViewInstanceExceptionOccurrence                  *calendarcalendarviewinstanceexceptionoccurrence.CalendarCalendarViewInstanceExceptionOccurrenceClient
	CalendarCalendarViewInstanceExceptionOccurrenceAttachment        *calendarcalendarviewinstanceexceptionoccurrenceattachment.CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	CalendarCalendarViewInstanceExceptionOccurrenceCalendar          *calendarcalendarviewinstanceexceptionoccurrencecalendar.CalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	CalendarCalendarViewInstanceExceptionOccurrenceExtension         *calendarcalendarviewinstanceexceptionoccurrenceextension.CalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	CalendarCalendarViewInstanceExtension                            *calendarcalendarviewinstanceextension.CalendarCalendarViewInstanceExtensionClient
	CalendarEvent                                                    *calendarevent.CalendarEventClient
	CalendarEventAttachment                                          *calendareventattachment.CalendarEventAttachmentClient
	CalendarEventCalendar                                            *calendareventcalendar.CalendarEventCalendarClient
	CalendarEventExceptionOccurrence                                 *calendareventexceptionoccurrence.CalendarEventExceptionOccurrenceClient
	CalendarEventExceptionOccurrenceAttachment                       *calendareventexceptionoccurrenceattachment.CalendarEventExceptionOccurrenceAttachmentClient
	CalendarEventExceptionOccurrenceCalendar                         *calendareventexceptionoccurrencecalendar.CalendarEventExceptionOccurrenceCalendarClient
	CalendarEventExceptionOccurrenceExtension                        *calendareventexceptionoccurrenceextension.CalendarEventExceptionOccurrenceExtensionClient
	CalendarEventExceptionOccurrenceInstance                         *calendareventexceptionoccurrenceinstance.CalendarEventExceptionOccurrenceInstanceClient
	CalendarEventExceptionOccurrenceInstanceAttachment               *calendareventexceptionoccurrenceinstanceattachment.CalendarEventExceptionOccurrenceInstanceAttachmentClient
	CalendarEventExceptionOccurrenceInstanceCalendar                 *calendareventexceptionoccurrenceinstancecalendar.CalendarEventExceptionOccurrenceInstanceCalendarClient
	CalendarEventExceptionOccurrenceInstanceExtension                *calendareventexceptionoccurrenceinstanceextension.CalendarEventExceptionOccurrenceInstanceExtensionClient
	CalendarEventExtension                                           *calendareventextension.CalendarEventExtensionClient
	CalendarEventInstance                                            *calendareventinstance.CalendarEventInstanceClient
	CalendarEventInstanceAttachment                                  *calendareventinstanceattachment.CalendarEventInstanceAttachmentClient
	CalendarEventInstanceCalendar                                    *calendareventinstancecalendar.CalendarEventInstanceCalendarClient
	CalendarEventInstanceExceptionOccurrence                         *calendareventinstanceexceptionoccurrence.CalendarEventInstanceExceptionOccurrenceClient
	CalendarEventInstanceExceptionOccurrenceAttachment               *calendareventinstanceexceptionoccurrenceattachment.CalendarEventInstanceExceptionOccurrenceAttachmentClient
	CalendarEventInstanceExceptionOccurrenceCalendar                 *calendareventinstanceexceptionoccurrencecalendar.CalendarEventInstanceExceptionOccurrenceCalendarClient
	CalendarEventInstanceExceptionOccurrenceExtension                *calendareventinstanceexceptionoccurrenceextension.CalendarEventInstanceExceptionOccurrenceExtensionClient
	CalendarEventInstanceExtension                                   *calendareventinstanceextension.CalendarEventInstanceExtensionClient
	CalendarView                                                     *calendarview.CalendarViewClient
	CalendarViewAttachment                                           *calendarviewattachment.CalendarViewAttachmentClient
	CalendarViewCalendar                                             *calendarviewcalendar.CalendarViewCalendarClient
	CalendarViewExceptionOccurrence                                  *calendarviewexceptionoccurrence.CalendarViewExceptionOccurrenceClient
	CalendarViewExceptionOccurrenceAttachment                        *calendarviewexceptionoccurrenceattachment.CalendarViewExceptionOccurrenceAttachmentClient
	CalendarViewExceptionOccurrenceCalendar                          *calendarviewexceptionoccurrencecalendar.CalendarViewExceptionOccurrenceCalendarClient
	CalendarViewExceptionOccurrenceExtension                         *calendarviewexceptionoccurrenceextension.CalendarViewExceptionOccurrenceExtensionClient
	CalendarViewExceptionOccurrenceInstance                          *calendarviewexceptionoccurrenceinstance.CalendarViewExceptionOccurrenceInstanceClient
	CalendarViewExceptionOccurrenceInstanceAttachment                *calendarviewexceptionoccurrenceinstanceattachment.CalendarViewExceptionOccurrenceInstanceAttachmentClient
	CalendarViewExceptionOccurrenceInstanceCalendar                  *calendarviewexceptionoccurrenceinstancecalendar.CalendarViewExceptionOccurrenceInstanceCalendarClient
	CalendarViewExceptionOccurrenceInstanceExtension                 *calendarviewexceptionoccurrenceinstanceextension.CalendarViewExceptionOccurrenceInstanceExtensionClient
	CalendarViewExtension                                            *calendarviewextension.CalendarViewExtensionClient
	CalendarViewInstance                                             *calendarviewinstance.CalendarViewInstanceClient
	CalendarViewInstanceAttachment                                   *calendarviewinstanceattachment.CalendarViewInstanceAttachmentClient
	CalendarViewInstanceCalendar                                     *calendarviewinstancecalendar.CalendarViewInstanceCalendarClient
	CalendarViewInstanceExceptionOccurrence                          *calendarviewinstanceexceptionoccurrence.CalendarViewInstanceExceptionOccurrenceClient
	CalendarViewInstanceExceptionOccurrenceAttachment                *calendarviewinstanceexceptionoccurrenceattachment.CalendarViewInstanceExceptionOccurrenceAttachmentClient
	CalendarViewInstanceExceptionOccurrenceCalendar                  *calendarviewinstanceexceptionoccurrencecalendar.CalendarViewInstanceExceptionOccurrenceCalendarClient
	CalendarViewInstanceExceptionOccurrenceExtension                 *calendarviewinstanceexceptionoccurrenceextension.CalendarViewInstanceExceptionOccurrenceExtensionClient
	CalendarViewInstanceExtension                                    *calendarviewinstanceextension.CalendarViewInstanceExtensionClient
	Conversation                                                     *conversation.ConversationClient
	ConversationThread                                               *conversationthread.ConversationThreadClient
	ConversationThreadPost                                           *conversationthreadpost.ConversationThreadPostClient
	ConversationThreadPostAttachment                                 *conversationthreadpostattachment.ConversationThreadPostAttachmentClient
	ConversationThreadPostExtension                                  *conversationthreadpostextension.ConversationThreadPostExtensionClient
	ConversationThreadPostInReplyTo                                  *conversationthreadpostinreplyto.ConversationThreadPostInReplyToClient
	ConversationThreadPostInReplyToAttachment                        *conversationthreadpostinreplytoattachment.ConversationThreadPostInReplyToAttachmentClient
	ConversationThreadPostInReplyToExtension                         *conversationthreadpostinreplytoextension.ConversationThreadPostInReplyToExtensionClient
	ConversationThreadPostInReplyToMention                           *conversationthreadpostinreplytomention.ConversationThreadPostInReplyToMentionClient
	ConversationThreadPostMention                                    *conversationthreadpostmention.ConversationThreadPostMentionClient
	CreatedOnBehalfOf                                                *createdonbehalfof.CreatedOnBehalfOfClient
	Drive                                                            *drive.DriveClient
	DriveActivity                                                    *driveactivity.DriveActivityClient
	DriveActivityDriveItem                                           *driveactivitydriveitem.DriveActivityDriveItemClient
	DriveActivityDriveItemContent                                    *driveactivitydriveitemcontent.DriveActivityDriveItemContentClient
	DriveActivityDriveItemContentStream                              *driveactivitydriveitemcontentstream.DriveActivityDriveItemContentStreamClient
	DriveActivityListItem                                            *driveactivitylistitem.DriveActivityListItemClient
	DriveBundle                                                      *drivebundle.DriveBundleClient
	DriveBundleContent                                               *drivebundlecontent.DriveBundleContentClient
	DriveBundleContentStream                                         *drivebundlecontentstream.DriveBundleContentStreamClient
	DriveCreatedByUser                                               *drivecreatedbyuser.DriveCreatedByUserClient
	DriveCreatedByUserMailboxSetting                                 *drivecreatedbyusermailboxsetting.DriveCreatedByUserMailboxSettingClient
	DriveCreatedByUserServiceProvisioningError                       *drivecreatedbyuserserviceprovisioningerror.DriveCreatedByUserServiceProvisioningErrorClient
	DriveFollowing                                                   *drivefollowing.DriveFollowingClient
	DriveFollowingContent                                            *drivefollowingcontent.DriveFollowingContentClient
	DriveFollowingContentStream                                      *drivefollowingcontentstream.DriveFollowingContentStreamClient
	DriveItem                                                        *driveitem.DriveItemClient
	DriveItemActivity                                                *driveitemactivity.DriveItemActivityClient
	DriveItemAnalytics                                               *driveitemanalytics.DriveItemAnalyticsClient
	DriveItemAnalyticsAllTime                                        *driveitemanalyticsalltime.DriveItemAnalyticsAllTimeClient
	DriveItemAnalyticsItemActivityStat                               *driveitemanalyticsitemactivitystat.DriveItemAnalyticsItemActivityStatClient
	DriveItemAnalyticsItemActivityStatActivity                       *driveitemanalyticsitemactivitystatactivity.DriveItemAnalyticsItemActivityStatActivityClient
	DriveItemAnalyticsItemActivityStatActivityDriveItem              *driveitemanalyticsitemactivitystatactivitydriveitem.DriveItemAnalyticsItemActivityStatActivityDriveItemClient
	DriveItemAnalyticsItemActivityStatActivityDriveItemContent       *driveitemanalyticsitemactivitystatactivitydriveitemcontent.DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient
	DriveItemAnalyticsItemActivityStatActivityDriveItemContentStream *driveitemanalyticsitemactivitystatactivitydriveitemcontentstream.DriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient
	DriveItemAnalyticsLastSevenDay                                   *driveitemanalyticslastsevenday.DriveItemAnalyticsLastSevenDayClient
	DriveItemChild                                                   *driveitemchild.DriveItemChildClient
	DriveItemChildContent                                            *driveitemchildcontent.DriveItemChildContentClient
	DriveItemChildContentStream                                      *driveitemchildcontentstream.DriveItemChildContentStreamClient
	DriveItemContent                                                 *driveitemcontent.DriveItemContentClient
	DriveItemContentStream                                           *driveitemcontentstream.DriveItemContentStreamClient
	DriveItemCreatedByUser                                           *driveitemcreatedbyuser.DriveItemCreatedByUserClient
	DriveItemCreatedByUserMailboxSetting                             *driveitemcreatedbyusermailboxsetting.DriveItemCreatedByUserMailboxSettingClient
	DriveItemCreatedByUserServiceProvisioningError                   *driveitemcreatedbyuserserviceprovisioningerror.DriveItemCreatedByUserServiceProvisioningErrorClient
	DriveItemLastModifiedByUser                                      *driveitemlastmodifiedbyuser.DriveItemLastModifiedByUserClient
	DriveItemLastModifiedByUserMailboxSetting                        *driveitemlastmodifiedbyusermailboxsetting.DriveItemLastModifiedByUserMailboxSettingClient
	DriveItemLastModifiedByUserServiceProvisioningError              *driveitemlastmodifiedbyuserserviceprovisioningerror.DriveItemLastModifiedByUserServiceProvisioningErrorClient
	DriveItemListItem                                                *driveitemlistitem.DriveItemListItemClient
	DriveItemListItemActivity                                        *driveitemlistitemactivity.DriveItemListItemActivityClient
	DriveItemListItemActivityDriveItem                               *driveitemlistitemactivitydriveitem.DriveItemListItemActivityDriveItemClient
	DriveItemListItemActivityDriveItemContent                        *driveitemlistitemactivitydriveitemcontent.DriveItemListItemActivityDriveItemContentClient
	DriveItemListItemActivityDriveItemContentStream                  *driveitemlistitemactivitydriveitemcontentstream.DriveItemListItemActivityDriveItemContentStreamClient
	DriveItemListItemActivityListItem                                *driveitemlistitemactivitylistitem.DriveItemListItemActivityListItemClient
	DriveItemListItemAnalytics                                       *driveitemlistitemanalytics.DriveItemListItemAnalyticsClient
	DriveItemListItemCreatedByUser                                   *driveitemlistitemcreatedbyuser.DriveItemListItemCreatedByUserClient
	DriveItemListItemCreatedByUserMailboxSetting                     *driveitemlistitemcreatedbyusermailboxsetting.DriveItemListItemCreatedByUserMailboxSettingClient
	DriveItemListItemCreatedByUserServiceProvisioningError           *driveitemlistitemcreatedbyuserserviceprovisioningerror.DriveItemListItemCreatedByUserServiceProvisioningErrorClient
	DriveItemListItemDocumentSetVersion                              *driveitemlistitemdocumentsetversion.DriveItemListItemDocumentSetVersionClient
	DriveItemListItemDocumentSetVersionField                         *driveitemlistitemdocumentsetversionfield.DriveItemListItemDocumentSetVersionFieldClient
	DriveItemListItemDriveItem                                       *driveitemlistitemdriveitem.DriveItemListItemDriveItemClient
	DriveItemListItemDriveItemContent                                *driveitemlistitemdriveitemcontent.DriveItemListItemDriveItemContentClient
	DriveItemListItemDriveItemContentStream                          *driveitemlistitemdriveitemcontentstream.DriveItemListItemDriveItemContentStreamClient
	DriveItemListItemField                                           *driveitemlistitemfield.DriveItemListItemFieldClient
	DriveItemListItemLastModifiedByUser                              *driveitemlistitemlastmodifiedbyuser.DriveItemListItemLastModifiedByUserClient
	DriveItemListItemLastModifiedByUserMailboxSetting                *driveitemlistitemlastmodifiedbyusermailboxsetting.DriveItemListItemLastModifiedByUserMailboxSettingClient
	DriveItemListItemLastModifiedByUserServiceProvisioningError      *driveitemlistitemlastmodifiedbyuserserviceprovisioningerror.DriveItemListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveItemListItemPermission                                      *driveitemlistitempermission.DriveItemListItemPermissionClient
	DriveItemListItemVersion                                         *driveitemlistitemversion.DriveItemListItemVersionClient
	DriveItemListItemVersionField                                    *driveitemlistitemversionfield.DriveItemListItemVersionFieldClient
	DriveItemPermission                                              *driveitempermission.DriveItemPermissionClient
	DriveItemRetentionLabel                                          *driveitemretentionlabel.DriveItemRetentionLabelClient
	DriveItemSubscription                                            *driveitemsubscription.DriveItemSubscriptionClient
	DriveItemThumbnail                                               *driveitemthumbnail.DriveItemThumbnailClient
	DriveItemVersion                                                 *driveitemversion.DriveItemVersionClient
	DriveItemVersionContent                                          *driveitemversioncontent.DriveItemVersionContentClient
	DriveLastModifiedByUser                                          *drivelastmodifiedbyuser.DriveLastModifiedByUserClient
	DriveLastModifiedByUserMailboxSetting                            *drivelastmodifiedbyusermailboxsetting.DriveLastModifiedByUserMailboxSettingClient
	DriveLastModifiedByUserServiceProvisioningError                  *drivelastmodifiedbyuserserviceprovisioningerror.DriveLastModifiedByUserServiceProvisioningErrorClient
	DriveList                                                        *drivelist.DriveListClient
	DriveListActivity                                                *drivelistactivity.DriveListActivityClient
	DriveListColumn                                                  *drivelistcolumn.DriveListColumnClient
	DriveListColumnSourceColumn                                      *drivelistcolumnsourcecolumn.DriveListColumnSourceColumnClient
	DriveListContentType                                             *drivelistcontenttype.DriveListContentTypeClient
	DriveListContentTypeBase                                         *drivelistcontenttypebase.DriveListContentTypeBaseClient
	DriveListContentTypeBaseType                                     *drivelistcontenttypebasetype.DriveListContentTypeBaseTypeClient
	DriveListContentTypeColumn                                       *drivelistcontenttypecolumn.DriveListContentTypeColumnClient
	DriveListContentTypeColumnLink                                   *drivelistcontenttypecolumnlink.DriveListContentTypeColumnLinkClient
	DriveListContentTypeColumnPosition                               *drivelistcontenttypecolumnposition.DriveListContentTypeColumnPositionClient
	DriveListContentTypeColumnSourceColumn                           *drivelistcontenttypecolumnsourcecolumn.DriveListContentTypeColumnSourceColumnClient
	DriveListCreatedByUser                                           *drivelistcreatedbyuser.DriveListCreatedByUserClient
	DriveListCreatedByUserMailboxSetting                             *drivelistcreatedbyusermailboxsetting.DriveListCreatedByUserMailboxSettingClient
	DriveListCreatedByUserServiceProvisioningError                   *drivelistcreatedbyuserserviceprovisioningerror.DriveListCreatedByUserServiceProvisioningErrorClient
	DriveListDrive                                                   *drivelistdrive.DriveListDriveClient
	DriveListItem                                                    *drivelistitem.DriveListItemClient
	DriveListItemActivity                                            *drivelistitemactivity.DriveListItemActivityClient
	DriveListItemActivityDriveItem                                   *drivelistitemactivitydriveitem.DriveListItemActivityDriveItemClient
	DriveListItemActivityDriveItemContent                            *drivelistitemactivitydriveitemcontent.DriveListItemActivityDriveItemContentClient
	DriveListItemActivityDriveItemContentStream                      *drivelistitemactivitydriveitemcontentstream.DriveListItemActivityDriveItemContentStreamClient
	DriveListItemActivityListItem                                    *drivelistitemactivitylistitem.DriveListItemActivityListItemClient
	DriveListItemAnalytics                                           *drivelistitemanalytics.DriveListItemAnalyticsClient
	DriveListItemCreatedByUser                                       *drivelistitemcreatedbyuser.DriveListItemCreatedByUserClient
	DriveListItemCreatedByUserMailboxSetting                         *drivelistitemcreatedbyusermailboxsetting.DriveListItemCreatedByUserMailboxSettingClient
	DriveListItemCreatedByUserServiceProvisioningError               *drivelistitemcreatedbyuserserviceprovisioningerror.DriveListItemCreatedByUserServiceProvisioningErrorClient
	DriveListItemDocumentSetVersion                                  *drivelistitemdocumentsetversion.DriveListItemDocumentSetVersionClient
	DriveListItemDocumentSetVersionField                             *drivelistitemdocumentsetversionfield.DriveListItemDocumentSetVersionFieldClient
	DriveListItemDriveItem                                           *drivelistitemdriveitem.DriveListItemDriveItemClient
	DriveListItemDriveItemContent                                    *drivelistitemdriveitemcontent.DriveListItemDriveItemContentClient
	DriveListItemDriveItemContentStream                              *drivelistitemdriveitemcontentstream.DriveListItemDriveItemContentStreamClient
	DriveListItemField                                               *drivelistitemfield.DriveListItemFieldClient
	DriveListItemLastModifiedByUser                                  *drivelistitemlastmodifiedbyuser.DriveListItemLastModifiedByUserClient
	DriveListItemLastModifiedByUserMailboxSetting                    *drivelistitemlastmodifiedbyusermailboxsetting.DriveListItemLastModifiedByUserMailboxSettingClient
	DriveListItemLastModifiedByUserServiceProvisioningError          *drivelistitemlastmodifiedbyuserserviceprovisioningerror.DriveListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveListItemPermission                                          *drivelistitempermission.DriveListItemPermissionClient
	DriveListItemVersion                                             *drivelistitemversion.DriveListItemVersionClient
	DriveListItemVersionField                                        *drivelistitemversionfield.DriveListItemVersionFieldClient
	DriveListLastModifiedByUser                                      *drivelistlastmodifiedbyuser.DriveListLastModifiedByUserClient
	DriveListLastModifiedByUserMailboxSetting                        *drivelistlastmodifiedbyusermailboxsetting.DriveListLastModifiedByUserMailboxSettingClient
	DriveListLastModifiedByUserServiceProvisioningError              *drivelistlastmodifiedbyuserserviceprovisioningerror.DriveListLastModifiedByUserServiceProvisioningErrorClient
	DriveListOperation                                               *drivelistoperation.DriveListOperationClient
	DriveListPermission                                              *drivelistpermission.DriveListPermissionClient
	DriveListSubscription                                            *drivelistsubscription.DriveListSubscriptionClient
	DriveRoot                                                        *driveroot.DriveRootClient
	DriveRootActivity                                                *driverootactivity.DriveRootActivityClient
	DriveRootAnalytics                                               *driverootanalytics.DriveRootAnalyticsClient
	DriveRootAnalyticsAllTime                                        *driverootanalyticsalltime.DriveRootAnalyticsAllTimeClient
	DriveRootAnalyticsItemActivityStat                               *driverootanalyticsitemactivitystat.DriveRootAnalyticsItemActivityStatClient
	DriveRootAnalyticsItemActivityStatActivity                       *driverootanalyticsitemactivitystatactivity.DriveRootAnalyticsItemActivityStatActivityClient
	DriveRootAnalyticsItemActivityStatActivityDriveItem              *driverootanalyticsitemactivitystatactivitydriveitem.DriveRootAnalyticsItemActivityStatActivityDriveItemClient
	DriveRootAnalyticsItemActivityStatActivityDriveItemContent       *driverootanalyticsitemactivitystatactivitydriveitemcontent.DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient
	DriveRootAnalyticsItemActivityStatActivityDriveItemContentStream *driverootanalyticsitemactivitystatactivitydriveitemcontentstream.DriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClient
	DriveRootAnalyticsLastSevenDay                                   *driverootanalyticslastsevenday.DriveRootAnalyticsLastSevenDayClient
	DriveRootChild                                                   *driverootchild.DriveRootChildClient
	DriveRootChildContent                                            *driverootchildcontent.DriveRootChildContentClient
	DriveRootChildContentStream                                      *driverootchildcontentstream.DriveRootChildContentStreamClient
	DriveRootContent                                                 *driverootcontent.DriveRootContentClient
	DriveRootContentStream                                           *driverootcontentstream.DriveRootContentStreamClient
	DriveRootCreatedByUser                                           *driverootcreatedbyuser.DriveRootCreatedByUserClient
	DriveRootCreatedByUserMailboxSetting                             *driverootcreatedbyusermailboxsetting.DriveRootCreatedByUserMailboxSettingClient
	DriveRootCreatedByUserServiceProvisioningError                   *driverootcreatedbyuserserviceprovisioningerror.DriveRootCreatedByUserServiceProvisioningErrorClient
	DriveRootLastModifiedByUser                                      *driverootlastmodifiedbyuser.DriveRootLastModifiedByUserClient
	DriveRootLastModifiedByUserMailboxSetting                        *driverootlastmodifiedbyusermailboxsetting.DriveRootLastModifiedByUserMailboxSettingClient
	DriveRootLastModifiedByUserServiceProvisioningError              *driverootlastmodifiedbyuserserviceprovisioningerror.DriveRootLastModifiedByUserServiceProvisioningErrorClient
	DriveRootListItem                                                *driverootlistitem.DriveRootListItemClient
	DriveRootListItemActivity                                        *driverootlistitemactivity.DriveRootListItemActivityClient
	DriveRootListItemActivityDriveItem                               *driverootlistitemactivitydriveitem.DriveRootListItemActivityDriveItemClient
	DriveRootListItemActivityDriveItemContent                        *driverootlistitemactivitydriveitemcontent.DriveRootListItemActivityDriveItemContentClient
	DriveRootListItemActivityDriveItemContentStream                  *driverootlistitemactivitydriveitemcontentstream.DriveRootListItemActivityDriveItemContentStreamClient
	DriveRootListItemActivityListItem                                *driverootlistitemactivitylistitem.DriveRootListItemActivityListItemClient
	DriveRootListItemAnalytics                                       *driverootlistitemanalytics.DriveRootListItemAnalyticsClient
	DriveRootListItemCreatedByUser                                   *driverootlistitemcreatedbyuser.DriveRootListItemCreatedByUserClient
	DriveRootListItemCreatedByUserMailboxSetting                     *driverootlistitemcreatedbyusermailboxsetting.DriveRootListItemCreatedByUserMailboxSettingClient
	DriveRootListItemCreatedByUserServiceProvisioningError           *driverootlistitemcreatedbyuserserviceprovisioningerror.DriveRootListItemCreatedByUserServiceProvisioningErrorClient
	DriveRootListItemDocumentSetVersion                              *driverootlistitemdocumentsetversion.DriveRootListItemDocumentSetVersionClient
	DriveRootListItemDocumentSetVersionField                         *driverootlistitemdocumentsetversionfield.DriveRootListItemDocumentSetVersionFieldClient
	DriveRootListItemDriveItem                                       *driverootlistitemdriveitem.DriveRootListItemDriveItemClient
	DriveRootListItemDriveItemContent                                *driverootlistitemdriveitemcontent.DriveRootListItemDriveItemContentClient
	DriveRootListItemDriveItemContentStream                          *driverootlistitemdriveitemcontentstream.DriveRootListItemDriveItemContentStreamClient
	DriveRootListItemField                                           *driverootlistitemfield.DriveRootListItemFieldClient
	DriveRootListItemLastModifiedByUser                              *driverootlistitemlastmodifiedbyuser.DriveRootListItemLastModifiedByUserClient
	DriveRootListItemLastModifiedByUserMailboxSetting                *driverootlistitemlastmodifiedbyusermailboxsetting.DriveRootListItemLastModifiedByUserMailboxSettingClient
	DriveRootListItemLastModifiedByUserServiceProvisioningError      *driverootlistitemlastmodifiedbyuserserviceprovisioningerror.DriveRootListItemLastModifiedByUserServiceProvisioningErrorClient
	DriveRootListItemPermission                                      *driverootlistitempermission.DriveRootListItemPermissionClient
	DriveRootListItemVersion                                         *driverootlistitemversion.DriveRootListItemVersionClient
	DriveRootListItemVersionField                                    *driverootlistitemversionfield.DriveRootListItemVersionFieldClient
	DriveRootPermission                                              *driverootpermission.DriveRootPermissionClient
	DriveRootRetentionLabel                                          *driverootretentionlabel.DriveRootRetentionLabelClient
	DriveRootSubscription                                            *driverootsubscription.DriveRootSubscriptionClient
	DriveRootThumbnail                                               *driverootthumbnail.DriveRootThumbnailClient
	DriveRootVersion                                                 *driverootversion.DriveRootVersionClient
	DriveRootVersionContent                                          *driverootversioncontent.DriveRootVersionContentClient
	DriveSpecial                                                     *drivespecial.DriveSpecialClient
	DriveSpecialContent                                              *drivespecialcontent.DriveSpecialContentClient
	DriveSpecialContentStream                                        *drivespecialcontentstream.DriveSpecialContentStreamClient
	Endpoint                                                         *endpoint.EndpointClient
	Event                                                            *event.EventClient
	EventAttachment                                                  *eventattachment.EventAttachmentClient
	EventCalendar                                                    *eventcalendar.EventCalendarClient
	EventExceptionOccurrence                                         *eventexceptionoccurrence.EventExceptionOccurrenceClient
	EventExceptionOccurrenceAttachment                               *eventexceptionoccurrenceattachment.EventExceptionOccurrenceAttachmentClient
	EventExceptionOccurrenceCalendar                                 *eventexceptionoccurrencecalendar.EventExceptionOccurrenceCalendarClient
	EventExceptionOccurrenceExtension                                *eventexceptionoccurrenceextension.EventExceptionOccurrenceExtensionClient
	EventExceptionOccurrenceInstance                                 *eventexceptionoccurrenceinstance.EventExceptionOccurrenceInstanceClient
	EventExceptionOccurrenceInstanceAttachment                       *eventexceptionoccurrenceinstanceattachment.EventExceptionOccurrenceInstanceAttachmentClient
	EventExceptionOccurrenceInstanceCalendar                         *eventexceptionoccurrenceinstancecalendar.EventExceptionOccurrenceInstanceCalendarClient
	EventExceptionOccurrenceInstanceExtension                        *eventexceptionoccurrenceinstanceextension.EventExceptionOccurrenceInstanceExtensionClient
	EventExtension                                                   *eventextension.EventExtensionClient
	EventInstance                                                    *eventinstance.EventInstanceClient
	EventInstanceAttachment                                          *eventinstanceattachment.EventInstanceAttachmentClient
	EventInstanceCalendar                                            *eventinstancecalendar.EventInstanceCalendarClient
	EventInstanceExceptionOccurrence                                 *eventinstanceexceptionoccurrence.EventInstanceExceptionOccurrenceClient
	EventInstanceExceptionOccurrenceAttachment                       *eventinstanceexceptionoccurrenceattachment.EventInstanceExceptionOccurrenceAttachmentClient
	EventInstanceExceptionOccurrenceCalendar                         *eventinstanceexceptionoccurrencecalendar.EventInstanceExceptionOccurrenceCalendarClient
	EventInstanceExceptionOccurrenceExtension                        *eventinstanceexceptionoccurrenceextension.EventInstanceExceptionOccurrenceExtensionClient
	EventInstanceExtension                                           *eventinstanceextension.EventInstanceExtensionClient
	Extension                                                        *extension.ExtensionClient
	Group                                                            *group.GroupClient
	GroupLifecyclePolicy                                             *grouplifecyclepolicy.GroupLifecyclePolicyClient
	Member                                                           *member.MemberClient
	MemberOf                                                         *memberof.MemberOfClient
	MembersWithLicenseError                                          *memberswithlicenseerror.MembersWithLicenseErrorClient
	Owner                                                            *owner.OwnerClient
	PermissionGrant                                                  *permissiongrant.PermissionGrantClient
	Photo                                                            *photo.PhotoClient
	Planner                                                          *planner.PlannerClient
	PlannerPlan                                                      *plannerplan.PlannerPlanClient
	PlannerPlanBucket                                                *plannerplanbucket.PlannerPlanBucketClient
	PlannerPlanBucketTask                                            *plannerplanbuckettask.PlannerPlanBucketTaskClient
	PlannerPlanBucketTaskAssignedToTaskBoardFormat                   *plannerplanbuckettaskassignedtotaskboardformat.PlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	PlannerPlanBucketTaskBucketTaskBoardFormat                       *plannerplanbuckettaskbuckettaskboardformat.PlannerPlanBucketTaskBucketTaskBoardFormatClient
	PlannerPlanBucketTaskDetail                                      *plannerplanbuckettaskdetail.PlannerPlanBucketTaskDetailClient
	PlannerPlanBucketTaskProgressTaskBoardFormat                     *plannerplanbuckettaskprogresstaskboardformat.PlannerPlanBucketTaskProgressTaskBoardFormatClient
	PlannerPlanDetail                                                *plannerplandetail.PlannerPlanDetailClient
	PlannerPlanTask                                                  *plannerplantask.PlannerPlanTaskClient
	PlannerPlanTaskAssignedToTaskBoardFormat                         *plannerplantaskassignedtotaskboardformat.PlannerPlanTaskAssignedToTaskBoardFormatClient
	PlannerPlanTaskBucketTaskBoardFormat                             *plannerplantaskbuckettaskboardformat.PlannerPlanTaskBucketTaskBoardFormatClient
	PlannerPlanTaskDetail                                            *plannerplantaskdetail.PlannerPlanTaskDetailClient
	PlannerPlanTaskProgressTaskBoardFormat                           *plannerplantaskprogresstaskboardformat.PlannerPlanTaskProgressTaskBoardFormatClient
	RejectedSender                                                   *rejectedsender.RejectedSenderClient
	ServiceProvisioningError                                         *serviceprovisioningerror.ServiceProvisioningErrorClient
	Setting                                                          *setting.SettingClient
	Site                                                             *site.SiteClient
	SiteAnalytics                                                    *siteanalytics.SiteAnalyticsClient
	SiteAnalyticsAllTime                                             *siteanalyticsalltime.SiteAnalyticsAllTimeClient
	SiteAnalyticsItemActivityStat                                    *siteanalyticsitemactivitystat.SiteAnalyticsItemActivityStatClient
	SiteAnalyticsItemActivityStatActivity                            *siteanalyticsitemactivitystatactivity.SiteAnalyticsItemActivityStatActivityClient
	SiteAnalyticsItemActivityStatActivityDriveItem                   *siteanalyticsitemactivitystatactivitydriveitem.SiteAnalyticsItemActivityStatActivityDriveItemClient
	SiteAnalyticsItemActivityStatActivityDriveItemContent            *siteanalyticsitemactivitystatactivitydriveitemcontent.SiteAnalyticsItemActivityStatActivityDriveItemContentClient
	SiteAnalyticsItemActivityStatActivityDriveItemContentStream      *siteanalyticsitemactivitystatactivitydriveitemcontentstream.SiteAnalyticsItemActivityStatActivityDriveItemContentStreamClient
	SiteAnalyticsLastSevenDay                                        *siteanalyticslastsevenday.SiteAnalyticsLastSevenDayClient
	SiteColumn                                                       *sitecolumn.SiteColumnClient
	SiteColumnSourceColumn                                           *sitecolumnsourcecolumn.SiteColumnSourceColumnClient
	SiteContentModel                                                 *sitecontentmodel.SiteContentModelClient
	SiteContentType                                                  *sitecontenttype.SiteContentTypeClient
	SiteContentTypeBase                                              *sitecontenttypebase.SiteContentTypeBaseClient
	SiteContentTypeBaseType                                          *sitecontenttypebasetype.SiteContentTypeBaseTypeClient
	SiteContentTypeColumn                                            *sitecontenttypecolumn.SiteContentTypeColumnClient
	SiteContentTypeColumnLink                                        *sitecontenttypecolumnlink.SiteContentTypeColumnLinkClient
	SiteContentTypeColumnPosition                                    *sitecontenttypecolumnposition.SiteContentTypeColumnPositionClient
	SiteContentTypeColumnSourceColumn                                *sitecontenttypecolumnsourcecolumn.SiteContentTypeColumnSourceColumnClient
	SiteCreatedByUser                                                *sitecreatedbyuser.SiteCreatedByUserClient
	SiteCreatedByUserMailboxSetting                                  *sitecreatedbyusermailboxsetting.SiteCreatedByUserMailboxSettingClient
	SiteCreatedByUserServiceProvisioningError                        *sitecreatedbyuserserviceprovisioningerror.SiteCreatedByUserServiceProvisioningErrorClient
	SiteDocumentProcessingJob                                        *sitedocumentprocessingjob.SiteDocumentProcessingJobClient
	SiteDrive                                                        *sitedrive.SiteDriveClient
	SiteExternalColumn                                               *siteexternalcolumn.SiteExternalColumnClient
	SiteInformationProtection                                        *siteinformationprotection.SiteInformationProtectionClient
	SiteInformationProtectionBitlocker                               *siteinformationprotectionbitlocker.SiteInformationProtectionBitlockerClient
	SiteInformationProtectionBitlockerRecoveryKey                    *siteinformationprotectionbitlockerrecoverykey.SiteInformationProtectionBitlockerRecoveryKeyClient
	SiteInformationProtectionDataLossPreventionPolicy                *siteinformationprotectiondatalosspreventionpolicy.SiteInformationProtectionDataLossPreventionPolicyClient
	SiteInformationProtectionPolicy                                  *siteinformationprotectionpolicy.SiteInformationProtectionPolicyClient
	SiteInformationProtectionPolicyLabel                             *siteinformationprotectionpolicylabel.SiteInformationProtectionPolicyLabelClient
	SiteInformationProtectionSensitivityLabel                        *siteinformationprotectionsensitivitylabel.SiteInformationProtectionSensitivityLabelClient
	SiteInformationProtectionSensitivityLabelSublabel                *siteinformationprotectionsensitivitylabelsublabel.SiteInformationProtectionSensitivityLabelSublabelClient
	SiteInformationProtectionSensitivityPolicySetting                *siteinformationprotectionsensitivitypolicysetting.SiteInformationProtectionSensitivityPolicySettingClient
	SiteInformationProtectionThreatAssessmentRequest                 *siteinformationprotectionthreatassessmentrequest.SiteInformationProtectionThreatAssessmentRequestClient
	SiteInformationProtectionThreatAssessmentRequestResult           *siteinformationprotectionthreatassessmentrequestresult.SiteInformationProtectionThreatAssessmentRequestResultClient
	SiteItem                                                         *siteitem.SiteItemClient
	SiteLastModifiedByUser                                           *sitelastmodifiedbyuser.SiteLastModifiedByUserClient
	SiteLastModifiedByUserMailboxSetting                             *sitelastmodifiedbyusermailboxsetting.SiteLastModifiedByUserMailboxSettingClient
	SiteLastModifiedByUserServiceProvisioningError                   *sitelastmodifiedbyuserserviceprovisioningerror.SiteLastModifiedByUserServiceProvisioningErrorClient
	SiteList                                                         *sitelist.SiteListClient
	SiteListActivity                                                 *sitelistactivity.SiteListActivityClient
	SiteListColumn                                                   *sitelistcolumn.SiteListColumnClient
	SiteListColumnSourceColumn                                       *sitelistcolumnsourcecolumn.SiteListColumnSourceColumnClient
	SiteListContentType                                              *sitelistcontenttype.SiteListContentTypeClient
	SiteListContentTypeBase                                          *sitelistcontenttypebase.SiteListContentTypeBaseClient
	SiteListContentTypeBaseType                                      *sitelistcontenttypebasetype.SiteListContentTypeBaseTypeClient
	SiteListContentTypeColumn                                        *sitelistcontenttypecolumn.SiteListContentTypeColumnClient
	SiteListContentTypeColumnLink                                    *sitelistcontenttypecolumnlink.SiteListContentTypeColumnLinkClient
	SiteListContentTypeColumnPosition                                *sitelistcontenttypecolumnposition.SiteListContentTypeColumnPositionClient
	SiteListContentTypeColumnSourceColumn                            *sitelistcontenttypecolumnsourcecolumn.SiteListContentTypeColumnSourceColumnClient
	SiteListCreatedByUser                                            *sitelistcreatedbyuser.SiteListCreatedByUserClient
	SiteListCreatedByUserMailboxSetting                              *sitelistcreatedbyusermailboxsetting.SiteListCreatedByUserMailboxSettingClient
	SiteListCreatedByUserServiceProvisioningError                    *sitelistcreatedbyuserserviceprovisioningerror.SiteListCreatedByUserServiceProvisioningErrorClient
	SiteListDrive                                                    *sitelistdrive.SiteListDriveClient
	SiteListItem                                                     *sitelistitem.SiteListItemClient
	SiteListItemActivity                                             *sitelistitemactivity.SiteListItemActivityClient
	SiteListItemActivityDriveItem                                    *sitelistitemactivitydriveitem.SiteListItemActivityDriveItemClient
	SiteListItemActivityDriveItemContent                             *sitelistitemactivitydriveitemcontent.SiteListItemActivityDriveItemContentClient
	SiteListItemActivityDriveItemContentStream                       *sitelistitemactivitydriveitemcontentstream.SiteListItemActivityDriveItemContentStreamClient
	SiteListItemActivityListItem                                     *sitelistitemactivitylistitem.SiteListItemActivityListItemClient
	SiteListItemAnalytics                                            *sitelistitemanalytics.SiteListItemAnalyticsClient
	SiteListItemCreatedByUser                                        *sitelistitemcreatedbyuser.SiteListItemCreatedByUserClient
	SiteListItemCreatedByUserMailboxSetting                          *sitelistitemcreatedbyusermailboxsetting.SiteListItemCreatedByUserMailboxSettingClient
	SiteListItemCreatedByUserServiceProvisioningError                *sitelistitemcreatedbyuserserviceprovisioningerror.SiteListItemCreatedByUserServiceProvisioningErrorClient
	SiteListItemDocumentSetVersion                                   *sitelistitemdocumentsetversion.SiteListItemDocumentSetVersionClient
	SiteListItemDocumentSetVersionField                              *sitelistitemdocumentsetversionfield.SiteListItemDocumentSetVersionFieldClient
	SiteListItemDriveItem                                            *sitelistitemdriveitem.SiteListItemDriveItemClient
	SiteListItemDriveItemContent                                     *sitelistitemdriveitemcontent.SiteListItemDriveItemContentClient
	SiteListItemDriveItemContentStream                               *sitelistitemdriveitemcontentstream.SiteListItemDriveItemContentStreamClient
	SiteListItemField                                                *sitelistitemfield.SiteListItemFieldClient
	SiteListItemLastModifiedByUser                                   *sitelistitemlastmodifiedbyuser.SiteListItemLastModifiedByUserClient
	SiteListItemLastModifiedByUserMailboxSetting                     *sitelistitemlastmodifiedbyusermailboxsetting.SiteListItemLastModifiedByUserMailboxSettingClient
	SiteListItemLastModifiedByUserServiceProvisioningError           *sitelistitemlastmodifiedbyuserserviceprovisioningerror.SiteListItemLastModifiedByUserServiceProvisioningErrorClient
	SiteListItemPermission                                           *sitelistitempermission.SiteListItemPermissionClient
	SiteListItemVersion                                              *sitelistitemversion.SiteListItemVersionClient
	SiteListItemVersionField                                         *sitelistitemversionfield.SiteListItemVersionFieldClient
	SiteListLastModifiedByUser                                       *sitelistlastmodifiedbyuser.SiteListLastModifiedByUserClient
	SiteListLastModifiedByUserMailboxSetting                         *sitelistlastmodifiedbyusermailboxsetting.SiteListLastModifiedByUserMailboxSettingClient
	SiteListLastModifiedByUserServiceProvisioningError               *sitelistlastmodifiedbyuserserviceprovisioningerror.SiteListLastModifiedByUserServiceProvisioningErrorClient
	SiteListOperation                                                *sitelistoperation.SiteListOperationClient
	SiteListPermission                                               *sitelistpermission.SiteListPermissionClient
	SiteListSubscription                                             *sitelistsubscription.SiteListSubscriptionClient
	SiteOperation                                                    *siteoperation.SiteOperationClient
	SitePage                                                         *sitepage.SitePageClient
	SitePageCreatedByUser                                            *sitepagecreatedbyuser.SitePageCreatedByUserClient
	SitePageCreatedByUserMailboxSetting                              *sitepagecreatedbyusermailboxsetting.SitePageCreatedByUserMailboxSettingClient
	SitePageCreatedByUserServiceProvisioningError                    *sitepagecreatedbyuserserviceprovisioningerror.SitePageCreatedByUserServiceProvisioningErrorClient
	SitePageLastModifiedByUser                                       *sitepagelastmodifiedbyuser.SitePageLastModifiedByUserClient
	SitePageLastModifiedByUserMailboxSetting                         *sitepagelastmodifiedbyusermailboxsetting.SitePageLastModifiedByUserMailboxSettingClient
	SitePageLastModifiedByUserServiceProvisioningError               *sitepagelastmodifiedbyuserserviceprovisioningerror.SitePageLastModifiedByUserServiceProvisioningErrorClient
	SitePermission                                                   *sitepermission.SitePermissionClient
	SiteRecycleBin                                                   *siterecyclebin.SiteRecycleBinClient
	SiteRecycleBinCreatedByUser                                      *siterecyclebincreatedbyuser.SiteRecycleBinCreatedByUserClient
	SiteRecycleBinCreatedByUserMailboxSetting                        *siterecyclebincreatedbyusermailboxsetting.SiteRecycleBinCreatedByUserMailboxSettingClient
	SiteRecycleBinCreatedByUserServiceProvisioningError              *siterecyclebincreatedbyuserserviceprovisioningerror.SiteRecycleBinCreatedByUserServiceProvisioningErrorClient
	SiteRecycleBinItem                                               *siterecyclebinitem.SiteRecycleBinItemClient
	SiteRecycleBinItemCreatedByUser                                  *siterecyclebinitemcreatedbyuser.SiteRecycleBinItemCreatedByUserClient
	SiteRecycleBinItemCreatedByUserMailboxSetting                    *siterecyclebinitemcreatedbyusermailboxsetting.SiteRecycleBinItemCreatedByUserMailboxSettingClient
	SiteRecycleBinItemCreatedByUserServiceProvisioningError          *siterecyclebinitemcreatedbyuserserviceprovisioningerror.SiteRecycleBinItemCreatedByUserServiceProvisioningErrorClient
	SiteRecycleBinItemLastModifiedByUser                             *siterecyclebinitemlastmodifiedbyuser.SiteRecycleBinItemLastModifiedByUserClient
	SiteRecycleBinItemLastModifiedByUserMailboxSetting               *siterecyclebinitemlastmodifiedbyusermailboxsetting.SiteRecycleBinItemLastModifiedByUserMailboxSettingClient
	SiteRecycleBinItemLastModifiedByUserServiceProvisioningError     *siterecyclebinitemlastmodifiedbyuserserviceprovisioningerror.SiteRecycleBinItemLastModifiedByUserServiceProvisioningErrorClient
	SiteRecycleBinLastModifiedByUser                                 *siterecyclebinlastmodifiedbyuser.SiteRecycleBinLastModifiedByUserClient
	SiteRecycleBinLastModifiedByUserMailboxSetting                   *siterecyclebinlastmodifiedbyusermailboxsetting.SiteRecycleBinLastModifiedByUserMailboxSettingClient
	SiteRecycleBinLastModifiedByUserServiceProvisioningError         *siterecyclebinlastmodifiedbyuserserviceprovisioningerror.SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient
	Team                                                             *team.TeamClient
	TeamAllChannel                                                   *teamallchannel.TeamAllChannelClient
	TeamChannel                                                      *teamchannel.TeamChannelClient
	TeamChannelFilesFolder                                           *teamchannelfilesfolder.TeamChannelFilesFolderClient
	TeamChannelFilesFolderContent                                    *teamchannelfilesfoldercontent.TeamChannelFilesFolderContentClient
	TeamChannelFilesFolderContentStream                              *teamchannelfilesfoldercontentstream.TeamChannelFilesFolderContentStreamClient
	TeamChannelMember                                                *teamchannelmember.TeamChannelMemberClient
	TeamChannelMessage                                               *teamchannelmessage.TeamChannelMessageClient
	TeamChannelMessageHostedContent                                  *teamchannelmessagehostedcontent.TeamChannelMessageHostedContentClient
	TeamChannelMessageReply                                          *teamchannelmessagereply.TeamChannelMessageReplyClient
	TeamChannelMessageReplyHostedContent                             *teamchannelmessagereplyhostedcontent.TeamChannelMessageReplyHostedContentClient
	TeamChannelSharedWithTeam                                        *teamchannelsharedwithteam.TeamChannelSharedWithTeamClient
	TeamChannelSharedWithTeamAllowedMember                           *teamchannelsharedwithteamallowedmember.TeamChannelSharedWithTeamAllowedMemberClient
	TeamChannelSharedWithTeamTeam                                    *teamchannelsharedwithteamteam.TeamChannelSharedWithTeamTeamClient
	TeamChannelTab                                                   *teamchanneltab.TeamChannelTabClient
	TeamChannelTabTeamsApp                                           *teamchanneltabteamsapp.TeamChannelTabTeamsAppClient
	TeamGroup                                                        *teamgroup.TeamGroupClient
	TeamGroupServiceProvisioningError                                *teamgroupserviceprovisioningerror.TeamGroupServiceProvisioningErrorClient
	TeamIncomingChannel                                              *teamincomingchannel.TeamIncomingChannelClient
	TeamInstalledApp                                                 *teaminstalledapp.TeamInstalledAppClient
	TeamInstalledAppTeamsApp                                         *teaminstalledappteamsapp.TeamInstalledAppTeamsAppClient
	TeamInstalledAppTeamsAppDefinition                               *teaminstalledappteamsappdefinition.TeamInstalledAppTeamsAppDefinitionClient
	TeamMember                                                       *teammember.TeamMemberClient
	TeamOperation                                                    *teamoperation.TeamOperationClient
	TeamOwner                                                        *teamowner.TeamOwnerClient
	TeamOwnerMailboxSetting                                          *teamownermailboxsetting.TeamOwnerMailboxSettingClient
	TeamOwnerServiceProvisioningError                                *teamownerserviceprovisioningerror.TeamOwnerServiceProvisioningErrorClient
	TeamPermissionGrant                                              *teampermissiongrant.TeamPermissionGrantClient
	TeamPhoto                                                        *teamphoto.TeamPhotoClient
	TeamPrimaryChannel                                               *teamprimarychannel.TeamPrimaryChannelClient
	TeamPrimaryChannelFilesFolder                                    *teamprimarychannelfilesfolder.TeamPrimaryChannelFilesFolderClient
	TeamPrimaryChannelFilesFolderContent                             *teamprimarychannelfilesfoldercontent.TeamPrimaryChannelFilesFolderContentClient
	TeamPrimaryChannelFilesFolderContentStream                       *teamprimarychannelfilesfoldercontentstream.TeamPrimaryChannelFilesFolderContentStreamClient
	TeamPrimaryChannelMember                                         *teamprimarychannelmember.TeamPrimaryChannelMemberClient
	TeamPrimaryChannelMessage                                        *teamprimarychannelmessage.TeamPrimaryChannelMessageClient
	TeamPrimaryChannelMessageHostedContent                           *teamprimarychannelmessagehostedcontent.TeamPrimaryChannelMessageHostedContentClient
	TeamPrimaryChannelMessageReply                                   *teamprimarychannelmessagereply.TeamPrimaryChannelMessageReplyClient
	TeamPrimaryChannelMessageReplyHostedContent                      *teamprimarychannelmessagereplyhostedcontent.TeamPrimaryChannelMessageReplyHostedContentClient
	TeamPrimaryChannelSharedWithTeam                                 *teamprimarychannelsharedwithteam.TeamPrimaryChannelSharedWithTeamClient
	TeamPrimaryChannelSharedWithTeamAllowedMember                    *teamprimarychannelsharedwithteamallowedmember.TeamPrimaryChannelSharedWithTeamAllowedMemberClient
	TeamPrimaryChannelSharedWithTeamTeam                             *teamprimarychannelsharedwithteamteam.TeamPrimaryChannelSharedWithTeamTeamClient
	TeamPrimaryChannelTab                                            *teamprimarychanneltab.TeamPrimaryChannelTabClient
	TeamPrimaryChannelTabTeamsApp                                    *teamprimarychanneltabteamsapp.TeamPrimaryChannelTabTeamsAppClient
	TeamSchedule                                                     *teamschedule.TeamScheduleClient
	TeamScheduleDayNote                                              *teamscheduledaynote.TeamScheduleDayNoteClient
	TeamScheduleOfferShiftRequest                                    *teamscheduleoffershiftrequest.TeamScheduleOfferShiftRequestClient
	TeamScheduleOpenShift                                            *teamscheduleopenshift.TeamScheduleOpenShiftClient
	TeamScheduleOpenShiftChangeRequest                               *teamscheduleopenshiftchangerequest.TeamScheduleOpenShiftChangeRequestClient
	TeamScheduleSchedulingGroup                                      *teamscheduleschedulinggroup.TeamScheduleSchedulingGroupClient
	TeamScheduleShift                                                *teamscheduleshift.TeamScheduleShiftClient
	TeamScheduleShiftsRoleDefinition                                 *teamscheduleshiftsroledefinition.TeamScheduleShiftsRoleDefinitionClient
	TeamScheduleSwapShiftsChangeRequest                              *teamscheduleswapshiftschangerequest.TeamScheduleSwapShiftsChangeRequestClient
	TeamScheduleTimeCard                                             *teamscheduletimecard.TeamScheduleTimeCardClient
	TeamScheduleTimeOffReason                                        *teamscheduletimeoffreason.TeamScheduleTimeOffReasonClient
	TeamScheduleTimeOffRequest                                       *teamscheduletimeoffrequest.TeamScheduleTimeOffRequestClient
	TeamScheduleTimesOff                                             *teamscheduletimesoff.TeamScheduleTimesOffClient
	TeamTag                                                          *teamtag.TeamTagClient
	TeamTagMember                                                    *teamtagmember.TeamTagMemberClient
	TeamTemplate                                                     *teamtemplate.TeamTemplateClient
	TeamTemplateDefinition                                           *teamtemplatedefinition.TeamTemplateDefinitionClient
	Thread                                                           *thread.ThreadClient
	ThreadPost                                                       *threadpost.ThreadPostClient
	ThreadPostAttachment                                             *threadpostattachment.ThreadPostAttachmentClient
	ThreadPostExtension                                              *threadpostextension.ThreadPostExtensionClient
	ThreadPostInReplyTo                                              *threadpostinreplyto.ThreadPostInReplyToClient
	ThreadPostInReplyToAttachment                                    *threadpostinreplytoattachment.ThreadPostInReplyToAttachmentClient
	ThreadPostInReplyToExtension                                     *threadpostinreplytoextension.ThreadPostInReplyToExtensionClient
	ThreadPostInReplyToMention                                       *threadpostinreplytomention.ThreadPostInReplyToMentionClient
	ThreadPostMention                                                *threadpostmention.ThreadPostMentionClient
	TransitiveMember                                                 *transitivemember.TransitiveMemberClient
	TransitiveMemberOf                                               *transitivememberof.TransitiveMemberOfClient
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

	conversationThreadPostInReplyToMentionClient, err := conversationthreadpostinreplytomention.NewConversationThreadPostInReplyToMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPostInReplyToMention client: %+v", err)
	}
	configureFunc(conversationThreadPostInReplyToMentionClient.Client)

	conversationThreadPostMentionClient, err := conversationthreadpostmention.NewConversationThreadPostMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConversationThreadPostMention client: %+v", err)
	}
	configureFunc(conversationThreadPostMentionClient.Client)

	createdOnBehalfOfClient, err := createdonbehalfof.NewCreatedOnBehalfOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreatedOnBehalfOf client: %+v", err)
	}
	configureFunc(createdOnBehalfOfClient.Client)

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

	endpointClient, err := endpoint.NewEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoint client: %+v", err)
	}
	configureFunc(endpointClient.Client)

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

	siteAnalyticsItemActivityStatActivityDriveItemContentStreamClient, err := siteanalyticsitemactivitystatactivitydriveitemcontentstream.NewSiteAnalyticsItemActivityStatActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticsItemActivityStatActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(siteAnalyticsItemActivityStatActivityDriveItemContentStreamClient.Client)

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

	siteContentModelClient, err := sitecontentmodel.NewSiteContentModelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteContentModel client: %+v", err)
	}
	configureFunc(siteContentModelClient.Client)

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

	siteDocumentProcessingJobClient, err := sitedocumentprocessingjob.NewSiteDocumentProcessingJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteDocumentProcessingJob client: %+v", err)
	}
	configureFunc(siteDocumentProcessingJobClient.Client)

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

	siteInformationProtectionBitlockerClient, err := siteinformationprotectionbitlocker.NewSiteInformationProtectionBitlockerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionBitlocker client: %+v", err)
	}
	configureFunc(siteInformationProtectionBitlockerClient.Client)

	siteInformationProtectionBitlockerRecoveryKeyClient, err := siteinformationprotectionbitlockerrecoverykey.NewSiteInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionBitlockerRecoveryKey client: %+v", err)
	}
	configureFunc(siteInformationProtectionBitlockerRecoveryKeyClient.Client)

	siteInformationProtectionClient, err := siteinformationprotection.NewSiteInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtection client: %+v", err)
	}
	configureFunc(siteInformationProtectionClient.Client)

	siteInformationProtectionDataLossPreventionPolicyClient, err := siteinformationprotectiondatalosspreventionpolicy.NewSiteInformationProtectionDataLossPreventionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionDataLossPreventionPolicy client: %+v", err)
	}
	configureFunc(siteInformationProtectionDataLossPreventionPolicyClient.Client)

	siteInformationProtectionPolicyClient, err := siteinformationprotectionpolicy.NewSiteInformationProtectionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionPolicy client: %+v", err)
	}
	configureFunc(siteInformationProtectionPolicyClient.Client)

	siteInformationProtectionPolicyLabelClient, err := siteinformationprotectionpolicylabel.NewSiteInformationProtectionPolicyLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionPolicyLabel client: %+v", err)
	}
	configureFunc(siteInformationProtectionPolicyLabelClient.Client)

	siteInformationProtectionSensitivityLabelClient, err := siteinformationprotectionsensitivitylabel.NewSiteInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(siteInformationProtectionSensitivityLabelClient.Client)

	siteInformationProtectionSensitivityLabelSublabelClient, err := siteinformationprotectionsensitivitylabelsublabel.NewSiteInformationProtectionSensitivityLabelSublabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionSensitivityLabelSublabel client: %+v", err)
	}
	configureFunc(siteInformationProtectionSensitivityLabelSublabelClient.Client)

	siteInformationProtectionSensitivityPolicySettingClient, err := siteinformationprotectionsensitivitypolicysetting.NewSiteInformationProtectionSensitivityPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionSensitivityPolicySetting client: %+v", err)
	}
	configureFunc(siteInformationProtectionSensitivityPolicySettingClient.Client)

	siteInformationProtectionThreatAssessmentRequestClient, err := siteinformationprotectionthreatassessmentrequest.NewSiteInformationProtectionThreatAssessmentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionThreatAssessmentRequest client: %+v", err)
	}
	configureFunc(siteInformationProtectionThreatAssessmentRequestClient.Client)

	siteInformationProtectionThreatAssessmentRequestResultClient, err := siteinformationprotectionthreatassessmentrequestresult.NewSiteInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteInformationProtectionThreatAssessmentRequestResult client: %+v", err)
	}
	configureFunc(siteInformationProtectionThreatAssessmentRequestResultClient.Client)

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

	siteListActivityClient, err := sitelistactivity.NewSiteListActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListActivity client: %+v", err)
	}
	configureFunc(siteListActivityClient.Client)

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

	siteListItemActivityClient, err := sitelistitemactivity.NewSiteListItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemActivity client: %+v", err)
	}
	configureFunc(siteListItemActivityClient.Client)

	siteListItemActivityDriveItemClient, err := sitelistitemactivitydriveitem.NewSiteListItemActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemActivityDriveItem client: %+v", err)
	}
	configureFunc(siteListItemActivityDriveItemClient.Client)

	siteListItemActivityDriveItemContentClient, err := sitelistitemactivitydriveitemcontent.NewSiteListItemActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemActivityDriveItemContent client: %+v", err)
	}
	configureFunc(siteListItemActivityDriveItemContentClient.Client)

	siteListItemActivityDriveItemContentStreamClient, err := sitelistitemactivitydriveitemcontentstream.NewSiteListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemActivityDriveItemContentStream client: %+v", err)
	}
	configureFunc(siteListItemActivityDriveItemContentStreamClient.Client)

	siteListItemActivityListItemClient, err := sitelistitemactivitylistitem.NewSiteListItemActivityListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemActivityListItem client: %+v", err)
	}
	configureFunc(siteListItemActivityListItemClient.Client)

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

	siteListItemDriveItemContentStreamClient, err := sitelistitemdriveitemcontentstream.NewSiteListItemDriveItemContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemDriveItemContentStream client: %+v", err)
	}
	configureFunc(siteListItemDriveItemContentStreamClient.Client)

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

	siteListItemPermissionClient, err := sitelistitempermission.NewSiteListItemPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemPermission client: %+v", err)
	}
	configureFunc(siteListItemPermissionClient.Client)

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

	siteListPermissionClient, err := sitelistpermission.NewSiteListPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListPermission client: %+v", err)
	}
	configureFunc(siteListPermissionClient.Client)

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

	siteRecycleBinClient, err := siterecyclebin.NewSiteRecycleBinClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBin client: %+v", err)
	}
	configureFunc(siteRecycleBinClient.Client)

	siteRecycleBinCreatedByUserClient, err := siterecyclebincreatedbyuser.NewSiteRecycleBinCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinCreatedByUser client: %+v", err)
	}
	configureFunc(siteRecycleBinCreatedByUserClient.Client)

	siteRecycleBinCreatedByUserMailboxSettingClient, err := siterecyclebincreatedbyusermailboxsetting.NewSiteRecycleBinCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteRecycleBinCreatedByUserMailboxSettingClient.Client)

	siteRecycleBinCreatedByUserServiceProvisioningErrorClient, err := siterecyclebincreatedbyuserserviceprovisioningerror.NewSiteRecycleBinCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteRecycleBinCreatedByUserServiceProvisioningErrorClient.Client)

	siteRecycleBinItemClient, err := siterecyclebinitem.NewSiteRecycleBinItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinItem client: %+v", err)
	}
	configureFunc(siteRecycleBinItemClient.Client)

	siteRecycleBinItemCreatedByUserClient, err := siterecyclebinitemcreatedbyuser.NewSiteRecycleBinItemCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinItemCreatedByUser client: %+v", err)
	}
	configureFunc(siteRecycleBinItemCreatedByUserClient.Client)

	siteRecycleBinItemCreatedByUserMailboxSettingClient, err := siterecyclebinitemcreatedbyusermailboxsetting.NewSiteRecycleBinItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinItemCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteRecycleBinItemCreatedByUserMailboxSettingClient.Client)

	siteRecycleBinItemCreatedByUserServiceProvisioningErrorClient, err := siterecyclebinitemcreatedbyuserserviceprovisioningerror.NewSiteRecycleBinItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinItemCreatedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteRecycleBinItemCreatedByUserServiceProvisioningErrorClient.Client)

	siteRecycleBinItemLastModifiedByUserClient, err := siterecyclebinitemlastmodifiedbyuser.NewSiteRecycleBinItemLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinItemLastModifiedByUser client: %+v", err)
	}
	configureFunc(siteRecycleBinItemLastModifiedByUserClient.Client)

	siteRecycleBinItemLastModifiedByUserMailboxSettingClient, err := siterecyclebinitemlastmodifiedbyusermailboxsetting.NewSiteRecycleBinItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinItemLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteRecycleBinItemLastModifiedByUserMailboxSettingClient.Client)

	siteRecycleBinItemLastModifiedByUserServiceProvisioningErrorClient, err := siterecyclebinitemlastmodifiedbyuserserviceprovisioningerror.NewSiteRecycleBinItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinItemLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteRecycleBinItemLastModifiedByUserServiceProvisioningErrorClient.Client)

	siteRecycleBinLastModifiedByUserClient, err := siterecyclebinlastmodifiedbyuser.NewSiteRecycleBinLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinLastModifiedByUser client: %+v", err)
	}
	configureFunc(siteRecycleBinLastModifiedByUserClient.Client)

	siteRecycleBinLastModifiedByUserMailboxSettingClient, err := siterecyclebinlastmodifiedbyusermailboxsetting.NewSiteRecycleBinLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(siteRecycleBinLastModifiedByUserMailboxSettingClient.Client)

	siteRecycleBinLastModifiedByUserServiceProvisioningErrorClient, err := siterecyclebinlastmodifiedbyuserserviceprovisioningerror.NewSiteRecycleBinLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteRecycleBinLastModifiedByUserServiceProvisioningError client: %+v", err)
	}
	configureFunc(siteRecycleBinLastModifiedByUserServiceProvisioningErrorClient.Client)

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

	teamChannelFilesFolderContentStreamClient, err := teamchannelfilesfoldercontentstream.NewTeamChannelFilesFolderContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamChannelFilesFolderContentStream client: %+v", err)
	}
	configureFunc(teamChannelFilesFolderContentStreamClient.Client)

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

	teamOwnerClient, err := teamowner.NewTeamOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamOwner client: %+v", err)
	}
	configureFunc(teamOwnerClient.Client)

	teamOwnerMailboxSettingClient, err := teamownermailboxsetting.NewTeamOwnerMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamOwnerMailboxSetting client: %+v", err)
	}
	configureFunc(teamOwnerMailboxSettingClient.Client)

	teamOwnerServiceProvisioningErrorClient, err := teamownerserviceprovisioningerror.NewTeamOwnerServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamOwnerServiceProvisioningError client: %+v", err)
	}
	configureFunc(teamOwnerServiceProvisioningErrorClient.Client)

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

	teamPrimaryChannelFilesFolderContentStreamClient, err := teamprimarychannelfilesfoldercontentstream.NewTeamPrimaryChannelFilesFolderContentStreamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamPrimaryChannelFilesFolderContentStream client: %+v", err)
	}
	configureFunc(teamPrimaryChannelFilesFolderContentStreamClient.Client)

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

	teamScheduleDayNoteClient, err := teamscheduledaynote.NewTeamScheduleDayNoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleDayNote client: %+v", err)
	}
	configureFunc(teamScheduleDayNoteClient.Client)

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

	teamScheduleShiftsRoleDefinitionClient, err := teamscheduleshiftsroledefinition.NewTeamScheduleShiftsRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleShiftsRoleDefinition client: %+v", err)
	}
	configureFunc(teamScheduleShiftsRoleDefinitionClient.Client)

	teamScheduleSwapShiftsChangeRequestClient, err := teamscheduleswapshiftschangerequest.NewTeamScheduleSwapShiftsChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleSwapShiftsChangeRequest client: %+v", err)
	}
	configureFunc(teamScheduleSwapShiftsChangeRequestClient.Client)

	teamScheduleTimeCardClient, err := teamscheduletimecard.NewTeamScheduleTimeCardClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamScheduleTimeCard client: %+v", err)
	}
	configureFunc(teamScheduleTimeCardClient.Client)

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

	teamTemplateDefinitionClient, err := teamtemplatedefinition.NewTeamTemplateDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TeamTemplateDefinition client: %+v", err)
	}
	configureFunc(teamTemplateDefinitionClient.Client)

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

	threadPostInReplyToMentionClient, err := threadpostinreplytomention.NewThreadPostInReplyToMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPostInReplyToMention client: %+v", err)
	}
	configureFunc(threadPostInReplyToMentionClient.Client)

	threadPostMentionClient, err := threadpostmention.NewThreadPostMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ThreadPostMention client: %+v", err)
	}
	configureFunc(threadPostMentionClient.Client)

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
		AcceptedSender:                                                   acceptedSenderClient,
		AppRoleAssignment:                                                appRoleAssignmentClient,
		Calendar:                                                         calendarClient,
		CalendarCalendarPermission:                                       calendarCalendarPermissionClient,
		CalendarCalendarView:                                             calendarCalendarViewClient,
		CalendarCalendarViewAttachment:                                   calendarCalendarViewAttachmentClient,
		CalendarCalendarViewCalendar:                                     calendarCalendarViewCalendarClient,
		CalendarCalendarViewExceptionOccurrence:                          calendarCalendarViewExceptionOccurrenceClient,
		CalendarCalendarViewExceptionOccurrenceAttachment:                calendarCalendarViewExceptionOccurrenceAttachmentClient,
		CalendarCalendarViewExceptionOccurrenceCalendar:                  calendarCalendarViewExceptionOccurrenceCalendarClient,
		CalendarCalendarViewExceptionOccurrenceExtension:                 calendarCalendarViewExceptionOccurrenceExtensionClient,
		CalendarCalendarViewExceptionOccurrenceInstance:                  calendarCalendarViewExceptionOccurrenceInstanceClient,
		CalendarCalendarViewExceptionOccurrenceInstanceAttachment:        calendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		CalendarCalendarViewExceptionOccurrenceInstanceCalendar:          calendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		CalendarCalendarViewExceptionOccurrenceInstanceExtension:         calendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		CalendarCalendarViewExtension:                                    calendarCalendarViewExtensionClient,
		CalendarCalendarViewInstance:                                     calendarCalendarViewInstanceClient,
		CalendarCalendarViewInstanceAttachment:                           calendarCalendarViewInstanceAttachmentClient,
		CalendarCalendarViewInstanceCalendar:                             calendarCalendarViewInstanceCalendarClient,
		CalendarCalendarViewInstanceExceptionOccurrence:                  calendarCalendarViewInstanceExceptionOccurrenceClient,
		CalendarCalendarViewInstanceExceptionOccurrenceAttachment:        calendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		CalendarCalendarViewInstanceExceptionOccurrenceCalendar:          calendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		CalendarCalendarViewInstanceExceptionOccurrenceExtension:         calendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		CalendarCalendarViewInstanceExtension:                            calendarCalendarViewInstanceExtensionClient,
		CalendarEvent:                                                    calendarEventClient,
		CalendarEventAttachment:                                          calendarEventAttachmentClient,
		CalendarEventCalendar:                                            calendarEventCalendarClient,
		CalendarEventExceptionOccurrence:                                 calendarEventExceptionOccurrenceClient,
		CalendarEventExceptionOccurrenceAttachment:                       calendarEventExceptionOccurrenceAttachmentClient,
		CalendarEventExceptionOccurrenceCalendar:                         calendarEventExceptionOccurrenceCalendarClient,
		CalendarEventExceptionOccurrenceExtension:                        calendarEventExceptionOccurrenceExtensionClient,
		CalendarEventExceptionOccurrenceInstance:                         calendarEventExceptionOccurrenceInstanceClient,
		CalendarEventExceptionOccurrenceInstanceAttachment:               calendarEventExceptionOccurrenceInstanceAttachmentClient,
		CalendarEventExceptionOccurrenceInstanceCalendar:                 calendarEventExceptionOccurrenceInstanceCalendarClient,
		CalendarEventExceptionOccurrenceInstanceExtension:                calendarEventExceptionOccurrenceInstanceExtensionClient,
		CalendarEventExtension:                                           calendarEventExtensionClient,
		CalendarEventInstance:                                            calendarEventInstanceClient,
		CalendarEventInstanceAttachment:                                  calendarEventInstanceAttachmentClient,
		CalendarEventInstanceCalendar:                                    calendarEventInstanceCalendarClient,
		CalendarEventInstanceExceptionOccurrence:                         calendarEventInstanceExceptionOccurrenceClient,
		CalendarEventInstanceExceptionOccurrenceAttachment:               calendarEventInstanceExceptionOccurrenceAttachmentClient,
		CalendarEventInstanceExceptionOccurrenceCalendar:                 calendarEventInstanceExceptionOccurrenceCalendarClient,
		CalendarEventInstanceExceptionOccurrenceExtension:                calendarEventInstanceExceptionOccurrenceExtensionClient,
		CalendarEventInstanceExtension:                                   calendarEventInstanceExtensionClient,
		CalendarView:                                                     calendarViewClient,
		CalendarViewAttachment:                                           calendarViewAttachmentClient,
		CalendarViewCalendar:                                             calendarViewCalendarClient,
		CalendarViewExceptionOccurrence:                                  calendarViewExceptionOccurrenceClient,
		CalendarViewExceptionOccurrenceAttachment:                        calendarViewExceptionOccurrenceAttachmentClient,
		CalendarViewExceptionOccurrenceCalendar:                          calendarViewExceptionOccurrenceCalendarClient,
		CalendarViewExceptionOccurrenceExtension:                         calendarViewExceptionOccurrenceExtensionClient,
		CalendarViewExceptionOccurrenceInstance:                          calendarViewExceptionOccurrenceInstanceClient,
		CalendarViewExceptionOccurrenceInstanceAttachment:                calendarViewExceptionOccurrenceInstanceAttachmentClient,
		CalendarViewExceptionOccurrenceInstanceCalendar:                  calendarViewExceptionOccurrenceInstanceCalendarClient,
		CalendarViewExceptionOccurrenceInstanceExtension:                 calendarViewExceptionOccurrenceInstanceExtensionClient,
		CalendarViewExtension:                                            calendarViewExtensionClient,
		CalendarViewInstance:                                             calendarViewInstanceClient,
		CalendarViewInstanceAttachment:                                   calendarViewInstanceAttachmentClient,
		CalendarViewInstanceCalendar:                                     calendarViewInstanceCalendarClient,
		CalendarViewInstanceExceptionOccurrence:                          calendarViewInstanceExceptionOccurrenceClient,
		CalendarViewInstanceExceptionOccurrenceAttachment:                calendarViewInstanceExceptionOccurrenceAttachmentClient,
		CalendarViewInstanceExceptionOccurrenceCalendar:                  calendarViewInstanceExceptionOccurrenceCalendarClient,
		CalendarViewInstanceExceptionOccurrenceExtension:                 calendarViewInstanceExceptionOccurrenceExtensionClient,
		CalendarViewInstanceExtension:                                    calendarViewInstanceExtensionClient,
		Conversation:                                                     conversationClient,
		ConversationThread:                                               conversationThreadClient,
		ConversationThreadPost:                                           conversationThreadPostClient,
		ConversationThreadPostAttachment:                                 conversationThreadPostAttachmentClient,
		ConversationThreadPostExtension:                                  conversationThreadPostExtensionClient,
		ConversationThreadPostInReplyTo:                                  conversationThreadPostInReplyToClient,
		ConversationThreadPostInReplyToAttachment:                        conversationThreadPostInReplyToAttachmentClient,
		ConversationThreadPostInReplyToExtension:                         conversationThreadPostInReplyToExtensionClient,
		ConversationThreadPostInReplyToMention:                           conversationThreadPostInReplyToMentionClient,
		ConversationThreadPostMention:                                    conversationThreadPostMentionClient,
		CreatedOnBehalfOf:                                                createdOnBehalfOfClient,
		Drive:                                                            driveClient,
		DriveActivity:                                                    driveActivityClient,
		DriveActivityDriveItem:                                           driveActivityDriveItemClient,
		DriveActivityDriveItemContent:                                    driveActivityDriveItemContentClient,
		DriveActivityDriveItemContentStream:                              driveActivityDriveItemContentStreamClient,
		DriveActivityListItem:                                            driveActivityListItemClient,
		DriveBundle:                                                      driveBundleClient,
		DriveBundleContent:                                               driveBundleContentClient,
		DriveBundleContentStream:                                         driveBundleContentStreamClient,
		DriveCreatedByUser:                                               driveCreatedByUserClient,
		DriveCreatedByUserMailboxSetting:                                 driveCreatedByUserMailboxSettingClient,
		DriveCreatedByUserServiceProvisioningError:                       driveCreatedByUserServiceProvisioningErrorClient,
		DriveFollowing:                                                   driveFollowingClient,
		DriveFollowingContent:                                            driveFollowingContentClient,
		DriveFollowingContentStream:                                      driveFollowingContentStreamClient,
		DriveItem:                                                        driveItemClient,
		DriveItemActivity:                                                driveItemActivityClient,
		DriveItemAnalytics:                                               driveItemAnalyticsClient,
		DriveItemAnalyticsAllTime:                                        driveItemAnalyticsAllTimeClient,
		DriveItemAnalyticsItemActivityStat:                               driveItemAnalyticsItemActivityStatClient,
		DriveItemAnalyticsItemActivityStatActivity:                       driveItemAnalyticsItemActivityStatActivityClient,
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
		Endpoint:                                                         endpointClient,
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
		Group:                                                            groupClient,
		GroupLifecyclePolicy:                                             groupLifecyclePolicyClient,
		Member:                                                           memberClient,
		MemberOf:                                                         memberOfClient,
		MembersWithLicenseError:                                          membersWithLicenseErrorClient,
		Owner:                                                            ownerClient,
		PermissionGrant:                                                  permissionGrantClient,
		Photo:                                                            photoClient,
		Planner:                                                          plannerClient,
		PlannerPlan:                                                      plannerPlanClient,
		PlannerPlanBucket:                                                plannerPlanBucketClient,
		PlannerPlanBucketTask:                                            plannerPlanBucketTaskClient,
		PlannerPlanBucketTaskAssignedToTaskBoardFormat:                   plannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		PlannerPlanBucketTaskBucketTaskBoardFormat:                       plannerPlanBucketTaskBucketTaskBoardFormatClient,
		PlannerPlanBucketTaskDetail:                                      plannerPlanBucketTaskDetailClient,
		PlannerPlanBucketTaskProgressTaskBoardFormat:                     plannerPlanBucketTaskProgressTaskBoardFormatClient,
		PlannerPlanDetail:                                                plannerPlanDetailClient,
		PlannerPlanTask:                                                  plannerPlanTaskClient,
		PlannerPlanTaskAssignedToTaskBoardFormat:                         plannerPlanTaskAssignedToTaskBoardFormatClient,
		PlannerPlanTaskBucketTaskBoardFormat:                             plannerPlanTaskBucketTaskBoardFormatClient,
		PlannerPlanTaskDetail:                                            plannerPlanTaskDetailClient,
		PlannerPlanTaskProgressTaskBoardFormat:                           plannerPlanTaskProgressTaskBoardFormatClient,
		RejectedSender:                                                   rejectedSenderClient,
		ServiceProvisioningError:                                         serviceProvisioningErrorClient,
		Setting:                                                          settingClient,
		Site:                                                             siteClient,
		SiteAnalytics:                                                    siteAnalyticsClient,
		SiteAnalyticsAllTime:                                             siteAnalyticsAllTimeClient,
		SiteAnalyticsItemActivityStat:                                    siteAnalyticsItemActivityStatClient,
		SiteAnalyticsItemActivityStatActivity:                            siteAnalyticsItemActivityStatActivityClient,
		SiteAnalyticsItemActivityStatActivityDriveItem:                   siteAnalyticsItemActivityStatActivityDriveItemClient,
		SiteAnalyticsItemActivityStatActivityDriveItemContent:            siteAnalyticsItemActivityStatActivityDriveItemContentClient,
		SiteAnalyticsItemActivityStatActivityDriveItemContentStream:      siteAnalyticsItemActivityStatActivityDriveItemContentStreamClient,
		SiteAnalyticsLastSevenDay:                                        siteAnalyticsLastSevenDayClient,
		SiteColumn:                                                       siteColumnClient,
		SiteColumnSourceColumn:                                           siteColumnSourceColumnClient,
		SiteContentModel:                                                 siteContentModelClient,
		SiteContentType:                                                  siteContentTypeClient,
		SiteContentTypeBase:                                              siteContentTypeBaseClient,
		SiteContentTypeBaseType:                                          siteContentTypeBaseTypeClient,
		SiteContentTypeColumn:                                            siteContentTypeColumnClient,
		SiteContentTypeColumnLink:                                        siteContentTypeColumnLinkClient,
		SiteContentTypeColumnPosition:                                    siteContentTypeColumnPositionClient,
		SiteContentTypeColumnSourceColumn:                                siteContentTypeColumnSourceColumnClient,
		SiteCreatedByUser:                                                siteCreatedByUserClient,
		SiteCreatedByUserMailboxSetting:                                  siteCreatedByUserMailboxSettingClient,
		SiteCreatedByUserServiceProvisioningError:                        siteCreatedByUserServiceProvisioningErrorClient,
		SiteDocumentProcessingJob:                                        siteDocumentProcessingJobClient,
		SiteDrive:                                                        siteDriveClient,
		SiteExternalColumn:                                               siteExternalColumnClient,
		SiteInformationProtection:                                        siteInformationProtectionClient,
		SiteInformationProtectionBitlocker:                               siteInformationProtectionBitlockerClient,
		SiteInformationProtectionBitlockerRecoveryKey:                    siteInformationProtectionBitlockerRecoveryKeyClient,
		SiteInformationProtectionDataLossPreventionPolicy:                siteInformationProtectionDataLossPreventionPolicyClient,
		SiteInformationProtectionPolicy:                                  siteInformationProtectionPolicyClient,
		SiteInformationProtectionPolicyLabel:                             siteInformationProtectionPolicyLabelClient,
		SiteInformationProtectionSensitivityLabel:                        siteInformationProtectionSensitivityLabelClient,
		SiteInformationProtectionSensitivityLabelSublabel:                siteInformationProtectionSensitivityLabelSublabelClient,
		SiteInformationProtectionSensitivityPolicySetting:                siteInformationProtectionSensitivityPolicySettingClient,
		SiteInformationProtectionThreatAssessmentRequest:                 siteInformationProtectionThreatAssessmentRequestClient,
		SiteInformationProtectionThreatAssessmentRequestResult:           siteInformationProtectionThreatAssessmentRequestResultClient,
		SiteItem:                                       siteItemClient,
		SiteLastModifiedByUser:                         siteLastModifiedByUserClient,
		SiteLastModifiedByUserMailboxSetting:           siteLastModifiedByUserMailboxSettingClient,
		SiteLastModifiedByUserServiceProvisioningError: siteLastModifiedByUserServiceProvisioningErrorClient,
		SiteList:                                                     siteListClient,
		SiteListActivity:                                             siteListActivityClient,
		SiteListColumn:                                               siteListColumnClient,
		SiteListColumnSourceColumn:                                   siteListColumnSourceColumnClient,
		SiteListContentType:                                          siteListContentTypeClient,
		SiteListContentTypeBase:                                      siteListContentTypeBaseClient,
		SiteListContentTypeBaseType:                                  siteListContentTypeBaseTypeClient,
		SiteListContentTypeColumn:                                    siteListContentTypeColumnClient,
		SiteListContentTypeColumnLink:                                siteListContentTypeColumnLinkClient,
		SiteListContentTypeColumnPosition:                            siteListContentTypeColumnPositionClient,
		SiteListContentTypeColumnSourceColumn:                        siteListContentTypeColumnSourceColumnClient,
		SiteListCreatedByUser:                                        siteListCreatedByUserClient,
		SiteListCreatedByUserMailboxSetting:                          siteListCreatedByUserMailboxSettingClient,
		SiteListCreatedByUserServiceProvisioningError:                siteListCreatedByUserServiceProvisioningErrorClient,
		SiteListDrive:                                                siteListDriveClient,
		SiteListItem:                                                 siteListItemClient,
		SiteListItemActivity:                                         siteListItemActivityClient,
		SiteListItemActivityDriveItem:                                siteListItemActivityDriveItemClient,
		SiteListItemActivityDriveItemContent:                         siteListItemActivityDriveItemContentClient,
		SiteListItemActivityDriveItemContentStream:                   siteListItemActivityDriveItemContentStreamClient,
		SiteListItemActivityListItem:                                 siteListItemActivityListItemClient,
		SiteListItemAnalytics:                                        siteListItemAnalyticsClient,
		SiteListItemCreatedByUser:                                    siteListItemCreatedByUserClient,
		SiteListItemCreatedByUserMailboxSetting:                      siteListItemCreatedByUserMailboxSettingClient,
		SiteListItemCreatedByUserServiceProvisioningError:            siteListItemCreatedByUserServiceProvisioningErrorClient,
		SiteListItemDocumentSetVersion:                               siteListItemDocumentSetVersionClient,
		SiteListItemDocumentSetVersionField:                          siteListItemDocumentSetVersionFieldClient,
		SiteListItemDriveItem:                                        siteListItemDriveItemClient,
		SiteListItemDriveItemContent:                                 siteListItemDriveItemContentClient,
		SiteListItemDriveItemContentStream:                           siteListItemDriveItemContentStreamClient,
		SiteListItemField:                                            siteListItemFieldClient,
		SiteListItemLastModifiedByUser:                               siteListItemLastModifiedByUserClient,
		SiteListItemLastModifiedByUserMailboxSetting:                 siteListItemLastModifiedByUserMailboxSettingClient,
		SiteListItemLastModifiedByUserServiceProvisioningError:       siteListItemLastModifiedByUserServiceProvisioningErrorClient,
		SiteListItemPermission:                                       siteListItemPermissionClient,
		SiteListItemVersion:                                          siteListItemVersionClient,
		SiteListItemVersionField:                                     siteListItemVersionFieldClient,
		SiteListLastModifiedByUser:                                   siteListLastModifiedByUserClient,
		SiteListLastModifiedByUserMailboxSetting:                     siteListLastModifiedByUserMailboxSettingClient,
		SiteListLastModifiedByUserServiceProvisioningError:           siteListLastModifiedByUserServiceProvisioningErrorClient,
		SiteListOperation:                                            siteListOperationClient,
		SiteListPermission:                                           siteListPermissionClient,
		SiteListSubscription:                                         siteListSubscriptionClient,
		SiteOperation:                                                siteOperationClient,
		SitePage:                                                     sitePageClient,
		SitePageCreatedByUser:                                        sitePageCreatedByUserClient,
		SitePageCreatedByUserMailboxSetting:                          sitePageCreatedByUserMailboxSettingClient,
		SitePageCreatedByUserServiceProvisioningError:                sitePageCreatedByUserServiceProvisioningErrorClient,
		SitePageLastModifiedByUser:                                   sitePageLastModifiedByUserClient,
		SitePageLastModifiedByUserMailboxSetting:                     sitePageLastModifiedByUserMailboxSettingClient,
		SitePageLastModifiedByUserServiceProvisioningError:           sitePageLastModifiedByUserServiceProvisioningErrorClient,
		SitePermission:                                               sitePermissionClient,
		SiteRecycleBin:                                               siteRecycleBinClient,
		SiteRecycleBinCreatedByUser:                                  siteRecycleBinCreatedByUserClient,
		SiteRecycleBinCreatedByUserMailboxSetting:                    siteRecycleBinCreatedByUserMailboxSettingClient,
		SiteRecycleBinCreatedByUserServiceProvisioningError:          siteRecycleBinCreatedByUserServiceProvisioningErrorClient,
		SiteRecycleBinItem:                                           siteRecycleBinItemClient,
		SiteRecycleBinItemCreatedByUser:                              siteRecycleBinItemCreatedByUserClient,
		SiteRecycleBinItemCreatedByUserMailboxSetting:                siteRecycleBinItemCreatedByUserMailboxSettingClient,
		SiteRecycleBinItemCreatedByUserServiceProvisioningError:      siteRecycleBinItemCreatedByUserServiceProvisioningErrorClient,
		SiteRecycleBinItemLastModifiedByUser:                         siteRecycleBinItemLastModifiedByUserClient,
		SiteRecycleBinItemLastModifiedByUserMailboxSetting:           siteRecycleBinItemLastModifiedByUserMailboxSettingClient,
		SiteRecycleBinItemLastModifiedByUserServiceProvisioningError: siteRecycleBinItemLastModifiedByUserServiceProvisioningErrorClient,
		SiteRecycleBinLastModifiedByUser:                             siteRecycleBinLastModifiedByUserClient,
		SiteRecycleBinLastModifiedByUserMailboxSetting:               siteRecycleBinLastModifiedByUserMailboxSettingClient,
		SiteRecycleBinLastModifiedByUserServiceProvisioningError:     siteRecycleBinLastModifiedByUserServiceProvisioningErrorClient,
		Team:                                          teamClient,
		TeamAllChannel:                                teamAllChannelClient,
		TeamChannel:                                   teamChannelClient,
		TeamChannelFilesFolder:                        teamChannelFilesFolderClient,
		TeamChannelFilesFolderContent:                 teamChannelFilesFolderContentClient,
		TeamChannelFilesFolderContentStream:           teamChannelFilesFolderContentStreamClient,
		TeamChannelMember:                             teamChannelMemberClient,
		TeamChannelMessage:                            teamChannelMessageClient,
		TeamChannelMessageHostedContent:               teamChannelMessageHostedContentClient,
		TeamChannelMessageReply:                       teamChannelMessageReplyClient,
		TeamChannelMessageReplyHostedContent:          teamChannelMessageReplyHostedContentClient,
		TeamChannelSharedWithTeam:                     teamChannelSharedWithTeamClient,
		TeamChannelSharedWithTeamAllowedMember:        teamChannelSharedWithTeamAllowedMemberClient,
		TeamChannelSharedWithTeamTeam:                 teamChannelSharedWithTeamTeamClient,
		TeamChannelTab:                                teamChannelTabClient,
		TeamChannelTabTeamsApp:                        teamChannelTabTeamsAppClient,
		TeamGroup:                                     teamGroupClient,
		TeamGroupServiceProvisioningError:             teamGroupServiceProvisioningErrorClient,
		TeamIncomingChannel:                           teamIncomingChannelClient,
		TeamInstalledApp:                              teamInstalledAppClient,
		TeamInstalledAppTeamsApp:                      teamInstalledAppTeamsAppClient,
		TeamInstalledAppTeamsAppDefinition:            teamInstalledAppTeamsAppDefinitionClient,
		TeamMember:                                    teamMemberClient,
		TeamOperation:                                 teamOperationClient,
		TeamOwner:                                     teamOwnerClient,
		TeamOwnerMailboxSetting:                       teamOwnerMailboxSettingClient,
		TeamOwnerServiceProvisioningError:             teamOwnerServiceProvisioningErrorClient,
		TeamPermissionGrant:                           teamPermissionGrantClient,
		TeamPhoto:                                     teamPhotoClient,
		TeamPrimaryChannel:                            teamPrimaryChannelClient,
		TeamPrimaryChannelFilesFolder:                 teamPrimaryChannelFilesFolderClient,
		TeamPrimaryChannelFilesFolderContent:          teamPrimaryChannelFilesFolderContentClient,
		TeamPrimaryChannelFilesFolderContentStream:    teamPrimaryChannelFilesFolderContentStreamClient,
		TeamPrimaryChannelMember:                      teamPrimaryChannelMemberClient,
		TeamPrimaryChannelMessage:                     teamPrimaryChannelMessageClient,
		TeamPrimaryChannelMessageHostedContent:        teamPrimaryChannelMessageHostedContentClient,
		TeamPrimaryChannelMessageReply:                teamPrimaryChannelMessageReplyClient,
		TeamPrimaryChannelMessageReplyHostedContent:   teamPrimaryChannelMessageReplyHostedContentClient,
		TeamPrimaryChannelSharedWithTeam:              teamPrimaryChannelSharedWithTeamClient,
		TeamPrimaryChannelSharedWithTeamAllowedMember: teamPrimaryChannelSharedWithTeamAllowedMemberClient,
		TeamPrimaryChannelSharedWithTeamTeam:          teamPrimaryChannelSharedWithTeamTeamClient,
		TeamPrimaryChannelTab:                         teamPrimaryChannelTabClient,
		TeamPrimaryChannelTabTeamsApp:                 teamPrimaryChannelTabTeamsAppClient,
		TeamSchedule:                                  teamScheduleClient,
		TeamScheduleDayNote:                           teamScheduleDayNoteClient,
		TeamScheduleOfferShiftRequest:                 teamScheduleOfferShiftRequestClient,
		TeamScheduleOpenShift:                         teamScheduleOpenShiftClient,
		TeamScheduleOpenShiftChangeRequest:            teamScheduleOpenShiftChangeRequestClient,
		TeamScheduleSchedulingGroup:                   teamScheduleSchedulingGroupClient,
		TeamScheduleShift:                             teamScheduleShiftClient,
		TeamScheduleShiftsRoleDefinition:              teamScheduleShiftsRoleDefinitionClient,
		TeamScheduleSwapShiftsChangeRequest:           teamScheduleSwapShiftsChangeRequestClient,
		TeamScheduleTimeCard:                          teamScheduleTimeCardClient,
		TeamScheduleTimeOffReason:                     teamScheduleTimeOffReasonClient,
		TeamScheduleTimeOffRequest:                    teamScheduleTimeOffRequestClient,
		TeamScheduleTimesOff:                          teamScheduleTimesOffClient,
		TeamTag:                                       teamTagClient,
		TeamTagMember:                                 teamTagMemberClient,
		TeamTemplate:                                  teamTemplateClient,
		TeamTemplateDefinition:                        teamTemplateDefinitionClient,
		Thread:                                        threadClient,
		ThreadPost:                                    threadPostClient,
		ThreadPostAttachment:                          threadPostAttachmentClient,
		ThreadPostExtension:                           threadPostExtensionClient,
		ThreadPostInReplyTo:                           threadPostInReplyToClient,
		ThreadPostInReplyToAttachment:                 threadPostInReplyToAttachmentClient,
		ThreadPostInReplyToExtension:                  threadPostInReplyToExtensionClient,
		ThreadPostInReplyToMention:                    threadPostInReplyToMentionClient,
		ThreadPostMention:                             threadPostMentionClient,
		TransitiveMember:                              transitiveMemberClient,
		TransitiveMemberOf:                            transitiveMemberOfClient,
	}, nil
}
