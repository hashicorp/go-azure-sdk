package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookGeoCoordinates struct {
	Accuracy         *float64 `json:"accuracy,omitempty"`
	Altitude         *float64 `json:"altitude,omitempty"`
	AltitudeAccuracy *float64 `json:"altitudeAccuracy,omitempty"`
	Latitude         *float64 `json:"latitude,omitempty"`
	Longitude        *float64 `json:"longitude,omitempty"`
	ODataType        *string  `json:"@odata.type,omitempty"`
}
