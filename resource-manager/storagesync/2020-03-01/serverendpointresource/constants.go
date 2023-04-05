package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureStatus string

const (
	FeatureStatusOff FeatureStatus = "off"
	FeatureStatusOn  FeatureStatus = "on"
)

func PossibleValuesForFeatureStatus() []string {
	return []string{
		string(FeatureStatusOff),
		string(FeatureStatusOn),
	}
}

type InitialDownloadPolicy string

const (
	InitialDownloadPolicyAvoidTieredFiles           InitialDownloadPolicy = "AvoidTieredFiles"
	InitialDownloadPolicyNamespaceOnly              InitialDownloadPolicy = "NamespaceOnly"
	InitialDownloadPolicyNamespaceThenModifiedFiles InitialDownloadPolicy = "NamespaceThenModifiedFiles"
)

func PossibleValuesForInitialDownloadPolicy() []string {
	return []string{
		string(InitialDownloadPolicyAvoidTieredFiles),
		string(InitialDownloadPolicyNamespaceOnly),
		string(InitialDownloadPolicyNamespaceThenModifiedFiles),
	}
}

type LocalCacheMode string

const (
	LocalCacheModeDownloadNewAndModifiedFiles LocalCacheMode = "DownloadNewAndModifiedFiles"
	LocalCacheModeUpdateLocallyCachedFiles    LocalCacheMode = "UpdateLocallyCachedFiles"
)

func PossibleValuesForLocalCacheMode() []string {
	return []string{
		string(LocalCacheModeDownloadNewAndModifiedFiles),
		string(LocalCacheModeUpdateLocallyCachedFiles),
	}
}

type ServerEndpointCloudTieringHealthState string

const (
	ServerEndpointCloudTieringHealthStateError   ServerEndpointCloudTieringHealthState = "Error"
	ServerEndpointCloudTieringHealthStateHealthy ServerEndpointCloudTieringHealthState = "Healthy"
)

func PossibleValuesForServerEndpointCloudTieringHealthState() []string {
	return []string{
		string(ServerEndpointCloudTieringHealthStateError),
		string(ServerEndpointCloudTieringHealthStateHealthy),
	}
}

type ServerEndpointOfflineDataTransferState string

const (
	ServerEndpointOfflineDataTransferStateComplete   ServerEndpointOfflineDataTransferState = "Complete"
	ServerEndpointOfflineDataTransferStateInProgress ServerEndpointOfflineDataTransferState = "InProgress"
	ServerEndpointOfflineDataTransferStateNotRunning ServerEndpointOfflineDataTransferState = "NotRunning"
	ServerEndpointOfflineDataTransferStateStopping   ServerEndpointOfflineDataTransferState = "Stopping"
)

func PossibleValuesForServerEndpointOfflineDataTransferState() []string {
	return []string{
		string(ServerEndpointOfflineDataTransferStateComplete),
		string(ServerEndpointOfflineDataTransferStateInProgress),
		string(ServerEndpointOfflineDataTransferStateNotRunning),
		string(ServerEndpointOfflineDataTransferStateStopping),
	}
}

type ServerEndpointSyncActivityState string

const (
	ServerEndpointSyncActivityStateDownload          ServerEndpointSyncActivityState = "Download"
	ServerEndpointSyncActivityStateUpload            ServerEndpointSyncActivityState = "Upload"
	ServerEndpointSyncActivityStateUploadAndDownload ServerEndpointSyncActivityState = "UploadAndDownload"
)

func PossibleValuesForServerEndpointSyncActivityState() []string {
	return []string{
		string(ServerEndpointSyncActivityStateDownload),
		string(ServerEndpointSyncActivityStateUpload),
		string(ServerEndpointSyncActivityStateUploadAndDownload),
	}
}

type ServerEndpointSyncHealthState string

const (
	ServerEndpointSyncHealthStateError                                    ServerEndpointSyncHealthState = "Error"
	ServerEndpointSyncHealthStateHealthy                                  ServerEndpointSyncHealthState = "Healthy"
	ServerEndpointSyncHealthStateNoActivity                               ServerEndpointSyncHealthState = "NoActivity"
	ServerEndpointSyncHealthStateSyncBlockedForChangeDetectionPostRestore ServerEndpointSyncHealthState = "SyncBlockedForChangeDetectionPostRestore"
	ServerEndpointSyncHealthStateSyncBlockedForRestore                    ServerEndpointSyncHealthState = "SyncBlockedForRestore"
)

func PossibleValuesForServerEndpointSyncHealthState() []string {
	return []string{
		string(ServerEndpointSyncHealthStateError),
		string(ServerEndpointSyncHealthStateHealthy),
		string(ServerEndpointSyncHealthStateNoActivity),
		string(ServerEndpointSyncHealthStateSyncBlockedForChangeDetectionPostRestore),
		string(ServerEndpointSyncHealthStateSyncBlockedForRestore),
	}
}

type ServerEndpointSyncMode string

const (
	ServerEndpointSyncModeInitialFullDownload ServerEndpointSyncMode = "InitialFullDownload"
	ServerEndpointSyncModeInitialUpload       ServerEndpointSyncMode = "InitialUpload"
	ServerEndpointSyncModeNamespaceDownload   ServerEndpointSyncMode = "NamespaceDownload"
	ServerEndpointSyncModeRegular             ServerEndpointSyncMode = "Regular"
	ServerEndpointSyncModeSnapshotUpload      ServerEndpointSyncMode = "SnapshotUpload"
)

func PossibleValuesForServerEndpointSyncMode() []string {
	return []string{
		string(ServerEndpointSyncModeInitialFullDownload),
		string(ServerEndpointSyncModeInitialUpload),
		string(ServerEndpointSyncModeNamespaceDownload),
		string(ServerEndpointSyncModeRegular),
		string(ServerEndpointSyncModeSnapshotUpload),
	}
}
