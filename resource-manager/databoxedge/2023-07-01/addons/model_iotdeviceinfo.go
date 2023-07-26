package addons

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTDeviceInfo struct {
	Authentication *Authentication `json:"authentication,omitempty"`
	DeviceId       string          `json:"deviceId"`
	IoTHostHub     string          `json:"ioTHostHub"`
	IoTHostHubId   *string         `json:"ioTHostHubId,omitempty"`
}
