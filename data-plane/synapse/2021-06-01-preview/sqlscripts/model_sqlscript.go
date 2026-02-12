package sqlscripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlScript struct {
	Content     SqlScriptContent `json:"content"`
	Description *string          `json:"description,omitempty"`
	Folder      *SqlScriptFolder `json:"folder,omitempty"`
	Type        *SqlScriptType   `json:"type,omitempty"`
}
