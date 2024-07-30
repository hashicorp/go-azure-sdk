package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Video struct {
	AudioBitsPerSample    *int64   `json:"audioBitsPerSample,omitempty"`
	AudioChannels         *int64   `json:"audioChannels,omitempty"`
	AudioFormat           *string  `json:"audioFormat,omitempty"`
	AudioSamplesPerSecond *int64   `json:"audioSamplesPerSecond,omitempty"`
	Bitrate               *int64   `json:"bitrate,omitempty"`
	Duration              *int64   `json:"duration,omitempty"`
	FourCC                *string  `json:"fourCC,omitempty"`
	FrameRate             *float64 `json:"frameRate,omitempty"`
	Height                *int64   `json:"height,omitempty"`
	ODataType             *string  `json:"@odata.type,omitempty"`
	Width                 *int64   `json:"width,omitempty"`
}
