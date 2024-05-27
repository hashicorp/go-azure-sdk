package datatypes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataTypeProperties struct {
	DatabaseCacheRetention *int64             `json:"databaseCacheRetention,omitempty"`
	DatabaseRetention      *int64             `json:"databaseRetention,omitempty"`
	ProvisioningState      *ProvisioningState `json:"provisioningState,omitempty"`
	State                  *DataTypeState     `json:"state,omitempty"`
	StateReason            *string            `json:"stateReason,omitempty"`
	StorageOutputRetention *int64             `json:"storageOutputRetention,omitempty"`
	VisualizationUrl       *string            `json:"visualizationUrl,omitempty"`
}
