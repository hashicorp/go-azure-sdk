package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpApplicationSegment struct {
	DestinationHost *string   `json:"destinationHost,omitempty"`
	Id              *string   `json:"id,omitempty"`
	ODataType       *string   `json:"@odata.type,omitempty"`
	Port            *int64    `json:"port,omitempty"`
	Ports           *[]string `json:"ports,omitempty"`
}
