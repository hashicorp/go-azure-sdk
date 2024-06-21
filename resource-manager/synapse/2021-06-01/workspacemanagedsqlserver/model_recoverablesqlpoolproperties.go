package workspacemanagedsqlserver

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableSqlPoolProperties struct {
	Edition                 *string `json:"edition,omitempty"`
	ElasticPoolName         *string `json:"elasticPoolName,omitempty"`
	LastAvailableBackupDate *string `json:"lastAvailableBackupDate,omitempty"`
	ServiceLevelObjective   *string `json:"serviceLevelObjective,omitempty"`
}
