package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosKerberosSingleSignOnExtensionOperationPredicate struct {
	ActiveDirectorySiteCode                  *string
	BlockActiveDirectorySiteAutoDiscovery    *bool
	BlockAutomaticLogin                      *bool
	CacheName                                *string
	IsDefaultRealm                           *bool
	ManagedAppsInBundleIdACLIncluded         *bool
	ODataType                                *string
	PasswordBlockModification                *bool
	PasswordChangeUrl                        *string
	PasswordEnableLocalSync                  *bool
	PasswordExpirationDays                   *int64
	PasswordExpirationNotificationDays       *int64
	PasswordMinimumAgeDays                   *int64
	PasswordMinimumLength                    *int64
	PasswordPreviousPasswordBlockCount       *int64
	PasswordRequireActiveDirectoryComplexity *bool
	PasswordRequirementsDescription          *string
	Realm                                    *string
	RequireUserPresence                      *bool
	SignInHelpText                           *string
	UserPrincipalName                        *string
}

func (p IosKerberosSingleSignOnExtensionOperationPredicate) Matches(input IosKerberosSingleSignOnExtension) bool {

	if p.ActiveDirectorySiteCode != nil && (input.ActiveDirectorySiteCode == nil || *p.ActiveDirectorySiteCode != *input.ActiveDirectorySiteCode) {
		return false
	}

	if p.BlockActiveDirectorySiteAutoDiscovery != nil && (input.BlockActiveDirectorySiteAutoDiscovery == nil || *p.BlockActiveDirectorySiteAutoDiscovery != *input.BlockActiveDirectorySiteAutoDiscovery) {
		return false
	}

	if p.BlockAutomaticLogin != nil && (input.BlockAutomaticLogin == nil || *p.BlockAutomaticLogin != *input.BlockAutomaticLogin) {
		return false
	}

	if p.CacheName != nil && (input.CacheName == nil || *p.CacheName != *input.CacheName) {
		return false
	}

	if p.IsDefaultRealm != nil && (input.IsDefaultRealm == nil || *p.IsDefaultRealm != *input.IsDefaultRealm) {
		return false
	}

	if p.ManagedAppsInBundleIdACLIncluded != nil && (input.ManagedAppsInBundleIdACLIncluded == nil || *p.ManagedAppsInBundleIdACLIncluded != *input.ManagedAppsInBundleIdACLIncluded) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PasswordBlockModification != nil && (input.PasswordBlockModification == nil || *p.PasswordBlockModification != *input.PasswordBlockModification) {
		return false
	}

	if p.PasswordChangeUrl != nil && (input.PasswordChangeUrl == nil || *p.PasswordChangeUrl != *input.PasswordChangeUrl) {
		return false
	}

	if p.PasswordEnableLocalSync != nil && (input.PasswordEnableLocalSync == nil || *p.PasswordEnableLocalSync != *input.PasswordEnableLocalSync) {
		return false
	}

	if p.PasswordExpirationDays != nil && (input.PasswordExpirationDays == nil || *p.PasswordExpirationDays != *input.PasswordExpirationDays) {
		return false
	}

	if p.PasswordExpirationNotificationDays != nil && (input.PasswordExpirationNotificationDays == nil || *p.PasswordExpirationNotificationDays != *input.PasswordExpirationNotificationDays) {
		return false
	}

	if p.PasswordMinimumAgeDays != nil && (input.PasswordMinimumAgeDays == nil || *p.PasswordMinimumAgeDays != *input.PasswordMinimumAgeDays) {
		return false
	}

	if p.PasswordMinimumLength != nil && (input.PasswordMinimumLength == nil || *p.PasswordMinimumLength != *input.PasswordMinimumLength) {
		return false
	}

	if p.PasswordPreviousPasswordBlockCount != nil && (input.PasswordPreviousPasswordBlockCount == nil || *p.PasswordPreviousPasswordBlockCount != *input.PasswordPreviousPasswordBlockCount) {
		return false
	}

	if p.PasswordRequireActiveDirectoryComplexity != nil && (input.PasswordRequireActiveDirectoryComplexity == nil || *p.PasswordRequireActiveDirectoryComplexity != *input.PasswordRequireActiveDirectoryComplexity) {
		return false
	}

	if p.PasswordRequirementsDescription != nil && (input.PasswordRequirementsDescription == nil || *p.PasswordRequirementsDescription != *input.PasswordRequirementsDescription) {
		return false
	}

	if p.Realm != nil && (input.Realm == nil || *p.Realm != *input.Realm) {
		return false
	}

	if p.RequireUserPresence != nil && (input.RequireUserPresence == nil || *p.RequireUserPresence != *input.RequireUserPresence) {
		return false
	}

	if p.SignInHelpText != nil && (input.SignInHelpText == nil || *p.SignInHelpText != *input.SignInHelpText) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
