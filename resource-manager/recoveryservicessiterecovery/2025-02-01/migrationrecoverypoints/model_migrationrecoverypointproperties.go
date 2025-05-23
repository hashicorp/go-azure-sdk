package migrationrecoverypoints

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationRecoveryPointProperties struct {
	RecoveryPointTime *string                     `json:"recoveryPointTime,omitempty"`
	RecoveryPointType *MigrationRecoveryPointType `json:"recoveryPointType,omitempty"`
}

func (o *MigrationRecoveryPointProperties) GetRecoveryPointTimeAsTime() (*time.Time, error) {
	if o.RecoveryPointTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RecoveryPointTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MigrationRecoveryPointProperties) SetRecoveryPointTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RecoveryPointTime = &formatted
}
