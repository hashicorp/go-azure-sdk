package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupacceptedsender"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarpermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendareventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcalendarviewinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthread"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpost"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpostattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpostextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpostinreplyto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpostinreplytoattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpostinreplytoextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpostinreplytomention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupconversationthreadpostmention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupcreatedonbehalfof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupendpoint"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupevent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventcalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrenceinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrenceinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrenceinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventexceptionoccurrenceinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstanceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstancecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstanceexceptionoccurrence"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstanceexceptionoccurrenceattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstanceexceptionoccurrencecalendar"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstanceexceptionoccurrenceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupeventinstanceextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupgrouplifecyclepolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupmemberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupmemberswithlicenseerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotenotebooksectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenoteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotepagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotepageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotepageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenoteresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenoteresourcecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouponenotesectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouppermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplanner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplan"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplanbucket"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplanbuckettask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplanbuckettaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplanbuckettaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplanbuckettaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplanbuckettaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplandetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplantask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplantaskassignedtotaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplantaskbuckettaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplantaskdetail"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupplannerplantaskprogresstaskboardformat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouprejectedsender"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteanalytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteanalyticalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteanalyticitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteanalyticitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteanalyticitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteanalyticitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteanalyticlastsevenday"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitedrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteexternalcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionbitlocker"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionbitlockerrecoverykey"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectiondatalosspreventionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionpolicylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionsensitivitylabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionsensitivitylabelsublabel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionsensitivitypolicysetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionthreatassessmentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteinformationprotectionthreatassessmentrequestresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelist"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcontenttype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcontenttypebase"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcontenttypebasetype"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcontenttypecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcontenttypecolumnlink"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcontenttypecolumnposition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcontenttypecolumnsourcecolumn"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistdrive"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemactivitylistitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemanalytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemcreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemcreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemdocumentsetversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemdocumentsetversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemdriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemdriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistitemversionfield"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistlastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistlastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitelistsubscription"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenote"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotenotebooksectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenoteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotepagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotepageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotepageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenoteresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenoteresourcecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectiongroupsectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectionpage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectionpagecontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectionpageparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectionpageparentsection"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectionparentnotebook"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteonenotesectionparentsectiongroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsiteoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitepage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitepagecreatedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitepagecreatedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitepagelastmodifiedbyuser"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitepagelastmodifiedbyusermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitepermission"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitesite"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstore"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildrenchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildrenchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildrenchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildrenchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetparentgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsetterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsettermchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsettermchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsettermchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsettermchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsettermrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsettermrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoregroupsettermrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoreset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildrenchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildrenchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildrenchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildrenchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupset"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildrenchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildrenchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildrenchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildrenchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsetterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsettermchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsettermchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsettermchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsettermchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsettermrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsettermrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetparentgroupsettermrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresetterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresettermchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresettermchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresettermchildrenrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresettermchildrenrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresettermrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresettermrelationfromterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupsitetermstoresettermrelationtoterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamallchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamchanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamincomingchannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteaminstalledapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteaminstalledappteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteaminstalledappteamsappdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteammember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamoperation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamowner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamownermailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteampermissiongrant"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamphoto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannel"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelfilesfolder"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelfilesfoldercontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelmessage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelmessagehostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelmessagereply"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelmessagereplyhostedcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelsharedwithteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelsharedwithteamallowedmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychannelsharedwithteamteam"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychanneltab"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamprimarychanneltabteamsapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduleoffershiftrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduleopenshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduleopenshiftchangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduleschedulinggroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduleshift"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduleswapshiftschangerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduletimecard"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduletimeoffreason"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduletimeoffrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamscheduletimesoff"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamtag"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamtagmember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupteamtemplatedefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthread"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpost"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpostattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpostextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpostinreplyto"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpostinreplytoattachment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpostinreplytoextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpostinreplytomention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/groupthreadpostmention"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouptransitivemember"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/grouptransitivememberof"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsetchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsetchildrenchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsetchildrenchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsetchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsetrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsetterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsettermchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsettermchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoregroupsettermrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetchildrenchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetchildrenchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsetchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsetchildrenchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsetchildrenchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsetchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsetrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsetterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsettermchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsettermchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetparentgroupsettermrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresetterm"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresettermchildren"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresettermchildrenrelation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/setgroupsitetermstoresettermrelation"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	Group                                                               *group.GroupClient
	GroupAcceptedSender                                                 *groupacceptedsender.GroupAcceptedSenderClient
	GroupAppRoleAssignment                                              *groupapproleassignment.GroupAppRoleAssignmentClient
	GroupCalendar                                                       *groupcalendar.GroupCalendarClient
	GroupCalendarCalendarPermission                                     *groupcalendarcalendarpermission.GroupCalendarCalendarPermissionClient
	GroupCalendarCalendarView                                           *groupcalendarcalendarview.GroupCalendarCalendarViewClient
	GroupCalendarCalendarViewAttachment                                 *groupcalendarcalendarviewattachment.GroupCalendarCalendarViewAttachmentClient
	GroupCalendarCalendarViewCalendar                                   *groupcalendarcalendarviewcalendar.GroupCalendarCalendarViewCalendarClient
	GroupCalendarCalendarViewExceptionOccurrence                        *groupcalendarcalendarviewexceptionoccurrence.GroupCalendarCalendarViewExceptionOccurrenceClient
	GroupCalendarCalendarViewExceptionOccurrenceAttachment              *groupcalendarcalendarviewexceptionoccurrenceattachment.GroupCalendarCalendarViewExceptionOccurrenceAttachmentClient
	GroupCalendarCalendarViewExceptionOccurrenceCalendar                *groupcalendarcalendarviewexceptionoccurrencecalendar.GroupCalendarCalendarViewExceptionOccurrenceCalendarClient
	GroupCalendarCalendarViewExceptionOccurrenceExtension               *groupcalendarcalendarviewexceptionoccurrenceextension.GroupCalendarCalendarViewExceptionOccurrenceExtensionClient
	GroupCalendarCalendarViewExceptionOccurrenceInstance                *groupcalendarcalendarviewexceptionoccurrenceinstance.GroupCalendarCalendarViewExceptionOccurrenceInstanceClient
	GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment      *groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient
	GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar        *groupcalendarcalendarviewexceptionoccurrenceinstancecalendar.GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient
	GroupCalendarCalendarViewExceptionOccurrenceInstanceExtension       *groupcalendarcalendarviewexceptionoccurrenceinstanceextension.GroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient
	GroupCalendarCalendarViewExtension                                  *groupcalendarcalendarviewextension.GroupCalendarCalendarViewExtensionClient
	GroupCalendarCalendarViewInstance                                   *groupcalendarcalendarviewinstance.GroupCalendarCalendarViewInstanceClient
	GroupCalendarCalendarViewInstanceAttachment                         *groupcalendarcalendarviewinstanceattachment.GroupCalendarCalendarViewInstanceAttachmentClient
	GroupCalendarCalendarViewInstanceCalendar                           *groupcalendarcalendarviewinstancecalendar.GroupCalendarCalendarViewInstanceCalendarClient
	GroupCalendarCalendarViewInstanceExceptionOccurrence                *groupcalendarcalendarviewinstanceexceptionoccurrence.GroupCalendarCalendarViewInstanceExceptionOccurrenceClient
	GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment      *groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient
	GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar        *groupcalendarcalendarviewinstanceexceptionoccurrencecalendar.GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient
	GroupCalendarCalendarViewInstanceExceptionOccurrenceExtension       *groupcalendarcalendarviewinstanceexceptionoccurrenceextension.GroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient
	GroupCalendarCalendarViewInstanceExtension                          *groupcalendarcalendarviewinstanceextension.GroupCalendarCalendarViewInstanceExtensionClient
	GroupCalendarEvent                                                  *groupcalendarevent.GroupCalendarEventClient
	GroupCalendarEventAttachment                                        *groupcalendareventattachment.GroupCalendarEventAttachmentClient
	GroupCalendarEventCalendar                                          *groupcalendareventcalendar.GroupCalendarEventCalendarClient
	GroupCalendarEventExceptionOccurrence                               *groupcalendareventexceptionoccurrence.GroupCalendarEventExceptionOccurrenceClient
	GroupCalendarEventExceptionOccurrenceAttachment                     *groupcalendareventexceptionoccurrenceattachment.GroupCalendarEventExceptionOccurrenceAttachmentClient
	GroupCalendarEventExceptionOccurrenceCalendar                       *groupcalendareventexceptionoccurrencecalendar.GroupCalendarEventExceptionOccurrenceCalendarClient
	GroupCalendarEventExceptionOccurrenceExtension                      *groupcalendareventexceptionoccurrenceextension.GroupCalendarEventExceptionOccurrenceExtensionClient
	GroupCalendarEventExceptionOccurrenceInstance                       *groupcalendareventexceptionoccurrenceinstance.GroupCalendarEventExceptionOccurrenceInstanceClient
	GroupCalendarEventExceptionOccurrenceInstanceAttachment             *groupcalendareventexceptionoccurrenceinstanceattachment.GroupCalendarEventExceptionOccurrenceInstanceAttachmentClient
	GroupCalendarEventExceptionOccurrenceInstanceCalendar               *groupcalendareventexceptionoccurrenceinstancecalendar.GroupCalendarEventExceptionOccurrenceInstanceCalendarClient
	GroupCalendarEventExceptionOccurrenceInstanceExtension              *groupcalendareventexceptionoccurrenceinstanceextension.GroupCalendarEventExceptionOccurrenceInstanceExtensionClient
	GroupCalendarEventExtension                                         *groupcalendareventextension.GroupCalendarEventExtensionClient
	GroupCalendarEventInstance                                          *groupcalendareventinstance.GroupCalendarEventInstanceClient
	GroupCalendarEventInstanceAttachment                                *groupcalendareventinstanceattachment.GroupCalendarEventInstanceAttachmentClient
	GroupCalendarEventInstanceCalendar                                  *groupcalendareventinstancecalendar.GroupCalendarEventInstanceCalendarClient
	GroupCalendarEventInstanceExceptionOccurrence                       *groupcalendareventinstanceexceptionoccurrence.GroupCalendarEventInstanceExceptionOccurrenceClient
	GroupCalendarEventInstanceExceptionOccurrenceAttachment             *groupcalendareventinstanceexceptionoccurrenceattachment.GroupCalendarEventInstanceExceptionOccurrenceAttachmentClient
	GroupCalendarEventInstanceExceptionOccurrenceCalendar               *groupcalendareventinstanceexceptionoccurrencecalendar.GroupCalendarEventInstanceExceptionOccurrenceCalendarClient
	GroupCalendarEventInstanceExceptionOccurrenceExtension              *groupcalendareventinstanceexceptionoccurrenceextension.GroupCalendarEventInstanceExceptionOccurrenceExtensionClient
	GroupCalendarEventInstanceExtension                                 *groupcalendareventinstanceextension.GroupCalendarEventInstanceExtensionClient
	GroupCalendarView                                                   *groupcalendarview.GroupCalendarViewClient
	GroupCalendarViewAttachment                                         *groupcalendarviewattachment.GroupCalendarViewAttachmentClient
	GroupCalendarViewCalendar                                           *groupcalendarviewcalendar.GroupCalendarViewCalendarClient
	GroupCalendarViewExceptionOccurrence                                *groupcalendarviewexceptionoccurrence.GroupCalendarViewExceptionOccurrenceClient
	GroupCalendarViewExceptionOccurrenceAttachment                      *groupcalendarviewexceptionoccurrenceattachment.GroupCalendarViewExceptionOccurrenceAttachmentClient
	GroupCalendarViewExceptionOccurrenceCalendar                        *groupcalendarviewexceptionoccurrencecalendar.GroupCalendarViewExceptionOccurrenceCalendarClient
	GroupCalendarViewExceptionOccurrenceExtension                       *groupcalendarviewexceptionoccurrenceextension.GroupCalendarViewExceptionOccurrenceExtensionClient
	GroupCalendarViewExceptionOccurrenceInstance                        *groupcalendarviewexceptionoccurrenceinstance.GroupCalendarViewExceptionOccurrenceInstanceClient
	GroupCalendarViewExceptionOccurrenceInstanceAttachment              *groupcalendarviewexceptionoccurrenceinstanceattachment.GroupCalendarViewExceptionOccurrenceInstanceAttachmentClient
	GroupCalendarViewExceptionOccurrenceInstanceCalendar                *groupcalendarviewexceptionoccurrenceinstancecalendar.GroupCalendarViewExceptionOccurrenceInstanceCalendarClient
	GroupCalendarViewExceptionOccurrenceInstanceExtension               *groupcalendarviewexceptionoccurrenceinstanceextension.GroupCalendarViewExceptionOccurrenceInstanceExtensionClient
	GroupCalendarViewExtension                                          *groupcalendarviewextension.GroupCalendarViewExtensionClient
	GroupCalendarViewInstance                                           *groupcalendarviewinstance.GroupCalendarViewInstanceClient
	GroupCalendarViewInstanceAttachment                                 *groupcalendarviewinstanceattachment.GroupCalendarViewInstanceAttachmentClient
	GroupCalendarViewInstanceCalendar                                   *groupcalendarviewinstancecalendar.GroupCalendarViewInstanceCalendarClient
	GroupCalendarViewInstanceExceptionOccurrence                        *groupcalendarviewinstanceexceptionoccurrence.GroupCalendarViewInstanceExceptionOccurrenceClient
	GroupCalendarViewInstanceExceptionOccurrenceAttachment              *groupcalendarviewinstanceexceptionoccurrenceattachment.GroupCalendarViewInstanceExceptionOccurrenceAttachmentClient
	GroupCalendarViewInstanceExceptionOccurrenceCalendar                *groupcalendarviewinstanceexceptionoccurrencecalendar.GroupCalendarViewInstanceExceptionOccurrenceCalendarClient
	GroupCalendarViewInstanceExceptionOccurrenceExtension               *groupcalendarviewinstanceexceptionoccurrenceextension.GroupCalendarViewInstanceExceptionOccurrenceExtensionClient
	GroupCalendarViewInstanceExtension                                  *groupcalendarviewinstanceextension.GroupCalendarViewInstanceExtensionClient
	GroupConversation                                                   *groupconversation.GroupConversationClient
	GroupConversationThread                                             *groupconversationthread.GroupConversationThreadClient
	GroupConversationThreadPost                                         *groupconversationthreadpost.GroupConversationThreadPostClient
	GroupConversationThreadPostAttachment                               *groupconversationthreadpostattachment.GroupConversationThreadPostAttachmentClient
	GroupConversationThreadPostExtension                                *groupconversationthreadpostextension.GroupConversationThreadPostExtensionClient
	GroupConversationThreadPostInReplyTo                                *groupconversationthreadpostinreplyto.GroupConversationThreadPostInReplyToClient
	GroupConversationThreadPostInReplyToAttachment                      *groupconversationthreadpostinreplytoattachment.GroupConversationThreadPostInReplyToAttachmentClient
	GroupConversationThreadPostInReplyToExtension                       *groupconversationthreadpostinreplytoextension.GroupConversationThreadPostInReplyToExtensionClient
	GroupConversationThreadPostInReplyToMention                         *groupconversationthreadpostinreplytomention.GroupConversationThreadPostInReplyToMentionClient
	GroupConversationThreadPostMention                                  *groupconversationthreadpostmention.GroupConversationThreadPostMentionClient
	GroupCreatedOnBehalfOf                                              *groupcreatedonbehalfof.GroupCreatedOnBehalfOfClient
	GroupDrive                                                          *groupdrive.GroupDriveClient
	GroupEndpoint                                                       *groupendpoint.GroupEndpointClient
	GroupEvent                                                          *groupevent.GroupEventClient
	GroupEventAttachment                                                *groupeventattachment.GroupEventAttachmentClient
	GroupEventCalendar                                                  *groupeventcalendar.GroupEventCalendarClient
	GroupEventExceptionOccurrence                                       *groupeventexceptionoccurrence.GroupEventExceptionOccurrenceClient
	GroupEventExceptionOccurrenceAttachment                             *groupeventexceptionoccurrenceattachment.GroupEventExceptionOccurrenceAttachmentClient
	GroupEventExceptionOccurrenceCalendar                               *groupeventexceptionoccurrencecalendar.GroupEventExceptionOccurrenceCalendarClient
	GroupEventExceptionOccurrenceExtension                              *groupeventexceptionoccurrenceextension.GroupEventExceptionOccurrenceExtensionClient
	GroupEventExceptionOccurrenceInstance                               *groupeventexceptionoccurrenceinstance.GroupEventExceptionOccurrenceInstanceClient
	GroupEventExceptionOccurrenceInstanceAttachment                     *groupeventexceptionoccurrenceinstanceattachment.GroupEventExceptionOccurrenceInstanceAttachmentClient
	GroupEventExceptionOccurrenceInstanceCalendar                       *groupeventexceptionoccurrenceinstancecalendar.GroupEventExceptionOccurrenceInstanceCalendarClient
	GroupEventExceptionOccurrenceInstanceExtension                      *groupeventexceptionoccurrenceinstanceextension.GroupEventExceptionOccurrenceInstanceExtensionClient
	GroupEventExtension                                                 *groupeventextension.GroupEventExtensionClient
	GroupEventInstance                                                  *groupeventinstance.GroupEventInstanceClient
	GroupEventInstanceAttachment                                        *groupeventinstanceattachment.GroupEventInstanceAttachmentClient
	GroupEventInstanceCalendar                                          *groupeventinstancecalendar.GroupEventInstanceCalendarClient
	GroupEventInstanceExceptionOccurrence                               *groupeventinstanceexceptionoccurrence.GroupEventInstanceExceptionOccurrenceClient
	GroupEventInstanceExceptionOccurrenceAttachment                     *groupeventinstanceexceptionoccurrenceattachment.GroupEventInstanceExceptionOccurrenceAttachmentClient
	GroupEventInstanceExceptionOccurrenceCalendar                       *groupeventinstanceexceptionoccurrencecalendar.GroupEventInstanceExceptionOccurrenceCalendarClient
	GroupEventInstanceExceptionOccurrenceExtension                      *groupeventinstanceexceptionoccurrenceextension.GroupEventInstanceExceptionOccurrenceExtensionClient
	GroupEventInstanceExtension                                         *groupeventinstanceextension.GroupEventInstanceExtensionClient
	GroupExtension                                                      *groupextension.GroupExtensionClient
	GroupGroupLifecyclePolicy                                           *groupgrouplifecyclepolicy.GroupGroupLifecyclePolicyClient
	GroupMember                                                         *groupmember.GroupMemberClient
	GroupMemberOf                                                       *groupmemberof.GroupMemberOfClient
	GroupMembersWithLicenseError                                        *groupmemberswithlicenseerror.GroupMembersWithLicenseErrorClient
	GroupOnenote                                                        *grouponenote.GroupOnenoteClient
	GroupOnenoteNotebook                                                *grouponenotenotebook.GroupOnenoteNotebookClient
	GroupOnenoteNotebookSection                                         *grouponenotenotebooksection.GroupOnenoteNotebookSectionClient
	GroupOnenoteNotebookSectionGroup                                    *grouponenotenotebooksectiongroup.GroupOnenoteNotebookSectionGroupClient
	GroupOnenoteNotebookSectionGroupParentNotebook                      *grouponenotenotebooksectiongroupparentnotebook.GroupOnenoteNotebookSectionGroupParentNotebookClient
	GroupOnenoteNotebookSectionGroupParentSectionGroup                  *grouponenotenotebooksectiongroupparentsectiongroup.GroupOnenoteNotebookSectionGroupParentSectionGroupClient
	GroupOnenoteNotebookSectionGroupSection                             *grouponenotenotebooksectiongroupsection.GroupOnenoteNotebookSectionGroupSectionClient
	GroupOnenoteNotebookSectionGroupSectionGroup                        *grouponenotenotebooksectiongroupsectiongroup.GroupOnenoteNotebookSectionGroupSectionGroupClient
	GroupOnenoteNotebookSectionGroupSectionPage                         *grouponenotenotebooksectiongroupsectionpage.GroupOnenoteNotebookSectionGroupSectionPageClient
	GroupOnenoteNotebookSectionGroupSectionPageContent                  *grouponenotenotebooksectiongroupsectionpagecontent.GroupOnenoteNotebookSectionGroupSectionPageContentClient
	GroupOnenoteNotebookSectionGroupSectionPageParentNotebook           *grouponenotenotebooksectiongroupsectionpageparentnotebook.GroupOnenoteNotebookSectionGroupSectionPageParentNotebookClient
	GroupOnenoteNotebookSectionGroupSectionPageParentSection            *grouponenotenotebooksectiongroupsectionpageparentsection.GroupOnenoteNotebookSectionGroupSectionPageParentSectionClient
	GroupOnenoteNotebookSectionGroupSectionParentNotebook               *grouponenotenotebooksectiongroupsectionparentnotebook.GroupOnenoteNotebookSectionGroupSectionParentNotebookClient
	GroupOnenoteNotebookSectionGroupSectionParentSectionGroup           *grouponenotenotebooksectiongroupsectionparentsectiongroup.GroupOnenoteNotebookSectionGroupSectionParentSectionGroupClient
	GroupOnenoteNotebookSectionPage                                     *grouponenotenotebooksectionpage.GroupOnenoteNotebookSectionPageClient
	GroupOnenoteNotebookSectionPageContent                              *grouponenotenotebooksectionpagecontent.GroupOnenoteNotebookSectionPageContentClient
	GroupOnenoteNotebookSectionPageParentNotebook                       *grouponenotenotebooksectionpageparentnotebook.GroupOnenoteNotebookSectionPageParentNotebookClient
	GroupOnenoteNotebookSectionPageParentSection                        *grouponenotenotebooksectionpageparentsection.GroupOnenoteNotebookSectionPageParentSectionClient
	GroupOnenoteNotebookSectionParentNotebook                           *grouponenotenotebooksectionparentnotebook.GroupOnenoteNotebookSectionParentNotebookClient
	GroupOnenoteNotebookSectionParentSectionGroup                       *grouponenotenotebooksectionparentsectiongroup.GroupOnenoteNotebookSectionParentSectionGroupClient
	GroupOnenoteOperation                                               *grouponenoteoperation.GroupOnenoteOperationClient
	GroupOnenotePage                                                    *grouponenotepage.GroupOnenotePageClient
	GroupOnenotePageContent                                             *grouponenotepagecontent.GroupOnenotePageContentClient
	GroupOnenotePageParentNotebook                                      *grouponenotepageparentnotebook.GroupOnenotePageParentNotebookClient
	GroupOnenotePageParentSection                                       *grouponenotepageparentsection.GroupOnenotePageParentSectionClient
	GroupOnenoteResource                                                *grouponenoteresource.GroupOnenoteResourceClient
	GroupOnenoteResourceContent                                         *grouponenoteresourcecontent.GroupOnenoteResourceContentClient
	GroupOnenoteSection                                                 *grouponenotesection.GroupOnenoteSectionClient
	GroupOnenoteSectionGroup                                            *grouponenotesectiongroup.GroupOnenoteSectionGroupClient
	GroupOnenoteSectionGroupParentNotebook                              *grouponenotesectiongroupparentnotebook.GroupOnenoteSectionGroupParentNotebookClient
	GroupOnenoteSectionGroupParentSectionGroup                          *grouponenotesectiongroupparentsectiongroup.GroupOnenoteSectionGroupParentSectionGroupClient
	GroupOnenoteSectionGroupSection                                     *grouponenotesectiongroupsection.GroupOnenoteSectionGroupSectionClient
	GroupOnenoteSectionGroupSectionGroup                                *grouponenotesectiongroupsectiongroup.GroupOnenoteSectionGroupSectionGroupClient
	GroupOnenoteSectionGroupSectionPage                                 *grouponenotesectiongroupsectionpage.GroupOnenoteSectionGroupSectionPageClient
	GroupOnenoteSectionGroupSectionPageContent                          *grouponenotesectiongroupsectionpagecontent.GroupOnenoteSectionGroupSectionPageContentClient
	GroupOnenoteSectionGroupSectionPageParentNotebook                   *grouponenotesectiongroupsectionpageparentnotebook.GroupOnenoteSectionGroupSectionPageParentNotebookClient
	GroupOnenoteSectionGroupSectionPageParentSection                    *grouponenotesectiongroupsectionpageparentsection.GroupOnenoteSectionGroupSectionPageParentSectionClient
	GroupOnenoteSectionGroupSectionParentNotebook                       *grouponenotesectiongroupsectionparentnotebook.GroupOnenoteSectionGroupSectionParentNotebookClient
	GroupOnenoteSectionGroupSectionParentSectionGroup                   *grouponenotesectiongroupsectionparentsectiongroup.GroupOnenoteSectionGroupSectionParentSectionGroupClient
	GroupOnenoteSectionPage                                             *grouponenotesectionpage.GroupOnenoteSectionPageClient
	GroupOnenoteSectionPageContent                                      *grouponenotesectionpagecontent.GroupOnenoteSectionPageContentClient
	GroupOnenoteSectionPageParentNotebook                               *grouponenotesectionpageparentnotebook.GroupOnenoteSectionPageParentNotebookClient
	GroupOnenoteSectionPageParentSection                                *grouponenotesectionpageparentsection.GroupOnenoteSectionPageParentSectionClient
	GroupOnenoteSectionParentNotebook                                   *grouponenotesectionparentnotebook.GroupOnenoteSectionParentNotebookClient
	GroupOnenoteSectionParentSectionGroup                               *grouponenotesectionparentsectiongroup.GroupOnenoteSectionParentSectionGroupClient
	GroupOwner                                                          *groupowner.GroupOwnerClient
	GroupPermissionGrant                                                *grouppermissiongrant.GroupPermissionGrantClient
	GroupPhoto                                                          *groupphoto.GroupPhotoClient
	GroupPlanner                                                        *groupplanner.GroupPlannerClient
	GroupPlannerPlan                                                    *groupplannerplan.GroupPlannerPlanClient
	GroupPlannerPlanBucket                                              *groupplannerplanbucket.GroupPlannerPlanBucketClient
	GroupPlannerPlanBucketTask                                          *groupplannerplanbuckettask.GroupPlannerPlanBucketTaskClient
	GroupPlannerPlanBucketTaskAssignedToTaskBoardFormat                 *groupplannerplanbuckettaskassignedtotaskboardformat.GroupPlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	GroupPlannerPlanBucketTaskBucketTaskBoardFormat                     *groupplannerplanbuckettaskbuckettaskboardformat.GroupPlannerPlanBucketTaskBucketTaskBoardFormatClient
	GroupPlannerPlanBucketTaskDetail                                    *groupplannerplanbuckettaskdetail.GroupPlannerPlanBucketTaskDetailClient
	GroupPlannerPlanBucketTaskProgressTaskBoardFormat                   *groupplannerplanbuckettaskprogresstaskboardformat.GroupPlannerPlanBucketTaskProgressTaskBoardFormatClient
	GroupPlannerPlanDetail                                              *groupplannerplandetail.GroupPlannerPlanDetailClient
	GroupPlannerPlanTask                                                *groupplannerplantask.GroupPlannerPlanTaskClient
	GroupPlannerPlanTaskAssignedToTaskBoardFormat                       *groupplannerplantaskassignedtotaskboardformat.GroupPlannerPlanTaskAssignedToTaskBoardFormatClient
	GroupPlannerPlanTaskBucketTaskBoardFormat                           *groupplannerplantaskbuckettaskboardformat.GroupPlannerPlanTaskBucketTaskBoardFormatClient
	GroupPlannerPlanTaskDetail                                          *groupplannerplantaskdetail.GroupPlannerPlanTaskDetailClient
	GroupPlannerPlanTaskProgressTaskBoardFormat                         *groupplannerplantaskprogresstaskboardformat.GroupPlannerPlanTaskProgressTaskBoardFormatClient
	GroupRejectedSender                                                 *grouprejectedsender.GroupRejectedSenderClient
	GroupSetting                                                        *groupsetting.GroupSettingClient
	GroupSite                                                           *groupsite.GroupSiteClient
	GroupSiteAnalytic                                                   *groupsiteanalytic.GroupSiteAnalyticClient
	GroupSiteAnalyticAllTime                                            *groupsiteanalyticalltime.GroupSiteAnalyticAllTimeClient
	GroupSiteAnalyticItemActivityStat                                   *groupsiteanalyticitemactivitystat.GroupSiteAnalyticItemActivityStatClient
	GroupSiteAnalyticItemActivityStatActivity                           *groupsiteanalyticitemactivitystatactivity.GroupSiteAnalyticItemActivityStatActivityClient
	GroupSiteAnalyticItemActivityStatActivityDriveItem                  *groupsiteanalyticitemactivitystatactivitydriveitem.GroupSiteAnalyticItemActivityStatActivityDriveItemClient
	GroupSiteAnalyticItemActivityStatActivityDriveItemContent           *groupsiteanalyticitemactivitystatactivitydriveitemcontent.GroupSiteAnalyticItemActivityStatActivityDriveItemContentClient
	GroupSiteAnalyticLastSevenDay                                       *groupsiteanalyticlastsevenday.GroupSiteAnalyticLastSevenDayClient
	GroupSiteColumn                                                     *groupsitecolumn.GroupSiteColumnClient
	GroupSiteColumnSourceColumn                                         *groupsitecolumnsourcecolumn.GroupSiteColumnSourceColumnClient
	GroupSiteContentType                                                *groupsitecontenttype.GroupSiteContentTypeClient
	GroupSiteContentTypeBase                                            *groupsitecontenttypebase.GroupSiteContentTypeBaseClient
	GroupSiteContentTypeBaseType                                        *groupsitecontenttypebasetype.GroupSiteContentTypeBaseTypeClient
	GroupSiteContentTypeColumn                                          *groupsitecontenttypecolumn.GroupSiteContentTypeColumnClient
	GroupSiteContentTypeColumnLink                                      *groupsitecontenttypecolumnlink.GroupSiteContentTypeColumnLinkClient
	GroupSiteContentTypeColumnPosition                                  *groupsitecontenttypecolumnposition.GroupSiteContentTypeColumnPositionClient
	GroupSiteContentTypeColumnSourceColumn                              *groupsitecontenttypecolumnsourcecolumn.GroupSiteContentTypeColumnSourceColumnClient
	GroupSiteCreatedByUser                                              *groupsitecreatedbyuser.GroupSiteCreatedByUserClient
	GroupSiteCreatedByUserMailboxSetting                                *groupsitecreatedbyusermailboxsetting.GroupSiteCreatedByUserMailboxSettingClient
	GroupSiteDrive                                                      *groupsitedrive.GroupSiteDriveClient
	GroupSiteExternalColumn                                             *groupsiteexternalcolumn.GroupSiteExternalColumnClient
	GroupSiteInformationProtection                                      *groupsiteinformationprotection.GroupSiteInformationProtectionClient
	GroupSiteInformationProtectionBitlocker                             *groupsiteinformationprotectionbitlocker.GroupSiteInformationProtectionBitlockerClient
	GroupSiteInformationProtectionBitlockerRecoveryKey                  *groupsiteinformationprotectionbitlockerrecoverykey.GroupSiteInformationProtectionBitlockerRecoveryKeyClient
	GroupSiteInformationProtectionDataLossPreventionPolicy              *groupsiteinformationprotectiondatalosspreventionpolicy.GroupSiteInformationProtectionDataLossPreventionPolicyClient
	GroupSiteInformationProtectionPolicy                                *groupsiteinformationprotectionpolicy.GroupSiteInformationProtectionPolicyClient
	GroupSiteInformationProtectionPolicyLabel                           *groupsiteinformationprotectionpolicylabel.GroupSiteInformationProtectionPolicyLabelClient
	GroupSiteInformationProtectionSensitivityLabel                      *groupsiteinformationprotectionsensitivitylabel.GroupSiteInformationProtectionSensitivityLabelClient
	GroupSiteInformationProtectionSensitivityLabelSublabel              *groupsiteinformationprotectionsensitivitylabelsublabel.GroupSiteInformationProtectionSensitivityLabelSublabelClient
	GroupSiteInformationProtectionSensitivityPolicySetting              *groupsiteinformationprotectionsensitivitypolicysetting.GroupSiteInformationProtectionSensitivityPolicySettingClient
	GroupSiteInformationProtectionThreatAssessmentRequest               *groupsiteinformationprotectionthreatassessmentrequest.GroupSiteInformationProtectionThreatAssessmentRequestClient
	GroupSiteInformationProtectionThreatAssessmentRequestResult         *groupsiteinformationprotectionthreatassessmentrequestresult.GroupSiteInformationProtectionThreatAssessmentRequestResultClient
	GroupSiteItem                                                       *groupsiteitem.GroupSiteItemClient
	GroupSiteLastModifiedByUser                                         *groupsitelastmodifiedbyuser.GroupSiteLastModifiedByUserClient
	GroupSiteLastModifiedByUserMailboxSetting                           *groupsitelastmodifiedbyusermailboxsetting.GroupSiteLastModifiedByUserMailboxSettingClient
	GroupSiteList                                                       *groupsitelist.GroupSiteListClient
	GroupSiteListActivity                                               *groupsitelistactivity.GroupSiteListActivityClient
	GroupSiteListColumn                                                 *groupsitelistcolumn.GroupSiteListColumnClient
	GroupSiteListColumnSourceColumn                                     *groupsitelistcolumnsourcecolumn.GroupSiteListColumnSourceColumnClient
	GroupSiteListContentType                                            *groupsitelistcontenttype.GroupSiteListContentTypeClient
	GroupSiteListContentTypeBase                                        *groupsitelistcontenttypebase.GroupSiteListContentTypeBaseClient
	GroupSiteListContentTypeBaseType                                    *groupsitelistcontenttypebasetype.GroupSiteListContentTypeBaseTypeClient
	GroupSiteListContentTypeColumn                                      *groupsitelistcontenttypecolumn.GroupSiteListContentTypeColumnClient
	GroupSiteListContentTypeColumnLink                                  *groupsitelistcontenttypecolumnlink.GroupSiteListContentTypeColumnLinkClient
	GroupSiteListContentTypeColumnPosition                              *groupsitelistcontenttypecolumnposition.GroupSiteListContentTypeColumnPositionClient
	GroupSiteListContentTypeColumnSourceColumn                          *groupsitelistcontenttypecolumnsourcecolumn.GroupSiteListContentTypeColumnSourceColumnClient
	GroupSiteListCreatedByUser                                          *groupsitelistcreatedbyuser.GroupSiteListCreatedByUserClient
	GroupSiteListCreatedByUserMailboxSetting                            *groupsitelistcreatedbyusermailboxsetting.GroupSiteListCreatedByUserMailboxSettingClient
	GroupSiteListDrive                                                  *groupsitelistdrive.GroupSiteListDriveClient
	GroupSiteListItem                                                   *groupsitelistitem.GroupSiteListItemClient
	GroupSiteListItemActivity                                           *groupsitelistitemactivity.GroupSiteListItemActivityClient
	GroupSiteListItemActivityDriveItem                                  *groupsitelistitemactivitydriveitem.GroupSiteListItemActivityDriveItemClient
	GroupSiteListItemActivityDriveItemContent                           *groupsitelistitemactivitydriveitemcontent.GroupSiteListItemActivityDriveItemContentClient
	GroupSiteListItemActivityListItem                                   *groupsitelistitemactivitylistitem.GroupSiteListItemActivityListItemClient
	GroupSiteListItemAnalytic                                           *groupsitelistitemanalytic.GroupSiteListItemAnalyticClient
	GroupSiteListItemCreatedByUser                                      *groupsitelistitemcreatedbyuser.GroupSiteListItemCreatedByUserClient
	GroupSiteListItemCreatedByUserMailboxSetting                        *groupsitelistitemcreatedbyusermailboxsetting.GroupSiteListItemCreatedByUserMailboxSettingClient
	GroupSiteListItemDocumentSetVersion                                 *groupsitelistitemdocumentsetversion.GroupSiteListItemDocumentSetVersionClient
	GroupSiteListItemDocumentSetVersionField                            *groupsitelistitemdocumentsetversionfield.GroupSiteListItemDocumentSetVersionFieldClient
	GroupSiteListItemDriveItem                                          *groupsitelistitemdriveitem.GroupSiteListItemDriveItemClient
	GroupSiteListItemDriveItemContent                                   *groupsitelistitemdriveitemcontent.GroupSiteListItemDriveItemContentClient
	GroupSiteListItemField                                              *groupsitelistitemfield.GroupSiteListItemFieldClient
	GroupSiteListItemLastModifiedByUser                                 *groupsitelistitemlastmodifiedbyuser.GroupSiteListItemLastModifiedByUserClient
	GroupSiteListItemLastModifiedByUserMailboxSetting                   *groupsitelistitemlastmodifiedbyusermailboxsetting.GroupSiteListItemLastModifiedByUserMailboxSettingClient
	GroupSiteListItemVersion                                            *groupsitelistitemversion.GroupSiteListItemVersionClient
	GroupSiteListItemVersionField                                       *groupsitelistitemversionfield.GroupSiteListItemVersionFieldClient
	GroupSiteListLastModifiedByUser                                     *groupsitelistlastmodifiedbyuser.GroupSiteListLastModifiedByUserClient
	GroupSiteListLastModifiedByUserMailboxSetting                       *groupsitelistlastmodifiedbyusermailboxsetting.GroupSiteListLastModifiedByUserMailboxSettingClient
	GroupSiteListOperation                                              *groupsitelistoperation.GroupSiteListOperationClient
	GroupSiteListSubscription                                           *groupsitelistsubscription.GroupSiteListSubscriptionClient
	GroupSiteOnenote                                                    *groupsiteonenote.GroupSiteOnenoteClient
	GroupSiteOnenoteNotebook                                            *groupsiteonenotenotebook.GroupSiteOnenoteNotebookClient
	GroupSiteOnenoteNotebookSection                                     *groupsiteonenotenotebooksection.GroupSiteOnenoteNotebookSectionClient
	GroupSiteOnenoteNotebookSectionGroup                                *groupsiteonenotenotebooksectiongroup.GroupSiteOnenoteNotebookSectionGroupClient
	GroupSiteOnenoteNotebookSectionGroupParentNotebook                  *groupsiteonenotenotebooksectiongroupparentnotebook.GroupSiteOnenoteNotebookSectionGroupParentNotebookClient
	GroupSiteOnenoteNotebookSectionGroupParentSectionGroup              *groupsiteonenotenotebooksectiongroupparentsectiongroup.GroupSiteOnenoteNotebookSectionGroupParentSectionGroupClient
	GroupSiteOnenoteNotebookSectionGroupSection                         *groupsiteonenotenotebooksectiongroupsection.GroupSiteOnenoteNotebookSectionGroupSectionClient
	GroupSiteOnenoteNotebookSectionGroupSectionGroup                    *groupsiteonenotenotebooksectiongroupsectiongroup.GroupSiteOnenoteNotebookSectionGroupSectionGroupClient
	GroupSiteOnenoteNotebookSectionGroupSectionPage                     *groupsiteonenotenotebooksectiongroupsectionpage.GroupSiteOnenoteNotebookSectionGroupSectionPageClient
	GroupSiteOnenoteNotebookSectionGroupSectionPageContent              *groupsiteonenotenotebooksectiongroupsectionpagecontent.GroupSiteOnenoteNotebookSectionGroupSectionPageContentClient
	GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebook       *groupsiteonenotenotebooksectiongroupsectionpageparentnotebook.GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient
	GroupSiteOnenoteNotebookSectionGroupSectionPageParentSection        *groupsiteonenotenotebooksectiongroupsectionpageparentsection.GroupSiteOnenoteNotebookSectionGroupSectionPageParentSectionClient
	GroupSiteOnenoteNotebookSectionGroupSectionParentNotebook           *groupsiteonenotenotebooksectiongroupsectionparentnotebook.GroupSiteOnenoteNotebookSectionGroupSectionParentNotebookClient
	GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroup       *groupsiteonenotenotebooksectiongroupsectionparentsectiongroup.GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient
	GroupSiteOnenoteNotebookSectionPage                                 *groupsiteonenotenotebooksectionpage.GroupSiteOnenoteNotebookSectionPageClient
	GroupSiteOnenoteNotebookSectionPageContent                          *groupsiteonenotenotebooksectionpagecontent.GroupSiteOnenoteNotebookSectionPageContentClient
	GroupSiteOnenoteNotebookSectionPageParentNotebook                   *groupsiteonenotenotebooksectionpageparentnotebook.GroupSiteOnenoteNotebookSectionPageParentNotebookClient
	GroupSiteOnenoteNotebookSectionPageParentSection                    *groupsiteonenotenotebooksectionpageparentsection.GroupSiteOnenoteNotebookSectionPageParentSectionClient
	GroupSiteOnenoteNotebookSectionParentNotebook                       *groupsiteonenotenotebooksectionparentnotebook.GroupSiteOnenoteNotebookSectionParentNotebookClient
	GroupSiteOnenoteNotebookSectionParentSectionGroup                   *groupsiteonenotenotebooksectionparentsectiongroup.GroupSiteOnenoteNotebookSectionParentSectionGroupClient
	GroupSiteOnenoteOperation                                           *groupsiteonenoteoperation.GroupSiteOnenoteOperationClient
	GroupSiteOnenotePage                                                *groupsiteonenotepage.GroupSiteOnenotePageClient
	GroupSiteOnenotePageContent                                         *groupsiteonenotepagecontent.GroupSiteOnenotePageContentClient
	GroupSiteOnenotePageParentNotebook                                  *groupsiteonenotepageparentnotebook.GroupSiteOnenotePageParentNotebookClient
	GroupSiteOnenotePageParentSection                                   *groupsiteonenotepageparentsection.GroupSiteOnenotePageParentSectionClient
	GroupSiteOnenoteResource                                            *groupsiteonenoteresource.GroupSiteOnenoteResourceClient
	GroupSiteOnenoteResourceContent                                     *groupsiteonenoteresourcecontent.GroupSiteOnenoteResourceContentClient
	GroupSiteOnenoteSection                                             *groupsiteonenotesection.GroupSiteOnenoteSectionClient
	GroupSiteOnenoteSectionGroup                                        *groupsiteonenotesectiongroup.GroupSiteOnenoteSectionGroupClient
	GroupSiteOnenoteSectionGroupParentNotebook                          *groupsiteonenotesectiongroupparentnotebook.GroupSiteOnenoteSectionGroupParentNotebookClient
	GroupSiteOnenoteSectionGroupParentSectionGroup                      *groupsiteonenotesectiongroupparentsectiongroup.GroupSiteOnenoteSectionGroupParentSectionGroupClient
	GroupSiteOnenoteSectionGroupSection                                 *groupsiteonenotesectiongroupsection.GroupSiteOnenoteSectionGroupSectionClient
	GroupSiteOnenoteSectionGroupSectionGroup                            *groupsiteonenotesectiongroupsectiongroup.GroupSiteOnenoteSectionGroupSectionGroupClient
	GroupSiteOnenoteSectionGroupSectionPage                             *groupsiteonenotesectiongroupsectionpage.GroupSiteOnenoteSectionGroupSectionPageClient
	GroupSiteOnenoteSectionGroupSectionPageContent                      *groupsiteonenotesectiongroupsectionpagecontent.GroupSiteOnenoteSectionGroupSectionPageContentClient
	GroupSiteOnenoteSectionGroupSectionPageParentNotebook               *groupsiteonenotesectiongroupsectionpageparentnotebook.GroupSiteOnenoteSectionGroupSectionPageParentNotebookClient
	GroupSiteOnenoteSectionGroupSectionPageParentSection                *groupsiteonenotesectiongroupsectionpageparentsection.GroupSiteOnenoteSectionGroupSectionPageParentSectionClient
	GroupSiteOnenoteSectionGroupSectionParentNotebook                   *groupsiteonenotesectiongroupsectionparentnotebook.GroupSiteOnenoteSectionGroupSectionParentNotebookClient
	GroupSiteOnenoteSectionGroupSectionParentSectionGroup               *groupsiteonenotesectiongroupsectionparentsectiongroup.GroupSiteOnenoteSectionGroupSectionParentSectionGroupClient
	GroupSiteOnenoteSectionPage                                         *groupsiteonenotesectionpage.GroupSiteOnenoteSectionPageClient
	GroupSiteOnenoteSectionPageContent                                  *groupsiteonenotesectionpagecontent.GroupSiteOnenoteSectionPageContentClient
	GroupSiteOnenoteSectionPageParentNotebook                           *groupsiteonenotesectionpageparentnotebook.GroupSiteOnenoteSectionPageParentNotebookClient
	GroupSiteOnenoteSectionPageParentSection                            *groupsiteonenotesectionpageparentsection.GroupSiteOnenoteSectionPageParentSectionClient
	GroupSiteOnenoteSectionParentNotebook                               *groupsiteonenotesectionparentnotebook.GroupSiteOnenoteSectionParentNotebookClient
	GroupSiteOnenoteSectionParentSectionGroup                           *groupsiteonenotesectionparentsectiongroup.GroupSiteOnenoteSectionParentSectionGroupClient
	GroupSiteOperation                                                  *groupsiteoperation.GroupSiteOperationClient
	GroupSitePage                                                       *groupsitepage.GroupSitePageClient
	GroupSitePageCreatedByUser                                          *groupsitepagecreatedbyuser.GroupSitePageCreatedByUserClient
	GroupSitePageCreatedByUserMailboxSetting                            *groupsitepagecreatedbyusermailboxsetting.GroupSitePageCreatedByUserMailboxSettingClient
	GroupSitePageLastModifiedByUser                                     *groupsitepagelastmodifiedbyuser.GroupSitePageLastModifiedByUserClient
	GroupSitePageLastModifiedByUserMailboxSetting                       *groupsitepagelastmodifiedbyusermailboxsetting.GroupSitePageLastModifiedByUserMailboxSettingClient
	GroupSitePermission                                                 *groupsitepermission.GroupSitePermissionClient
	GroupSiteSite                                                       *groupsitesite.GroupSiteSiteClient
	GroupSiteTermStore                                                  *groupsitetermstore.GroupSiteTermStoreClient
	GroupSiteTermStoreGroup                                             *groupsitetermstoregroup.GroupSiteTermStoreGroupClient
	GroupSiteTermStoreGroupSet                                          *groupsitetermstoregroupset.GroupSiteTermStoreGroupSetClient
	GroupSiteTermStoreGroupSetChildren                                  *groupsitetermstoregroupsetchildren.GroupSiteTermStoreGroupSetChildrenClient
	GroupSiteTermStoreGroupSetChildrenChildren                          *groupsitetermstoregroupsetchildrenchildren.GroupSiteTermStoreGroupSetChildrenChildrenClient
	GroupSiteTermStoreGroupSetChildrenChildrenRelation                  *groupsitetermstoregroupsetchildrenchildrenrelation.GroupSiteTermStoreGroupSetChildrenChildrenRelationClient
	GroupSiteTermStoreGroupSetChildrenChildrenRelationFromTerm          *groupsitetermstoregroupsetchildrenchildrenrelationfromterm.GroupSiteTermStoreGroupSetChildrenChildrenRelationFromTermClient
	GroupSiteTermStoreGroupSetChildrenChildrenRelationToTerm            *groupsitetermstoregroupsetchildrenchildrenrelationtoterm.GroupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient
	GroupSiteTermStoreGroupSetChildrenRelation                          *groupsitetermstoregroupsetchildrenrelation.GroupSiteTermStoreGroupSetChildrenRelationClient
	GroupSiteTermStoreGroupSetChildrenRelationFromTerm                  *groupsitetermstoregroupsetchildrenrelationfromterm.GroupSiteTermStoreGroupSetChildrenRelationFromTermClient
	GroupSiteTermStoreGroupSetChildrenRelationToTerm                    *groupsitetermstoregroupsetchildrenrelationtoterm.GroupSiteTermStoreGroupSetChildrenRelationToTermClient
	GroupSiteTermStoreGroupSetParentGroup                               *groupsitetermstoregroupsetparentgroup.GroupSiteTermStoreGroupSetParentGroupClient
	GroupSiteTermStoreGroupSetRelation                                  *groupsitetermstoregroupsetrelation.GroupSiteTermStoreGroupSetRelationClient
	GroupSiteTermStoreGroupSetRelationFromTerm                          *groupsitetermstoregroupsetrelationfromterm.GroupSiteTermStoreGroupSetRelationFromTermClient
	GroupSiteTermStoreGroupSetRelationToTerm                            *groupsitetermstoregroupsetrelationtoterm.GroupSiteTermStoreGroupSetRelationToTermClient
	GroupSiteTermStoreGroupSetTerm                                      *groupsitetermstoregroupsetterm.GroupSiteTermStoreGroupSetTermClient
	GroupSiteTermStoreGroupSetTermChildren                              *groupsitetermstoregroupsettermchildren.GroupSiteTermStoreGroupSetTermChildrenClient
	GroupSiteTermStoreGroupSetTermChildrenRelation                      *groupsitetermstoregroupsettermchildrenrelation.GroupSiteTermStoreGroupSetTermChildrenRelationClient
	GroupSiteTermStoreGroupSetTermChildrenRelationFromTerm              *groupsitetermstoregroupsettermchildrenrelationfromterm.GroupSiteTermStoreGroupSetTermChildrenRelationFromTermClient
	GroupSiteTermStoreGroupSetTermChildrenRelationToTerm                *groupsitetermstoregroupsettermchildrenrelationtoterm.GroupSiteTermStoreGroupSetTermChildrenRelationToTermClient
	GroupSiteTermStoreGroupSetTermRelation                              *groupsitetermstoregroupsettermrelation.GroupSiteTermStoreGroupSetTermRelationClient
	GroupSiteTermStoreGroupSetTermRelationFromTerm                      *groupsitetermstoregroupsettermrelationfromterm.GroupSiteTermStoreGroupSetTermRelationFromTermClient
	GroupSiteTermStoreGroupSetTermRelationToTerm                        *groupsitetermstoregroupsettermrelationtoterm.GroupSiteTermStoreGroupSetTermRelationToTermClient
	GroupSiteTermStoreSet                                               *groupsitetermstoreset.GroupSiteTermStoreSetClient
	GroupSiteTermStoreSetChildren                                       *groupsitetermstoresetchildren.GroupSiteTermStoreSetChildrenClient
	GroupSiteTermStoreSetChildrenChildren                               *groupsitetermstoresetchildrenchildren.GroupSiteTermStoreSetChildrenChildrenClient
	GroupSiteTermStoreSetChildrenChildrenRelation                       *groupsitetermstoresetchildrenchildrenrelation.GroupSiteTermStoreSetChildrenChildrenRelationClient
	GroupSiteTermStoreSetChildrenChildrenRelationFromTerm               *groupsitetermstoresetchildrenchildrenrelationfromterm.GroupSiteTermStoreSetChildrenChildrenRelationFromTermClient
	GroupSiteTermStoreSetChildrenChildrenRelationToTerm                 *groupsitetermstoresetchildrenchildrenrelationtoterm.GroupSiteTermStoreSetChildrenChildrenRelationToTermClient
	GroupSiteTermStoreSetChildrenRelation                               *groupsitetermstoresetchildrenrelation.GroupSiteTermStoreSetChildrenRelationClient
	GroupSiteTermStoreSetChildrenRelationFromTerm                       *groupsitetermstoresetchildrenrelationfromterm.GroupSiteTermStoreSetChildrenRelationFromTermClient
	GroupSiteTermStoreSetChildrenRelationToTerm                         *groupsitetermstoresetchildrenrelationtoterm.GroupSiteTermStoreSetChildrenRelationToTermClient
	GroupSiteTermStoreSetParentGroup                                    *groupsitetermstoresetparentgroup.GroupSiteTermStoreSetParentGroupClient
	GroupSiteTermStoreSetParentGroupSet                                 *groupsitetermstoresetparentgroupset.GroupSiteTermStoreSetParentGroupSetClient
	GroupSiteTermStoreSetParentGroupSetChildren                         *groupsitetermstoresetparentgroupsetchildren.GroupSiteTermStoreSetParentGroupSetChildrenClient
	GroupSiteTermStoreSetParentGroupSetChildrenChildren                 *groupsitetermstoresetparentgroupsetchildrenchildren.GroupSiteTermStoreSetParentGroupSetChildrenChildrenClient
	GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelation         *groupsitetermstoresetparentgroupsetchildrenchildrenrelation.GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient
	GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTerm *groupsitetermstoresetparentgroupsetchildrenchildrenrelationfromterm.GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTermClient
	GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTerm   *groupsitetermstoresetparentgroupsetchildrenchildrenrelationtoterm.GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTermClient
	GroupSiteTermStoreSetParentGroupSetChildrenRelation                 *groupsitetermstoresetparentgroupsetchildrenrelation.GroupSiteTermStoreSetParentGroupSetChildrenRelationClient
	GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTerm         *groupsitetermstoresetparentgroupsetchildrenrelationfromterm.GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient
	GroupSiteTermStoreSetParentGroupSetChildrenRelationToTerm           *groupsitetermstoresetparentgroupsetchildrenrelationtoterm.GroupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient
	GroupSiteTermStoreSetParentGroupSetRelation                         *groupsitetermstoresetparentgroupsetrelation.GroupSiteTermStoreSetParentGroupSetRelationClient
	GroupSiteTermStoreSetParentGroupSetRelationFromTerm                 *groupsitetermstoresetparentgroupsetrelationfromterm.GroupSiteTermStoreSetParentGroupSetRelationFromTermClient
	GroupSiteTermStoreSetParentGroupSetRelationToTerm                   *groupsitetermstoresetparentgroupsetrelationtoterm.GroupSiteTermStoreSetParentGroupSetRelationToTermClient
	GroupSiteTermStoreSetParentGroupSetTerm                             *groupsitetermstoresetparentgroupsetterm.GroupSiteTermStoreSetParentGroupSetTermClient
	GroupSiteTermStoreSetParentGroupSetTermChildren                     *groupsitetermstoresetparentgroupsettermchildren.GroupSiteTermStoreSetParentGroupSetTermChildrenClient
	GroupSiteTermStoreSetParentGroupSetTermChildrenRelation             *groupsitetermstoresetparentgroupsettermchildrenrelation.GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient
	GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTerm     *groupsitetermstoresetparentgroupsettermchildrenrelationfromterm.GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient
	GroupSiteTermStoreSetParentGroupSetTermChildrenRelationToTerm       *groupsitetermstoresetparentgroupsettermchildrenrelationtoterm.GroupSiteTermStoreSetParentGroupSetTermChildrenRelationToTermClient
	GroupSiteTermStoreSetParentGroupSetTermRelation                     *groupsitetermstoresetparentgroupsettermrelation.GroupSiteTermStoreSetParentGroupSetTermRelationClient
	GroupSiteTermStoreSetParentGroupSetTermRelationFromTerm             *groupsitetermstoresetparentgroupsettermrelationfromterm.GroupSiteTermStoreSetParentGroupSetTermRelationFromTermClient
	GroupSiteTermStoreSetParentGroupSetTermRelationToTerm               *groupsitetermstoresetparentgroupsettermrelationtoterm.GroupSiteTermStoreSetParentGroupSetTermRelationToTermClient
	GroupSiteTermStoreSetRelation                                       *groupsitetermstoresetrelation.GroupSiteTermStoreSetRelationClient
	GroupSiteTermStoreSetRelationFromTerm                               *groupsitetermstoresetrelationfromterm.GroupSiteTermStoreSetRelationFromTermClient
	GroupSiteTermStoreSetRelationToTerm                                 *groupsitetermstoresetrelationtoterm.GroupSiteTermStoreSetRelationToTermClient
	GroupSiteTermStoreSetTerm                                           *groupsitetermstoresetterm.GroupSiteTermStoreSetTermClient
	GroupSiteTermStoreSetTermChildren                                   *groupsitetermstoresettermchildren.GroupSiteTermStoreSetTermChildrenClient
	GroupSiteTermStoreSetTermChildrenRelation                           *groupsitetermstoresettermchildrenrelation.GroupSiteTermStoreSetTermChildrenRelationClient
	GroupSiteTermStoreSetTermChildrenRelationFromTerm                   *groupsitetermstoresettermchildrenrelationfromterm.GroupSiteTermStoreSetTermChildrenRelationFromTermClient
	GroupSiteTermStoreSetTermChildrenRelationToTerm                     *groupsitetermstoresettermchildrenrelationtoterm.GroupSiteTermStoreSetTermChildrenRelationToTermClient
	GroupSiteTermStoreSetTermRelation                                   *groupsitetermstoresettermrelation.GroupSiteTermStoreSetTermRelationClient
	GroupSiteTermStoreSetTermRelationFromTerm                           *groupsitetermstoresettermrelationfromterm.GroupSiteTermStoreSetTermRelationFromTermClient
	GroupSiteTermStoreSetTermRelationToTerm                             *groupsitetermstoresettermrelationtoterm.GroupSiteTermStoreSetTermRelationToTermClient
	GroupTeam                                                           *groupteam.GroupTeamClient
	GroupTeamAllChannel                                                 *groupteamallchannel.GroupTeamAllChannelClient
	GroupTeamChannel                                                    *groupteamchannel.GroupTeamChannelClient
	GroupTeamChannelFilesFolder                                         *groupteamchannelfilesfolder.GroupTeamChannelFilesFolderClient
	GroupTeamChannelFilesFolderContent                                  *groupteamchannelfilesfoldercontent.GroupTeamChannelFilesFolderContentClient
	GroupTeamChannelMember                                              *groupteamchannelmember.GroupTeamChannelMemberClient
	GroupTeamChannelMessage                                             *groupteamchannelmessage.GroupTeamChannelMessageClient
	GroupTeamChannelMessageHostedContent                                *groupteamchannelmessagehostedcontent.GroupTeamChannelMessageHostedContentClient
	GroupTeamChannelMessageReply                                        *groupteamchannelmessagereply.GroupTeamChannelMessageReplyClient
	GroupTeamChannelMessageReplyHostedContent                           *groupteamchannelmessagereplyhostedcontent.GroupTeamChannelMessageReplyHostedContentClient
	GroupTeamChannelSharedWithTeam                                      *groupteamchannelsharedwithteam.GroupTeamChannelSharedWithTeamClient
	GroupTeamChannelSharedWithTeamAllowedMember                         *groupteamchannelsharedwithteamallowedmember.GroupTeamChannelSharedWithTeamAllowedMemberClient
	GroupTeamChannelSharedWithTeamTeam                                  *groupteamchannelsharedwithteamteam.GroupTeamChannelSharedWithTeamTeamClient
	GroupTeamChannelTab                                                 *groupteamchanneltab.GroupTeamChannelTabClient
	GroupTeamChannelTabTeamsApp                                         *groupteamchanneltabteamsapp.GroupTeamChannelTabTeamsAppClient
	GroupTeamGroup                                                      *groupteamgroup.GroupTeamGroupClient
	GroupTeamIncomingChannel                                            *groupteamincomingchannel.GroupTeamIncomingChannelClient
	GroupTeamInstalledApp                                               *groupteaminstalledapp.GroupTeamInstalledAppClient
	GroupTeamInstalledAppTeamsApp                                       *groupteaminstalledappteamsapp.GroupTeamInstalledAppTeamsAppClient
	GroupTeamInstalledAppTeamsAppDefinition                             *groupteaminstalledappteamsappdefinition.GroupTeamInstalledAppTeamsAppDefinitionClient
	GroupTeamMember                                                     *groupteammember.GroupTeamMemberClient
	GroupTeamOperation                                                  *groupteamoperation.GroupTeamOperationClient
	GroupTeamOwner                                                      *groupteamowner.GroupTeamOwnerClient
	GroupTeamOwnerMailboxSetting                                        *groupteamownermailboxsetting.GroupTeamOwnerMailboxSettingClient
	GroupTeamPermissionGrant                                            *groupteampermissiongrant.GroupTeamPermissionGrantClient
	GroupTeamPhoto                                                      *groupteamphoto.GroupTeamPhotoClient
	GroupTeamPrimaryChannel                                             *groupteamprimarychannel.GroupTeamPrimaryChannelClient
	GroupTeamPrimaryChannelFilesFolder                                  *groupteamprimarychannelfilesfolder.GroupTeamPrimaryChannelFilesFolderClient
	GroupTeamPrimaryChannelFilesFolderContent                           *groupteamprimarychannelfilesfoldercontent.GroupTeamPrimaryChannelFilesFolderContentClient
	GroupTeamPrimaryChannelMember                                       *groupteamprimarychannelmember.GroupTeamPrimaryChannelMemberClient
	GroupTeamPrimaryChannelMessage                                      *groupteamprimarychannelmessage.GroupTeamPrimaryChannelMessageClient
	GroupTeamPrimaryChannelMessageHostedContent                         *groupteamprimarychannelmessagehostedcontent.GroupTeamPrimaryChannelMessageHostedContentClient
	GroupTeamPrimaryChannelMessageReply                                 *groupteamprimarychannelmessagereply.GroupTeamPrimaryChannelMessageReplyClient
	GroupTeamPrimaryChannelMessageReplyHostedContent                    *groupteamprimarychannelmessagereplyhostedcontent.GroupTeamPrimaryChannelMessageReplyHostedContentClient
	GroupTeamPrimaryChannelSharedWithTeam                               *groupteamprimarychannelsharedwithteam.GroupTeamPrimaryChannelSharedWithTeamClient
	GroupTeamPrimaryChannelSharedWithTeamAllowedMember                  *groupteamprimarychannelsharedwithteamallowedmember.GroupTeamPrimaryChannelSharedWithTeamAllowedMemberClient
	GroupTeamPrimaryChannelSharedWithTeamTeam                           *groupteamprimarychannelsharedwithteamteam.GroupTeamPrimaryChannelSharedWithTeamTeamClient
	GroupTeamPrimaryChannelTab                                          *groupteamprimarychanneltab.GroupTeamPrimaryChannelTabClient
	GroupTeamPrimaryChannelTabTeamsApp                                  *groupteamprimarychanneltabteamsapp.GroupTeamPrimaryChannelTabTeamsAppClient
	GroupTeamSchedule                                                   *groupteamschedule.GroupTeamScheduleClient
	GroupTeamScheduleOfferShiftRequest                                  *groupteamscheduleoffershiftrequest.GroupTeamScheduleOfferShiftRequestClient
	GroupTeamScheduleOpenShift                                          *groupteamscheduleopenshift.GroupTeamScheduleOpenShiftClient
	GroupTeamScheduleOpenShiftChangeRequest                             *groupteamscheduleopenshiftchangerequest.GroupTeamScheduleOpenShiftChangeRequestClient
	GroupTeamScheduleSchedulingGroup                                    *groupteamscheduleschedulinggroup.GroupTeamScheduleSchedulingGroupClient
	GroupTeamScheduleShift                                              *groupteamscheduleshift.GroupTeamScheduleShiftClient
	GroupTeamScheduleSwapShiftsChangeRequest                            *groupteamscheduleswapshiftschangerequest.GroupTeamScheduleSwapShiftsChangeRequestClient
	GroupTeamScheduleTimeCard                                           *groupteamscheduletimecard.GroupTeamScheduleTimeCardClient
	GroupTeamScheduleTimeOffReason                                      *groupteamscheduletimeoffreason.GroupTeamScheduleTimeOffReasonClient
	GroupTeamScheduleTimeOffRequest                                     *groupteamscheduletimeoffrequest.GroupTeamScheduleTimeOffRequestClient
	GroupTeamScheduleTimesOff                                           *groupteamscheduletimesoff.GroupTeamScheduleTimesOffClient
	GroupTeamTag                                                        *groupteamtag.GroupTeamTagClient
	GroupTeamTagMember                                                  *groupteamtagmember.GroupTeamTagMemberClient
	GroupTeamTemplate                                                   *groupteamtemplate.GroupTeamTemplateClient
	GroupTeamTemplateDefinition                                         *groupteamtemplatedefinition.GroupTeamTemplateDefinitionClient
	GroupThread                                                         *groupthread.GroupThreadClient
	GroupThreadPost                                                     *groupthreadpost.GroupThreadPostClient
	GroupThreadPostAttachment                                           *groupthreadpostattachment.GroupThreadPostAttachmentClient
	GroupThreadPostExtension                                            *groupthreadpostextension.GroupThreadPostExtensionClient
	GroupThreadPostInReplyTo                                            *groupthreadpostinreplyto.GroupThreadPostInReplyToClient
	GroupThreadPostInReplyToAttachment                                  *groupthreadpostinreplytoattachment.GroupThreadPostInReplyToAttachmentClient
	GroupThreadPostInReplyToExtension                                   *groupthreadpostinreplytoextension.GroupThreadPostInReplyToExtensionClient
	GroupThreadPostInReplyToMention                                     *groupthreadpostinreplytomention.GroupThreadPostInReplyToMentionClient
	GroupThreadPostMention                                              *groupthreadpostmention.GroupThreadPostMentionClient
	GroupTransitiveMember                                               *grouptransitivemember.GroupTransitiveMemberClient
	GroupTransitiveMemberOf                                             *grouptransitivememberof.GroupTransitiveMemberOfClient
	SetGroupSiteTermStoreGroupSetChildren                               *setgroupsitetermstoregroupsetchildren.SetGroupSiteTermStoreGroupSetChildrenClient
	SetGroupSiteTermStoreGroupSetChildrenChildren                       *setgroupsitetermstoregroupsetchildrenchildren.SetGroupSiteTermStoreGroupSetChildrenChildrenClient
	SetGroupSiteTermStoreGroupSetChildrenChildrenRelation               *setgroupsitetermstoregroupsetchildrenchildrenrelation.SetGroupSiteTermStoreGroupSetChildrenChildrenRelationClient
	SetGroupSiteTermStoreGroupSetChildrenRelation                       *setgroupsitetermstoregroupsetchildrenrelation.SetGroupSiteTermStoreGroupSetChildrenRelationClient
	SetGroupSiteTermStoreGroupSetRelation                               *setgroupsitetermstoregroupsetrelation.SetGroupSiteTermStoreGroupSetRelationClient
	SetGroupSiteTermStoreGroupSetTerm                                   *setgroupsitetermstoregroupsetterm.SetGroupSiteTermStoreGroupSetTermClient
	SetGroupSiteTermStoreGroupSetTermChildren                           *setgroupsitetermstoregroupsettermchildren.SetGroupSiteTermStoreGroupSetTermChildrenClient
	SetGroupSiteTermStoreGroupSetTermChildrenRelation                   *setgroupsitetermstoregroupsettermchildrenrelation.SetGroupSiteTermStoreGroupSetTermChildrenRelationClient
	SetGroupSiteTermStoreGroupSetTermRelation                           *setgroupsitetermstoregroupsettermrelation.SetGroupSiteTermStoreGroupSetTermRelationClient
	SetGroupSiteTermStoreSetChildren                                    *setgroupsitetermstoresetchildren.SetGroupSiteTermStoreSetChildrenClient
	SetGroupSiteTermStoreSetChildrenChildren                            *setgroupsitetermstoresetchildrenchildren.SetGroupSiteTermStoreSetChildrenChildrenClient
	SetGroupSiteTermStoreSetChildrenChildrenRelation                    *setgroupsitetermstoresetchildrenchildrenrelation.SetGroupSiteTermStoreSetChildrenChildrenRelationClient
	SetGroupSiteTermStoreSetChildrenRelation                            *setgroupsitetermstoresetchildrenrelation.SetGroupSiteTermStoreSetChildrenRelationClient
	SetGroupSiteTermStoreSetParentGroupSetChildren                      *setgroupsitetermstoresetparentgroupsetchildren.SetGroupSiteTermStoreSetParentGroupSetChildrenClient
	SetGroupSiteTermStoreSetParentGroupSetChildrenChildren              *setgroupsitetermstoresetparentgroupsetchildrenchildren.SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenClient
	SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelation      *setgroupsitetermstoresetparentgroupsetchildrenchildrenrelation.SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient
	SetGroupSiteTermStoreSetParentGroupSetChildrenRelation              *setgroupsitetermstoresetparentgroupsetchildrenrelation.SetGroupSiteTermStoreSetParentGroupSetChildrenRelationClient
	SetGroupSiteTermStoreSetParentGroupSetRelation                      *setgroupsitetermstoresetparentgroupsetrelation.SetGroupSiteTermStoreSetParentGroupSetRelationClient
	SetGroupSiteTermStoreSetParentGroupSetTerm                          *setgroupsitetermstoresetparentgroupsetterm.SetGroupSiteTermStoreSetParentGroupSetTermClient
	SetGroupSiteTermStoreSetParentGroupSetTermChildren                  *setgroupsitetermstoresetparentgroupsettermchildren.SetGroupSiteTermStoreSetParentGroupSetTermChildrenClient
	SetGroupSiteTermStoreSetParentGroupSetTermChildrenRelation          *setgroupsitetermstoresetparentgroupsettermchildrenrelation.SetGroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient
	SetGroupSiteTermStoreSetParentGroupSetTermRelation                  *setgroupsitetermstoresetparentgroupsettermrelation.SetGroupSiteTermStoreSetParentGroupSetTermRelationClient
	SetGroupSiteTermStoreSetRelation                                    *setgroupsitetermstoresetrelation.SetGroupSiteTermStoreSetRelationClient
	SetGroupSiteTermStoreSetTerm                                        *setgroupsitetermstoresetterm.SetGroupSiteTermStoreSetTermClient
	SetGroupSiteTermStoreSetTermChildren                                *setgroupsitetermstoresettermchildren.SetGroupSiteTermStoreSetTermChildrenClient
	SetGroupSiteTermStoreSetTermChildrenRelation                        *setgroupsitetermstoresettermchildrenrelation.SetGroupSiteTermStoreSetTermChildrenRelationClient
	SetGroupSiteTermStoreSetTermRelation                                *setgroupsitetermstoresettermrelation.SetGroupSiteTermStoreSetTermRelationClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	groupAcceptedSenderClient, err := groupacceptedsender.NewGroupAcceptedSenderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupAcceptedSender client: %+v", err)
	}
	configureFunc(groupAcceptedSenderClient.Client)

	groupAppRoleAssignmentClient, err := groupapproleassignment.NewGroupAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupAppRoleAssignment client: %+v", err)
	}
	configureFunc(groupAppRoleAssignmentClient.Client)

	groupCalendarCalendarPermissionClient, err := groupcalendarcalendarpermission.NewGroupCalendarCalendarPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarPermission client: %+v", err)
	}
	configureFunc(groupCalendarCalendarPermissionClient.Client)

	groupCalendarCalendarViewAttachmentClient, err := groupcalendarcalendarviewattachment.NewGroupCalendarCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewAttachment client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewAttachmentClient.Client)

	groupCalendarCalendarViewCalendarClient, err := groupcalendarcalendarviewcalendar.NewGroupCalendarCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewCalendar client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewCalendarClient.Client)

	groupCalendarCalendarViewClient, err := groupcalendarcalendarview.NewGroupCalendarCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarView client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceAttachmentClient, err := groupcalendarcalendarviewexceptionoccurrenceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceAttachmentClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceCalendarClient, err := groupcalendarcalendarviewexceptionoccurrencecalendar.NewGroupCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceCalendarClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceClient, err := groupcalendarcalendarviewexceptionoccurrence.NewGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceExtensionClient, err := groupcalendarcalendarviewexceptionoccurrenceextension.NewGroupCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceExtensionClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := groupcalendarcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, err := groupcalendarcalendarviewexceptionoccurrenceinstancecalendar.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceInstanceClient, err := groupcalendarcalendarviewexceptionoccurrenceinstance.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceInstanceClient.Client)

	groupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient, err := groupcalendarcalendarviewexceptionoccurrenceinstanceextension.NewGroupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

	groupCalendarCalendarViewExtensionClient, err := groupcalendarcalendarviewextension.NewGroupCalendarCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewExtension client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewExtensionClient.Client)

	groupCalendarCalendarViewInstanceAttachmentClient, err := groupcalendarcalendarviewinstanceattachment.NewGroupCalendarCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceAttachmentClient.Client)

	groupCalendarCalendarViewInstanceCalendarClient, err := groupcalendarcalendarviewinstancecalendar.NewGroupCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceCalendarClient.Client)

	groupCalendarCalendarViewInstanceClient, err := groupcalendarcalendarviewinstance.NewGroupCalendarCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstance client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceClient.Client)

	groupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := groupcalendarcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	groupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, err := groupcalendarcalendarviewinstanceexceptionoccurrencecalendar.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	groupCalendarCalendarViewInstanceExceptionOccurrenceClient, err := groupcalendarcalendarviewinstanceexceptionoccurrence.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceExceptionOccurrenceClient.Client)

	groupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, err := groupcalendarcalendarviewinstanceexceptionoccurrenceextension.NewGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

	groupCalendarCalendarViewInstanceExtensionClient, err := groupcalendarcalendarviewinstanceextension.NewGroupCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(groupCalendarCalendarViewInstanceExtensionClient.Client)

	groupCalendarClient, err := groupcalendar.NewGroupCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendar client: %+v", err)
	}
	configureFunc(groupCalendarClient.Client)

	groupCalendarEventAttachmentClient, err := groupcalendareventattachment.NewGroupCalendarEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventAttachment client: %+v", err)
	}
	configureFunc(groupCalendarEventAttachmentClient.Client)

	groupCalendarEventCalendarClient, err := groupcalendareventcalendar.NewGroupCalendarEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventCalendar client: %+v", err)
	}
	configureFunc(groupCalendarEventCalendarClient.Client)

	groupCalendarEventClient, err := groupcalendarevent.NewGroupCalendarEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEvent client: %+v", err)
	}
	configureFunc(groupCalendarEventClient.Client)

	groupCalendarEventExceptionOccurrenceAttachmentClient, err := groupcalendareventexceptionoccurrenceattachment.NewGroupCalendarEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceAttachmentClient.Client)

	groupCalendarEventExceptionOccurrenceCalendarClient, err := groupcalendareventexceptionoccurrencecalendar.NewGroupCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceCalendarClient.Client)

	groupCalendarEventExceptionOccurrenceClient, err := groupcalendareventexceptionoccurrence.NewGroupCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceClient.Client)

	groupCalendarEventExceptionOccurrenceExtensionClient, err := groupcalendareventexceptionoccurrenceextension.NewGroupCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceExtensionClient.Client)

	groupCalendarEventExceptionOccurrenceInstanceAttachmentClient, err := groupcalendareventexceptionoccurrenceinstanceattachment.NewGroupCalendarEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceInstanceAttachmentClient.Client)

	groupCalendarEventExceptionOccurrenceInstanceCalendarClient, err := groupcalendareventexceptionoccurrenceinstancecalendar.NewGroupCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceInstanceCalendarClient.Client)

	groupCalendarEventExceptionOccurrenceInstanceClient, err := groupcalendareventexceptionoccurrenceinstance.NewGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceInstanceClient.Client)

	groupCalendarEventExceptionOccurrenceInstanceExtensionClient, err := groupcalendareventexceptionoccurrenceinstanceextension.NewGroupCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(groupCalendarEventExceptionOccurrenceInstanceExtensionClient.Client)

	groupCalendarEventExtensionClient, err := groupcalendareventextension.NewGroupCalendarEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventExtension client: %+v", err)
	}
	configureFunc(groupCalendarEventExtensionClient.Client)

	groupCalendarEventInstanceAttachmentClient, err := groupcalendareventinstanceattachment.NewGroupCalendarEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstanceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceAttachmentClient.Client)

	groupCalendarEventInstanceCalendarClient, err := groupcalendareventinstancecalendar.NewGroupCalendarEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstanceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceCalendarClient.Client)

	groupCalendarEventInstanceClient, err := groupcalendareventinstance.NewGroupCalendarEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstance client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceClient.Client)

	groupCalendarEventInstanceExceptionOccurrenceAttachmentClient, err := groupcalendareventinstanceexceptionoccurrenceattachment.NewGroupCalendarEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceExceptionOccurrenceAttachmentClient.Client)

	groupCalendarEventInstanceExceptionOccurrenceCalendarClient, err := groupcalendareventinstanceexceptionoccurrencecalendar.NewGroupCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceExceptionOccurrenceCalendarClient.Client)

	groupCalendarEventInstanceExceptionOccurrenceClient, err := groupcalendareventinstanceexceptionoccurrence.NewGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceExceptionOccurrenceClient.Client)

	groupCalendarEventInstanceExceptionOccurrenceExtensionClient, err := groupcalendareventinstanceexceptionoccurrenceextension.NewGroupCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceExceptionOccurrenceExtensionClient.Client)

	groupCalendarEventInstanceExtensionClient, err := groupcalendareventinstanceextension.NewGroupCalendarEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarEventInstanceExtension client: %+v", err)
	}
	configureFunc(groupCalendarEventInstanceExtensionClient.Client)

	groupCalendarViewAttachmentClient, err := groupcalendarviewattachment.NewGroupCalendarViewAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewAttachment client: %+v", err)
	}
	configureFunc(groupCalendarViewAttachmentClient.Client)

	groupCalendarViewCalendarClient, err := groupcalendarviewcalendar.NewGroupCalendarViewCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewCalendar client: %+v", err)
	}
	configureFunc(groupCalendarViewCalendarClient.Client)

	groupCalendarViewClient, err := groupcalendarview.NewGroupCalendarViewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarView client: %+v", err)
	}
	configureFunc(groupCalendarViewClient.Client)

	groupCalendarViewExceptionOccurrenceAttachmentClient, err := groupcalendarviewexceptionoccurrenceattachment.NewGroupCalendarViewExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceAttachmentClient.Client)

	groupCalendarViewExceptionOccurrenceCalendarClient, err := groupcalendarviewexceptionoccurrencecalendar.NewGroupCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceCalendarClient.Client)

	groupCalendarViewExceptionOccurrenceClient, err := groupcalendarviewexceptionoccurrence.NewGroupCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceClient.Client)

	groupCalendarViewExceptionOccurrenceExtensionClient, err := groupcalendarviewexceptionoccurrenceextension.NewGroupCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceExtensionClient.Client)

	groupCalendarViewExceptionOccurrenceInstanceAttachmentClient, err := groupcalendarviewexceptionoccurrenceinstanceattachment.NewGroupCalendarViewExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceInstanceAttachmentClient.Client)

	groupCalendarViewExceptionOccurrenceInstanceCalendarClient, err := groupcalendarviewexceptionoccurrenceinstancecalendar.NewGroupCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceInstanceCalendarClient.Client)

	groupCalendarViewExceptionOccurrenceInstanceClient, err := groupcalendarviewexceptionoccurrenceinstance.NewGroupCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceInstanceClient.Client)

	groupCalendarViewExceptionOccurrenceInstanceExtensionClient, err := groupcalendarviewexceptionoccurrenceinstanceextension.NewGroupCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(groupCalendarViewExceptionOccurrenceInstanceExtensionClient.Client)

	groupCalendarViewExtensionClient, err := groupcalendarviewextension.NewGroupCalendarViewExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewExtension client: %+v", err)
	}
	configureFunc(groupCalendarViewExtensionClient.Client)

	groupCalendarViewInstanceAttachmentClient, err := groupcalendarviewinstanceattachment.NewGroupCalendarViewInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstanceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceAttachmentClient.Client)

	groupCalendarViewInstanceCalendarClient, err := groupcalendarviewinstancecalendar.NewGroupCalendarViewInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstanceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceCalendarClient.Client)

	groupCalendarViewInstanceClient, err := groupcalendarviewinstance.NewGroupCalendarViewInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstance client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceClient.Client)

	groupCalendarViewInstanceExceptionOccurrenceAttachmentClient, err := groupcalendarviewinstanceexceptionoccurrenceattachment.NewGroupCalendarViewInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceExceptionOccurrenceAttachmentClient.Client)

	groupCalendarViewInstanceExceptionOccurrenceCalendarClient, err := groupcalendarviewinstanceexceptionoccurrencecalendar.NewGroupCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceExceptionOccurrenceCalendarClient.Client)

	groupCalendarViewInstanceExceptionOccurrenceClient, err := groupcalendarviewinstanceexceptionoccurrence.NewGroupCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceExceptionOccurrenceClient.Client)

	groupCalendarViewInstanceExceptionOccurrenceExtensionClient, err := groupcalendarviewinstanceexceptionoccurrenceextension.NewGroupCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceExceptionOccurrenceExtensionClient.Client)

	groupCalendarViewInstanceExtensionClient, err := groupcalendarviewinstanceextension.NewGroupCalendarViewInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCalendarViewInstanceExtension client: %+v", err)
	}
	configureFunc(groupCalendarViewInstanceExtensionClient.Client)

	groupClient, err := group.NewGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Group client: %+v", err)
	}
	configureFunc(groupClient.Client)

	groupConversationClient, err := groupconversation.NewGroupConversationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversation client: %+v", err)
	}
	configureFunc(groupConversationClient.Client)

	groupConversationThreadClient, err := groupconversationthread.NewGroupConversationThreadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThread client: %+v", err)
	}
	configureFunc(groupConversationThreadClient.Client)

	groupConversationThreadPostAttachmentClient, err := groupconversationthreadpostattachment.NewGroupConversationThreadPostAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPostAttachment client: %+v", err)
	}
	configureFunc(groupConversationThreadPostAttachmentClient.Client)

	groupConversationThreadPostClient, err := groupconversationthreadpost.NewGroupConversationThreadPostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPost client: %+v", err)
	}
	configureFunc(groupConversationThreadPostClient.Client)

	groupConversationThreadPostExtensionClient, err := groupconversationthreadpostextension.NewGroupConversationThreadPostExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPostExtension client: %+v", err)
	}
	configureFunc(groupConversationThreadPostExtensionClient.Client)

	groupConversationThreadPostInReplyToAttachmentClient, err := groupconversationthreadpostinreplytoattachment.NewGroupConversationThreadPostInReplyToAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPostInReplyToAttachment client: %+v", err)
	}
	configureFunc(groupConversationThreadPostInReplyToAttachmentClient.Client)

	groupConversationThreadPostInReplyToClient, err := groupconversationthreadpostinreplyto.NewGroupConversationThreadPostInReplyToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPostInReplyTo client: %+v", err)
	}
	configureFunc(groupConversationThreadPostInReplyToClient.Client)

	groupConversationThreadPostInReplyToExtensionClient, err := groupconversationthreadpostinreplytoextension.NewGroupConversationThreadPostInReplyToExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPostInReplyToExtension client: %+v", err)
	}
	configureFunc(groupConversationThreadPostInReplyToExtensionClient.Client)

	groupConversationThreadPostInReplyToMentionClient, err := groupconversationthreadpostinreplytomention.NewGroupConversationThreadPostInReplyToMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPostInReplyToMention client: %+v", err)
	}
	configureFunc(groupConversationThreadPostInReplyToMentionClient.Client)

	groupConversationThreadPostMentionClient, err := groupconversationthreadpostmention.NewGroupConversationThreadPostMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupConversationThreadPostMention client: %+v", err)
	}
	configureFunc(groupConversationThreadPostMentionClient.Client)

	groupCreatedOnBehalfOfClient, err := groupcreatedonbehalfof.NewGroupCreatedOnBehalfOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupCreatedOnBehalfOf client: %+v", err)
	}
	configureFunc(groupCreatedOnBehalfOfClient.Client)

	groupDriveClient, err := groupdrive.NewGroupDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupDrive client: %+v", err)
	}
	configureFunc(groupDriveClient.Client)

	groupEndpointClient, err := groupendpoint.NewGroupEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEndpoint client: %+v", err)
	}
	configureFunc(groupEndpointClient.Client)

	groupEventAttachmentClient, err := groupeventattachment.NewGroupEventAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventAttachment client: %+v", err)
	}
	configureFunc(groupEventAttachmentClient.Client)

	groupEventCalendarClient, err := groupeventcalendar.NewGroupEventCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventCalendar client: %+v", err)
	}
	configureFunc(groupEventCalendarClient.Client)

	groupEventClient, err := groupevent.NewGroupEventClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEvent client: %+v", err)
	}
	configureFunc(groupEventClient.Client)

	groupEventExceptionOccurrenceAttachmentClient, err := groupeventexceptionoccurrenceattachment.NewGroupEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceAttachmentClient.Client)

	groupEventExceptionOccurrenceCalendarClient, err := groupeventexceptionoccurrencecalendar.NewGroupEventExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceCalendarClient.Client)

	groupEventExceptionOccurrenceClient, err := groupeventexceptionoccurrence.NewGroupEventExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceClient.Client)

	groupEventExceptionOccurrenceExtensionClient, err := groupeventexceptionoccurrenceextension.NewGroupEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceExtensionClient.Client)

	groupEventExceptionOccurrenceInstanceAttachmentClient, err := groupeventexceptionoccurrenceinstanceattachment.NewGroupEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrenceInstanceAttachment client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceInstanceAttachmentClient.Client)

	groupEventExceptionOccurrenceInstanceCalendarClient, err := groupeventexceptionoccurrenceinstancecalendar.NewGroupEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrenceInstanceCalendar client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceInstanceCalendarClient.Client)

	groupEventExceptionOccurrenceInstanceClient, err := groupeventexceptionoccurrenceinstance.NewGroupEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrenceInstance client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceInstanceClient.Client)

	groupEventExceptionOccurrenceInstanceExtensionClient, err := groupeventexceptionoccurrenceinstanceextension.NewGroupEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExceptionOccurrenceInstanceExtension client: %+v", err)
	}
	configureFunc(groupEventExceptionOccurrenceInstanceExtensionClient.Client)

	groupEventExtensionClient, err := groupeventextension.NewGroupEventExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventExtension client: %+v", err)
	}
	configureFunc(groupEventExtensionClient.Client)

	groupEventInstanceAttachmentClient, err := groupeventinstanceattachment.NewGroupEventInstanceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstanceAttachment client: %+v", err)
	}
	configureFunc(groupEventInstanceAttachmentClient.Client)

	groupEventInstanceCalendarClient, err := groupeventinstancecalendar.NewGroupEventInstanceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstanceCalendar client: %+v", err)
	}
	configureFunc(groupEventInstanceCalendarClient.Client)

	groupEventInstanceClient, err := groupeventinstance.NewGroupEventInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstance client: %+v", err)
	}
	configureFunc(groupEventInstanceClient.Client)

	groupEventInstanceExceptionOccurrenceAttachmentClient, err := groupeventinstanceexceptionoccurrenceattachment.NewGroupEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstanceExceptionOccurrenceAttachment client: %+v", err)
	}
	configureFunc(groupEventInstanceExceptionOccurrenceAttachmentClient.Client)

	groupEventInstanceExceptionOccurrenceCalendarClient, err := groupeventinstanceexceptionoccurrencecalendar.NewGroupEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstanceExceptionOccurrenceCalendar client: %+v", err)
	}
	configureFunc(groupEventInstanceExceptionOccurrenceCalendarClient.Client)

	groupEventInstanceExceptionOccurrenceClient, err := groupeventinstanceexceptionoccurrence.NewGroupEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstanceExceptionOccurrence client: %+v", err)
	}
	configureFunc(groupEventInstanceExceptionOccurrenceClient.Client)

	groupEventInstanceExceptionOccurrenceExtensionClient, err := groupeventinstanceexceptionoccurrenceextension.NewGroupEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstanceExceptionOccurrenceExtension client: %+v", err)
	}
	configureFunc(groupEventInstanceExceptionOccurrenceExtensionClient.Client)

	groupEventInstanceExtensionClient, err := groupeventinstanceextension.NewGroupEventInstanceExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupEventInstanceExtension client: %+v", err)
	}
	configureFunc(groupEventInstanceExtensionClient.Client)

	groupExtensionClient, err := groupextension.NewGroupExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupExtension client: %+v", err)
	}
	configureFunc(groupExtensionClient.Client)

	groupGroupLifecyclePolicyClient, err := groupgrouplifecyclepolicy.NewGroupGroupLifecyclePolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupGroupLifecyclePolicy client: %+v", err)
	}
	configureFunc(groupGroupLifecyclePolicyClient.Client)

	groupMemberClient, err := groupmember.NewGroupMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupMember client: %+v", err)
	}
	configureFunc(groupMemberClient.Client)

	groupMemberOfClient, err := groupmemberof.NewGroupMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupMemberOf client: %+v", err)
	}
	configureFunc(groupMemberOfClient.Client)

	groupMembersWithLicenseErrorClient, err := groupmemberswithlicenseerror.NewGroupMembersWithLicenseErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupMembersWithLicenseError client: %+v", err)
	}
	configureFunc(groupMembersWithLicenseErrorClient.Client)

	groupOnenoteClient, err := grouponenote.NewGroupOnenoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenote client: %+v", err)
	}
	configureFunc(groupOnenoteClient.Client)

	groupOnenoteNotebookClient, err := grouponenotenotebook.NewGroupOnenoteNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookClient.Client)

	groupOnenoteNotebookSectionClient, err := grouponenotenotebooksection.NewGroupOnenoteNotebookSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSection client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionClient.Client)

	groupOnenoteNotebookSectionGroupClient, err := grouponenotenotebooksectiongroup.NewGroupOnenoteNotebookSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupClient.Client)

	groupOnenoteNotebookSectionGroupParentNotebookClient, err := grouponenotenotebooksectiongroupparentnotebook.NewGroupOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupParentNotebookClient.Client)

	groupOnenoteNotebookSectionGroupParentSectionGroupClient, err := grouponenotenotebooksectiongroupparentsectiongroup.NewGroupOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupParentSectionGroupClient.Client)

	groupOnenoteNotebookSectionGroupSectionClient, err := grouponenotenotebooksectiongroupsection.NewGroupOnenoteNotebookSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSection client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionClient.Client)

	groupOnenoteNotebookSectionGroupSectionGroupClient, err := grouponenotenotebooksectiongroupsectiongroup.NewGroupOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionGroupClient.Client)

	groupOnenoteNotebookSectionGroupSectionPageClient, err := grouponenotenotebooksectiongroupsectionpage.NewGroupOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionPageClient.Client)

	groupOnenoteNotebookSectionGroupSectionPageContentClient, err := grouponenotenotebooksectiongroupsectionpagecontent.NewGroupOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionPageContentClient.Client)

	groupOnenoteNotebookSectionGroupSectionPageParentNotebookClient, err := grouponenotenotebooksectiongroupsectionpageparentnotebook.NewGroupOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionPageParentNotebookClient.Client)

	groupOnenoteNotebookSectionGroupSectionPageParentSectionClient, err := grouponenotenotebooksectiongroupsectionpageparentsection.NewGroupOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionPageParentSectionClient.Client)

	groupOnenoteNotebookSectionGroupSectionParentNotebookClient, err := grouponenotenotebooksectiongroupsectionparentnotebook.NewGroupOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionParentNotebookClient.Client)

	groupOnenoteNotebookSectionGroupSectionParentSectionGroupClient, err := grouponenotenotebooksectiongroupsectionparentsectiongroup.NewGroupOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionGroupSectionParentSectionGroupClient.Client)

	groupOnenoteNotebookSectionPageClient, err := grouponenotenotebooksectionpage.NewGroupOnenoteNotebookSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionPage client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionPageClient.Client)

	groupOnenoteNotebookSectionPageContentClient, err := grouponenotenotebooksectionpagecontent.NewGroupOnenoteNotebookSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionPageContent client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionPageContentClient.Client)

	groupOnenoteNotebookSectionPageParentNotebookClient, err := grouponenotenotebooksectionpageparentnotebook.NewGroupOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionPageParentNotebookClient.Client)

	groupOnenoteNotebookSectionPageParentSectionClient, err := grouponenotenotebooksectionpageparentsection.NewGroupOnenoteNotebookSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionPageParentSectionClient.Client)

	groupOnenoteNotebookSectionParentNotebookClient, err := grouponenotenotebooksectionparentnotebook.NewGroupOnenoteNotebookSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionParentNotebookClient.Client)

	groupOnenoteNotebookSectionParentSectionGroupClient, err := grouponenotenotebooksectionparentsectiongroup.NewGroupOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteNotebookSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteNotebookSectionParentSectionGroupClient.Client)

	groupOnenoteOperationClient, err := grouponenoteoperation.NewGroupOnenoteOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteOperation client: %+v", err)
	}
	configureFunc(groupOnenoteOperationClient.Client)

	groupOnenotePageClient, err := grouponenotepage.NewGroupOnenotePageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenotePage client: %+v", err)
	}
	configureFunc(groupOnenotePageClient.Client)

	groupOnenotePageContentClient, err := grouponenotepagecontent.NewGroupOnenotePageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenotePageContent client: %+v", err)
	}
	configureFunc(groupOnenotePageContentClient.Client)

	groupOnenotePageParentNotebookClient, err := grouponenotepageparentnotebook.NewGroupOnenotePageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenotePageParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenotePageParentNotebookClient.Client)

	groupOnenotePageParentSectionClient, err := grouponenotepageparentsection.NewGroupOnenotePageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenotePageParentSection client: %+v", err)
	}
	configureFunc(groupOnenotePageParentSectionClient.Client)

	groupOnenoteResourceClient, err := grouponenoteresource.NewGroupOnenoteResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteResource client: %+v", err)
	}
	configureFunc(groupOnenoteResourceClient.Client)

	groupOnenoteResourceContentClient, err := grouponenoteresourcecontent.NewGroupOnenoteResourceContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteResourceContent client: %+v", err)
	}
	configureFunc(groupOnenoteResourceContentClient.Client)

	groupOnenoteSectionClient, err := grouponenotesection.NewGroupOnenoteSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSection client: %+v", err)
	}
	configureFunc(groupOnenoteSectionClient.Client)

	groupOnenoteSectionGroupClient, err := grouponenotesectiongroup.NewGroupOnenoteSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupClient.Client)

	groupOnenoteSectionGroupParentNotebookClient, err := grouponenotesectiongroupparentnotebook.NewGroupOnenoteSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupParentNotebookClient.Client)

	groupOnenoteSectionGroupParentSectionGroupClient, err := grouponenotesectiongroupparentsectiongroup.NewGroupOnenoteSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupParentSectionGroupClient.Client)

	groupOnenoteSectionGroupSectionClient, err := grouponenotesectiongroupsection.NewGroupOnenoteSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSection client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionClient.Client)

	groupOnenoteSectionGroupSectionGroupClient, err := grouponenotesectiongroupsectiongroup.NewGroupOnenoteSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionGroupClient.Client)

	groupOnenoteSectionGroupSectionPageClient, err := grouponenotesectiongroupsectionpage.NewGroupOnenoteSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionPageClient.Client)

	groupOnenoteSectionGroupSectionPageContentClient, err := grouponenotesectiongroupsectionpagecontent.NewGroupOnenoteSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionPageContentClient.Client)

	groupOnenoteSectionGroupSectionPageParentNotebookClient, err := grouponenotesectiongroupsectionpageparentnotebook.NewGroupOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionPageParentNotebookClient.Client)

	groupOnenoteSectionGroupSectionPageParentSectionClient, err := grouponenotesectiongroupsectionpageparentsection.NewGroupOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionPageParentSectionClient.Client)

	groupOnenoteSectionGroupSectionParentNotebookClient, err := grouponenotesectiongroupsectionparentnotebook.NewGroupOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionParentNotebookClient.Client)

	groupOnenoteSectionGroupSectionParentSectionGroupClient, err := grouponenotesectiongroupsectionparentsectiongroup.NewGroupOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteSectionGroupSectionParentSectionGroupClient.Client)

	groupOnenoteSectionPageClient, err := grouponenotesectionpage.NewGroupOnenoteSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionPage client: %+v", err)
	}
	configureFunc(groupOnenoteSectionPageClient.Client)

	groupOnenoteSectionPageContentClient, err := grouponenotesectionpagecontent.NewGroupOnenoteSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionPageContent client: %+v", err)
	}
	configureFunc(groupOnenoteSectionPageContentClient.Client)

	groupOnenoteSectionPageParentNotebookClient, err := grouponenotesectionpageparentnotebook.NewGroupOnenoteSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteSectionPageParentNotebookClient.Client)

	groupOnenoteSectionPageParentSectionClient, err := grouponenotesectionpageparentsection.NewGroupOnenoteSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupOnenoteSectionPageParentSectionClient.Client)

	groupOnenoteSectionParentNotebookClient, err := grouponenotesectionparentnotebook.NewGroupOnenoteSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupOnenoteSectionParentNotebookClient.Client)

	groupOnenoteSectionParentSectionGroupClient, err := grouponenotesectionparentsectiongroup.NewGroupOnenoteSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOnenoteSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupOnenoteSectionParentSectionGroupClient.Client)

	groupOwnerClient, err := groupowner.NewGroupOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupOwner client: %+v", err)
	}
	configureFunc(groupOwnerClient.Client)

	groupPermissionGrantClient, err := grouppermissiongrant.NewGroupPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPermissionGrant client: %+v", err)
	}
	configureFunc(groupPermissionGrantClient.Client)

	groupPhotoClient, err := groupphoto.NewGroupPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPhoto client: %+v", err)
	}
	configureFunc(groupPhotoClient.Client)

	groupPlannerClient, err := groupplanner.NewGroupPlannerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlanner client: %+v", err)
	}
	configureFunc(groupPlannerClient.Client)

	groupPlannerPlanBucketClient, err := groupplannerplanbucket.NewGroupPlannerPlanBucketClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanBucket client: %+v", err)
	}
	configureFunc(groupPlannerPlanBucketClient.Client)

	groupPlannerPlanBucketTaskAssignedToTaskBoardFormatClient, err := groupplannerplanbuckettaskassignedtotaskboardformat.NewGroupPlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanBucketTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(groupPlannerPlanBucketTaskAssignedToTaskBoardFormatClient.Client)

	groupPlannerPlanBucketTaskBucketTaskBoardFormatClient, err := groupplannerplanbuckettaskbuckettaskboardformat.NewGroupPlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanBucketTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(groupPlannerPlanBucketTaskBucketTaskBoardFormatClient.Client)

	groupPlannerPlanBucketTaskClient, err := groupplannerplanbuckettask.NewGroupPlannerPlanBucketTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanBucketTask client: %+v", err)
	}
	configureFunc(groupPlannerPlanBucketTaskClient.Client)

	groupPlannerPlanBucketTaskDetailClient, err := groupplannerplanbuckettaskdetail.NewGroupPlannerPlanBucketTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanBucketTaskDetail client: %+v", err)
	}
	configureFunc(groupPlannerPlanBucketTaskDetailClient.Client)

	groupPlannerPlanBucketTaskProgressTaskBoardFormatClient, err := groupplannerplanbuckettaskprogresstaskboardformat.NewGroupPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanBucketTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(groupPlannerPlanBucketTaskProgressTaskBoardFormatClient.Client)

	groupPlannerPlanClient, err := groupplannerplan.NewGroupPlannerPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlan client: %+v", err)
	}
	configureFunc(groupPlannerPlanClient.Client)

	groupPlannerPlanDetailClient, err := groupplannerplandetail.NewGroupPlannerPlanDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanDetail client: %+v", err)
	}
	configureFunc(groupPlannerPlanDetailClient.Client)

	groupPlannerPlanTaskAssignedToTaskBoardFormatClient, err := groupplannerplantaskassignedtotaskboardformat.NewGroupPlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanTaskAssignedToTaskBoardFormat client: %+v", err)
	}
	configureFunc(groupPlannerPlanTaskAssignedToTaskBoardFormatClient.Client)

	groupPlannerPlanTaskBucketTaskBoardFormatClient, err := groupplannerplantaskbuckettaskboardformat.NewGroupPlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanTaskBucketTaskBoardFormat client: %+v", err)
	}
	configureFunc(groupPlannerPlanTaskBucketTaskBoardFormatClient.Client)

	groupPlannerPlanTaskClient, err := groupplannerplantask.NewGroupPlannerPlanTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanTask client: %+v", err)
	}
	configureFunc(groupPlannerPlanTaskClient.Client)

	groupPlannerPlanTaskDetailClient, err := groupplannerplantaskdetail.NewGroupPlannerPlanTaskDetailClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanTaskDetail client: %+v", err)
	}
	configureFunc(groupPlannerPlanTaskDetailClient.Client)

	groupPlannerPlanTaskProgressTaskBoardFormatClient, err := groupplannerplantaskprogresstaskboardformat.NewGroupPlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupPlannerPlanTaskProgressTaskBoardFormat client: %+v", err)
	}
	configureFunc(groupPlannerPlanTaskProgressTaskBoardFormatClient.Client)

	groupRejectedSenderClient, err := grouprejectedsender.NewGroupRejectedSenderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupRejectedSender client: %+v", err)
	}
	configureFunc(groupRejectedSenderClient.Client)

	groupSettingClient, err := groupsetting.NewGroupSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSetting client: %+v", err)
	}
	configureFunc(groupSettingClient.Client)

	groupSiteAnalyticAllTimeClient, err := groupsiteanalyticalltime.NewGroupSiteAnalyticAllTimeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteAnalyticAllTime client: %+v", err)
	}
	configureFunc(groupSiteAnalyticAllTimeClient.Client)

	groupSiteAnalyticClient, err := groupsiteanalytic.NewGroupSiteAnalyticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteAnalytic client: %+v", err)
	}
	configureFunc(groupSiteAnalyticClient.Client)

	groupSiteAnalyticItemActivityStatActivityClient, err := groupsiteanalyticitemactivitystatactivity.NewGroupSiteAnalyticItemActivityStatActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteAnalyticItemActivityStatActivity client: %+v", err)
	}
	configureFunc(groupSiteAnalyticItemActivityStatActivityClient.Client)

	groupSiteAnalyticItemActivityStatActivityDriveItemClient, err := groupsiteanalyticitemactivitystatactivitydriveitem.NewGroupSiteAnalyticItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteAnalyticItemActivityStatActivityDriveItem client: %+v", err)
	}
	configureFunc(groupSiteAnalyticItemActivityStatActivityDriveItemClient.Client)

	groupSiteAnalyticItemActivityStatActivityDriveItemContentClient, err := groupsiteanalyticitemactivitystatactivitydriveitemcontent.NewGroupSiteAnalyticItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteAnalyticItemActivityStatActivityDriveItemContent client: %+v", err)
	}
	configureFunc(groupSiteAnalyticItemActivityStatActivityDriveItemContentClient.Client)

	groupSiteAnalyticItemActivityStatClient, err := groupsiteanalyticitemactivitystat.NewGroupSiteAnalyticItemActivityStatClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteAnalyticItemActivityStat client: %+v", err)
	}
	configureFunc(groupSiteAnalyticItemActivityStatClient.Client)

	groupSiteAnalyticLastSevenDayClient, err := groupsiteanalyticlastsevenday.NewGroupSiteAnalyticLastSevenDayClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteAnalyticLastSevenDay client: %+v", err)
	}
	configureFunc(groupSiteAnalyticLastSevenDayClient.Client)

	groupSiteClient, err := groupsite.NewGroupSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSite client: %+v", err)
	}
	configureFunc(groupSiteClient.Client)

	groupSiteColumnClient, err := groupsitecolumn.NewGroupSiteColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteColumn client: %+v", err)
	}
	configureFunc(groupSiteColumnClient.Client)

	groupSiteColumnSourceColumnClient, err := groupsitecolumnsourcecolumn.NewGroupSiteColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteColumnSourceColumn client: %+v", err)
	}
	configureFunc(groupSiteColumnSourceColumnClient.Client)

	groupSiteContentTypeBaseClient, err := groupsitecontenttypebase.NewGroupSiteContentTypeBaseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteContentTypeBase client: %+v", err)
	}
	configureFunc(groupSiteContentTypeBaseClient.Client)

	groupSiteContentTypeBaseTypeClient, err := groupsitecontenttypebasetype.NewGroupSiteContentTypeBaseTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteContentTypeBaseType client: %+v", err)
	}
	configureFunc(groupSiteContentTypeBaseTypeClient.Client)

	groupSiteContentTypeClient, err := groupsitecontenttype.NewGroupSiteContentTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteContentType client: %+v", err)
	}
	configureFunc(groupSiteContentTypeClient.Client)

	groupSiteContentTypeColumnClient, err := groupsitecontenttypecolumn.NewGroupSiteContentTypeColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteContentTypeColumn client: %+v", err)
	}
	configureFunc(groupSiteContentTypeColumnClient.Client)

	groupSiteContentTypeColumnLinkClient, err := groupsitecontenttypecolumnlink.NewGroupSiteContentTypeColumnLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteContentTypeColumnLink client: %+v", err)
	}
	configureFunc(groupSiteContentTypeColumnLinkClient.Client)

	groupSiteContentTypeColumnPositionClient, err := groupsitecontenttypecolumnposition.NewGroupSiteContentTypeColumnPositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteContentTypeColumnPosition client: %+v", err)
	}
	configureFunc(groupSiteContentTypeColumnPositionClient.Client)

	groupSiteContentTypeColumnSourceColumnClient, err := groupsitecontenttypecolumnsourcecolumn.NewGroupSiteContentTypeColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteContentTypeColumnSourceColumn client: %+v", err)
	}
	configureFunc(groupSiteContentTypeColumnSourceColumnClient.Client)

	groupSiteCreatedByUserClient, err := groupsitecreatedbyuser.NewGroupSiteCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteCreatedByUser client: %+v", err)
	}
	configureFunc(groupSiteCreatedByUserClient.Client)

	groupSiteCreatedByUserMailboxSettingClient, err := groupsitecreatedbyusermailboxsetting.NewGroupSiteCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSiteCreatedByUserMailboxSettingClient.Client)

	groupSiteDriveClient, err := groupsitedrive.NewGroupSiteDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteDrive client: %+v", err)
	}
	configureFunc(groupSiteDriveClient.Client)

	groupSiteExternalColumnClient, err := groupsiteexternalcolumn.NewGroupSiteExternalColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteExternalColumn client: %+v", err)
	}
	configureFunc(groupSiteExternalColumnClient.Client)

	groupSiteInformationProtectionBitlockerClient, err := groupsiteinformationprotectionbitlocker.NewGroupSiteInformationProtectionBitlockerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionBitlocker client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionBitlockerClient.Client)

	groupSiteInformationProtectionBitlockerRecoveryKeyClient, err := groupsiteinformationprotectionbitlockerrecoverykey.NewGroupSiteInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionBitlockerRecoveryKey client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionBitlockerRecoveryKeyClient.Client)

	groupSiteInformationProtectionClient, err := groupsiteinformationprotection.NewGroupSiteInformationProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtection client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionClient.Client)

	groupSiteInformationProtectionDataLossPreventionPolicyClient, err := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteInformationProtectionDataLossPreventionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionDataLossPreventionPolicy client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionDataLossPreventionPolicyClient.Client)

	groupSiteInformationProtectionPolicyClient, err := groupsiteinformationprotectionpolicy.NewGroupSiteInformationProtectionPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionPolicy client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionPolicyClient.Client)

	groupSiteInformationProtectionPolicyLabelClient, err := groupsiteinformationprotectionpolicylabel.NewGroupSiteInformationProtectionPolicyLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionPolicyLabel client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionPolicyLabelClient.Client)

	groupSiteInformationProtectionSensitivityLabelClient, err := groupsiteinformationprotectionsensitivitylabel.NewGroupSiteInformationProtectionSensitivityLabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionSensitivityLabel client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionSensitivityLabelClient.Client)

	groupSiteInformationProtectionSensitivityLabelSublabelClient, err := groupsiteinformationprotectionsensitivitylabelsublabel.NewGroupSiteInformationProtectionSensitivityLabelSublabelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionSensitivityLabelSublabel client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionSensitivityLabelSublabelClient.Client)

	groupSiteInformationProtectionSensitivityPolicySettingClient, err := groupsiteinformationprotectionsensitivitypolicysetting.NewGroupSiteInformationProtectionSensitivityPolicySettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionSensitivityPolicySetting client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionSensitivityPolicySettingClient.Client)

	groupSiteInformationProtectionThreatAssessmentRequestClient, err := groupsiteinformationprotectionthreatassessmentrequest.NewGroupSiteInformationProtectionThreatAssessmentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionThreatAssessmentRequest client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionThreatAssessmentRequestClient.Client)

	groupSiteInformationProtectionThreatAssessmentRequestResultClient, err := groupsiteinformationprotectionthreatassessmentrequestresult.NewGroupSiteInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteInformationProtectionThreatAssessmentRequestResult client: %+v", err)
	}
	configureFunc(groupSiteInformationProtectionThreatAssessmentRequestResultClient.Client)

	groupSiteItemClient, err := groupsiteitem.NewGroupSiteItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteItem client: %+v", err)
	}
	configureFunc(groupSiteItemClient.Client)

	groupSiteLastModifiedByUserClient, err := groupsitelastmodifiedbyuser.NewGroupSiteLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteLastModifiedByUser client: %+v", err)
	}
	configureFunc(groupSiteLastModifiedByUserClient.Client)

	groupSiteLastModifiedByUserMailboxSettingClient, err := groupsitelastmodifiedbyusermailboxsetting.NewGroupSiteLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSiteLastModifiedByUserMailboxSettingClient.Client)

	groupSiteListActivityClient, err := groupsitelistactivity.NewGroupSiteListActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListActivity client: %+v", err)
	}
	configureFunc(groupSiteListActivityClient.Client)

	groupSiteListClient, err := groupsitelist.NewGroupSiteListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteList client: %+v", err)
	}
	configureFunc(groupSiteListClient.Client)

	groupSiteListColumnClient, err := groupsitelistcolumn.NewGroupSiteListColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListColumn client: %+v", err)
	}
	configureFunc(groupSiteListColumnClient.Client)

	groupSiteListColumnSourceColumnClient, err := groupsitelistcolumnsourcecolumn.NewGroupSiteListColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListColumnSourceColumn client: %+v", err)
	}
	configureFunc(groupSiteListColumnSourceColumnClient.Client)

	groupSiteListContentTypeBaseClient, err := groupsitelistcontenttypebase.NewGroupSiteListContentTypeBaseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListContentTypeBase client: %+v", err)
	}
	configureFunc(groupSiteListContentTypeBaseClient.Client)

	groupSiteListContentTypeBaseTypeClient, err := groupsitelistcontenttypebasetype.NewGroupSiteListContentTypeBaseTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListContentTypeBaseType client: %+v", err)
	}
	configureFunc(groupSiteListContentTypeBaseTypeClient.Client)

	groupSiteListContentTypeClient, err := groupsitelistcontenttype.NewGroupSiteListContentTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListContentType client: %+v", err)
	}
	configureFunc(groupSiteListContentTypeClient.Client)

	groupSiteListContentTypeColumnClient, err := groupsitelistcontenttypecolumn.NewGroupSiteListContentTypeColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListContentTypeColumn client: %+v", err)
	}
	configureFunc(groupSiteListContentTypeColumnClient.Client)

	groupSiteListContentTypeColumnLinkClient, err := groupsitelistcontenttypecolumnlink.NewGroupSiteListContentTypeColumnLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListContentTypeColumnLink client: %+v", err)
	}
	configureFunc(groupSiteListContentTypeColumnLinkClient.Client)

	groupSiteListContentTypeColumnPositionClient, err := groupsitelistcontenttypecolumnposition.NewGroupSiteListContentTypeColumnPositionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListContentTypeColumnPosition client: %+v", err)
	}
	configureFunc(groupSiteListContentTypeColumnPositionClient.Client)

	groupSiteListContentTypeColumnSourceColumnClient, err := groupsitelistcontenttypecolumnsourcecolumn.NewGroupSiteListContentTypeColumnSourceColumnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListContentTypeColumnSourceColumn client: %+v", err)
	}
	configureFunc(groupSiteListContentTypeColumnSourceColumnClient.Client)

	groupSiteListCreatedByUserClient, err := groupsitelistcreatedbyuser.NewGroupSiteListCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListCreatedByUser client: %+v", err)
	}
	configureFunc(groupSiteListCreatedByUserClient.Client)

	groupSiteListCreatedByUserMailboxSettingClient, err := groupsitelistcreatedbyusermailboxsetting.NewGroupSiteListCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSiteListCreatedByUserMailboxSettingClient.Client)

	groupSiteListDriveClient, err := groupsitelistdrive.NewGroupSiteListDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListDrive client: %+v", err)
	}
	configureFunc(groupSiteListDriveClient.Client)

	groupSiteListItemActivityClient, err := groupsitelistitemactivity.NewGroupSiteListItemActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemActivity client: %+v", err)
	}
	configureFunc(groupSiteListItemActivityClient.Client)

	groupSiteListItemActivityDriveItemClient, err := groupsitelistitemactivitydriveitem.NewGroupSiteListItemActivityDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemActivityDriveItem client: %+v", err)
	}
	configureFunc(groupSiteListItemActivityDriveItemClient.Client)

	groupSiteListItemActivityDriveItemContentClient, err := groupsitelistitemactivitydriveitemcontent.NewGroupSiteListItemActivityDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemActivityDriveItemContent client: %+v", err)
	}
	configureFunc(groupSiteListItemActivityDriveItemContentClient.Client)

	groupSiteListItemActivityListItemClient, err := groupsitelistitemactivitylistitem.NewGroupSiteListItemActivityListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemActivityListItem client: %+v", err)
	}
	configureFunc(groupSiteListItemActivityListItemClient.Client)

	groupSiteListItemAnalyticClient, err := groupsitelistitemanalytic.NewGroupSiteListItemAnalyticClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemAnalytic client: %+v", err)
	}
	configureFunc(groupSiteListItemAnalyticClient.Client)

	groupSiteListItemClient, err := groupsitelistitem.NewGroupSiteListItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItem client: %+v", err)
	}
	configureFunc(groupSiteListItemClient.Client)

	groupSiteListItemCreatedByUserClient, err := groupsitelistitemcreatedbyuser.NewGroupSiteListItemCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemCreatedByUser client: %+v", err)
	}
	configureFunc(groupSiteListItemCreatedByUserClient.Client)

	groupSiteListItemCreatedByUserMailboxSettingClient, err := groupsitelistitemcreatedbyusermailboxsetting.NewGroupSiteListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSiteListItemCreatedByUserMailboxSettingClient.Client)

	groupSiteListItemDocumentSetVersionClient, err := groupsitelistitemdocumentsetversion.NewGroupSiteListItemDocumentSetVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemDocumentSetVersion client: %+v", err)
	}
	configureFunc(groupSiteListItemDocumentSetVersionClient.Client)

	groupSiteListItemDocumentSetVersionFieldClient, err := groupsitelistitemdocumentsetversionfield.NewGroupSiteListItemDocumentSetVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemDocumentSetVersionField client: %+v", err)
	}
	configureFunc(groupSiteListItemDocumentSetVersionFieldClient.Client)

	groupSiteListItemDriveItemClient, err := groupsitelistitemdriveitem.NewGroupSiteListItemDriveItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemDriveItem client: %+v", err)
	}
	configureFunc(groupSiteListItemDriveItemClient.Client)

	groupSiteListItemDriveItemContentClient, err := groupsitelistitemdriveitemcontent.NewGroupSiteListItemDriveItemContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemDriveItemContent client: %+v", err)
	}
	configureFunc(groupSiteListItemDriveItemContentClient.Client)

	groupSiteListItemFieldClient, err := groupsitelistitemfield.NewGroupSiteListItemFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemField client: %+v", err)
	}
	configureFunc(groupSiteListItemFieldClient.Client)

	groupSiteListItemLastModifiedByUserClient, err := groupsitelistitemlastmodifiedbyuser.NewGroupSiteListItemLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemLastModifiedByUser client: %+v", err)
	}
	configureFunc(groupSiteListItemLastModifiedByUserClient.Client)

	groupSiteListItemLastModifiedByUserMailboxSettingClient, err := groupsitelistitemlastmodifiedbyusermailboxsetting.NewGroupSiteListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSiteListItemLastModifiedByUserMailboxSettingClient.Client)

	groupSiteListItemVersionClient, err := groupsitelistitemversion.NewGroupSiteListItemVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemVersion client: %+v", err)
	}
	configureFunc(groupSiteListItemVersionClient.Client)

	groupSiteListItemVersionFieldClient, err := groupsitelistitemversionfield.NewGroupSiteListItemVersionFieldClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListItemVersionField client: %+v", err)
	}
	configureFunc(groupSiteListItemVersionFieldClient.Client)

	groupSiteListLastModifiedByUserClient, err := groupsitelistlastmodifiedbyuser.NewGroupSiteListLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListLastModifiedByUser client: %+v", err)
	}
	configureFunc(groupSiteListLastModifiedByUserClient.Client)

	groupSiteListLastModifiedByUserMailboxSettingClient, err := groupsitelistlastmodifiedbyusermailboxsetting.NewGroupSiteListLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSiteListLastModifiedByUserMailboxSettingClient.Client)

	groupSiteListOperationClient, err := groupsitelistoperation.NewGroupSiteListOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListOperation client: %+v", err)
	}
	configureFunc(groupSiteListOperationClient.Client)

	groupSiteListSubscriptionClient, err := groupsitelistsubscription.NewGroupSiteListSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteListSubscription client: %+v", err)
	}
	configureFunc(groupSiteListSubscriptionClient.Client)

	groupSiteOnenoteClient, err := groupsiteonenote.NewGroupSiteOnenoteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenote client: %+v", err)
	}
	configureFunc(groupSiteOnenoteClient.Client)

	groupSiteOnenoteNotebookClient, err := groupsiteonenotenotebook.NewGroupSiteOnenoteNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookClient.Client)

	groupSiteOnenoteNotebookSectionClient, err := groupsiteonenotenotebooksection.NewGroupSiteOnenoteNotebookSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionClient.Client)

	groupSiteOnenoteNotebookSectionGroupClient, err := groupsiteonenotenotebooksectiongroup.NewGroupSiteOnenoteNotebookSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupClient.Client)

	groupSiteOnenoteNotebookSectionGroupParentNotebookClient, err := groupsiteonenotenotebooksectiongroupparentnotebook.NewGroupSiteOnenoteNotebookSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupParentNotebookClient.Client)

	groupSiteOnenoteNotebookSectionGroupParentSectionGroupClient, err := groupsiteonenotenotebooksectiongroupparentsectiongroup.NewGroupSiteOnenoteNotebookSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupParentSectionGroupClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionClient, err := groupsiteonenotenotebooksectiongroupsection.NewGroupSiteOnenoteNotebookSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionGroupClient, err := groupsiteonenotenotebooksectiongroupsectiongroup.NewGroupSiteOnenoteNotebookSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionGroupClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionPageClient, err := groupsiteonenotenotebooksectiongroupsectionpage.NewGroupSiteOnenoteNotebookSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionPageClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionPageContentClient, err := groupsiteonenotenotebooksectiongroupsectionpagecontent.NewGroupSiteOnenoteNotebookSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionPageContentClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient, err := groupsiteonenotenotebooksectiongroupsectionpageparentnotebook.NewGroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionPageParentSectionClient, err := groupsiteonenotenotebooksectiongroupsectionpageparentsection.NewGroupSiteOnenoteNotebookSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionPageParentSectionClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionParentNotebookClient, err := groupsiteonenotenotebooksectiongroupsectionparentnotebook.NewGroupSiteOnenoteNotebookSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionParentNotebookClient.Client)

	groupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient, err := groupsiteonenotenotebooksectiongroupsectionparentsectiongroup.NewGroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient.Client)

	groupSiteOnenoteNotebookSectionPageClient, err := groupsiteonenotenotebooksectionpage.NewGroupSiteOnenoteNotebookSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionPage client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionPageClient.Client)

	groupSiteOnenoteNotebookSectionPageContentClient, err := groupsiteonenotenotebooksectionpagecontent.NewGroupSiteOnenoteNotebookSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionPageContent client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionPageContentClient.Client)

	groupSiteOnenoteNotebookSectionPageParentNotebookClient, err := groupsiteonenotenotebooksectionpageparentnotebook.NewGroupSiteOnenoteNotebookSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionPageParentNotebookClient.Client)

	groupSiteOnenoteNotebookSectionPageParentSectionClient, err := groupsiteonenotenotebooksectionpageparentsection.NewGroupSiteOnenoteNotebookSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionPageParentSectionClient.Client)

	groupSiteOnenoteNotebookSectionParentNotebookClient, err := groupsiteonenotenotebooksectionparentnotebook.NewGroupSiteOnenoteNotebookSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionParentNotebookClient.Client)

	groupSiteOnenoteNotebookSectionParentSectionGroupClient, err := groupsiteonenotenotebooksectionparentsectiongroup.NewGroupSiteOnenoteNotebookSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteNotebookSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteNotebookSectionParentSectionGroupClient.Client)

	groupSiteOnenoteOperationClient, err := groupsiteonenoteoperation.NewGroupSiteOnenoteOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteOperation client: %+v", err)
	}
	configureFunc(groupSiteOnenoteOperationClient.Client)

	groupSiteOnenotePageClient, err := groupsiteonenotepage.NewGroupSiteOnenotePageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenotePage client: %+v", err)
	}
	configureFunc(groupSiteOnenotePageClient.Client)

	groupSiteOnenotePageContentClient, err := groupsiteonenotepagecontent.NewGroupSiteOnenotePageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenotePageContent client: %+v", err)
	}
	configureFunc(groupSiteOnenotePageContentClient.Client)

	groupSiteOnenotePageParentNotebookClient, err := groupsiteonenotepageparentnotebook.NewGroupSiteOnenotePageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenotePageParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenotePageParentNotebookClient.Client)

	groupSiteOnenotePageParentSectionClient, err := groupsiteonenotepageparentsection.NewGroupSiteOnenotePageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenotePageParentSection client: %+v", err)
	}
	configureFunc(groupSiteOnenotePageParentSectionClient.Client)

	groupSiteOnenoteResourceClient, err := groupsiteonenoteresource.NewGroupSiteOnenoteResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteResource client: %+v", err)
	}
	configureFunc(groupSiteOnenoteResourceClient.Client)

	groupSiteOnenoteResourceContentClient, err := groupsiteonenoteresourcecontent.NewGroupSiteOnenoteResourceContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteResourceContent client: %+v", err)
	}
	configureFunc(groupSiteOnenoteResourceContentClient.Client)

	groupSiteOnenoteSectionClient, err := groupsiteonenotesection.NewGroupSiteOnenoteSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionClient.Client)

	groupSiteOnenoteSectionGroupClient, err := groupsiteonenotesectiongroup.NewGroupSiteOnenoteSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupClient.Client)

	groupSiteOnenoteSectionGroupParentNotebookClient, err := groupsiteonenotesectiongroupparentnotebook.NewGroupSiteOnenoteSectionGroupParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupParentNotebookClient.Client)

	groupSiteOnenoteSectionGroupParentSectionGroupClient, err := groupsiteonenotesectiongroupparentsectiongroup.NewGroupSiteOnenoteSectionGroupParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupParentSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupParentSectionGroupClient.Client)

	groupSiteOnenoteSectionGroupSectionClient, err := groupsiteonenotesectiongroupsection.NewGroupSiteOnenoteSectionGroupSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionClient.Client)

	groupSiteOnenoteSectionGroupSectionGroupClient, err := groupsiteonenotesectiongroupsectiongroup.NewGroupSiteOnenoteSectionGroupSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionGroupClient.Client)

	groupSiteOnenoteSectionGroupSectionPageClient, err := groupsiteonenotesectiongroupsectionpage.NewGroupSiteOnenoteSectionGroupSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSectionPage client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionPageClient.Client)

	groupSiteOnenoteSectionGroupSectionPageContentClient, err := groupsiteonenotesectiongroupsectionpagecontent.NewGroupSiteOnenoteSectionGroupSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSectionPageContent client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionPageContentClient.Client)

	groupSiteOnenoteSectionGroupSectionPageParentNotebookClient, err := groupsiteonenotesectiongroupsectionpageparentnotebook.NewGroupSiteOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionPageParentNotebookClient.Client)

	groupSiteOnenoteSectionGroupSectionPageParentSectionClient, err := groupsiteonenotesectiongroupsectionpageparentsection.NewGroupSiteOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionPageParentSectionClient.Client)

	groupSiteOnenoteSectionGroupSectionParentNotebookClient, err := groupsiteonenotesectiongroupsectionparentnotebook.NewGroupSiteOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionParentNotebookClient.Client)

	groupSiteOnenoteSectionGroupSectionParentSectionGroupClient, err := groupsiteonenotesectiongroupsectionparentsectiongroup.NewGroupSiteOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionGroupSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionGroupSectionParentSectionGroupClient.Client)

	groupSiteOnenoteSectionPageClient, err := groupsiteonenotesectionpage.NewGroupSiteOnenoteSectionPageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionPage client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionPageClient.Client)

	groupSiteOnenoteSectionPageContentClient, err := groupsiteonenotesectionpagecontent.NewGroupSiteOnenoteSectionPageContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionPageContent client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionPageContentClient.Client)

	groupSiteOnenoteSectionPageParentNotebookClient, err := groupsiteonenotesectionpageparentnotebook.NewGroupSiteOnenoteSectionPageParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionPageParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionPageParentNotebookClient.Client)

	groupSiteOnenoteSectionPageParentSectionClient, err := groupsiteonenotesectionpageparentsection.NewGroupSiteOnenoteSectionPageParentSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionPageParentSection client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionPageParentSectionClient.Client)

	groupSiteOnenoteSectionParentNotebookClient, err := groupsiteonenotesectionparentnotebook.NewGroupSiteOnenoteSectionParentNotebookClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionParentNotebook client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionParentNotebookClient.Client)

	groupSiteOnenoteSectionParentSectionGroupClient, err := groupsiteonenotesectionparentsectiongroup.NewGroupSiteOnenoteSectionParentSectionGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOnenoteSectionParentSectionGroup client: %+v", err)
	}
	configureFunc(groupSiteOnenoteSectionParentSectionGroupClient.Client)

	groupSiteOperationClient, err := groupsiteoperation.NewGroupSiteOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteOperation client: %+v", err)
	}
	configureFunc(groupSiteOperationClient.Client)

	groupSitePageClient, err := groupsitepage.NewGroupSitePageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSitePage client: %+v", err)
	}
	configureFunc(groupSitePageClient.Client)

	groupSitePageCreatedByUserClient, err := groupsitepagecreatedbyuser.NewGroupSitePageCreatedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSitePageCreatedByUser client: %+v", err)
	}
	configureFunc(groupSitePageCreatedByUserClient.Client)

	groupSitePageCreatedByUserMailboxSettingClient, err := groupsitepagecreatedbyusermailboxsetting.NewGroupSitePageCreatedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSitePageCreatedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSitePageCreatedByUserMailboxSettingClient.Client)

	groupSitePageLastModifiedByUserClient, err := groupsitepagelastmodifiedbyuser.NewGroupSitePageLastModifiedByUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSitePageLastModifiedByUser client: %+v", err)
	}
	configureFunc(groupSitePageLastModifiedByUserClient.Client)

	groupSitePageLastModifiedByUserMailboxSettingClient, err := groupsitepagelastmodifiedbyusermailboxsetting.NewGroupSitePageLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSitePageLastModifiedByUserMailboxSetting client: %+v", err)
	}
	configureFunc(groupSitePageLastModifiedByUserMailboxSettingClient.Client)

	groupSitePermissionClient, err := groupsitepermission.NewGroupSitePermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSitePermission client: %+v", err)
	}
	configureFunc(groupSitePermissionClient.Client)

	groupSiteSiteClient, err := groupsitesite.NewGroupSiteSiteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteSite client: %+v", err)
	}
	configureFunc(groupSiteSiteClient.Client)

	groupSiteTermStoreClient, err := groupsitetermstore.NewGroupSiteTermStoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStore client: %+v", err)
	}
	configureFunc(groupSiteTermStoreClient.Client)

	groupSiteTermStoreGroupClient, err := groupsitetermstoregroup.NewGroupSiteTermStoreGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroup client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupClient.Client)

	groupSiteTermStoreGroupSetChildrenChildrenClient, err := groupsitetermstoregroupsetchildrenchildren.NewGroupSiteTermStoreGroupSetChildrenChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildrenChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenChildrenClient.Client)

	groupSiteTermStoreGroupSetChildrenChildrenRelationClient, err := groupsitetermstoregroupsetchildrenchildrenrelation.NewGroupSiteTermStoreGroupSetChildrenChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildrenChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenChildrenRelationClient.Client)

	groupSiteTermStoreGroupSetChildrenChildrenRelationFromTermClient, err := groupsitetermstoregroupsetchildrenchildrenrelationfromterm.NewGroupSiteTermStoreGroupSetChildrenChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildrenChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenChildrenRelationFromTermClient.Client)

	groupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient, err := groupsitetermstoregroupsetchildrenchildrenrelationtoterm.NewGroupSiteTermStoreGroupSetChildrenChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildrenChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient.Client)

	groupSiteTermStoreGroupSetChildrenClient, err := groupsitetermstoregroupsetchildren.NewGroupSiteTermStoreGroupSetChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenClient.Client)

	groupSiteTermStoreGroupSetChildrenRelationClient, err := groupsitetermstoregroupsetchildrenrelation.NewGroupSiteTermStoreGroupSetChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenRelationClient.Client)

	groupSiteTermStoreGroupSetChildrenRelationFromTermClient, err := groupsitetermstoregroupsetchildrenrelationfromterm.NewGroupSiteTermStoreGroupSetChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenRelationFromTermClient.Client)

	groupSiteTermStoreGroupSetChildrenRelationToTermClient, err := groupsitetermstoregroupsetchildrenrelationtoterm.NewGroupSiteTermStoreGroupSetChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetChildrenRelationToTermClient.Client)

	groupSiteTermStoreGroupSetClient, err := groupsitetermstoregroupset.NewGroupSiteTermStoreGroupSetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSet client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetClient.Client)

	groupSiteTermStoreGroupSetParentGroupClient, err := groupsitetermstoregroupsetparentgroup.NewGroupSiteTermStoreGroupSetParentGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetParentGroup client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetParentGroupClient.Client)

	groupSiteTermStoreGroupSetRelationClient, err := groupsitetermstoregroupsetrelation.NewGroupSiteTermStoreGroupSetRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetRelationClient.Client)

	groupSiteTermStoreGroupSetRelationFromTermClient, err := groupsitetermstoregroupsetrelationfromterm.NewGroupSiteTermStoreGroupSetRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetRelationFromTermClient.Client)

	groupSiteTermStoreGroupSetRelationToTermClient, err := groupsitetermstoregroupsetrelationtoterm.NewGroupSiteTermStoreGroupSetRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetRelationToTermClient.Client)

	groupSiteTermStoreGroupSetTermChildrenClient, err := groupsitetermstoregroupsettermchildren.NewGroupSiteTermStoreGroupSetTermChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTermChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermChildrenClient.Client)

	groupSiteTermStoreGroupSetTermChildrenRelationClient, err := groupsitetermstoregroupsettermchildrenrelation.NewGroupSiteTermStoreGroupSetTermChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTermChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermChildrenRelationClient.Client)

	groupSiteTermStoreGroupSetTermChildrenRelationFromTermClient, err := groupsitetermstoregroupsettermchildrenrelationfromterm.NewGroupSiteTermStoreGroupSetTermChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTermChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermChildrenRelationFromTermClient.Client)

	groupSiteTermStoreGroupSetTermChildrenRelationToTermClient, err := groupsitetermstoregroupsettermchildrenrelationtoterm.NewGroupSiteTermStoreGroupSetTermChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTermChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermChildrenRelationToTermClient.Client)

	groupSiteTermStoreGroupSetTermClient, err := groupsitetermstoregroupsetterm.NewGroupSiteTermStoreGroupSetTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermClient.Client)

	groupSiteTermStoreGroupSetTermRelationClient, err := groupsitetermstoregroupsettermrelation.NewGroupSiteTermStoreGroupSetTermRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTermRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermRelationClient.Client)

	groupSiteTermStoreGroupSetTermRelationFromTermClient, err := groupsitetermstoregroupsettermrelationfromterm.NewGroupSiteTermStoreGroupSetTermRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTermRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermRelationFromTermClient.Client)

	groupSiteTermStoreGroupSetTermRelationToTermClient, err := groupsitetermstoregroupsettermrelationtoterm.NewGroupSiteTermStoreGroupSetTermRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreGroupSetTermRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreGroupSetTermRelationToTermClient.Client)

	groupSiteTermStoreSetChildrenChildrenClient, err := groupsitetermstoresetchildrenchildren.NewGroupSiteTermStoreSetChildrenChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildrenChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenChildrenClient.Client)

	groupSiteTermStoreSetChildrenChildrenRelationClient, err := groupsitetermstoresetchildrenchildrenrelation.NewGroupSiteTermStoreSetChildrenChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildrenChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenChildrenRelationClient.Client)

	groupSiteTermStoreSetChildrenChildrenRelationFromTermClient, err := groupsitetermstoresetchildrenchildrenrelationfromterm.NewGroupSiteTermStoreSetChildrenChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildrenChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenChildrenRelationFromTermClient.Client)

	groupSiteTermStoreSetChildrenChildrenRelationToTermClient, err := groupsitetermstoresetchildrenchildrenrelationtoterm.NewGroupSiteTermStoreSetChildrenChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildrenChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenChildrenRelationToTermClient.Client)

	groupSiteTermStoreSetChildrenClient, err := groupsitetermstoresetchildren.NewGroupSiteTermStoreSetChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenClient.Client)

	groupSiteTermStoreSetChildrenRelationClient, err := groupsitetermstoresetchildrenrelation.NewGroupSiteTermStoreSetChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenRelationClient.Client)

	groupSiteTermStoreSetChildrenRelationFromTermClient, err := groupsitetermstoresetchildrenrelationfromterm.NewGroupSiteTermStoreSetChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenRelationFromTermClient.Client)

	groupSiteTermStoreSetChildrenRelationToTermClient, err := groupsitetermstoresetchildrenrelationtoterm.NewGroupSiteTermStoreSetChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetChildrenRelationToTermClient.Client)

	groupSiteTermStoreSetClient, err := groupsitetermstoreset.NewGroupSiteTermStoreSetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSet client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetClient.Client)

	groupSiteTermStoreSetParentGroupClient, err := groupsitetermstoresetparentgroup.NewGroupSiteTermStoreSetParentGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroup client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenChildrenClient, err := groupsitetermstoresetparentgroupsetchildrenchildren.NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildrenChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenChildrenClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient, err := groupsitetermstoresetparentgroupsetchildrenchildrenrelation.NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTermClient, err := groupsitetermstoresetparentgroupsetchildrenchildrenrelationfromterm.NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTermClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTermClient, err := groupsitetermstoresetparentgroupsetchildrenchildrenrelationtoterm.NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTermClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenClient, err := groupsitetermstoresetparentgroupsetchildren.NewGroupSiteTermStoreSetParentGroupSetChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenRelationClient, err := groupsitetermstoresetparentgroupsetchildrenrelation.NewGroupSiteTermStoreSetParentGroupSetChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenRelationClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient, err := groupsitetermstoresetparentgroupsetchildrenrelationfromterm.NewGroupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient.Client)

	groupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient, err := groupsitetermstoresetparentgroupsetchildrenrelationtoterm.NewGroupSiteTermStoreSetParentGroupSetChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient.Client)

	groupSiteTermStoreSetParentGroupSetClient, err := groupsitetermstoresetparentgroupset.NewGroupSiteTermStoreSetParentGroupSetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSet client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetClient.Client)

	groupSiteTermStoreSetParentGroupSetRelationClient, err := groupsitetermstoresetparentgroupsetrelation.NewGroupSiteTermStoreSetParentGroupSetRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetRelationClient.Client)

	groupSiteTermStoreSetParentGroupSetRelationFromTermClient, err := groupsitetermstoresetparentgroupsetrelationfromterm.NewGroupSiteTermStoreSetParentGroupSetRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetRelationFromTermClient.Client)

	groupSiteTermStoreSetParentGroupSetRelationToTermClient, err := groupsitetermstoresetparentgroupsetrelationtoterm.NewGroupSiteTermStoreSetParentGroupSetRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetRelationToTermClient.Client)

	groupSiteTermStoreSetParentGroupSetTermChildrenClient, err := groupsitetermstoresetparentgroupsettermchildren.NewGroupSiteTermStoreSetParentGroupSetTermChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTermChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermChildrenClient.Client)

	groupSiteTermStoreSetParentGroupSetTermChildrenRelationClient, err := groupsitetermstoresetparentgroupsettermchildrenrelation.NewGroupSiteTermStoreSetParentGroupSetTermChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTermChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermChildrenRelationClient.Client)

	groupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient, err := groupsitetermstoresetparentgroupsettermchildrenrelationfromterm.NewGroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient.Client)

	groupSiteTermStoreSetParentGroupSetTermChildrenRelationToTermClient, err := groupsitetermstoresetparentgroupsettermchildrenrelationtoterm.NewGroupSiteTermStoreSetParentGroupSetTermChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTermChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermChildrenRelationToTermClient.Client)

	groupSiteTermStoreSetParentGroupSetTermClient, err := groupsitetermstoresetparentgroupsetterm.NewGroupSiteTermStoreSetParentGroupSetTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermClient.Client)

	groupSiteTermStoreSetParentGroupSetTermRelationClient, err := groupsitetermstoresetparentgroupsettermrelation.NewGroupSiteTermStoreSetParentGroupSetTermRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTermRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermRelationClient.Client)

	groupSiteTermStoreSetParentGroupSetTermRelationFromTermClient, err := groupsitetermstoresetparentgroupsettermrelationfromterm.NewGroupSiteTermStoreSetParentGroupSetTermRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTermRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermRelationFromTermClient.Client)

	groupSiteTermStoreSetParentGroupSetTermRelationToTermClient, err := groupsitetermstoresetparentgroupsettermrelationtoterm.NewGroupSiteTermStoreSetParentGroupSetTermRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetParentGroupSetTermRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetParentGroupSetTermRelationToTermClient.Client)

	groupSiteTermStoreSetRelationClient, err := groupsitetermstoresetrelation.NewGroupSiteTermStoreSetRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetRelationClient.Client)

	groupSiteTermStoreSetRelationFromTermClient, err := groupsitetermstoresetrelationfromterm.NewGroupSiteTermStoreSetRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetRelationFromTermClient.Client)

	groupSiteTermStoreSetRelationToTermClient, err := groupsitetermstoresetrelationtoterm.NewGroupSiteTermStoreSetRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetRelationToTermClient.Client)

	groupSiteTermStoreSetTermChildrenClient, err := groupsitetermstoresettermchildren.NewGroupSiteTermStoreSetTermChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTermChildren client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermChildrenClient.Client)

	groupSiteTermStoreSetTermChildrenRelationClient, err := groupsitetermstoresettermchildrenrelation.NewGroupSiteTermStoreSetTermChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTermChildrenRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermChildrenRelationClient.Client)

	groupSiteTermStoreSetTermChildrenRelationFromTermClient, err := groupsitetermstoresettermchildrenrelationfromterm.NewGroupSiteTermStoreSetTermChildrenRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTermChildrenRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermChildrenRelationFromTermClient.Client)

	groupSiteTermStoreSetTermChildrenRelationToTermClient, err := groupsitetermstoresettermchildrenrelationtoterm.NewGroupSiteTermStoreSetTermChildrenRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTermChildrenRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermChildrenRelationToTermClient.Client)

	groupSiteTermStoreSetTermClient, err := groupsitetermstoresetterm.NewGroupSiteTermStoreSetTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermClient.Client)

	groupSiteTermStoreSetTermRelationClient, err := groupsitetermstoresettermrelation.NewGroupSiteTermStoreSetTermRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTermRelation client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermRelationClient.Client)

	groupSiteTermStoreSetTermRelationFromTermClient, err := groupsitetermstoresettermrelationfromterm.NewGroupSiteTermStoreSetTermRelationFromTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTermRelationFromTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermRelationFromTermClient.Client)

	groupSiteTermStoreSetTermRelationToTermClient, err := groupsitetermstoresettermrelationtoterm.NewGroupSiteTermStoreSetTermRelationToTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupSiteTermStoreSetTermRelationToTerm client: %+v", err)
	}
	configureFunc(groupSiteTermStoreSetTermRelationToTermClient.Client)

	groupTeamAllChannelClient, err := groupteamallchannel.NewGroupTeamAllChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamAllChannel client: %+v", err)
	}
	configureFunc(groupTeamAllChannelClient.Client)

	groupTeamChannelClient, err := groupteamchannel.NewGroupTeamChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannel client: %+v", err)
	}
	configureFunc(groupTeamChannelClient.Client)

	groupTeamChannelFilesFolderClient, err := groupteamchannelfilesfolder.NewGroupTeamChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelFilesFolder client: %+v", err)
	}
	configureFunc(groupTeamChannelFilesFolderClient.Client)

	groupTeamChannelFilesFolderContentClient, err := groupteamchannelfilesfoldercontent.NewGroupTeamChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(groupTeamChannelFilesFolderContentClient.Client)

	groupTeamChannelMemberClient, err := groupteamchannelmember.NewGroupTeamChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelMember client: %+v", err)
	}
	configureFunc(groupTeamChannelMemberClient.Client)

	groupTeamChannelMessageClient, err := groupteamchannelmessage.NewGroupTeamChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelMessage client: %+v", err)
	}
	configureFunc(groupTeamChannelMessageClient.Client)

	groupTeamChannelMessageHostedContentClient, err := groupteamchannelmessagehostedcontent.NewGroupTeamChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(groupTeamChannelMessageHostedContentClient.Client)

	groupTeamChannelMessageReplyClient, err := groupteamchannelmessagereply.NewGroupTeamChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelMessageReply client: %+v", err)
	}
	configureFunc(groupTeamChannelMessageReplyClient.Client)

	groupTeamChannelMessageReplyHostedContentClient, err := groupteamchannelmessagereplyhostedcontent.NewGroupTeamChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(groupTeamChannelMessageReplyHostedContentClient.Client)

	groupTeamChannelSharedWithTeamAllowedMemberClient, err := groupteamchannelsharedwithteamallowedmember.NewGroupTeamChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(groupTeamChannelSharedWithTeamAllowedMemberClient.Client)

	groupTeamChannelSharedWithTeamClient, err := groupteamchannelsharedwithteam.NewGroupTeamChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(groupTeamChannelSharedWithTeamClient.Client)

	groupTeamChannelSharedWithTeamTeamClient, err := groupteamchannelsharedwithteamteam.NewGroupTeamChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(groupTeamChannelSharedWithTeamTeamClient.Client)

	groupTeamChannelTabClient, err := groupteamchanneltab.NewGroupTeamChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelTab client: %+v", err)
	}
	configureFunc(groupTeamChannelTabClient.Client)

	groupTeamChannelTabTeamsAppClient, err := groupteamchanneltabteamsapp.NewGroupTeamChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(groupTeamChannelTabTeamsAppClient.Client)

	groupTeamClient, err := groupteam.NewGroupTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeam client: %+v", err)
	}
	configureFunc(groupTeamClient.Client)

	groupTeamGroupClient, err := groupteamgroup.NewGroupTeamGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamGroup client: %+v", err)
	}
	configureFunc(groupTeamGroupClient.Client)

	groupTeamIncomingChannelClient, err := groupteamincomingchannel.NewGroupTeamIncomingChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamIncomingChannel client: %+v", err)
	}
	configureFunc(groupTeamIncomingChannelClient.Client)

	groupTeamInstalledAppClient, err := groupteaminstalledapp.NewGroupTeamInstalledAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamInstalledApp client: %+v", err)
	}
	configureFunc(groupTeamInstalledAppClient.Client)

	groupTeamInstalledAppTeamsAppClient, err := groupteaminstalledappteamsapp.NewGroupTeamInstalledAppTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamInstalledAppTeamsApp client: %+v", err)
	}
	configureFunc(groupTeamInstalledAppTeamsAppClient.Client)

	groupTeamInstalledAppTeamsAppDefinitionClient, err := groupteaminstalledappteamsappdefinition.NewGroupTeamInstalledAppTeamsAppDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamInstalledAppTeamsAppDefinition client: %+v", err)
	}
	configureFunc(groupTeamInstalledAppTeamsAppDefinitionClient.Client)

	groupTeamMemberClient, err := groupteammember.NewGroupTeamMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamMember client: %+v", err)
	}
	configureFunc(groupTeamMemberClient.Client)

	groupTeamOperationClient, err := groupteamoperation.NewGroupTeamOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamOperation client: %+v", err)
	}
	configureFunc(groupTeamOperationClient.Client)

	groupTeamOwnerClient, err := groupteamowner.NewGroupTeamOwnerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamOwner client: %+v", err)
	}
	configureFunc(groupTeamOwnerClient.Client)

	groupTeamOwnerMailboxSettingClient, err := groupteamownermailboxsetting.NewGroupTeamOwnerMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamOwnerMailboxSetting client: %+v", err)
	}
	configureFunc(groupTeamOwnerMailboxSettingClient.Client)

	groupTeamPermissionGrantClient, err := groupteampermissiongrant.NewGroupTeamPermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPermissionGrant client: %+v", err)
	}
	configureFunc(groupTeamPermissionGrantClient.Client)

	groupTeamPhotoClient, err := groupteamphoto.NewGroupTeamPhotoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPhoto client: %+v", err)
	}
	configureFunc(groupTeamPhotoClient.Client)

	groupTeamPrimaryChannelClient, err := groupteamprimarychannel.NewGroupTeamPrimaryChannelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannel client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelClient.Client)

	groupTeamPrimaryChannelFilesFolderClient, err := groupteamprimarychannelfilesfolder.NewGroupTeamPrimaryChannelFilesFolderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelFilesFolder client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelFilesFolderClient.Client)

	groupTeamPrimaryChannelFilesFolderContentClient, err := groupteamprimarychannelfilesfoldercontent.NewGroupTeamPrimaryChannelFilesFolderContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelFilesFolderContent client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelFilesFolderContentClient.Client)

	groupTeamPrimaryChannelMemberClient, err := groupteamprimarychannelmember.NewGroupTeamPrimaryChannelMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelMember client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelMemberClient.Client)

	groupTeamPrimaryChannelMessageClient, err := groupteamprimarychannelmessage.NewGroupTeamPrimaryChannelMessageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelMessage client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelMessageClient.Client)

	groupTeamPrimaryChannelMessageHostedContentClient, err := groupteamprimarychannelmessagehostedcontent.NewGroupTeamPrimaryChannelMessageHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelMessageHostedContent client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelMessageHostedContentClient.Client)

	groupTeamPrimaryChannelMessageReplyClient, err := groupteamprimarychannelmessagereply.NewGroupTeamPrimaryChannelMessageReplyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelMessageReply client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelMessageReplyClient.Client)

	groupTeamPrimaryChannelMessageReplyHostedContentClient, err := groupteamprimarychannelmessagereplyhostedcontent.NewGroupTeamPrimaryChannelMessageReplyHostedContentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelMessageReplyHostedContent client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelMessageReplyHostedContentClient.Client)

	groupTeamPrimaryChannelSharedWithTeamAllowedMemberClient, err := groupteamprimarychannelsharedwithteamallowedmember.NewGroupTeamPrimaryChannelSharedWithTeamAllowedMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelSharedWithTeamAllowedMember client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelSharedWithTeamAllowedMemberClient.Client)

	groupTeamPrimaryChannelSharedWithTeamClient, err := groupteamprimarychannelsharedwithteam.NewGroupTeamPrimaryChannelSharedWithTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelSharedWithTeam client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelSharedWithTeamClient.Client)

	groupTeamPrimaryChannelSharedWithTeamTeamClient, err := groupteamprimarychannelsharedwithteamteam.NewGroupTeamPrimaryChannelSharedWithTeamTeamClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelSharedWithTeamTeam client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelSharedWithTeamTeamClient.Client)

	groupTeamPrimaryChannelTabClient, err := groupteamprimarychanneltab.NewGroupTeamPrimaryChannelTabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelTab client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelTabClient.Client)

	groupTeamPrimaryChannelTabTeamsAppClient, err := groupteamprimarychanneltabteamsapp.NewGroupTeamPrimaryChannelTabTeamsAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamPrimaryChannelTabTeamsApp client: %+v", err)
	}
	configureFunc(groupTeamPrimaryChannelTabTeamsAppClient.Client)

	groupTeamScheduleClient, err := groupteamschedule.NewGroupTeamScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamSchedule client: %+v", err)
	}
	configureFunc(groupTeamScheduleClient.Client)

	groupTeamScheduleOfferShiftRequestClient, err := groupteamscheduleoffershiftrequest.NewGroupTeamScheduleOfferShiftRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleOfferShiftRequest client: %+v", err)
	}
	configureFunc(groupTeamScheduleOfferShiftRequestClient.Client)

	groupTeamScheduleOpenShiftChangeRequestClient, err := groupteamscheduleopenshiftchangerequest.NewGroupTeamScheduleOpenShiftChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleOpenShiftChangeRequest client: %+v", err)
	}
	configureFunc(groupTeamScheduleOpenShiftChangeRequestClient.Client)

	groupTeamScheduleOpenShiftClient, err := groupteamscheduleopenshift.NewGroupTeamScheduleOpenShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleOpenShift client: %+v", err)
	}
	configureFunc(groupTeamScheduleOpenShiftClient.Client)

	groupTeamScheduleSchedulingGroupClient, err := groupteamscheduleschedulinggroup.NewGroupTeamScheduleSchedulingGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleSchedulingGroup client: %+v", err)
	}
	configureFunc(groupTeamScheduleSchedulingGroupClient.Client)

	groupTeamScheduleShiftClient, err := groupteamscheduleshift.NewGroupTeamScheduleShiftClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleShift client: %+v", err)
	}
	configureFunc(groupTeamScheduleShiftClient.Client)

	groupTeamScheduleSwapShiftsChangeRequestClient, err := groupteamscheduleswapshiftschangerequest.NewGroupTeamScheduleSwapShiftsChangeRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleSwapShiftsChangeRequest client: %+v", err)
	}
	configureFunc(groupTeamScheduleSwapShiftsChangeRequestClient.Client)

	groupTeamScheduleTimeCardClient, err := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleTimeCard client: %+v", err)
	}
	configureFunc(groupTeamScheduleTimeCardClient.Client)

	groupTeamScheduleTimeOffReasonClient, err := groupteamscheduletimeoffreason.NewGroupTeamScheduleTimeOffReasonClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleTimeOffReason client: %+v", err)
	}
	configureFunc(groupTeamScheduleTimeOffReasonClient.Client)

	groupTeamScheduleTimeOffRequestClient, err := groupteamscheduletimeoffrequest.NewGroupTeamScheduleTimeOffRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleTimeOffRequest client: %+v", err)
	}
	configureFunc(groupTeamScheduleTimeOffRequestClient.Client)

	groupTeamScheduleTimesOffClient, err := groupteamscheduletimesoff.NewGroupTeamScheduleTimesOffClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamScheduleTimesOff client: %+v", err)
	}
	configureFunc(groupTeamScheduleTimesOffClient.Client)

	groupTeamTagClient, err := groupteamtag.NewGroupTeamTagClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamTag client: %+v", err)
	}
	configureFunc(groupTeamTagClient.Client)

	groupTeamTagMemberClient, err := groupteamtagmember.NewGroupTeamTagMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamTagMember client: %+v", err)
	}
	configureFunc(groupTeamTagMemberClient.Client)

	groupTeamTemplateClient, err := groupteamtemplate.NewGroupTeamTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamTemplate client: %+v", err)
	}
	configureFunc(groupTeamTemplateClient.Client)

	groupTeamTemplateDefinitionClient, err := groupteamtemplatedefinition.NewGroupTeamTemplateDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTeamTemplateDefinition client: %+v", err)
	}
	configureFunc(groupTeamTemplateDefinitionClient.Client)

	groupThreadClient, err := groupthread.NewGroupThreadClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThread client: %+v", err)
	}
	configureFunc(groupThreadClient.Client)

	groupThreadPostAttachmentClient, err := groupthreadpostattachment.NewGroupThreadPostAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPostAttachment client: %+v", err)
	}
	configureFunc(groupThreadPostAttachmentClient.Client)

	groupThreadPostClient, err := groupthreadpost.NewGroupThreadPostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPost client: %+v", err)
	}
	configureFunc(groupThreadPostClient.Client)

	groupThreadPostExtensionClient, err := groupthreadpostextension.NewGroupThreadPostExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPostExtension client: %+v", err)
	}
	configureFunc(groupThreadPostExtensionClient.Client)

	groupThreadPostInReplyToAttachmentClient, err := groupthreadpostinreplytoattachment.NewGroupThreadPostInReplyToAttachmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPostInReplyToAttachment client: %+v", err)
	}
	configureFunc(groupThreadPostInReplyToAttachmentClient.Client)

	groupThreadPostInReplyToClient, err := groupthreadpostinreplyto.NewGroupThreadPostInReplyToClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPostInReplyTo client: %+v", err)
	}
	configureFunc(groupThreadPostInReplyToClient.Client)

	groupThreadPostInReplyToExtensionClient, err := groupthreadpostinreplytoextension.NewGroupThreadPostInReplyToExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPostInReplyToExtension client: %+v", err)
	}
	configureFunc(groupThreadPostInReplyToExtensionClient.Client)

	groupThreadPostInReplyToMentionClient, err := groupthreadpostinreplytomention.NewGroupThreadPostInReplyToMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPostInReplyToMention client: %+v", err)
	}
	configureFunc(groupThreadPostInReplyToMentionClient.Client)

	groupThreadPostMentionClient, err := groupthreadpostmention.NewGroupThreadPostMentionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupThreadPostMention client: %+v", err)
	}
	configureFunc(groupThreadPostMentionClient.Client)

	groupTransitiveMemberClient, err := grouptransitivemember.NewGroupTransitiveMemberClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTransitiveMember client: %+v", err)
	}
	configureFunc(groupTransitiveMemberClient.Client)

	groupTransitiveMemberOfClient, err := grouptransitivememberof.NewGroupTransitiveMemberOfClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupTransitiveMemberOf client: %+v", err)
	}
	configureFunc(groupTransitiveMemberOfClient.Client)

	setGroupSiteTermStoreGroupSetChildrenChildrenClient, err := setgroupsitetermstoregroupsetchildrenchildren.NewSetGroupSiteTermStoreGroupSetChildrenChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetChildrenChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetChildrenChildrenClient.Client)

	setGroupSiteTermStoreGroupSetChildrenChildrenRelationClient, err := setgroupsitetermstoregroupsetchildrenchildrenrelation.NewSetGroupSiteTermStoreGroupSetChildrenChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetChildrenChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetChildrenChildrenRelationClient.Client)

	setGroupSiteTermStoreGroupSetChildrenClient, err := setgroupsitetermstoregroupsetchildren.NewSetGroupSiteTermStoreGroupSetChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetChildrenClient.Client)

	setGroupSiteTermStoreGroupSetChildrenRelationClient, err := setgroupsitetermstoregroupsetchildrenrelation.NewSetGroupSiteTermStoreGroupSetChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetChildrenRelationClient.Client)

	setGroupSiteTermStoreGroupSetRelationClient, err := setgroupsitetermstoregroupsetrelation.NewSetGroupSiteTermStoreGroupSetRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetRelationClient.Client)

	setGroupSiteTermStoreGroupSetTermChildrenClient, err := setgroupsitetermstoregroupsettermchildren.NewSetGroupSiteTermStoreGroupSetTermChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetTermChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetTermChildrenClient.Client)

	setGroupSiteTermStoreGroupSetTermChildrenRelationClient, err := setgroupsitetermstoregroupsettermchildrenrelation.NewSetGroupSiteTermStoreGroupSetTermChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetTermChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetTermChildrenRelationClient.Client)

	setGroupSiteTermStoreGroupSetTermClient, err := setgroupsitetermstoregroupsetterm.NewSetGroupSiteTermStoreGroupSetTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetTerm client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetTermClient.Client)

	setGroupSiteTermStoreGroupSetTermRelationClient, err := setgroupsitetermstoregroupsettermrelation.NewSetGroupSiteTermStoreGroupSetTermRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreGroupSetTermRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreGroupSetTermRelationClient.Client)

	setGroupSiteTermStoreSetChildrenChildrenClient, err := setgroupsitetermstoresetchildrenchildren.NewSetGroupSiteTermStoreSetChildrenChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetChildrenChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetChildrenChildrenClient.Client)

	setGroupSiteTermStoreSetChildrenChildrenRelationClient, err := setgroupsitetermstoresetchildrenchildrenrelation.NewSetGroupSiteTermStoreSetChildrenChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetChildrenChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetChildrenChildrenRelationClient.Client)

	setGroupSiteTermStoreSetChildrenClient, err := setgroupsitetermstoresetchildren.NewSetGroupSiteTermStoreSetChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetChildrenClient.Client)

	setGroupSiteTermStoreSetChildrenRelationClient, err := setgroupsitetermstoresetchildrenrelation.NewSetGroupSiteTermStoreSetChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetChildrenRelationClient.Client)

	setGroupSiteTermStoreSetParentGroupSetChildrenChildrenClient, err := setgroupsitetermstoresetparentgroupsetchildrenchildren.NewSetGroupSiteTermStoreSetParentGroupSetChildrenChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetChildrenChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetChildrenChildrenClient.Client)

	setGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient, err := setgroupsitetermstoresetparentgroupsetchildrenchildrenrelation.NewSetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient.Client)

	setGroupSiteTermStoreSetParentGroupSetChildrenClient, err := setgroupsitetermstoresetparentgroupsetchildren.NewSetGroupSiteTermStoreSetParentGroupSetChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetChildrenClient.Client)

	setGroupSiteTermStoreSetParentGroupSetChildrenRelationClient, err := setgroupsitetermstoresetparentgroupsetchildrenrelation.NewSetGroupSiteTermStoreSetParentGroupSetChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetChildrenRelationClient.Client)

	setGroupSiteTermStoreSetParentGroupSetRelationClient, err := setgroupsitetermstoresetparentgroupsetrelation.NewSetGroupSiteTermStoreSetParentGroupSetRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetRelationClient.Client)

	setGroupSiteTermStoreSetParentGroupSetTermChildrenClient, err := setgroupsitetermstoresetparentgroupsettermchildren.NewSetGroupSiteTermStoreSetParentGroupSetTermChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetTermChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetTermChildrenClient.Client)

	setGroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient, err := setgroupsitetermstoresetparentgroupsettermchildrenrelation.NewSetGroupSiteTermStoreSetParentGroupSetTermChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetTermChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient.Client)

	setGroupSiteTermStoreSetParentGroupSetTermClient, err := setgroupsitetermstoresetparentgroupsetterm.NewSetGroupSiteTermStoreSetParentGroupSetTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetTerm client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetTermClient.Client)

	setGroupSiteTermStoreSetParentGroupSetTermRelationClient, err := setgroupsitetermstoresetparentgroupsettermrelation.NewSetGroupSiteTermStoreSetParentGroupSetTermRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetParentGroupSetTermRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetParentGroupSetTermRelationClient.Client)

	setGroupSiteTermStoreSetRelationClient, err := setgroupsitetermstoresetrelation.NewSetGroupSiteTermStoreSetRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetRelationClient.Client)

	setGroupSiteTermStoreSetTermChildrenClient, err := setgroupsitetermstoresettermchildren.NewSetGroupSiteTermStoreSetTermChildrenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetTermChildren client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetTermChildrenClient.Client)

	setGroupSiteTermStoreSetTermChildrenRelationClient, err := setgroupsitetermstoresettermchildrenrelation.NewSetGroupSiteTermStoreSetTermChildrenRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetTermChildrenRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetTermChildrenRelationClient.Client)

	setGroupSiteTermStoreSetTermClient, err := setgroupsitetermstoresetterm.NewSetGroupSiteTermStoreSetTermClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetTerm client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetTermClient.Client)

	setGroupSiteTermStoreSetTermRelationClient, err := setgroupsitetermstoresettermrelation.NewSetGroupSiteTermStoreSetTermRelationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SetGroupSiteTermStoreSetTermRelation client: %+v", err)
	}
	configureFunc(setGroupSiteTermStoreSetTermRelationClient.Client)

	return &Client{
		Group:                                                               groupClient,
		GroupAcceptedSender:                                                 groupAcceptedSenderClient,
		GroupAppRoleAssignment:                                              groupAppRoleAssignmentClient,
		GroupCalendar:                                                       groupCalendarClient,
		GroupCalendarCalendarPermission:                                     groupCalendarCalendarPermissionClient,
		GroupCalendarCalendarView:                                           groupCalendarCalendarViewClient,
		GroupCalendarCalendarViewAttachment:                                 groupCalendarCalendarViewAttachmentClient,
		GroupCalendarCalendarViewCalendar:                                   groupCalendarCalendarViewCalendarClient,
		GroupCalendarCalendarViewExceptionOccurrence:                        groupCalendarCalendarViewExceptionOccurrenceClient,
		GroupCalendarCalendarViewExceptionOccurrenceAttachment:              groupCalendarCalendarViewExceptionOccurrenceAttachmentClient,
		GroupCalendarCalendarViewExceptionOccurrenceCalendar:                groupCalendarCalendarViewExceptionOccurrenceCalendarClient,
		GroupCalendarCalendarViewExceptionOccurrenceExtension:               groupCalendarCalendarViewExceptionOccurrenceExtensionClient,
		GroupCalendarCalendarViewExceptionOccurrenceInstance:                groupCalendarCalendarViewExceptionOccurrenceInstanceClient,
		GroupCalendarCalendarViewExceptionOccurrenceInstanceAttachment:      groupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendar:        groupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient,
		GroupCalendarCalendarViewExceptionOccurrenceInstanceExtension:       groupCalendarCalendarViewExceptionOccurrenceInstanceExtensionClient,
		GroupCalendarCalendarViewExtension:                                  groupCalendarCalendarViewExtensionClient,
		GroupCalendarCalendarViewInstance:                                   groupCalendarCalendarViewInstanceClient,
		GroupCalendarCalendarViewInstanceAttachment:                         groupCalendarCalendarViewInstanceAttachmentClient,
		GroupCalendarCalendarViewInstanceCalendar:                           groupCalendarCalendarViewInstanceCalendarClient,
		GroupCalendarCalendarViewInstanceExceptionOccurrence:                groupCalendarCalendarViewInstanceExceptionOccurrenceClient,
		GroupCalendarCalendarViewInstanceExceptionOccurrenceAttachment:      groupCalendarCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendar:        groupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient,
		GroupCalendarCalendarViewInstanceExceptionOccurrenceExtension:       groupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient,
		GroupCalendarCalendarViewInstanceExtension:                          groupCalendarCalendarViewInstanceExtensionClient,
		GroupCalendarEvent:                                                  groupCalendarEventClient,
		GroupCalendarEventAttachment:                                        groupCalendarEventAttachmentClient,
		GroupCalendarEventCalendar:                                          groupCalendarEventCalendarClient,
		GroupCalendarEventExceptionOccurrence:                               groupCalendarEventExceptionOccurrenceClient,
		GroupCalendarEventExceptionOccurrenceAttachment:                     groupCalendarEventExceptionOccurrenceAttachmentClient,
		GroupCalendarEventExceptionOccurrenceCalendar:                       groupCalendarEventExceptionOccurrenceCalendarClient,
		GroupCalendarEventExceptionOccurrenceExtension:                      groupCalendarEventExceptionOccurrenceExtensionClient,
		GroupCalendarEventExceptionOccurrenceInstance:                       groupCalendarEventExceptionOccurrenceInstanceClient,
		GroupCalendarEventExceptionOccurrenceInstanceAttachment:             groupCalendarEventExceptionOccurrenceInstanceAttachmentClient,
		GroupCalendarEventExceptionOccurrenceInstanceCalendar:               groupCalendarEventExceptionOccurrenceInstanceCalendarClient,
		GroupCalendarEventExceptionOccurrenceInstanceExtension:              groupCalendarEventExceptionOccurrenceInstanceExtensionClient,
		GroupCalendarEventExtension:                                         groupCalendarEventExtensionClient,
		GroupCalendarEventInstance:                                          groupCalendarEventInstanceClient,
		GroupCalendarEventInstanceAttachment:                                groupCalendarEventInstanceAttachmentClient,
		GroupCalendarEventInstanceCalendar:                                  groupCalendarEventInstanceCalendarClient,
		GroupCalendarEventInstanceExceptionOccurrence:                       groupCalendarEventInstanceExceptionOccurrenceClient,
		GroupCalendarEventInstanceExceptionOccurrenceAttachment:             groupCalendarEventInstanceExceptionOccurrenceAttachmentClient,
		GroupCalendarEventInstanceExceptionOccurrenceCalendar:               groupCalendarEventInstanceExceptionOccurrenceCalendarClient,
		GroupCalendarEventInstanceExceptionOccurrenceExtension:              groupCalendarEventInstanceExceptionOccurrenceExtensionClient,
		GroupCalendarEventInstanceExtension:                                 groupCalendarEventInstanceExtensionClient,
		GroupCalendarView:                                                   groupCalendarViewClient,
		GroupCalendarViewAttachment:                                         groupCalendarViewAttachmentClient,
		GroupCalendarViewCalendar:                                           groupCalendarViewCalendarClient,
		GroupCalendarViewExceptionOccurrence:                                groupCalendarViewExceptionOccurrenceClient,
		GroupCalendarViewExceptionOccurrenceAttachment:                      groupCalendarViewExceptionOccurrenceAttachmentClient,
		GroupCalendarViewExceptionOccurrenceCalendar:                        groupCalendarViewExceptionOccurrenceCalendarClient,
		GroupCalendarViewExceptionOccurrenceExtension:                       groupCalendarViewExceptionOccurrenceExtensionClient,
		GroupCalendarViewExceptionOccurrenceInstance:                        groupCalendarViewExceptionOccurrenceInstanceClient,
		GroupCalendarViewExceptionOccurrenceInstanceAttachment:              groupCalendarViewExceptionOccurrenceInstanceAttachmentClient,
		GroupCalendarViewExceptionOccurrenceInstanceCalendar:                groupCalendarViewExceptionOccurrenceInstanceCalendarClient,
		GroupCalendarViewExceptionOccurrenceInstanceExtension:               groupCalendarViewExceptionOccurrenceInstanceExtensionClient,
		GroupCalendarViewExtension:                                          groupCalendarViewExtensionClient,
		GroupCalendarViewInstance:                                           groupCalendarViewInstanceClient,
		GroupCalendarViewInstanceAttachment:                                 groupCalendarViewInstanceAttachmentClient,
		GroupCalendarViewInstanceCalendar:                                   groupCalendarViewInstanceCalendarClient,
		GroupCalendarViewInstanceExceptionOccurrence:                        groupCalendarViewInstanceExceptionOccurrenceClient,
		GroupCalendarViewInstanceExceptionOccurrenceAttachment:              groupCalendarViewInstanceExceptionOccurrenceAttachmentClient,
		GroupCalendarViewInstanceExceptionOccurrenceCalendar:                groupCalendarViewInstanceExceptionOccurrenceCalendarClient,
		GroupCalendarViewInstanceExceptionOccurrenceExtension:               groupCalendarViewInstanceExceptionOccurrenceExtensionClient,
		GroupCalendarViewInstanceExtension:                                  groupCalendarViewInstanceExtensionClient,
		GroupConversation:                                                   groupConversationClient,
		GroupConversationThread:                                             groupConversationThreadClient,
		GroupConversationThreadPost:                                         groupConversationThreadPostClient,
		GroupConversationThreadPostAttachment:                               groupConversationThreadPostAttachmentClient,
		GroupConversationThreadPostExtension:                                groupConversationThreadPostExtensionClient,
		GroupConversationThreadPostInReplyTo:                                groupConversationThreadPostInReplyToClient,
		GroupConversationThreadPostInReplyToAttachment:                      groupConversationThreadPostInReplyToAttachmentClient,
		GroupConversationThreadPostInReplyToExtension:                       groupConversationThreadPostInReplyToExtensionClient,
		GroupConversationThreadPostInReplyToMention:                         groupConversationThreadPostInReplyToMentionClient,
		GroupConversationThreadPostMention:                                  groupConversationThreadPostMentionClient,
		GroupCreatedOnBehalfOf:                                              groupCreatedOnBehalfOfClient,
		GroupDrive:                                                          groupDriveClient,
		GroupEndpoint:                                                       groupEndpointClient,
		GroupEvent:                                                          groupEventClient,
		GroupEventAttachment:                                                groupEventAttachmentClient,
		GroupEventCalendar:                                                  groupEventCalendarClient,
		GroupEventExceptionOccurrence:                                       groupEventExceptionOccurrenceClient,
		GroupEventExceptionOccurrenceAttachment:                             groupEventExceptionOccurrenceAttachmentClient,
		GroupEventExceptionOccurrenceCalendar:                               groupEventExceptionOccurrenceCalendarClient,
		GroupEventExceptionOccurrenceExtension:                              groupEventExceptionOccurrenceExtensionClient,
		GroupEventExceptionOccurrenceInstance:                               groupEventExceptionOccurrenceInstanceClient,
		GroupEventExceptionOccurrenceInstanceAttachment:                     groupEventExceptionOccurrenceInstanceAttachmentClient,
		GroupEventExceptionOccurrenceInstanceCalendar:                       groupEventExceptionOccurrenceInstanceCalendarClient,
		GroupEventExceptionOccurrenceInstanceExtension:                      groupEventExceptionOccurrenceInstanceExtensionClient,
		GroupEventExtension:                                                 groupEventExtensionClient,
		GroupEventInstance:                                                  groupEventInstanceClient,
		GroupEventInstanceAttachment:                                        groupEventInstanceAttachmentClient,
		GroupEventInstanceCalendar:                                          groupEventInstanceCalendarClient,
		GroupEventInstanceExceptionOccurrence:                               groupEventInstanceExceptionOccurrenceClient,
		GroupEventInstanceExceptionOccurrenceAttachment:                     groupEventInstanceExceptionOccurrenceAttachmentClient,
		GroupEventInstanceExceptionOccurrenceCalendar:                       groupEventInstanceExceptionOccurrenceCalendarClient,
		GroupEventInstanceExceptionOccurrenceExtension:                      groupEventInstanceExceptionOccurrenceExtensionClient,
		GroupEventInstanceExtension:                                         groupEventInstanceExtensionClient,
		GroupExtension:                                                      groupExtensionClient,
		GroupGroupLifecyclePolicy:                                           groupGroupLifecyclePolicyClient,
		GroupMember:                                                         groupMemberClient,
		GroupMemberOf:                                                       groupMemberOfClient,
		GroupMembersWithLicenseError:                                        groupMembersWithLicenseErrorClient,
		GroupOnenote:                                                        groupOnenoteClient,
		GroupOnenoteNotebook:                                                groupOnenoteNotebookClient,
		GroupOnenoteNotebookSection:                                         groupOnenoteNotebookSectionClient,
		GroupOnenoteNotebookSectionGroup:                                    groupOnenoteNotebookSectionGroupClient,
		GroupOnenoteNotebookSectionGroupParentNotebook:                      groupOnenoteNotebookSectionGroupParentNotebookClient,
		GroupOnenoteNotebookSectionGroupParentSectionGroup:                  groupOnenoteNotebookSectionGroupParentSectionGroupClient,
		GroupOnenoteNotebookSectionGroupSection:                             groupOnenoteNotebookSectionGroupSectionClient,
		GroupOnenoteNotebookSectionGroupSectionGroup:                        groupOnenoteNotebookSectionGroupSectionGroupClient,
		GroupOnenoteNotebookSectionGroupSectionPage:                         groupOnenoteNotebookSectionGroupSectionPageClient,
		GroupOnenoteNotebookSectionGroupSectionPageContent:                  groupOnenoteNotebookSectionGroupSectionPageContentClient,
		GroupOnenoteNotebookSectionGroupSectionPageParentNotebook:           groupOnenoteNotebookSectionGroupSectionPageParentNotebookClient,
		GroupOnenoteNotebookSectionGroupSectionPageParentSection:            groupOnenoteNotebookSectionGroupSectionPageParentSectionClient,
		GroupOnenoteNotebookSectionGroupSectionParentNotebook:               groupOnenoteNotebookSectionGroupSectionParentNotebookClient,
		GroupOnenoteNotebookSectionGroupSectionParentSectionGroup:           groupOnenoteNotebookSectionGroupSectionParentSectionGroupClient,
		GroupOnenoteNotebookSectionPage:                                     groupOnenoteNotebookSectionPageClient,
		GroupOnenoteNotebookSectionPageContent:                              groupOnenoteNotebookSectionPageContentClient,
		GroupOnenoteNotebookSectionPageParentNotebook:                       groupOnenoteNotebookSectionPageParentNotebookClient,
		GroupOnenoteNotebookSectionPageParentSection:                        groupOnenoteNotebookSectionPageParentSectionClient,
		GroupOnenoteNotebookSectionParentNotebook:                           groupOnenoteNotebookSectionParentNotebookClient,
		GroupOnenoteNotebookSectionParentSectionGroup:                       groupOnenoteNotebookSectionParentSectionGroupClient,
		GroupOnenoteOperation:                                               groupOnenoteOperationClient,
		GroupOnenotePage:                                                    groupOnenotePageClient,
		GroupOnenotePageContent:                                             groupOnenotePageContentClient,
		GroupOnenotePageParentNotebook:                                      groupOnenotePageParentNotebookClient,
		GroupOnenotePageParentSection:                                       groupOnenotePageParentSectionClient,
		GroupOnenoteResource:                                                groupOnenoteResourceClient,
		GroupOnenoteResourceContent:                                         groupOnenoteResourceContentClient,
		GroupOnenoteSection:                                                 groupOnenoteSectionClient,
		GroupOnenoteSectionGroup:                                            groupOnenoteSectionGroupClient,
		GroupOnenoteSectionGroupParentNotebook:                              groupOnenoteSectionGroupParentNotebookClient,
		GroupOnenoteSectionGroupParentSectionGroup:                          groupOnenoteSectionGroupParentSectionGroupClient,
		GroupOnenoteSectionGroupSection:                                     groupOnenoteSectionGroupSectionClient,
		GroupOnenoteSectionGroupSectionGroup:                                groupOnenoteSectionGroupSectionGroupClient,
		GroupOnenoteSectionGroupSectionPage:                                 groupOnenoteSectionGroupSectionPageClient,
		GroupOnenoteSectionGroupSectionPageContent:                          groupOnenoteSectionGroupSectionPageContentClient,
		GroupOnenoteSectionGroupSectionPageParentNotebook:                   groupOnenoteSectionGroupSectionPageParentNotebookClient,
		GroupOnenoteSectionGroupSectionPageParentSection:                    groupOnenoteSectionGroupSectionPageParentSectionClient,
		GroupOnenoteSectionGroupSectionParentNotebook:                       groupOnenoteSectionGroupSectionParentNotebookClient,
		GroupOnenoteSectionGroupSectionParentSectionGroup:                   groupOnenoteSectionGroupSectionParentSectionGroupClient,
		GroupOnenoteSectionPage:                                             groupOnenoteSectionPageClient,
		GroupOnenoteSectionPageContent:                                      groupOnenoteSectionPageContentClient,
		GroupOnenoteSectionPageParentNotebook:                               groupOnenoteSectionPageParentNotebookClient,
		GroupOnenoteSectionPageParentSection:                                groupOnenoteSectionPageParentSectionClient,
		GroupOnenoteSectionParentNotebook:                                   groupOnenoteSectionParentNotebookClient,
		GroupOnenoteSectionParentSectionGroup:                               groupOnenoteSectionParentSectionGroupClient,
		GroupOwner:                                                          groupOwnerClient,
		GroupPermissionGrant:                                                groupPermissionGrantClient,
		GroupPhoto:                                                          groupPhotoClient,
		GroupPlanner:                                                        groupPlannerClient,
		GroupPlannerPlan:                                                    groupPlannerPlanClient,
		GroupPlannerPlanBucket:                                              groupPlannerPlanBucketClient,
		GroupPlannerPlanBucketTask:                                          groupPlannerPlanBucketTaskClient,
		GroupPlannerPlanBucketTaskAssignedToTaskBoardFormat:                 groupPlannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		GroupPlannerPlanBucketTaskBucketTaskBoardFormat:                     groupPlannerPlanBucketTaskBucketTaskBoardFormatClient,
		GroupPlannerPlanBucketTaskDetail:                                    groupPlannerPlanBucketTaskDetailClient,
		GroupPlannerPlanBucketTaskProgressTaskBoardFormat:                   groupPlannerPlanBucketTaskProgressTaskBoardFormatClient,
		GroupPlannerPlanDetail:                                              groupPlannerPlanDetailClient,
		GroupPlannerPlanTask:                                                groupPlannerPlanTaskClient,
		GroupPlannerPlanTaskAssignedToTaskBoardFormat:                       groupPlannerPlanTaskAssignedToTaskBoardFormatClient,
		GroupPlannerPlanTaskBucketTaskBoardFormat:                           groupPlannerPlanTaskBucketTaskBoardFormatClient,
		GroupPlannerPlanTaskDetail:                                          groupPlannerPlanTaskDetailClient,
		GroupPlannerPlanTaskProgressTaskBoardFormat:                         groupPlannerPlanTaskProgressTaskBoardFormatClient,
		GroupRejectedSender:                                                 groupRejectedSenderClient,
		GroupSetting:                                                        groupSettingClient,
		GroupSite:                                                           groupSiteClient,
		GroupSiteAnalytic:                                                   groupSiteAnalyticClient,
		GroupSiteAnalyticAllTime:                                            groupSiteAnalyticAllTimeClient,
		GroupSiteAnalyticItemActivityStat:                                   groupSiteAnalyticItemActivityStatClient,
		GroupSiteAnalyticItemActivityStatActivity:                           groupSiteAnalyticItemActivityStatActivityClient,
		GroupSiteAnalyticItemActivityStatActivityDriveItem:                  groupSiteAnalyticItemActivityStatActivityDriveItemClient,
		GroupSiteAnalyticItemActivityStatActivityDriveItemContent:           groupSiteAnalyticItemActivityStatActivityDriveItemContentClient,
		GroupSiteAnalyticLastSevenDay:                                       groupSiteAnalyticLastSevenDayClient,
		GroupSiteColumn:                                                     groupSiteColumnClient,
		GroupSiteColumnSourceColumn:                                         groupSiteColumnSourceColumnClient,
		GroupSiteContentType:                                                groupSiteContentTypeClient,
		GroupSiteContentTypeBase:                                            groupSiteContentTypeBaseClient,
		GroupSiteContentTypeBaseType:                                        groupSiteContentTypeBaseTypeClient,
		GroupSiteContentTypeColumn:                                          groupSiteContentTypeColumnClient,
		GroupSiteContentTypeColumnLink:                                      groupSiteContentTypeColumnLinkClient,
		GroupSiteContentTypeColumnPosition:                                  groupSiteContentTypeColumnPositionClient,
		GroupSiteContentTypeColumnSourceColumn:                              groupSiteContentTypeColumnSourceColumnClient,
		GroupSiteCreatedByUser:                                              groupSiteCreatedByUserClient,
		GroupSiteCreatedByUserMailboxSetting:                                groupSiteCreatedByUserMailboxSettingClient,
		GroupSiteDrive:                                                      groupSiteDriveClient,
		GroupSiteExternalColumn:                                             groupSiteExternalColumnClient,
		GroupSiteInformationProtection:                                      groupSiteInformationProtectionClient,
		GroupSiteInformationProtectionBitlocker:                             groupSiteInformationProtectionBitlockerClient,
		GroupSiteInformationProtectionBitlockerRecoveryKey:                  groupSiteInformationProtectionBitlockerRecoveryKeyClient,
		GroupSiteInformationProtectionDataLossPreventionPolicy:              groupSiteInformationProtectionDataLossPreventionPolicyClient,
		GroupSiteInformationProtectionPolicy:                                groupSiteInformationProtectionPolicyClient,
		GroupSiteInformationProtectionPolicyLabel:                           groupSiteInformationProtectionPolicyLabelClient,
		GroupSiteInformationProtectionSensitivityLabel:                      groupSiteInformationProtectionSensitivityLabelClient,
		GroupSiteInformationProtectionSensitivityLabelSublabel:              groupSiteInformationProtectionSensitivityLabelSublabelClient,
		GroupSiteInformationProtectionSensitivityPolicySetting:              groupSiteInformationProtectionSensitivityPolicySettingClient,
		GroupSiteInformationProtectionThreatAssessmentRequest:               groupSiteInformationProtectionThreatAssessmentRequestClient,
		GroupSiteInformationProtectionThreatAssessmentRequestResult:         groupSiteInformationProtectionThreatAssessmentRequestResultClient,
		GroupSiteItem:                                                       groupSiteItemClient,
		GroupSiteLastModifiedByUser:                                         groupSiteLastModifiedByUserClient,
		GroupSiteLastModifiedByUserMailboxSetting:                           groupSiteLastModifiedByUserMailboxSettingClient,
		GroupSiteList:                                                       groupSiteListClient,
		GroupSiteListActivity:                                               groupSiteListActivityClient,
		GroupSiteListColumn:                                                 groupSiteListColumnClient,
		GroupSiteListColumnSourceColumn:                                     groupSiteListColumnSourceColumnClient,
		GroupSiteListContentType:                                            groupSiteListContentTypeClient,
		GroupSiteListContentTypeBase:                                        groupSiteListContentTypeBaseClient,
		GroupSiteListContentTypeBaseType:                                    groupSiteListContentTypeBaseTypeClient,
		GroupSiteListContentTypeColumn:                                      groupSiteListContentTypeColumnClient,
		GroupSiteListContentTypeColumnLink:                                  groupSiteListContentTypeColumnLinkClient,
		GroupSiteListContentTypeColumnPosition:                              groupSiteListContentTypeColumnPositionClient,
		GroupSiteListContentTypeColumnSourceColumn:                          groupSiteListContentTypeColumnSourceColumnClient,
		GroupSiteListCreatedByUser:                                          groupSiteListCreatedByUserClient,
		GroupSiteListCreatedByUserMailboxSetting:                            groupSiteListCreatedByUserMailboxSettingClient,
		GroupSiteListDrive:                                                  groupSiteListDriveClient,
		GroupSiteListItem:                                                   groupSiteListItemClient,
		GroupSiteListItemActivity:                                           groupSiteListItemActivityClient,
		GroupSiteListItemActivityDriveItem:                                  groupSiteListItemActivityDriveItemClient,
		GroupSiteListItemActivityDriveItemContent:                           groupSiteListItemActivityDriveItemContentClient,
		GroupSiteListItemActivityListItem:                                   groupSiteListItemActivityListItemClient,
		GroupSiteListItemAnalytic:                                           groupSiteListItemAnalyticClient,
		GroupSiteListItemCreatedByUser:                                      groupSiteListItemCreatedByUserClient,
		GroupSiteListItemCreatedByUserMailboxSetting:                        groupSiteListItemCreatedByUserMailboxSettingClient,
		GroupSiteListItemDocumentSetVersion:                                 groupSiteListItemDocumentSetVersionClient,
		GroupSiteListItemDocumentSetVersionField:                            groupSiteListItemDocumentSetVersionFieldClient,
		GroupSiteListItemDriveItem:                                          groupSiteListItemDriveItemClient,
		GroupSiteListItemDriveItemContent:                                   groupSiteListItemDriveItemContentClient,
		GroupSiteListItemField:                                              groupSiteListItemFieldClient,
		GroupSiteListItemLastModifiedByUser:                                 groupSiteListItemLastModifiedByUserClient,
		GroupSiteListItemLastModifiedByUserMailboxSetting:                   groupSiteListItemLastModifiedByUserMailboxSettingClient,
		GroupSiteListItemVersion:                                            groupSiteListItemVersionClient,
		GroupSiteListItemVersionField:                                       groupSiteListItemVersionFieldClient,
		GroupSiteListLastModifiedByUser:                                     groupSiteListLastModifiedByUserClient,
		GroupSiteListLastModifiedByUserMailboxSetting:                       groupSiteListLastModifiedByUserMailboxSettingClient,
		GroupSiteListOperation:                                              groupSiteListOperationClient,
		GroupSiteListSubscription:                                           groupSiteListSubscriptionClient,
		GroupSiteOnenote:                                                    groupSiteOnenoteClient,
		GroupSiteOnenoteNotebook:                                            groupSiteOnenoteNotebookClient,
		GroupSiteOnenoteNotebookSection:                                     groupSiteOnenoteNotebookSectionClient,
		GroupSiteOnenoteNotebookSectionGroup:                                groupSiteOnenoteNotebookSectionGroupClient,
		GroupSiteOnenoteNotebookSectionGroupParentNotebook:                  groupSiteOnenoteNotebookSectionGroupParentNotebookClient,
		GroupSiteOnenoteNotebookSectionGroupParentSectionGroup:              groupSiteOnenoteNotebookSectionGroupParentSectionGroupClient,
		GroupSiteOnenoteNotebookSectionGroupSection:                         groupSiteOnenoteNotebookSectionGroupSectionClient,
		GroupSiteOnenoteNotebookSectionGroupSectionGroup:                    groupSiteOnenoteNotebookSectionGroupSectionGroupClient,
		GroupSiteOnenoteNotebookSectionGroupSectionPage:                     groupSiteOnenoteNotebookSectionGroupSectionPageClient,
		GroupSiteOnenoteNotebookSectionGroupSectionPageContent:              groupSiteOnenoteNotebookSectionGroupSectionPageContentClient,
		GroupSiteOnenoteNotebookSectionGroupSectionPageParentNotebook:       groupSiteOnenoteNotebookSectionGroupSectionPageParentNotebookClient,
		GroupSiteOnenoteNotebookSectionGroupSectionPageParentSection:        groupSiteOnenoteNotebookSectionGroupSectionPageParentSectionClient,
		GroupSiteOnenoteNotebookSectionGroupSectionParentNotebook:           groupSiteOnenoteNotebookSectionGroupSectionParentNotebookClient,
		GroupSiteOnenoteNotebookSectionGroupSectionParentSectionGroup:       groupSiteOnenoteNotebookSectionGroupSectionParentSectionGroupClient,
		GroupSiteOnenoteNotebookSectionPage:                                 groupSiteOnenoteNotebookSectionPageClient,
		GroupSiteOnenoteNotebookSectionPageContent:                          groupSiteOnenoteNotebookSectionPageContentClient,
		GroupSiteOnenoteNotebookSectionPageParentNotebook:                   groupSiteOnenoteNotebookSectionPageParentNotebookClient,
		GroupSiteOnenoteNotebookSectionPageParentSection:                    groupSiteOnenoteNotebookSectionPageParentSectionClient,
		GroupSiteOnenoteNotebookSectionParentNotebook:                       groupSiteOnenoteNotebookSectionParentNotebookClient,
		GroupSiteOnenoteNotebookSectionParentSectionGroup:                   groupSiteOnenoteNotebookSectionParentSectionGroupClient,
		GroupSiteOnenoteOperation:                                           groupSiteOnenoteOperationClient,
		GroupSiteOnenotePage:                                                groupSiteOnenotePageClient,
		GroupSiteOnenotePageContent:                                         groupSiteOnenotePageContentClient,
		GroupSiteOnenotePageParentNotebook:                                  groupSiteOnenotePageParentNotebookClient,
		GroupSiteOnenotePageParentSection:                                   groupSiteOnenotePageParentSectionClient,
		GroupSiteOnenoteResource:                                            groupSiteOnenoteResourceClient,
		GroupSiteOnenoteResourceContent:                                     groupSiteOnenoteResourceContentClient,
		GroupSiteOnenoteSection:                                             groupSiteOnenoteSectionClient,
		GroupSiteOnenoteSectionGroup:                                        groupSiteOnenoteSectionGroupClient,
		GroupSiteOnenoteSectionGroupParentNotebook:                          groupSiteOnenoteSectionGroupParentNotebookClient,
		GroupSiteOnenoteSectionGroupParentSectionGroup:                      groupSiteOnenoteSectionGroupParentSectionGroupClient,
		GroupSiteOnenoteSectionGroupSection:                                 groupSiteOnenoteSectionGroupSectionClient,
		GroupSiteOnenoteSectionGroupSectionGroup:                            groupSiteOnenoteSectionGroupSectionGroupClient,
		GroupSiteOnenoteSectionGroupSectionPage:                             groupSiteOnenoteSectionGroupSectionPageClient,
		GroupSiteOnenoteSectionGroupSectionPageContent:                      groupSiteOnenoteSectionGroupSectionPageContentClient,
		GroupSiteOnenoteSectionGroupSectionPageParentNotebook:               groupSiteOnenoteSectionGroupSectionPageParentNotebookClient,
		GroupSiteOnenoteSectionGroupSectionPageParentSection:                groupSiteOnenoteSectionGroupSectionPageParentSectionClient,
		GroupSiteOnenoteSectionGroupSectionParentNotebook:                   groupSiteOnenoteSectionGroupSectionParentNotebookClient,
		GroupSiteOnenoteSectionGroupSectionParentSectionGroup:               groupSiteOnenoteSectionGroupSectionParentSectionGroupClient,
		GroupSiteOnenoteSectionPage:                                         groupSiteOnenoteSectionPageClient,
		GroupSiteOnenoteSectionPageContent:                                  groupSiteOnenoteSectionPageContentClient,
		GroupSiteOnenoteSectionPageParentNotebook:                           groupSiteOnenoteSectionPageParentNotebookClient,
		GroupSiteOnenoteSectionPageParentSection:                            groupSiteOnenoteSectionPageParentSectionClient,
		GroupSiteOnenoteSectionParentNotebook:                               groupSiteOnenoteSectionParentNotebookClient,
		GroupSiteOnenoteSectionParentSectionGroup:                           groupSiteOnenoteSectionParentSectionGroupClient,
		GroupSiteOperation:                                                  groupSiteOperationClient,
		GroupSitePage:                                                       groupSitePageClient,
		GroupSitePageCreatedByUser:                                          groupSitePageCreatedByUserClient,
		GroupSitePageCreatedByUserMailboxSetting:                            groupSitePageCreatedByUserMailboxSettingClient,
		GroupSitePageLastModifiedByUser:                                     groupSitePageLastModifiedByUserClient,
		GroupSitePageLastModifiedByUserMailboxSetting:                       groupSitePageLastModifiedByUserMailboxSettingClient,
		GroupSitePermission:                                                 groupSitePermissionClient,
		GroupSiteSite:                                                       groupSiteSiteClient,
		GroupSiteTermStore:                                                  groupSiteTermStoreClient,
		GroupSiteTermStoreGroup:                                             groupSiteTermStoreGroupClient,
		GroupSiteTermStoreGroupSet:                                          groupSiteTermStoreGroupSetClient,
		GroupSiteTermStoreGroupSetChildren:                                  groupSiteTermStoreGroupSetChildrenClient,
		GroupSiteTermStoreGroupSetChildrenChildren:                          groupSiteTermStoreGroupSetChildrenChildrenClient,
		GroupSiteTermStoreGroupSetChildrenChildrenRelation:                  groupSiteTermStoreGroupSetChildrenChildrenRelationClient,
		GroupSiteTermStoreGroupSetChildrenChildrenRelationFromTerm:          groupSiteTermStoreGroupSetChildrenChildrenRelationFromTermClient,
		GroupSiteTermStoreGroupSetChildrenChildrenRelationToTerm:            groupSiteTermStoreGroupSetChildrenChildrenRelationToTermClient,
		GroupSiteTermStoreGroupSetChildrenRelation:                          groupSiteTermStoreGroupSetChildrenRelationClient,
		GroupSiteTermStoreGroupSetChildrenRelationFromTerm:                  groupSiteTermStoreGroupSetChildrenRelationFromTermClient,
		GroupSiteTermStoreGroupSetChildrenRelationToTerm:                    groupSiteTermStoreGroupSetChildrenRelationToTermClient,
		GroupSiteTermStoreGroupSetParentGroup:                               groupSiteTermStoreGroupSetParentGroupClient,
		GroupSiteTermStoreGroupSetRelation:                                  groupSiteTermStoreGroupSetRelationClient,
		GroupSiteTermStoreGroupSetRelationFromTerm:                          groupSiteTermStoreGroupSetRelationFromTermClient,
		GroupSiteTermStoreGroupSetRelationToTerm:                            groupSiteTermStoreGroupSetRelationToTermClient,
		GroupSiteTermStoreGroupSetTerm:                                      groupSiteTermStoreGroupSetTermClient,
		GroupSiteTermStoreGroupSetTermChildren:                              groupSiteTermStoreGroupSetTermChildrenClient,
		GroupSiteTermStoreGroupSetTermChildrenRelation:                      groupSiteTermStoreGroupSetTermChildrenRelationClient,
		GroupSiteTermStoreGroupSetTermChildrenRelationFromTerm:              groupSiteTermStoreGroupSetTermChildrenRelationFromTermClient,
		GroupSiteTermStoreGroupSetTermChildrenRelationToTerm:                groupSiteTermStoreGroupSetTermChildrenRelationToTermClient,
		GroupSiteTermStoreGroupSetTermRelation:                              groupSiteTermStoreGroupSetTermRelationClient,
		GroupSiteTermStoreGroupSetTermRelationFromTerm:                      groupSiteTermStoreGroupSetTermRelationFromTermClient,
		GroupSiteTermStoreGroupSetTermRelationToTerm:                        groupSiteTermStoreGroupSetTermRelationToTermClient,
		GroupSiteTermStoreSet:                                               groupSiteTermStoreSetClient,
		GroupSiteTermStoreSetChildren:                                       groupSiteTermStoreSetChildrenClient,
		GroupSiteTermStoreSetChildrenChildren:                               groupSiteTermStoreSetChildrenChildrenClient,
		GroupSiteTermStoreSetChildrenChildrenRelation:                       groupSiteTermStoreSetChildrenChildrenRelationClient,
		GroupSiteTermStoreSetChildrenChildrenRelationFromTerm:               groupSiteTermStoreSetChildrenChildrenRelationFromTermClient,
		GroupSiteTermStoreSetChildrenChildrenRelationToTerm:                 groupSiteTermStoreSetChildrenChildrenRelationToTermClient,
		GroupSiteTermStoreSetChildrenRelation:                               groupSiteTermStoreSetChildrenRelationClient,
		GroupSiteTermStoreSetChildrenRelationFromTerm:                       groupSiteTermStoreSetChildrenRelationFromTermClient,
		GroupSiteTermStoreSetChildrenRelationToTerm:                         groupSiteTermStoreSetChildrenRelationToTermClient,
		GroupSiteTermStoreSetParentGroup:                                    groupSiteTermStoreSetParentGroupClient,
		GroupSiteTermStoreSetParentGroupSet:                                 groupSiteTermStoreSetParentGroupSetClient,
		GroupSiteTermStoreSetParentGroupSetChildren:                         groupSiteTermStoreSetParentGroupSetChildrenClient,
		GroupSiteTermStoreSetParentGroupSetChildrenChildren:                 groupSiteTermStoreSetParentGroupSetChildrenChildrenClient,
		GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelation:         groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient,
		GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTerm: groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationFromTermClient,
		GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTerm:   groupSiteTermStoreSetParentGroupSetChildrenChildrenRelationToTermClient,
		GroupSiteTermStoreSetParentGroupSetChildrenRelation:                 groupSiteTermStoreSetParentGroupSetChildrenRelationClient,
		GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTerm:         groupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient,
		GroupSiteTermStoreSetParentGroupSetChildrenRelationToTerm:           groupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient,
		GroupSiteTermStoreSetParentGroupSetRelation:                         groupSiteTermStoreSetParentGroupSetRelationClient,
		GroupSiteTermStoreSetParentGroupSetRelationFromTerm:                 groupSiteTermStoreSetParentGroupSetRelationFromTermClient,
		GroupSiteTermStoreSetParentGroupSetRelationToTerm:                   groupSiteTermStoreSetParentGroupSetRelationToTermClient,
		GroupSiteTermStoreSetParentGroupSetTerm:                             groupSiteTermStoreSetParentGroupSetTermClient,
		GroupSiteTermStoreSetParentGroupSetTermChildren:                     groupSiteTermStoreSetParentGroupSetTermChildrenClient,
		GroupSiteTermStoreSetParentGroupSetTermChildrenRelation:             groupSiteTermStoreSetParentGroupSetTermChildrenRelationClient,
		GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTerm:     groupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient,
		GroupSiteTermStoreSetParentGroupSetTermChildrenRelationToTerm:       groupSiteTermStoreSetParentGroupSetTermChildrenRelationToTermClient,
		GroupSiteTermStoreSetParentGroupSetTermRelation:                     groupSiteTermStoreSetParentGroupSetTermRelationClient,
		GroupSiteTermStoreSetParentGroupSetTermRelationFromTerm:             groupSiteTermStoreSetParentGroupSetTermRelationFromTermClient,
		GroupSiteTermStoreSetParentGroupSetTermRelationToTerm:               groupSiteTermStoreSetParentGroupSetTermRelationToTermClient,
		GroupSiteTermStoreSetRelation:                                       groupSiteTermStoreSetRelationClient,
		GroupSiteTermStoreSetRelationFromTerm:                               groupSiteTermStoreSetRelationFromTermClient,
		GroupSiteTermStoreSetRelationToTerm:                                 groupSiteTermStoreSetRelationToTermClient,
		GroupSiteTermStoreSetTerm:                                           groupSiteTermStoreSetTermClient,
		GroupSiteTermStoreSetTermChildren:                                   groupSiteTermStoreSetTermChildrenClient,
		GroupSiteTermStoreSetTermChildrenRelation:                           groupSiteTermStoreSetTermChildrenRelationClient,
		GroupSiteTermStoreSetTermChildrenRelationFromTerm:                   groupSiteTermStoreSetTermChildrenRelationFromTermClient,
		GroupSiteTermStoreSetTermChildrenRelationToTerm:                     groupSiteTermStoreSetTermChildrenRelationToTermClient,
		GroupSiteTermStoreSetTermRelation:                                   groupSiteTermStoreSetTermRelationClient,
		GroupSiteTermStoreSetTermRelationFromTerm:                           groupSiteTermStoreSetTermRelationFromTermClient,
		GroupSiteTermStoreSetTermRelationToTerm:                             groupSiteTermStoreSetTermRelationToTermClient,
		GroupTeam:                                                           groupTeamClient,
		GroupTeamAllChannel:                                                 groupTeamAllChannelClient,
		GroupTeamChannel:                                                    groupTeamChannelClient,
		GroupTeamChannelFilesFolder:                                         groupTeamChannelFilesFolderClient,
		GroupTeamChannelFilesFolderContent:                                  groupTeamChannelFilesFolderContentClient,
		GroupTeamChannelMember:                                              groupTeamChannelMemberClient,
		GroupTeamChannelMessage:                                             groupTeamChannelMessageClient,
		GroupTeamChannelMessageHostedContent:                                groupTeamChannelMessageHostedContentClient,
		GroupTeamChannelMessageReply:                                        groupTeamChannelMessageReplyClient,
		GroupTeamChannelMessageReplyHostedContent:                           groupTeamChannelMessageReplyHostedContentClient,
		GroupTeamChannelSharedWithTeam:                                      groupTeamChannelSharedWithTeamClient,
		GroupTeamChannelSharedWithTeamAllowedMember:                         groupTeamChannelSharedWithTeamAllowedMemberClient,
		GroupTeamChannelSharedWithTeamTeam:                                  groupTeamChannelSharedWithTeamTeamClient,
		GroupTeamChannelTab:                                                 groupTeamChannelTabClient,
		GroupTeamChannelTabTeamsApp:                                         groupTeamChannelTabTeamsAppClient,
		GroupTeamGroup:                                                      groupTeamGroupClient,
		GroupTeamIncomingChannel:                                            groupTeamIncomingChannelClient,
		GroupTeamInstalledApp:                                               groupTeamInstalledAppClient,
		GroupTeamInstalledAppTeamsApp:                                       groupTeamInstalledAppTeamsAppClient,
		GroupTeamInstalledAppTeamsAppDefinition:                             groupTeamInstalledAppTeamsAppDefinitionClient,
		GroupTeamMember:                                                     groupTeamMemberClient,
		GroupTeamOperation:                                                  groupTeamOperationClient,
		GroupTeamOwner:                                                      groupTeamOwnerClient,
		GroupTeamOwnerMailboxSetting:                                        groupTeamOwnerMailboxSettingClient,
		GroupTeamPermissionGrant:                                            groupTeamPermissionGrantClient,
		GroupTeamPhoto:                                                      groupTeamPhotoClient,
		GroupTeamPrimaryChannel:                                             groupTeamPrimaryChannelClient,
		GroupTeamPrimaryChannelFilesFolder:                                  groupTeamPrimaryChannelFilesFolderClient,
		GroupTeamPrimaryChannelFilesFolderContent:                           groupTeamPrimaryChannelFilesFolderContentClient,
		GroupTeamPrimaryChannelMember:                                       groupTeamPrimaryChannelMemberClient,
		GroupTeamPrimaryChannelMessage:                                      groupTeamPrimaryChannelMessageClient,
		GroupTeamPrimaryChannelMessageHostedContent:                         groupTeamPrimaryChannelMessageHostedContentClient,
		GroupTeamPrimaryChannelMessageReply:                                 groupTeamPrimaryChannelMessageReplyClient,
		GroupTeamPrimaryChannelMessageReplyHostedContent:                    groupTeamPrimaryChannelMessageReplyHostedContentClient,
		GroupTeamPrimaryChannelSharedWithTeam:                               groupTeamPrimaryChannelSharedWithTeamClient,
		GroupTeamPrimaryChannelSharedWithTeamAllowedMember:                  groupTeamPrimaryChannelSharedWithTeamAllowedMemberClient,
		GroupTeamPrimaryChannelSharedWithTeamTeam:                           groupTeamPrimaryChannelSharedWithTeamTeamClient,
		GroupTeamPrimaryChannelTab:                                          groupTeamPrimaryChannelTabClient,
		GroupTeamPrimaryChannelTabTeamsApp:                                  groupTeamPrimaryChannelTabTeamsAppClient,
		GroupTeamSchedule:                                                   groupTeamScheduleClient,
		GroupTeamScheduleOfferShiftRequest:                                  groupTeamScheduleOfferShiftRequestClient,
		GroupTeamScheduleOpenShift:                                          groupTeamScheduleOpenShiftClient,
		GroupTeamScheduleOpenShiftChangeRequest:                             groupTeamScheduleOpenShiftChangeRequestClient,
		GroupTeamScheduleSchedulingGroup:                                    groupTeamScheduleSchedulingGroupClient,
		GroupTeamScheduleShift:                                              groupTeamScheduleShiftClient,
		GroupTeamScheduleSwapShiftsChangeRequest:                            groupTeamScheduleSwapShiftsChangeRequestClient,
		GroupTeamScheduleTimeCard:                                           groupTeamScheduleTimeCardClient,
		GroupTeamScheduleTimeOffReason:                                      groupTeamScheduleTimeOffReasonClient,
		GroupTeamScheduleTimeOffRequest:                                     groupTeamScheduleTimeOffRequestClient,
		GroupTeamScheduleTimesOff:                                           groupTeamScheduleTimesOffClient,
		GroupTeamTag:                                                        groupTeamTagClient,
		GroupTeamTagMember:                                                  groupTeamTagMemberClient,
		GroupTeamTemplate:                                                   groupTeamTemplateClient,
		GroupTeamTemplateDefinition:                                         groupTeamTemplateDefinitionClient,
		GroupThread:                                                         groupThreadClient,
		GroupThreadPost:                                                     groupThreadPostClient,
		GroupThreadPostAttachment:                                           groupThreadPostAttachmentClient,
		GroupThreadPostExtension:                                            groupThreadPostExtensionClient,
		GroupThreadPostInReplyTo:                                            groupThreadPostInReplyToClient,
		GroupThreadPostInReplyToAttachment:                                  groupThreadPostInReplyToAttachmentClient,
		GroupThreadPostInReplyToExtension:                                   groupThreadPostInReplyToExtensionClient,
		GroupThreadPostInReplyToMention:                                     groupThreadPostInReplyToMentionClient,
		GroupThreadPostMention:                                              groupThreadPostMentionClient,
		GroupTransitiveMember:                                               groupTransitiveMemberClient,
		GroupTransitiveMemberOf:                                             groupTransitiveMemberOfClient,
		SetGroupSiteTermStoreGroupSetChildren:                               setGroupSiteTermStoreGroupSetChildrenClient,
		SetGroupSiteTermStoreGroupSetChildrenChildren:                       setGroupSiteTermStoreGroupSetChildrenChildrenClient,
		SetGroupSiteTermStoreGroupSetChildrenChildrenRelation:               setGroupSiteTermStoreGroupSetChildrenChildrenRelationClient,
		SetGroupSiteTermStoreGroupSetChildrenRelation:                       setGroupSiteTermStoreGroupSetChildrenRelationClient,
		SetGroupSiteTermStoreGroupSetRelation:                               setGroupSiteTermStoreGroupSetRelationClient,
		SetGroupSiteTermStoreGroupSetTerm:                                   setGroupSiteTermStoreGroupSetTermClient,
		SetGroupSiteTermStoreGroupSetTermChildren:                           setGroupSiteTermStoreGroupSetTermChildrenClient,
		SetGroupSiteTermStoreGroupSetTermChildrenRelation:                   setGroupSiteTermStoreGroupSetTermChildrenRelationClient,
		SetGroupSiteTermStoreGroupSetTermRelation:                           setGroupSiteTermStoreGroupSetTermRelationClient,
		SetGroupSiteTermStoreSetChildren:                                    setGroupSiteTermStoreSetChildrenClient,
		SetGroupSiteTermStoreSetChildrenChildren:                            setGroupSiteTermStoreSetChildrenChildrenClient,
		SetGroupSiteTermStoreSetChildrenChildrenRelation:                    setGroupSiteTermStoreSetChildrenChildrenRelationClient,
		SetGroupSiteTermStoreSetChildrenRelation:                            setGroupSiteTermStoreSetChildrenRelationClient,
		SetGroupSiteTermStoreSetParentGroupSetChildren:                      setGroupSiteTermStoreSetParentGroupSetChildrenClient,
		SetGroupSiteTermStoreSetParentGroupSetChildrenChildren:              setGroupSiteTermStoreSetParentGroupSetChildrenChildrenClient,
		SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelation:      setGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient,
		SetGroupSiteTermStoreSetParentGroupSetChildrenRelation:              setGroupSiteTermStoreSetParentGroupSetChildrenRelationClient,
		SetGroupSiteTermStoreSetParentGroupSetRelation:                      setGroupSiteTermStoreSetParentGroupSetRelationClient,
		SetGroupSiteTermStoreSetParentGroupSetTerm:                          setGroupSiteTermStoreSetParentGroupSetTermClient,
		SetGroupSiteTermStoreSetParentGroupSetTermChildren:                  setGroupSiteTermStoreSetParentGroupSetTermChildrenClient,
		SetGroupSiteTermStoreSetParentGroupSetTermChildrenRelation:          setGroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient,
		SetGroupSiteTermStoreSetParentGroupSetTermRelation:                  setGroupSiteTermStoreSetParentGroupSetTermRelationClient,
		SetGroupSiteTermStoreSetRelation:                                    setGroupSiteTermStoreSetRelationClient,
		SetGroupSiteTermStoreSetTerm:                                        setGroupSiteTermStoreSetTermClient,
		SetGroupSiteTermStoreSetTermChildren:                                setGroupSiteTermStoreSetTermChildrenClient,
		SetGroupSiteTermStoreSetTermChildrenRelation:                        setGroupSiteTermStoreSetTermChildrenRelationClient,
		SetGroupSiteTermStoreSetTermRelation:                                setGroupSiteTermStoreSetTermRelationClient,
	}, nil
}
