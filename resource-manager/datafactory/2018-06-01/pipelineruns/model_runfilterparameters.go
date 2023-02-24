package pipelineruns

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunFilterParameters struct {
	ContinuationToken *string            `json:"continuationToken,omitempty"`
	Filters           *[]RunQueryFilter  `json:"filters,omitempty"`
	LastUpdatedAfter  string             `json:"lastUpdatedAfter"`
	LastUpdatedBefore string             `json:"lastUpdatedBefore"`
	OrderBy           *[]RunQueryOrderBy `json:"orderBy,omitempty"`
}

func (o *RunFilterParameters) GetLastUpdatedAfterAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastUpdatedAfter, "2006-01-02T15:04:05Z07:00")
}

func (o *RunFilterParameters) SetLastUpdatedAfterAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedAfter = formatted
}

func (o *RunFilterParameters) GetLastUpdatedBeforeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastUpdatedBefore, "2006-01-02T15:04:05Z07:00")
}

func (o *RunFilterParameters) SetLastUpdatedBeforeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedBefore = formatted
}
