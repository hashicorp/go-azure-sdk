package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingApplicability struct {
	Description  *string                                                        `json:"description,omitempty"`
	DeviceMode   *DeviceManagementConfigurationSettingApplicabilityDeviceMode   `json:"deviceMode,omitempty"`
	ODataType    *string                                                        `json:"@odata.type,omitempty"`
	Platform     *DeviceManagementConfigurationSettingApplicabilityPlatform     `json:"platform,omitempty"`
	Technologies *DeviceManagementConfigurationSettingApplicabilityTechnologies `json:"technologies,omitempty"`
}
