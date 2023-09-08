package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealth struct {
	LastConnectionTime *string `json:"lastConnectionTime,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
}
