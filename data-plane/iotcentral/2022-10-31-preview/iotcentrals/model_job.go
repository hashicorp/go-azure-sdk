package iotcentrals

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Job struct {
	Batch                 *JobBatch                 `json:"batch,omitempty"`
	CancellationThreshold *JobCancellationThreshold `json:"cancellationThreshold,omitempty"`
	Data                  []JobData                 `json:"data"`
	Description           *string                   `json:"description,omitempty"`
	DisplayName           *string                   `json:"displayName,omitempty"`
	End                   *string                   `json:"end,omitempty"`
	Group                 string                    `json:"group"`
	Id                    *string                   `json:"id,omitempty"`
	Organizations         *[]string                 `json:"organizations,omitempty"`
	Progress              *JobProgress              `json:"progress,omitempty"`
	ScheduledJobId        *string                   `json:"scheduledJobId,omitempty"`
	Start                 *string                   `json:"start,omitempty"`
	Status                *string                   `json:"status,omitempty"`
}

func (o *Job) GetEndAsTime() (*time.Time, error) {
	if o.End == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.End, "2006-01-02T15:04:05Z07:00")
}

func (o *Job) SetEndAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.End = &formatted
}

func (o *Job) GetStartAsTime() (*time.Time, error) {
	if o.Start == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Start, "2006-01-02T15:04:05Z07:00")
}

func (o *Job) SetStartAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Start = &formatted
}

var _ json.Unmarshaler = &Job{}

func (s *Job) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Batch                 *JobBatch                 `json:"batch,omitempty"`
		CancellationThreshold *JobCancellationThreshold `json:"cancellationThreshold,omitempty"`
		Description           *string                   `json:"description,omitempty"`
		DisplayName           *string                   `json:"displayName,omitempty"`
		End                   *string                   `json:"end,omitempty"`
		Group                 string                    `json:"group"`
		Id                    *string                   `json:"id,omitempty"`
		Organizations         *[]string                 `json:"organizations,omitempty"`
		Progress              *JobProgress              `json:"progress,omitempty"`
		ScheduledJobId        *string                   `json:"scheduledJobId,omitempty"`
		Start                 *string                   `json:"start,omitempty"`
		Status                *string                   `json:"status,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Batch = decoded.Batch
	s.CancellationThreshold = decoded.CancellationThreshold
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.End = decoded.End
	s.Group = decoded.Group
	s.Id = decoded.Id
	s.Organizations = decoded.Organizations
	s.Progress = decoded.Progress
	s.ScheduledJobId = decoded.ScheduledJobId
	s.Start = decoded.Start
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Job into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["data"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Data into list []json.RawMessage: %+v", err)
		}

		output := make([]JobData, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalJobDataImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Data' for 'Job': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Data = output
	}

	return nil
}
