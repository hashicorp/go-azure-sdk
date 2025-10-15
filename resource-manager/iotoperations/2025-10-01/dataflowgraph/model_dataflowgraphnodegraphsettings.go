package dataflowgraph

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphNodeGraphSettings struct {
	Artifact            string                                 `json:"artifact"`
	Configuration       *[]DataflowGraphGraphNodeConfiguration `json:"configuration,omitempty"`
	RegistryEndpointRef string                                 `json:"registryEndpointRef"`
}
