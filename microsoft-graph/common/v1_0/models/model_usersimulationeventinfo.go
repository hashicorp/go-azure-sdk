package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSimulationEventInfo struct {
	Browser                 *string `json:"browser,omitempty"`
	EventDateTime           *string `json:"eventDateTime,omitempty"`
	EventName               *string `json:"eventName,omitempty"`
	IpAddress               *string `json:"ipAddress,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	OsPlatformDeviceDetails *string `json:"osPlatformDeviceDetails,omitempty"`
}
