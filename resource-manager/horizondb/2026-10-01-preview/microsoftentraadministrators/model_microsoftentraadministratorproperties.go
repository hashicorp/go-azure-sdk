package microsoftentraadministrators

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftEntraAdministratorProperties struct {
	ObjectId          *string            `json:"objectId,omitempty"`
	PrincipalName     string             `json:"principalName"`
	PrincipalType     PrincipalTypes     `json:"principalType"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	TenantId          *string            `json:"tenantId,omitempty"`
}
