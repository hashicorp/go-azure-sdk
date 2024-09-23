package windowsautopilotdeviceidentity

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignWindowsAutopilotDeviceIdentityUserToDeviceRequest struct {
	AddressableUserName nullable.Type[string] `json:"addressableUserName,omitempty"`
	UserPrincipalName   nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
