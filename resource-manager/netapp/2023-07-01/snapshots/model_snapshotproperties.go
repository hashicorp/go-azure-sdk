package snapshots

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnapshotProperties struct {
	Created           *string `json:"created,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
	SnapshotId        *string `json:"snapshotId,omitempty"`
}
