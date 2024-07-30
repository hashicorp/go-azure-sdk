package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationConstraintItem struct {
	Address              *PhysicalAddress                    `json:"address,omitempty"`
	Coordinates          *OutlookGeoCoordinates              `json:"coordinates,omitempty"`
	DisplayName          *string                             `json:"displayName,omitempty"`
	LocationEmailAddress *string                             `json:"locationEmailAddress,omitempty"`
	LocationType         *LocationConstraintItemLocationType `json:"locationType,omitempty"`
	LocationUri          *string                             `json:"locationUri,omitempty"`
	ODataType            *string                             `json:"@odata.type,omitempty"`
	ResolveAvailability  *bool                               `json:"resolveAvailability,omitempty"`
	UniqueId             *string                             `json:"uniqueId,omitempty"`
	UniqueIdType         *LocationConstraintItemUniqueIdType `json:"uniqueIdType,omitempty"`
}
