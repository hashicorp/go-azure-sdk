package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceGeoLocation struct {
	LastCollectedDateTime *string `json:"lastCollectedDateTime,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
