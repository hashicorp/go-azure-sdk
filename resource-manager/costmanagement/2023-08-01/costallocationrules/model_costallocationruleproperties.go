package costallocationrules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostAllocationRuleProperties struct {
	CreatedDate *string                   `json:"createdDate,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Details     CostAllocationRuleDetails `json:"details"`
	Status      RuleStatus                `json:"status"`
	UpdatedDate *string                   `json:"updatedDate,omitempty"`
}

func (o *CostAllocationRuleProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *CostAllocationRuleProperties) GetUpdatedDateAsTime() (*time.Time, error) {
	if o.UpdatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UpdatedDate, "2006-01-02T15:04:05Z07:00")
}
