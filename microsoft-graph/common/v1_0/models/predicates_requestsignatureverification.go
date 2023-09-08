package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestSignatureVerificationOperationPredicate struct {
	IsSignedRequestRequired *bool
	ODataType               *string
}

func (p RequestSignatureVerificationOperationPredicate) Matches(input RequestSignatureVerification) bool {

	if p.IsSignedRequestRequired != nil && (input.IsSignedRequestRequired == nil || *p.IsSignedRequestRequired != *input.IsSignedRequestRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
