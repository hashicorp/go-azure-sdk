package encodings

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
	OdataType string `json:"@odata.type"`
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
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.JobInputAsset") {
		var out JobInputAsset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobInputAsset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.JobInputClip") {
		var out JobInputClip
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobInputClip: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.JobInputHttp") {
		var out JobInputHTTP
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobInputHTTP: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.JobInputSequence") {
		var out JobInputSequence
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobInputSequence: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.JobInputs") {
		var out JobInputs
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobInputs: %+v", err)
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
