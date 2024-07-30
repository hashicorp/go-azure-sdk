package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptUserState struct {
	DeviceRunStates    *[]DeviceManagementScriptDeviceState `json:"deviceRunStates,omitempty"`
	ErrorDeviceCount   *int64                               `json:"errorDeviceCount,omitempty"`
	Id                 *string                              `json:"id,omitempty"`
	ODataType          *string                              `json:"@odata.type,omitempty"`
	SuccessDeviceCount *int64                               `json:"successDeviceCount,omitempty"`
	UserPrincipalName  *string                              `json:"userPrincipalName,omitempty"`
}
