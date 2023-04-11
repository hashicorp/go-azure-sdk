package devicesecuritygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowlistCustomAlertRule struct {
	AllowlistValues []string   `json:"allowlistValues"`
	Description     *string    `json:"description,omitempty"`
	DisplayName     *string    `json:"displayName,omitempty"`
	IsEnabled       *bool      `json:"isEnabled,omitempty"`
	RuleType        *string    `json:"ruleType,omitempty"`
	ValueType       *ValueType `json:"valueType,omitempty"`
}
