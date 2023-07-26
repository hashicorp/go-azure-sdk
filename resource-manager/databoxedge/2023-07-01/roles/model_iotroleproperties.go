package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTRoleProperties struct {
	ComputeResource      *ComputeResource  `json:"computeResource,omitempty"`
	HostPlatform         PlatformType      `json:"hostPlatform"`
	HostPlatformType     *HostPlatformType `json:"hostPlatformType,omitempty"`
	IoTDeviceDetails     IoTDeviceInfo     `json:"ioTDeviceDetails"`
	IoTEdgeAgentInfo     *IoTEdgeAgentInfo `json:"ioTEdgeAgentInfo,omitempty"`
	IoTEdgeDeviceDetails IoTDeviceInfo     `json:"ioTEdgeDeviceDetails"`
	RoleStatus           RoleStatus        `json:"roleStatus"`
	ShareMappings        *[]MountPointMap  `json:"shareMappings,omitempty"`
}
