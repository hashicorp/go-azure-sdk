package notebooks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Notebook struct {
	BigDataPool       *BigDataPoolReference      `json:"bigDataPool,omitempty"`
	Cells             []NotebookCell             `json:"cells"`
	Description       *string                    `json:"description,omitempty"`
	Metadata          NotebookMetadata           `json:"metadata"`
	Nbformat          int64                      `json:"nbformat"`
	NbformatMinor     int64                      `json:"nbformat_minor"`
	SessionProperties *NotebookSessionProperties `json:"sessionProperties,omitempty"`
}
