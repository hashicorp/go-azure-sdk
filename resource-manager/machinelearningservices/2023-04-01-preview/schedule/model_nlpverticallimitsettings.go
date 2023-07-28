package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NlpVerticalLimitSettings struct {
	MaxConcurrentTrials *int64  `json:"maxConcurrentTrials,omitempty"`
	MaxNodes            *int64  `json:"maxNodes,omitempty"`
	MaxTrials           *int64  `json:"maxTrials,omitempty"`
	Timeout             *string `json:"timeout,omitempty"`
	TrialTimeout        *string `json:"trialTimeout,omitempty"`
}
