package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateAuthorityOperationPredicate struct {
	Certificate                       *string
	CertificateRevocationListUrl      *string
	DeltaCertificateRevocationListUrl *string
	IsRootAuthority                   *bool
	Issuer                            *string
	IssuerSki                         *string
	ODataType                         *string
}

func (p CertificateAuthorityOperationPredicate) Matches(input CertificateAuthority) bool {

	if p.Certificate != nil && (input.Certificate == nil || *p.Certificate != *input.Certificate) {
		return false
	}

	if p.CertificateRevocationListUrl != nil && (input.CertificateRevocationListUrl == nil || *p.CertificateRevocationListUrl != *input.CertificateRevocationListUrl) {
		return false
	}

	if p.DeltaCertificateRevocationListUrl != nil && (input.DeltaCertificateRevocationListUrl == nil || *p.DeltaCertificateRevocationListUrl != *input.DeltaCertificateRevocationListUrl) {
		return false
	}

	if p.IsRootAuthority != nil && (input.IsRootAuthority == nil || *p.IsRootAuthority != *input.IsRootAuthority) {
		return false
	}

	if p.Issuer != nil && (input.Issuer == nil || *p.Issuer != *input.Issuer) {
		return false
	}

	if p.IssuerSki != nil && (input.IssuerSki == nil || *p.IssuerSki != *input.IssuerSki) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
