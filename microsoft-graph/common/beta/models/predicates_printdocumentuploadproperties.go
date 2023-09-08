package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintDocumentUploadPropertiesOperationPredicate struct {
	ContentType  *string
	DocumentName *string
	ODataType    *string
	Size         *int64
}

func (p PrintDocumentUploadPropertiesOperationPredicate) Matches(input PrintDocumentUploadProperties) bool {

	if p.ContentType != nil && (input.ContentType == nil || *p.ContentType != *input.ContentType) {
		return false
	}

	if p.DocumentName != nil && (input.DocumentName == nil || *p.DocumentName != *input.DocumentName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	return true
}
