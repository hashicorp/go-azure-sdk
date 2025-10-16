package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CcpResponseConfig struct {
	CompressionAlgo               *string  `json:"compressionAlgo,omitempty"`
	ConvertChildPropertiesToArray *bool    `json:"convertChildPropertiesToArray,omitempty"`
	CsvDelimiter                  *string  `json:"csvDelimiter,omitempty"`
	CsvEscape                     *string  `json:"csvEscape,omitempty"`
	EventsJsonPaths               []string `json:"eventsJsonPaths"`
	Format                        *string  `json:"format,omitempty"`
	HasCsvBoundary                *bool    `json:"hasCsvBoundary,omitempty"`
	HasCsvHeader                  *bool    `json:"hasCsvHeader,omitempty"`
	IsGzipCompressed              *bool    `json:"isGzipCompressed,omitempty"`
	SuccessStatusJsonPath         *string  `json:"successStatusJsonPath,omitempty"`
	SuccessStatusValue            *string  `json:"successStatusValue,omitempty"`
}
