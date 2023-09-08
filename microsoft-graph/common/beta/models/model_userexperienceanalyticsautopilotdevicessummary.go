package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAutopilotDevicesSummary struct {
	DevicesNotAutopilotRegistered              *int64  `json:"devicesNotAutopilotRegistered,omitempty"`
	DevicesWithoutAutopilotProfileAssigned     *int64  `json:"devicesWithoutAutopilotProfileAssigned,omitempty"`
	ODataType                                  *string `json:"@odata.type,omitempty"`
	TotalWindows10DevicesWithoutTenantAttached *int64  `json:"totalWindows10DevicesWithoutTenantAttached,omitempty"`
}
