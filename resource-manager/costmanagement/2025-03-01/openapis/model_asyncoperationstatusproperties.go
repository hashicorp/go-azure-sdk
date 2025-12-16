package openapis

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AsyncOperationStatusProperties struct {
	ReportURL          *BenefitUtilizationSummaryReportSchema `json:"reportUrl,omitempty"`
	SecondaryReportURL *BenefitUtilizationSummaryReportSchema `json:"secondaryReportUrl,omitempty"`
	ValidUntil         *string                                `json:"validUntil,omitempty"`
}

func (o *AsyncOperationStatusProperties) GetValidUntilAsTime() (*time.Time, error) {
	if o.ValidUntil == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidUntil, "2006-01-02T15:04:05Z07:00")
}

func (o *AsyncOperationStatusProperties) SetValidUntilAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidUntil = &formatted
}
