package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesforceV2Source struct {
	AdditionalColumns        *interface{} `json:"additionalColumns,omitempty"`
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	IncludeDeletedObjects    *interface{} `json:"includeDeletedObjects,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	QueryTimeout             *interface{} `json:"queryTimeout,omitempty"`
	SOQLQuery                *interface{} `json:"SOQLQuery,omitempty"`
	SourceRetryCount         *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
	Type                     string       `json:"type"`
}
