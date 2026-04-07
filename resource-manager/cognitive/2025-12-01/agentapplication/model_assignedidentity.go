package agentapplication

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedIdentity struct {
	ClientId          string                     `json:"clientId"`
	Kind              IdentityKind               `json:"kind"`
	PrincipalId       string                     `json:"principalId"`
	ProvisioningState *IdentityProvisioningState `json:"provisioningState,omitempty"`
	Subject           *string                    `json:"subject,omitempty"`
	TenantId          string                     `json:"tenantId"`
	Type              IdentityManagementType     `json:"type"`
}
