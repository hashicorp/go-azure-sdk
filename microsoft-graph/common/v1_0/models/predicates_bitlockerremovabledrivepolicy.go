package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRemovableDrivePolicyOperationPredicate struct {
	BlockCrossOrganizationWriteAccess *bool
	ODataType                         *string
	RequireEncryptionForWriteAccess   *bool
}

func (p BitLockerRemovableDrivePolicyOperationPredicate) Matches(input BitLockerRemovableDrivePolicy) bool {

	if p.BlockCrossOrganizationWriteAccess != nil && (input.BlockCrossOrganizationWriteAccess == nil || *p.BlockCrossOrganizationWriteAccess != *input.BlockCrossOrganizationWriteAccess) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequireEncryptionForWriteAccess != nil && (input.RequireEncryptionForWriteAccess == nil || *p.RequireEncryptionForWriteAccess != *input.RequireEncryptionForWriteAccess) {
		return false
	}

	return true
}
