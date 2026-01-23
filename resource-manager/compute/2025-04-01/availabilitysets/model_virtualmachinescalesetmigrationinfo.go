package availabilitysets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetMigrationInfo struct {
	DefaultVirtualMachineScaleSetInfo *DefaultVirtualMachineScaleSetInfo `json:"defaultVirtualMachineScaleSetInfo,omitempty"`
	MigrateToVirtualMachineScaleSet   *SubResource                       `json:"migrateToVirtualMachineScaleSet,omitempty"`
}
