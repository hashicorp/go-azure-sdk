package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepIOSEnrollmentProfileOperationPredicate struct {
	AppearanceScreenDisabled                            *bool
	AppleIdDisabled                                     *bool
	ApplePayDisabled                                    *bool
	AwaitDeviceConfiguredConfirmation                   *bool
	CarrierActivationUrl                                *string
	CompanyPortalVppTokenId                             *string
	ConfigurationEndpointUrl                            *string
	ConfigurationWebUrl                                 *bool
	Description                                         *string
	DeviceNameTemplate                                  *string
	DeviceToDeviceMigrationDisabled                     *bool
	DiagnosticsDisabled                                 *bool
	DisplayName                                         *string
	DisplayToneSetupDisabled                            *bool
	EnableAuthenticationViaCompanyPortal                *bool
	EnableSharedIPad                                    *bool
	EnableSingleAppEnrollmentMode                       *bool
	ExpressLanguageScreenDisabled                       *bool
	ForceTemporarySession                               *bool
	HomeButtonScreenDisabled                            *bool
	IMessageAndFaceTimeScreenDisabled                   *bool
	Id                                                  *string
	IsDefault                                           *bool
	IsMandatory                                         *bool
	LocationDisabled                                    *bool
	ODataType                                           *string
	OnBoardingScreenDisabled                            *bool
	PassCodeDisabled                                    *bool
	PasscodeLockGracePeriodInSeconds                    *int64
	PreferredLanguageScreenDisabled                     *bool
	PrivacyPaneDisabled                                 *bool
	ProfileRemovalDisabled                              *bool
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool
	RequiresUserAuthentication                          *bool
	RestoreBlocked                                      *bool
	RestoreCompletedScreenDisabled                      *bool
	RestoreFromAndroidDisabled                          *bool
	ScreenTimeScreenDisabled                            *bool
	SharedIPadMaximumUserCount                          *int64
	SimSetupScreenDisabled                              *bool
	SiriDisabled                                        *bool
	SoftwareUpdateScreenDisabled                        *bool
	SupervisedModeEnabled                               *bool
	SupportDepartment                                   *string
	SupportPhoneNumber                                  *string
	TemporarySessionTimeoutInSeconds                    *int64
	TermsAndConditionsDisabled                          *bool
	TouchIdDisabled                                     *bool
	UpdateCompleteScreenDisabled                        *bool
	UserSessionTimeoutInSeconds                         *int64
	UserlessSharedAadModeEnabled                        *bool
	WatchMigrationScreenDisabled                        *bool
	WelcomeScreenDisabled                               *bool
	ZoomDisabled                                        *bool
}

