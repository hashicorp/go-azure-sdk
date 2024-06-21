package pricesheets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EAPricesheetDownloadProperties struct {
	DownloadFileProperties *EAPriceSheetProperties `json:"downloadFileProperties,omitempty"`
	DownloadUrl            *string                 `json:"downloadUrl,omitempty"`
	ValidTill              *string                 `json:"validTill,omitempty"`
}
