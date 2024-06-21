package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DbLevelValidationStatus struct {
	DatabaseName *string                  `json:"databaseName,omitempty"`
	EndedOn      *string                  `json:"endedOn,omitempty"`
	StartedOn    *string                  `json:"startedOn,omitempty"`
	Summary      *[]ValidationSummaryItem `json:"summary,omitempty"`
}
