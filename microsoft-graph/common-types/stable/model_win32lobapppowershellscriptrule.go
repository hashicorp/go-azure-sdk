package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRule struct {
	ComparisonValue       *string                                       `json:"comparisonValue,omitempty"`
	DisplayName           *string                                       `json:"displayName,omitempty"`
	EnforceSignatureCheck *bool                                         `json:"enforceSignatureCheck,omitempty"`
	ODataType             *string                                       `json:"@odata.type,omitempty"`
	OperationType         *Win32LobAppPowerShellScriptRuleOperationType `json:"operationType,omitempty"`
	Operator              *Win32LobAppPowerShellScriptRuleOperator      `json:"operator,omitempty"`
	RuleType              *Win32LobAppPowerShellScriptRuleRuleType      `json:"ruleType,omitempty"`
	RunAs32Bit            *bool                                         `json:"runAs32Bit,omitempty"`
	RunAsAccount          *Win32LobAppPowerShellScriptRuleRunAsAccount  `json:"runAsAccount,omitempty"`
	ScriptContent         *string                                       `json:"scriptContent,omitempty"`
}
