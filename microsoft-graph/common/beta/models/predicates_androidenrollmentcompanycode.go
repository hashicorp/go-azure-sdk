package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnrollmentCompanyCodeOperationPredicate struct {
	EnrollmentToken *string
	ODataType       *string
	QrCodeContent   *string
}

func (p AndroidEnrollmentCompanyCodeOperationPredicate) Matches(input AndroidEnrollmentCompanyCode) bool {

	if p.EnrollmentToken != nil && (input.EnrollmentToken == nil || *p.EnrollmentToken != *input.EnrollmentToken) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.QrCodeContent != nil && (input.QrCodeContent == nil || *p.QrCodeContent != *input.QrCodeContent) {
		return false
	}

	return true
}
