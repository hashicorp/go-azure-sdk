package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationOperationPredicate struct {
	BlockUnifiedPasswordForWorkProfile                        *bool
	CreatedDateTime                                           *string
	Description                                               *string
	DisplayName                                               *string
	Id                                                        *string
	LastModifiedDateTime                                      *string
	ODataType                                                 *string
	PasswordBlockFaceUnlock                                   *bool
	PasswordBlockFingerprintUnlock                            *bool
	PasswordBlockIrisUnlock                                   *bool
	PasswordBlockTrustAgents                                  *bool
	PasswordExpirationDays                                    *int64
	PasswordMinimumLength                                     *int64
	PasswordMinutesOfInactivityBeforeScreenTimeout            *int64
	PasswordPreviousPasswordBlockCount                        *int64
	PasswordSignInFailureCountBeforeFactoryReset              *int64
	SecurityRequireVerifyApps                                 *bool
	SupportsScopeTags                                         *bool
	Version                                                   *int64
	VpnAlwaysOnPackageIdentifier                              *string
	VpnEnableAlwaysOnLockdownMode                             *bool
	WorkProfileAllowWidgets                                   *bool
	WorkProfileBlockAddingAccounts                            *bool
	WorkProfileBlockCamera                                    *bool
	WorkProfileBlockCrossProfileCallerId                      *bool
	WorkProfileBlockCrossProfileContactsSearch                *bool
	WorkProfileBlockCrossProfileCopyPaste                     *bool
	WorkProfileBlockNotificationsWhileDeviceLocked            *bool
	WorkProfileBlockPersonalAppInstallsFromUnknownSources     *bool
	WorkProfileBlockScreenCapture                             *bool
	WorkProfileBluetoothEnableContactSharing                  *bool
	WorkProfilePasswordBlockFaceUnlock                        *bool
	WorkProfilePasswordBlockFingerprintUnlock                 *bool
	WorkProfilePasswordBlockIrisUnlock                        *bool
	WorkProfilePasswordBlockTrustAgents                       *bool
	WorkProfilePasswordExpirationDays                         *int64
	WorkProfilePasswordMinLetterCharacters                    *int64
	WorkProfilePasswordMinLowerCaseCharacters                 *int64
	WorkProfilePasswordMinNonLetterCharacters                 *int64
	WorkProfilePasswordMinNumericCharacters                   *int64
	WorkProfilePasswordMinSymbolCharacters                    *int64
	WorkProfilePasswordMinUpperCaseCharacters                 *int64
	WorkProfilePasswordMinimumLength                          *int64
	WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout *int64
	WorkProfilePasswordPreviousPasswordBlockCount             *int64
	WorkProfilePasswordSignInFailureCountBeforeFactoryReset   *int64
	WorkProfileRequirePassword                                *bool
}

