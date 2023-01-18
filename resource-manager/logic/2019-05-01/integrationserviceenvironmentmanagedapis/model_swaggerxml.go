package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwaggerXml struct {
	Attribute  *bool                   `json:"attribute,omitempty"`
	Extensions *map[string]interface{} `json:"extensions,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Namespace  *string                 `json:"namespace,omitempty"`
	Prefix     *string                 `json:"prefix,omitempty"`
	Wrapped    *bool                   `json:"wrapped,omitempty"`
}
