package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosTrustedRootCertificateOperationPredicate struct {
	CertFileName           *string
	CreatedDateTime        *string
	Description            *string
	DisplayName            *string
	Id                     *string
	LastModifiedDateTime   *string
	ODataType              *string
	SupportsScopeTags      *bool
	TrustedRootCertificate *string
	Version                *int64
}

func (p IosTrustedRootCertificateOperationPredicate) Matches(input IosTrustedRootCertificate) bool {

	if p.CertFileName != nil && (input.CertFileName == nil || *p.CertFileName != *input.CertFileName) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.TrustedRootCertificate != nil && (input.TrustedRootCertificate == nil || *p.TrustedRootCertificate != *input.TrustedRootCertificate) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
