package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OsVersionCount struct {
	DeviceCount        *int64  `json:"deviceCount,omitempty"`
	LastUpdateDateTime *string `json:"lastUpdateDateTime,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	OsVersion          *string `json:"osVersion,omitempty"`
}
