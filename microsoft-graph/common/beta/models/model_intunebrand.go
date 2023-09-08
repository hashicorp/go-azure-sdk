package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrand struct {
	CompanyPortalBlockedActions               *[]CompanyPortalBlockedAction      `json:"companyPortalBlockedActions,omitempty"`
	ContactITEmailAddress                     *string                            `json:"contactITEmailAddress,omitempty"`
	ContactITName                             *string                            `json:"contactITName,omitempty"`
	ContactITNotes                            *string                            `json:"contactITNotes,omitempty"`
	ContactITPhoneNumber                      *string                            `json:"contactITPhoneNumber,omitempty"`
	CustomCanSeePrivacyMessage                *string                            `json:"customCanSeePrivacyMessage,omitempty"`
	CustomCantSeePrivacyMessage               *string                            `json:"customCantSeePrivacyMessage,omitempty"`
	CustomPrivacyMessage                      *string                            `json:"customPrivacyMessage,omitempty"`
	DarkBackgroundLogo                        *MimeContent                       `json:"darkBackgroundLogo,omitempty"`
	DisableClientTelemetry                    *bool                              `json:"disableClientTelemetry,omitempty"`
	DisableDeviceCategorySelection            *bool                              `json:"disableDeviceCategorySelection,omitempty"`
	DisplayName                               *string                            `json:"displayName,omitempty"`
	EnrollmentAvailability                    *IntuneBrandEnrollmentAvailability `json:"enrollmentAvailability,omitempty"`
	IsFactoryResetDisabled                    *bool                              `json:"isFactoryResetDisabled,omitempty"`
	IsRemoveDeviceDisabled                    *bool                              `json:"isRemoveDeviceDisabled,omitempty"`
	LandingPageCustomizedImage                *MimeContent                       `json:"landingPageCustomizedImage,omitempty"`
	LightBackgroundLogo                       *MimeContent                       `json:"lightBackgroundLogo,omitempty"`
	ODataType                                 *string                            `json:"@odata.type,omitempty"`
	OnlineSupportSiteName                     *string                            `json:"onlineSupportSiteName,omitempty"`
	OnlineSupportSiteUrl                      *string                            `json:"onlineSupportSiteUrl,omitempty"`
	PrivacyUrl                                *string                            `json:"privacyUrl,omitempty"`
	RoleScopeTagIds                           *[]string                          `json:"roleScopeTagIds,omitempty"`
	SendDeviceOwnershipChangePushNotification *bool                              `json:"sendDeviceOwnershipChangePushNotification,omitempty"`
	ShowAzureADEnterpriseApps                 *bool                              `json:"showAzureADEnterpriseApps,omitempty"`
	ShowConfigurationManagerApps              *bool                              `json:"showConfigurationManagerApps,omitempty"`
	ShowDisplayNameNextToLogo                 *bool                              `json:"showDisplayNameNextToLogo,omitempty"`
	ShowLogo                                  *bool                              `json:"showLogo,omitempty"`
	ShowNameNextToLogo                        *bool                              `json:"showNameNextToLogo,omitempty"`
	ShowOfficeWebApps                         *bool                              `json:"showOfficeWebApps,omitempty"`
	ThemeColor                                *RgbColor                          `json:"themeColor,omitempty"`
}
