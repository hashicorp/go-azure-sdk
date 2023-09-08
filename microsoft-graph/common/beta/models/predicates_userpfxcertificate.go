package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPFXCertificateOperationPredicate struct {
	CreatedDateTime      *string
	EncryptedPfxBlob     *string
	EncryptedPfxPassword *string
	ExpirationDateTime   *string
	Id                   *string
	KeyName              *string
	LastModifiedDateTime *string
	ODataType            *string
	ProviderName         *string
	StartDateTime        *string
	Thumbprint           *string
	UserPrincipalName    *string
}

func (p UserPFXCertificateOperationPredicate) Matches(input UserPFXCertificate) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.EncryptedPfxBlob != nil && (input.EncryptedPfxBlob == nil || *p.EncryptedPfxBlob != *input.EncryptedPfxBlob) {
		return false
	}

	if p.EncryptedPfxPassword != nil && (input.EncryptedPfxPassword == nil || *p.EncryptedPfxPassword != *input.EncryptedPfxPassword) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.KeyName != nil && (input.KeyName == nil || *p.KeyName != *input.KeyName) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProviderName != nil && (input.ProviderName == nil || *p.ProviderName != *input.ProviderName) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Thumbprint != nil && (input.Thumbprint == nil || *p.Thumbprint != *input.Thumbprint) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
