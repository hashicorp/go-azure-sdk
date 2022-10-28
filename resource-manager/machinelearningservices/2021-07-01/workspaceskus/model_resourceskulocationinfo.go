package workspaceskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkuLocationInfo struct {
	Location    *string                   `json:"location,omitempty"`
	ZoneDetails *[]ResourceSkuZoneDetails `json:"zoneDetails,omitempty"`
	Zones       *Zones                    `json:"zones,omitempty"`
}
