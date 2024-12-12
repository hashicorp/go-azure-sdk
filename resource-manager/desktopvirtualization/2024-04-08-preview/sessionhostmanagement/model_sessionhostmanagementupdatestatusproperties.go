package sessionhostmanagement

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionHostManagementUpdateStatusProperties struct {
	CorrelationId         *string                                 `json:"correlationId,omitempty"`
	Progress              *SessionHostManagementOperationProgress `json:"progress,omitempty"`
	ScheduledDateTime     *string                                 `json:"scheduledDateTime,omitempty"`
	SessionHostManagement *SessionHostManagement                  `json:"sessionHostManagement,omitempty"`
}

func (o *SessionHostManagementUpdateStatusProperties) GetScheduledDateTimeAsTime() (*time.Time, error) {
	if o.ScheduledDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScheduledDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SessionHostManagementUpdateStatusProperties) SetScheduledDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScheduledDateTime = &formatted
}
