package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightsTableResult struct {
	Columns *[]InsightsTableResultColumnsInlined `json:"columns,omitempty"`
	Rows    *[][]string                          `json:"rows,omitempty"`
}
