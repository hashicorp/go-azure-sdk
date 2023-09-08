package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptValidationResult struct {
	ODataType    *string                            `json:"@odata.type,omitempty"`
	RuleErrors   *[]DeviceComplianceScriptRuleError `json:"ruleErrors,omitempty"`
	Rules        *[]DeviceComplianceScriptRule      `json:"rules,omitempty"`
	ScriptErrors *[]DeviceComplianceScriptError     `json:"scriptErrors,omitempty"`
}
