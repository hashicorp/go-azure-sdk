package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemFacet struct {
	AllowedAudiences     *ItemFacetAllowedAudiences `json:"allowedAudiences,omitempty"`
	CreatedBy            *IdentitySet               `json:"createdBy,omitempty"`
	CreatedDateTime      *string                    `json:"createdDateTime,omitempty"`
	Id                   *string                    `json:"id,omitempty"`
	Inference            *InferenceData             `json:"inference,omitempty"`
	IsSearchable         *bool                      `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet               `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                    `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                    `json:"@odata.type,omitempty"`
	Source               *PersonDataSources         `json:"source,omitempty"`
}
