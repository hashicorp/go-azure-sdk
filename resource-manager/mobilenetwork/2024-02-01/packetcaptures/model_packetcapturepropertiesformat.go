package packetcaptures

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCapturePropertiesFormat struct {
	BytesToCapturePerPacket *int64               `json:"bytesToCapturePerPacket,omitempty"`
	CaptureStartTime        *string              `json:"captureStartTime,omitempty"`
	NetworkInterfaces       *[]string            `json:"networkInterfaces,omitempty"`
	OutputFiles             *[]string            `json:"outputFiles,omitempty"`
	ProvisioningState       *ProvisioningState   `json:"provisioningState,omitempty"`
	Reason                  *string              `json:"reason,omitempty"`
	Status                  *PacketCaptureStatus `json:"status,omitempty"`
	TimeLimitInSeconds      *int64               `json:"timeLimitInSeconds,omitempty"`
	TotalBytesPerSession    *int64               `json:"totalBytesPerSession,omitempty"`
}

func (o *PacketCapturePropertiesFormat) GetCaptureStartTimeAsTime() (*time.Time, error) {
	if o.CaptureStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CaptureStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PacketCapturePropertiesFormat) SetCaptureStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CaptureStartTime = &formatted
}
