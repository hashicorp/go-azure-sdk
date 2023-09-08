package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationDetailOperationPredicate struct {
	EmailContent      *string
	Id                *string
	IsDefaultLangauge *bool
	Language          *string
	Locale            *string
	ODataType         *string
	Subject           *string
}

func (p EndUserNotificationDetailOperationPredicate) Matches(input EndUserNotificationDetail) bool {

	if p.EmailContent != nil && (input.EmailContent == nil || *p.EmailContent != *input.EmailContent) {
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

	if p.Locale != nil && (input.Locale == nil || *p.Locale != *input.Locale) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	return true
}
