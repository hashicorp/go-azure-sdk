package virtualmachinescalesets

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetUpdate struct {
	Identity   *identity.SystemAndUserAssignedMap      `json:"identity,omitempty"`
	Plan       *Plan                                   `json:"plan"`
	Properties *VirtualMachineScaleSetUpdateProperties `json:"properties"`
	Sku        *Sku                                    `json:"sku"`
	Tags       *map[string]string                      `json:"tags,omitempty"`
}
