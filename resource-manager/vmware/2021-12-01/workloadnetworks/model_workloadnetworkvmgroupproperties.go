package workloadnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkVMGroupProperties struct {
	DisplayName       *string                                  `json:"displayName,omitempty"`
	Members           *[]string                                `json:"members,omitempty"`
	ProvisioningState *WorkloadNetworkVMGroupProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                   `json:"revision,omitempty"`
	Status            *VMGroupStatusEnum                       `json:"status,omitempty"`
}
