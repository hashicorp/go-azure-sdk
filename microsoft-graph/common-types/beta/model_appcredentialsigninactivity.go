package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppCredentialSignInActivity struct {
	AppId                    *string                                      `json:"appId,omitempty"`
	AppObjectId              *string                                      `json:"appObjectId,omitempty"`
	CreatedDateTime          *string                                      `json:"createdDateTime,omitempty"`
	CredentialOrigin         *AppCredentialSignInActivityCredentialOrigin `json:"credentialOrigin,omitempty"`
	ExpirationDateTime       *string                                      `json:"expirationDateTime,omitempty"`
	Id                       *string                                      `json:"id,omitempty"`
	KeyId                    *string                                      `json:"keyId,omitempty"`
	KeyType                  *AppCredentialSignInActivityKeyType          `json:"keyType,omitempty"`
	KeyUsage                 *AppCredentialSignInActivityKeyUsage         `json:"keyUsage,omitempty"`
	ODataType                *string                                      `json:"@odata.type,omitempty"`
	ResourceId               *string                                      `json:"resourceId,omitempty"`
	ServicePrincipalObjectId *string                                      `json:"servicePrincipalObjectId,omitempty"`
	SignInActivity           *SignInActivity                              `json:"signInActivity,omitempty"`
}
