package networkclouds

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretRotationStatus struct {
	ExpirePeriodDays       *int64                  `json:"expirePeriodDays,omitempty"`
	LastRotationTime       *string                 `json:"lastRotationTime,omitempty"`
	RotationPeriodDays     *int64                  `json:"rotationPeriodDays,omitempty"`
	SecretArchiveReference *SecretArchiveReference `json:"secretArchiveReference,omitempty"`
	SecretType             *string                 `json:"secretType,omitempty"`
}

func (o *SecretRotationStatus) GetLastRotationTimeAsTime() (*time.Time, error) {
	if o.LastRotationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRotationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SecretRotationStatus) SetLastRotationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRotationTime = &formatted
}
