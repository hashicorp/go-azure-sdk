package backupusagesummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupManagementUsage struct {
	CurrentValue  *int64      `json:"currentValue,omitempty"`
	Limit         *int64      `json:"limit,omitempty"`
	Name          *NameInfo   `json:"name,omitempty"`
	NextResetTime *string     `json:"nextResetTime,omitempty"`
	QuotaPeriod   *string     `json:"quotaPeriod,omitempty"`
	Unit          *UsagesUnit `json:"unit,omitempty"`
}
