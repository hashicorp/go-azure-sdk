package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedCredentialData struct {
	Authority *string                   `json:"authority,omitempty"`
	Claims    *VerifiedCredentialClaims `json:"claims,omitempty"`
	ODataType *string                   `json:"@odata.type,omitempty"`
	Type      *[]string                 `json:"type,omitempty"`
}
