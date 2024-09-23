package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobOutput interface {
	JobOutput() BaseJobOutputImpl
}

var _ JobOutput = BaseJobOutputImpl{}

type BaseJobOutputImpl struct {
	EndTime        *string   `json:"endTime,omitempty"`
	Error          *JobError `json:"error,omitempty"`
	Label          *string   `json:"label,omitempty"`
	OdataType      string    `json:"@odata.type"`
	PresetOverride Preset    `json:"presetOverride"`
	Progress       *int64    `json:"progress,omitempty"`
	StartTime      *string   `json:"startTime,omitempty"`
	State          *JobState `json:"state,omitempty"`
}

func (s BaseJobOutputImpl) JobOutput() BaseJobOutputImpl {
	return s
}

var _ JobOutput = RawJobOutputImpl{}

// RawJobOutputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobOutputImpl struct {
	jobOutput BaseJobOutputImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawJobOutputImpl) JobOutput() BaseJobOutputImpl {
	return s.jobOutput
}

var _ json.Unmarshaler = &BaseJobOutputImpl{}

func (s *BaseJobOutputImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EndTime   *string   `json:"endTime,omitempty"`
		Error     *JobError `json:"error,omitempty"`
		Label     *string   `json:"label,omitempty"`
		OdataType string    `json:"@odata.type"`
		Progress  *int64    `json:"progress,omitempty"`
		StartTime *string   `json:"startTime,omitempty"`
		State     *JobState `json:"state,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EndTime = decoded.EndTime
	s.Error = decoded.Error
	s.Label = decoded.Label
	s.OdataType = decoded.OdataType
	s.Progress = decoded.Progress
	s.StartTime = decoded.StartTime
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseJobOutputImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["presetOverride"]; ok {
		impl, err := UnmarshalPresetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'PresetOverride' for 'BaseJobOutputImpl': %+v", err)
		}
		s.PresetOverride = impl
	}

	return nil
}

func UnmarshalJobOutputImplementation(input []byte) (JobOutput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobOutput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.JobOutputAsset") {
		var out JobOutputAsset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobOutputAsset: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobOutputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobOutputImpl: %+v", err)
	}

	return RawJobOutputImpl{
		jobOutput: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
