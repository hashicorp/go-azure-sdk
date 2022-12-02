package backupprotecteditems

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureIaaSVMProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `json:"oldestRecoveryPoint,omitempty"`
	PolicyInconsistent  *bool   `json:"policyInconsistent,omitempty"`
	RecoveryPointCount  *int64  `json:"recoveryPointCount,omitempty"`
}

func (o *AzureIaaSVMProtectedItemExtendedInfo) GetOldestRecoveryPointAsTime() (*time.Time, error) {
	if o.OldestRecoveryPoint == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.OldestRecoveryPoint, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureIaaSVMProtectedItemExtendedInfo) SetOldestRecoveryPointAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.OldestRecoveryPoint = &formatted
}
