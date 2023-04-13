package operations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Dimension struct {
	DisplayName           *string `json:"displayName,omitempty"`
	InternalMetricName    *string `json:"internalMetricName,omitempty"`
	InternalName          *string `json:"internalName,omitempty"`
	Name                  *string `json:"name,omitempty"`
	SourceMdmNamespace    *string `json:"sourceMdmNamespace,omitempty"`
	ToBeExportedToShoebox *bool   `json:"toBeExportedToShoebox,omitempty"`
}
