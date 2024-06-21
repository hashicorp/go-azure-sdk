package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointSyncStatus struct {
	BackgroundDataDownloadActivity      *ServerEndpointBackgroundDataDownloadActivity `json:"backgroundDataDownloadActivity,omitempty"`
	CombinedHealth                      *ServerEndpointHealthState                    `json:"combinedHealth,omitempty"`
	DownloadActivity                    *ServerEndpointSyncActivityStatus             `json:"downloadActivity,omitempty"`
	DownloadHealth                      *ServerEndpointHealthState                    `json:"downloadHealth,omitempty"`
	DownloadStatus                      *ServerEndpointSyncSessionStatus              `json:"downloadStatus,omitempty"`
	LastUpdatedTimestamp                *string                                       `json:"lastUpdatedTimestamp,omitempty"`
	OfflineDataTransferStatus           *ServerEndpointOfflineDataTransferState       `json:"offlineDataTransferStatus,omitempty"`
	SyncActivity                        *ServerEndpointSyncActivityState              `json:"syncActivity,omitempty"`
	TotalPersistentFilesNotSyncingCount *int64                                        `json:"totalPersistentFilesNotSyncingCount,omitempty"`
	UploadActivity                      *ServerEndpointSyncActivityStatus             `json:"uploadActivity,omitempty"`
	UploadHealth                        *ServerEndpointHealthState                    `json:"uploadHealth,omitempty"`
	UploadStatus                        *ServerEndpointSyncSessionStatus              `json:"uploadStatus,omitempty"`
}
