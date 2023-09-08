package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleItem struct {
	End       *DateTimeTimeZone   `json:"end,omitempty"`
	IsPrivate *bool               `json:"isPrivate,omitempty"`
	Location  *string             `json:"location,omitempty"`
	ODataType *string             `json:"@odata.type,omitempty"`
	Start     *DateTimeTimeZone   `json:"start,omitempty"`
	Status    *ScheduleItemStatus `json:"status,omitempty"`
	Subject   *string             `json:"subject,omitempty"`
}
