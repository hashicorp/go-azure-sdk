package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Preset interface {
	Preset() BasePresetImpl
}

var _ Preset = BasePresetImpl{}

type BasePresetImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BasePresetImpl) Preset() BasePresetImpl {
	return s
}

var _ Preset = RawPresetImpl{}

// RawPresetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPresetImpl struct {
	preset BasePresetImpl
	Type   string
	Values map[string]interface{}
}

func (s RawPresetImpl) Preset() BasePresetImpl {
	return s.preset
}

func UnmarshalPresetImplementation(input []byte) (Preset, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Preset into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.AudioAnalyzerPreset") {
		var out AudioAnalyzerPreset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AudioAnalyzerPreset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.BuiltInStandardEncoderPreset") {
		var out BuiltInStandardEncoderPreset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BuiltInStandardEncoderPreset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.FaceDetectorPreset") {
		var out FaceDetectorPreset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FaceDetectorPreset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.StandardEncoderPreset") {
		var out StandardEncoderPreset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StandardEncoderPreset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.VideoAnalyzerPreset") {
		var out VideoAnalyzerPreset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VideoAnalyzerPreset: %+v", err)
		}
		return out, nil
	}

	var parent BasePresetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePresetImpl: %+v", err)
	}

	return RawPresetImpl{
		preset: parent,
		Type:   value,
		Values: temp,
	}, nil

}
