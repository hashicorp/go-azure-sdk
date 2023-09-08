package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalLockConfigurationOperationPredicate struct {
	AllProperties              *bool
	CredentialsWithUsageSign   *bool
	CredentialsWithUsageVerify *bool
	IsEnabled                  *bool
	ODataType                  *string
	TokenEncryptionKeyId       *bool
}

func (p ServicePrincipalLockConfigurationOperationPredicate) Matches(input ServicePrincipalLockConfiguration) bool {

	if p.AllProperties != nil && (input.AllProperties == nil || *p.AllProperties != *input.AllProperties) {
		return false
	}

	if p.CredentialsWithUsageSign != nil && (input.CredentialsWithUsageSign == nil || *p.CredentialsWithUsageSign != *input.CredentialsWithUsageSign) {
		return false
	}

	if p.CredentialsWithUsageVerify != nil && (input.CredentialsWithUsageVerify == nil || *p.CredentialsWithUsageVerify != *input.CredentialsWithUsageVerify) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TokenEncryptionKeyId != nil && (input.TokenEncryptionKeyId == nil || *p.TokenEncryptionKeyId != *input.TokenEncryptionKeyId) {
		return false
	}

	return true
}
