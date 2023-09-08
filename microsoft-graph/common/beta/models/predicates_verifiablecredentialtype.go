package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialTypeOperationPredicate struct {
	CredentialType *string
	ODataType      *string
}

func (p VerifiableCredentialTypeOperationPredicate) Matches(input VerifiableCredentialType) bool {

	if p.CredentialType != nil && (input.CredentialType == nil || *p.CredentialType != *input.CredentialType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
