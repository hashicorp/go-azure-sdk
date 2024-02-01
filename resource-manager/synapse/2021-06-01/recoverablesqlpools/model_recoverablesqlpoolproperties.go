package recoverablesqlpools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableSqlPoolProperties struct {
	Edition                 *string `json:"edition,omitempty"`
	ElasticPoolName         *string `json:"elasticPoolName,omitempty"`
	LastAvailableBackupDate *string `json:"lastAvailableBackupDate,omitempty"`
	ServiceLevelObjective   *string `json:"serviceLevelObjective,omitempty"`
}

func (o *RecoverableSqlPoolProperties) GetLastAvailableBackupDateAsTime() (*time.Time, error) {
	if o.LastAvailableBackupDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastAvailableBackupDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RecoverableSqlPoolProperties) SetLastAvailableBackupDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastAvailableBackupDate = &formatted
}
