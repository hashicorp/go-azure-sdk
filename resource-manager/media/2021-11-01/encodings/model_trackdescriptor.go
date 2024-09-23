package encodings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrackDescriptor interface {
	TrackDescriptor() BaseTrackDescriptorImpl
}

var _ TrackDescriptor = BaseTrackDescriptorImpl{}

type BaseTrackDescriptorImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BaseTrackDescriptorImpl) TrackDescriptor() BaseTrackDescriptorImpl {
	return s
}

var _ TrackDescriptor = RawTrackDescriptorImpl{}

// RawTrackDescriptorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTrackDescriptorImpl struct {
	trackDescriptor BaseTrackDescriptorImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawTrackDescriptorImpl) TrackDescriptor() BaseTrackDescriptorImpl {
	return s.trackDescriptor
}

func UnmarshalTrackDescriptorImplementation(input []byte) (TrackDescriptor, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TrackDescriptor into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.AudioTrackDescriptor") {
		var out AudioTrackDescriptor
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AudioTrackDescriptor: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.SelectAudioTrackByAttribute") {
		var out SelectAudioTrackByAttribute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelectAudioTrackByAttribute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.SelectAudioTrackById") {
		var out SelectAudioTrackById
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelectAudioTrackById: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.SelectVideoTrackByAttribute") {
		var out SelectVideoTrackByAttribute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelectVideoTrackByAttribute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.SelectVideoTrackById") {
		var out SelectVideoTrackById
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelectVideoTrackById: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.VideoTrackDescriptor") {
		var out VideoTrackDescriptor
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VideoTrackDescriptor: %+v", err)
		}
		return out, nil
	}

	var parent BaseTrackDescriptorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTrackDescriptorImpl: %+v", err)
	}

	return RawTrackDescriptorImpl{
		trackDescriptor: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
