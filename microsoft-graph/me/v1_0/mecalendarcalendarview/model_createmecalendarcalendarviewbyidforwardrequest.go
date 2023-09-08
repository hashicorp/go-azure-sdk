package mecalendarcalendarview

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeCalendarCalendarViewByIdForwardRequest struct {
	Comment      *string      `json:"Comment,omitempty"`
	ToRecipients *[]Recipient `json:"ToRecipients,omitempty"`
}
