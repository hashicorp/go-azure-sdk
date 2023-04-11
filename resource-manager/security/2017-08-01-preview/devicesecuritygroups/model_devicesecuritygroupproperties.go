package devicesecuritygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceSecurityGroupProperties struct {
	AllowlistRules  *[]AllowlistCustomAlertRule  `json:"allowlistRules,omitempty"`
	DenylistRules   *[]DenylistCustomAlertRule   `json:"denylistRules,omitempty"`
	ThresholdRules  *[]ThresholdCustomAlertRule  `json:"thresholdRules,omitempty"`
	TimeWindowRules *[]TimeWindowCustomAlertRule `json:"timeWindowRules,omitempty"`
}
