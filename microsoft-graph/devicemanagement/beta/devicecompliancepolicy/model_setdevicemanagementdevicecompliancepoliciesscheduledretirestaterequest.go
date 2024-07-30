package devicecompliancepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequest struct {
	ManagedDeviceIds   *[]string                                                                    `json:"managedDeviceIds,omitempty"`
	ScopedToAllDevices *bool                                                                        `json:"scopedToAllDevices,omitempty"`
	State              *SetDeviceManagementDeviceCompliancePoliciesScheduledRetireStateRequestState `json:"state,omitempty"`
}
