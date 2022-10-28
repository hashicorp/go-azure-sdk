package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiManagementSkuRestrictionInfo struct {
	Locations *[]string `json:"locations,omitempty"`
	Zones     *Zones    `json:"zones,omitempty"`
}
