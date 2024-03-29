package regions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoscaleRecurrence struct {
	Schedule *[]AutoscaleSchedule `json:"schedule,omitempty"`
	TimeZone *string              `json:"timeZone,omitempty"`
}
