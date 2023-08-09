package workloadnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkPortMirroringProperties struct {
	Destination       *string                                        `json:"destination,omitempty"`
	Direction         *PortMirroringDirectionEnum                    `json:"direction,omitempty"`
	DisplayName       *string                                        `json:"displayName,omitempty"`
	ProvisioningState *WorkloadNetworkPortMirroringProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                         `json:"revision,omitempty"`
	Source            *string                                        `json:"source,omitempty"`
	Status            *PortMirroringStatusEnum                       `json:"status,omitempty"`
}
