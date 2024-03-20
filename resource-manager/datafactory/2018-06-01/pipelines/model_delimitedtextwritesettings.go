package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelimitedTextWriteSettings struct {
	FileExtension  interface{}  `json:"fileExtension"`
	FileNamePrefix *interface{} `json:"fileNamePrefix,omitempty"`
	MaxRowsPerFile *interface{} `json:"maxRowsPerFile,omitempty"`
	QuoteAllText   *interface{} `json:"quoteAllText,omitempty"`
	Type           string       `json:"type"`
}
