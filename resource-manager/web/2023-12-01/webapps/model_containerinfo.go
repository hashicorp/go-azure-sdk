package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerInfo struct {
	CurrentCPUStats   *ContainerCPUStatistics              `json:"currentCpuStats,omitempty"`
	CurrentTimeStamp  *string                              `json:"currentTimeStamp,omitempty"`
	Eth0              *ContainerNetworkInterfaceStatistics `json:"eth0,omitempty"`
	Id                *string                              `json:"id,omitempty"`
	MemoryStats       *ContainerMemoryStatistics           `json:"memoryStats,omitempty"`
	Name              *string                              `json:"name,omitempty"`
	PreviousCPUStats  *ContainerCPUStatistics              `json:"previousCpuStats,omitempty"`
	PreviousTimeStamp *string                              `json:"previousTimeStamp,omitempty"`
}
