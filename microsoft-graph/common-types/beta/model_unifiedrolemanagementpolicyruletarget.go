package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyRuleTarget struct {
	Caller              *string            `json:"caller,omitempty"`
	EnforcedSettings    *[]string          `json:"enforcedSettings,omitempty"`
	InheritableSettings *[]string          `json:"inheritableSettings,omitempty"`
	Level               *string            `json:"level,omitempty"`
	ODataType           *string            `json:"@odata.type,omitempty"`
	Operations          *[]string          `json:"operations,omitempty"`
	TargetObjects       *[]DirectoryObject `json:"targetObjects,omitempty"`
}
