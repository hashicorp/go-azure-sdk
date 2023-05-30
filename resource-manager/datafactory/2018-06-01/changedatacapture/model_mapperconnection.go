package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MapperConnection struct {
	CommonDslConnectorProperties *[]MapperDslConnectorProperties `json:"commonDslConnectorProperties,omitempty"`
	IsInlineDataset              *bool                           `json:"isInlineDataset,omitempty"`
	LinkedService                *LinkedServiceReference         `json:"linkedService,omitempty"`
	LinkedServiceType            *string                         `json:"linkedServiceType,omitempty"`
	Type                         ConnectionType                  `json:"type"`
}
