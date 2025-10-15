package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobInput interface {
	JobInput() BaseJobInputImpl
}

var _ JobInput = BaseJobInputImpl{}

type BaseJobInputImpl struct {
	Description  *string      `json:"description,omitempty"`
	JobInputType JobInputType `json:"jobInputType"`
}

func (s BaseJobInputImpl) JobInput() BaseJobInputImpl {
	return s
}

var _ JobInput = RawJobInputImpl{}

// RawJobInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobInputImpl struct {
	jobInput BaseJobInputImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawJobInputImpl) JobInput() BaseJobInputImpl {
	return s.jobInput
}

func UnmarshalJobInputImplementation(input []byte) (JobInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["jobInputType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "custom_model") {
		var out CustomModelJobInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomModelJobInput: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "literal") {
		var out LiteralJobInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LiteralJobInput: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "mlflow_model") {
		var out MLFlowModelJobInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MLFlowModelJobInput: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "mltable") {
		var out MLTableJobInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MLTableJobInput: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "triton_model") {
		var out TritonModelJobInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TritonModelJobInput: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "uri_file") {
		var out UriFileJobInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UriFileJobInput: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "uri_folder") {
		var out UriFolderJobInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UriFolderJobInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobInputImpl: %+v", err)
	}

	return RawJobInputImpl{
		jobInput: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
