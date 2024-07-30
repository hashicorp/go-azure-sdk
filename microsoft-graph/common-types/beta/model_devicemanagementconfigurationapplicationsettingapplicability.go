package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationApplicationSettingApplicability struct {
	Description  *string                                                                   `json:"description,omitempty"`
	DeviceMode   *DeviceManagementConfigurationApplicationSettingApplicabilityDeviceMode   `json:"deviceMode,omitempty"`
	ODataType    *string                                                                   `json:"@odata.type,omitempty"`
	Platform     *DeviceManagementConfigurationApplicationSettingApplicabilityPlatform     `json:"platform,omitempty"`
	Technologies *DeviceManagementConfigurationApplicationSettingApplicabilityTechnologies `json:"technologies,omitempty"`
}
