package tables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoredLogs struct {
	AzureAsyncOperationId *string `json:"azureAsyncOperationId,omitempty"`
	EndRestoreTime        *string `json:"endRestoreTime,omitempty"`
	SourceTable           *string `json:"sourceTable,omitempty"`
	StartRestoreTime      *string `json:"startRestoreTime,omitempty"`
}
