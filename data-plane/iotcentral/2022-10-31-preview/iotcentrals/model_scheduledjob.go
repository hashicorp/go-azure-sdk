package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledJob struct {
	Batch                 *JobBatch                 `json:"batch,omitempty"`
	CancellationThreshold *JobCancellationThreshold `json:"cancellationThreshold,omitempty"`
	Completed             *bool                     `json:"completed,omitempty"`
	Data                  []JobData                 `json:"data"`
	Description           *string                   `json:"description,omitempty"`
	DisplayName           *string                   `json:"displayName,omitempty"`
	Enabled               *bool                     `json:"enabled,omitempty"`
	Etag                  *string                   `json:"etag,omitempty"`
	Group                 string                    `json:"group"`
	Id                    *string                   `json:"id,omitempty"`
	Organizations         *[]string                 `json:"organizations,omitempty"`
	Schedule              JobSchedule               `json:"schedule"`
}

var _ json.Unmarshaler = &ScheduledJob{}

func (s *ScheduledJob) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Batch                 *JobBatch                 `json:"batch,omitempty"`
		CancellationThreshold *JobCancellationThreshold `json:"cancellationThreshold,omitempty"`
		Completed             *bool                     `json:"completed,omitempty"`
		Description           *string                   `json:"description,omitempty"`
		DisplayName           *string                   `json:"displayName,omitempty"`
		Enabled               *bool                     `json:"enabled,omitempty"`
		Etag                  *string                   `json:"etag,omitempty"`
		Group                 string                    `json:"group"`
		Id                    *string                   `json:"id,omitempty"`
		Organizations         *[]string                 `json:"organizations,omitempty"`
		Schedule              JobSchedule               `json:"schedule"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Batch = decoded.Batch
	s.CancellationThreshold = decoded.CancellationThreshold
	s.Completed = decoded.Completed
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Enabled = decoded.Enabled
	s.Etag = decoded.Etag
	s.Group = decoded.Group
	s.Id = decoded.Id
	s.Organizations = decoded.Organizations
	s.Schedule = decoded.Schedule

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ScheduledJob into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Data' for 'ScheduledJob': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Data = output
	}

	return nil
}
