package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepEnrollmentProfile struct {
	AppleIdDisabled                                     *bool                                  `json:"appleIdDisabled,omitempty"`
	ApplePayDisabled                                    *bool                                  `json:"applePayDisabled,omitempty"`
	AwaitDeviceConfiguredConfirmation                   *bool                                  `json:"awaitDeviceConfiguredConfirmation,omitempty"`
	ConfigurationEndpointUrl                            *string                                `json:"configurationEndpointUrl,omitempty"`
	Description                                         *string                                `json:"description,omitempty"`
	DiagnosticsDisabled                                 *bool                                  `json:"diagnosticsDisabled,omitempty"`
	DisplayName                                         *string                                `json:"displayName,omitempty"`
	EnableAuthenticationViaCompanyPortal                *bool                                  `json:"enableAuthenticationViaCompanyPortal,omitempty"`
	EnableSharedIPad                                    *bool                                  `json:"enableSharedIPad,omitempty"`
	ITunesPairingMode                                   *DepEnrollmentProfileITunesPairingMode `json:"iTunesPairingMode,omitempty"`
	Id                                                  *string                                `json:"id,omitempty"`
	IsDefault                                           *bool                                  `json:"isDefault,omitempty"`
	IsMandatory                                         *bool                                  `json:"isMandatory,omitempty"`
	LocationDisabled                                    *bool                                  `json:"locationDisabled,omitempty"`
	MacOSFileVaultDisabled                              *bool                                  `json:"macOSFileVaultDisabled,omitempty"`
	MacOSRegistrationDisabled                           *bool                                  `json:"macOSRegistrationDisabled,omitempty"`
	ManagementCertificates                              *[]ManagementCertificateWithThumbprint `json:"managementCertificates,omitempty"`
	ODataType                                           *string                                `json:"@odata.type,omitempty"`
	PassCodeDisabled                                    *bool                                  `json:"passCodeDisabled,omitempty"`
	ProfileRemovalDisabled                              *bool                                  `json:"profileRemovalDisabled,omitempty"`
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool                                  `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`
	RequiresUserAuthentication                          *bool                                  `json:"requiresUserAuthentication,omitempty"`
	RestoreBlocked                                      *bool                                  `json:"restoreBlocked,omitempty"`
	RestoreFromAndroidDisabled                          *bool                                  `json:"restoreFromAndroidDisabled,omitempty"`
	SharedIPadMaximumUserCount                          *int64                                 `json:"sharedIPadMaximumUserCount,omitempty"`
	SiriDisabled                                        *bool                                  `json:"siriDisabled,omitempty"`
	SupervisedModeEnabled                               *bool                                  `json:"supervisedModeEnabled,omitempty"`
	SupportDepartment                                   *string                                `json:"supportDepartment,omitempty"`
	SupportPhoneNumber                                  *string                                `json:"supportPhoneNumber,omitempty"`
	TermsAndConditionsDisabled                          *bool                                  `json:"termsAndConditionsDisabled,omitempty"`
	TouchIdDisabled                                     *bool                                  `json:"touchIdDisabled,omitempty"`
	ZoomDisabled                                        *bool                                  `json:"zoomDisabled,omitempty"`
}
