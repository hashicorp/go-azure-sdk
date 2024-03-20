package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CosmosDbSqlApiSource struct {
	AdditionalColumns        *interface{} `json:"additionalColumns,omitempty"`
	DetectDatetime           *interface{} `json:"detectDatetime,omitempty"`
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	PageSize                 *interface{} `json:"pageSize,omitempty"`
	PreferredRegions         *interface{} `json:"preferredRegions,omitempty"`
	Query                    *interface{} `json:"query,omitempty"`
	SourceRetryCount         *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
	Type                     string       `json:"type"`
}
