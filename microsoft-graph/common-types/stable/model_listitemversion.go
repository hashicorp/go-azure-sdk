package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListItemVersion struct {
	Fields               *FieldValueSet    `json:"fields,omitempty"`
	Id                   *string           `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet      `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string           `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string           `json:"@odata.type,omitempty"`
	Publication          *PublicationFacet `json:"publication,omitempty"`
}
