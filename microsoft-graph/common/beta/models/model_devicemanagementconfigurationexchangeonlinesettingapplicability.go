package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicability struct {
	Description  *string                                                                      `json:"description,omitempty"`
	DeviceMode   *DeviceManagementConfigurationExchangeOnlineSettingApplicabilityDeviceMode   `json:"deviceMode,omitempty"`
	ODataType    *string                                                                      `json:"@odata.type,omitempty"`
	Platform     *DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform     `json:"platform,omitempty"`
	Technologies *DeviceManagementConfigurationExchangeOnlineSettingApplicabilityTechnologies `json:"technologies,omitempty"`
}
