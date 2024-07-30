package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAssociatedIdentities struct {
	All               *[]AzureIdentity         `json:"all,omitempty"`
	ManagedIdentities *[]AzureManagedIdentity  `json:"managedIdentities,omitempty"`
	ODataType         *string                  `json:"@odata.type,omitempty"`
	ServicePrincipals *[]AzureServicePrincipal `json:"servicePrincipals,omitempty"`
	Users             *[]AzureUser             `json:"users,omitempty"`
}
