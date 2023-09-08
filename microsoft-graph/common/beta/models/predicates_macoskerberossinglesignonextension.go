package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSKerberosSingleSignOnExtensionOperationPredicate struct {
	ActiveDirectorySiteCode                  *string
	BlockActiveDirectorySiteAutoDiscovery    *bool
	BlockAutomaticLogin                      *bool
	CacheName                                *string
	CredentialsCacheMonitored                *bool
	IsDefaultRealm                           *bool
	KerberosAppsInBundleIdACLIncluded        *bool
	ManagedAppsInBundleIdACLIncluded         *bool
	ModeCredentialUsed                       *string
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
	TlsForLDAPRequired                       *bool
	UserPrincipalName                        *string
	UserSetupDelayed                         *bool
	UsernameLabelCustom                      *string
}

func (p MacOSKerberosSingleSignOnExtensionOperationPredicate) Matches(input MacOSKerberosSingleSignOnExtension) bool {

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

	if p.CredentialsCacheMonitored != nil && (input.CredentialsCacheMonitored == nil || *p.CredentialsCacheMonitored != *input.CredentialsCacheMonitored) {
		return false
	}

	if p.IsDefaultRealm != nil && (input.IsDefaultRealm == nil || *p.IsDefaultRealm != *input.IsDefaultRealm) {
		return false
	}

	if p.KerberosAppsInBundleIdACLIncluded != nil && (input.KerberosAppsInBundleIdACLIncluded == nil || *p.KerberosAppsInBundleIdACLIncluded != *input.KerberosAppsInBundleIdACLIncluded) {
		return false
	}

	if p.ManagedAppsInBundleIdACLIncluded != nil && (input.ManagedAppsInBundleIdACLIncluded == nil || *p.ManagedAppsInBundleIdACLIncluded != *input.ManagedAppsInBundleIdACLIncluded) {
		return false
	}

	if p.ModeCredentialUsed != nil && (input.ModeCredentialUsed == nil || *p.ModeCredentialUsed != *input.ModeCredentialUsed) {
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

	if p.TlsForLDAPRequired != nil && (input.TlsForLDAPRequired == nil || *p.TlsForLDAPRequired != *input.TlsForLDAPRequired) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	if p.UserSetupDelayed != nil && (input.UserSetupDelayed == nil || *p.UserSetupDelayed != *input.UserSetupDelayed) {
		return false
	}

	if p.UsernameLabelCustom != nil && (input.UsernameLabelCustom == nil || *p.UsernameLabelCustom != *input.UsernameLabelCustom) {
		return false
	}

	return true
}
