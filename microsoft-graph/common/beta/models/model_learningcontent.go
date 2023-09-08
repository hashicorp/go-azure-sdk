package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningContent struct {
	AdditionalTags       *[]string `json:"additionalTags,omitempty"`
	ContentWebUrl        *string   `json:"contentWebUrl,omitempty"`
	Contributors         *[]string `json:"contributors,omitempty"`
	CreatedDateTime      *string   `json:"createdDateTime,omitempty"`
	Description          *string   `json:"description,omitempty"`
	Duration             *string   `json:"duration,omitempty"`
	ExternalId           *string   `json:"externalId,omitempty"`
	Format               *string   `json:"format,omitempty"`
	Id                   *string   `json:"id,omitempty"`
	IsActive             *bool     `json:"isActive,omitempty"`
	IsPremium            *bool     `json:"isPremium,omitempty"`
	IsSearchable         *bool     `json:"isSearchable,omitempty"`
	LanguageTag          *string   `json:"languageTag,omitempty"`
	LastModifiedDateTime *string   `json:"lastModifiedDateTime,omitempty"`
	NumberOfPages        *int64    `json:"numberOfPages,omitempty"`
	ODataType            *string   `json:"@odata.type,omitempty"`
	SkillTags            *[]string `json:"skillTags,omitempty"`
	SourceName           *string   `json:"sourceName,omitempty"`
	ThumbnailWebUrl      *string   `json:"thumbnailWebUrl,omitempty"`
	Title                *string   `json:"title,omitempty"`
}
