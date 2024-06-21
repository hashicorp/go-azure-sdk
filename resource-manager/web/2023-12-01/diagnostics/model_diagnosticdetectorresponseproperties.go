package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticDetectorResponseProperties struct {
	AbnormalTimePeriods *[]DetectorAbnormalTimePeriod `json:"abnormalTimePeriods,omitempty"`
	Data                *[][]NameValuePair            `json:"data,omitempty"`
	DetectorDefinition  *DetectorDefinition           `json:"detectorDefinition,omitempty"`
	EndTime             *string                       `json:"endTime,omitempty"`
	IssueDetected       *bool                         `json:"issueDetected,omitempty"`
	Metrics             *[]DiagnosticMetricSet        `json:"metrics,omitempty"`
	ResponseMetaData    *ResponseMetaData             `json:"responseMetaData,omitempty"`
	StartTime           *string                       `json:"startTime,omitempty"`
}
