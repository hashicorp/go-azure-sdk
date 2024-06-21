package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProperties struct {
	CapacityAllocated *int64                    `json:"capacityAllocated,omitempty"`
	CapacityAssigned  *int64                    `json:"capacityAssigned,omitempty"`
	ClusterId         *string                   `json:"clusterId,omitempty"`
	CreatedDate       *string                   `json:"createdDate,omitempty"`
	ProvisioningState *ClusterProvisioningState `json:"provisioningState,omitempty"`
}
