package costdetails

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostDetailsOperationResults struct {
	Error     *ErrorDetails          `json:"error"`
	Id        *string                `json:"id,omitempty"`
	Manifest  *ReportManifest        `json:"manifest"`
	Name      *string                `json:"name,omitempty"`
	Status    *CostDetailsStatusType `json:"status,omitempty"`
	Type      *string                `json:"type,omitempty"`
	ValidTill *string                `json:"validTill,omitempty"`
}

func (o *CostDetailsOperationResults) GetValidTillAsTime() (*time.Time, error) {
	if o.ValidTill == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidTill, "2006-01-02T15:04:05Z07:00")
}

func (o *CostDetailsOperationResults) SetValidTillAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidTill = &formatted
}
