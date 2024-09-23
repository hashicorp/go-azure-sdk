package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClipTime interface {
	ClipTime() BaseClipTimeImpl
}

var _ ClipTime = BaseClipTimeImpl{}

type BaseClipTimeImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BaseClipTimeImpl) ClipTime() BaseClipTimeImpl {
	return s
}

var _ ClipTime = RawClipTimeImpl{}

// RawClipTimeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawClipTimeImpl struct {
	clipTime BaseClipTimeImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawClipTimeImpl) ClipTime() BaseClipTimeImpl {
	return s.clipTime
}

func UnmarshalClipTimeImplementation(input []byte) (ClipTime, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ClipTime into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.AbsoluteClipTime") {
		var out AbsoluteClipTime
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AbsoluteClipTime: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.UtcClipTime") {
		var out UtcClipTime
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UtcClipTime: %+v", err)
		}
		return out, nil
	}

	var parent BaseClipTimeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseClipTimeImpl: %+v", err)
	}

	return RawClipTimeImpl{
		clipTime: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
