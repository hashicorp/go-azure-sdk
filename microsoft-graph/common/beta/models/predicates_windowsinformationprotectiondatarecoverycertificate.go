package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionDataRecoveryCertificateOperationPredicate struct {
	Certificate        *string
	Description        *string
	ExpirationDateTime *string
	ODataType          *string
	SubjectName        *string
}

func (p WindowsInformationProtectionDataRecoveryCertificateOperationPredicate) Matches(input WindowsInformationProtectionDataRecoveryCertificate) bool {

	if p.Certificate != nil && (input.Certificate == nil || *p.Certificate != *input.Certificate) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SubjectName != nil && (input.SubjectName == nil || *p.SubjectName != *input.SubjectName) {
		return false
	}

	return true
}
