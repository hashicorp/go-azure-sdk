package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionDeploymentStatusOperationPredicate struct {
	ManagementActionId        *string
	ManagementTemplateId      *string
	ManagementTemplateVersion *int64
	ODataType                 *string
}

func (p ManagedTenantsManagementActionDeploymentStatusOperationPredicate) Matches(input ManagedTenantsManagementActionDeploymentStatus) bool {

	if p.ManagementActionId != nil && (input.ManagementActionId == nil || *p.ManagementActionId != *input.ManagementActionId) {
		return false
	}

	if p.ManagementTemplateId != nil && (input.ManagementTemplateId == nil || *p.ManagementTemplateId != *input.ManagementTemplateId) {
		return false
	}

	if p.ManagementTemplateVersion != nil && (input.ManagementTemplateVersion == nil || *p.ManagementTemplateVersion != *input.ManagementTemplateVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
