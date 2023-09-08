package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileOperationPredicate struct {
	MimeType           *string
	ODataType          *string
	ProcessingMetadata *bool
}

func (p FileOperationPredicate) Matches(input File) bool {

	if p.MimeType != nil && (input.MimeType == nil || *p.MimeType != *input.MimeType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProcessingMetadata != nil && (input.ProcessingMetadata == nil || *p.ProcessingMetadata != *input.ProcessingMetadata) {
		return false
	}

	return true
}
