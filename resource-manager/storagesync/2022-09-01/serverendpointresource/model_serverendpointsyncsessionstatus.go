package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointSyncSessionStatus struct {
	FilesNotSyncingErrors          *[]ServerEndpointFilesNotSyncingError `json:"filesNotSyncingErrors,omitempty"`
	LastSyncMode                   *ServerEndpointSyncMode               `json:"lastSyncMode,omitempty"`
	LastSyncPerItemErrorCount      *int64                                `json:"lastSyncPerItemErrorCount,omitempty"`
	LastSyncResult                 *int64                                `json:"lastSyncResult,omitempty"`
	LastSyncSuccessTimestamp       *string                               `json:"lastSyncSuccessTimestamp,omitempty"`
	LastSyncTimestamp              *string                               `json:"lastSyncTimestamp,omitempty"`
	PersistentFilesNotSyncingCount *int64                                `json:"persistentFilesNotSyncingCount,omitempty"`
	TransientFilesNotSyncingCount  *int64                                `json:"transientFilesNotSyncingCount,omitempty"`
}
