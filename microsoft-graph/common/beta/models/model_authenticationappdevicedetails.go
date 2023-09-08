package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppDeviceDetails struct {
	AppVersion      *string `json:"appVersion,omitempty"`
	ClientApp       *string `json:"clientApp,omitempty"`
	DeviceId        *string `json:"deviceId,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	OperatingSystem *string `json:"operatingSystem,omitempty"`
}
