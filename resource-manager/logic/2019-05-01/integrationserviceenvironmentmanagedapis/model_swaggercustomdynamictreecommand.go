package integrationserviceenvironmentmanagedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwaggerCustomDynamicTreeCommand struct {
	ItemFullTitlePath *string                                       `json:"itemFullTitlePath,omitempty"`
	ItemIsParent      *string                                       `json:"itemIsParent,omitempty"`
	ItemTitlePath     *string                                       `json:"itemTitlePath,omitempty"`
	ItemValuePath     *string                                       `json:"itemValuePath,omitempty"`
	ItemsPath         *string                                       `json:"itemsPath,omitempty"`
	OperationId       *string                                       `json:"operationId,omitempty"`
	Parameters        *map[string]SwaggerCustomDynamicTreeParameter `json:"parameters,omitempty"`
	SelectableFilter  *string                                       `json:"selectableFilter,omitempty"`
}
