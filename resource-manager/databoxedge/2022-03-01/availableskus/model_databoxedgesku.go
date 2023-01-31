package availableskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataBoxEdgeSku struct {
	ApiVersions   *[]string          `json:"apiVersions,omitempty"`
	Availability  *SkuAvailability   `json:"availability,omitempty"`
	Capabilities  *[]SkuCapability   `json:"capabilities,omitempty"`
	Costs         *[]SkuCost         `json:"costs,omitempty"`
	Family        *string            `json:"family,omitempty"`
	Kind          *string            `json:"kind,omitempty"`
	LocationInfo  *[]SkuLocationInfo `json:"locationInfo,omitempty"`
	Locations     *[]string          `json:"locations,omitempty"`
	Name          *SkuName           `json:"name,omitempty"`
	ResourceType  *string            `json:"resourceType,omitempty"`
	ShipmentTypes *[]ShipmentType    `json:"shipmentTypes,omitempty"`
	SignupOption  *SkuSignupOption   `json:"signupOption,omitempty"`
	Size          *string            `json:"size,omitempty"`
	Tier          *SkuTier           `json:"tier,omitempty"`
	Version       *SkuVersion        `json:"version,omitempty"`
}
