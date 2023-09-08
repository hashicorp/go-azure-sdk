package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandOperationPredicate struct {
	ContactITEmailAddress                     *string
	ContactITName                             *string
	ContactITNotes                            *string
	ContactITPhoneNumber                      *string
	CustomCanSeePrivacyMessage                *string
	CustomCantSeePrivacyMessage               *string
	CustomPrivacyMessage                      *string
	DisableClientTelemetry                    *bool
	DisableDeviceCategorySelection            *bool
	DisplayName                               *string
	IsFactoryResetDisabled                    *bool
	IsRemoveDeviceDisabled                    *bool
	ODataType                                 *string
	OnlineSupportSiteName                     *string
	OnlineSupportSiteUrl                      *string
	PrivacyUrl                                *string
	SendDeviceOwnershipChangePushNotification *bool
	ShowAzureADEnterpriseApps                 *bool
	ShowConfigurationManagerApps              *bool
	ShowDisplayNameNextToLogo                 *bool
	ShowLogo                                  *bool
	ShowNameNextToLogo                        *bool
	ShowOfficeWebApps                         *bool
}

func (p IntuneBrandOperationPredicate) Matches(input IntuneBrand) bool {

	if p.ContactITEmailAddress != nil && (input.ContactITEmailAddress == nil || *p.ContactITEmailAddress != *input.ContactITEmailAddress) {
		return false
	}

	if p.ContactITName != nil && (input.ContactITName == nil || *p.ContactITName != *input.ContactITName) {
		return false
	}

	if p.ContactITNotes != nil && (input.ContactITNotes == nil || *p.ContactITNotes != *input.ContactITNotes) {
		return false
	}

	if p.ContactITPhoneNumber != nil && (input.ContactITPhoneNumber == nil || *p.ContactITPhoneNumber != *input.ContactITPhoneNumber) {
		return false
	}

	if p.CustomCanSeePrivacyMessage != nil && (input.CustomCanSeePrivacyMessage == nil || *p.CustomCanSeePrivacyMessage != *input.CustomCanSeePrivacyMessage) {
		return false
	}

	if p.CustomCantSeePrivacyMessage != nil && (input.CustomCantSeePrivacyMessage == nil || *p.CustomCantSeePrivacyMessage != *input.CustomCantSeePrivacyMessage) {
		return false
	}

	if p.CustomPrivacyMessage != nil && (input.CustomPrivacyMessage == nil || *p.CustomPrivacyMessage != *input.CustomPrivacyMessage) {
		return false
	}

	if p.DisableClientTelemetry != nil && (input.DisableClientTelemetry == nil || *p.DisableClientTelemetry != *input.DisableClientTelemetry) {
		return false
	}

	if p.DisableDeviceCategorySelection != nil && (input.DisableDeviceCategorySelection == nil || *p.DisableDeviceCategorySelection != *input.DisableDeviceCategorySelection) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.IsFactoryResetDisabled != nil && (input.IsFactoryResetDisabled == nil || *p.IsFactoryResetDisabled != *input.IsFactoryResetDisabled) {
		return false
	}

	if p.IsRemoveDeviceDisabled != nil && (input.IsRemoveDeviceDisabled == nil || *p.IsRemoveDeviceDisabled != *input.IsRemoveDeviceDisabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnlineSupportSiteName != nil && (input.OnlineSupportSiteName == nil || *p.OnlineSupportSiteName != *input.OnlineSupportSiteName) {
		return false
	}

	if p.OnlineSupportSiteUrl != nil && (input.OnlineSupportSiteUrl == nil || *p.OnlineSupportSiteUrl != *input.OnlineSupportSiteUrl) {
		return false
	}

	if p.PrivacyUrl != nil && (input.PrivacyUrl == nil || *p.PrivacyUrl != *input.PrivacyUrl) {
		return false
	}

	if p.SendDeviceOwnershipChangePushNotification != nil && (input.SendDeviceOwnershipChangePushNotification == nil || *p.SendDeviceOwnershipChangePushNotification != *input.SendDeviceOwnershipChangePushNotification) {
		return false
	}

	if p.ShowAzureADEnterpriseApps != nil && (input.ShowAzureADEnterpriseApps == nil || *p.ShowAzureADEnterpriseApps != *input.ShowAzureADEnterpriseApps) {
		return false
	}

	if p.ShowConfigurationManagerApps != nil && (input.ShowConfigurationManagerApps == nil || *p.ShowConfigurationManagerApps != *input.ShowConfigurationManagerApps) {
		return false
	}

	if p.ShowDisplayNameNextToLogo != nil && (input.ShowDisplayNameNextToLogo == nil || *p.ShowDisplayNameNextToLogo != *input.ShowDisplayNameNextToLogo) {
		return false
	}

	if p.ShowLogo != nil && (input.ShowLogo == nil || *p.ShowLogo != *input.ShowLogo) {
		return false
	}

	if p.ShowNameNextToLogo != nil && (input.ShowNameNextToLogo == nil || *p.ShowNameNextToLogo != *input.ShowNameNextToLogo) {
		return false
	}

	if p.ShowOfficeWebApps != nil && (input.ShowOfficeWebApps == nil || *p.ShowOfficeWebApps != *input.ShowOfficeWebApps) {
		return false
	}

	return true
}
