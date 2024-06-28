package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TypeConversionSettings struct {
	AllowDataTruncation  *bool   `json:"allowDataTruncation,omitempty"`
	Culture              *string `json:"culture,omitempty"`
	DateTimeFormat       *string `json:"dateTimeFormat,omitempty"`
	DateTimeOffsetFormat *string `json:"dateTimeOffsetFormat,omitempty"`
	TimeSpanFormat       *string `json:"timeSpanFormat,omitempty"`
	TreatBooleanAsNumber *bool   `json:"treatBooleanAsNumber,omitempty"`
}
