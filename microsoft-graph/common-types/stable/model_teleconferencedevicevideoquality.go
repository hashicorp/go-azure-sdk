package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeleconferenceDeviceVideoQuality struct {
	AverageInboundBitRate                     *float64 `json:"averageInboundBitRate,omitempty"`
	AverageInboundFrameRate                   *float64 `json:"averageInboundFrameRate,omitempty"`
	AverageInboundJitter                      *string  `json:"averageInboundJitter,omitempty"`
	AverageInboundPacketLossRateInPercentage  *float64 `json:"averageInboundPacketLossRateInPercentage,omitempty"`
	AverageInboundRoundTripDelay              *string  `json:"averageInboundRoundTripDelay,omitempty"`
	AverageOutboundBitRate                    *float64 `json:"averageOutboundBitRate,omitempty"`
	AverageOutboundFrameRate                  *float64 `json:"averageOutboundFrameRate,omitempty"`
	AverageOutboundJitter                     *string  `json:"averageOutboundJitter,omitempty"`
	AverageOutboundPacketLossRateInPercentage *float64 `json:"averageOutboundPacketLossRateInPercentage,omitempty"`
	AverageOutboundRoundTripDelay             *string  `json:"averageOutboundRoundTripDelay,omitempty"`
	ChannelIndex                              *int64   `json:"channelIndex,omitempty"`
	InboundPackets                            *int64   `json:"inboundPackets,omitempty"`
	LocalIPAddress                            *string  `json:"localIPAddress,omitempty"`
	LocalPort                                 *int64   `json:"localPort,omitempty"`
	MaximumInboundJitter                      *string  `json:"maximumInboundJitter,omitempty"`
	MaximumInboundPacketLossRateInPercentage  *float64 `json:"maximumInboundPacketLossRateInPercentage,omitempty"`
	MaximumInboundRoundTripDelay              *string  `json:"maximumInboundRoundTripDelay,omitempty"`
	MaximumOutboundJitter                     *string  `json:"maximumOutboundJitter,omitempty"`
	MaximumOutboundPacketLossRateInPercentage *float64 `json:"maximumOutboundPacketLossRateInPercentage,omitempty"`
	MaximumOutboundRoundTripDelay             *string  `json:"maximumOutboundRoundTripDelay,omitempty"`
	MediaDuration                             *string  `json:"mediaDuration,omitempty"`
	NetworkLinkSpeedInBytes                   *int64   `json:"networkLinkSpeedInBytes,omitempty"`
	ODataType                                 *string  `json:"@odata.type,omitempty"`
	OutboundPackets                           *int64   `json:"outboundPackets,omitempty"`
	RemoteIPAddress                           *string  `json:"remoteIPAddress,omitempty"`
	RemotePort                                *int64   `json:"remotePort,omitempty"`
}
