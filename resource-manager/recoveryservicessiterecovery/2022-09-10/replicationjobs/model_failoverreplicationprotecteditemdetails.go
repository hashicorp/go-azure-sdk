package replicationjobs

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailoverReplicationProtectedItemDetails struct {
	FriendlyName            *string `json:"friendlyName,omitempty"`
	Name                    *string `json:"name,omitempty"`
	NetworkConnectionStatus *string `json:"networkConnectionStatus,omitempty"`
	NetworkFriendlyName     *string `json:"networkFriendlyName,omitempty"`
	RecoveryPointId         *string `json:"recoveryPointId,omitempty"`
	RecoveryPointTime       *string `json:"recoveryPointTime,omitempty"`
	Subnet                  *string `json:"subnet,omitempty"`
	TestVmFriendlyName      *string `json:"testVmFriendlyName,omitempty"`
	TestVmName              *string `json:"testVmName,omitempty"`
}

func (o *FailoverReplicationProtectedItemDetails) GetRecoveryPointTimeAsTime() (*time.Time, error) {
	if o.RecoveryPointTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RecoveryPointTime, "2006-01-02T15:04:05Z07:00")
}

func (o *FailoverReplicationProtectedItemDetails) SetRecoveryPointTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RecoveryPointTime = &formatted
}
