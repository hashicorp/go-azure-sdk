package topology

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopologySingleResource struct {
	Children             *[]TopologySingleResourceChild  `json:"children,omitempty"`
	Location             *string                         `json:"location,omitempty"`
	NetworkZones         *string                         `json:"networkZones,omitempty"`
	Parents              *[]TopologySingleResourceParent `json:"parents,omitempty"`
	RecommendationsExist *bool                           `json:"recommendationsExist,omitempty"`
	ResourceId           *string                         `json:"resourceId,omitempty"`
	Severity             *string                         `json:"severity,omitempty"`
	TopologyScore        *int64                          `json:"topologyScore,omitempty"`
}
