package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterLocation struct {
	AltitudeInMeters *int64    `json:"altitudeInMeters,omitempty"`
	Building         *string   `json:"building,omitempty"`
	City             *string   `json:"city,omitempty"`
	CountryOrRegion  *string   `json:"countryOrRegion,omitempty"`
	Floor            *string   `json:"floor,omitempty"`
	FloorDescription *string   `json:"floorDescription,omitempty"`
	Latitude         *float64  `json:"latitude,omitempty"`
	Longitude        *float64  `json:"longitude,omitempty"`
	ODataType        *string   `json:"@odata.type,omitempty"`
	Organization     *[]string `json:"organization,omitempty"`
	PostalCode       *string   `json:"postalCode,omitempty"`
	RoomDescription  *string   `json:"roomDescription,omitempty"`
	RoomName         *string   `json:"roomName,omitempty"`
	Site             *string   `json:"site,omitempty"`
	StateOrProvince  *string   `json:"stateOrProvince,omitempty"`
	StreetAddress    *string   `json:"streetAddress,omitempty"`
	Subdivision      *[]string `json:"subdivision,omitempty"`
	Subunit          *[]string `json:"subunit,omitempty"`
}
