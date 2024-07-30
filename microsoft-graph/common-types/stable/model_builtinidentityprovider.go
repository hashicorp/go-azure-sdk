package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuiltInIdentityProvider struct {
	DisplayName          *string `json:"displayName,omitempty"`
	Id                   *string `json:"id,omitempty"`
	IdentityProviderType *string `json:"identityProviderType,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
