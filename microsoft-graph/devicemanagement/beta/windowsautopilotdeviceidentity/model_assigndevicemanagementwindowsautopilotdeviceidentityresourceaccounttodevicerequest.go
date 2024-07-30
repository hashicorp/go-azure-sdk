package windowsautopilotdeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementWindowsAutopilotDeviceIdentityResourceAccountToDeviceRequest struct {
	AddressableUserName *string `json:"addressableUserName,omitempty"`
	ResourceAccountName *string `json:"resourceAccountName,omitempty"`
	UserPrincipalName   *string `json:"userPrincipalName,omitempty"`
}
