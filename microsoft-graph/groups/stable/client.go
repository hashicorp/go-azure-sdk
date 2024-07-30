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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalytic"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticalltime"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticitemactivitystat"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticitemactivitystatactivity"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticitemactivitystatactivitydriveitem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticitemactivitystatactivitydriveitemcontent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/siteanalyticlastsevenday"
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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/sitelistitemanalytic"
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
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AcceptedSender                                         *acceptedsender.AcceptedSenderClient
	AppRoleAssignment                                      *approleassignment.AppRoleAssignmentClient
	Calendar                                               *calendar.CalendarClient
	CalendarCalendarPermission                             *calendarcalendarpermission.CalendarCalendarPermissionClient
	CalendarCalendarView                                   *calendarcalendarview.CalendarCalendarViewClient
	CalendarCalendarViewAttachment                         *calendarcalendarviewattachment.CalendarCalendarViewAttachmentClient
	CalendarCalendarViewCalendar                           *calendarcalendarviewcalendar.CalendarCalendarViewCalendarClient
	CalendarCalendarViewExtension                          *calendarcalendarviewextension.CalendarCalendarViewExtensionClient
	CalendarCalendarViewInstance                           *calendarcalendarviewinstance.CalendarCalendarViewInstanceClient
	CalendarCalendarViewInstanceAttachment                 *calendarcalendarviewinstanceattachment.CalendarCalendarViewInstanceAttachmentClient
	CalendarCalendarViewInstanceCalendar                   *calendarcalendarviewinstancecalendar.CalendarCalendarViewInstanceCalendarClient
	CalendarCalendarViewInstanceExtension                  *calendarcalendarviewinstanceextension.CalendarCalendarViewInstanceExtensionClient
	CalendarEvent                                          *calendarevent.CalendarEventClient
	CalendarEventAttachment                                *calendareventattachment.CalendarEventAttachmentClient
	CalendarEventCalendar                                  *calendareventcalendar.CalendarEventCalendarClient
	CalendarEventExtension                                 *calendareventextension.CalendarEventExtensionClient
	CalendarEventInstance                                  *calendareventinstance.CalendarEventInstanceClient
	CalendarEventInstanceAttachment                        *calendareventinstanceattachment.CalendarEventInstanceAttachmentClient
	CalendarEventInstanceCalendar                          *calendareventinstancecalendar.CalendarEventInstanceCalendarClient
	CalendarEventInstanceExtension                         *calendareventinstanceextension.CalendarEventInstanceExtensionClient
	CalendarView                                           *calendarview.CalendarViewClient
	CalendarViewAttachment                                 *calendarviewattachment.CalendarViewAttachmentClient
	CalendarViewCalendar                                   *calendarviewcalendar.CalendarViewCalendarClient
	CalendarViewExtension                                  *calendarviewextension.CalendarViewExtensionClient
	CalendarViewInstance                                   *calendarviewinstance.CalendarViewInstanceClient
	CalendarViewInstanceAttachment                         *calendarviewinstanceattachment.CalendarViewInstanceAttachmentClient
	CalendarViewInstanceCalendar                           *calendarviewinstancecalendar.CalendarViewInstanceCalendarClient
	CalendarViewInstanceExtension                          *calendarviewinstanceextension.CalendarViewInstanceExtensionClient
	Conversation                                           *conversation.ConversationClient
	ConversationThread                                     *conversationthread.ConversationThreadClient
	ConversationThreadPost                                 *conversationthreadpost.ConversationThreadPostClient
	ConversationThreadPostAttachment                       *conversationthreadpostattachment.ConversationThreadPostAttachmentClient
	ConversationThreadPostExtension                        *conversationthreadpostextension.ConversationThreadPostExtensionClient
	ConversationThreadPostInReplyTo                        *conversationthreadpostinreplyto.ConversationThreadPostInReplyToClient
	ConversationThreadPostInReplyToAttachment              *conversationthreadpostinreplytoattachment.ConversationThreadPostInReplyToAttachmentClient
	ConversationThreadPostInReplyToExtension               *conversationthreadpostinreplytoextension.ConversationThreadPostInReplyToExtensionClient
	CreatedOnBehalfOf                                      *createdonbehalfof.CreatedOnBehalfOfClient
	Drive                                                  *drive.DriveClient
	Event                                                  *event.EventClient
	EventAttachment                                        *eventattachment.EventAttachmentClient
	EventCalendar                                          *eventcalendar.EventCalendarClient
	EventExtension                                         *eventextension.EventExtensionClient
	EventInstance                                          *eventinstance.EventInstanceClient
	EventInstanceAttachment                                *eventinstanceattachment.EventInstanceAttachmentClient
	EventInstanceCalendar                                  *eventinstancecalendar.EventInstanceCalendarClient
	EventInstanceExtension                                 *eventinstanceextension.EventInstanceExtensionClient
	Extension                                              *extension.ExtensionClient
	Group                                                  *group.GroupClient
	GroupLifecyclePolicy                                   *grouplifecyclepolicy.GroupLifecyclePolicyClient
	Member                                                 *member.MemberClient
	MemberOf                                               *memberof.MemberOfClient
	MembersWithLicenseError                                *memberswithlicenseerror.MembersWithLicenseErrorClient
	Owner                                                  *owner.OwnerClient
	PermissionGrant                                        *permissiongrant.PermissionGrantClient
	Photo                                                  *photo.PhotoClient
	Planner                                                *planner.PlannerClient
	PlannerPlan                                            *plannerplan.PlannerPlanClient
	PlannerPlanBucket                                      *plannerplanbucket.PlannerPlanBucketClient
	PlannerPlanBucketTask                                  *plannerplanbuckettask.PlannerPlanBucketTaskClient
	PlannerPlanBucketTaskAssignedToTaskBoardFormat         *plannerplanbuckettaskassignedtotaskboardformat.PlannerPlanBucketTaskAssignedToTaskBoardFormatClient
	PlannerPlanBucketTaskBucketTaskBoardFormat             *plannerplanbuckettaskbuckettaskboardformat.PlannerPlanBucketTaskBucketTaskBoardFormatClient
	PlannerPlanBucketTaskDetail                            *plannerplanbuckettaskdetail.PlannerPlanBucketTaskDetailClient
	PlannerPlanBucketTaskProgressTaskBoardFormat           *plannerplanbuckettaskprogresstaskboardformat.PlannerPlanBucketTaskProgressTaskBoardFormatClient
	PlannerPlanDetail                                      *plannerplandetail.PlannerPlanDetailClient
	PlannerPlanTask                                        *plannerplantask.PlannerPlanTaskClient
	PlannerPlanTaskAssignedToTaskBoardFormat               *plannerplantaskassignedtotaskboardformat.PlannerPlanTaskAssignedToTaskBoardFormatClient
	PlannerPlanTaskBucketTaskBoardFormat                   *plannerplantaskbuckettaskboardformat.PlannerPlanTaskBucketTaskBoardFormatClient
	PlannerPlanTaskDetail                                  *plannerplantaskdetail.PlannerPlanTaskDetailClient
	PlannerPlanTaskProgressTaskBoardFormat                 *plannerplantaskprogresstaskboardformat.PlannerPlanTaskProgressTaskBoardFormatClient
	RejectedSender                                         *rejectedsender.RejectedSenderClient
	ServiceProvisioningError                               *serviceprovisioningerror.ServiceProvisioningErrorClient
	Setting                                                *setting.SettingClient
	Site                                                   *site.SiteClient
	SiteAnalytic                                           *siteanalytic.SiteAnalyticClient
	SiteAnalyticAllTime                                    *siteanalyticalltime.SiteAnalyticAllTimeClient
	SiteAnalyticItemActivityStat                           *siteanalyticitemactivitystat.SiteAnalyticItemActivityStatClient
	SiteAnalyticItemActivityStatActivity                   *siteanalyticitemactivitystatactivity.SiteAnalyticItemActivityStatActivityClient
	SiteAnalyticItemActivityStatActivityDriveItem          *siteanalyticitemactivitystatactivitydriveitem.SiteAnalyticItemActivityStatActivityDriveItemClient
	SiteAnalyticItemActivityStatActivityDriveItemContent   *siteanalyticitemactivitystatactivitydriveitemcontent.SiteAnalyticItemActivityStatActivityDriveItemContentClient
	SiteAnalyticLastSevenDay                               *siteanalyticlastsevenday.SiteAnalyticLastSevenDayClient
	SiteColumn                                             *sitecolumn.SiteColumnClient
	SiteColumnSourceColumn                                 *sitecolumnsourcecolumn.SiteColumnSourceColumnClient
	SiteContentType                                        *sitecontenttype.SiteContentTypeClient
	SiteContentTypeBase                                    *sitecontenttypebase.SiteContentTypeBaseClient
	SiteContentTypeBaseType                                *sitecontenttypebasetype.SiteContentTypeBaseTypeClient
	SiteContentTypeColumn                                  *sitecontenttypecolumn.SiteContentTypeColumnClient
	SiteContentTypeColumnLink                              *sitecontenttypecolumnlink.SiteContentTypeColumnLinkClient
	SiteContentTypeColumnPosition                          *sitecontenttypecolumnposition.SiteContentTypeColumnPositionClient
	SiteContentTypeColumnSourceColumn                      *sitecontenttypecolumnsourcecolumn.SiteContentTypeColumnSourceColumnClient
	SiteCreatedByUser                                      *sitecreatedbyuser.SiteCreatedByUserClient
	SiteCreatedByUserMailboxSetting                        *sitecreatedbyusermailboxsetting.SiteCreatedByUserMailboxSettingClient
	SiteCreatedByUserServiceProvisioningError              *sitecreatedbyuserserviceprovisioningerror.SiteCreatedByUserServiceProvisioningErrorClient
	SiteDrive                                              *sitedrive.SiteDriveClient
	SiteExternalColumn                                     *siteexternalcolumn.SiteExternalColumnClient
	SiteItem                                               *siteitem.SiteItemClient
	SiteLastModifiedByUser                                 *sitelastmodifiedbyuser.SiteLastModifiedByUserClient
	SiteLastModifiedByUserMailboxSetting                   *sitelastmodifiedbyusermailboxsetting.SiteLastModifiedByUserMailboxSettingClient
	SiteLastModifiedByUserServiceProvisioningError         *sitelastmodifiedbyuserserviceprovisioningerror.SiteLastModifiedByUserServiceProvisioningErrorClient
	SiteList                                               *sitelist.SiteListClient
	SiteListColumn                                         *sitelistcolumn.SiteListColumnClient
	SiteListColumnSourceColumn                             *sitelistcolumnsourcecolumn.SiteListColumnSourceColumnClient
	SiteListContentType                                    *sitelistcontenttype.SiteListContentTypeClient
	SiteListContentTypeBase                                *sitelistcontenttypebase.SiteListContentTypeBaseClient
	SiteListContentTypeBaseType                            *sitelistcontenttypebasetype.SiteListContentTypeBaseTypeClient
	SiteListContentTypeColumn                              *sitelistcontenttypecolumn.SiteListContentTypeColumnClient
	SiteListContentTypeColumnLink                          *sitelistcontenttypecolumnlink.SiteListContentTypeColumnLinkClient
	SiteListContentTypeColumnPosition                      *sitelistcontenttypecolumnposition.SiteListContentTypeColumnPositionClient
	SiteListContentTypeColumnSourceColumn                  *sitelistcontenttypecolumnsourcecolumn.SiteListContentTypeColumnSourceColumnClient
	SiteListCreatedByUser                                  *sitelistcreatedbyuser.SiteListCreatedByUserClient
	SiteListCreatedByUserMailboxSetting                    *sitelistcreatedbyusermailboxsetting.SiteListCreatedByUserMailboxSettingClient
	SiteListCreatedByUserServiceProvisioningError          *sitelistcreatedbyuserserviceprovisioningerror.SiteListCreatedByUserServiceProvisioningErrorClient
	SiteListDrive                                          *sitelistdrive.SiteListDriveClient
	SiteListItem                                           *sitelistitem.SiteListItemClient
	SiteListItemAnalytic                                   *sitelistitemanalytic.SiteListItemAnalyticClient
	SiteListItemCreatedByUser                              *sitelistitemcreatedbyuser.SiteListItemCreatedByUserClient
	SiteListItemCreatedByUserMailboxSetting                *sitelistitemcreatedbyusermailboxsetting.SiteListItemCreatedByUserMailboxSettingClient
	SiteListItemCreatedByUserServiceProvisioningError      *sitelistitemcreatedbyuserserviceprovisioningerror.SiteListItemCreatedByUserServiceProvisioningErrorClient
	SiteListItemDocumentSetVersion                         *sitelistitemdocumentsetversion.SiteListItemDocumentSetVersionClient
	SiteListItemDocumentSetVersionField                    *sitelistitemdocumentsetversionfield.SiteListItemDocumentSetVersionFieldClient
	SiteListItemDriveItem                                  *sitelistitemdriveitem.SiteListItemDriveItemClient
	SiteListItemDriveItemContent                           *sitelistitemdriveitemcontent.SiteListItemDriveItemContentClient
	SiteListItemField                                      *sitelistitemfield.SiteListItemFieldClient
	SiteListItemLastModifiedByUser                         *sitelistitemlastmodifiedbyuser.SiteListItemLastModifiedByUserClient
	SiteListItemLastModifiedByUserMailboxSetting           *sitelistitemlastmodifiedbyusermailboxsetting.SiteListItemLastModifiedByUserMailboxSettingClient
	SiteListItemLastModifiedByUserServiceProvisioningError *sitelistitemlastmodifiedbyuserserviceprovisioningerror.SiteListItemLastModifiedByUserServiceProvisioningErrorClient
	SiteListItemVersion                                    *sitelistitemversion.SiteListItemVersionClient
	SiteListItemVersionField                               *sitelistitemversionfield.SiteListItemVersionFieldClient
	SiteListLastModifiedByUser                             *sitelistlastmodifiedbyuser.SiteListLastModifiedByUserClient
	SiteListLastModifiedByUserMailboxSetting               *sitelistlastmodifiedbyusermailboxsetting.SiteListLastModifiedByUserMailboxSettingClient
	SiteListLastModifiedByUserServiceProvisioningError     *sitelistlastmodifiedbyuserserviceprovisioningerror.SiteListLastModifiedByUserServiceProvisioningErrorClient
	SiteListOperation                                      *sitelistoperation.SiteListOperationClient
	SiteListSubscription                                   *sitelistsubscription.SiteListSubscriptionClient
	SiteOperation                                          *siteoperation.SiteOperationClient
	SitePermission                                         *sitepermission.SitePermissionClient
	Team                                                   *team.TeamClient
	TeamAllChannel                                         *teamallchannel.TeamAllChannelClient
	TeamChannel                                            *teamchannel.TeamChannelClient
	TeamChannelFilesFolder                                 *teamchannelfilesfolder.TeamChannelFilesFolderClient
	TeamChannelFilesFolderContent                          *teamchannelfilesfoldercontent.TeamChannelFilesFolderContentClient
	TeamChannelMember                                      *teamchannelmember.TeamChannelMemberClient
	TeamChannelMessage                                     *teamchannelmessage.TeamChannelMessageClient
	TeamChannelMessageHostedContent                        *teamchannelmessagehostedcontent.TeamChannelMessageHostedContentClient
	TeamChannelMessageReply                                *teamchannelmessagereply.TeamChannelMessageReplyClient
	TeamChannelMessageReplyHostedContent                   *teamchannelmessagereplyhostedcontent.TeamChannelMessageReplyHostedContentClient
	TeamChannelSharedWithTeam                              *teamchannelsharedwithteam.TeamChannelSharedWithTeamClient
	TeamChannelSharedWithTeamAllowedMember                 *teamchannelsharedwithteamallowedmember.TeamChannelSharedWithTeamAllowedMemberClient
	TeamChannelSharedWithTeamTeam                          *teamchannelsharedwithteamteam.TeamChannelSharedWithTeamTeamClient
	TeamChannelTab                                         *teamchanneltab.TeamChannelTabClient
	TeamChannelTabTeamsApp                                 *teamchanneltabteamsapp.TeamChannelTabTeamsAppClient
	TeamGroup                                              *teamgroup.TeamGroupClient
	TeamGroupServiceProvisioningError                      *teamgroupserviceprovisioningerror.TeamGroupServiceProvisioningErrorClient
	TeamIncomingChannel                                    *teamincomingchannel.TeamIncomingChannelClient
	TeamInstalledApp                                       *teaminstalledapp.TeamInstalledAppClient
	TeamInstalledAppTeamsApp                               *teaminstalledappteamsapp.TeamInstalledAppTeamsAppClient
	TeamInstalledAppTeamsAppDefinition                     *teaminstalledappteamsappdefinition.TeamInstalledAppTeamsAppDefinitionClient
	TeamMember                                             *teammember.TeamMemberClient
	TeamOperation                                          *teamoperation.TeamOperationClient
	TeamPermissionGrant                                    *teampermissiongrant.TeamPermissionGrantClient
	TeamPhoto                                              *teamphoto.TeamPhotoClient
	TeamPrimaryChannel                                     *teamprimarychannel.TeamPrimaryChannelClient
	TeamPrimaryChannelFilesFolder                          *teamprimarychannelfilesfolder.TeamPrimaryChannelFilesFolderClient
	TeamPrimaryChannelFilesFolderContent                   *teamprimarychannelfilesfoldercontent.TeamPrimaryChannelFilesFolderContentClient
	TeamPrimaryChannelMember                               *teamprimarychannelmember.TeamPrimaryChannelMemberClient
	TeamPrimaryChannelMessage                              *teamprimarychannelmessage.TeamPrimaryChannelMessageClient
	TeamPrimaryChannelMessageHostedContent                 *teamprimarychannelmessagehostedcontent.TeamPrimaryChannelMessageHostedContentClient
	TeamPrimaryChannelMessageReply                         *teamprimarychannelmessagereply.TeamPrimaryChannelMessageReplyClient
	TeamPrimaryChannelMessageReplyHostedContent            *teamprimarychannelmessagereplyhostedcontent.TeamPrimaryChannelMessageReplyHostedContentClient
	TeamPrimaryChannelSharedWithTeam                       *teamprimarychannelsharedwithteam.TeamPrimaryChannelSharedWithTeamClient
	TeamPrimaryChannelSharedWithTeamAllowedMember          *teamprimarychannelsharedwithteamallowedmember.TeamPrimaryChannelSharedWithTeamAllowedMemberClient
	TeamPrimaryChannelSharedWithTeamTeam                   *teamprimarychannelsharedwithteamteam.TeamPrimaryChannelSharedWithTeamTeamClient
	TeamPrimaryChannelTab                                  *teamprimarychanneltab.TeamPrimaryChannelTabClient
	TeamPrimaryChannelTabTeamsApp                          *teamprimarychanneltabteamsapp.TeamPrimaryChannelTabTeamsAppClient
	TeamSchedule                                           *teamschedule.TeamScheduleClient
	TeamScheduleOfferShiftRequest                          *teamscheduleoffershiftrequest.TeamScheduleOfferShiftRequestClient
	TeamScheduleOpenShift                                  *teamscheduleopenshift.TeamScheduleOpenShiftClient
	TeamScheduleOpenShiftChangeRequest                     *teamscheduleopenshiftchangerequest.TeamScheduleOpenShiftChangeRequestClient
	TeamScheduleSchedulingGroup                            *teamscheduleschedulinggroup.TeamScheduleSchedulingGroupClient
	TeamScheduleShift                                      *teamscheduleshift.TeamScheduleShiftClient
	TeamScheduleSwapShiftsChangeRequest                    *teamscheduleswapshiftschangerequest.TeamScheduleSwapShiftsChangeRequestClient
	TeamScheduleTimeOffReason                              *teamscheduletimeoffreason.TeamScheduleTimeOffReasonClient
	TeamScheduleTimeOffRequest                             *teamscheduletimeoffrequest.TeamScheduleTimeOffRequestClient
	TeamScheduleTimesOff                                   *teamscheduletimesoff.TeamScheduleTimesOffClient
	TeamTag                                                *teamtag.TeamTagClient
	TeamTagMember                                          *teamtagmember.TeamTagMemberClient
	TeamTemplate                                           *teamtemplate.TeamTemplateClient
	Thread                                                 *thread.ThreadClient
	ThreadPost                                             *threadpost.ThreadPostClient
	ThreadPostAttachment                                   *threadpostattachment.ThreadPostAttachmentClient
	ThreadPostExtension                                    *threadpostextension.ThreadPostExtensionClient
	ThreadPostInReplyTo                                    *threadpostinreplyto.ThreadPostInReplyToClient
	ThreadPostInReplyToAttachment                          *threadpostinreplytoattachment.ThreadPostInReplyToAttachmentClient
	ThreadPostInReplyToExtension                           *threadpostinreplytoextension.ThreadPostInReplyToExtensionClient
	TransitiveMember                                       *transitivemember.TransitiveMemberClient
	TransitiveMemberOf                                     *transitivememberof.TransitiveMemberOfClient
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

	driveClient, err := drive.NewDriveClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Drive client: %+v", err)
	}
	configureFunc(driveClient.Client)

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
		AcceptedSender:                            acceptedSenderClient,
		AppRoleAssignment:                         appRoleAssignmentClient,
		Calendar:                                  calendarClient,
		CalendarCalendarPermission:                calendarCalendarPermissionClient,
		CalendarCalendarView:                      calendarCalendarViewClient,
		CalendarCalendarViewAttachment:            calendarCalendarViewAttachmentClient,
		CalendarCalendarViewCalendar:              calendarCalendarViewCalendarClient,
		CalendarCalendarViewExtension:             calendarCalendarViewExtensionClient,
		CalendarCalendarViewInstance:              calendarCalendarViewInstanceClient,
		CalendarCalendarViewInstanceAttachment:    calendarCalendarViewInstanceAttachmentClient,
		CalendarCalendarViewInstanceCalendar:      calendarCalendarViewInstanceCalendarClient,
		CalendarCalendarViewInstanceExtension:     calendarCalendarViewInstanceExtensionClient,
		CalendarEvent:                             calendarEventClient,
		CalendarEventAttachment:                   calendarEventAttachmentClient,
		CalendarEventCalendar:                     calendarEventCalendarClient,
		CalendarEventExtension:                    calendarEventExtensionClient,
		CalendarEventInstance:                     calendarEventInstanceClient,
		CalendarEventInstanceAttachment:           calendarEventInstanceAttachmentClient,
		CalendarEventInstanceCalendar:             calendarEventInstanceCalendarClient,
		CalendarEventInstanceExtension:            calendarEventInstanceExtensionClient,
		CalendarView:                              calendarViewClient,
		CalendarViewAttachment:                    calendarViewAttachmentClient,
		CalendarViewCalendar:                      calendarViewCalendarClient,
		CalendarViewExtension:                     calendarViewExtensionClient,
		CalendarViewInstance:                      calendarViewInstanceClient,
		CalendarViewInstanceAttachment:            calendarViewInstanceAttachmentClient,
		CalendarViewInstanceCalendar:              calendarViewInstanceCalendarClient,
		CalendarViewInstanceExtension:             calendarViewInstanceExtensionClient,
		Conversation:                              conversationClient,
		ConversationThread:                        conversationThreadClient,
		ConversationThreadPost:                    conversationThreadPostClient,
		ConversationThreadPostAttachment:          conversationThreadPostAttachmentClient,
		ConversationThreadPostExtension:           conversationThreadPostExtensionClient,
		ConversationThreadPostInReplyTo:           conversationThreadPostInReplyToClient,
		ConversationThreadPostInReplyToAttachment: conversationThreadPostInReplyToAttachmentClient,
		ConversationThreadPostInReplyToExtension:  conversationThreadPostInReplyToExtensionClient,
		CreatedOnBehalfOf:                         createdOnBehalfOfClient,
		Drive:                                     driveClient,
		Event:                                     eventClient,
		EventAttachment:                           eventAttachmentClient,
		EventCalendar:                             eventCalendarClient,
		EventExtension:                            eventExtensionClient,
		EventInstance:                             eventInstanceClient,
		EventInstanceAttachment:                   eventInstanceAttachmentClient,
		EventInstanceCalendar:                     eventInstanceCalendarClient,
		EventInstanceExtension:                    eventInstanceExtensionClient,
		Extension:                                 extensionClient,
		Group:                                     groupClient,
		GroupLifecyclePolicy:                      groupLifecyclePolicyClient,
		Member:                                    memberClient,
		MemberOf:                                  memberOfClient,
		MembersWithLicenseError:                   membersWithLicenseErrorClient,
		Owner:                                     ownerClient,
		PermissionGrant:                           permissionGrantClient,
		Photo:                                     photoClient,
		Planner:                                   plannerClient,
		PlannerPlan:                               plannerPlanClient,
		PlannerPlanBucket:                         plannerPlanBucketClient,
		PlannerPlanBucketTask:                     plannerPlanBucketTaskClient,
		PlannerPlanBucketTaskAssignedToTaskBoardFormat:       plannerPlanBucketTaskAssignedToTaskBoardFormatClient,
		PlannerPlanBucketTaskBucketTaskBoardFormat:           plannerPlanBucketTaskBucketTaskBoardFormatClient,
		PlannerPlanBucketTaskDetail:                          plannerPlanBucketTaskDetailClient,
		PlannerPlanBucketTaskProgressTaskBoardFormat:         plannerPlanBucketTaskProgressTaskBoardFormatClient,
		PlannerPlanDetail:                                    plannerPlanDetailClient,
		PlannerPlanTask:                                      plannerPlanTaskClient,
		PlannerPlanTaskAssignedToTaskBoardFormat:             plannerPlanTaskAssignedToTaskBoardFormatClient,
		PlannerPlanTaskBucketTaskBoardFormat:                 plannerPlanTaskBucketTaskBoardFormatClient,
		PlannerPlanTaskDetail:                                plannerPlanTaskDetailClient,
		PlannerPlanTaskProgressTaskBoardFormat:               plannerPlanTaskProgressTaskBoardFormatClient,
		RejectedSender:                                       rejectedSenderClient,
		ServiceProvisioningError:                             serviceProvisioningErrorClient,
		Setting:                                              settingClient,
		Site:                                                 siteClient,
		SiteAnalytic:                                         siteAnalyticClient,
		SiteAnalyticAllTime:                                  siteAnalyticAllTimeClient,
		SiteAnalyticItemActivityStat:                         siteAnalyticItemActivityStatClient,
		SiteAnalyticItemActivityStatActivity:                 siteAnalyticItemActivityStatActivityClient,
		SiteAnalyticItemActivityStatActivityDriveItem:        siteAnalyticItemActivityStatActivityDriveItemClient,
		SiteAnalyticItemActivityStatActivityDriveItemContent: siteAnalyticItemActivityStatActivityDriveItemContentClient,
		SiteAnalyticLastSevenDay:                             siteAnalyticLastSevenDayClient,
		SiteColumn:                                           siteColumnClient,
		SiteColumnSourceColumn:                               siteColumnSourceColumnClient,
		SiteContentType:                                      siteContentTypeClient,
		SiteContentTypeBase:                                  siteContentTypeBaseClient,
		SiteContentTypeBaseType:                              siteContentTypeBaseTypeClient,
		SiteContentTypeColumn:                                siteContentTypeColumnClient,
		SiteContentTypeColumnLink:                            siteContentTypeColumnLinkClient,
		SiteContentTypeColumnPosition:                        siteContentTypeColumnPositionClient,
		SiteContentTypeColumnSourceColumn:                    siteContentTypeColumnSourceColumnClient,
		SiteCreatedByUser:                                    siteCreatedByUserClient,
		SiteCreatedByUserMailboxSetting:                      siteCreatedByUserMailboxSettingClient,
		SiteCreatedByUserServiceProvisioningError:            siteCreatedByUserServiceProvisioningErrorClient,
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
		SiteListItemAnalytic:                                   siteListItemAnalyticClient,
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
