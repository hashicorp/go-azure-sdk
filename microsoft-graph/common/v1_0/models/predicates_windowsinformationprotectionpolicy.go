package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyOperationPredicate struct {
	AzureRightsManagementServicesAllowed   *bool
	CreatedDateTime                        *string
	DaysWithoutContactBeforeUnenroll       *int64
	Description                            *string
	DisplayName                            *string
	EnterpriseDomain                       *string
	EnterpriseIPRangesAreAuthoritative     *bool
	EnterpriseProxyServersAreAuthoritative *bool
	IconsVisible                           *bool
	Id                                     *string
	IndexingEncryptedStoresOrItemsBlocked  *bool
	IsAssigned                             *bool
	LastModifiedDateTime                   *string
	MdmEnrollmentUrl                       *string
	MinutesOfInactivityBeforeDeviceLock    *int64
	NumberOfPastPinsRemembered             *int64
	ODataType                              *string
	PasswordMaximumAttemptCount            *int64
	PinExpirationDays                      *int64
	PinMinimumLength                       *int64
	ProtectionUnderLockConfigRequired      *bool
	RevokeOnMdmHandoffDisabled             *bool
	RevokeOnUnenrollDisabled               *bool
	RightsManagementServicesTemplateId     *string
	Version                                *string
	WindowsHelloForBusinessBlocked         *bool
}

func (p WindowsInformationProtectionPolicyOperationPredicate) Matches(input WindowsInformationProtectionPolicy) bool {

	if p.AzureRightsManagementServicesAllowed != nil && (input.AzureRightsManagementServicesAllowed == nil || *p.AzureRightsManagementServicesAllowed != *input.AzureRightsManagementServicesAllowed) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DaysWithoutContactBeforeUnenroll != nil && (input.DaysWithoutContactBeforeUnenroll == nil || *p.DaysWithoutContactBeforeUnenroll != *input.DaysWithoutContactBeforeUnenroll) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnterpriseDomain != nil && (input.EnterpriseDomain == nil || *p.EnterpriseDomain != *input.EnterpriseDomain) {
		return false
	}

	if p.EnterpriseIPRangesAreAuthoritative != nil && (input.EnterpriseIPRangesAreAuthoritative == nil || *p.EnterpriseIPRangesAreAuthoritative != *input.EnterpriseIPRangesAreAuthoritative) {
		return false
	}

	if p.EnterpriseProxyServersAreAuthoritative != nil && (input.EnterpriseProxyServersAreAuthoritative == nil || *p.EnterpriseProxyServersAreAuthoritative != *input.EnterpriseProxyServersAreAuthoritative) {
		return false
	}

	if p.IconsVisible != nil && (input.IconsVisible == nil || *p.IconsVisible != *input.IconsVisible) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IndexingEncryptedStoresOrItemsBlocked != nil && (input.IndexingEncryptedStoresOrItemsBlocked == nil || *p.IndexingEncryptedStoresOrItemsBlocked != *input.IndexingEncryptedStoresOrItemsBlocked) {
		return false
	}

	if p.IsAssigned != nil && (input.IsAssigned == nil || *p.IsAssigned != *input.IsAssigned) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MdmEnrollmentUrl != nil && (input.MdmEnrollmentUrl == nil || *p.MdmEnrollmentUrl != *input.MdmEnrollmentUrl) {
		return false
	}

	if p.MinutesOfInactivityBeforeDeviceLock != nil && (input.MinutesOfInactivityBeforeDeviceLock == nil || *p.MinutesOfInactivityBeforeDeviceLock != *input.MinutesOfInactivityBeforeDeviceLock) {
		return false
	}

	if p.NumberOfPastPinsRemembered != nil && (input.NumberOfPastPinsRemembered == nil || *p.NumberOfPastPinsRemembered != *input.NumberOfPastPinsRemembered) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PasswordMaximumAttemptCount != nil && (input.PasswordMaximumAttemptCount == nil || *p.PasswordMaximumAttemptCount != *input.PasswordMaximumAttemptCount) {
		return false
	}

	if p.PinExpirationDays != nil && (input.PinExpirationDays == nil || *p.PinExpirationDays != *input.PinExpirationDays) {
		return false
	}

	if p.PinMinimumLength != nil && (input.PinMinimumLength == nil || *p.PinMinimumLength != *input.PinMinimumLength) {
		return false
	}

	if p.ProtectionUnderLockConfigRequired != nil && (input.ProtectionUnderLockConfigRequired == nil || *p.ProtectionUnderLockConfigRequired != *input.ProtectionUnderLockConfigRequired) {
		return false
	}

	if p.RevokeOnMdmHandoffDisabled != nil && (input.RevokeOnMdmHandoffDisabled == nil || *p.RevokeOnMdmHandoffDisabled != *input.RevokeOnMdmHandoffDisabled) {
		return false
	}

	if p.RevokeOnUnenrollDisabled != nil && (input.RevokeOnUnenrollDisabled == nil || *p.RevokeOnUnenrollDisabled != *input.RevokeOnUnenrollDisabled) {
		return false
	}

	if p.RightsManagementServicesTemplateId != nil && (input.RightsManagementServicesTemplateId == nil || *p.RightsManagementServicesTemplateId != *input.RightsManagementServicesTemplateId) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	if p.WindowsHelloForBusinessBlocked != nil && (input.WindowsHelloForBusinessBlocked == nil || *p.WindowsHelloForBusinessBlocked != *input.WindowsHelloForBusinessBlocked) {
		return false
	}

	return true
}
