package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonName struct {
	AllowedAudiences     *PersonNameAllowedAudiences `json:"allowedAudiences,omitempty"`
	CreatedBy            *IdentitySet                `json:"createdBy,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	DisplayName          *string                     `json:"displayName,omitempty"`
	First                *string                     `json:"first,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Inference            *InferenceData              `json:"inference,omitempty"`
	Initials             *string                     `json:"initials,omitempty"`
	IsSearchable         *bool                       `json:"isSearchable,omitempty"`
	LanguageTag          *string                     `json:"languageTag,omitempty"`
	Last                 *string                     `json:"last,omitempty"`
	LastModifiedBy       *IdentitySet                `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	Maiden               *string                     `json:"maiden,omitempty"`
	Middle               *string                     `json:"middle,omitempty"`
	Nickname             *string                     `json:"nickname,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Pronunciation        *PersonNamePronounciation   `json:"pronunciation,omitempty"`
	Source               *PersonDataSources          `json:"source,omitempty"`
	Suffix               *string                     `json:"suffix,omitempty"`
	Title                *string                     `json:"title,omitempty"`
}
