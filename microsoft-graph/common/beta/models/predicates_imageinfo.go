package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageInfoOperationPredicate struct {
	AddImageQuery   *bool
	AlternateText   *string
	AlternativeText *string
	IconUrl         *string
	ODataType       *string
}

func (p ImageInfoOperationPredicate) Matches(input ImageInfo) bool {

	if p.AddImageQuery != nil && (input.AddImageQuery == nil || *p.AddImageQuery != *input.AddImageQuery) {
		return false
	}

	if p.AlternateText != nil && (input.AlternateText == nil || *p.AlternateText != *input.AlternateText) {
		return false
	}

	if p.AlternativeText != nil && (input.AlternativeText == nil || *p.AlternativeText != *input.AlternativeText) {
		return false
	}

	if p.IconUrl != nil && (input.IconUrl == nil || *p.IconUrl != *input.IconUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
