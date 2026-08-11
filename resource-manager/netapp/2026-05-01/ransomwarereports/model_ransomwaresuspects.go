package ransomwarereports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RansomwareSuspects struct {
	Extension    *string                      `json:"extension,omitempty"`
	FileCount    *int64                       `json:"fileCount,omitempty"`
	Resolution   *RansomwareSuspectResolution `json:"resolution,omitempty"`
	SuspectFiles *[]SuspectFile               `json:"suspectFiles,omitempty"`
}
