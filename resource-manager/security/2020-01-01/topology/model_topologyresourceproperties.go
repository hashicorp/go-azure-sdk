package topology

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopologyResourceProperties struct {
	CalculatedDateTime *string                   `json:"calculatedDateTime,omitempty"`
	TopologyResources  *[]TopologySingleResource `json:"topologyResources,omitempty"`
}
