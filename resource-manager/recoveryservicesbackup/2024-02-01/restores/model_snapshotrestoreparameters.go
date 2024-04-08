package restores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnapshotRestoreParameters struct {
	LogPointInTimeForDBRecovery *string `json:"logPointInTimeForDBRecovery,omitempty"`
	SkipAttachAndMount          *bool   `json:"skipAttachAndMount,omitempty"`
}
