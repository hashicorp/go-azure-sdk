package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCForensicStorageAccountOperationPredicate struct {
	ODataType          *string
	StorageAccountId   *string
	StorageAccountName *string
}

func (p CloudPCForensicStorageAccountOperationPredicate) Matches(input CloudPCForensicStorageAccount) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StorageAccountId != nil && (input.StorageAccountId == nil || *p.StorageAccountId != *input.StorageAccountId) {
		return false
	}

	if p.StorageAccountName != nil && (input.StorageAccountName == nil || *p.StorageAccountName != *input.StorageAccountName) {
		return false
	}

	return true
}
