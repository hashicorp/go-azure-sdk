package devicecapacityinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMMemory struct {
	CurrentMemoryUsageMB *int64 `json:"currentMemoryUsageMB,omitempty"`
	StartupMemoryMB      *int64 `json:"startupMemoryMB,omitempty"`
}
