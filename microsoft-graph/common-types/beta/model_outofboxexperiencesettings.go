package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfBoxExperienceSettings struct {
	DeviceUsageType           *OutOfBoxExperienceSettingsDeviceUsageType `json:"deviceUsageType,omitempty"`
	HideEULA                  *bool                                      `json:"hideEULA,omitempty"`
	HideEscapeLink            *bool                                      `json:"hideEscapeLink,omitempty"`
	HidePrivacySettings       *bool                                      `json:"hidePrivacySettings,omitempty"`
	ODataType                 *string                                    `json:"@odata.type,omitempty"`
	SkipKeyboardSelectionPage *bool                                      `json:"skipKeyboardSelectionPage,omitempty"`
	UserType                  *OutOfBoxExperienceSettingsUserType        `json:"userType,omitempty"`
}
