package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePagePreviewOperationPredicate struct {
	ODataType   *string
	PreviewText *string
}

func (p OnenotePagePreviewOperationPredicate) Matches(input OnenotePagePreview) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PreviewText != nil && (input.PreviewText == nil || *p.PreviewText != *input.PreviewText) {
		return false
	}

	return true
}
