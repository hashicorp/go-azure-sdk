package namespaceassets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceAssetStatus struct {
	Config           *StatusConfig                          `json:"config,omitempty"`
	Datasets         *[]NamespaceAssetStatusDataset         `json:"datasets,omitempty"`
	EventGroups      *[]NamespaceAssetStatusEventGroup      `json:"eventGroups,omitempty"`
	ManagementGroups *[]NamespaceAssetStatusManagementGroup `json:"managementGroups,omitempty"`
	Streams          *[]NamespaceAssetStatusStream          `json:"streams,omitempty"`
}
