package topology

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopologyResourceProperties struct {
	CalculatedDateTime *string                   `json:"calculatedDateTime,omitempty"`
	TopologyResources  *[]TopologySingleResource `json:"topologyResources,omitempty"`
}

func (o *TopologyResourceProperties) GetCalculatedDateTimeAsTime() (*time.Time, error) {
	if o.CalculatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CalculatedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TopologyResourceProperties) SetCalculatedDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CalculatedDateTime = &formatted
}
