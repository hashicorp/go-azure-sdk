package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettings struct {
	DeviceComplianceCheckinThresholdDays *int64  `json:"deviceComplianceCheckinThresholdDays,omitempty"`
	IsScheduledActionEnabled             *bool   `json:"isScheduledActionEnabled,omitempty"`
	ODataType                            *string `json:"@odata.type,omitempty"`
	SecureByDefault                      *bool   `json:"secureByDefault,omitempty"`
}
