package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVerificationResultOperationPredicate struct {
	ODataType      *string
	SignatureValid *bool
}

func (p SecurityVerificationResultOperationPredicate) Matches(input SecurityVerificationResult) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SignatureValid != nil && (input.SignatureValid == nil || *p.SignatureValid != *input.SignatureValid) {
		return false
	}

	return true
}
