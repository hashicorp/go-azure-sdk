package clusters

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservationProperties struct {
	LastSkuUpdate *string `json:"lastSkuUpdate,omitempty"`
	MinCapacity   *int64  `json:"minCapacity,omitempty"`
}

func (o *CapacityReservationProperties) GetLastSkuUpdateAsTime() (*time.Time, error) {
	if o.LastSkuUpdate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastSkuUpdate, "2006-01-02T15:04:05Z07:00")
}

func (o *CapacityReservationProperties) SetLastSkuUpdateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastSkuUpdate = &formatted
}
