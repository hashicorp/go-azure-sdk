package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BufferDecryptionResultOperationPredicate struct {
	DecryptedBuffer *string
	ODataType       *string
}

func (p BufferDecryptionResultOperationPredicate) Matches(input BufferDecryptionResult) bool {

	if p.DecryptedBuffer != nil && (input.DecryptedBuffer == nil || *p.DecryptedBuffer != *input.DecryptedBuffer) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
