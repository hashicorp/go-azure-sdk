package intent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceManagementIntentSettingRequest struct {
	Settings *[]DeviceManagementSettingInstance `json:"settings,omitempty"`
}
