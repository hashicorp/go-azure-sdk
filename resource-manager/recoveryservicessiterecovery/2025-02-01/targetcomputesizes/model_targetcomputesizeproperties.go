package targetcomputesizes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetComputeSizeProperties struct {
	CpuCoresCount     *int64                     `json:"cpuCoresCount,omitempty"`
	Errors            *[]ComputeSizeErrorDetails `json:"errors,omitempty"`
	FriendlyName      *string                    `json:"friendlyName,omitempty"`
	HighIopsSupported *string                    `json:"highIopsSupported,omitempty"`
	HyperVGenerations *[]string                  `json:"hyperVGenerations,omitempty"`
	MaxDataDiskCount  *int64                     `json:"maxDataDiskCount,omitempty"`
	MaxNicsCount      *int64                     `json:"maxNicsCount,omitempty"`
	MemoryInGB        *float64                   `json:"memoryInGB,omitempty"`
	Name              *string                    `json:"name,omitempty"`
	VCPUsAvailable    *int64                     `json:"vCPUsAvailable,omitempty"`
}
