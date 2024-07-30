package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPatent struct {
	AllowedAudiences     *ItemPatentAllowedAudiences `json:"allowedAudiences,omitempty"`
	CreatedBy            *IdentitySet                `json:"createdBy,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	DisplayName          *string                     `json:"displayName,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Inference            *InferenceData              `json:"inference,omitempty"`
	IsPending            *bool                       `json:"isPending,omitempty"`
	IsSearchable         *bool                       `json:"isSearchable,omitempty"`
	IssuedDate           *string                     `json:"issuedDate,omitempty"`
	IssuingAuthority     *string                     `json:"issuingAuthority,omitempty"`
	LastModifiedBy       *IdentitySet                `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	Number               *string                     `json:"number,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Source               *PersonDataSources          `json:"source,omitempty"`
	WebUrl               *string                     `json:"webUrl,omitempty"`
}
