package virtualmachinescalesetvms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVMInstanceView struct {
	AssignedHost              *string                                `json:"assignedHost,omitempty"`
	BootDiagnostics           *BootDiagnosticsInstanceView           `json:"bootDiagnostics"`
	Disks                     *[]DiskInstanceView                    `json:"disks,omitempty"`
	Extensions                *[]VirtualMachineExtensionInstanceView `json:"extensions,omitempty"`
	MaintenanceRedeployStatus *MaintenanceRedeployStatus             `json:"maintenanceRedeployStatus"`
	PlacementGroupId          *string                                `json:"placementGroupId,omitempty"`
	PlatformFaultDomain       *int64                                 `json:"platformFaultDomain,omitempty"`
	PlatformUpdateDomain      *int64                                 `json:"platformUpdateDomain,omitempty"`
	RdpThumbPrint             *string                                `json:"rdpThumbPrint,omitempty"`
	Statuses                  *[]InstanceViewStatus                  `json:"statuses,omitempty"`
	VmAgent                   *VirtualMachineAgentInstanceView       `json:"vmAgent"`
	VmHealth                  *VirtualMachineHealthStatus            `json:"vmHealth"`
}
