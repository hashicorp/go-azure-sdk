package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LandingPageDetailOperationPredicate struct {
	Content           *string
	Id                *string
	IsDefaultLangauge *bool
	Language          *string
	ODataType         *string
}

func (p LandingPageDetailOperationPredicate) Matches(input LandingPageDetail) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefaultLangauge != nil && (input.IsDefaultLangauge == nil || *p.IsDefaultLangauge != *input.IsDefaultLangauge) {
		return false
	}

	if p.Language != nil && (input.Language == nil || *p.Language != *input.Language) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
