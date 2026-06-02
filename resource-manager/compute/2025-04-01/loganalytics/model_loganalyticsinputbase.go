package loganalytics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogAnalyticsInputBase struct {
	BlobContainerSasUri        string `json:"blobContainerSasUri"`
	FromTime                   string `json:"fromTime"`
	GroupByClientApplicationId *bool  `json:"groupByClientApplicationId,omitempty"`
	GroupByOperationName       *bool  `json:"groupByOperationName,omitempty"`
	GroupByResourceName        *bool  `json:"groupByResourceName,omitempty"`
	GroupByThrottlePolicy      *bool  `json:"groupByThrottlePolicy,omitempty"`
	GroupByUserAgent           *bool  `json:"groupByUserAgent,omitempty"`
	ToTime                     string `json:"toTime"`
}

func (o *LogAnalyticsInputBase) GetFromTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.FromTime, "2006-01-02T15:04:05Z07:00")
}

func (o *LogAnalyticsInputBase) SetFromTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FromTime = formatted
}

func (o *LogAnalyticsInputBase) GetToTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.ToTime, "2006-01-02T15:04:05Z07:00")
}

func (o *LogAnalyticsInputBase) SetToTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ToTime = formatted
}
