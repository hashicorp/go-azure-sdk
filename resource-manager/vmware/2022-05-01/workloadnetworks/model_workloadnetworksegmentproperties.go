package workloadnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkSegmentProperties struct {
	ConnectedGateway  *string                                  `json:"connectedGateway,omitempty"`
	DisplayName       *string                                  `json:"displayName,omitempty"`
	PortVif           *[]WorkloadNetworkSegmentPortVif         `json:"portVif,omitempty"`
	ProvisioningState *WorkloadNetworkSegmentProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                   `json:"revision,omitempty"`
	Status            *SegmentStatusEnum                       `json:"status,omitempty"`
	Subnet            *WorkloadNetworkSegmentSubnet            `json:"subnet"`
}
