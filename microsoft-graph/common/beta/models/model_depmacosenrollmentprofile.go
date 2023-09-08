package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepMacOSEnrollmentProfile struct {
	AccessibilityScreenDisabled                         *bool     `json:"accessibilityScreenDisabled,omitempty"`
	AdminAccountFullName                                *string   `json:"adminAccountFullName,omitempty"`
	AdminAccountPassword                                *string   `json:"adminAccountPassword,omitempty"`
	AdminAccountUserName                                *string   `json:"adminAccountUserName,omitempty"`
	AppleIdDisabled                                     *bool     `json:"appleIdDisabled,omitempty"`
	ApplePayDisabled                                    *bool     `json:"applePayDisabled,omitempty"`
	AutoAdvanceSetupEnabled                             *bool     `json:"autoAdvanceSetupEnabled,omitempty"`
	AutoUnlockWithWatchDisabled                         *bool     `json:"autoUnlockWithWatchDisabled,omitempty"`
	ChooseYourLockScreenDisabled                        *bool     `json:"chooseYourLockScreenDisabled,omitempty"`
	ConfigurationEndpointUrl                            *string   `json:"configurationEndpointUrl,omitempty"`
	ConfigurationWebUrl                                 *bool     `json:"configurationWebUrl,omitempty"`
	Description                                         *string   `json:"description,omitempty"`
	DeviceNameTemplate                                  *string   `json:"deviceNameTemplate,omitempty"`
	DiagnosticsDisabled                                 *bool     `json:"diagnosticsDisabled,omitempty"`
	DisplayName                                         *string   `json:"displayName,omitempty"`
	DisplayToneSetupDisabled                            *bool     `json:"displayToneSetupDisabled,omitempty"`
	DontAutoPopulatePrimaryAccountInfo                  *bool     `json:"dontAutoPopulatePrimaryAccountInfo,omitempty"`
	EnableAuthenticationViaCompanyPortal                *bool     `json:"enableAuthenticationViaCompanyPortal,omitempty"`
	EnableRestrictEditing                               *bool     `json:"enableRestrictEditing,omitempty"`
	EnabledSkipKeys                                     *[]string `json:"enabledSkipKeys,omitempty"`
	EnrollmentTimeAzureAdGroupIds                       *[]string `json:"enrollmentTimeAzureAdGroupIds,omitempty"`
	FileVaultDisabled                                   *bool     `json:"fileVaultDisabled,omitempty"`
	HideAdminAccount                                    *bool     `json:"hideAdminAccount,omitempty"`
	ICloudDiagnosticsDisabled                           *bool     `json:"iCloudDiagnosticsDisabled,omitempty"`
	ICloudStorageDisabled                               *bool     `json:"iCloudStorageDisabled,omitempty"`
	Id                                                  *string   `json:"id,omitempty"`
	IsDefault                                           *bool     `json:"isDefault,omitempty"`
	IsMandatory                                         *bool     `json:"isMandatory,omitempty"`
	LocationDisabled                                    *bool     `json:"locationDisabled,omitempty"`
	ODataType                                           *string   `json:"@odata.type,omitempty"`
	PassCodeDisabled                                    *bool     `json:"passCodeDisabled,omitempty"`
	PrimaryAccountFullName                              *string   `json:"primaryAccountFullName,omitempty"`
	PrimaryAccountUserName                              *string   `json:"primaryAccountUserName,omitempty"`
	PrivacyPaneDisabled                                 *bool     `json:"privacyPaneDisabled,omitempty"`
	ProfileRemovalDisabled                              *bool     `json:"profileRemovalDisabled,omitempty"`
	RegistrationDisabled                                *bool     `json:"registrationDisabled,omitempty"`
	RequestRequiresNetworkTether                        *bool     `json:"requestRequiresNetworkTether,omitempty"`
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool     `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`
	RequiresUserAuthentication                          *bool     `json:"requiresUserAuthentication,omitempty"`
	RestoreBlocked                                      *bool     `json:"restoreBlocked,omitempty"`
	ScreenTimeScreenDisabled                            *bool     `json:"screenTimeScreenDisabled,omitempty"`
	SetPrimarySetupAccountAsRegularUser                 *bool     `json:"setPrimarySetupAccountAsRegularUser,omitempty"`
	SiriDisabled                                        *bool     `json:"siriDisabled,omitempty"`
	SkipPrimarySetupAccountCreation                     *bool     `json:"skipPrimarySetupAccountCreation,omitempty"`
	SupervisedModeEnabled                               *bool     `json:"supervisedModeEnabled,omitempty"`
	SupportDepartment                                   *string   `json:"supportDepartment,omitempty"`
	SupportPhoneNumber                                  *string   `json:"supportPhoneNumber,omitempty"`
	TermsAndConditionsDisabled                          *bool     `json:"termsAndConditionsDisabled,omitempty"`
	TouchIdDisabled                                     *bool     `json:"touchIdDisabled,omitempty"`
	ZoomDisabled                                        *bool     `json:"zoomDisabled,omitempty"`
}
