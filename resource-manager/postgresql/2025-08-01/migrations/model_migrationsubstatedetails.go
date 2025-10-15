package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationSubstateDetails struct {
	CurrentSubState   *MigrationSubstate                 `json:"currentSubState,omitempty"`
	DbDetails         *map[string]DatabaseMigrationState `json:"dbDetails,omitempty"`
	ValidationDetails *ValidationDetails                 `json:"validationDetails,omitempty"`
}
