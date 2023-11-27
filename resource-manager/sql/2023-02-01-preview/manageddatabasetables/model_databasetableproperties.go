package manageddatabasetables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseTableProperties struct {
	MemoryOptimized *bool              `json:"memoryOptimized,omitempty"`
	TemporalType    *TableTemporalType `json:"temporalType,omitempty"`
}
