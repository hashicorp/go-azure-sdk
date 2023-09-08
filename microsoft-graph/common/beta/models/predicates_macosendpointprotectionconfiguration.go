package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationOperationPredicate struct {
	CreatedDateTime                                      *string
	Description                                          *string
	DisplayName                                          *string
	FileVaultAllowDeferralUntilSignOut                   *bool
	FileVaultDisablePromptAtSignOut                      *bool
	FileVaultEnabled                                     *bool
	FileVaultHidePersonalRecoveryKey                     *bool
	FileVaultInstitutionalRecoveryKeyCertificate         *string
	FileVaultInstitutionalRecoveryKeyCertificateFileName *string
	FileVaultNumberOfTimesUserCanIgnore                  *int64
	FileVaultPersonalRecoveryKeyHelpMessage              *string
	FileVaultPersonalRecoveryKeyRotationInMonths         *int64
	FirewallBlockAllIncoming                             *bool
	FirewallEnableStealthMode                            *bool
	FirewallEnabled                                      *bool
	GatekeeperBlockOverride                              *bool
	Id                                                   *string
	LastModifiedDateTime                                 *string
	ODataType                                            *string
	SupportsScopeTags                                    *bool
	Version                                              *int64
}

func (p MacOSEndpointProtectionConfigurationOperationPredicate) Matches(input MacOSEndpointProtectionConfiguration) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.FileVaultAllowDeferralUntilSignOut != nil && (input.FileVaultAllowDeferralUntilSignOut == nil || *p.FileVaultAllowDeferralUntilSignOut != *input.FileVaultAllowDeferralUntilSignOut) {
		return false
	}

	if p.FileVaultDisablePromptAtSignOut != nil && (input.FileVaultDisablePromptAtSignOut == nil || *p.FileVaultDisablePromptAtSignOut != *input.FileVaultDisablePromptAtSignOut) {
		return false
	}

	if p.FileVaultEnabled != nil && (input.FileVaultEnabled == nil || *p.FileVaultEnabled != *input.FileVaultEnabled) {
		return false
	}

	if p.FileVaultHidePersonalRecoveryKey != nil && (input.FileVaultHidePersonalRecoveryKey == nil || *p.FileVaultHidePersonalRecoveryKey != *input.FileVaultHidePersonalRecoveryKey) {
		return false
	}

	if p.FileVaultInstitutionalRecoveryKeyCertificate != nil && (input.FileVaultInstitutionalRecoveryKeyCertificate == nil || *p.FileVaultInstitutionalRecoveryKeyCertificate != *input.FileVaultInstitutionalRecoveryKeyCertificate) {
		return false
	}

	if p.FileVaultInstitutionalRecoveryKeyCertificateFileName != nil && (input.FileVaultInstitutionalRecoveryKeyCertificateFileName == nil || *p.FileVaultInstitutionalRecoveryKeyCertificateFileName != *input.FileVaultInstitutionalRecoveryKeyCertificateFileName) {
		return false
	}

	if p.FileVaultNumberOfTimesUserCanIgnore != nil && (input.FileVaultNumberOfTimesUserCanIgnore == nil || *p.FileVaultNumberOfTimesUserCanIgnore != *input.FileVaultNumberOfTimesUserCanIgnore) {
		return false
	}

	if p.FileVaultPersonalRecoveryKeyHelpMessage != nil && (input.FileVaultPersonalRecoveryKeyHelpMessage == nil || *p.FileVaultPersonalRecoveryKeyHelpMessage != *input.FileVaultPersonalRecoveryKeyHelpMessage) {
		return false
	}

	if p.FileVaultPersonalRecoveryKeyRotationInMonths != nil && (input.FileVaultPersonalRecoveryKeyRotationInMonths == nil || *p.FileVaultPersonalRecoveryKeyRotationInMonths != *input.FileVaultPersonalRecoveryKeyRotationInMonths) {
		return false
	}

	if p.FirewallBlockAllIncoming != nil && (input.FirewallBlockAllIncoming == nil || *p.FirewallBlockAllIncoming != *input.FirewallBlockAllIncoming) {
		return false
	}

	if p.FirewallEnableStealthMode != nil && (input.FirewallEnableStealthMode == nil || *p.FirewallEnableStealthMode != *input.FirewallEnableStealthMode) {
		return false
	}

	if p.FirewallEnabled != nil && (input.FirewallEnabled == nil || *p.FirewallEnabled != *input.FirewallEnabled) {
		return false
	}

	if p.GatekeeperBlockOverride != nil && (input.GatekeeperBlockOverride == nil || *p.GatekeeperBlockOverride != *input.GatekeeperBlockOverride) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
