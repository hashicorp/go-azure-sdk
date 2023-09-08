package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Shared struct {
	ODataType      *string      `json:"@odata.type,omitempty"`
	Owner          *IdentitySet `json:"owner,omitempty"`
	Scope          *string      `json:"scope,omitempty"`
	SharedBy       *IdentitySet `json:"sharedBy,omitempty"`
	SharedDateTime *string      `json:"sharedDateTime,omitempty"`
}
