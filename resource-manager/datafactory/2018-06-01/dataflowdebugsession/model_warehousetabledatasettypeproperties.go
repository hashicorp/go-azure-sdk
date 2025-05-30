package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WarehouseTableDatasetTypeProperties struct {
	Schema *interface{} `json:"schema,omitempty"`
	Table  *interface{} `json:"table,omitempty"`
}
