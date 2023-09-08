package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialSettings struct {
	CredentialTypes *[]VerifiableCredentialType `json:"credentialTypes,omitempty"`
	ODataType       *string                     `json:"@odata.type,omitempty"`
}
