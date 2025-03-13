package proxyoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListWorkspaceKeysResult struct {
	AppInsightsInstrumentationKey *string                        `json:"appInsightsInstrumentationKey,omitempty"`
	ContainerRegistryCredentials  *RegistryListCredentialsResult `json:"containerRegistryCredentials,omitempty"`
	NotebookAccessKeys            *ListNotebookKeysResult        `json:"notebookAccessKeys,omitempty"`
	UserStorageArmId              *string                        `json:"userStorageArmId,omitempty"`
	UserStorageKey                *string                        `json:"userStorageKey,omitempty"`
}
