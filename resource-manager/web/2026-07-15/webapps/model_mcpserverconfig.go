package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type McpServerConfig struct {
	Description *string   `json:"description,omitempty"`
	Enabled     *bool     `json:"enabled,omitempty"`
	Endpoint    *string   `json:"endpoint,omitempty"`
	Name        *string   `json:"name,omitempty"`
	ToolList    *[]string `json:"toolList,omitempty"`
}
