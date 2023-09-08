package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBufferEncryptionResultOperationPredicate struct {
	EncryptedBuffer   *string
	ODataType         *string
	PublishingLicense *string
}

func (p SecurityBufferEncryptionResultOperationPredicate) Matches(input SecurityBufferEncryptionResult) bool {

	if p.EncryptedBuffer != nil && (input.EncryptedBuffer == nil || *p.EncryptedBuffer != *input.EncryptedBuffer) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PublishingLicense != nil && (input.PublishingLicense == nil || *p.PublishingLicense != *input.PublishingLicense) {
		return false
	}

	return true
}
