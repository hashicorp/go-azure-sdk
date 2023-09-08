package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityHistoryItem struct {
	ActiveDurationSeconds *int64                     `json:"activeDurationSeconds,omitempty"`
	Activity              *UserActivity              `json:"activity,omitempty"`
	CreatedDateTime       *string                    `json:"createdDateTime,omitempty"`
	ExpirationDateTime    *string                    `json:"expirationDateTime,omitempty"`
	Id                    *string                    `json:"id,omitempty"`
	LastActiveDateTime    *string                    `json:"lastActiveDateTime,omitempty"`
	LastModifiedDateTime  *string                    `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string                    `json:"@odata.type,omitempty"`
	StartedDateTime       *string                    `json:"startedDateTime,omitempty"`
	Status                *ActivityHistoryItemStatus `json:"status,omitempty"`
	UserTimezone          *string                    `json:"userTimezone,omitempty"`
}
