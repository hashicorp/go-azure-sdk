package datasets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextFormat struct {
	ColumnDelimiter  *interface{} `json:"columnDelimiter,omitempty"`
	Deserializer     *interface{} `json:"deserializer,omitempty"`
	EncodingName     *interface{} `json:"encodingName,omitempty"`
	EscapeChar       *interface{} `json:"escapeChar,omitempty"`
	FirstRowAsHeader *interface{} `json:"firstRowAsHeader,omitempty"`
	NullValue        *interface{} `json:"nullValue,omitempty"`
	QuoteChar        *interface{} `json:"quoteChar,omitempty"`
	RowDelimiter     *interface{} `json:"rowDelimiter,omitempty"`
	Serializer       *interface{} `json:"serializer,omitempty"`
	SkipLineCount    *interface{} `json:"skipLineCount,omitempty"`
	TreatEmptyAsNull *interface{} `json:"treatEmptyAsNull,omitempty"`
	Type             string       `json:"type"`
}
