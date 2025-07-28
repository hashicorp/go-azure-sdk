package wafloganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WafRankingsResponseDataInlined struct {
	GroupValues *[]string                                       `json:"groupValues,omitempty"`
	Metrics     *[]WafRankingsResponseDataInlinedMetricsInlined `json:"metrics,omitempty"`
}
