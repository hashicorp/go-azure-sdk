package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceKeyDetails struct {
	KeyVaultUrl *string `json:"keyVaultUrl,omitempty"`
	Name        *string `json:"name,omitempty"`
}
