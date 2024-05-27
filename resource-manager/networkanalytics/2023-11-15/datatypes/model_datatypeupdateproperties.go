package datatypes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataTypeUpdateProperties struct {
	DatabaseCacheRetention *int64         `json:"databaseCacheRetention,omitempty"`
	DatabaseRetention      *int64         `json:"databaseRetention,omitempty"`
	State                  *DataTypeState `json:"state,omitempty"`
	StorageOutputRetention *int64         `json:"storageOutputRetention,omitempty"`
}
