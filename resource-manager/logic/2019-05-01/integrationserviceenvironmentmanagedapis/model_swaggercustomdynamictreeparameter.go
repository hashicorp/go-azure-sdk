package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwaggerCustomDynamicTreeParameter struct {
	ParameterReference    *string      `json:"parameterReference,omitempty"`
	Required              *bool        `json:"required,omitempty"`
	SelectedItemValuePath *string      `json:"selectedItemValuePath,omitempty"`
	Value                 *interface{} `json:"value,omitempty"`
}
