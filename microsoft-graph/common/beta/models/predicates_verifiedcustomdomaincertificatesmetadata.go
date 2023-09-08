package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedCustomDomainCertificatesMetadataOperationPredicate struct {
	ExpiryDate  *string
	IssueDate   *string
	IssuerName  *string
	ODataType   *string
	SubjectName *string
	Thumbprint  *string
}

func (p VerifiedCustomDomainCertificatesMetadataOperationPredicate) Matches(input VerifiedCustomDomainCertificatesMetadata) bool {

	if p.ExpiryDate != nil && (input.ExpiryDate == nil || *p.ExpiryDate != *input.ExpiryDate) {
		return false
	}

	if p.IssueDate != nil && (input.IssueDate == nil || *p.IssueDate != *input.IssueDate) {
		return false
	}

	if p.IssuerName != nil && (input.IssuerName == nil || *p.IssuerName != *input.IssuerName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SubjectName != nil && (input.SubjectName == nil || *p.SubjectName != *input.SubjectName) {
		return false
	}

	if p.Thumbprint != nil && (input.Thumbprint == nil || *p.Thumbprint != *input.Thumbprint) {
		return false
	}

	return true
}
