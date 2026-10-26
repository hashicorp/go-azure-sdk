package identitybindings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBindingManagedIdentityProfile struct {
	ClientId   *string `json:"clientId,omitempty"`
	ObjectId   *string `json:"objectId,omitempty"`
	ResourceId string  `json:"resourceId"`
	TenantId   *string `json:"tenantId,omitempty"`
}
