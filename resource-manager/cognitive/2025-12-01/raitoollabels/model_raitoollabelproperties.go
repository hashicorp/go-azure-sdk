package raitoollabels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiToolLabelProperties struct {
	AccountScope       *RaiToolLabelPropertiesAccountScope        `json:"accountScope,omitempty"`
	ProjectScopes      *[]RaiToolLabelPropertiesProjectScopesItem `json:"projectScopes,omitempty"`
	ToolConnectionName string                                     `json:"toolConnectionName"`
}
