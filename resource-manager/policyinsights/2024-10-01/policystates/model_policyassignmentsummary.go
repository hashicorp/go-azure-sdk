package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentSummary struct {
	PolicyAssignmentId    *string                    `json:"policyAssignmentId,omitempty"`
	PolicyDefinitions     *[]PolicyDefinitionSummary `json:"policyDefinitions,omitempty"`
	PolicyGroups          *[]PolicyGroupSummary      `json:"policyGroups,omitempty"`
	PolicySetDefinitionId *string                    `json:"policySetDefinitionId,omitempty"`
	Results               *SummaryResults            `json:"results,omitempty"`
}
