package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SapCloudForCustomerSource struct {
	AdditionalColumns        *interface{} `json:"additionalColumns,omitempty"`
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	HTTPRequestTimeout       *interface{} `json:"httpRequestTimeout,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	Query                    *interface{} `json:"query,omitempty"`
	QueryTimeout             *interface{} `json:"queryTimeout,omitempty"`
	SourceRetryCount         *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
	Type                     string       `json:"type"`
}
