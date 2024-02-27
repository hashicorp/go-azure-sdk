package sqlpoolsusages

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolUsage struct {
	CurrentValue  *float64 `json:"currentValue,omitempty"`
	DisplayName   *string  `json:"displayName,omitempty"`
	Limit         *float64 `json:"limit,omitempty"`
	Name          *string  `json:"name,omitempty"`
	NextResetTime *string  `json:"nextResetTime,omitempty"`
	ResourceName  *string  `json:"resourceName,omitempty"`
	Unit          *string  `json:"unit,omitempty"`
}

func (o *SqlPoolUsage) GetNextResetTimeAsTime() (*time.Time, error) {
	if o.NextResetTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NextResetTime, "2006-01-02T15:04:05Z07:00")
}
