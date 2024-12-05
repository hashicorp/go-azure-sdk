package endpoint

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointModelSkuProperties struct {
	Capacity        *EndpointModelSkuCapacityProperties    `json:"capacity,omitempty"`
	ConnectionIds   *[]string                              `json:"connectionIds,omitempty"`
	DeprecationDate *string                                `json:"deprecationDate,omitempty"`
	Name            *string                                `json:"name,omitempty"`
	RateLimits      *[]EndpointModelSkuRateLimitProperties `json:"rateLimits,omitempty"`
	UsageName       *string                                `json:"usageName,omitempty"`
}

func (o *EndpointModelSkuProperties) GetDeprecationDateAsTime() (*time.Time, error) {
	if o.DeprecationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeprecationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EndpointModelSkuProperties) SetDeprecationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeprecationDate = &formatted
}
