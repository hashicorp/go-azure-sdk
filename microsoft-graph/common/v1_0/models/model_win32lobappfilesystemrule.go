package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRule struct {
	Check32BitOn64System *bool                                   `json:"check32BitOn64System,omitempty"`
	ComparisonValue      *string                                 `json:"comparisonValue,omitempty"`
	FileOrFolderName     *string                                 `json:"fileOrFolderName,omitempty"`
	ODataType            *string                                 `json:"@odata.type,omitempty"`
	OperationType        *Win32LobAppFileSystemRuleOperationType `json:"operationType,omitempty"`
	Operator             *Win32LobAppFileSystemRuleOperator      `json:"operator,omitempty"`
	Path                 *string                                 `json:"path,omitempty"`
	RuleType             *Win32LobAppFileSystemRuleRuleType      `json:"ruleType,omitempty"`
}
