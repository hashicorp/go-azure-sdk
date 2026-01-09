package windowsautopilotdeviceidentity

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignWindowsAutopilotDeviceIdentityResourceAccountToDeviceRequest struct {
	AddressableUserName nullable.Type[string] `json:"addressableUserName,omitempty"`
	ResourceAccountName nullable.Type[string] `json:"resourceAccountName,omitempty"`
	UserPrincipalName   nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
