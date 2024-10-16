package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkAttachment struct {
	AttachedNetworkId     string                           `json:"attachedNetworkId"`
	DefaultGateway        *DefaultGateway                  `json:"defaultGateway,omitempty"`
	IPAllocationMethod    VirtualMachineIPAllocationMethod `json:"ipAllocationMethod"`
	IPv4Address           *string                          `json:"ipv4Address,omitempty"`
	IPv6Address           *string                          `json:"ipv6Address,omitempty"`
	MacAddress            *string                          `json:"macAddress,omitempty"`
	NetworkAttachmentName *string                          `json:"networkAttachmentName,omitempty"`
}
