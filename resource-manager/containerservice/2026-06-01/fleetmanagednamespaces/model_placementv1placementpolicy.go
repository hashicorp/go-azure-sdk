package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementV1PlacementPolicy struct {
	Affinity      *PlacementV1Affinity     `json:"affinity,omitempty"`
	ClusterNames  *[]string                `json:"clusterNames,omitempty"`
	PlacementType *PlacementType           `json:"placementType,omitempty"`
	Tolerations   *[]PlacementV1Toleration `json:"tolerations,omitempty"`
}
