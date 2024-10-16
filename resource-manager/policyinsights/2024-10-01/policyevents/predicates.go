package policyevents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyEventOperationPredicate struct {
	ComplianceState               *string
	EffectiveParameters           *string
	IsCompliant                   *bool
	ManagementGroupIds            *string
	OdataContext                  *string
	OdataId                       *string
	PolicyAssignmentId            *string
	PolicyAssignmentName          *string
	PolicyAssignmentOwner         *string
	PolicyAssignmentParameters    *string
	PolicyAssignmentScope         *string
	PolicyDefinitionAction        *string
	PolicyDefinitionCategory      *string
	PolicyDefinitionId            *string
	PolicyDefinitionName          *string
	PolicyDefinitionReferenceId   *string
	PolicySetDefinitionCategory   *string
	PolicySetDefinitionId         *string
	PolicySetDefinitionName       *string
	PolicySetDefinitionOwner      *string
	PolicySetDefinitionParameters *string
	PrincipalOid                  *string
	ResourceGroup                 *string
	ResourceId                    *string
	ResourceLocation              *string
	ResourceTags                  *string
	ResourceType                  *string
	SubscriptionId                *string
	TenantId                      *string
	Timestamp                     *string
}

func (p PolicyEventOperationPredicate) Matches(input PolicyEvent) bool {

	if p.ComplianceState != nil && (input.ComplianceState == nil || *p.ComplianceState != *input.ComplianceState) {
		return false
	}

	if p.EffectiveParameters != nil && (input.EffectiveParameters == nil || *p.EffectiveParameters != *input.EffectiveParameters) {
		return false
	}

	if p.IsCompliant != nil && (input.IsCompliant == nil || *p.IsCompliant != *input.IsCompliant) {
		return false
	}

	if p.ManagementGroupIds != nil && (input.ManagementGroupIds == nil || *p.ManagementGroupIds != *input.ManagementGroupIds) {
		return false
	}

	if p.OdataContext != nil && (input.OdataContext == nil || *p.OdataContext != *input.OdataContext) {
		return false
	}

	if p.OdataId != nil && (input.OdataId == nil || *p.OdataId != *input.OdataId) {
		return false
	}

	if p.PolicyAssignmentId != nil && (input.PolicyAssignmentId == nil || *p.PolicyAssignmentId != *input.PolicyAssignmentId) {
		return false
	}

	if p.PolicyAssignmentName != nil && (input.PolicyAssignmentName == nil || *p.PolicyAssignmentName != *input.PolicyAssignmentName) {
		return false
	}

	if p.PolicyAssignmentOwner != nil && (input.PolicyAssignmentOwner == nil || *p.PolicyAssignmentOwner != *input.PolicyAssignmentOwner) {
		return false
	}

	if p.PolicyAssignmentParameters != nil && (input.PolicyAssignmentParameters == nil || *p.PolicyAssignmentParameters != *input.PolicyAssignmentParameters) {
		return false
	}

	if p.PolicyAssignmentScope != nil && (input.PolicyAssignmentScope == nil || *p.PolicyAssignmentScope != *input.PolicyAssignmentScope) {
		return false
	}

	if p.PolicyDefinitionAction != nil && (input.PolicyDefinitionAction == nil || *p.PolicyDefinitionAction != *input.PolicyDefinitionAction) {
		return false
	}

	if p.PolicyDefinitionCategory != nil && (input.PolicyDefinitionCategory == nil || *p.PolicyDefinitionCategory != *input.PolicyDefinitionCategory) {
		return false
	}

	if p.PolicyDefinitionId != nil && (input.PolicyDefinitionId == nil || *p.PolicyDefinitionId != *input.PolicyDefinitionId) {
		return false
	}

	if p.PolicyDefinitionName != nil && (input.PolicyDefinitionName == nil || *p.PolicyDefinitionName != *input.PolicyDefinitionName) {
		return false
	}

	if p.PolicyDefinitionReferenceId != nil && (input.PolicyDefinitionReferenceId == nil || *p.PolicyDefinitionReferenceId != *input.PolicyDefinitionReferenceId) {
		return false
	}

	if p.PolicySetDefinitionCategory != nil && (input.PolicySetDefinitionCategory == nil || *p.PolicySetDefinitionCategory != *input.PolicySetDefinitionCategory) {
		return false
	}

	if p.PolicySetDefinitionId != nil && (input.PolicySetDefinitionId == nil || *p.PolicySetDefinitionId != *input.PolicySetDefinitionId) {
		return false
	}

	if p.PolicySetDefinitionName != nil && (input.PolicySetDefinitionName == nil || *p.PolicySetDefinitionName != *input.PolicySetDefinitionName) {
		return false
	}

	if p.PolicySetDefinitionOwner != nil && (input.PolicySetDefinitionOwner == nil || *p.PolicySetDefinitionOwner != *input.PolicySetDefinitionOwner) {
		return false
	}

	if p.PolicySetDefinitionParameters != nil && (input.PolicySetDefinitionParameters == nil || *p.PolicySetDefinitionParameters != *input.PolicySetDefinitionParameters) {
		return false
	}

	if p.PrincipalOid != nil && (input.PrincipalOid == nil || *p.PrincipalOid != *input.PrincipalOid) {
		return false
	}

	if p.ResourceGroup != nil && (input.ResourceGroup == nil || *p.ResourceGroup != *input.ResourceGroup) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	if p.ResourceLocation != nil && (input.ResourceLocation == nil || *p.ResourceLocation != *input.ResourceLocation) {
		return false
	}

	if p.ResourceTags != nil && (input.ResourceTags == nil || *p.ResourceTags != *input.ResourceTags) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil || *p.ResourceType != *input.ResourceType) {
		return false
	}

	if p.SubscriptionId != nil && (input.SubscriptionId == nil || *p.SubscriptionId != *input.SubscriptionId) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.Timestamp != nil && (input.Timestamp == nil || *p.Timestamp != *input.Timestamp) {
		return false
	}

	return true
}
