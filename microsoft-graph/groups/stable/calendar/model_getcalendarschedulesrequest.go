package calendar

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCalendarSchedulesRequest struct {
	AvailabilityViewInterval nullable.Type[int64]     `json:"AvailabilityViewInterval,omitempty"`
	EndTime                  *stable.DateTimeTimeZone `json:"EndTime,omitempty"`
	Schedules                *[]string                `json:"Schedules,omitempty"`
	StartTime                *stable.DateTimeTimeZone `json:"StartTime,omitempty"`
}
