package autoscalesettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoscaleProfile struct {
	Capacity   ScaleCapacity `json:"capacity"`
	FixedDate  *TimeWindow   `json:"fixedDate"`
	Name       string        `json:"name"`
	Recurrence *Recurrence   `json:"recurrence"`
	Rules      []ScaleRule   `json:"rules"`
}
