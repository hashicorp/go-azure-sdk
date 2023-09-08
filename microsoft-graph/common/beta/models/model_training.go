package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Training struct {
	AvailabilityStatus   *TrainingAvailabilityStatus `json:"availabilityStatus,omitempty"`
	CreatedBy            *EmailIdentity              `json:"createdBy,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	DisplayName          *string                     `json:"displayName,omitempty"`
	DurationInMinutes    *int64                      `json:"durationInMinutes,omitempty"`
	HasEvaluation        *bool                       `json:"hasEvaluation,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	LanguageDetails      *[]TrainingLanguageDetail   `json:"languageDetails,omitempty"`
	LastModifiedBy       *EmailIdentity              `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Source               *TrainingSource             `json:"source,omitempty"`
	SupportedLocales     *[]string                   `json:"supportedLocales,omitempty"`
	Tags                 *[]string                   `json:"tags,omitempty"`
	Type                 *TrainingType               `json:"type,omitempty"`
}
