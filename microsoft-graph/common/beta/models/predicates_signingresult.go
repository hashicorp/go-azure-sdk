package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SigningResultOperationPredicate struct {
	ODataType    *string
	Signature    *string
	SigningKeyId *string
}

func (p SigningResultOperationPredicate) Matches(input SigningResult) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Signature != nil && (input.Signature == nil || *p.Signature != *input.Signature) {
		return false
	}

	if p.SigningKeyId != nil && (input.SigningKeyId == nil || *p.SigningKeyId != *input.SigningKeyId) {
		return false
	}

	return true
}
