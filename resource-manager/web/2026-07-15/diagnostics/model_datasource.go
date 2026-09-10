package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSource struct {
	DataSourceUri *[]NameValuePair `json:"dataSourceUri,omitempty"`
	Instructions  *[]string        `json:"instructions,omitempty"`
}
