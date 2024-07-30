package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Reminder struct {
	ChangeKey        *string           `json:"changeKey,omitempty"`
	EventEndTime     *DateTimeTimeZone `json:"eventEndTime,omitempty"`
	EventId          *string           `json:"eventId,omitempty"`
	EventLocation    *Location         `json:"eventLocation,omitempty"`
	EventStartTime   *DateTimeTimeZone `json:"eventStartTime,omitempty"`
	EventSubject     *string           `json:"eventSubject,omitempty"`
	EventWebLink     *string           `json:"eventWebLink,omitempty"`
	ODataType        *string           `json:"@odata.type,omitempty"`
	ReminderFireTime *DateTimeTimeZone `json:"reminderFireTime,omitempty"`
}
