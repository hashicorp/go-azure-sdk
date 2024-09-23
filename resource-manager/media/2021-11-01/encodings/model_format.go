package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Format interface {
	Format() BaseFormatImpl
}

var _ Format = BaseFormatImpl{}

type BaseFormatImpl struct {
	FilenamePattern string `json:"filenamePattern"`
	OdataType       string `json:"@odata.type"`
}

func (s BaseFormatImpl) Format() BaseFormatImpl {
	return s
}

var _ Format = RawFormatImpl{}

// RawFormatImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawFormatImpl struct {
	format BaseFormatImpl
	Type   string
	Values map[string]interface{}
}

func (s RawFormatImpl) Format() BaseFormatImpl {
	return s.format
}

func UnmarshalFormatImplementation(input []byte) (Format, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Format into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.ImageFormat") {
		var out ImageFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImageFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.JpgFormat") {
		var out JpgFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JpgFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.Mp4Format") {
		var out Mp4Format
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Mp4Format: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.MultiBitrateFormat") {
		var out MultiBitrateFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MultiBitrateFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.PngFormat") {
		var out PngFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PngFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.TransportStreamFormat") {
		var out TransportStreamFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TransportStreamFormat: %+v", err)
		}
		return out, nil
	}

	var parent BaseFormatImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFormatImpl: %+v", err)
	}

	return RawFormatImpl{
		format: parent,
		Type:   value,
		Values: temp,
	}, nil

}
