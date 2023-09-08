package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateUserBindingOperationPredicate struct {
	ODataType            *string
	Priority             *int64
	UserProperty         *string
	X509CertificateField *string
}

func (p X509CertificateUserBindingOperationPredicate) Matches(input X509CertificateUserBinding) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	if p.UserProperty != nil && (input.UserProperty == nil || *p.UserProperty != *input.UserProperty) {
		return false
	}

	if p.X509CertificateField != nil && (input.X509CertificateField == nil || *p.X509CertificateField != *input.X509CertificateField) {
		return false
	}

	return true
}
