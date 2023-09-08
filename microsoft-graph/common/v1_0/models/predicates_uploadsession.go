package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UploadSessionOperationPredicate struct {
	ExpirationDateTime *string
	ODataType          *string
	UploadUrl          *string
}

func (p UploadSessionOperationPredicate) Matches(input UploadSession) bool {

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UploadUrl != nil && (input.UploadUrl == nil || *p.UploadUrl != *input.UploadUrl) {
		return false
	}

	return true
}
