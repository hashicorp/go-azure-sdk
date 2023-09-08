package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddContentHeaderActionOperationPredicate struct {
	FontColor     *string
	FontName      *string
	FontSize      *int64
	Margin        *int64
	ODataType     *string
	Text          *string
	UiElementName *string
}

func (p SecurityAddContentHeaderActionOperationPredicate) Matches(input SecurityAddContentHeaderAction) bool {

	if p.FontColor != nil && (input.FontColor == nil || *p.FontColor != *input.FontColor) {
		return false
	}

	if p.FontName != nil && (input.FontName == nil || *p.FontName != *input.FontName) {
		return false
	}

	if p.FontSize != nil && (input.FontSize == nil || *p.FontSize != *input.FontSize) {
		return false
	}

	if p.Margin != nil && (input.Margin == nil || *p.Margin != *input.Margin) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Text != nil && (input.Text == nil || *p.Text != *input.Text) {
		return false
	}

	if p.UiElementName != nil && (input.UiElementName == nil || *p.UiElementName != *input.UiElementName) {
		return false
	}

	return true
}
