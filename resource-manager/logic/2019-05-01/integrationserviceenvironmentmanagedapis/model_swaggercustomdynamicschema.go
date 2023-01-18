package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwaggerCustomDynamicSchema struct {
	OperationId *string                 `json:"operationId,omitempty"`
	Parameters  *map[string]interface{} `json:"parameters,omitempty"`
	ValuePath   *string                 `json:"valuePath,omitempty"`
}
