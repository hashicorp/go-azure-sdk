package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationStatus struct {
	CurrentSubStateDetails *MigrationSubstateDetails `json:"currentSubStateDetails,omitempty"`
	Error                  *string                   `json:"error,omitempty"`
	State                  *MigrationState           `json:"state,omitempty"`
}
