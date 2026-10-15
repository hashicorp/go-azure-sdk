package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropagationPolicy struct {
	PlacementProfile *PlacementProfile `json:"placementProfile,omitempty"`
	Type             PropagationType   `json:"type"`
}
