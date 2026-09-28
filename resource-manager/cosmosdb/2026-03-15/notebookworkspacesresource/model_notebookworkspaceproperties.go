package notebookworkspacesresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookWorkspaceProperties struct {
	NotebookServerEndpoint *string `json:"notebookServerEndpoint,omitempty"`
	Status                 *string `json:"status,omitempty"`
}
