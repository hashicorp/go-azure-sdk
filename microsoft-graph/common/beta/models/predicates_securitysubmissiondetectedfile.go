package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionDetectedFileOperationPredicate struct {
	FileHash  *string
	FileName  *string
	ODataType *string
}

func (p SecuritySubmissionDetectedFileOperationPredicate) Matches(input SecuritySubmissionDetectedFile) bool {

	if p.FileHash != nil && (input.FileHash == nil || *p.FileHash != *input.FileHash) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
