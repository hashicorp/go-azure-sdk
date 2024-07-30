package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepIOSEnrollmentProfile struct {
	AppearanceScreenDisabled                            *bool                                     `json:"appearanceScreenDisabled,omitempty"`
	AppleIdDisabled                                     *bool                                     `json:"appleIdDisabled,omitempty"`
	ApplePayDisabled                                    *bool                                     `json:"applePayDisabled,omitempty"`
	AwaitDeviceConfiguredConfirmation                   *bool                                     `json:"awaitDeviceConfiguredConfirmation,omitempty"`
	CarrierActivationUrl                                *string                                   `json:"carrierActivationUrl,omitempty"`
	CompanyPortalVppTokenId                             *string                                   `json:"companyPortalVppTokenId,omitempty"`
	ConfigurationEndpointUrl                            *string                                   `json:"configurationEndpointUrl,omitempty"`
	ConfigurationWebUrl                                 *bool                                     `json:"configurationWebUrl,omitempty"`
	Description                                         *string                                   `json:"description,omitempty"`
	DeviceNameTemplate                                  *string                                   `json:"deviceNameTemplate,omitempty"`
	DeviceToDeviceMigrationDisabled                     *bool                                     `json:"deviceToDeviceMigrationDisabled,omitempty"`
	DiagnosticsDisabled                                 *bool                                     `json:"diagnosticsDisabled,omitempty"`
	DisplayName                                         *string                                   `json:"displayName,omitempty"`
	DisplayToneSetupDisabled                            *bool                                     `json:"displayToneSetupDisabled,omitempty"`
	EnableAuthenticationViaCompanyPortal                *bool                                     `json:"enableAuthenticationViaCompanyPortal,omitempty"`
	EnableSharedIPad                                    *bool                                     `json:"enableSharedIPad,omitempty"`
	EnableSingleAppEnrollmentMode                       *bool                                     `json:"enableSingleAppEnrollmentMode,omitempty"`
	EnabledSkipKeys                                     *[]string                                 `json:"enabledSkipKeys,omitempty"`
	EnrollmentTimeAzureAdGroupIds                       *[]string                                 `json:"enrollmentTimeAzureAdGroupIds,omitempty"`
	ExpressLanguageScreenDisabled                       *bool                                     `json:"expressLanguageScreenDisabled,omitempty"`
	ForceTemporarySession                               *bool                                     `json:"forceTemporarySession,omitempty"`
	HomeButtonScreenDisabled                            *bool                                     `json:"homeButtonScreenDisabled,omitempty"`
	IMessageAndFaceTimeScreenDisabled                   *bool                                     `json:"iMessageAndFaceTimeScreenDisabled,omitempty"`
	ITunesPairingMode                                   *DepIOSEnrollmentProfileITunesPairingMode `json:"iTunesPairingMode,omitempty"`
	Id                                                  *string                                   `json:"id,omitempty"`
	IsDefault                                           *bool                                     `json:"isDefault,omitempty"`
	IsMandatory                                         *bool                                     `json:"isMandatory,omitempty"`
	LocationDisabled                                    *bool                                     `json:"locationDisabled,omitempty"`
	ManagementCertificates                              *[]ManagementCertificateWithThumbprint    `json:"managementCertificates,omitempty"`
	ODataType                                           *string                                   `json:"@odata.type,omitempty"`
	OnBoardingScreenDisabled                            *bool                                     `json:"onBoardingScreenDisabled,omitempty"`
	PassCodeDisabled                                    *bool                                     `json:"passCodeDisabled,omitempty"`
	PasscodeLockGracePeriodInSeconds                    *int64                                    `json:"passcodeLockGracePeriodInSeconds,omitempty"`
	PreferredLanguageScreenDisabled                     *bool                                     `json:"preferredLanguageScreenDisabled,omitempty"`
	PrivacyPaneDisabled                                 *bool                                     `json:"privacyPaneDisabled,omitempty"`
	ProfileRemovalDisabled                              *bool                                     `json:"profileRemovalDisabled,omitempty"`
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool                                     `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`
	RequiresUserAuthentication                          *bool                                     `json:"requiresUserAuthentication,omitempty"`
	RestoreBlocked                                      *bool                                     `json:"restoreBlocked,omitempty"`
	RestoreCompletedScreenDisabled                      *bool                                     `json:"restoreCompletedScreenDisabled,omitempty"`
	RestoreFromAndroidDisabled                          *bool                                     `json:"restoreFromAndroidDisabled,omitempty"`
	ScreenTimeScreenDisabled                            *bool                                     `json:"screenTimeScreenDisabled,omitempty"`
	SharedIPadMaximumUserCount                          *int64                                    `json:"sharedIPadMaximumUserCount,omitempty"`
	SimSetupScreenDisabled                              *bool                                     `json:"simSetupScreenDisabled,omitempty"`
	SiriDisabled                                        *bool                                     `json:"siriDisabled,omitempty"`
	SoftwareUpdateScreenDisabled                        *bool                                     `json:"softwareUpdateScreenDisabled,omitempty"`
	SupervisedModeEnabled                               *bool                                     `json:"supervisedModeEnabled,omitempty"`
	SupportDepartment                                   *string                                   `json:"supportDepartment,omitempty"`
	SupportPhoneNumber                                  *string                                   `json:"supportPhoneNumber,omitempty"`
	TemporarySessionTimeoutInSeconds                    *int64                                    `json:"temporarySessionTimeoutInSeconds,omitempty"`
	TermsAndConditionsDisabled                          *bool                                     `json:"termsAndConditionsDisabled,omitempty"`
	TouchIdDisabled                                     *bool                                     `json:"touchIdDisabled,omitempty"`
	UpdateCompleteScreenDisabled                        *bool                                     `json:"updateCompleteScreenDisabled,omitempty"`
	UserSessionTimeoutInSeconds                         *int64                                    `json:"userSessionTimeoutInSeconds,omitempty"`
	UserlessSharedAadModeEnabled                        *bool                                     `json:"userlessSharedAadModeEnabled,omitempty"`
	WaitForDeviceConfiguredConfirmation                 *bool                                     `json:"waitForDeviceConfiguredConfirmation,omitempty"`
	WatchMigrationScreenDisabled                        *bool                                     `json:"watchMigrationScreenDisabled,omitempty"`
	WelcomeScreenDisabled                               *bool                                     `json:"welcomeScreenDisabled,omitempty"`
	ZoomDisabled                                        *bool                                     `json:"zoomDisabled,omitempty"`
}
