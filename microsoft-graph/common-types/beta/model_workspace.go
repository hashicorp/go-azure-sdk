package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Workspace struct {
	Address                *PhysicalAddress       `json:"address,omitempty"`
	Building               *string                `json:"building,omitempty"`
	Capacity               *int64                 `json:"capacity,omitempty"`
	DisplayName            *string                `json:"displayName,omitempty"`
	EmailAddress           *string                `json:"emailAddress,omitempty"`
	FloorLabel             *string                `json:"floorLabel,omitempty"`
	FloorNumber            *int64                 `json:"floorNumber,omitempty"`
	GeoCoordinates         *OutlookGeoCoordinates `json:"geoCoordinates,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	IsWheelChairAccessible *bool                  `json:"isWheelChairAccessible,omitempty"`
	Label                  *string                `json:"label,omitempty"`
	Nickname               *string                `json:"nickname,omitempty"`
	ODataType              *string                `json:"@odata.type,omitempty"`
	Phone                  *string                `json:"phone,omitempty"`
	Tags                   *[]string              `json:"tags,omitempty"`
}
