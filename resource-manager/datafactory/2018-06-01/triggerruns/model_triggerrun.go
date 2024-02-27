package triggerruns

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerRun struct {
	DependencyStatus    *map[string]interface{} `json:"dependencyStatus,omitempty"`
	Message             *string                 `json:"message,omitempty"`
	Properties          *map[string]string      `json:"properties,omitempty"`
	RunDimension        *map[string]string      `json:"runDimension,omitempty"`
	Status              *TriggerRunStatus       `json:"status,omitempty"`
	TriggerName         *string                 `json:"triggerName,omitempty"`
	TriggerRunId        *string                 `json:"triggerRunId,omitempty"`
	TriggerRunTimestamp *string                 `json:"triggerRunTimestamp,omitempty"`
	TriggerType         *string                 `json:"triggerType,omitempty"`
	TriggeredPipelines  *map[string]string      `json:"triggeredPipelines,omitempty"`
}

func (o *TriggerRun) GetTriggerRunTimestampAsTime() (*time.Time, error) {
	if o.TriggerRunTimestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TriggerRunTimestamp, "2006-01-02T15:04:05Z07:00")
}
