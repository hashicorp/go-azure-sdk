package groupevent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupByIdEventByIdSnoozeReminderRequest struct {
	NewReminderTime *DateTimeTimeZone `json:"NewReminderTime,omitempty"`
}
