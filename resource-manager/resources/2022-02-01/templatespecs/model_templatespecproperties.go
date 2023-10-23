package templatespecs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateSpecProperties struct {
	Description *string                             `json:"description,omitempty"`
	DisplayName *string                             `json:"displayName,omitempty"`
	Metadata    *interface{}                        `json:"metadata,omitempty"`
	Versions    *map[string]TemplateSpecVersionInfo `json:"versions,omitempty"`
}
