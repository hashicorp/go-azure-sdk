package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingReminder struct {
	Message    *string                    `json:"message,omitempty"`
	ODataType  *string                    `json:"@odata.type,omitempty"`
	Offset     *string                    `json:"offset,omitempty"`
	Recipients *BookingReminderRecipients `json:"recipients,omitempty"`
}
