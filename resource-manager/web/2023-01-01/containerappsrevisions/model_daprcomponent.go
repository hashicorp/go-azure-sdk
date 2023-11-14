package containerappsrevisions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprComponent struct {
	Metadata *[]DaprMetadata `json:"metadata,omitempty"`
	Name     *string         `json:"name,omitempty"`
	Type     *string         `json:"type,omitempty"`
	Version  *string         `json:"version,omitempty"`
}
