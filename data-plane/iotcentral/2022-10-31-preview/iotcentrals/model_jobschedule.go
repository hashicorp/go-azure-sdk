package iotcentrals

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobSchedule struct {
	End        JobScheduleEnd `json:"end"`
	Recurrence *JobRecurrence `json:"recurrence,omitempty"`
	Start      string         `json:"start"`
}

func (o *JobSchedule) GetStartAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.Start, "2006-01-02T15:04:05Z07:00")
}

func (o *JobSchedule) SetStartAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Start = formatted
}

var _ json.Unmarshaler = &JobSchedule{}

func (s *JobSchedule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Recurrence *JobRecurrence `json:"recurrence,omitempty"`
		Start      string         `json:"start"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Recurrence = decoded.Recurrence
	s.Start = decoded.Start

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling JobSchedule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["end"]; ok {
		impl, err := UnmarshalJobScheduleEndImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'End' for 'JobSchedule': %+v", err)
		}
		s.End = impl
	}

	return nil
}
