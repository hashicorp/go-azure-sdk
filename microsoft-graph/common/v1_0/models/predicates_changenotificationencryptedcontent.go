package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotificationEncryptedContentOperationPredicate struct {
	Data                            *string
	DataKey                         *string
	DataSignature                   *string
	EncryptionCertificateId         *string
	EncryptionCertificateThumbprint *string
	ODataType                       *string
}

func (p ChangeNotificationEncryptedContentOperationPredicate) Matches(input ChangeNotificationEncryptedContent) bool {

	if p.Data != nil && (input.Data == nil || *p.Data != *input.Data) {
		return false
	}

	if p.DataKey != nil && (input.DataKey == nil || *p.DataKey != *input.DataKey) {
		return false
	}

	if p.DataSignature != nil && (input.DataSignature == nil || *p.DataSignature != *input.DataSignature) {
		return false
	}

	if p.EncryptionCertificateId != nil && (input.EncryptionCertificateId == nil || *p.EncryptionCertificateId != *input.EncryptionCertificateId) {
		return false
	}

	if p.EncryptionCertificateThumbprint != nil && (input.EncryptionCertificateThumbprint == nil || *p.EncryptionCertificateThumbprint != *input.EncryptionCertificateThumbprint) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
