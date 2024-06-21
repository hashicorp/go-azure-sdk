package syncmembers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncFullSchemaProperties struct {
	LastUpdateTime *string                `json:"lastUpdateTime,omitempty"`
	Tables         *[]SyncFullSchemaTable `json:"tables,omitempty"`
}
