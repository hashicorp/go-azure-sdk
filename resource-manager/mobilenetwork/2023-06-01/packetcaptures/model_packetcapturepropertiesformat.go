package packetcaptures

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCapturePropertiesFormat struct {
	BytesToCapturePerPacket *int64               `json:"bytesToCapturePerPacket,omitempty"`
	CaptureStartTime        *string              `json:"captureStartTime,omitempty"`
	NetworkInterfaces       *[]string            `json:"networkInterfaces,omitempty"`
	ProvisioningState       *ProvisioningState   `json:"provisioningState,omitempty"`
	Reason                  *string              `json:"reason,omitempty"`
	Status                  *PacketCaptureStatus `json:"status,omitempty"`
	TimeLimitInSeconds      *int64               `json:"timeLimitInSeconds,omitempty"`
	TotalBytesPerSession    *int64               `json:"totalBytesPerSession,omitempty"`
}
