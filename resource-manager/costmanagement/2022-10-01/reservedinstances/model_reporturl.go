package reservedinstances

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportURL struct {
	ReportUrl  *ReservationReportSchema `json:"reportUrl,omitempty"`
	ValidUntil *string                  `json:"validUntil,omitempty"`
}

func (o *ReportURL) GetValidUntilAsTime() (*time.Time, error) {
	if o.ValidUntil == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidUntil, "2006-01-02T15:04:05Z07:00")
}

func (o *ReportURL) SetValidUntilAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidUntil = &formatted
}
