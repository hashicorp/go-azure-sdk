package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsDeviceInfo struct {
	CaptureDeviceDriver              *string  `json:"captureDeviceDriver,omitempty"`
	CaptureDeviceName                *string  `json:"captureDeviceName,omitempty"`
	CaptureNotFunctioningEventRatio  *float64 `json:"captureNotFunctioningEventRatio,omitempty"`
	CpuInsufficentEventRatio         *float64 `json:"cpuInsufficentEventRatio,omitempty"`
	DeviceClippingEventRatio         *float64 `json:"deviceClippingEventRatio,omitempty"`
	DeviceGlitchEventRatio           *float64 `json:"deviceGlitchEventRatio,omitempty"`
	HowlingEventCount                *int64   `json:"howlingEventCount,omitempty"`
	InitialSignalLevelRootMeanSquare *float64 `json:"initialSignalLevelRootMeanSquare,omitempty"`
	LowSpeechLevelEventRatio         *float64 `json:"lowSpeechLevelEventRatio,omitempty"`
	LowSpeechToNoiseEventRatio       *float64 `json:"lowSpeechToNoiseEventRatio,omitempty"`
	MicGlitchRate                    *float64 `json:"micGlitchRate,omitempty"`
	ODataType                        *string  `json:"@odata.type,omitempty"`
	ReceivedNoiseLevel               *int64   `json:"receivedNoiseLevel,omitempty"`
	ReceivedSignalLevel              *int64   `json:"receivedSignalLevel,omitempty"`
	RenderDeviceDriver               *string  `json:"renderDeviceDriver,omitempty"`
	RenderDeviceName                 *string  `json:"renderDeviceName,omitempty"`
	RenderMuteEventRatio             *float64 `json:"renderMuteEventRatio,omitempty"`
	RenderNotFunctioningEventRatio   *float64 `json:"renderNotFunctioningEventRatio,omitempty"`
	RenderZeroVolumeEventRatio       *float64 `json:"renderZeroVolumeEventRatio,omitempty"`
	SentNoiseLevel                   *int64   `json:"sentNoiseLevel,omitempty"`
	SentSignalLevel                  *int64   `json:"sentSignalLevel,omitempty"`
	SpeakerGlitchRate                *float64 `json:"speakerGlitchRate,omitempty"`
}
