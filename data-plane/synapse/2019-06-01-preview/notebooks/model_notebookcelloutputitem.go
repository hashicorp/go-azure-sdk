package notebooks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookCellOutputItem struct {
	Data           *interface{}   `json:"data,omitempty"`
	ExecutionCount *int64         `json:"execution_count,omitempty"`
	Metadata       *interface{}   `json:"metadata,omitempty"`
	Name           *string        `json:"name,omitempty"`
	OutputType     CellOutputType `json:"output_type"`
	Text           *interface{}   `json:"text,omitempty"`
}
