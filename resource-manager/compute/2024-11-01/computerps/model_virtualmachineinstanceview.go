package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineInstanceView struct {
	AssignedHost              *string                                `json:"assignedHost,omitempty"`
	BootDiagnostics           *BootDiagnosticsInstanceView           `json:"bootDiagnostics,omitempty"`
	ComputerName              *string                                `json:"computerName,omitempty"`
	Disks                     *[]DiskInstanceView                    `json:"disks,omitempty"`
	Extensions                *[]VirtualMachineExtensionInstanceView `json:"extensions,omitempty"`
	HyperVGeneration          *HyperVGenerationType                  `json:"hyperVGeneration,omitempty"`
	IsVMInStandbyPool         *bool                                  `json:"isVMInStandbyPool,omitempty"`
	MaintenanceRedeployStatus *MaintenanceRedeployStatus             `json:"maintenanceRedeployStatus,omitempty"`
	OsName                    *string                                `json:"osName,omitempty"`
	OsVersion                 *string                                `json:"osVersion,omitempty"`
	PatchStatus               *VirtualMachinePatchStatus             `json:"patchStatus,omitempty"`
	PlatformFaultDomain       *int64                                 `json:"platformFaultDomain,omitempty"`
	PlatformUpdateDomain      *int64                                 `json:"platformUpdateDomain,omitempty"`
	RdpThumbPrint             *string                                `json:"rdpThumbPrint,omitempty"`
	Statuses                  *[]InstanceViewStatus                  `json:"statuses,omitempty"`
	VMAgent                   *VirtualMachineAgentInstanceView       `json:"vmAgent,omitempty"`
	VMHealth                  *VirtualMachineHealthStatus            `json:"vmHealth,omitempty"`
}
