package bot

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Bot struct {
	Etag       *string            `json:"etag,omitempty"`
	Id         *string            `json:"id,omitempty"`
	Kind       *Kind              `json:"kind,omitempty"`
	Location   *string            `json:"location,omitempty"`
	Name       *string            `json:"name,omitempty"`
	Properties *BotProperties     `json:"properties,omitempty"`
	Sku        *Sku               `json:"sku,omitempty"`
	Tags       *map[string]string `json:"tags,omitempty"`
	Type       *string            `json:"type,omitempty"`
	Zones      *zones.Schema      `json:"zones,omitempty"`
}
