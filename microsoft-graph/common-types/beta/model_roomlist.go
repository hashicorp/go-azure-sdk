package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoomList struct {
	Address        *PhysicalAddress       `json:"address,omitempty"`
	DisplayName    *string                `json:"displayName,omitempty"`
	EmailAddress   *string                `json:"emailAddress,omitempty"`
	GeoCoordinates *OutlookGeoCoordinates `json:"geoCoordinates,omitempty"`
	Id             *string                `json:"id,omitempty"`
	ODataType      *string                `json:"@odata.type,omitempty"`
	Phone          *string                `json:"phone,omitempty"`
	Rooms          *[]Room                `json:"rooms,omitempty"`
	Workspaces     *[]Workspace           `json:"workspaces,omitempty"`
}
