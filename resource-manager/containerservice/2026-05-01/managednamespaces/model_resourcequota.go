package managednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceQuota struct {
	CpuLimit      *string `json:"cpuLimit,omitempty"`
	CpuRequest    *string `json:"cpuRequest,omitempty"`
	MemoryLimit   *string `json:"memoryLimit,omitempty"`
	MemoryRequest *string `json:"memoryRequest,omitempty"`
}
