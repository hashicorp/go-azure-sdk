package projectcapabilityhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectCapabilityHost struct {
	AiServicesConnections    *[]string                        `json:"aiServicesConnections,omitempty"`
	ProvisioningState        *CapabilityHostProvisioningState `json:"provisioningState,omitempty"`
	StorageConnections       *[]string                        `json:"storageConnections,omitempty"`
	ThreadStorageConnections *[]string                        `json:"threadStorageConnections,omitempty"`
	VectorStoreConnections   *[]string                        `json:"vectorStoreConnections,omitempty"`
}
