package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDownloadProgress struct {
	DownloadPhase             *DownloadPhase `json:"downloadPhase,omitempty"`
	NumberOfUpdatesDownloaded *int64         `json:"numberOfUpdatesDownloaded,omitempty"`
	NumberOfUpdatesToDownload *int64         `json:"numberOfUpdatesToDownload,omitempty"`
	PercentComplete           *int64         `json:"percentComplete,omitempty"`
	TotalBytesDownloaded      *float64       `json:"totalBytesDownloaded,omitempty"`
	TotalBytesToDownload      *float64       `json:"totalBytesToDownload,omitempty"`
}
