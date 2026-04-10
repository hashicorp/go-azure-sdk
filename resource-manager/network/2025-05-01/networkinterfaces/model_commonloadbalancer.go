package networkinterfaces

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/edgezones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonLoadBalancer struct {
	Etag             *string                             `json:"etag,omitempty"`
	ExtendedLocation *edgezones.Model                    `json:"extendedLocation,omitempty"`
	Id               *string                             `json:"id,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             *string                             `json:"name,omitempty"`
	Properties       *CommonLoadBalancerPropertiesFormat `json:"properties,omitempty"`
	Sku              *CommonLoadBalancerSku              `json:"sku,omitempty"`
	Tags             *map[string]string                  `json:"tags,omitempty"`
	Type             *string                             `json:"type,omitempty"`
}
