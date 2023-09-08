package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateCatalogItemOperationPredicate struct {
	DisplayName      *string
	EndOfSupportDate *string
	Id               *string
	IsExpeditable    *bool
	KbArticleId      *string
	ODataType        *string
	ReleaseDateTime  *string
}

func (p WindowsQualityUpdateCatalogItemOperationPredicate) Matches(input WindowsQualityUpdateCatalogItem) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EndOfSupportDate != nil && (input.EndOfSupportDate == nil || *p.EndOfSupportDate != *input.EndOfSupportDate) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsExpeditable != nil && (input.IsExpeditable == nil || *p.IsExpeditable != *input.IsExpeditable) {
		return false
	}

	if p.KbArticleId != nil && (input.KbArticleId == nil || *p.KbArticleId != *input.KbArticleId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReleaseDateTime != nil && (input.ReleaseDateTime == nil || *p.ReleaseDateTime != *input.ReleaseDateTime) {
		return false
	}

	return true
}
