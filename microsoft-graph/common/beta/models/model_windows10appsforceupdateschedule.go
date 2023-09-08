package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AppsForceUpdateSchedule struct {
	ODataType                          *string                                     `json:"@odata.type,omitempty"`
	Recurrence                         *Windows10AppsForceUpdateScheduleRecurrence `json:"recurrence,omitempty"`
	RunImmediatelyIfAfterStartDateTime *bool                                       `json:"runImmediatelyIfAfterStartDateTime,omitempty"`
	StartDateTime                      *string                                     `json:"startDateTime,omitempty"`
}
