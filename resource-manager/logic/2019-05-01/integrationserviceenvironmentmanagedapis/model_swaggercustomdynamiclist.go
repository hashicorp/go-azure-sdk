package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwaggerCustomDynamicList struct {
	BuiltInOperation *string                                    `json:"builtInOperation,omitempty"`
	ItemTitlePath    *string                                    `json:"itemTitlePath,omitempty"`
	ItemValuePath    *string                                    `json:"itemValuePath,omitempty"`
	ItemsPath        *string                                    `json:"itemsPath,omitempty"`
	OperationId      *string                                    `json:"operationId,omitempty"`
	Parameters       *map[string]SwaggerCustomDynamicProperties `json:"parameters,omitempty"`
}
