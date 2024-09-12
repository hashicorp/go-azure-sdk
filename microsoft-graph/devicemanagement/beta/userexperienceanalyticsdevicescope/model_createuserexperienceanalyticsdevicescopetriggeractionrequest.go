package userexperienceanalyticsdevicescope

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsDeviceScopeTriggerActionRequest struct {
	// Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
	ActionName *beta.DeviceScopeAction `json:"actionName,omitempty"`

	DeviceScopeId *string `json:"deviceScopeId,omitempty"`
}
