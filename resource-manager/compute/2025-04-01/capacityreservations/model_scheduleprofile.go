package capacityreservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleProfile struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}
