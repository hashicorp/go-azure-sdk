package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningContentOperationPredicate struct {
	ContentWebUrl        *string
	CreatedDateTime      *string
	Description          *string
	Duration             *string
	ExternalId           *string
	Format               *string
	Id                   *string
	IsActive             *bool
	IsPremium            *bool
	IsSearchable         *bool
	LanguageTag          *string
	LastModifiedDateTime *string
	NumberOfPages        *int64
	ODataType            *string
	SourceName           *string
	ThumbnailWebUrl      *string
	Title                *string
}

func (p LearningContentOperationPredicate) Matches(input LearningContent) bool {

	if p.ContentWebUrl != nil && (input.ContentWebUrl == nil || *p.ContentWebUrl != *input.ContentWebUrl) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.Format != nil && (input.Format == nil || *p.Format != *input.Format) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsActive != nil && (input.IsActive == nil || *p.IsActive != *input.IsActive) {
		return false
	}

	if p.IsPremium != nil && (input.IsPremium == nil || *p.IsPremium != *input.IsPremium) {
		return false
	}

	if p.IsSearchable != nil && (input.IsSearchable == nil || *p.IsSearchable != *input.IsSearchable) {
		return false
	}

	if p.LanguageTag != nil && (input.LanguageTag == nil || *p.LanguageTag != *input.LanguageTag) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.NumberOfPages != nil && (input.NumberOfPages == nil || *p.NumberOfPages != *input.NumberOfPages) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SourceName != nil && (input.SourceName == nil || *p.SourceName != *input.SourceName) {
		return false
	}

	if p.ThumbnailWebUrl != nil && (input.ThumbnailWebUrl == nil || *p.ThumbnailWebUrl != *input.ThumbnailWebUrl) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
