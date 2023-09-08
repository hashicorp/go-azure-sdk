package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPreviewInfoOperationPredicate struct {
	GetUrl         *string
	ODataType      *string
	PostParameters *string
	PostUrl        *string
}

func (p ItemPreviewInfoOperationPredicate) Matches(input ItemPreviewInfo) bool {

	if p.GetUrl != nil && (input.GetUrl == nil || *p.GetUrl != *input.GetUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PostParameters != nil && (input.PostParameters == nil || *p.PostParameters != *input.PostParameters) {
		return false
	}

	if p.PostUrl != nil && (input.PostUrl == nil || *p.PostUrl != *input.PostUrl) {
		return false
	}

	return true
}
