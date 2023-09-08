package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleManagedIdentityProviderOperationPredicate struct {
	CertificateData *string
	DeveloperId     *string
	DisplayName     *string
	Id              *string
	KeyId           *string
	ODataType       *string
	ServiceId       *string
}

func (p AppleManagedIdentityProviderOperationPredicate) Matches(input AppleManagedIdentityProvider) bool {

	if p.CertificateData != nil && (input.CertificateData == nil || *p.CertificateData != *input.CertificateData) {
		return false
	}

	if p.DeveloperId != nil && (input.DeveloperId == nil || *p.DeveloperId != *input.DeveloperId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.KeyId != nil && (input.KeyId == nil || *p.KeyId != *input.KeyId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ServiceId != nil && (input.ServiceId == nil || *p.ServiceId != *input.ServiceId) {
		return false
	}

	return true
}
