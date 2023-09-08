package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileEncryptionInfoOperationPredicate struct {
	EncryptionKey        *string
	FileDigest           *string
	FileDigestAlgorithm  *string
	InitializationVector *string
	Mac                  *string
	MacKey               *string
	ODataType            *string
	ProfileIdentifier    *string
}

func (p FileEncryptionInfoOperationPredicate) Matches(input FileEncryptionInfo) bool {

	if p.EncryptionKey != nil && (input.EncryptionKey == nil || *p.EncryptionKey != *input.EncryptionKey) {
		return false
	}

	if p.FileDigest != nil && (input.FileDigest == nil || *p.FileDigest != *input.FileDigest) {
		return false
	}

	if p.FileDigestAlgorithm != nil && (input.FileDigestAlgorithm == nil || *p.FileDigestAlgorithm != *input.FileDigestAlgorithm) {
		return false
	}

	if p.InitializationVector != nil && (input.InitializationVector == nil || *p.InitializationVector != *input.InitializationVector) {
		return false
	}

	if p.Mac != nil && (input.Mac == nil || *p.Mac != *input.Mac) {
		return false
	}

	if p.MacKey != nil && (input.MacKey == nil || *p.MacKey != *input.MacKey) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProfileIdentifier != nil && (input.ProfileIdentifier == nil || *p.ProfileIdentifier != *input.ProfileIdentifier) {
		return false
	}

	return true
}
