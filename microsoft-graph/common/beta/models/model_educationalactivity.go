package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationalActivity struct {
	AllowedAudiences     *EducationalActivityAllowedAudiences `json:"allowedAudiences,omitempty"`
	CompletionMonthYear  *string                              `json:"completionMonthYear,omitempty"`
	CreatedBy            *IdentitySet                         `json:"createdBy,omitempty"`
	CreatedDateTime      *string                              `json:"createdDateTime,omitempty"`
	EndMonthYear         *string                              `json:"endMonthYear,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	Inference            *InferenceData                       `json:"inference,omitempty"`
	Institution          *InstitutionData                     `json:"institution,omitempty"`
	IsSearchable         *bool                                `json:"isSearchable,omitempty"`
	LastModifiedBy       *IdentitySet                         `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                              `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	Program              *EducationalActivityDetail           `json:"program,omitempty"`
	Source               *PersonDataSources                   `json:"source,omitempty"`
	StartMonthYear       *string                              `json:"startMonthYear,omitempty"`
}
