package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupSchedule struct {
	FrequencyInterval     int64         `json:"frequencyInterval"`
	FrequencyUnit         FrequencyUnit `json:"frequencyUnit"`
	KeepAtLeastOneBackup  bool          `json:"keepAtLeastOneBackup"`
	LastExecutionTime     *string       `json:"lastExecutionTime,omitempty"`
	RetentionPeriodInDays int64         `json:"retentionPeriodInDays"`
	StartTime             *string       `json:"startTime,omitempty"`
}
