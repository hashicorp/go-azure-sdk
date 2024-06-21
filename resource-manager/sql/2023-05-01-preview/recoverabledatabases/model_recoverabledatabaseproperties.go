package recoverabledatabases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableDatabaseProperties struct {
	Edition                 *string                 `json:"edition,omitempty"`
	ElasticPoolName         *string                 `json:"elasticPoolName,omitempty"`
	Keys                    *map[string]DatabaseKey `json:"keys,omitempty"`
	LastAvailableBackupDate *string                 `json:"lastAvailableBackupDate,omitempty"`
	ServiceLevelObjective   *string                 `json:"serviceLevelObjective,omitempty"`
}
