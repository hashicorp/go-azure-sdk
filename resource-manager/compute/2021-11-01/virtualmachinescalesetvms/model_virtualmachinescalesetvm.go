package virtualmachinescalesetvms

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVM struct {
	Id         *string                             `json:"id,omitempty"`
	InstanceId *string                             `json:"instanceId,omitempty"`
	Location   string                              `json:"location"`
	Name       *string                             `json:"name,omitempty"`
	Plan       *Plan                               `json:"plan"`
	Properties *VirtualMachineScaleSetVMProperties `json:"properties"`
	Resources  *[]VirtualMachineExtension          `json:"resources,omitempty"`
	Sku        *Sku                                `json:"sku"`
	Tags       *map[string]string                  `json:"tags,omitempty"`
	Type       *string                             `json:"type,omitempty"`
	Zones      *zones.Schema                       `json:"zones,omitempty"`
}
