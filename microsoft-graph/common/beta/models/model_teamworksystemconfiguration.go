package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSystemConfiguration struct {
	DateTimeConfiguration  *TeamworkDateTimeConfiguration `json:"dateTimeConfiguration,omitempty"`
	DefaultPassword        *string                        `json:"defaultPassword,omitempty"`
	DeviceLockTimeout      *string                        `json:"deviceLockTimeout,omitempty"`
	IsDeviceLockEnabled    *bool                          `json:"isDeviceLockEnabled,omitempty"`
	IsLoggingEnabled       *bool                          `json:"isLoggingEnabled,omitempty"`
	IsPowerSavingEnabled   *bool                          `json:"isPowerSavingEnabled,omitempty"`
	IsScreenCaptureEnabled *bool                          `json:"isScreenCaptureEnabled,omitempty"`
	IsSilentModeEnabled    *bool                          `json:"isSilentModeEnabled,omitempty"`
	Language               *string                        `json:"language,omitempty"`
	LockPin                *string                        `json:"lockPin,omitempty"`
	LoggingLevel           *string                        `json:"loggingLevel,omitempty"`
	NetworkConfiguration   *TeamworkNetworkConfiguration  `json:"networkConfiguration,omitempty"`
	ODataType              *string                        `json:"@odata.type,omitempty"`
}
