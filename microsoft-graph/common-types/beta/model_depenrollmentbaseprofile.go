package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepEnrollmentBaseProfile struct {
	AppleIdDisabled                                     *bool     `json:"appleIdDisabled,omitempty"`
	ApplePayDisabled                                    *bool     `json:"applePayDisabled,omitempty"`
	ConfigurationEndpointUrl                            *string   `json:"configurationEndpointUrl,omitempty"`
	ConfigurationWebUrl                                 *bool     `json:"configurationWebUrl,omitempty"`
	Description                                         *string   `json:"description,omitempty"`
	DeviceNameTemplate                                  *string   `json:"deviceNameTemplate,omitempty"`
	DiagnosticsDisabled                                 *bool     `json:"diagnosticsDisabled,omitempty"`
	DisplayName                                         *string   `json:"displayName,omitempty"`
	DisplayToneSetupDisabled                            *bool     `json:"displayToneSetupDisabled,omitempty"`
	EnableAuthenticationViaCompanyPortal                *bool     `json:"enableAuthenticationViaCompanyPortal,omitempty"`
	EnabledSkipKeys                                     *[]string `json:"enabledSkipKeys,omitempty"`
	EnrollmentTimeAzureAdGroupIds                       *[]string `json:"enrollmentTimeAzureAdGroupIds,omitempty"`
	Id                                                  *string   `json:"id,omitempty"`
	IsDefault                                           *bool     `json:"isDefault,omitempty"`
	IsMandatory                                         *bool     `json:"isMandatory,omitempty"`
	LocationDisabled                                    *bool     `json:"locationDisabled,omitempty"`
	ODataType                                           *string   `json:"@odata.type,omitempty"`
	PrivacyPaneDisabled                                 *bool     `json:"privacyPaneDisabled,omitempty"`
	ProfileRemovalDisabled                              *bool     `json:"profileRemovalDisabled,omitempty"`
	RequireCompanyPortalOnSetupAssistantEnrolledDevices *bool     `json:"requireCompanyPortalOnSetupAssistantEnrolledDevices,omitempty"`
	RequiresUserAuthentication                          *bool     `json:"requiresUserAuthentication,omitempty"`
	RestoreBlocked                                      *bool     `json:"restoreBlocked,omitempty"`
	ScreenTimeScreenDisabled                            *bool     `json:"screenTimeScreenDisabled,omitempty"`
	SiriDisabled                                        *bool     `json:"siriDisabled,omitempty"`
	SupervisedModeEnabled                               *bool     `json:"supervisedModeEnabled,omitempty"`
	SupportDepartment                                   *string   `json:"supportDepartment,omitempty"`
	SupportPhoneNumber                                  *string   `json:"supportPhoneNumber,omitempty"`
	TermsAndConditionsDisabled                          *bool     `json:"termsAndConditionsDisabled,omitempty"`
	TouchIdDisabled                                     *bool     `json:"touchIdDisabled,omitempty"`
	WaitForDeviceConfiguredConfirmation                 *bool     `json:"waitForDeviceConfiguredConfirmation,omitempty"`
}
