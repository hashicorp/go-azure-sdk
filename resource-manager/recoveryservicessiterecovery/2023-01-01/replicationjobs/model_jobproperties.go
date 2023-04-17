package replicationjobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobProperties struct {
	ActivityId         *string            `json:"activityId,omitempty"`
	AllowedActions     *[]string          `json:"allowedActions,omitempty"`
	CustomDetails      *JobDetails        `json:"customDetails,omitempty"`
	EndTime            *string            `json:"endTime,omitempty"`
	Errors             *[]JobErrorDetails `json:"errors,omitempty"`
	FriendlyName       *string            `json:"friendlyName,omitempty"`
	ScenarioName       *string            `json:"scenarioName,omitempty"`
	StartTime          *string            `json:"startTime,omitempty"`
	State              *string            `json:"state,omitempty"`
	StateDescription   *string            `json:"stateDescription,omitempty"`
	TargetInstanceType *string            `json:"targetInstanceType,omitempty"`
	TargetObjectId     *string            `json:"targetObjectId,omitempty"`
	TargetObjectName   *string            `json:"targetObjectName,omitempty"`
	Tasks              *[]ASRTask         `json:"tasks,omitempty"`
}

func (o *JobProperties) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobProperties) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *JobProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}

var _ json.Unmarshaler = &JobProperties{}

func (s *JobProperties) UnmarshalJSON(bytes []byte) error {
	type alias JobProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into JobProperties: %+v", err)
	}

	s.ActivityId = decoded.ActivityId
	s.AllowedActions = decoded.AllowedActions
	s.EndTime = decoded.EndTime
	s.Errors = decoded.Errors
	s.FriendlyName = decoded.FriendlyName
	s.ScenarioName = decoded.ScenarioName
	s.StartTime = decoded.StartTime
	s.State = decoded.State
	s.StateDescription = decoded.StateDescription
	s.TargetInstanceType = decoded.TargetInstanceType
	s.TargetObjectId = decoded.TargetObjectId
	s.TargetObjectName = decoded.TargetObjectName
	s.Tasks = decoded.Tasks

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling JobProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["customDetails"]; ok {
		impl, err := unmarshalJobDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CustomDetails' for 'JobProperties': %+v", err)
		}
		s.CustomDetails = &impl
	}
	return nil
}
