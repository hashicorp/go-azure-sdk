package restorepoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointProperties struct {
	EarliestRestoreDate      *string           `json:"earliestRestoreDate,omitempty"`
	RestorePointCreationDate *string           `json:"restorePointCreationDate,omitempty"`
	RestorePointLabel        *string           `json:"restorePointLabel,omitempty"`
	RestorePointType         *RestorePointType `json:"restorePointType,omitempty"`
}
