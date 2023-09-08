package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirementOperationPredicate struct {
	DetectionValue        *string
	DisplayName           *string
	EnforceSignatureCheck *bool
	ODataType             *string
	RunAs32Bit            *bool
	ScriptContent         *string
}

func (p Win32LobAppPowerShellScriptRequirementOperationPredicate) Matches(input Win32LobAppPowerShellScriptRequirement) bool {

	if p.DetectionValue != nil && (input.DetectionValue == nil || *p.DetectionValue != *input.DetectionValue) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnforceSignatureCheck != nil && (input.EnforceSignatureCheck == nil || *p.EnforceSignatureCheck != *input.EnforceSignatureCheck) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RunAs32Bit != nil && (input.RunAs32Bit == nil || *p.RunAs32Bit != *input.RunAs32Bit) {
		return false
	}

	if p.ScriptContent != nil && (input.ScriptContent == nil || *p.ScriptContent != *input.ScriptContent) {
		return false
	}

	return true
}
