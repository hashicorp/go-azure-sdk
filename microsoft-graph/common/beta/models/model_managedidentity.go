package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIdentity struct {
	AssociatedResourceId *string                 `json:"associatedResourceId,omitempty"`
	FederatedTokenId     *string                 `json:"federatedTokenId,omitempty"`
	FederatedTokenIssuer *string                 `json:"federatedTokenIssuer,omitempty"`
	MsiType              *ManagedIdentityMsiType `json:"msiType,omitempty"`
	ODataType            *string                 `json:"@odata.type,omitempty"`
}
