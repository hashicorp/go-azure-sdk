package resourceskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkuInfo struct {
	ApiVersion   *string                    `json:"apiVersion,omitempty"`
	Capabilities *[]ResourceSkuCapability   `json:"capabilities,omitempty"`
	LocationInfo *ResourceSkuLocationInfo   `json:"locationInfo"`
	Name         *string                    `json:"name,omitempty"`
	ResourceType *string                    `json:"resourceType,omitempty"`
	Restrictions *[]ResourceSkuRestrictions `json:"restrictions,omitempty"`
	Tier         *string                    `json:"tier,omitempty"`
}
