package vaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultPropertiesMoveDetails struct {
	CompletionTimeUtc *string `json:"completionTimeUtc,omitempty"`
	OperationId       *string `json:"operationId,omitempty"`
	SourceResourceId  *string `json:"sourceResourceId,omitempty"`
	StartTimeUtc      *string `json:"startTimeUtc,omitempty"`
	TargetResourceId  *string `json:"targetResourceId,omitempty"`
}
