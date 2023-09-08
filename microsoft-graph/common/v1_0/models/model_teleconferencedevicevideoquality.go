package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeleconferenceDeviceVideoQuality struct {
	AverageInboundJitter          *string `json:"averageInboundJitter,omitempty"`
	AverageInboundRoundTripDelay  *string `json:"averageInboundRoundTripDelay,omitempty"`
	AverageOutboundJitter         *string `json:"averageOutboundJitter,omitempty"`
	AverageOutboundRoundTripDelay *string `json:"averageOutboundRoundTripDelay,omitempty"`
	ChannelIndex                  *int64  `json:"channelIndex,omitempty"`
	InboundPackets                *int64  `json:"inboundPackets,omitempty"`
	LocalIPAddress                *string `json:"localIPAddress,omitempty"`
	LocalPort                     *int64  `json:"localPort,omitempty"`
	MaximumInboundJitter          *string `json:"maximumInboundJitter,omitempty"`
	MaximumInboundRoundTripDelay  *string `json:"maximumInboundRoundTripDelay,omitempty"`
	MaximumOutboundJitter         *string `json:"maximumOutboundJitter,omitempty"`
	MaximumOutboundRoundTripDelay *string `json:"maximumOutboundRoundTripDelay,omitempty"`
	MediaDuration                 *string `json:"mediaDuration,omitempty"`
	NetworkLinkSpeedInBytes       *int64  `json:"networkLinkSpeedInBytes,omitempty"`
	ODataType                     *string `json:"@odata.type,omitempty"`
	OutboundPackets               *int64  `json:"outboundPackets,omitempty"`
	RemoteIPAddress               *string `json:"remoteIPAddress,omitempty"`
	RemotePort                    *int64  `json:"remotePort,omitempty"`
}
