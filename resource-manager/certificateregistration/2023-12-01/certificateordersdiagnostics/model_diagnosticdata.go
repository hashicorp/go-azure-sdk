package certificateordersdiagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticData struct {
	RenderingProperties *Rendering               `json:"renderingProperties,omitempty"`
	Table               *DataTableResponseObject `json:"table,omitempty"`
}
