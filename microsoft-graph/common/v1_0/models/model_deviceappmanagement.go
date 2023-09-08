package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagement struct {
	AndroidManagedAppProtections                              *[]AndroidManagedAppProtection           `json:"androidManagedAppProtections,omitempty"`
	DefaultManagedAppProtections                              *[]DefaultManagedAppProtection           `json:"defaultManagedAppProtections,omitempty"`
	Id                                                        *string                                  `json:"id,omitempty"`
	IosManagedAppProtections                                  *[]IosManagedAppProtection               `json:"iosManagedAppProtections,omitempty"`
	IsEnabledForMicrosoftStoreForBusiness                     *bool                                    `json:"isEnabledForMicrosoftStoreForBusiness,omitempty"`
	ManagedAppPolicies                                        *[]ManagedAppPolicy                      `json:"managedAppPolicies,omitempty"`
	ManagedAppRegistrations                                   *[]ManagedAppRegistration                `json:"managedAppRegistrations,omitempty"`
	ManagedAppStatuses                                        *[]ManagedAppStatus                      `json:"managedAppStatuses,omitempty"`
	ManagedEBooks                                             *[]ManagedEBook                          `json:"managedEBooks,omitempty"`
	MdmWindowsInformationProtectionPolicies                   *[]MdmWindowsInformationProtectionPolicy `json:"mdmWindowsInformationProtectionPolicies,omitempty"`
	MicrosoftStoreForBusinessLanguage                         *string                                  `json:"microsoftStoreForBusinessLanguage,omitempty"`
	MicrosoftStoreForBusinessLastCompletedApplicationSyncTime *string                                  `json:"microsoftStoreForBusinessLastCompletedApplicationSyncTime,omitempty"`
	MicrosoftStoreForBusinessLastSuccessfulSyncDateTime       *string                                  `json:"microsoftStoreForBusinessLastSuccessfulSyncDateTime,omitempty"`
	MobileAppCategories                                       *[]MobileAppCategory                     `json:"mobileAppCategories,omitempty"`
	MobileAppConfigurations                                   *[]ManagedDeviceMobileAppConfiguration   `json:"mobileAppConfigurations,omitempty"`
	MobileApps                                                *[]MobileApp                             `json:"mobileApps,omitempty"`
	ODataType                                                 *string                                  `json:"@odata.type,omitempty"`
	TargetedManagedAppConfigurations                          *[]TargetedManagedAppConfiguration       `json:"targetedManagedAppConfigurations,omitempty"`
	VppTokens                                                 *[]VppToken                              `json:"vppTokens,omitempty"`
	WindowsInformationProtectionPolicies                      *[]WindowsInformationProtectionPolicy    `json:"windowsInformationProtectionPolicies,omitempty"`
}
