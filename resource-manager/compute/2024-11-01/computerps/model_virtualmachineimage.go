package computerps

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/edgezones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineImage struct {
	ExtendedLocation *edgezones.Model               `json:"extendedLocation,omitempty"`
	Id               *string                        `json:"id,omitempty"`
	Location         string                         `json:"location"`
	Name             string                         `json:"name"`
	Properties       *VirtualMachineImageProperties `json:"properties,omitempty"`
	Tags             *map[string]string             `json:"tags,omitempty"`
}
