package hybridrunbookworker

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridRunbookWorkerProperties struct {
	IP                 *string     `json:"ip,omitempty"`
	LastSeenDateTime   *string     `json:"lastSeenDateTime,omitempty"`
	RegisteredDateTime *string     `json:"registeredDateTime,omitempty"`
	VMResourceId       *string     `json:"vmResourceId,omitempty"`
	WorkerName         *string     `json:"workerName,omitempty"`
	WorkerType         *WorkerType `json:"workerType,omitempty"`
}
