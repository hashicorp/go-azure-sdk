package addons

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTAddonProperties struct {
	HostPlatform         *PlatformType     `json:"hostPlatform,omitempty"`
	HostPlatformType     *HostPlatformType `json:"hostPlatformType,omitempty"`
	IoTDeviceDetails     IoTDeviceInfo     `json:"ioTDeviceDetails"`
	IoTEdgeDeviceDetails IoTDeviceInfo     `json:"ioTEdgeDeviceDetails"`
	ProvisioningState    *AddonState       `json:"provisioningState,omitempty"`
	Version              *string           `json:"version,omitempty"`
}
