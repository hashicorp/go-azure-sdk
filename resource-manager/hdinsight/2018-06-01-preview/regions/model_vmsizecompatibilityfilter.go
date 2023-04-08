package regions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMSizeCompatibilityFilter struct {
	ClusterFlavors            *[]string `json:"ClusterFlavors,omitempty"`
	ClusterVersions           *[]string `json:"ClusterVersions,omitempty"`
	ComputeIsolationSupported *string   `json:"ComputeIsolationSupported,omitempty"`
	ESPApplied                *string   `json:"ESPApplied,omitempty"`
	FilterMode                *string   `json:"FilterMode,omitempty"`
	NodeTypes                 *[]string `json:"NodeTypes,omitempty"`
	OsType                    *[]string `json:"OsType,omitempty"`
	Regions                   *[]string `json:"Regions,omitempty"`
	VMSizes                   *[]string `json:"VMSizes,omitempty"`
}
