package namespaceassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDatasetDataPoint struct {
	DataPointConfiguration *string `json:"dataPointConfiguration,omitempty"`
	DataSource             string  `json:"dataSource"`
	Name                   string  `json:"name"`
	TypeRef                *string `json:"typeRef,omitempty"`
}
