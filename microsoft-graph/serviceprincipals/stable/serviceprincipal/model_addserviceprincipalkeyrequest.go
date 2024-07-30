package serviceprincipal

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddServicePrincipalKeyRequest struct {
	KeyCredential      *KeyCredential      `json:"keyCredential,omitempty"`
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
	Proof              *string             `json:"proof,omitempty"`
}
