package windowsautopilotdeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceManagementWindowsAutopilotDeviceIdentityDevicePropertyRequest struct {
	AddressableUserName *string `json:"addressableUserName,omitempty"`
	DisplayName         *string `json:"displayName,omitempty"`
	GroupTag            *string `json:"groupTag,omitempty"`
	UserPrincipalName   *string `json:"userPrincipalName,omitempty"`
}
