package replicationjobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ASRTask struct {
	AllowedActions         *[]string          `json:"allowedActions,omitempty"`
	CustomDetails          TaskTypeDetails    `json:"customDetails"`
	EndTime                *string            `json:"endTime,omitempty"`
	Errors                 *[]JobErrorDetails `json:"errors,omitempty"`
	FriendlyName           *string            `json:"friendlyName,omitempty"`
	GroupTaskCustomDetails GroupTaskDetails   `json:"groupTaskCustomDetails"`
	Name                   *string            `json:"name,omitempty"`
	StartTime              *string            `json:"startTime,omitempty"`
	State                  *string            `json:"state,omitempty"`
	StateDescription       *string            `json:"stateDescription,omitempty"`
	TaskId                 *string            `json:"taskId,omitempty"`
	TaskType               *string            `json:"taskType,omitempty"`
}

func (o *ASRTask) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ASRTask) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *ASRTask) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ASRTask) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}

var _ json.Unmarshaler = &ASRTask{}

func (s *ASRTask) UnmarshalJSON(bytes []byte) error {
	type alias ASRTask
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ASRTask: %+v", err)
	}

	s.AllowedActions = decoded.AllowedActions
	s.EndTime = decoded.EndTime
	s.Errors = decoded.Errors
	s.FriendlyName = decoded.FriendlyName
	s.Name = decoded.Name
	s.StartTime = decoded.StartTime
	s.State = decoded.State
	s.StateDescription = decoded.StateDescription
	s.TaskId = decoded.TaskId
	s.TaskType = decoded.TaskType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ASRTask into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["customDetails"]; ok {
		impl, err := unmarshalTaskTypeDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CustomDetails' for 'ASRTask': %+v", err)
		}
		s.CustomDetails = impl
	}

	if v, ok := temp["groupTaskCustomDetails"]; ok {
		impl, err := unmarshalGroupTaskDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GroupTaskCustomDetails' for 'ASRTask': %+v", err)
		}
		s.GroupTaskCustomDetails = impl
	}
	return nil
}
