package devicesecuritygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeWindowCustomAlertRule struct {
	Description    *string `json:"description,omitempty"`
	DisplayName    *string `json:"displayName,omitempty"`
	IsEnabled      *bool   `json:"isEnabled,omitempty"`
	MaxThreshold   int64   `json:"maxThreshold"`
	MinThreshold   int64   `json:"minThreshold"`
	RuleType       *string `json:"ruleType,omitempty"`
	TimeWindowSize string  `json:"timeWindowSize"`
}
