package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageTileConfigurationFormat struct {
	BackgroundColor *string             `json:"backgroundColor,omitempty"`
	FitImage        *bool               `json:"fitImage,omitempty"`
	ShowTitle       *bool               `json:"showTitle,omitempty"`
	TextColor       *string             `json:"textColor,omitempty"`
	TextSize        *float64            `json:"textSize,omitempty"`
	TextSizeUnit    *ImageTileTextUnits `json:"textSizeUnit,omitempty"`
}
