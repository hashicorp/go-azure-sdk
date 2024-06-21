package blobcontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImmutableStorageWithVersioning struct {
	Enabled        *bool           `json:"enabled,omitempty"`
	MigrationState *MigrationState `json:"migrationState,omitempty"`
	TimeStamp      *string         `json:"timeStamp,omitempty"`
}
