package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandingProfile struct {
	Assignments                               *[]IntuneBrandingProfileAssignment           `json:"assignments,omitempty"`
	CompanyPortalBlockedActions               *[]CompanyPortalBlockedAction                `json:"companyPortalBlockedActions,omitempty"`
	ContactITEmailAddress                     *string                                      `json:"contactITEmailAddress,omitempty"`
	ContactITName                             *string                                      `json:"contactITName,omitempty"`
	ContactITNotes                            *string                                      `json:"contactITNotes,omitempty"`
	ContactITPhoneNumber                      *string                                      `json:"contactITPhoneNumber,omitempty"`
	CreatedDateTime                           *string                                      `json:"createdDateTime,omitempty"`
	CustomCanSeePrivacyMessage                *string                                      `json:"customCanSeePrivacyMessage,omitempty"`
	CustomCantSeePrivacyMessage               *string                                      `json:"customCantSeePrivacyMessage,omitempty"`
	CustomPrivacyMessage                      *string                                      `json:"customPrivacyMessage,omitempty"`
	DisableClientTelemetry                    *bool                                        `json:"disableClientTelemetry,omitempty"`
	DisableDeviceCategorySelection            *bool                                        `json:"disableDeviceCategorySelection,omitempty"`
	DisplayName                               *string                                      `json:"displayName,omitempty"`
	EnrollmentAvailability                    *IntuneBrandingProfileEnrollmentAvailability `json:"enrollmentAvailability,omitempty"`
	Id                                        *string                                      `json:"id,omitempty"`
	IsDefaultProfile                          *bool                                        `json:"isDefaultProfile,omitempty"`
	IsFactoryResetDisabled                    *bool                                        `json:"isFactoryResetDisabled,omitempty"`
	IsRemoveDeviceDisabled                    *bool                                        `json:"isRemoveDeviceDisabled,omitempty"`
	LandingPageCustomizedImage                *MimeContent                                 `json:"landingPageCustomizedImage,omitempty"`
	LastModifiedDateTime                      *string                                      `json:"lastModifiedDateTime,omitempty"`
	LightBackgroundLogo                       *MimeContent                                 `json:"lightBackgroundLogo,omitempty"`
	ODataType                                 *string                                      `json:"@odata.type,omitempty"`
	OnlineSupportSiteName                     *string                                      `json:"onlineSupportSiteName,omitempty"`
	OnlineSupportSiteUrl                      *string                                      `json:"onlineSupportSiteUrl,omitempty"`
	PrivacyUrl                                *string                                      `json:"privacyUrl,omitempty"`
	ProfileDescription                        *string                                      `json:"profileDescription,omitempty"`
	ProfileName                               *string                                      `json:"profileName,omitempty"`
	RoleScopeTagIds                           *[]string                                    `json:"roleScopeTagIds,omitempty"`
	SendDeviceOwnershipChangePushNotification *bool                                        `json:"sendDeviceOwnershipChangePushNotification,omitempty"`
	ShowAzureADEnterpriseApps                 *bool                                        `json:"showAzureADEnterpriseApps,omitempty"`
	ShowConfigurationManagerApps              *bool                                        `json:"showConfigurationManagerApps,omitempty"`
	ShowDisplayNameNextToLogo                 *bool                                        `json:"showDisplayNameNextToLogo,omitempty"`
	ShowLogo                                  *bool                                        `json:"showLogo,omitempty"`
	ShowOfficeWebApps                         *bool                                        `json:"showOfficeWebApps,omitempty"`
	ThemeColor                                *RgbColor                                    `json:"themeColor,omitempty"`
	ThemeColorLogo                            *MimeContent                                 `json:"themeColorLogo,omitempty"`
}
