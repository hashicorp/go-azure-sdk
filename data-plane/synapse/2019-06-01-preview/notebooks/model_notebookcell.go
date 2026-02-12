package notebooks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookCell struct {
	Attachments *interface{}              `json:"attachments,omitempty"`
	CellType    string                    `json:"cell_type"`
	Metadata    interface{}               `json:"metadata"`
	Outputs     *[]NotebookCellOutputItem `json:"outputs,omitempty"`
	Source      []string                  `json:"source"`
}
