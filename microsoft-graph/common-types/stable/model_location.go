package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Location struct {
	Address              *PhysicalAddress       `json:"address,omitempty"`
	Coordinates          *OutlookGeoCoordinates `json:"coordinates,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	LocationEmailAddress *string                `json:"locationEmailAddress,omitempty"`
	LocationType         *LocationLocationType  `json:"locationType,omitempty"`
	LocationUri          *string                `json:"locationUri,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	UniqueId             *string                `json:"uniqueId,omitempty"`
	UniqueIdType         *LocationUniqueIdType  `json:"uniqueIdType,omitempty"`
}
