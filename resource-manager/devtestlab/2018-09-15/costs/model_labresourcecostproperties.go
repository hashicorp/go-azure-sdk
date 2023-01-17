package costs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabResourceCostProperties struct {
	ExternalResourceId  *string  `json:"externalResourceId,omitempty"`
	ResourceCost        *float64 `json:"resourceCost,omitempty"`
	ResourceId          *string  `json:"resourceId,omitempty"`
	ResourceOwner       *string  `json:"resourceOwner,omitempty"`
	ResourcePricingTier *string  `json:"resourcePricingTier,omitempty"`
	ResourceStatus      *string  `json:"resourceStatus,omitempty"`
	ResourceType        *string  `json:"resourceType,omitempty"`
	ResourceUId         *string  `json:"resourceUId,omitempty"`
	Resourcename        *string  `json:"resourcename,omitempty"`
}
