package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSetVersionOperationPredicate struct {
	Comment                   *string
	CreatedDateTime           *string
	Id                        *string
	LastModifiedDateTime      *string
	ODataType                 *string
	ShouldCaptureMinorVersion *bool
}

func (p DocumentSetVersionOperationPredicate) Matches(input DocumentSetVersion) bool {

	if p.Comment != nil && (input.Comment == nil || *p.Comment != *input.Comment) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ShouldCaptureMinorVersion != nil && (input.ShouldCaptureMinorVersion == nil || *p.ShouldCaptureMinorVersion != *input.ShouldCaptureMinorVersion) {
		return false
	}

	return true
}
