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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteanalyticlastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitecolumnsourcecolumn"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemanalytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemcreatedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistlastmodifiedbyuserserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/sitelistoperation"
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
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AcceptedSender                                               *acceptedsender.AcceptedSenderClient
	AppRoleAssignment                                            *approleassignment.AppRoleAssignmentClient
	Calendar                                                     *calendar.CalendarClient
	CalendarCalendarPermission                                   *calendarcalendarpermission.CalendarCalendarPermissionClient
	CalendarCalendarView                                         *calendarcalendarview.CalendarCalendarViewClient
	CalendarCalendarViewAttachment                               *calendarcalendarviewattachment.CalendarCalendarViewAttachmentClient
	CalendarCalendarViewCalendar                                 *calendarcalendarviewcalendar.CalendarCalendarViewCalendarClient
	CalendarCalendarViewExceptionOccurrence                      *calendarcalendarviewexceptionoccurrence.CalendarCalendarViewExceptionOccurrenceClient
	CalendarCalendarViewExceptionOccurrenceAttachment            *calendarcalendarviewexceptionoccurrenceattachment.CalendarCalendarViewExceptionOccurrenceAttachmentClient
	CalendarCalendarViewExceptionOccurrenceCalendar              *calendarcalendarviewexceptionoccurrencecalendar.CalendarCalendarViewExceptionOccurrenceCalendarClient
	CalendarCalendarViewExceptionOccurrenceExtension             *calendarcalendarviewexceptionoccurrenceextension.CalendarCalendarViewExceptionOccurrenceExtensionClient
	CalendarCalendarViewExceptionOccurrenceInstance              *calendarcalendarviewexceptionoccurrenceinstance.CalendarCalendarViewExceptionOccurrenceInstanceClient
	CalendarCalendarViewExceptionOccurrenceInstanceAttachment    *calendarcalendarviewexceptionoccurrenceinstanceattachment.CalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	CalendarCalendarViewExceptionOccurrenceInstanceCalendar      *calendarcalendarviewexceptionoccurrenceinstancecalendar.CalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	CalendarCalendarViewExceptionOccurrenceInstanceExtension     *calendarcalendarviewexceptionoccurrenceinstanceextension.CalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	CalendarCalendarViewExtension                                *calendarcalendarviewextension.CalendarCalendarViewExtensionClient
	CalendarCalendarViewInstance                                 *calendarcalendarviewinstance.CalendarCalendarViewInstanceClient
	CalendarCalendarViewInstanceAttachment                       *calendarcalendarviewinstanceattachment.CalendarCalendarViewInstanceAttachmentClient
	CalendarCalendarViewInstanceCalendar                         *calendarcalendarviewinstancecalendar.CalendarCalendarViewInstanceCalendarClient
	CalendarCalendarViewInstanceExceptionOccurrence              *calendarcalendarviewinstanceexceptionoccurrence.CalendarCalendarViewInstanceExceptionOccurrenceClient
	CalendarCalendarViewInstanceExceptionOccurrenceAttachment    *calendarcalendarviewinstanceexceptionoccurrenceattachment.CalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	CalendarCalendarViewInstanceExceptionOccurrenceCalendar      *calendarcalendarviewinstanceexceptionoccurrencecalendar.CalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	CalendarCalendarViewInstanceExceptionOccurrenceExtension     *calendarcalendarviewinstanceexceptionoccurrenceextension.CalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	CalendarCalendarViewInstanceExtension                        *calendarcalendarviewinstanceextension.CalendarCalendarViewInstanceExtensionClient
	CalendarEvent                                                *calendarevent.CalendarEventClient
	CalendarEventAttachment                                      *calendareventattachment.CalendarEventAttachmentClient
	CalendarEventCalendar                                        *calendareventcalendar.CalendarEventCalendarClient
	CalendarEventExceptionOccurrence                             *calendareventexceptionoccurrence.CalendarEventExceptionOccurrenceClient
	CalendarEventExceptionOccurrenceAttachment                   *calendareventexceptionoccurrenceattachment.CalendarEventExceptionOccurrenceAttachmentClient
	CalendarEventExceptionOccurrenceCalendar                     *calendareventexceptionoccurrencecalendar.CalendarEventExceptionOccurrenceCalendarClient
	CalendarEventExceptionOccurrenceExtension                    *calendareventexceptionoccurrenceextension.CalendarEventExceptionOccurrenceExtensionClient
	CalendarEventExceptionOccurrenceInstance                     *calendareventexceptionoccurrenceinstance.CalendarEventExceptionOccurrenceInstanceClient
	CalendarEventExceptionOccurrenceInstanceAttachment           *calendareventexceptionoccurrenceinstanceattachment.CalendarEventExceptionOccurrenceInstanceAttachmentClient
	CalendarEventExceptionOccurrenceInstanceCalendar             *calendareventexceptionoccurrenceinstancecalendar.CalendarEventExceptionOccurrenceInstanceCalendarClient
	CalendarEventExceptionOccurrenceInstanceExtension            *calendareventexceptionoccurrenceinstanceextension.CalendarEventExceptionOccurrenceInstanceExtensionClient
	CalendarEventExtension                                       *calendareventextension.CalendarEventExtensionClient
	CalendarEventInstance                                        *calendareventinstance.CalendarEventInstanceClient
	CalendarEventInstanceAttachment                              *calendareventinstanceattachment.CalendarEventInstanceAttachmentClient
	CalendarEventInstanceCalendar                                *calendareventinstancecalendar.CalendarEventInstanceCalendarClient
	CalendarEventInstanceExceptionOccurrence                     *calendareventinstanceexceptionoccurrence.CalendarEventInstanceExceptionOccurrenceClient
	CalendarEventInstanceExceptionOccurrenceAttachment           *calendareventinstanceexceptionoccurrenceattachment.CalendarEventInstanceExceptionOccurrenceAttachmentClient
	CalendarEventInstanceExceptionOccurrenceCalendar             *calendareventinstanceexceptionoccurrencecalendar.CalendarEventInstanceExceptionOccurrenceCalendarClient
	CalendarEventInstanceExceptionOccurrenceExtension            *calendareventinstanceexceptionoccurrenceextension.CalendarEventInstanceExceptionOccurrenceExtensionClient
	CalendarEventInstanceExtension                               *calendareventinstanceextension.CalendarEventInstanceExtensionClient
	CalendarView                                                 *calendarview.CalendarViewClient
	CalendarViewAttachment                                       *calendarviewattachment.CalendarViewAttachmentClient
	CalendarViewCalendar                                         *calendarviewcalendar.CalendarViewCalendarClient
	CalendarViewExceptionOccurrence                              *calendarviewexceptionoccurrence.CalendarViewExceptionOccurrenceClient
	CalendarViewExceptionOccurrenceAttachment                    *calendarviewexceptionoccurrenceattachment.CalendarViewExceptionOccurrenceAttachmentClient
	CalendarViewExceptionOccurrenceCalendar                      *calendarviewexceptionoccurrencecalendar.CalendarViewExceptionOccurrenceCalendarClient
	CalendarViewExceptionOccurrenceExtension                     *calendarviewexceptionoccurrenceextension.CalendarViewExceptionOccurrenceExtensionClient
	CalendarViewExceptionOccurrenceInstance                      *calendarviewexceptionoccurrenceinstance.CalendarViewExceptionOccurrenceInstanceClient
	CalendarViewExceptionOccurrenceInstanceAttachment            *calendarviewexceptionoccurrenceinstanceattachment.CalendarViewExceptionOccurrenceInstanceAttachmentClient
	CalendarViewExceptionOccurrenceInstanceCalendar              *calendarviewexceptionoccurrenceinstancecalendar.CalendarViewExceptionOccurrenceInstanceCalendarClient
	CalendarViewExceptionOccurrenceInstanceExtension             *calendarviewexceptionoccurrenceinstanceextension.CalendarViewExceptionOccurrenceInstanceExtensionClient
	CalendarViewExtension                                        *calendarviewextension.CalendarViewExtensionClient
	CalendarViewInstance                                         *calendarviewinstance.CalendarViewInstanceClient
	CalendarViewInstanceAttachment                               *calendarviewinstanceattachment.CalendarViewInstanceAttachmentClient
	CalendarViewInstanceCalendar                                 *calendarviewinstancecalendar.CalendarViewInstanceCalendarClient
	CalendarViewInstanceExceptionOccurrence                      *calendarviewinstanceexceptionoccurrence.CalendarViewInstanceExceptionOccurrenceClient
	CalendarViewInstanceExceptionOccurrenceAttachment            *calendarviewinstanceexceptionoccurrenceattachment.CalendarViewInstanceExceptionOccurrenceAttachmentClient
	CalendarViewInstanceExceptionOccurrenceCalendar              *calendarviewinstanceexceptionoccurrencecalendar.CalendarViewInstanceExceptionOccurrenceCalendarClient
	CalendarViewInstanceExceptionOccurrenceExtension             *calendarviewinstanceexceptionoccurrenceextension.CalendarViewInstanceExceptionOccurrenceExtensionClient
	CalendarViewInstanceExtension                                *calendarviewinstanceextension.CalendarViewInstanceExtensionClient
	Conversation                                                 *conversation.ConversationClient
	ConversationThread                                           *conversationthread.ConversationThreadClient
	ConversationThreadPost                                       *conversationthreadpost.ConversationThreadPostClient
	ConversationThreadPostAttachment                             *conversationthreadpostattachment.ConversationThreadPostAttachmentClient
	ConversationThreadPostExtension                              *conversationthreadpostextension.ConversationThreadPostExtensionClient
	ConversationThreadPostInReplyTo                              *conversationthreadpostinreplyto.ConversationThreadPostInReplyToClient
	ConversationThreadPostInReplyToAttachment                    *conversationthreadpostinreplytoattachment.ConversationThreadPostInReplyToAttachmentClient
	ConversationThreadPostInReplyToExtension                     *conversationthreadpostinreplytoextension.ConversationThreadPostInReplyToExtensionClient
	ConversationThreadPostInReplyToMention                       *conversationthreadpostinreplytomention.ConversationThreadPostInReplyToMentionClient
	ConversationThreadPostMention                                *conversationthreadpostmention.ConversationThreadPostMentionClient
	CreatedOnBehalfOf                                            *createdonbehalfof.CreatedOnBehalfOfClient
	Drive                                                        *drive.DriveClient
	Endpoint                                                     *endpoint.EndpointClient
	Event                                                        *event.EventClient
	EventAttachment                                              *eventattachment.EventAttachmentClient
	EventCalendar                                                *eventcalendar.EventCalendarClient
	EventExceptionOccurrence                                     *eventexceptionoccurrence.EventExceptionOccurrenceClient
	EventExceptionOccurrenceAttachment                           *eventexceptionoccurrenceattachment.EventExceptionOccurrenceAttachmentClient
	EventExceptionOccurrenceCalendar                             *eventexceptionoccurrencecalendar.EventExceptionOccurrenceCalendarClient
	EventExceptionOccurrenceExtension                            *eventexceptionoccurrenceextension.EventExceptionOccurrenceExtensionClient
	EventExceptionOccurrenceInstance                             *eventexceptionoccurrenceinstance.EventExceptionOccurrenceInstanceClient
	EventExceptionOccurrenceInstanceAttachment                   *eventexceptionoccurrenceinstanceattachment.EventExceptionOccurrenceInstanceAttachmentClient
	EventExceptionOccurrenceInstanceCalendar                     *eventexceptionoccurrenceinstancecalendar.EventExceptionOccurrenceInstanceCalendarClient
	EventExceptionOccurrenceInstanceExtension                    *eventexceptionoccurrenceinstanceextension.EventExceptionOccurrenceInstanceExtensionClient
	EventExtension                                               *eventextension.EventExtensionClient
	EventInstance                                                *eventinstance.EventInstanceClient
	EventInstanceAttachment                                      *eventinstanceattachment.EventInstanceAttachmentClient
	EventInstanceCalendar                                        *eventinstancecalendar.EventInstanceCalendarClient
	EventInstanceExceptionOccurrence                             *eventinstanceexceptionoccurrence.EventInstanceExceptionOccurrenceClient
	EventInstanceExceptionOccurrenceAttachment                   *eventinstanceexceptionoccurrenceattachment.EventInstanceExceptionOccurrenceAttachmentClient
	EventInstanceExceptionOccurrenceCalendar                     *eventinstanceexceptionoccurrencecalendar.EventInstanceExceptionOccurrenceCalendarClient
	EventInstanceExceptionOccurrenceExtension                    *eventinstanceexceptionoccurrenceextension.EventInstanceExceptionOccurrenceExtensionClient
	EventInstanceExtension                                       *eventinstanceextension.EventInstanceExtensionClient
	Extension                                                    *extension.ExtensionClient
	Group                                                        *group.GroupClient
	GroupLifecyclePolicy                                         *grouplifecyclepolicy.GroupLifecyclePolicyClient
	Member                                                       *member.MemberClient
	MemberOf                                                     *memberof.MemberOfClient
	MembersWithLicenseError                                      *memberswithlicenseerror.MembersWithLicenseErrorClient
	Owner                                                        *owner.OwnerClient
	PermissionGrant                                              *permissiongrant.PermissionGrantClient
	Photo                                                        *photo.PhotoClient
	Planner                                                      *planner.PlannerClient
	PlannerPlan                                                  *plannerplan.PlannerPlanClient
	PlannerPlanBucket                                            *plannerplanbucket.PlannerPlanBucketClient
	PlannerPlanBucketTask                                        *plannerplanbuckettask.PlannerPlanBucketTaskClient
	PlannerPlanBucketTaskAssignedToTaskBoardFormat               *plannerplanbuckettaskassignedtotaskboardformat.PlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	PlannerPlanBucketTaskBucketTaskBoardFormat                   *plannerplanbuckettaskbuckettaskboardformat.PlannerPlanBucketTaskBucketTaskBoardFormatClient
	PlannerPlanBucketTaskDetail                                  *plannerplanbuckettaskdetail.PlannerPlanBucketTaskDetailClient
	PlannerPlanBucketTaskProgressTaskBoardFormat                 *plannerplanbuckettaskprogresstaskboardformat.PlannerPlanBucketTaskProgressTaskBoardFormatClient
	PlannerPlanDetail                                            *plannerplandetail.PlannerPlanDetailClient
	PlannerPlanTask                                              *plannerplantask.PlannerPlanTaskClient
	PlannerPlanTaskAssignedToTaskBoardFormat                     *plannerplantaskassignedtotaskboardformat.PlannerPlanTaskAssignedToTaskBoardFormatClient
	PlannerPlanTaskBucketTaskBoardFormat                         *plannerplantaskbuckettaskboardformat.PlannerPlanTaskBucketTaskBoardFormatClient
	PlannerPlanTaskDetail                                        *plannerplantaskdetail.PlannerPlanTaskDetailClient
	PlannerPlanTaskProgressTaskBoardFormat                       *plannerplantaskprogresstaskboardformat.PlannerPlanTaskProgressTaskBoardFormatClient
	RejectedSender                                               *rejectedsender.RejectedSenderClient
	ServiceProvisioningError                                     *serviceprovisioningerror.ServiceProvisioningErrorClient
	Setting                                                      *setting.SettingClient
	Site                                                         *site.SiteClient
	SiteAnalytic                                                 *siteanalytic.SiteAnalyticClient
	SiteAnalyticAllTime                                          *siteanalyticalltime.SiteAnalyticAllTimeClient
	SiteAnalyticItemActivityStat                                 *siteanalyticitemactivitystat.SiteAnalyticItemActivityStatClient
	SiteAnalyticItemActivityStatActivity                         *siteanalyticitemactivitystatactivity.SiteAnalyticItemActivityStatActivityClient
	SiteAnalyticItemActivityStatActivityDriveItem                *siteanalyticitemactivitystatactivitydriveitem.SiteAnalyticItemActivityStatActivityDriveItemClient
	SiteAnalyticItemActivityStatActivityDriveItemContent         *siteanalyticitemactivitystatactivitydriveitemcontent.SiteAnalyticItemActivityStatActivityDriveItemContentClient
	SiteAnalyticLastSevenDay                                     *siteanalyticlastsevenday.SiteAnalyticLastSevenDayClient
	SiteColumn                                                   *sitecolumn.SiteColumnClient
	SiteColumnSourceColumn                                       *sitecolumnsourcecolumn.SiteColumnSourceColumnClient
	SiteContentType                                              *sitecontenttype.SiteContentTypeClient
	SiteContentTypeBase                                          *sitecontenttypebase.SiteContentTypeBaseClient
	SiteContentTypeBaseType                                      *sitecontenttypebasetype.SiteContentTypeBaseTypeClient
	SiteContentTypeColumn                                        *sitecontenttypecolumn.SiteContentTypeColumnClient
	SiteContentTypeColumnLink                                    *sitecontenttypecolumnlink.SiteContentTypeColumnLinkClient
	SiteContentTypeColumnPosition                                *sitecontenttypecolumnposition.SiteContentTypeColumnPositionClient
	SiteContentTypeColumnSourceColumn                            *sitecontenttypecolumnsourcecolumn.SiteContentTypeColumnSourceColumnClient
	SiteCreatedByUser                                            *sitecreatedbyuser.SiteCreatedByUserClient
	SiteCreatedByUserMailboxSetting                              *sitecreatedbyusermailboxsetting.SiteCreatedByUserMailboxSettingClient
	SiteCreatedByUserServiceProvisioningError                    *sitecreatedbyuserserviceprovisioningerror.SiteCreatedByUserServiceProvisioningErrorClient
	SiteDrive                                                    *sitedrive.SiteDriveClient
	SiteExternalColumn                                           *siteexternalcolumn.SiteExternalColumnClient
	SiteInformationProtection                                    *siteinformationprotection.SiteInformationProtectionClient
	SiteInformationProtectionBitlocker                           *siteinformationprotectionbitlocker.SiteInformationProtectionBitlockerClient
	SiteInformationProtectionBitlockerRecoveryKey                *siteinformationprotectionbitlockerrecoverykey.SiteInformationProtectionBitlockerRecoveryKeyClient
	SiteInformationProtectionDataLossPreventionPolicy            *siteinformationprotectiondatalosspreventionpolicy.SiteInformationProtectionDataLossPreventionPolicyClient
	SiteInformationProtectionPolicy                              *siteinformationprotectionpolicy.SiteInformationProtectionPolicyClient
	SiteInformationProtectionPolicyLabel                         *siteinformationprotectionpolicylabel.SiteInformationProtectionPolicyLabelClient
	SiteInformationProtectionSensitivityLabel                    *siteinformationprotectionsensitivitylabel.SiteInformationProtectionSensitivityLabelClient
	SiteInformationProtectionSensitivityLabelSublabel            *siteinformationprotectionsensitivitylabelsublabel.SiteInformationProtectionSensitivityLabelSublabelClient
	SiteInformationProtectionSensitivityPolicySetting            *siteinformationprotectionsensitivitypolicysetting.SiteInformationProtectionSensitivityPolicySettingClient
	SiteInformationProtectionThreatAssessmentRequest             *siteinformationprotectionthreatassessmentrequest.SiteInformationProtectionThreatAssessmentRequestClient
	SiteInformationProtectionThreatAssessmentRequestResult       *siteinformationprotectionthreatassessmentrequestresult.SiteInformationProtectionThreatAssessmentRequestResultClient
	SiteItem                                                     *siteitem.SiteItemClient
	SiteLastModifiedByUser                                       *sitelastmodifiedbyuser.SiteLastModifiedByUserClient
	SiteLastModifiedByUserMailboxSetting                         *sitelastmodifiedbyusermailboxsetting.SiteLastModifiedByUserMailboxSettingClient
	SiteLastModifiedByUserServiceProvisioningError               *sitelastmodifiedbyuserserviceprovisioningerror.SiteLastModifiedByUserServiceProvisioningErrorClient
	SiteList                                                     *sitelist.SiteListClient
	SiteListActivity                                             *sitelistactivity.SiteListActivityClient
	SiteListColumn                                               *sitelistcolumn.SiteListColumnClient
	SiteListColumnSourceColumn                                   *sitelistcolumnsourcecolumn.SiteListColumnSourceColumnClient
	SiteListContentType                                          *sitelistcontenttype.SiteListContentTypeClient
	SiteListContentTypeBase                                      *sitelistcontenttypebase.SiteListContentTypeBaseClient
	SiteListContentTypeBaseType                                  *sitelistcontenttypebasetype.SiteListContentTypeBaseTypeClient
	SiteListContentTypeColumn                                    *sitelistcontenttypecolumn.SiteListContentTypeColumnClient
	SiteListContentTypeColumnLink                                *sitelistcontenttypecolumnlink.SiteListContentTypeColumnLinkClient
	SiteListContentTypeColumnPosition                            *sitelistcontenttypecolumnposition.SiteListContentTypeColumnPositionClient
	SiteListContentTypeColumnSourceColumn                        *sitelistcontenttypecolumnsourcecolumn.SiteListContentTypeColumnSourceColumnClient
	SiteListCreatedByUser                                        *sitelistcreatedbyuser.SiteListCreatedByUserClient
	SiteListCreatedByUserMailboxSetting                          *sitelistcreatedbyusermailboxsetting.SiteListCreatedByUserMailboxSettingClient
	SiteListCreatedByUserServiceProvisioningError                *sitelistcreatedbyuserserviceprovisioningerror.SiteListCreatedByUserServiceProvisioningErrorClient
	SiteListDrive                                                *sitelistdrive.SiteListDriveClient
	SiteListItem                                                 *sitelistitem.SiteListItemClient
	SiteListItemActivity                                         *sitelistitemactivity.SiteListItemActivityClient
	SiteListItemActivityDriveItem                                *sitelistitemactivitydriveitem.SiteListItemActivityDriveItemClient
	SiteListItemActivityDriveItemContent                         *sitelistitemactivitydriveitemcontent.SiteListItemActivityDriveItemContentClient
	SiteListItemActivityListItem                                 *sitelistitemactivitylistitem.SiteListItemActivityListItemClient
	SiteListItemAnalytic                                         *sitelistitemanalytic.SiteListItemAnalyticClient
	SiteListItemCreatedByUser                                    *sitelistitemcreatedbyuser.SiteListItemCreatedByUserClient
	SiteListItemCreatedByUserMailboxSetting                      *sitelistitemcreatedbyusermailboxsetting.SiteListItemCreatedByUserMailboxSettingClient
	SiteListItemCreatedByUserServiceProvisioningError            *sitelistitemcreatedbyuserserviceprovisioningerror.SiteListItemCreatedByUserServiceProvisioningErrorClient
	SiteListItemDocumentSetVersion                               *sitelistitemdocumentsetversion.SiteListItemDocumentSetVersionClient
	SiteListItemDocumentSetVersionField                          *sitelistitemdocumentsetversionfield.SiteListItemDocumentSetVersionFieldClient
	SiteListItemDriveItem                                        *sitelistitemdriveitem.SiteListItemDriveItemClient
	SiteListItemDriveItemContent                                 *sitelistitemdriveitemcontent.SiteListItemDriveItemContentClient
	SiteListItemField                                            *sitelistitemfield.SiteListItemFieldClient
	SiteListItemLastModifiedByUser                               *sitelistitemlastmodifiedbyuser.SiteListItemLastModifiedByUserClient
	SiteListItemLastModifiedByUserMailboxSetting                 *sitelistitemlastmodifiedbyusermailboxsetting.SiteListItemLastModifiedByUserMailboxSettingClient
	SiteListItemLastModifiedByUserServiceProvisioningError       *sitelistitemlastmodifiedbyuserserviceprovisioningerror.SiteListItemLastModifiedByUserServiceProvisioningErrorClient
	SiteListItemVersion                                          *sitelistitemversion.SiteListItemVersionClient
	SiteListItemVersionField                                     *sitelistitemversionfield.SiteListItemVersionFieldClient
	SiteListLastModifiedByUser                                   *sitelistlastmodifiedbyuser.SiteListLastModifiedByUserClient
	SiteListLastModifiedByUserMailboxSetting                     *sitelistlastmodifiedbyusermailboxsetting.SiteListLastModifiedByUserMailboxSettingClient
	SiteListLastModifiedByUserServiceProvisioningError           *sitelistlastmodifiedbyuserserviceprovisioningerror.SiteListLastModifiedByUserServiceProvisioningErrorClient
	SiteListOperation                                            *sitelistoperation.SiteListOperationClient
	SiteListSubscription                                         *sitelistsubscription.SiteListSubscriptionClient
	SiteOperation                                                *siteoperation.SiteOperationClient
	SitePage                                                     *sitepage.SitePageClient
	SitePageCreatedByUser                                        *sitepagecreatedbyuser.SitePageCreatedByUserClient
	SitePageCreatedByUserMailboxSetting                          *sitepagecreatedbyusermailboxsetting.SitePageCreatedByUserMailboxSettingClient
	SitePageCreatedByUserServiceProvisioningError                *sitepagecreatedbyuserserviceprovisioningerror.SitePageCreatedByUserServiceProvisioningErrorClient
	SitePageLastModifiedByUser                                   *sitepagelastmodifiedbyuser.SitePageLastModifiedByUserClient
	SitePageLastModifiedByUserMailboxSetting                     *sitepagelastmodifiedbyusermailboxsetting.SitePageLastModifiedByUserMailboxSettingClient
	SitePageLastModifiedByUserServiceProvisioningError           *sitepagelastmodifiedbyuserserviceprovisioningerror.SitePageLastModifiedByUserServiceProvisioningErrorClient
	SitePermission                                               *sitepermission.SitePermissionClient
	SiteRecycleBin                                               *siterecyclebin.SiteRecycleBinClient
	SiteRecycleBinCreatedByUser                                  *siterecyclebincreatedbyuser.SiteRecycleBinCreatedByUserClient
	SiteRecycleBinCreatedByUserMailboxSetting                    *siterecyclebincreatedbyusermailboxsetting.SiteRecycleBinCreatedByUserMailboxSettingClient
	SiteRecycleBinCreatedByUserServiceProvisioningError          *siterecyclebincreatedbyuserserviceprovisioningerror.SiteRecycleBinCreatedByUserServiceProvisioningErrorClient
	SiteRecycleBinItem                                           *siterecyclebinitem.SiteRecycleBinItemClient
	SiteRecycleBinItemCreatedByUser                              *siterecyclebinitemcreatedbyuser.SiteRecycleBinItemCreatedByUserClient
	SiteRecycleBinItemCreatedByUserMailboxSetting                *siterecyclebinitemcreatedbyusermailboxsetting.SiteRecycleBinItemCreatedByUserMailboxSettingClient
	SiteRecycleBinItemCreatedByUserServiceProvisioningError      *siterecyclebinitemcreatedbyuserserviceprovisioningerror.SiteRecycleBinItemCreatedByUserServiceProvisioningErrorClient
	SiteRecycleBinItemLastModifiedByUser                         *siterecyclebinitemlastmodifiedbyuser.SiteRecycleBinItemLastModifiedByUserClient
	SiteRecycleBinItemLastModifiedByUserMailboxSetting           *siterecyclebinitemlastmodifiedbyusermailboxsetting.SiteRecycleBinItemLastModifiedByUserMailboxSettingClient
	SiteRecycleBinItemLastModifiedByUserServiceProvisioningError *siterecyclebinitemlastmodifiedbyuserserviceprovisioningerror.SiteRecycleBinItemLastModifiedByUserServiceProvisioningErrorClient
	SiteRecycleBinLastModifiedByUser                             *siterecyclebinlastmodifiedbyuser.SiteRecycleBinLastModifiedByUserClient
	SiteRecycleBinLastModifiedByUserMailboxSetting               *siterecyclebinlastmodifiedbyusermailboxsetting.SiteRecycleBinLastModifiedByUserMailboxSettingClient
	SiteRecycleBinLastModifiedByUserServiceProvisioningError     *siterecyclebinlastmodifiedbyuserserviceprovisioningerror.SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient
	Team                                                         *team.TeamClient
	TeamAllChannel                                               *teamallchannel.TeamAllChannelClient
	TeamChannel                                                  *teamchannel.TeamChannelClient
	TeamChannelFilesFolder                                       *teamchannelfilesfolder.TeamChannelFilesFolderClient
	TeamChannelFilesFolderContent                                *teamchannelfilesfoldercontent.TeamChannelFilesFolderContentClient
	TeamChannelMember                                            *teamchannelmember.TeamChannelMemberClient
	TeamChannelMessage                                           *teamchannelmessage.TeamChannelMessageClient
	TeamChannelMessageHostedContent                              *teamchannelmessagehostedcontent.TeamChannelMessageHostedContentClient
	TeamChannelMessageReply                                      *teamchannelmessagereply.TeamChannelMessageReplyClient
	TeamChannelMessageReplyHostedContent                         *teamchannelmessagereplyhostedcontent.TeamChannelMessageReplyHostedContentClient
	TeamChannelSharedWithTeam                                    *teamchannelsharedwithteam.TeamChannelSharedWithTeamClient
	TeamChannelSharedWithTeamAllowedMember                       *teamchannelsharedwithteamallowedmember.TeamChannelSharedWithTeamAllowedMemberClient
	TeamChannelSharedWithTeamTeam                                *teamchannelsharedwithteamteam.TeamChannelSharedWithTeamTeamClient
	TeamChannelTab                                               *teamchanneltab.TeamChannelTabClient
	TeamChannelTabTeamsApp                                       *teamchanneltabteamsapp.TeamChannelTabTeamsAppClient
	TeamGroup                                                    *teamgroup.TeamGroupClient
	TeamGroupServiceProvisioningError                            *teamgroupserviceprovisioningerror.TeamGroupServiceProvisioningErrorClient
	TeamIncomingChannel                                          *teamincomingchannel.TeamIncomingChannelClient
	TeamInstalledApp                                             *teaminstalledapp.TeamInstalledAppClient
	TeamInstalledAppTeamsApp                                     *teaminstalledappteamsapp.TeamInstalledAppTeamsAppClient
	TeamInstalledAppTeamsAppDefinition                           *teaminstalledappteamsappdefinition.TeamInstalledAppTeamsAppDefinitionClient
	TeamMember                                                   *teammember.TeamMemberClient
	TeamOperation                                                *teamoperation.TeamOperationClient
	TeamOwner                                                    *teamowner.TeamOwnerClient
	TeamOwnerMailboxSetting                                      *teamownermailboxsetting.TeamOwnerMailboxSettingClient
	TeamOwnerServiceProvisioningError                            *teamownerserviceprovisioningerror.TeamOwnerServiceProvisioningErrorClient
	TeamPermissionGrant                                          *teampermissiongrant.TeamPermissionGrantClient
	TeamPhoto                                                    *teamphoto.TeamPhotoClient
	TeamPrimaryChannel                                           *teamprimarychannel.TeamPrimaryChannelClient
	TeamPrimaryChannelFilesFolder                                *teamprimarychannelfilesfolder.TeamPrimaryChannelFilesFolderClient
	TeamPrimaryChannelFilesFolderContent                         *teamprimarychannelfilesfoldercontent.TeamPrimaryChannelFilesFolderContentClient
	TeamPrimaryChannelMember                                     *teamprimarychannelmember.TeamPrimaryChannelMemberClient
	TeamPrimaryChannelMessage                                    *teamprimarychannelmessage.TeamPrimaryChannelMessageClient
	TeamPrimaryChannelMessageHostedContent                       *teamprimarychannelmessagehostedcontent.TeamPrimaryChannelMessageHostedContentClient
	TeamPrimaryChannelMessageReply                               *teamprimarychannelmessagereply.TeamPrimaryChannelMessageReplyClient
	TeamPrimaryChannelMessageReplyHostedContent                  *teamprimarychannelmessagereplyhostedcontent.TeamPrimaryChannelMessageReplyHostedContentClient
	TeamPrimaryChannelSharedWithTeam                             *teamprimarychannelsharedwithteam.TeamPrimaryChannelSharedWithTeamClient
	TeamPrimaryChannelSharedWithTeamAllowedMember                *teamprimarychannelsharedwithteamallowedmember.TeamPrimaryChannelSharedWithTeamAllowedMemberClient
	TeamPrimaryChannelSharedWithTeamTeam                         *teamprimarychannelsharedwithteamteam.TeamPrimaryChannelSharedWithTeamTeamClient
	TeamPrimaryChannelTab                                        *teamprimarychanneltab.TeamPrimaryChannelTabClient
	TeamPrimaryChannelTabTeamsApp                                *teamprimarychanneltabteamsapp.TeamPrimaryChannelTabTeamsAppClient
	TeamSchedule                                                 *teamschedule.TeamScheduleClient
	TeamScheduleDayNote                                          *teamscheduledaynote.TeamScheduleDayNoteClient
	TeamScheduleOfferShiftRequest                                *teamscheduleoffershiftrequest.TeamScheduleOfferShiftRequestClient
	TeamScheduleOpenShift                                        *teamscheduleopenshift.TeamScheduleOpenShiftClient
	TeamScheduleOpenShiftChangeRequest                           *teamscheduleopenshiftchangerequest.TeamScheduleOpenShiftChangeRequestClient
	TeamScheduleSchedulingGroup                                  *teamscheduleschedulinggroup.TeamScheduleSchedulingGroupClient
	TeamScheduleShift                                            *teamscheduleshift.TeamScheduleShiftClient
	TeamScheduleSwapShiftsChangeRequest                          *teamscheduleswapshiftschangerequest.TeamScheduleSwapShiftsChangeRequestClient
	TeamScheduleTimeCard                                         *teamscheduletimecard.TeamScheduleTimeCardClient
	TeamScheduleTimeOffReason                                    *teamscheduletimeoffreason.TeamScheduleTimeOffReasonClient
	TeamScheduleTimeOffRequest                                   *teamscheduletimeoffrequest.TeamScheduleTimeOffRequestClient
	TeamScheduleTimesOff                                         *teamscheduletimesoff.TeamScheduleTimesOffClient
	TeamTag                                                      *teamtag.TeamTagClient
	TeamTagMember                                                *teamtagmember.TeamTagMemberClient
	TeamTemplate                                                 *teamtemplate.TeamTemplateClient
	TeamTemplateDefinition                                       *teamtemplatedefinition.TeamTemplateDefinitionClient
	Thread                                                       *thread.ThreadClient
	ThreadPost                                                   *threadpost.ThreadPostClient
	ThreadPostAttachment                                         *threadpostattachment.ThreadPostAttachmentClient
	ThreadPostExtension                                          *threadpostextension.ThreadPostExtensionClient
	ThreadPostInReplyTo                                          *threadpostinreplyto.ThreadPostInReplyToClient
	ThreadPostInReplyToAttachment                                *threadpostinreplytoattachment.ThreadPostInReplyToAttachmentClient
	ThreadPostInReplyToExtension                                 *threadpostinreplytoextension.ThreadPostInReplyToExtensionClient
	ThreadPostInReplyToMention                                   *threadpostinreplytomention.ThreadPostInReplyToMentionClient
	ThreadPostMention                                            *threadpostmention.ThreadPostMentionClient
	TransitiveMember                                             *transitivemember.TransitiveMemberClient
	TransitiveMemberOf                                           *transitivememberof.TransitiveMemberOfClient
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

	driveClient, err := drive.NewDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Drive client: %+v", err)
	}
	configureFunc(driveClient.Client)

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

	siteAnalyticAllTimeClient, err := siteanalyticalltime.NewSiteAnalyticAllTimeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticAllTime client: %+v", err)
	}
	configureFunc(siteAnalyticAllTimeClient.Client)

	siteAnalyticClient, err := siteanalytic.NewSiteAnalyticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalytic client: %+v", err)
	}
	configureFunc(siteAnalyticClient.Client)

	siteAnalyticItemActivityStatActivityClient, err := siteanalyticitemactivitystatactivity.NewSiteAnalyticItemActivityStatActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticItemActivityStatActivity client: %+v", err)
	}
	configureFunc(siteAnalyticItemActivityStatActivityClient.Client)

	siteAnalyticItemActivityStatActivityDriveItemClient, err := siteanalyticitemactivitystatactivitydriveitem.NewSiteAnalyticItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticItemActivityStatActivityDriveItem client: %+v", err)
	}
	configureFunc(siteAnalyticItemActivityStatActivityDriveItemClient.Client)

	siteAnalyticItemActivityStatActivityDriveItemContentClient, err := siteanalyticitemactivitystatactivitydriveitemcontent.NewSiteAnalyticItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticItemActivityStatActivityDriveItemContent client: %+v", err)
	}
	configureFunc(siteAnalyticItemActivityStatActivityDriveItemContentClient.Client)

	siteAnalyticItemActivityStatClient, err := siteanalyticitemactivitystat.NewSiteAnalyticItemActivityStatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticItemActivityStat client: %+v", err)
	}
	configureFunc(siteAnalyticItemActivityStatClient.Client)

	siteAnalyticLastSevenDayClient, err := siteanalyticlastsevenday.NewSiteAnalyticLastSevenDayClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteAnalyticLastSevenDay client: %+v", err)
	}
	configureFunc(siteAnalyticLastSevenDayClient.Client)

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

	siteListItemActivityListItemClient, err := sitelistitemactivitylistitem.NewSiteListItemActivityListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemActivityListItem client: %+v", err)
	}
	configureFunc(siteListItemActivityListItemClient.Client)

	siteListItemAnalyticClient, err := sitelistitemanalytic.NewSiteListItemAnalyticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SiteListItemAnalytic client: %+v", err)
	}
	configureFunc(siteListItemAnalyticClient.Client)

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
		AcceptedSender:                                            acceptedSenderClient,
		AppRoleAssignment:                                         appRoleAssignmentClient,
		Calendar:                                                  calendarClient,
		CalendarCalendarPermission:                                calendarCalendarPermissionClient,
		CalendarCalendarView:                                      calendarCalendarViewClient,
		CalendarCalendarViewAttachment:                            calendarCalendarViewAttachmentClient,
		CalendarCalendarViewCalendar:                              calendarCalendarViewCalendarClient,
		CalendarCalendarViewExceptionOccurrence:                   calendarCalendarViewExceptionOccurrenceClient,
		CalendarCalendarViewExceptionOccurrenceAttachment:         calendarCalendarViewExceptionOccurrenceAttachmentClient,
		CalendarCalendarViewExceptionOccurrenceCalendar:           calendarCalendarViewExceptionOccurrenceCalendarClient,
		CalendarCalendarViewExceptionOccurrenceExtension:          calendarCalendarViewExceptionOccurrenceExtensionClient,
		CalendarCalendarViewExceptionOccurrenceInstance:           calendarCalendarViewExceptionOccurrenceInstanceClient,
		CalendarCalendarViewExceptionOccurrenceInstanceAttachment: calendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		CalendarCalendarViewExceptionOccurrenceInstanceCalendar:   calendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		CalendarCalendarViewExceptionOccurrenceInstanceExtension:  calendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		CalendarCalendarViewExtension:                             calendarCalendarViewExtensionClient,
		CalendarCalendarViewInstance:                              calendarCalendarViewInstanceClient,
		CalendarCalendarViewInstanceAttachment:                    calendarCalendarViewInstanceAttachmentClient,
		CalendarCalendarViewInstanceCalendar:                      calendarCalendarViewInstanceCalendarClient,
		CalendarCalendarViewInstanceExceptionOccurrence:           calendarCalendarViewInstanceExceptionOccurrenceClient,
		CalendarCalendarViewInstanceExceptionOccurrenceAttachment: calendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		CalendarCalendarViewInstanceExceptionOccurrenceCalendar:   calendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		CalendarCalendarViewInstanceExceptionOccurrenceExtension:  calendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		CalendarCalendarViewInstanceExtension:                     calendarCalendarViewInstanceExtensionClient,
		CalendarEvent:                                             calendarEventClient,
		CalendarEventAttachment:                                   calendarEventAttachmentClient,
		CalendarEventCalendar:                                     calendarEventCalendarClient,
		CalendarEventExceptionOccurrence:                          calendarEventExceptionOccurrenceClient,
		CalendarEventExceptionOccurrenceAttachment:                calendarEventExceptionOccurrenceAttachmentClient,
		CalendarEventExceptionOccurrenceCalendar:                  calendarEventExceptionOccurrenceCalendarClient,
		CalendarEventExceptionOccurrenceExtension:                 calendarEventExceptionOccurrenceExtensionClient,
		CalendarEventExceptionOccurrenceInstance:                  calendarEventExceptionOccurrenceInstanceClient,
		CalendarEventExceptionOccurrenceInstanceAttachment:        calendarEventExceptionOccurrenceInstanceAttachmentClient,
		CalendarEventExceptionOccurrenceInstanceCalendar:          calendarEventExceptionOccurrenceInstanceCalendarClient,
		CalendarEventExceptionOccurrenceInstanceExtension:         calendarEventExceptionOccurrenceInstanceExtensionClient,
		CalendarEventExtension:                                    calendarEventExtensionClient,
		CalendarEventInstance:                                     calendarEventInstanceClient,
		CalendarEventInstanceAttachment:                           calendarEventInstanceAttachmentClient,
		CalendarEventInstanceCalendar:                             calendarEventInstanceCalendarClient,
		CalendarEventInstanceExceptionOccurrence:                  calendarEventInstanceExceptionOccurrenceClient,
		CalendarEventInstanceExceptionOccurrenceAttachment:        calendarEventInstanceExceptionOccurrenceAttachmentClient,
		CalendarEventInstanceExceptionOccurrenceCalendar:          calendarEventInstanceExceptionOccurrenceCalendarClient,
		CalendarEventInstanceExceptionOccurrenceExtension:         calendarEventInstanceExceptionOccurrenceExtensionClient,
		CalendarEventInstanceExtension:                            calendarEventInstanceExtensionClient,
		CalendarView:                                              calendarViewClient,
		CalendarViewAttachment:                                    calendarViewAttachmentClient,
		CalendarViewCalendar:                                      calendarViewCalendarClient,
		CalendarViewExceptionOccurrence:                           calendarViewExceptionOccurrenceClient,
		CalendarViewExceptionOccurrenceAttachment:                 calendarViewExceptionOccurrenceAttachmentClient,
		CalendarViewExceptionOccurrenceCalendar:                   calendarViewExceptionOccurrenceCalendarClient,
		CalendarViewExceptionOccurrenceExtension:                  calendarViewExceptionOccurrenceExtensionClient,
		CalendarViewExceptionOccurrenceInstance:                   calendarViewExceptionOccurrenceInstanceClient,
		CalendarViewExceptionOccurrenceInstanceAttachment:         calendarViewExceptionOccurrenceInstanceAttachmentClient,
		CalendarViewExceptionOccurrenceInstanceCalendar:           calendarViewExceptionOccurrenceInstanceCalendarClient,
		CalendarViewExceptionOccurrenceInstanceExtension:          calendarViewExceptionOccurrenceInstanceExtensionClient,
		CalendarViewExtension:                                     calendarViewExtensionClient,
		CalendarViewInstance:                                      calendarViewInstanceClient,
		CalendarViewInstanceAttachment:                            calendarViewInstanceAttachmentClient,
		CalendarViewInstanceCalendar:                              calendarViewInstanceCalendarClient,
		CalendarViewInstanceExceptionOccurrence:                   calendarViewInstanceExceptionOccurrenceClient,
		CalendarViewInstanceExceptionOccurrenceAttachment:         calendarViewInstanceExceptionOccurrenceAttachmentClient,
		CalendarViewInstanceExceptionOccurrenceCalendar:           calendarViewInstanceExceptionOccurrenceCalendarClient,
		CalendarViewInstanceExceptionOccurrenceExtension:          calendarViewInstanceExceptionOccurrenceExtensionClient,
		CalendarViewInstanceExtension:                             calendarViewInstanceExtensionClient,
		Conversation:                                              conversationClient,
		ConversationThread:                                        conversationThreadClient,
		ConversationThreadPost:                                    conversationThreadPostClient,
		ConversationThreadPostAttachment:                          conversationThreadPostAttachmentClient,
		ConversationThreadPostExtension:                           conversationThreadPostExtensionClient,
		ConversationThreadPostInReplyTo:                           conversationThreadPostInReplyToClient,
		ConversationThreadPostInReplyToAttachment:                 conversationThreadPostInReplyToAttachmentClient,
		ConversationThreadPostInReplyToExtension:                  conversationThreadPostInReplyToExtensionClient,
		ConversationThreadPostInReplyToMention:                    conversationThreadPostInReplyToMentionClient,
		ConversationThreadPostMention:                             conversationThreadPostMentionClient,
		CreatedOnBehalfOf:                                         createdOnBehalfOfClient,
		Drive:                                                     driveClient,
		Endpoint:                                                  endpointClient,
		Event:                                                     eventClient,
		EventAttachment:                                           eventAttachmentClient,
		EventCalendar:                                             eventCalendarClient,
		EventExceptionOccurrence:                                  eventExceptionOccurrenceClient,
		EventExceptionOccurrenceAttachment:                        eventExceptionOccurrenceAttachmentClient,
		EventExceptionOccurrenceCalendar:                          eventExceptionOccurrenceCalendarClient,
		EventExceptionOccurrenceExtension:                         eventExceptionOccurrenceExtensionClient,
		EventExceptionOccurrenceInstance:                          eventExceptionOccurrenceInstanceClient,
		EventExceptionOccurrenceInstanceAttachment:                eventExceptionOccurrenceInstanceAttachmentClient,
		EventExceptionOccurrenceInstanceCalendar:                  eventExceptionOccurrenceInstanceCalendarClient,
		EventExceptionOccurrenceInstanceExtension:                 eventExceptionOccurrenceInstanceExtensionClient,
		EventExtension:                                            eventExtensionClient,
		EventInstance:                                             eventInstanceClient,
		EventInstanceAttachment:                                   eventInstanceAttachmentClient,
		EventInstanceCalendar:                                     eventInstanceCalendarClient,
		EventInstanceExceptionOccurrence:                          eventInstanceExceptionOccurrenceClient,
		EventInstanceExceptionOccurrenceAttachment:                eventInstanceExceptionOccurrenceAttachmentClient,
		EventInstanceExceptionOccurrenceCalendar:                  eventInstanceExceptionOccurrenceCalendarClient,
		EventInstanceExceptionOccurrenceExtension:                 eventInstanceExceptionOccurrenceExtensionClient,
		EventInstanceExtension:                                    eventInstanceExtensionClient,
		Extension:                                                 extensionClient,
		Group:                                                     groupClient,
		GroupLifecyclePolicy:                                      groupLifecyclePolicyClient,
		Member:                                                    memberClient,
		MemberOf:                                                  memberOfClient,
		MembersWithLicenseError:                                   membersWithLicenseErrorClient,
		Owner:                                                     ownerClient,
		PermissionGrant:                                           permissionGrantClient,
		Photo:                                                     photoClient,
		Planner:                                                   plannerClient,
		PlannerPlan:                                               plannerPlanClient,
		PlannerPlanBucket:                                         plannerPlanBucketClient,
		PlannerPlanBucketTask:                                     plannerPlanBucketTaskClient,
		PlannerPlanBucketTaskAssignedToTaskBoardFormat:            plannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		PlannerPlanBucketTaskBucketTaskBoardFormat:                plannerPlanBucketTaskBucketTaskBoardFormatClient,
		PlannerPlanBucketTaskDetail:                               plannerPlanBucketTaskDetailClient,
		PlannerPlanBucketTaskProgressTaskBoardFormat:              plannerPlanBucketTaskProgressTaskBoardFormatClient,
		PlannerPlanDetail:                                         plannerPlanDetailClient,
		PlannerPlanTask:                                           plannerPlanTaskClient,
		PlannerPlanTaskAssignedToTaskBoardFormat:                  plannerPlanTaskAssignedToTaskBoardFormatClient,
		PlannerPlanTaskBucketTaskBoardFormat:                      plannerPlanTaskBucketTaskBoardFormatClient,
		PlannerPlanTaskDetail:                                     plannerPlanTaskDetailClient,
		PlannerPlanTaskProgressTaskBoardFormat:                    plannerPlanTaskProgressTaskBoardFormatClient,
		RejectedSender:                                            rejectedSenderClient,
		ServiceProvisioningError:                                  serviceProvisioningErrorClient,
		Setting:                                                   settingClient,
		Site:                                                      siteClient,
		SiteAnalytic:                                              siteAnalyticClient,
		SiteAnalyticAllTime:                                       siteAnalyticAllTimeClient,
		SiteAnalyticItemActivityStat:                              siteAnalyticItemActivityStatClient,
		SiteAnalyticItemActivityStatActivity:                      siteAnalyticItemActivityStatActivityClient,
		SiteAnalyticItemActivityStatActivityDriveItem:             siteAnalyticItemActivityStatActivityDriveItemClient,
		SiteAnalyticItemActivityStatActivityDriveItemContent:      siteAnalyticItemActivityStatActivityDriveItemContentClient,
		SiteAnalyticLastSevenDay:                                  siteAnalyticLastSevenDayClient,
		SiteColumn:                                                siteColumnClient,
		SiteColumnSourceColumn:                                    siteColumnSourceColumnClient,
		SiteContentType:                                           siteContentTypeClient,
		SiteContentTypeBase:                                       siteContentTypeBaseClient,
		SiteContentTypeBaseType:                                   siteContentTypeBaseTypeClient,
		SiteContentTypeColumn:                                     siteContentTypeColumnClient,
		SiteContentTypeColumnLink:                                 siteContentTypeColumnLinkClient,
		SiteContentTypeColumnPosition:                             siteContentTypeColumnPositionClient,
		SiteContentTypeColumnSourceColumn:                         siteContentTypeColumnSourceColumnClient,
		SiteCreatedByUser:                                         siteCreatedByUserClient,
		SiteCreatedByUserMailboxSetting:                           siteCreatedByUserMailboxSettingClient,
		SiteCreatedByUserServiceProvisioningError:                 siteCreatedByUserServiceProvisioningErrorClient,
		SiteDrive:                                              siteDriveClient,
		SiteExternalColumn:                                     siteExternalColumnClient,
		SiteInformationProtection:                              siteInformationProtectionClient,
		SiteInformationProtectionBitlocker:                     siteInformationProtectionBitlockerClient,
		SiteInformationProtectionBitlockerRecoveryKey:          siteInformationProtectionBitlockerRecoveryKeyClient,
		SiteInformationProtectionDataLossPreventionPolicy:      siteInformationProtectionDataLossPreventionPolicyClient,
		SiteInformationProtectionPolicy:                        siteInformationProtectionPolicyClient,
		SiteInformationProtectionPolicyLabel:                   siteInformationProtectionPolicyLabelClient,
		SiteInformationProtectionSensitivityLabel:              siteInformationProtectionSensitivityLabelClient,
		SiteInformationProtectionSensitivityLabelSublabel:      siteInformationProtectionSensitivityLabelSublabelClient,
		SiteInformationProtectionSensitivityPolicySetting:      siteInformationProtectionSensitivityPolicySettingClient,
		SiteInformationProtectionThreatAssessmentRequest:       siteInformationProtectionThreatAssessmentRequestClient,
		SiteInformationProtectionThreatAssessmentRequestResult: siteInformationProtectionThreatAssessmentRequestResultClient,
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
		SiteListItemActivityListItem:                                 siteListItemActivityListItemClient,
		SiteListItemAnalytic:                                         siteListItemAnalyticClient,
		SiteListItemCreatedByUser:                                    siteListItemCreatedByUserClient,
		SiteListItemCreatedByUserMailboxSetting:                      siteListItemCreatedByUserMailboxSettingClient,
		SiteListItemCreatedByUserServiceProvisioningError:            siteListItemCreatedByUserServiceProvisioningErrorClient,
		SiteListItemDocumentSetVersion:                               siteListItemDocumentSetVersionClient,
		SiteListItemDocumentSetVersionField:                          siteListItemDocumentSetVersionFieldClient,
		SiteListItemDriveItem:                                        siteListItemDriveItemClient,
		SiteListItemDriveItemContent:                                 siteListItemDriveItemContentClient,
		SiteListItemField:                                            siteListItemFieldClient,
		SiteListItemLastModifiedByUser:                               siteListItemLastModifiedByUserClient,
		SiteListItemLastModifiedByUserMailboxSetting:                 siteListItemLastModifiedByUserMailboxSettingClient,
		SiteListItemLastModifiedByUserServiceProvisioningError:       siteListItemLastModifiedByUserServiceProvisioningErrorClient,
		SiteListItemVersion:                                          siteListItemVersionClient,
		SiteListItemVersionField:                                     siteListItemVersionFieldClient,
		SiteListLastModifiedByUser:                                   siteListLastModifiedByUserClient,
		SiteListLastModifiedByUserMailboxSetting:                     siteListLastModifiedByUserMailboxSettingClient,
		SiteListLastModifiedByUserServiceProvisioningError:           siteListLastModifiedByUserServiceProvisioningErrorClient,
		SiteListOperation:                                            siteListOperationClient,
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
