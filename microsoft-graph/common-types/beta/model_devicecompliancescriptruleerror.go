package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleError struct {
	Code                                       *DeviceComplianceScriptRuleErrorCode                                       `json:"code,omitempty"`
	DeviceComplianceScriptRulesValidationError *DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError `json:"deviceComplianceScriptRulesValidationError,omitempty"`
	Message                                    *string                                                                    `json:"message,omitempty"`
	ODataType                                  *string                                                                    `json:"@odata.type,omitempty"`
	SettingName                                *string                                                                    `json:"settingName,omitempty"`
}
