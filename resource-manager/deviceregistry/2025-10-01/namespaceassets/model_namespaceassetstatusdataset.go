package namespaceassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceAssetStatusDataset struct {
	Error                  *StatusError                     `json:"error,omitempty"`
	MessageSchemaReference *NamespaceMessageSchemaReference `json:"messageSchemaReference,omitempty"`
	Name                   string                           `json:"name"`
}
