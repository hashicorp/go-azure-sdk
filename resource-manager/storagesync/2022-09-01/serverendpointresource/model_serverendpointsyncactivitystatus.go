package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointSyncActivityStatus struct {
	AppliedBytes            *int64                  `json:"appliedBytes,omitempty"`
	AppliedItemCount        *int64                  `json:"appliedItemCount,omitempty"`
	PerItemErrorCount       *int64                  `json:"perItemErrorCount,omitempty"`
	SessionMinutesRemaining *int64                  `json:"sessionMinutesRemaining,omitempty"`
	SyncMode                *ServerEndpointSyncMode `json:"syncMode,omitempty"`
	Timestamp               *string                 `json:"timestamp,omitempty"`
	TotalBytes              *int64                  `json:"totalBytes,omitempty"`
	TotalItemCount          *int64                  `json:"totalItemCount,omitempty"`
}
