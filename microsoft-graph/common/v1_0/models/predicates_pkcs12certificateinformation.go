package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Pkcs12CertificateInformationOperationPredicate struct {
	IsActive   *bool
	NotAfter   *int64
	NotBefore  *int64
	ODataType  *string
	Thumbprint *string
}

func (p Pkcs12CertificateInformationOperationPredicate) Matches(input Pkcs12CertificateInformation) bool {

	if p.IsActive != nil && (input.IsActive == nil || *p.IsActive != *input.IsActive) {
		return false
	}

	if p.NotAfter != nil && (input.NotAfter == nil || *p.NotAfter != *input.NotAfter) {
		return false
	}

	if p.NotBefore != nil && (input.NotBefore == nil || *p.NotBefore != *input.NotBefore) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Thumbprint != nil && (input.Thumbprint == nil || *p.Thumbprint != *input.Thumbprint) {
		return false
	}

	return true
}
