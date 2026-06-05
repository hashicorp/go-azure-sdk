package dscpconfiguration

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/edgezones"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonPrivateLinkService struct {
	Etag             *string                             `json:"etag,omitempty"`
	ExtendedLocation *edgezones.Model                    `json:"extendedLocation,omitempty"`
	Id               *string                             `json:"id,omitempty"`
	Name             *string                             `json:"name,omitempty"`
	Properties       *CommonPrivateLinkServiceProperties `json:"properties,omitempty"`
	SystemData       *systemdata.SystemData              `json:"systemData,omitempty"`
	Type             *string                             `json:"type,omitempty"`
}
