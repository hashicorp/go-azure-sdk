package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareProfile struct {
	CpuHotAddEnabled    *bool  `json:"cpuHotAddEnabled,omitempty"`
	CpuHotRemoveEnabled *bool  `json:"cpuHotRemoveEnabled,omitempty"`
	MemoryHotAddEnabled *bool  `json:"memoryHotAddEnabled,omitempty"`
	MemorySizeMB        *int64 `json:"memorySizeMB,omitempty"`
	NumCPUs             *int64 `json:"numCPUs,omitempty"`
	NumCoresPerSocket   *int64 `json:"numCoresPerSocket,omitempty"`
}
