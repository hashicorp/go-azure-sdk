package privateclouds

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateCloud struct {
	Id         *string                  `json:"id,omitempty"`
	Identity   *identity.SystemAssigned `json:"identity,omitempty"`
	Location   string                   `json:"location"`
	Name       *string                  `json:"name,omitempty"`
	Properties *PrivateCloudProperties  `json:"properties,omitempty"`
	Sku        Sku                      `json:"sku"`
	SystemData *systemdata.SystemData   `json:"systemData,omitempty"`
	Tags       *map[string]string       `json:"tags,omitempty"`
	Type       *string                  `json:"type,omitempty"`
	Zones      *zones.Schema            `json:"zones,omitempty"`
}
