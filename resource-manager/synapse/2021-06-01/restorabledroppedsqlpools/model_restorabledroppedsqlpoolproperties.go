package restorabledroppedsqlpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedSqlPoolProperties struct {
	CreationDate          *string `json:"creationDate,omitempty"`
	DatabaseName          *string `json:"databaseName,omitempty"`
	DeletionDate          *string `json:"deletionDate,omitempty"`
	EarliestRestoreDate   *string `json:"earliestRestoreDate,omitempty"`
	Edition               *string `json:"edition,omitempty"`
	ElasticPoolName       *string `json:"elasticPoolName,omitempty"`
	MaxSizeBytes          *string `json:"maxSizeBytes,omitempty"`
	ServiceLevelObjective *string `json:"serviceLevelObjective,omitempty"`
}
