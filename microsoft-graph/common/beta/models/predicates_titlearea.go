package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleAreaOperationPredicate struct {
	AlternativeText         *string
	EnableGradientEffect    *bool
	ImageWebUrl             *string
	ODataType               *string
	ShowAuthor              *bool
	ShowPublishedDate       *bool
	ShowTextBlockAboveTitle *bool
	TextAboveTitle          *string
}

func (p TitleAreaOperationPredicate) Matches(input TitleArea) bool {

	if p.AlternativeText != nil && (input.AlternativeText == nil || *p.AlternativeText != *input.AlternativeText) {
		return false
	}

	if p.EnableGradientEffect != nil && (input.EnableGradientEffect == nil || *p.EnableGradientEffect != *input.EnableGradientEffect) {
		return false
	}

	if p.ImageWebUrl != nil && (input.ImageWebUrl == nil || *p.ImageWebUrl != *input.ImageWebUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ShowAuthor != nil && (input.ShowAuthor == nil || *p.ShowAuthor != *input.ShowAuthor) {
		return false
	}

	if p.ShowPublishedDate != nil && (input.ShowPublishedDate == nil || *p.ShowPublishedDate != *input.ShowPublishedDate) {
		return false
	}

	if p.ShowTextBlockAboveTitle != nil && (input.ShowTextBlockAboveTitle == nil || *p.ShowTextBlockAboveTitle != *input.ShowTextBlockAboveTitle) {
		return false
	}

	if p.TextAboveTitle != nil && (input.TextAboveTitle == nil || *p.TextAboveTitle != *input.TextAboveTitle) {
		return false
	}

	return true
}
