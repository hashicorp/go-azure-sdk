package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HTTPReadSettings struct {
	AdditionalColumns        *interface{} `json:"additionalColumns,omitempty"`
	AdditionalHeaders        *interface{} `json:"additionalHeaders,omitempty"`
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	RequestBody              *interface{} `json:"requestBody,omitempty"`
	RequestMethod            *interface{} `json:"requestMethod,omitempty"`
	RequestTimeout           *interface{} `json:"requestTimeout,omitempty"`
	Type                     string       `json:"type"`
}
