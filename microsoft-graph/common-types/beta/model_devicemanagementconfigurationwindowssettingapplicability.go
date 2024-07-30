package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicability struct {
	ConfigurationServiceProviderVersion *string                                                                           `json:"configurationServiceProviderVersion,omitempty"`
	Description                         *string                                                                           `json:"description,omitempty"`
	DeviceMode                          *DeviceManagementConfigurationWindowsSettingApplicabilityDeviceMode               `json:"deviceMode,omitempty"`
	MaximumSupportedVersion             *string                                                                           `json:"maximumSupportedVersion,omitempty"`
	MinimumSupportedVersion             *string                                                                           `json:"minimumSupportedVersion,omitempty"`
	ODataType                           *string                                                                           `json:"@odata.type,omitempty"`
	Platform                            *DeviceManagementConfigurationWindowsSettingApplicabilityPlatform                 `json:"platform,omitempty"`
	RequiredAzureAdTrustType            *DeviceManagementConfigurationWindowsSettingApplicabilityRequiredAzureAdTrustType `json:"requiredAzureAdTrustType,omitempty"`
	RequiresAzureAd                     *bool                                                                             `json:"requiresAzureAd,omitempty"`
	Technologies                        *DeviceManagementConfigurationWindowsSettingApplicabilityTechnologies             `json:"technologies,omitempty"`
	WindowsSkus                         *DeviceManagementConfigurationWindowsSettingApplicabilityWindowsSkus              `json:"windowsSkus,omitempty"`
}
