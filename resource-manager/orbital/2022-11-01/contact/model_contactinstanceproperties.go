package contact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactInstanceProperties struct {
	EndAzimuthDegrees       *float64 `json:"endAzimuthDegrees,omitempty"`
	EndElevationDegrees     *float64 `json:"endElevationDegrees,omitempty"`
	MaximumElevationDegrees *float64 `json:"maximumElevationDegrees,omitempty"`
	RxEndTime               *string  `json:"rxEndTime,omitempty"`
	RxStartTime             *string  `json:"rxStartTime,omitempty"`
	StartAzimuthDegrees     *float64 `json:"startAzimuthDegrees,omitempty"`
	StartElevationDegrees   *float64 `json:"startElevationDegrees,omitempty"`
	TxEndTime               *string  `json:"txEndTime,omitempty"`
	TxStartTime             *string  `json:"txStartTime,omitempty"`
}
