package assetsandassetfilters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrackBase interface {
	TrackBase() BaseTrackBaseImpl
}

var _ TrackBase = BaseTrackBaseImpl{}

type BaseTrackBaseImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BaseTrackBaseImpl) TrackBase() BaseTrackBaseImpl {
	return s
}

var _ TrackBase = RawTrackBaseImpl{}

// RawTrackBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTrackBaseImpl struct {
	trackBase BaseTrackBaseImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawTrackBaseImpl) TrackBase() BaseTrackBaseImpl {
	return s.trackBase
}

func UnmarshalTrackBaseImplementation(input []byte) (TrackBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TrackBase into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.AudioTrack") {
		var out AudioTrack
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AudioTrack: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.TextTrack") {
		var out TextTrack
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TextTrack: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.VideoTrack") {
		var out VideoTrack
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VideoTrack: %+v", err)
		}
		return out, nil
	}

	var parent BaseTrackBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTrackBaseImpl: %+v", err)
	}

	return RawTrackBaseImpl{
		trackBase: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
