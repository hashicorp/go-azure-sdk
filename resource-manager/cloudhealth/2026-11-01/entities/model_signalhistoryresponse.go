package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalHistoryResponse struct {
	EntityName string                   `json:"entityName"`
	History    []SignalHistoryDataPoint `json:"history"`
	NextMarker *string                  `json:"nextMarker,omitempty"`
	SignalName string                   `json:"signalName"`
}
