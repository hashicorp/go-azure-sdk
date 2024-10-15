package appplatform

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobTriggerConfig interface {
	JobTriggerConfig() BaseJobTriggerConfigImpl
}

var _ JobTriggerConfig = BaseJobTriggerConfigImpl{}

type BaseJobTriggerConfigImpl struct {
	TriggerType TriggerType `json:"triggerType"`
}

func (s BaseJobTriggerConfigImpl) JobTriggerConfig() BaseJobTriggerConfigImpl {
	return s
}

var _ JobTriggerConfig = RawJobTriggerConfigImpl{}

// RawJobTriggerConfigImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobTriggerConfigImpl struct {
	jobTriggerConfig BaseJobTriggerConfigImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawJobTriggerConfigImpl) JobTriggerConfig() BaseJobTriggerConfigImpl {
	return s.jobTriggerConfig
}

func UnmarshalJobTriggerConfigImplementation(input []byte) (JobTriggerConfig, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobTriggerConfig into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["triggerType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Manual") {
		var out ManualJobTriggerConfig
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManualJobTriggerConfig: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobTriggerConfigImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobTriggerConfigImpl: %+v", err)
	}

	return RawJobTriggerConfigImpl{
		jobTriggerConfig: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
