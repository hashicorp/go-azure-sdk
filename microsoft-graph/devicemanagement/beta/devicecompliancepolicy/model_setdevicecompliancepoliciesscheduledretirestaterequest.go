package devicecompliancepolicy

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetDeviceCompliancePoliciesScheduledRetireStateRequest struct {
	ManagedDeviceIds   *[]string           `json:"managedDeviceIds,omitempty"`
	ScopedToAllDevices nullable.Type[bool] `json:"scopedToAllDevices,omitempty"`

	// Cancel or confirm scheduled retire
	State *beta.ScheduledRetireState `json:"state,omitempty"`
}
