package costallocationrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostAllocationRuleProperties struct {
	CreatedDate *string                   `json:"createdDate,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Details     CostAllocationRuleDetails `json:"details"`
	Status      RuleStatus                `json:"status"`
	UpdatedDate *string                   `json:"updatedDate,omitempty"`
}
