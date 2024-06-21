package containers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RefreshDetails struct {
	ErrorManifestFile                *string `json:"errorManifestFile,omitempty"`
	InProgressRefreshJobId           *string `json:"inProgressRefreshJobId,omitempty"`
	LastCompletedRefreshJobTimeInUTC *string `json:"lastCompletedRefreshJobTimeInUTC,omitempty"`
	LastJob                          *string `json:"lastJob,omitempty"`
}
