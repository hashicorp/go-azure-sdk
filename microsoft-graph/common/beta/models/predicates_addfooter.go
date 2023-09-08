package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddFooterOperationPredicate struct {
	FontColor *string
	FontSize  *int64
	Margin    *int64
	Name      *string
	ODataType *string
	Text      *string
}

func (p AddFooterOperationPredicate) Matches(input AddFooter) bool {

	if p.FontColor != nil && (input.FontColor == nil || *p.FontColor != *input.FontColor) {
		return false
	}

	if p.FontSize != nil && (input.FontSize == nil || *p.FontSize != *input.FontSize) {
		return false
	}

	if p.Margin != nil && (input.Margin == nil || *p.Margin != *input.Margin) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Text != nil && (input.Text == nil || *p.Text != *input.Text) {
		return false
	}

	return true
}
