package operationalizationclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeStartStopSchedule struct {
	Action             *ComputePowerAction `json:"action,omitempty"`
	Cron               TriggerBase         `json:"cron"`
	Id                 *string             `json:"id,omitempty"`
	ProvisioningStatus *ProvisioningStatus `json:"provisioningStatus,omitempty"`
	Recurrence         TriggerBase         `json:"recurrence"`
	Schedule           *ScheduleBase       `json:"schedule"`
	Status             *ScheduleStatus     `json:"status,omitempty"`
	TriggerType        *TriggerType        `json:"triggerType,omitempty"`
}

var _ json.Unmarshaler = &ComputeStartStopSchedule{}

func (s *ComputeStartStopSchedule) UnmarshalJSON(bytes []byte) error {
	type alias ComputeStartStopSchedule
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ComputeStartStopSchedule: %+v", err)
	}

	s.Action = decoded.Action
	s.Id = decoded.Id
	s.ProvisioningStatus = decoded.ProvisioningStatus
	s.Schedule = decoded.Schedule
	s.Status = decoded.Status
	s.TriggerType = decoded.TriggerType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ComputeStartStopSchedule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["cron"]; ok {
		impl, err := unmarshalTriggerBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Cron' for 'ComputeStartStopSchedule': %+v", err)
		}
		s.Cron = impl
	}

	if v, ok := temp["recurrence"]; ok {
		impl, err := unmarshalTriggerBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Recurrence' for 'ComputeStartStopSchedule': %+v", err)
		}
		s.Recurrence = impl
	}
	return nil
}
