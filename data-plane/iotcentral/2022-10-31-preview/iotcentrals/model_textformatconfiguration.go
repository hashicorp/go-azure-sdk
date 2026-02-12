package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextFormatConfiguration struct {
	AbbreviateValue *bool             `json:"abbreviateValue,omitempty"`
	DecimalPlaces   *int64            `json:"decimalPlaces,omitempty"`
	TextSize        *float64          `json:"textSize,omitempty"`
	Unit            *TileTextSizeUnit `json:"unit,omitempty"`
	WordWrap        *bool             `json:"wordWrap,omitempty"`
}
