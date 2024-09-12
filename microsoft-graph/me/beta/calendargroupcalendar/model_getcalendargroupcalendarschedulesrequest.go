package calendargroupcalendar

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCalendarGroupCalendarSchedulesRequest struct {
	AvailabilityViewInterval nullable.Type[int64]   `json:"AvailabilityViewInterval,omitempty"`
	EndTime                  *beta.DateTimeTimeZone `json:"EndTime,omitempty"`
	Schedules                *[]string              `json:"Schedules,omitempty"`
	StartTime                *beta.DateTimeTimeZone `json:"StartTime,omitempty"`
}
