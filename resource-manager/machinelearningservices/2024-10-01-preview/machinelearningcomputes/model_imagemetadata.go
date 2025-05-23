package machinelearningcomputes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageMetadata struct {
	CurrentImageVersion    *string           `json:"currentImageVersion,omitempty"`
	IsLatestOsImageVersion *bool             `json:"isLatestOsImageVersion,omitempty"`
	LatestImageVersion     *string           `json:"latestImageVersion,omitempty"`
	OsPatchingStatus       *OsPatchingStatus `json:"osPatchingStatus,omitempty"`
}
