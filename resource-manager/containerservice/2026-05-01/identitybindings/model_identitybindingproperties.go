package identitybindings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBindingProperties struct {
	ManagedIdentity   IdentityBindingManagedIdentityProfile `json:"managedIdentity"`
	OidcIssuer        *IdentityBindingOidcIssuerProfile     `json:"oidcIssuer,omitempty"`
	ProvisioningState *IdentityBindingProvisioningState     `json:"provisioningState,omitempty"`
}
