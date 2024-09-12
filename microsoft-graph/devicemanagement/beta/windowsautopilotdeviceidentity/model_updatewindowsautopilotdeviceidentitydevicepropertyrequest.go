package windowsautopilotdeviceidentity

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsAutopilotDeviceIdentityDevicePropertyRequest struct {
	AddressableUserName   nullable.Type[string] `json:"addressableUserName,omitempty"`
	DeviceAccountPassword nullable.Type[string] `json:"deviceAccountPassword,omitempty"`
	DeviceAccountUpn      nullable.Type[string] `json:"deviceAccountUpn,omitempty"`
	DeviceFriendlyName    nullable.Type[string] `json:"deviceFriendlyName,omitempty"`
	DisplayName           nullable.Type[string] `json:"displayName,omitempty"`
	GroupTag              nullable.Type[string] `json:"groupTag,omitempty"`
	UserPrincipalName     nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
