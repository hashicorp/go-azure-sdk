package summaryrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SummaryLogsProperties struct {
	Description       *string                `json:"description,omitempty"`
	DisplayName       *string                `json:"displayName,omitempty"`
	IsActive          *bool                  `json:"isActive,omitempty"`
	ProvisioningState *ProvisioningStateEnum `json:"provisioningState,omitempty"`
	RuleDefinition    *RuleDefinition        `json:"ruleDefinition,omitempty"`
	RuleType          *RuleTypeEnum          `json:"ruleType,omitempty"`
	StatusCode        *StatusCodeEnum        `json:"statusCode,omitempty"`
}
