package windowsautopilotdeploymentprofileassigneddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceManagementWindowsAutopilotDeploymentProfileAssignedDeviceDevicePropertyRequest struct {
	AddressableUserName   *string `json:"addressableUserName,omitempty"`
	DeviceAccountPassword *string `json:"deviceAccountPassword,omitempty"`
	DeviceAccountUpn      *string `json:"deviceAccountUpn,omitempty"`
	DeviceFriendlyName    *string `json:"deviceFriendlyName,omitempty"`
	DisplayName           *string `json:"displayName,omitempty"`
	GroupTag              *string `json:"groupTag,omitempty"`
	UserPrincipalName     *string `json:"userPrincipalName,omitempty"`
}
