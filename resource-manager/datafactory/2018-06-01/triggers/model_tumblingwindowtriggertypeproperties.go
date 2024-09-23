package triggers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TumblingWindowTriggerTypeProperties struct {
	Delay          *string                 `json:"delay,omitempty"`
	DependsOn      *[]DependencyReference  `json:"dependsOn,omitempty"`
	EndTime        *string                 `json:"endTime,omitempty"`
	Frequency      TumblingWindowFrequency `json:"frequency"`
	Interval       int64                   `json:"interval"`
	MaxConcurrency int64                   `json:"maxConcurrency"`
	RetryPolicy    *RetryPolicy            `json:"retryPolicy,omitempty"`
	StartTime      string                  `json:"startTime"`
}

func (o *TumblingWindowTriggerTypeProperties) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TumblingWindowTriggerTypeProperties) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *TumblingWindowTriggerTypeProperties) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TumblingWindowTriggerTypeProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}

var _ json.Unmarshaler = &TumblingWindowTriggerTypeProperties{}

func (s *TumblingWindowTriggerTypeProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Delay          *string                 `json:"delay,omitempty"`
		EndTime        *string                 `json:"endTime,omitempty"`
		Frequency      TumblingWindowFrequency `json:"frequency"`
		Interval       int64                   `json:"interval"`
		MaxConcurrency int64                   `json:"maxConcurrency"`
		RetryPolicy    *RetryPolicy            `json:"retryPolicy,omitempty"`
		StartTime      string                  `json:"startTime"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Delay = decoded.Delay
	s.EndTime = decoded.EndTime
	s.Frequency = decoded.Frequency
	s.Interval = decoded.Interval
	s.MaxConcurrency = decoded.MaxConcurrency
	s.RetryPolicy = decoded.RetryPolicy
	s.StartTime = decoded.StartTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TumblingWindowTriggerTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dependsOn"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DependsOn into list []json.RawMessage: %+v", err)
		}

		output := make([]DependencyReference, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDependencyReferenceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DependsOn' for 'TumblingWindowTriggerTypeProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DependsOn = &output
	}

	return nil
}
