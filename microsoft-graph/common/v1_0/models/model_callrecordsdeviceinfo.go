package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsDeviceInfo struct {
	CaptureDeviceDriver *string `json:"captureDeviceDriver,omitempty"`
	CaptureDeviceName   *string `json:"captureDeviceName,omitempty"`
	HowlingEventCount   *int64  `json:"howlingEventCount,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	ReceivedNoiseLevel  *int64  `json:"receivedNoiseLevel,omitempty"`
	ReceivedSignalLevel *int64  `json:"receivedSignalLevel,omitempty"`
	RenderDeviceDriver  *string `json:"renderDeviceDriver,omitempty"`
	RenderDeviceName    *string `json:"renderDeviceName,omitempty"`
	SentNoiseLevel      *int64  `json:"sentNoiseLevel,omitempty"`
	SentSignalLevel     *int64  `json:"sentSignalLevel,omitempty"`
}
