package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementUserRightsSetting struct {
	LocalUsersOrGroups *[]DeviceManagementUserRightsLocalUserOrGroup `json:"localUsersOrGroups,omitempty"`
	ODataType          *string                                       `json:"@odata.type,omitempty"`
	State              *DeviceManagementUserRightsSettingState       `json:"state,omitempty"`
}
