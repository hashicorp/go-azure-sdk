package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonCertification struct {
	AllowedAudiences     *PersonCertificationAllowedAudiences `json:"allowedAudiences,omitempty"`
	CertificationId      *string                              `json:"certificationId,omitempty"`
	CreatedBy            *IdentitySet                         `json:"createdBy,omitempty"`
	CreatedDateTime      *string                              `json:"createdDateTime,omitempty"`
	Description          *string                              `json:"description,omitempty"`
	DisplayName          *string                              `json:"displayName,omitempty"`
	EndDate              *string                              `json:"endDate,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	Inference            *InferenceData                       `json:"inference,omitempty"`
	IsSearchable         *bool                                `json:"isSearchable,omitempty"`
	IssuedDate           *string                              `json:"issuedDate,omitempty"`
	IssuingAuthority     *string                              `json:"issuingAuthority,omitempty"`
	IssuingCompany       *string                              `json:"issuingCompany,omitempty"`
	LastModifiedBy       *IdentitySet                         `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                              `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	Source               *PersonDataSources                   `json:"source,omitempty"`
	StartDate            *string                              `json:"startDate,omitempty"`
	ThumbnailUrl         *string                              `json:"thumbnailUrl,omitempty"`
	WebUrl               *string                              `json:"webUrl,omitempty"`
}
