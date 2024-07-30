package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FederatedIdentityCredential struct {
	Audiences   *[]string `json:"audiences,omitempty"`
	Description *string   `json:"description,omitempty"`
	Id          *string   `json:"id,omitempty"`
	Issuer      *string   `json:"issuer,omitempty"`
	Name        *string   `json:"name,omitempty"`
	ODataType   *string   `json:"@odata.type,omitempty"`
	Subject     *string   `json:"subject,omitempty"`
}
