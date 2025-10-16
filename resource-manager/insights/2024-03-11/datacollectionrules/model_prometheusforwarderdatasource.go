package datacollectionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrometheusForwarderDataSource struct {
	CustomVMScrapeConfig *[]interface{}                               `json:"customVMScrapeConfig,omitempty"`
	LabelIncludeFilter   *map[string]string                           `json:"labelIncludeFilter,omitempty"`
	Name                 *string                                      `json:"name,omitempty"`
	Streams              *[]KnownPrometheusForwarderDataSourceStreams `json:"streams,omitempty"`
}
