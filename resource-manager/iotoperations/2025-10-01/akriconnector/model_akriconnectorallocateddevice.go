package akriconnector

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorAllocatedDevice struct {
	DeviceInboundEndpointName string `json:"deviceInboundEndpointName"`
	DeviceName                string `json:"deviceName"`
}
