package bicepclient

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DecompileOperationSuccessResponse struct {
	EntryPoint string           `json:"entryPoint"`
	Files      []FileDefinition `json:"files"`
}
