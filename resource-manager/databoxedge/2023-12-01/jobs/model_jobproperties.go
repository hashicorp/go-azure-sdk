package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobProperties struct {
	CurrentStage       *UpdateOperationStage   `json:"currentStage,omitempty"`
	DownloadProgress   *UpdateDownloadProgress `json:"downloadProgress,omitempty"`
	ErrorManifestFile  *string                 `json:"errorManifestFile,omitempty"`
	Folder             *string                 `json:"folder,omitempty"`
	InstallProgress    *UpdateInstallProgress  `json:"installProgress,omitempty"`
	JobType            *JobType                `json:"jobType,omitempty"`
	RefreshedEntityId  *string                 `json:"refreshedEntityId,omitempty"`
	TotalRefreshErrors *int64                  `json:"totalRefreshErrors,omitempty"`
}
