package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialType struct {
	CredentialType *string   `json:"credentialType,omitempty"`
	Issuers        *[]string `json:"issuers,omitempty"`
	ODataType      *string   `json:"@odata.type,omitempty"`
}
