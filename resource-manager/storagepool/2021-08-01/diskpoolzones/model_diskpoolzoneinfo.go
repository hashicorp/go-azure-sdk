package diskpoolzones

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskPoolZoneInfo struct {
	AdditionalCapabilities *[]string     `json:"additionalCapabilities,omitempty"`
	AvailabilityZones      *zones.Schema `json:"availabilityZones,omitempty"`
	Sku                    *Sku          `json:"sku,omitempty"`
}
