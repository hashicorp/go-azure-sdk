package vmwares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSku struct {
	Capabilities *[]ResourceSkuCapabilities `json:"capabilities,omitempty"`
	Family       *string                    `json:"family,omitempty"`
	LocationInfo []ResourceSkuLocationInfo  `json:"locationInfo"`
	Locations    []string                   `json:"locations"`
	Name         string                     `json:"name"`
	ResourceType ResourceSkuResourceType    `json:"resourceType"`
	Restrictions []ResourceSkuRestrictions  `json:"restrictions"`
	Size         *string                    `json:"size,omitempty"`
	Tier         *string                    `json:"tier,omitempty"`
}
