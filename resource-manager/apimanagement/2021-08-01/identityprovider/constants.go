package identityprovider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProviderType string

const (
	IdentityProviderTypeAad       IdentityProviderType = "aad"
	IdentityProviderTypeAadBTwoC  IdentityProviderType = "aadB2C"
	IdentityProviderTypeFacebook  IdentityProviderType = "facebook"
	IdentityProviderTypeGoogle    IdentityProviderType = "google"
	IdentityProviderTypeMicrosoft IdentityProviderType = "microsoft"
	IdentityProviderTypeTwitter   IdentityProviderType = "twitter"
)

func PossibleValuesForIdentityProviderType() []string {
	return []string{
		string(IdentityProviderTypeAad),
		string(IdentityProviderTypeAadBTwoC),
		string(IdentityProviderTypeFacebook),
		string(IdentityProviderTypeGoogle),
		string(IdentityProviderTypeMicrosoft),
		string(IdentityProviderTypeTwitter),
	}
}
