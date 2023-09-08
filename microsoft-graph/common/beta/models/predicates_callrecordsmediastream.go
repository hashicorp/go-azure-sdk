package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamOperationPredicate struct {
	AverageAudioNetworkJitter         *string
	AverageBandwidthEstimate          *int64
	AverageFreezeDuration             *string
	AverageJitter                     *string
	AverageRoundTripTime              *string
	EndDateTime                       *string
	IsAudioForwardErrorCorrectionUsed *bool
	MaxAudioNetworkJitter             *string
	MaxJitter                         *string
	MaxRoundTripTime                  *string
	ODataType                         *string
	PacketUtilization                 *int64
	RmsFreezeDuration                 *string
	StartDateTime                     *string
	StreamId                          *string
	WasMediaBypassed                  *bool
}

func (p CallRecordsMediaStreamOperationPredicate) Matches(input CallRecordsMediaStream) bool {

	if p.AverageAudioNetworkJitter != nil && (input.AverageAudioNetworkJitter == nil || *p.AverageAudioNetworkJitter != *input.AverageAudioNetworkJitter) {
		return false
	}

	if p.AverageBandwidthEstimate != nil && (input.AverageBandwidthEstimate == nil || *p.AverageBandwidthEstimate != *input.AverageBandwidthEstimate) {
		return false
	}

	if p.AverageFreezeDuration != nil && (input.AverageFreezeDuration == nil || *p.AverageFreezeDuration != *input.AverageFreezeDuration) {
		return false
	}

	if p.AverageJitter != nil && (input.AverageJitter == nil || *p.AverageJitter != *input.AverageJitter) {
		return false
	}

	if p.AverageRoundTripTime != nil && (input.AverageRoundTripTime == nil || *p.AverageRoundTripTime != *input.AverageRoundTripTime) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.IsAudioForwardErrorCorrectionUsed != nil && (input.IsAudioForwardErrorCorrectionUsed == nil || *p.IsAudioForwardErrorCorrectionUsed != *input.IsAudioForwardErrorCorrectionUsed) {
		return false
	}

	if p.MaxAudioNetworkJitter != nil && (input.MaxAudioNetworkJitter == nil || *p.MaxAudioNetworkJitter != *input.MaxAudioNetworkJitter) {
		return false
	}

	if p.MaxJitter != nil && (input.MaxJitter == nil || *p.MaxJitter != *input.MaxJitter) {
		return false
	}

	if p.MaxRoundTripTime != nil && (input.MaxRoundTripTime == nil || *p.MaxRoundTripTime != *input.MaxRoundTripTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PacketUtilization != nil && (input.PacketUtilization == nil || *p.PacketUtilization != *input.PacketUtilization) {
		return false
	}

	if p.RmsFreezeDuration != nil && (input.RmsFreezeDuration == nil || *p.RmsFreezeDuration != *input.RmsFreezeDuration) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.StreamId != nil && (input.StreamId == nil || *p.StreamId != *input.StreamId) {
		return false
	}

	if p.WasMediaBypassed != nil && (input.WasMediaBypassed == nil || *p.WasMediaBypassed != *input.WasMediaBypassed) {
		return false
	}

	return true
}
