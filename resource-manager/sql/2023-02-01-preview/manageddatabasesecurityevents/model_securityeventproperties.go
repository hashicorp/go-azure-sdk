package manageddatabasesecurityevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventProperties struct {
	ApplicationName                               *string                                        `json:"applicationName,omitempty"`
	ClientIP                                      *string                                        `json:"clientIp,omitempty"`
	Database                                      *string                                        `json:"database,omitempty"`
	EventTime                                     *string                                        `json:"eventTime,omitempty"`
	PrincipalName                                 *string                                        `json:"principalName,omitempty"`
	SecurityEventSqlInjectionAdditionalProperties *SecurityEventSqlInjectionAdditionalProperties `json:"securityEventSqlInjectionAdditionalProperties,omitempty"`
	SecurityEventType                             *SecurityEventType                             `json:"securityEventType,omitempty"`
	Server                                        *string                                        `json:"server,omitempty"`
	Subscription                                  *string                                        `json:"subscription,omitempty"`
}

func (o *SecurityEventProperties) GetEventTimeAsTime() (*time.Time, error) {
	if o.EventTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EventTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SecurityEventProperties) SetEventTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EventTime = &formatted
}
