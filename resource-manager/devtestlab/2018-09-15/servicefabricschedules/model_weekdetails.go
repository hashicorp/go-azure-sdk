package servicefabricschedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WeekDetails struct {
	Time     *string   `json:"time,omitempty"`
	Weekdays *[]string `json:"weekdays,omitempty"`
}
