package billingroleassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRoleAssignmentProperties struct {
	CreatedByPrincipalId       *string `json:"createdByPrincipalId,omitempty"`
	CreatedByPrincipalTenantId *string `json:"createdByPrincipalTenantId,omitempty"`
	CreatedByUserEmailAddress  *string `json:"createdByUserEmailAddress,omitempty"`
	CreatedOn                  *string `json:"createdOn,omitempty"`
	Name                       *string `json:"name,omitempty"`
	PrincipalId                *string `json:"principalId,omitempty"`
	PrincipalTenantId          *string `json:"principalTenantId,omitempty"`
	RoleDefinitionId           *string `json:"roleDefinitionId,omitempty"`
	Scope                      *string `json:"scope,omitempty"`
	UserAuthenticationType     *string `json:"userAuthenticationType,omitempty"`
	UserEmailAddress           *string `json:"userEmailAddress,omitempty"`
}
