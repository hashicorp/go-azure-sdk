package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiency struct {
	AllowedAudiences     *LanguageProficiencyAllowedAudiences `json:"allowedAudiences,omitempty"`
	CreatedBy            *IdentitySet                         `json:"createdBy,omitempty"`
	CreatedDateTime      *string                              `json:"createdDateTime,omitempty"`
	DisplayName          *string                              `json:"displayName,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	Inference            *InferenceData                       `json:"inference,omitempty"`
	IsSearchable         *bool                                `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet                         `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                              `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	Proficiency          *LanguageProficiencyProficiency      `json:"proficiency,omitempty"`
	Reading              *LanguageProficiencyReading          `json:"reading,omitempty"`
	Source               *PersonDataSources                   `json:"source,omitempty"`
	Spoken               *LanguageProficiencySpoken           `json:"spoken,omitempty"`
	Tag                  *string                              `json:"tag,omitempty"`
	ThumbnailUrl         *string                              `json:"thumbnailUrl,omitempty"`
	Written              *LanguageProficiencyWritten          `json:"written,omitempty"`
}
