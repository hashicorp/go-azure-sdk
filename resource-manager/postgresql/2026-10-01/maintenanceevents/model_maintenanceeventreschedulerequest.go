package maintenanceevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceEventRescheduleRequest struct {
	PostponeToDateTime string `json:"postponeToDateTime"`
}

func (o *MaintenanceEventRescheduleRequest) GetPostponeToDateTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.PostponeToDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventRescheduleRequest) SetPostponeToDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PostponeToDateTime = formatted
}
