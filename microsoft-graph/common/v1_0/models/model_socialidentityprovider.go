package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SocialIdentityProvider struct {
	ClientId             *string `json:"clientId,omitempty"`
	ClientSecret         *string `json:"clientSecret,omitempty"`
	DisplayName          *string `json:"displayName,omitempty"`
	Id                   *string `json:"id,omitempty"`
	IdentityProviderType *string `json:"identityProviderType,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
