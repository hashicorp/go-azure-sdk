package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointBackgroundDataDownloadActivity struct {
	DownloadedBytes  *int64  `json:"downloadedBytes,omitempty"`
	PercentProgress  *int64  `json:"percentProgress,omitempty"`
	StartedTimestamp *string `json:"startedTimestamp,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty"`
}
