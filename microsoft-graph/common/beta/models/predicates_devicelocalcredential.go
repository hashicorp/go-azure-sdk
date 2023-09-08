package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLocalCredentialOperationPredicate struct {
	AccountName    *string
	AccountSid     *string
	BackupDateTime *string
	ODataType      *string
	PasswordBase64 *string
}

func (p DeviceLocalCredentialOperationPredicate) Matches(input DeviceLocalCredential) bool {

	if p.AccountName != nil && (input.AccountName == nil || *p.AccountName != *input.AccountName) {
		return false
	}

	if p.AccountSid != nil && (input.AccountSid == nil || *p.AccountSid != *input.AccountSid) {
		return false
	}

	if p.BackupDateTime != nil && (input.BackupDateTime == nil || *p.BackupDateTime != *input.BackupDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PasswordBase64 != nil && (input.PasswordBase64 == nil || *p.PasswordBase64 != *input.PasswordBase64) {
		return false
	}

	return true
}
