package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeManagedFolderOperationPredicate struct {
	FolderIdentifier *string
	FolderName       *string
	ODataType        *string
}

func (p AndroidDeviceOwnerKioskModeManagedFolderOperationPredicate) Matches(input AndroidDeviceOwnerKioskModeManagedFolder) bool {

	if p.FolderIdentifier != nil && (input.FolderIdentifier == nil || *p.FolderIdentifier != *input.FolderIdentifier) {
		return false
	}

	if p.FolderName != nil && (input.FolderName == nil || *p.FolderName != *input.FolderName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
