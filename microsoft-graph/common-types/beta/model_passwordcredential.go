package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordCredential struct {
	CustomKeyIdentifier *string `json:"customKeyIdentifier,omitempty"`
	DisplayName         *string `json:"displayName,omitempty"`
	EndDateTime         *string `json:"endDateTime,omitempty"`
	Hint                *string `json:"hint,omitempty"`
	KeyId               *string `json:"keyId,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	SecretText          *string `json:"secretText,omitempty"`
	StartDateTime       *string `json:"startDateTime,omitempty"`
}
