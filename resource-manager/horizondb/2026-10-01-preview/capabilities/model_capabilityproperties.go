package capabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityProperties struct {
	ClusterCapabilities ClusterCapability  `json:"clusterCapabilities"`
	PoolCapabilities    PoolCapability     `json:"poolCapabilities"`
	ProvisioningState   *ProvisioningState `json:"provisioningState,omitempty"`
}
