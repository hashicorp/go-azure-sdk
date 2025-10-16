package datacollectionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OtelTracesDirectDataSource struct {
	EnrichWithReference            *string                                  `json:"enrichWithReference,omitempty"`
	EnrichWithResourceAttributes   *[]string                                `json:"enrichWithResourceAttributes,omitempty"`
	Name                           *string                                  `json:"name,omitempty"`
	ReplaceResourceIdWithReference *bool                                    `json:"replaceResourceIdWithReference,omitempty"`
	Streams                        []KnownOtelTracesDirectDataSourceStreams `json:"streams"`
}
