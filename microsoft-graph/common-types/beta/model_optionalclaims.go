package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OptionalClaims struct {
	AccessToken *[]OptionalClaim `json:"accessToken,omitempty"`
	IdToken     *[]OptionalClaim `json:"idToken,omitempty"`
	ODataType   *string          `json:"@odata.type,omitempty"`
	Saml2Token  *[]OptionalClaim `json:"saml2Token,omitempty"`
}
