package virtualnetworks

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/edgezones"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonPublicIPAddress struct {
	Etag             *string                                `json:"etag,omitempty"`
	ExtendedLocation *edgezones.Model                       `json:"extendedLocation,omitempty"`
	Id               *string                                `json:"id,omitempty"`
	Name             *string                                `json:"name,omitempty"`
	Properties       *CommonPublicIPAddressPropertiesFormat `json:"properties,omitempty"`
	Sku              *CommonPublicIPAddressSku              `json:"sku,omitempty"`
	SystemData       *systemdata.SystemData                 `json:"systemData,omitempty"`
	Type             *string                                `json:"type,omitempty"`
	Zones            *zones.Schema                          `json:"zones,omitempty"`
}
