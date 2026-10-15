package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementProfile struct {
	DefaultClusterResourcePlacement *PlacementV1ClusterResourcePlacementSpec `json:"defaultClusterResourcePlacement,omitempty"`
}
