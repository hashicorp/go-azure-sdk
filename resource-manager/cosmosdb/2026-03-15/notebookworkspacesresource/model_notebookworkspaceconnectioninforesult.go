package notebookworkspacesresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookWorkspaceConnectionInfoResult struct {
	AuthToken              *string `json:"authToken,omitempty"`
	NotebookServerEndpoint *string `json:"notebookServerEndpoint,omitempty"`
}
