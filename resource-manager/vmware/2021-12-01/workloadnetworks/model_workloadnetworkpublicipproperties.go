package workloadnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkPublicIPProperties struct {
	DisplayName       *string                                   `json:"displayName,omitempty"`
	NumberOfPublicIPs *int64                                    `json:"numberOfPublicIPs,omitempty"`
	ProvisioningState *WorkloadNetworkPublicIPProvisioningState `json:"provisioningState,omitempty"`
	PublicIPBlock     *string                                   `json:"publicIPBlock,omitempty"`
}
