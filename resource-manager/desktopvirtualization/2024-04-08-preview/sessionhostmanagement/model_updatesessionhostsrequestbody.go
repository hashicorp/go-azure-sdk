package sessionhostmanagement

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSessionHostsRequestBody struct {
	ScheduledDateTime     *string                                     `json:"scheduledDateTime,omitempty"`
	ScheduledDateTimeZone *string                                     `json:"scheduledDateTimeZone,omitempty"`
	Update                *HostPoolUpdateConfigurationPatchProperties `json:"update,omitempty"`
}

func (o *UpdateSessionHostsRequestBody) GetScheduledDateTimeAsTime() (*time.Time, error) {
	if o.ScheduledDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScheduledDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *UpdateSessionHostsRequestBody) SetScheduledDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScheduledDateTime = &formatted
}
