package policytokens

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyLogInfo struct {
	Ancestors                      *string   `json:"ancestors,omitempty"`
	ComplianceReasonCode           *string   `json:"complianceReasonCode,omitempty"`
	PolicyAssignmentDisplayName    *string   `json:"policyAssignmentDisplayName,omitempty"`
	PolicyAssignmentId             *string   `json:"policyAssignmentId,omitempty"`
	PolicyAssignmentName           *string   `json:"policyAssignmentName,omitempty"`
	PolicyAssignmentScope          *string   `json:"policyAssignmentScope,omitempty"`
	PolicyAssignmentVersion        *string   `json:"policyAssignmentVersion,omitempty"`
	PolicyDefinitionDisplayName    *string   `json:"policyDefinitionDisplayName,omitempty"`
	PolicyDefinitionEffect         *string   `json:"policyDefinitionEffect,omitempty"`
	PolicyDefinitionGroupNames     *[]string `json:"policyDefinitionGroupNames,omitempty"`
	PolicyDefinitionId             *string   `json:"policyDefinitionId,omitempty"`
	PolicyDefinitionName           *string   `json:"policyDefinitionName,omitempty"`
	PolicyDefinitionReferenceId    *string   `json:"policyDefinitionReferenceId,omitempty"`
	PolicyDefinitionVersion        *string   `json:"policyDefinitionVersion,omitempty"`
	PolicyExemptionIds             *[]string `json:"policyExemptionIds,omitempty"`
	PolicySetDefinitionCategory    *string   `json:"policySetDefinitionCategory,omitempty"`
	PolicySetDefinitionDisplayName *string   `json:"policySetDefinitionDisplayName,omitempty"`
	PolicySetDefinitionId          *string   `json:"policySetDefinitionId,omitempty"`
	PolicySetDefinitionName        *string   `json:"policySetDefinitionName,omitempty"`
	PolicySetDefinitionVersion     *string   `json:"policySetDefinitionVersion,omitempty"`
	ResourceLocation               *string   `json:"resourceLocation,omitempty"`
}
