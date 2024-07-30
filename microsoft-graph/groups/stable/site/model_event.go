package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Event struct {
	AllowNewTimeProposals         *bool                                `json:"allowNewTimeProposals,omitempty"`
	Attachments                   *[]Attachment                        `json:"attachments,omitempty"`
	Attendees                     *[]Attendee                          `json:"attendees,omitempty"`
	Body                          *ItemBody                            `json:"body,omitempty"`
	BodyPreview                   *string                              `json:"bodyPreview,omitempty"`
	Calendar                      *Calendar                            `json:"calendar,omitempty"`
	Categories                    *[]string                            `json:"categories,omitempty"`
	ChangeKey                     *string                              `json:"changeKey,omitempty"`
	CreatedDateTime               *string                              `json:"createdDateTime,omitempty"`
	End                           *DateTimeTimeZone                    `json:"end,omitempty"`
	Extensions                    *[]Extension                         `json:"extensions,omitempty"`
	HasAttachments                *bool                                `json:"hasAttachments,omitempty"`
	HideAttendees                 *bool                                `json:"hideAttendees,omitempty"`
	ICalUId                       *string                              `json:"iCalUId,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	Importance                    *EventImportance                     `json:"importance,omitempty"`
	Instances                     *[]Event                             `json:"instances,omitempty"`
	IsAllDay                      *bool                                `json:"isAllDay,omitempty"`
	IsCancelled                   *bool                                `json:"isCancelled,omitempty"`
	IsDraft                       *bool                                `json:"isDraft,omitempty"`
	IsOnlineMeeting               *bool                                `json:"isOnlineMeeting,omitempty"`
	IsOrganizer                   *bool                                `json:"isOrganizer,omitempty"`
	IsReminderOn                  *bool                                `json:"isReminderOn,omitempty"`
	LastModifiedDateTime          *string                              `json:"lastModifiedDateTime,omitempty"`
	Location                      *Location                            `json:"location,omitempty"`
	Locations                     *[]Location                          `json:"locations,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	OnlineMeeting                 *OnlineMeetingInfo                   `json:"onlineMeeting,omitempty"`
	OnlineMeetingProvider         *EventOnlineMeetingProvider          `json:"onlineMeetingProvider,omitempty"`
	OnlineMeetingUrl              *string                              `json:"onlineMeetingUrl,omitempty"`
	Organizer                     *Recipient                           `json:"organizer,omitempty"`
	OriginalEndTimeZone           *string                              `json:"originalEndTimeZone,omitempty"`
	OriginalStart                 *string                              `json:"originalStart,omitempty"`
	OriginalStartTimeZone         *string                              `json:"originalStartTimeZone,omitempty"`
	Recurrence                    *PatternedRecurrence                 `json:"recurrence,omitempty"`
	ReminderMinutesBeforeStart    *int64                               `json:"reminderMinutesBeforeStart,omitempty"`
	ResponseRequested             *bool                                `json:"responseRequested,omitempty"`
	ResponseStatus                *ResponseStatus                      `json:"responseStatus,omitempty"`
	Sensitivity                   *EventSensitivity                    `json:"sensitivity,omitempty"`
	SeriesMasterId                *string                              `json:"seriesMasterId,omitempty"`
	ShowAs                        *EventShowAs                         `json:"showAs,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	Start                         *DateTimeTimeZone                    `json:"start,omitempty"`
	Subject                       *string                              `json:"subject,omitempty"`
	TransactionId                 *string                              `json:"transactionId,omitempty"`
	Type                          *EventType                           `json:"type,omitempty"`
	WebLink                       *string                              `json:"webLink,omitempty"`
}