func (p AndroidForWorkGeneralDeviceConfigurationOperationPredicate) Matches(input AndroidForWorkGeneralDeviceConfiguration) bool {

	if p.BlockUnifiedPasswordForWorkProfile != nil && (input.BlockUnifiedPasswordForWorkProfile == nil || *p.BlockUnifiedPasswordForWorkProfile != *input.BlockUnifiedPasswordForWorkProfile) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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

	if p.PasswordBlockFaceUnlock != nil && (input.PasswordBlockFaceUnlock == nil || *p.PasswordBlockFaceUnlock != *input.PasswordBlockFaceUnlock) {
		return false
	}

	if p.PasswordBlockFingerprintUnlock != nil && (input.PasswordBlockFingerprintUnlock == nil || *p.PasswordBlockFingerprintUnlock != *input.PasswordBlockFingerprintUnlock) {
		return false
	}

	if p.PasswordBlockIrisUnlock != nil && (input.PasswordBlockIrisUnlock == nil || *p.PasswordBlockIrisUnlock != *input.PasswordBlockIrisUnlock) {
		return false
	}

	if p.PasswordBlockTrustAgents != nil && (input.PasswordBlockTrustAgents == nil || *p.PasswordBlockTrustAgents != *input.PasswordBlockTrustAgents) {
		return false
	}

	if p.PasswordExpirationDays != nil && (input.PasswordExpirationDays == nil || *p.PasswordExpirationDays != *input.PasswordExpirationDays) {
		return false
	}

	if p.PasswordMinimumLength != nil && (input.PasswordMinimumLength == nil || *p.PasswordMinimumLength != *input.PasswordMinimumLength) {
		return false
	}

	if p.PasswordMinutesOfInactivityBeforeScreenTimeout != nil && (input.PasswordMinutesOfInactivityBeforeScreenTimeout == nil || *p.PasswordMinutesOfInactivityBeforeScreenTimeout != *input.PasswordMinutesOfInactivityBeforeScreenTimeout) {
		return false
	}

	if p.PasswordPreviousPasswordBlockCount != nil && (input.PasswordPreviousPasswordBlockCount == nil || *p.PasswordPreviousPasswordBlockCount != *input.PasswordPreviousPasswordBlockCount) {
		return false
	}

	if p.PasswordSignInFailureCountBeforeFactoryReset != nil && (input.PasswordSignInFailureCountBeforeFactoryReset == nil || *p.PasswordSignInFailureCountBeforeFactoryReset != *input.PasswordSignInFailureCountBeforeFactoryReset) {
		return false
	}

	if p.SecurityRequireVerifyApps != nil && (input.SecurityRequireVerifyApps == nil || *p.SecurityRequireVerifyApps != *input.SecurityRequireVerifyApps) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	if p.VpnAlwaysOnPackageIdentifier != nil && (input.VpnAlwaysOnPackageIdentifier == nil || *p.VpnAlwaysOnPackageIdentifier != *input.VpnAlwaysOnPackageIdentifier) {
		return false
	}

	if p.VpnEnableAlwaysOnLockdownMode != nil && (input.VpnEnableAlwaysOnLockdownMode == nil || *p.VpnEnableAlwaysOnLockdownMode != *input.VpnEnableAlwaysOnLockdownMode) {
		return false
	}

	if p.WorkProfileAllowWidgets != nil && (input.WorkProfileAllowWidgets == nil || *p.WorkProfileAllowWidgets != *input.WorkProfileAllowWidgets) {
		return false
	}

	if p.WorkProfileBlockAddingAccounts != nil && (input.WorkProfileBlockAddingAccounts == nil || *p.WorkProfileBlockAddingAccounts != *input.WorkProfileBlockAddingAccounts) {
		return false
	}

	if p.WorkProfileBlockCamera != nil && (input.WorkProfileBlockCamera == nil || *p.WorkProfileBlockCamera != *input.WorkProfileBlockCamera) {
		return false
	}

	if p.WorkProfileBlockCrossProfileCallerId != nil && (input.WorkProfileBlockCrossProfileCallerId == nil || *p.WorkProfileBlockCrossProfileCallerId != *input.WorkProfileBlockCrossProfileCallerId) {
		return false
	}

	if p.WorkProfileBlockCrossProfileContactsSearch != nil && (input.WorkProfileBlockCrossProfileContactsSearch == nil || *p.WorkProfileBlockCrossProfileContactsSearch != *input.WorkProfileBlockCrossProfileContactsSearch) {
		return false
	}

	if p.WorkProfileBlockCrossProfileCopyPaste != nil && (input.WorkProfileBlockCrossProfileCopyPaste == nil || *p.WorkProfileBlockCrossProfileCopyPaste != *input.WorkProfileBlockCrossProfileCopyPaste) {
		return false
	}

	if p.WorkProfileBlockNotificationsWhileDeviceLocked != nil && (input.WorkProfileBlockNotificationsWhileDeviceLocked == nil || *p.WorkProfileBlockNotificationsWhileDeviceLocked != *input.WorkProfileBlockNotificationsWhileDeviceLocked) {
		return false
	}

	if p.WorkProfileBlockPersonalAppInstallsFromUnknownSources != nil && (input.WorkProfileBlockPersonalAppInstallsFromUnknownSources == nil || *p.WorkProfileBlockPersonalAppInstallsFromUnknownSources != *input.WorkProfileBlockPersonalAppInstallsFromUnknownSources) {
		return false
	}

	if p.WorkProfileBlockScreenCapture != nil && (input.WorkProfileBlockScreenCapture == nil || *p.WorkProfileBlockScreenCapture != *input.WorkProfileBlockScreenCapture) {
		return false
	}

	if p.WorkProfileBluetoothEnableContactSharing != nil && (input.WorkProfileBluetoothEnableContactSharing == nil || *p.WorkProfileBluetoothEnableContactSharing != *input.WorkProfileBluetoothEnableContactSharing) {
		return false
	}

	if p.WorkProfilePasswordBlockFaceUnlock != nil && (input.WorkProfilePasswordBlockFaceUnlock == nil || *p.WorkProfilePasswordBlockFaceUnlock != *input.WorkProfilePasswordBlockFaceUnlock) {
		return false
	}

	if p.WorkProfilePasswordBlockFingerprintUnlock != nil && (input.WorkProfilePasswordBlockFingerprintUnlock == nil || *p.WorkProfilePasswordBlockFingerprintUnlock != *input.WorkProfilePasswordBlockFingerprintUnlock) {
		return false
	}

	if p.WorkProfilePasswordBlockIrisUnlock != nil && (input.WorkProfilePasswordBlockIrisUnlock == nil || *p.WorkProfilePasswordBlockIrisUnlock != *input.WorkProfilePasswordBlockIrisUnlock) {
		return false
	}

	if p.WorkProfilePasswordBlockTrustAgents != nil && (input.WorkProfilePasswordBlockTrustAgents == nil || *p.WorkProfilePasswordBlockTrustAgents != *input.WorkProfilePasswordBlockTrustAgents) {
		return false
	}

	if p.WorkProfilePasswordExpirationDays != nil && (input.WorkProfilePasswordExpirationDays == nil || *p.WorkProfilePasswordExpirationDays != *input.WorkProfilePasswordExpirationDays) {
		return false
	}

	if p.WorkProfilePasswordMinLetterCharacters != nil && (input.WorkProfilePasswordMinLetterCharacters == nil || *p.WorkProfilePasswordMinLetterCharacters != *input.WorkProfilePasswordMinLetterCharacters) {
		return false
	}

	if p.WorkProfilePasswordMinLowerCaseCharacters != nil && (input.WorkProfilePasswordMinLowerCaseCharacters == nil || *p.WorkProfilePasswordMinLowerCaseCharacters != *input.WorkProfilePasswordMinLowerCaseCharacters) {
		return false
	}

	if p.WorkProfilePasswordMinNonLetterCharacters != nil && (input.WorkProfilePasswordMinNonLetterCharacters == nil || *p.WorkProfilePasswordMinNonLetterCharacters != *input.WorkProfilePasswordMinNonLetterCharacters) {
		return false
	}

	if p.WorkProfilePasswordMinNumericCharacters != nil && (input.WorkProfilePasswordMinNumericCharacters == nil || *p.WorkProfilePasswordMinNumericCharacters != *input.WorkProfilePasswordMinNumericCharacters) {
		return false
	}

	if p.WorkProfilePasswordMinSymbolCharacters != nil && (input.WorkProfilePasswordMinSymbolCharacters == nil || *p.WorkProfilePasswordMinSymbolCharacters != *input.WorkProfilePasswordMinSymbolCharacters) {
		return false
	}

	if p.WorkProfilePasswordMinUpperCaseCharacters != nil && (input.WorkProfilePasswordMinUpperCaseCharacters == nil || *p.WorkProfilePasswordMinUpperCaseCharacters != *input.WorkProfilePasswordMinUpperCaseCharacters) {
		return false
	}

	if p.WorkProfilePasswordMinimumLength != nil && (input.WorkProfilePasswordMinimumLength == nil || *p.WorkProfilePasswordMinimumLength != *input.WorkProfilePasswordMinimumLength) {
		return false
	}

	if p.WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout != nil && (input.WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout == nil || *p.WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout != *input.WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout) {
		return false
	}

	if p.WorkProfilePasswordPreviousPasswordBlockCount != nil && (input.WorkProfilePasswordPreviousPasswordBlockCount == nil || *p.WorkProfilePasswordPreviousPasswordBlockCount != *input.WorkProfilePasswordPreviousPasswordBlockCount) {
		return false
	}

	if p.WorkProfilePasswordSignInFailureCountBeforeFactoryReset != nil && (input.WorkProfilePasswordSignInFailureCountBeforeFactoryReset == nil || *p.WorkProfilePasswordSignInFailureCountBeforeFactoryReset != *input.WorkProfilePasswordSignInFailureCountBeforeFactoryReset) {
		return false
	}

	if p.WorkProfileRequirePassword != nil && (input.WorkProfileRequirePassword == nil || *p.WorkProfileRequirePassword != *input.WorkProfileRequirePassword) {
		return false
	}

	return true
}
