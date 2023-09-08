package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsDeviceInfoOperationPredicate struct {
	CaptureDeviceDriver *string
	CaptureDeviceName   *string
	HowlingEventCount   *int64
	ODataType           *string
	ReceivedNoiseLevel  *int64
	ReceivedSignalLevel *int64
	RenderDeviceDriver  *string
	RenderDeviceName    *string
	SentNoiseLevel      *int64
	SentSignalLevel     *int64
}

func (p CallRecordsDeviceInfoOperationPredicate) Matches(input CallRecordsDeviceInfo) bool {

	if p.CaptureDeviceDriver != nil && (input.CaptureDeviceDriver == nil || *p.CaptureDeviceDriver != *input.CaptureDeviceDriver) {
		return false
	}

	if p.CaptureDeviceName != nil && (input.CaptureDeviceName == nil || *p.CaptureDeviceName != *input.CaptureDeviceName) {
		return false
	}

	if p.HowlingEventCount != nil && (input.HowlingEventCount == nil || *p.HowlingEventCount != *input.HowlingEventCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReceivedNoiseLevel != nil && (input.ReceivedNoiseLevel == nil || *p.ReceivedNoiseLevel != *input.ReceivedNoiseLevel) {
		return false
	}

	if p.ReceivedSignalLevel != nil && (input.ReceivedSignalLevel == nil || *p.ReceivedSignalLevel != *input.ReceivedSignalLevel) {
		return false
	}

	if p.RenderDeviceDriver != nil && (input.RenderDeviceDriver == nil || *p.RenderDeviceDriver != *input.RenderDeviceDriver) {
		return false
	}

	if p.RenderDeviceName != nil && (input.RenderDeviceName == nil || *p.RenderDeviceName != *input.RenderDeviceName) {
		return false
	}

	if p.SentNoiseLevel != nil && (input.SentNoiseLevel == nil || *p.SentNoiseLevel != *input.SentNoiseLevel) {
		return false
	}

	if p.SentSignalLevel != nil && (input.SentSignalLevel == nil || *p.SentSignalLevel != *input.SentSignalLevel) {
		return false
	}

	return true
}
