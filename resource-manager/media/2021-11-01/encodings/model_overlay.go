package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Overlay interface {
	Overlay() BaseOverlayImpl
}

var _ Overlay = BaseOverlayImpl{}

type BaseOverlayImpl struct {
	AudioGainLevel  *float64 `json:"audioGainLevel,omitempty"`
	End             *string  `json:"end,omitempty"`
	FadeInDuration  *string  `json:"fadeInDuration,omitempty"`
	FadeOutDuration *string  `json:"fadeOutDuration,omitempty"`
	InputLabel      string   `json:"inputLabel"`
	OdataType       string   `json:"@odata.type"`
	Start           *string  `json:"start,omitempty"`
}

func (s BaseOverlayImpl) Overlay() BaseOverlayImpl {
	return s
}

var _ Overlay = RawOverlayImpl{}

// RawOverlayImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOverlayImpl struct {
	overlay BaseOverlayImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawOverlayImpl) Overlay() BaseOverlayImpl {
	return s.overlay
}

func UnmarshalOverlayImplementation(input []byte) (Overlay, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Overlay into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.AudioOverlay") {
		var out AudioOverlay
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AudioOverlay: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.VideoOverlay") {
		var out VideoOverlay
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VideoOverlay: %+v", err)
		}
		return out, nil
	}

	var parent BaseOverlayImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOverlayImpl: %+v", err)
	}

	return RawOverlayImpl{
		overlay: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
