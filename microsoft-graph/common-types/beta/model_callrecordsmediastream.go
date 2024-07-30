package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStream struct {
	AudioCodec                               *CallRecordsMediaStreamAudioCodec      `json:"audioCodec,omitempty"`
	AverageAudioDegradation                  *float64                               `json:"averageAudioDegradation,omitempty"`
	AverageAudioNetworkJitter                *string                                `json:"averageAudioNetworkJitter,omitempty"`
	AverageBandwidthEstimate                 *int64                                 `json:"averageBandwidthEstimate,omitempty"`
	AverageFreezeDuration                    *string                                `json:"averageFreezeDuration,omitempty"`
	AverageJitter                            *string                                `json:"averageJitter,omitempty"`
	AveragePacketLossRate                    *float64                               `json:"averagePacketLossRate,omitempty"`
	AverageRatioOfConcealedSamples           *float64                               `json:"averageRatioOfConcealedSamples,omitempty"`
	AverageReceivedFrameRate                 *float64                               `json:"averageReceivedFrameRate,omitempty"`
	AverageRoundTripTime                     *string                                `json:"averageRoundTripTime,omitempty"`
	AverageVideoFrameLossPercentage          *float64                               `json:"averageVideoFrameLossPercentage,omitempty"`
	AverageVideoFrameRate                    *float64                               `json:"averageVideoFrameRate,omitempty"`
	AverageVideoPacketLossRate               *float64                               `json:"averageVideoPacketLossRate,omitempty"`
	EndDateTime                              *string                                `json:"endDateTime,omitempty"`
	IsAudioForwardErrorCorrectionUsed        *bool                                  `json:"isAudioForwardErrorCorrectionUsed,omitempty"`
	LowFrameRateRatio                        *float64                               `json:"lowFrameRateRatio,omitempty"`
	LowVideoProcessingCapabilityRatio        *float64                               `json:"lowVideoProcessingCapabilityRatio,omitempty"`
	MaxAudioNetworkJitter                    *string                                `json:"maxAudioNetworkJitter,omitempty"`
	MaxJitter                                *string                                `json:"maxJitter,omitempty"`
	MaxPacketLossRate                        *float64                               `json:"maxPacketLossRate,omitempty"`
	MaxRatioOfConcealedSamples               *float64                               `json:"maxRatioOfConcealedSamples,omitempty"`
	MaxRoundTripTime                         *string                                `json:"maxRoundTripTime,omitempty"`
	ODataType                                *string                                `json:"@odata.type,omitempty"`
	PacketUtilization                        *int64                                 `json:"packetUtilization,omitempty"`
	PostForwardErrorCorrectionPacketLossRate *float64                               `json:"postForwardErrorCorrectionPacketLossRate,omitempty"`
	RmsFreezeDuration                        *string                                `json:"rmsFreezeDuration,omitempty"`
	StartDateTime                            *string                                `json:"startDateTime,omitempty"`
	StreamDirection                          *CallRecordsMediaStreamStreamDirection `json:"streamDirection,omitempty"`
	StreamId                                 *string                                `json:"streamId,omitempty"`
	VideoCodec                               *CallRecordsMediaStreamVideoCodec      `json:"videoCodec,omitempty"`
	WasMediaBypassed                         *bool                                  `json:"wasMediaBypassed,omitempty"`
}
