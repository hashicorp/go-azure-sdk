package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRule struct {
	DataType                           *DeviceComplianceScriptRuleDataType                           `json:"dataType,omitempty"`
	DeviceComplianceScriptRulOperator  *DeviceComplianceScriptRuleDeviceComplianceScriptRulOperator  `json:"deviceComplianceScriptRulOperator,omitempty"`
	DeviceComplianceScriptRuleDataType *DeviceComplianceScriptRuleDeviceComplianceScriptRuleDataType `json:"deviceComplianceScriptRuleDataType,omitempty"`
	ODataType                          *string                                                       `json:"@odata.type,omitempty"`
	Operand                            *string                                                       `json:"operand,omitempty"`
	Operator                           *DeviceComplianceScriptRuleOperator                           `json:"operator,omitempty"`
	SettingName                        *string                                                       `json:"settingName,omitempty"`
}
