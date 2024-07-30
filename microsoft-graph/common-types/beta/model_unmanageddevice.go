package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDevice struct {
	DeviceName       *string `json:"deviceName,omitempty"`
	Domain           *string `json:"domain,omitempty"`
	IpAddress        *string `json:"ipAddress,omitempty"`
	LastLoggedOnUser *string `json:"lastLoggedOnUser,omitempty"`
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`
	Location         *string `json:"location,omitempty"`
	MacAddress       *string `json:"macAddress,omitempty"`
	Manufacturer     *string `json:"manufacturer,omitempty"`
	Model            *string `json:"model,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	Os               *string `json:"os,omitempty"`
	OsVersion        *string `json:"osVersion,omitempty"`
}
