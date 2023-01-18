package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwaggerExternalDocumentation struct {
	Description *string                 `json:"description,omitempty"`
	Extensions  *map[string]interface{} `json:"extensions,omitempty"`
	Uri         *string                 `json:"uri,omitempty"`
}
