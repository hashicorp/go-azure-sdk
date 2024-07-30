package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityWhoisContact struct {
	Address      *PhysicalAddress `json:"address,omitempty"`
	Email        *string          `json:"email,omitempty"`
	Fax          *string          `json:"fax,omitempty"`
	Name         *string          `json:"name,omitempty"`
	ODataType    *string          `json:"@odata.type,omitempty"`
	Organization *string          `json:"organization,omitempty"`
	Telephone    *string          `json:"telephone,omitempty"`
}