func (p DepIOSEnrollmentProfileOperationPredicate) Matches(input DepIOSEnrollmentProfile) bool {

	if p.AppearanceScreenDisabled != nil && (input.AppearanceScreenDisabled == nil || *p.AppearanceScreenDisabled != *input.AppearanceScreenDisabled) {
		return false
	}

	if p.AppleIdDisabled != nil && (input.AppleIdDisabled == nil || *p.AppleIdDisabled != *input.AppleIdDisabled) {
		return false
	}

	if p.ApplePayDisabled != nil && (input.ApplePayDisabled == nil || *p.ApplePayDisabled != *input.ApplePayDisabled) {
		return false
	}

	if p.AwaitDeviceConfiguredConfirmation != nil && (input.AwaitDeviceConfiguredConfirmation == nil || *p.AwaitDeviceConfiguredConfirmation != *input.AwaitDeviceConfiguredConfirmation) {
		return false
	}

	if p.CarrierActivationUrl != nil && (input.CarrierActivationUrl == nil || *p.CarrierActivationUrl != *input.CarrierActivationUrl) {
		return false
	}

	if p.CompanyPortalVppTokenId != nil && (input.CompanyPortalVppTokenId == nil || *p.CompanyPortalVppTokenId != *input.CompanyPortalVppTokenId) {
		return false
	}

	if p.ConfigurationEndpointUrl != nil && (input.ConfigurationEndpointUrl == nil || *p.ConfigurationEndpointUrl != *input.ConfigurationEndpointUrl) {
		return false
	}

	if p.ConfigurationWebUrl != nil && (input.ConfigurationWebUrl == nil || *p.ConfigurationWebUrl != *input.ConfigurationWebUrl) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceNameTemplate != nil && (input.DeviceNameTemplate == nil || *p.DeviceNameTemplate != *input.DeviceNameTemplate) {
		return false
	}

	if p.DeviceToDeviceMigrationDisabled != nil && (input.DeviceToDeviceMigrationDisabled == nil || *p.DeviceToDeviceMigrationDisabled != *input.DeviceToDeviceMigrationDisabled) {
		return false
	}

	if p.DiagnosticsDisabled != nil && (input.DiagnosticsDisabled == nil || *p.DiagnosticsDisabled != *input.DiagnosticsDisabled) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DisplayToneSetupDisabled != nil && (input.DisplayToneSetupDisabled == nil || *p.DisplayToneSetupDisabled != *input.DisplayToneSetupDisabled) {
		return false
	}

	if p.EnableAuthenticationViaCompanyPortal != nil && (input.EnableAuthenticationViaCompanyPortal == nil || *p.EnableAuthenticationViaCompanyPortal != *input.EnableAuthenticationViaCompanyPortal) {
		return false
	}

	if p.EnableSharedIPad != nil && (input.EnableSharedIPad == nil || *p.EnableSharedIPad != *input.EnableSharedIPad) {
		return false
	}

	if p.EnableSingleAppEnrollmentMode != nil && (input.EnableSingleAppEnrollmentMode == nil || *p.EnableSingleAppEnrollmentMode != *input.EnableSingleAppEnrollmentMode) {
		return false
	}

	if p.ExpressLanguageScreenDisabled != nil && (input.ExpressLanguageScreenDisabled == nil || *p.ExpressLanguageScreenDisabled != *input.ExpressLanguageScreenDisabled) {
		return false
	}

	if p.ForceTemporarySession != nil && (input.ForceTemporarySession == nil || *p.ForceTemporarySession != *input.ForceTemporarySession) {
		return false
	}

	if p.HomeButtonScreenDisabled != nil && (input.HomeButtonScreenDisabled == nil || *p.HomeButtonScreenDisabled != *input.HomeButtonScreenDisabled) {
		return false
	}

	if p.IMessageAndFaceTimeScreenDisabled != nil && (input.IMessageAndFaceTimeScreenDisabled == nil || *p.IMessageAndFaceTimeScreenDisabled != *input.IMessageAndFaceTimeScreenDisabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefault != nil && (input.IsDefault == nil || *p.IsDefault != *input.IsDefault) {
		return false
	}

	if p.IsMandatory != nil && (input.IsMandatory == nil || *p.IsMandatory != *input.IsMandatory) {
		return false
	}

	if p.LocationDisabled != nil && (input.LocationDisabled == nil || *p.LocationDisabled != *input.LocationDisabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnBoardingScreenDisabled != nil && (input.OnBoardingScreenDisabled == nil || *p.OnBoardingScreenDisabled != *input.OnBoardingScreenDisabled) {
		return false
	}

	if p.PassCodeDisabled != nil && (input.PassCodeDisabled == nil || *p.PassCodeDisabled != *input.PassCodeDisabled) {
		return false
	}

	if p.PasscodeLockGracePeriodInSeconds != nil && (input.PasscodeLockGracePeriodInSeconds == nil || *p.PasscodeLockGracePeriodInSeconds != *input.PasscodeLockGracePeriodInSeconds) {
		return false
	}

	if p.PreferredLanguageScreenDisabled != nil && (input.PreferredLanguageScreenDisabled == nil || *p.PreferredLanguageScreenDisabled != *input.PreferredLanguageScreenDisabled) {
		return false
	}

	if p.PrivacyPaneDisabled != nil && (input.PrivacyPaneDisabled == nil || *p.PrivacyPaneDisabled != *input.PrivacyPaneDisabled) {
		return false
	}

	if p.ProfileRemovalDisabled != nil && (input.ProfileRemovalDisabled == nil || *p.ProfileRemovalDisabled != *input.ProfileRemovalDisabled) {
		return false
	}

	if p.RequireCompanyPortalOnSetupAssistantEnrolledDevices != nil && (input.RequireCompanyPortalOnSetupAssistantEnrolledDevices == nil || *p.RequireCompanyPortalOnSetupAssistantEnrolledDevices != *input.RequireCompanyPortalOnSetupAssistantEnrolledDevices) {
		return false
	}

	if p.RequiresUserAuthentication != nil && (input.RequiresUserAuthentication == nil || *p.RequiresUserAuthentication != *input.RequiresUserAuthentication) {
		return false
	}

	if p.RestoreBlocked != nil && (input.RestoreBlocked == nil || *p.RestoreBlocked != *input.RestoreBlocked) {
		return false
	}

	if p.RestoreCompletedScreenDisabled != nil && (input.RestoreCompletedScreenDisabled == nil || *p.RestoreCompletedScreenDisabled != *input.RestoreCompletedScreenDisabled) {
		return false
	}

	if p.RestoreFromAndroidDisabled != nil && (input.RestoreFromAndroidDisabled == nil || *p.RestoreFromAndroidDisabled != *input.RestoreFromAndroidDisabled) {
		return false
	}

	if p.ScreenTimeScreenDisabled != nil && (input.ScreenTimeScreenDisabled == nil || *p.ScreenTimeScreenDisabled != *input.ScreenTimeScreenDisabled) {
		return false
	}

	if p.SharedIPadMaximumUserCount != nil && (input.SharedIPadMaximumUserCount == nil || *p.SharedIPadMaximumUserCount != *input.SharedIPadMaximumUserCount) {
		return false
	}

	if p.SimSetupScreenDisabled != nil && (input.SimSetupScreenDisabled == nil || *p.SimSetupScreenDisabled != *input.SimSetupScreenDisabled) {
		return false
	}

	if p.SiriDisabled != nil && (input.SiriDisabled == nil || *p.SiriDisabled != *input.SiriDisabled) {
		return false
	}

	if p.SoftwareUpdateScreenDisabled != nil && (input.SoftwareUpdateScreenDisabled == nil || *p.SoftwareUpdateScreenDisabled != *input.SoftwareUpdateScreenDisabled) {
		return false
	}

	if p.SupervisedModeEnabled != nil && (input.SupervisedModeEnabled == nil || *p.SupervisedModeEnabled != *input.SupervisedModeEnabled) {
		return false
	}

	if p.SupportDepartment != nil && (input.SupportDepartment == nil || *p.SupportDepartment != *input.SupportDepartment) {
		return false
	}

	if p.SupportPhoneNumber != nil && (input.SupportPhoneNumber == nil || *p.SupportPhoneNumber != *input.SupportPhoneNumber) {
		return false
	}

	if p.TemporarySessionTimeoutInSeconds != nil && (input.TemporarySessionTimeoutInSeconds == nil || *p.TemporarySessionTimeoutInSeconds != *input.TemporarySessionTimeoutInSeconds) {
		return false
	}

	if p.TermsAndConditionsDisabled != nil && (input.TermsAndConditionsDisabled == nil || *p.TermsAndConditionsDisabled != *input.TermsAndConditionsDisabled) {
		return false
	}

	if p.TouchIdDisabled != nil && (input.TouchIdDisabled == nil || *p.TouchIdDisabled != *input.TouchIdDisabled) {
		return false
	}

	if p.UpdateCompleteScreenDisabled != nil && (input.UpdateCompleteScreenDisabled == nil || *p.UpdateCompleteScreenDisabled != *input.UpdateCompleteScreenDisabled) {
		return false
	}

	if p.UserSessionTimeoutInSeconds != nil && (input.UserSessionTimeoutInSeconds == nil || *p.UserSessionTimeoutInSeconds != *input.UserSessionTimeoutInSeconds) {
		return false
	}

	if p.UserlessSharedAadModeEnabled != nil && (input.UserlessSharedAadModeEnabled == nil || *p.UserlessSharedAadModeEnabled != *input.UserlessSharedAadModeEnabled) {
		return false
	}

	if p.WatchMigrationScreenDisabled != nil && (input.WatchMigrationScreenDisabled == nil || *p.WatchMigrationScreenDisabled != *input.WatchMigrationScreenDisabled) {
		return false
	}

	if p.WelcomeScreenDisabled != nil && (input.WelcomeScreenDisabled == nil || *p.WelcomeScreenDisabled != *input.WelcomeScreenDisabled) {
		return false
	}

	if p.ZoomDisabled != nil && (input.ZoomDisabled == nil || *p.ZoomDisabled != *input.ZoomDisabled) {
		return false
	}

	return true
}
