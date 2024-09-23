package manageddevice

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateManagedDeviceWindowsDeviceAccountRequest struct {
	UpdateWindowsDeviceAccountActionParameter *stable.UpdateWindowsDeviceAccountActionParameter `json:"updateWindowsDeviceAccountActionParameter,omitempty"`
}
