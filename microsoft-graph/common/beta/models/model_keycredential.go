package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyCredential struct {
	CustomKeyIdentifier *string `json:"customKeyIdentifier,omitempty"`
	DisplayName         *string `json:"displayName,omitempty"`
	EndDateTime         *string `json:"endDateTime,omitempty"`
	Key                 *string `json:"key,omitempty"`
	KeyId               *string `json:"keyId,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	StartDateTime       *string `json:"startDateTime,omitempty"`
	Type                *string `json:"type,omitempty"`
	Usage               *string `json:"usage,omitempty"`
}
