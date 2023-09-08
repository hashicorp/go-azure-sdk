package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationOperationPredicate struct {
	AccountsBlockAddingNonMicrosoftAccountEmail    *bool
	ApplyOnlyToWindows81                           *bool
	BrowserBlockAutofill                           *bool
	BrowserBlockAutomaticDetectionOfIntranetSites  *bool
	BrowserBlockEnterpriseModeAccess               *bool
	BrowserBlockJavaScript                         *bool
	BrowserBlockPlugins                            *bool
	BrowserBlockPopups                             *bool
	BrowserBlockSendingDoNotTrackHeader            *bool
	BrowserBlockSingleWordEntryOnIntranetSites     *bool
	BrowserEnterpriseModeSiteListLocation          *string
	BrowserLoggingReportLocation                   *string
	BrowserRequireFirewall                         *bool
	BrowserRequireFraudWarning                     *bool
	BrowserRequireHighSecurityForRestrictedSites   *bool
	BrowserRequireSmartScreen                      *bool
	CellularBlockDataRoaming                       *bool
	CreatedDateTime                                *string
	Description                                    *string
	DiagnosticsBlockDataSubmission                 *bool
	DisplayName                                    *string
	Id                                             *string
	LastModifiedDateTime                           *string
	ODataType                                      *string
	PasswordBlockPicturePasswordAndPin             *bool
	PasswordExpirationDays                         *int64
	PasswordMinimumCharacterSetCount               *int64
	PasswordMinimumLength                          *int64
	PasswordMinutesOfInactivityBeforeScreenTimeout *int64
	PasswordPreviousPasswordBlockCount             *int64
	PasswordSignInFailureCountBeforeFactoryReset   *int64
	StorageRequireDeviceEncryption                 *bool
	SupportsScopeTags                              *bool
	UpdatesRequireAutomaticUpdates                 *bool
	Version                                        *int64
	WorkFoldersUrl                                 *string
}

func (p Windows81GeneralConfigurationOperationPredicate) Matches(input Windows81GeneralConfiguration) bool {

	if p.AccountsBlockAddingNonMicrosoftAccountEmail != nil && (input.AccountsBlockAddingNonMicrosoftAccountEmail == nil || *p.AccountsBlockAddingNonMicrosoftAccountEmail != *input.AccountsBlockAddingNonMicrosoftAccountEmail) {
		return false
	}

	if p.ApplyOnlyToWindows81 != nil && (input.ApplyOnlyToWindows81 == nil || *p.ApplyOnlyToWindows81 != *input.ApplyOnlyToWindows81) {
		return false
	}

	if p.BrowserBlockAutofill != nil && (input.BrowserBlockAutofill == nil || *p.BrowserBlockAutofill != *input.BrowserBlockAutofill) {
		return false
	}

	if p.BrowserBlockAutomaticDetectionOfIntranetSites != nil && (input.BrowserBlockAutomaticDetectionOfIntranetSites == nil || *p.BrowserBlockAutomaticDetectionOfIntranetSites != *input.BrowserBlockAutomaticDetectionOfIntranetSites) {
		return false
	}

	if p.BrowserBlockEnterpriseModeAccess != nil && (input.BrowserBlockEnterpriseModeAccess == nil || *p.BrowserBlockEnterpriseModeAccess != *input.BrowserBlockEnterpriseModeAccess) {
		return false
	}

	if p.BrowserBlockJavaScript != nil && (input.BrowserBlockJavaScript == nil || *p.BrowserBlockJavaScript != *input.BrowserBlockJavaScript) {
		return false
	}

	if p.BrowserBlockPlugins != nil && (input.BrowserBlockPlugins == nil || *p.BrowserBlockPlugins != *input.BrowserBlockPlugins) {
		return false
	}

	if p.BrowserBlockPopups != nil && (input.BrowserBlockPopups == nil || *p.BrowserBlockPopups != *input.BrowserBlockPopups) {
		return false
	}

	if p.BrowserBlockSendingDoNotTrackHeader != nil && (input.BrowserBlockSendingDoNotTrackHeader == nil || *p.BrowserBlockSendingDoNotTrackHeader != *input.BrowserBlockSendingDoNotTrackHeader) {
		return false
	}

	if p.BrowserBlockSingleWordEntryOnIntranetSites != nil && (input.BrowserBlockSingleWordEntryOnIntranetSites == nil || *p.BrowserBlockSingleWordEntryOnIntranetSites != *input.BrowserBlockSingleWordEntryOnIntranetSites) {
		return false
	}

	if p.BrowserEnterpriseModeSiteListLocation != nil && (input.BrowserEnterpriseModeSiteListLocation == nil || *p.BrowserEnterpriseModeSiteListLocation != *input.BrowserEnterpriseModeSiteListLocation) {
		return false
	}

	if p.BrowserLoggingReportLocation != nil && (input.BrowserLoggingReportLocation == nil || *p.BrowserLoggingReportLocation != *input.BrowserLoggingReportLocation) {
		return false
	}

	if p.BrowserRequireFirewall != nil && (input.BrowserRequireFirewall == nil || *p.BrowserRequireFirewall != *input.BrowserRequireFirewall) {
		return false
	}

	if p.BrowserRequireFraudWarning != nil && (input.BrowserRequireFraudWarning == nil || *p.BrowserRequireFraudWarning != *input.BrowserRequireFraudWarning) {
		return false
	}

	if p.BrowserRequireHighSecurityForRestrictedSites != nil && (input.BrowserRequireHighSecurityForRestrictedSites == nil || *p.BrowserRequireHighSecurityForRestrictedSites != *input.BrowserRequireHighSecurityForRestrictedSites) {
		return false
	}

	if p.BrowserRequireSmartScreen != nil && (input.BrowserRequireSmartScreen == nil || *p.BrowserRequireSmartScreen != *input.BrowserRequireSmartScreen) {
		return false
	}

	if p.CellularBlockDataRoaming != nil && (input.CellularBlockDataRoaming == nil || *p.CellularBlockDataRoaming != *input.CellularBlockDataRoaming) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DiagnosticsBlockDataSubmission != nil && (input.DiagnosticsBlockDataSubmission == nil || *p.DiagnosticsBlockDataSubmission != *input.DiagnosticsBlockDataSubmission) {
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

	if p.PasswordBlockPicturePasswordAndPin != nil && (input.PasswordBlockPicturePasswordAndPin == nil || *p.PasswordBlockPicturePasswordAndPin != *input.PasswordBlockPicturePasswordAndPin) {
		return false
	}

	if p.PasswordExpirationDays != nil && (input.PasswordExpirationDays == nil || *p.PasswordExpirationDays != *input.PasswordExpirationDays) {
		return false
	}

	if p.PasswordMinimumCharacterSetCount != nil && (input.PasswordMinimumCharacterSetCount == nil || *p.PasswordMinimumCharacterSetCount != *input.PasswordMinimumCharacterSetCount) {
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

	if p.StorageRequireDeviceEncryption != nil && (input.StorageRequireDeviceEncryption == nil || *p.StorageRequireDeviceEncryption != *input.StorageRequireDeviceEncryption) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.UpdatesRequireAutomaticUpdates != nil && (input.UpdatesRequireAutomaticUpdates == nil || *p.UpdatesRequireAutomaticUpdates != *input.UpdatesRequireAutomaticUpdates) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	if p.WorkFoldersUrl != nil && (input.WorkFoldersUrl == nil || *p.WorkFoldersUrl != *input.WorkFoldersUrl) {
		return false
	}

	return true
}
