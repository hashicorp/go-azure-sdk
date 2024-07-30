package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceGeoLocation struct {
	Altitude              *float64 `json:"altitude,omitempty"`
	Heading               *float64 `json:"heading,omitempty"`
	HorizontalAccuracy    *float64 `json:"horizontalAccuracy,omitempty"`
	LastCollectedDateTime *string  `json:"lastCollectedDateTime,omitempty"`
	Latitude              *float64 `json:"latitude,omitempty"`
	Longitude             *float64 `json:"longitude,omitempty"`
	ODataType             *string  `json:"@odata.type,omitempty"`
	Speed                 *float64 `json:"speed,omitempty"`
	VerticalAccuracy      *float64 `json:"verticalAccuracy,omitempty"`
}
