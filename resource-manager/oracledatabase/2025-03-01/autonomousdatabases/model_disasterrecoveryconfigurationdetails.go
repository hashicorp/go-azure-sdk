package autonomousdatabases

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisasterRecoveryConfigurationDetails struct {
	DisasterRecoveryType           *DisasterRecoveryType `json:"disasterRecoveryType,omitempty"`
	IsReplicateAutomaticBackups    *bool                 `json:"isReplicateAutomaticBackups,omitempty"`
	IsSnapshotStandby              *bool                 `json:"isSnapshotStandby,omitempty"`
	TimeSnapshotStandbyEnabledTill *string               `json:"timeSnapshotStandbyEnabledTill,omitempty"`
}

func (o *DisasterRecoveryConfigurationDetails) GetTimeSnapshotStandbyEnabledTillAsTime() (*time.Time, error) {
	if o.TimeSnapshotStandbyEnabledTill == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeSnapshotStandbyEnabledTill, "2006-01-02T15:04:05Z07:00")
}

func (o *DisasterRecoveryConfigurationDetails) SetTimeSnapshotStandbyEnabledTillAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeSnapshotStandbyEnabledTill = &formatted
}
