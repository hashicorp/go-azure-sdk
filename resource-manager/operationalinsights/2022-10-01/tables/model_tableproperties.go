package tables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TableProperties struct {
	ArchiveRetentionInDays        *int64                 `json:"archiveRetentionInDays,omitempty"`
	LastPlanModifiedDate          *string                `json:"lastPlanModifiedDate,omitempty"`
	Plan                          *TablePlanEnum         `json:"plan,omitempty"`
	ProvisioningState             *ProvisioningStateEnum `json:"provisioningState,omitempty"`
	RestoredLogs                  *RestoredLogs          `json:"restoredLogs"`
	ResultStatistics              *ResultStatistics      `json:"resultStatistics"`
	RetentionInDays               *int64                 `json:"retentionInDays,omitempty"`
	RetentionInDaysAsDefault      *bool                  `json:"retentionInDaysAsDefault,omitempty"`
	Schema                        *Schema                `json:"schema"`
	SearchResults                 *SearchResults         `json:"searchResults"`
	TotalRetentionInDays          *int64                 `json:"totalRetentionInDays,omitempty"`
	TotalRetentionInDaysAsDefault *bool                  `json:"totalRetentionInDaysAsDefault,omitempty"`
}
