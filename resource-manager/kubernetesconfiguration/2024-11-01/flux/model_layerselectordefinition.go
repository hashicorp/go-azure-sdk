package flux

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LayerSelectorDefinition struct {
	MediaType *string        `json:"mediaType,omitempty"`
	Operation *OperationType `json:"operation,omitempty"`
}
