package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerFixedDrivePolicyOperationPredicate struct {
	ODataType                       *string
	RequireEncryptionForWriteAccess *bool
}

func (p BitLockerFixedDrivePolicyOperationPredicate) Matches(input BitLockerFixedDrivePolicy) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequireEncryptionForWriteAccess != nil && (input.RequireEncryptionForWriteAccess == nil || *p.RequireEncryptionForWriteAccess != *input.RequireEncryptionForWriteAccess) {
		return false
	}

	return true
}
