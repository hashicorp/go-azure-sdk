package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyEnablementRule struct {
	EnabledRules *[]string                              `json:"enabledRules,omitempty"`
	Id           *string                                `json:"id,omitempty"`
	ODataType    *string                                `json:"@odata.type,omitempty"`
	Target       *UnifiedRoleManagementPolicyRuleTarget `json:"target,omitempty"`
}
