package databaseadvisors

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorProperties struct {
	AdvisorStatus                  *AdvisorStatus                  `json:"advisorStatus,omitempty"`
	AutoExecuteStatus              AutoExecuteStatus               `json:"autoExecuteStatus"`
	AutoExecuteStatusInheritedFrom *AutoExecuteStatusInheritedFrom `json:"autoExecuteStatusInheritedFrom,omitempty"`
	LastChecked                    *string                         `json:"lastChecked,omitempty"`
	RecommendationsStatus          *string                         `json:"recommendationsStatus,omitempty"`
	RecommendedActions             *[]RecommendedAction            `json:"recommendedActions,omitempty"`
}

func (o *AdvisorProperties) GetLastCheckedAsTime() (*time.Time, error) {
	if o.LastChecked == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastChecked, "2006-01-02T15:04:05Z07:00")
}

func (o *AdvisorProperties) SetLastCheckedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastChecked = &formatted
}
