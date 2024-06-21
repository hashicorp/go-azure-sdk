package loganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestRateByIntervalInput struct {
	BlobContainerSasUri        string         `json:"blobContainerSasUri"`
	FromTime                   string         `json:"fromTime"`
	GroupByClientApplicationId *bool          `json:"groupByClientApplicationId,omitempty"`
	GroupByOperationName       *bool          `json:"groupByOperationName,omitempty"`
	GroupByResourceName        *bool          `json:"groupByResourceName,omitempty"`
	GroupByThrottlePolicy      *bool          `json:"groupByThrottlePolicy,omitempty"`
	GroupByUserAgent           *bool          `json:"groupByUserAgent,omitempty"`
	IntervalLength             IntervalInMins `json:"intervalLength"`
	ToTime                     string         `json:"toTime"`
}
