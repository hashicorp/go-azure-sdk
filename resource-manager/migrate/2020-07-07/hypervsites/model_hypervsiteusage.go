package hypervsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVSiteUsage struct {
	ClusterCount      *int64 `json:"clusterCount,omitempty"`
	HostCount         *int64 `json:"hostCount,omitempty"`
	MachineCount      *int64 `json:"machineCount,omitempty"`
	RunAsAccountCount *int64 `json:"runAsAccountCount,omitempty"`
}
