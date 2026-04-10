package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalysisData struct {
	Data               *[][]NameValuePair     `json:"data,omitempty"`
	DetectorDefinition *DetectorDefinition    `json:"detectorDefinition,omitempty"`
	DetectorMetaData   *ResponseMetaData      `json:"detectorMetaData,omitempty"`
	Metrics            *[]DiagnosticMetricSet `json:"metrics,omitempty"`
	Source             *string                `json:"source,omitempty"`
}
