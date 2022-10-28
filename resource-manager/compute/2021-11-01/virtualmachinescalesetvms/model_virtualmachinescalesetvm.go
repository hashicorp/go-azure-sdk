package virtualmachinescalesetvms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVM struct {
	Id         *string                             `json:"id,omitempty"`
	InstanceId *string                             `json:"instanceId,omitempty"`
	Location   string                              `json:"location"`
	Name       *string                             `json:"name,omitempty"`
	Plan       *Plan                               `json:"plan,omitempty"`
	Properties *VirtualMachineScaleSetVMProperties `json:"properties,omitempty"`
	Resources  *[]VirtualMachineExtension          `json:"resources,omitempty"`
	Sku        *Sku                                `json:"sku,omitempty"`
	Tags       *map[string]string                  `json:"tags,omitempty"`
	Type       *string                             `json:"type,omitempty"`
	Zones      *Zones                              `json:"zones,omitempty"`
}
