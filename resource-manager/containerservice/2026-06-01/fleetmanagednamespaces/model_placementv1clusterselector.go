package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementV1ClusterSelector struct {
	ClusterSelectorTerms []PlacementV1ClusterSelectorTerm `json:"clusterSelectorTerms"`
}
