package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextClassificationRequestOperationPredicate struct {
	FileExtension *string
	Id            *string
	ODataType     *string
	Text          *string
}

func (p TextClassificationRequestOperationPredicate) Matches(input TextClassificationRequest) bool {

	if p.FileExtension != nil && (input.FileExtension == nil || *p.FileExtension != *input.FileExtension) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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
