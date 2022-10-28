package availabilitysets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkuRestrictionInfo struct {
	Locations *[]string `json:"locations,omitempty"`
	Zones     *Zones    `json:"zones,omitempty"`
}
