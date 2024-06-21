package packetcaptures

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCaptureQueryStatusResult struct {
	CaptureStartTime    *string    `json:"captureStartTime,omitempty"`
	Id                  *string    `json:"id,omitempty"`
	Name                *string    `json:"name,omitempty"`
	PacketCaptureError  *[]PcError `json:"packetCaptureError,omitempty"`
	PacketCaptureStatus *PcStatus  `json:"packetCaptureStatus,omitempty"`
	StopReason          *string    `json:"stopReason,omitempty"`
}
