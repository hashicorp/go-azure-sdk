package comanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PauseDeviceManagementComanagedDeviceConfigurationRefreshRequest struct {
	PauseTimePeriodInMinutes *int64 `json:"pauseTimePeriodInMinutes,omitempty"`
}
