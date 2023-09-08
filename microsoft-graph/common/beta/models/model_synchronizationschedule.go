package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationSchedule struct {
	Expiration *string                       `json:"expiration,omitempty"`
	Interval   *string                       `json:"interval,omitempty"`
	ODataType  *string                       `json:"@odata.type,omitempty"`
	State      *SynchronizationScheduleState `json:"state,omitempty"`
}
