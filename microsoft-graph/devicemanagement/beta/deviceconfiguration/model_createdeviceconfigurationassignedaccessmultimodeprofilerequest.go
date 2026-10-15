package deviceconfiguration

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright IBM Corp. 2022, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeviceConfigurationAssignedAccessMultiModeProfileRequest struct {
	AssignedAccessMultiModeProfiles *[]beta.WindowsAssignedAccessProfile `json:"assignedAccessMultiModeProfiles,omitempty"`
}
