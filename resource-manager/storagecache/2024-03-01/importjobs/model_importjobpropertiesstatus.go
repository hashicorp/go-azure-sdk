package importjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportJobPropertiesStatus struct {
	BlobsImportedPerSecond *int64            `json:"blobsImportedPerSecond,omitempty"`
	BlobsWalkedPerSecond   *int64            `json:"blobsWalkedPerSecond,omitempty"`
	LastCompletionTime     *string           `json:"lastCompletionTime,omitempty"`
	LastStartedTime        *string           `json:"lastStartedTime,omitempty"`
	State                  *ImportStatusType `json:"state,omitempty"`
	StatusMessage          *string           `json:"statusMessage,omitempty"`
	TotalBlobsImported     *int64            `json:"totalBlobsImported,omitempty"`
	TotalBlobsWalked       *int64            `json:"totalBlobsWalked,omitempty"`
	TotalConflicts         *int64            `json:"totalConflicts,omitempty"`
	TotalErrors            *int64            `json:"totalErrors,omitempty"`
}
