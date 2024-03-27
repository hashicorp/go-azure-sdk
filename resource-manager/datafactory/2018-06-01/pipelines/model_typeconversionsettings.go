package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TypeConversionSettings struct {
	AllowDataTruncation  *interface{} `json:"allowDataTruncation,omitempty"`
	Culture              *interface{} `json:"culture,omitempty"`
	DateTimeFormat       *interface{} `json:"dateTimeFormat,omitempty"`
	DateTimeOffsetFormat *interface{} `json:"dateTimeOffsetFormat,omitempty"`
	TimeSpanFormat       *interface{} `json:"timeSpanFormat,omitempty"`
	TreatBooleanAsNumber *interface{} `json:"treatBooleanAsNumber,omitempty"`
}
