package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabServicesSku struct {
	Capabilities *[]LabServicesSkuCapabilities `json:"capabilities,omitempty"`
	Capacity     *LabServicesSkuCapacity       `json:"capacity,omitempty"`
	Costs        *[]LabServicesSkuCost         `json:"costs,omitempty"`
	Family       *string                       `json:"family,omitempty"`
	Locations    *[]string                     `json:"locations,omitempty"`
	Name         *string                       `json:"name,omitempty"`
	ResourceType *string                       `json:"resourceType,omitempty"`
	Restrictions *[]LabServicesSkuRestrictions `json:"restrictions,omitempty"`
	Size         *string                       `json:"size,omitempty"`
	Tier         *LabServicesSkuTier           `json:"tier,omitempty"`
}
