package workloadnetworkvirtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkVirtualMachineProperties struct {
	DisplayName       *string                           `json:"displayName,omitempty"`
	ProvisioningState *WorkloadNetworkProvisioningState `json:"provisioningState,omitempty"`
	VMType            *VMTypeEnum                       `json:"vmType,omitempty"`
}
