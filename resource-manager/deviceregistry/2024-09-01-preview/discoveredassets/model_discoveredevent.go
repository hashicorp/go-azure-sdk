package discoveredassets

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredEvent struct {
	EventConfiguration *string `json:"eventConfiguration,omitempty"`
	EventNotifier      string  `json:"eventNotifier"`
	LastUpdatedOn      *string `json:"lastUpdatedOn,omitempty"`
	Name               string  `json:"name"`
	Topic              *Topic  `json:"topic,omitempty"`
}

func (o *DiscoveredEvent) GetLastUpdatedOnAsTime() (*time.Time, error) {
	if o.LastUpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DiscoveredEvent) SetLastUpdatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedOn = &formatted
}
