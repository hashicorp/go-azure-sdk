package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccountInformation struct {
	AgeGroup             *string                                 `json:"ageGroup,omitempty"`
	AllowedAudiences     *UserAccountInformationAllowedAudiences `json:"allowedAudiences,omitempty"`
	CountryCode          *string                                 `json:"countryCode,omitempty"`
	CreatedBy            *IdentitySet                            `json:"createdBy,omitempty"`
	CreatedDateTime      *string                                 `json:"createdDateTime,omitempty"`
	Id                   *string                                 `json:"id,omitempty"`
	Inference            *InferenceData                          `json:"inference,omitempty"`
	IsSearchable         *bool                                   `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet                            `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                                 `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                 `json:"@odata.type,omitempty"`
	PreferredLanguageTag *LocaleInfo                             `json:"preferredLanguageTag,omitempty"`
	Source               *PersonDataSources                      `json:"source,omitempty"`
	UserPrincipalName    *string                                 `json:"userPrincipalName,omitempty"`
}
