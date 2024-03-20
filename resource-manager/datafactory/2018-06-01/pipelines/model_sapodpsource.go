package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SapOdpSource struct {
	AdditionalColumns        *interface{} `json:"additionalColumns,omitempty"`
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	ExtractionMode           *interface{} `json:"extractionMode,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	Projection               *interface{} `json:"projection,omitempty"`
	QueryTimeout             *interface{} `json:"queryTimeout,omitempty"`
	Selection                *interface{} `json:"selection,omitempty"`
	SourceRetryCount         *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
	SubscriberProcess        *interface{} `json:"subscriberProcess,omitempty"`
	Type                     string       `json:"type"`
}
