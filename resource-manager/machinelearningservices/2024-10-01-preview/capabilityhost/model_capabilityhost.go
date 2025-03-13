package capabilityhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityHost struct {
	AiServicesConnections  *[]string                        `json:"aiServicesConnections,omitempty"`
	CapabilityHostKind     *CapabilityHostKind              `json:"capabilityHostKind,omitempty"`
	CustomerSubnet         *string                          `json:"customerSubnet,omitempty"`
	Description            *string                          `json:"description,omitempty"`
	Properties             *map[string]string               `json:"properties,omitempty"`
	ProvisioningState      *CapabilityHostProvisioningState `json:"provisioningState,omitempty"`
	StorageConnections     *[]string                        `json:"storageConnections,omitempty"`
	Tags                   *map[string]string               `json:"tags,omitempty"`
	VectorStoreConnections *[]string                        `json:"vectorStoreConnections,omitempty"`
}
