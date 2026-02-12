package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowSourceSetting struct {
	RowLimit   *int64  `json:"rowLimit,omitempty"`
	SourceName *string `json:"sourceName,omitempty"`
}
