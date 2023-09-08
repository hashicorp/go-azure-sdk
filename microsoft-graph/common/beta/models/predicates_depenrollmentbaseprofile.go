package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepEnrollmentBaseProfileOperationPredicate struct {
	AppleIdDisabled                                     *bool
	ApplePayDisabled                                    *bool
	ConfigurationEndpointUrl                            *string
	ConfigurationWebUrl                                 *bool
	Description                                         *string
	DeviceNameTemplate                                  *string
	DiagnosticsDisabled                                 *bool
	DisplayName                                         *string
	DisplayToneSetupDisabled                            *bool
	EnableAuthenticationViaCompanyPortal                *bool
	Id                                                  *string
	IsDefault                                           *bool
	IsMandatory                                         *bool
	LocationDisabled                                    *bool
	ODataType                                           *string
	PrivacyPaneDisabled                                 *bool
	ProfileRemovalDisabled                              *bool
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool
	RequiresUserAuthentication                          *bool
	RestoreBlocked                                      *bool
	ScreenTimeScreenDisabled                            *bool
	SiriDisabled                                        *bool
	SupervisedModeEnabled                               *bool
	SupportDepartment                                   *string
	SupportPhoneNumber                                  *string
	TermsAndConditionsDisabled                          *bool
	TouchIdDisabled                                     *bool
}

func (p DepEnrollmentBaseProfileOperationPredicate) Matches(input DepEnrollmentBaseProfile) bool {

	if p.AppleIdDisabled != nil && (input.AppleIdDisabled == nil || *p.AppleIdDisabled != *input.AppleIdDisabled) {
		return false
	}

	if p.ApplePayDisabled != nil && (input.ApplePayDisabled == nil || *p.ApplePayDisabled != *input.ApplePayDisabled) {
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

	if p.ScreenTimeScreenDisabled != nil && (input.ScreenTimeScreenDisabled == nil || *p.ScreenTimeScreenDisabled != *input.ScreenTimeScreenDisabled) {
		return false
	}

	if p.SiriDisabled != nil && (input.SiriDisabled == nil || *p.SiriDisabled != *input.SiriDisabled) {
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

	if p.TermsAndConditionsDisabled != nil && (input.TermsAndConditionsDisabled == nil || *p.TermsAndConditionsDisabled != *input.TermsAndConditionsDisabled) {
		return false
	}

	if p.TouchIdDisabled != nil && (input.TouchIdDisabled == nil || *p.TouchIdDisabled != *input.TouchIdDisabled) {
		return false
	}

	return true
}
