package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Room struct {
	Address                *PhysicalAddress       `json:"address,omitempty"`
	AudioDeviceName        *string                `json:"audioDeviceName,omitempty"`
	BookingType            *RoomBookingType       `json:"bookingType,omitempty"`
	Building               *string                `json:"building,omitempty"`
	Capacity               *int64                 `json:"capacity,omitempty"`
	DisplayDeviceName      *string                `json:"displayDeviceName,omitempty"`
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
	VideoDeviceName        *string                `json:"videoDeviceName,omitempty"`
}
