package bandwidthschedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BandwidthScheduleProperties struct {
	Days       []DayOfWeek `json:"days"`
	RateInMbps int64       `json:"rateInMbps"`
	Start      string      `json:"start"`
	Stop       string      `json:"stop"`
}
