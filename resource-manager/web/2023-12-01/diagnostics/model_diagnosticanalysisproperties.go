package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticAnalysisProperties struct {
	AbnormalTimePeriods    *[]AbnormalTimePeriod `json:"abnormalTimePeriods,omitempty"`
	EndTime                *string               `json:"endTime,omitempty"`
	NonCorrelatedDetectors *[]DetectorDefinition `json:"nonCorrelatedDetectors,omitempty"`
	Payload                *[]AnalysisData       `json:"payload,omitempty"`
	StartTime              *string               `json:"startTime,omitempty"`
}
