package privateendpoint

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatus struct {
	EndTime    *string                     `json:"endTime,omitempty"`
	Error      *OperationStatusError       `json:"error,omitempty"`
	Id         *string                     `json:"id,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties OperationStatusExtendedInfo `json:"properties"`
	StartTime  *string                     `json:"startTime,omitempty"`
	Status     *OperationStatusValues      `json:"status,omitempty"`
}

func (o *OperationStatus) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *OperationStatus) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *OperationStatus) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *OperationStatus) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}

var _ json.Unmarshaler = &OperationStatus{}

func (s *OperationStatus) UnmarshalJSON(bytes []byte) error {
	type alias OperationStatus
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into OperationStatus: %+v", err)
	}

	s.EndTime = decoded.EndTime
	s.Error = decoded.Error
	s.Id = decoded.Id
	s.Name = decoded.Name
	s.StartTime = decoded.StartTime
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OperationStatus into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalOperationStatusExtendedInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'OperationStatus': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
