package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirement struct {
	DetectionType         *Win32LobAppPowerShellScriptRequirementDetectionType `json:"detectionType,omitempty"`
	DetectionValue        *string                                              `json:"detectionValue,omitempty"`
	DisplayName           *string                                              `json:"displayName,omitempty"`
	EnforceSignatureCheck *bool                                                `json:"enforceSignatureCheck,omitempty"`
	ODataType             *string                                              `json:"@odata.type,omitempty"`
	Operator              *Win32LobAppPowerShellScriptRequirementOperator      `json:"operator,omitempty"`
	RunAs32Bit            *bool                                                `json:"runAs32Bit,omitempty"`
	RunAsAccount          *Win32LobAppPowerShellScriptRequirementRunAsAccount  `json:"runAsAccount,omitempty"`
	ScriptContent         *string                                              `json:"scriptContent,omitempty"`
}
