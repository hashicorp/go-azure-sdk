package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AiIntegration struct {
	ApiSpecPath *string       `json:"apiSpecPath,omitempty"`
	Mcp         *McpSettings  `json:"mcp,omitempty"`
	SiteAuth    *SiteAuthInfo `json:"siteAuth,omitempty"`
}
