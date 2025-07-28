package computerps

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/edgezones"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachine struct {
	Etag             *string                            `json:"etag,omitempty"`
	ExtendedLocation *edgezones.Model                   `json:"extendedLocation,omitempty"`
	Id               *string                            `json:"id,omitempty"`
	Identity         *identity.SystemAndUserAssignedMap `json:"identity,omitempty"`
	Location         string                             `json:"location"`
	ManagedBy        *string                            `json:"managedBy,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Placement        *Placement                         `json:"placement,omitempty"`
	Plan             *Plan                              `json:"plan,omitempty"`
	Properties       *VirtualMachineProperties          `json:"properties,omitempty"`
	Resources        *[]VirtualMachineExtension         `json:"resources,omitempty"`
	SystemData       *systemdata.SystemData             `json:"systemData,omitempty"`
	Tags             *map[string]string                 `json:"tags,omitempty"`
	Type             *string                            `json:"type,omitempty"`
	Zones            *zones.Schema                      `json:"zones,omitempty"`
}
