package bms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreValidateEnableBackupRequest struct {
	Properties   *string         `json:"properties,omitempty"`
	ResourceId   *string         `json:"resourceId,omitempty"`
	ResourceType *DataSourceType `json:"resourceType,omitempty"`
	VaultId      *string         `json:"vaultId,omitempty"`
}
