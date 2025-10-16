package namespaceassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceAssetStatusManagementAction struct {
	Error                          *StatusError                     `json:"error,omitempty"`
	Name                           string                           `json:"name"`
	RequestMessageSchemaReference  *NamespaceMessageSchemaReference `json:"requestMessageSchemaReference,omitempty"`
	ResponseMessageSchemaReference *NamespaceMessageSchemaReference `json:"responseMessageSchemaReference,omitempty"`
}
