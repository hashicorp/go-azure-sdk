package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRule struct {
	Check32BitOn64System *bool                                 `json:"check32BitOn64System,omitempty"`
	ComparisonValue      *string                               `json:"comparisonValue,omitempty"`
	KeyPath              *string                               `json:"keyPath,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	OperationType        *Win32LobAppRegistryRuleOperationType `json:"operationType,omitempty"`
	Operator             *Win32LobAppRegistryRuleOperator      `json:"operator,omitempty"`
	RuleType             *Win32LobAppRegistryRuleRuleType      `json:"ruleType,omitempty"`
	ValueName            *string                               `json:"valueName,omitempty"`
}
