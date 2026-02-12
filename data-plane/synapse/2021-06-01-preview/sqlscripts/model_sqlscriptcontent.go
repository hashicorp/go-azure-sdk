package sqlscripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlScriptContent struct {
	CurrentConnection *SqlConnection     `json:"currentConnection,omitempty"`
	Metadata          *SqlScriptMetadata `json:"metadata,omitempty"`
	Query             string             `json:"query"`
	ResultLimit       *int64             `json:"resultLimit,omitempty"`
}
