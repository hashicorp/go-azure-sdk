package changedatacapturedestinations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureDestinationDetails interface {
	ChangeDataCaptureDestinationDetails() BaseChangeDataCaptureDestinationDetailsImpl
}

var _ ChangeDataCaptureDestinationDetails = BaseChangeDataCaptureDestinationDetailsImpl{}

type BaseChangeDataCaptureDestinationDetailsImpl struct {
	DestinationType ChangeDataCaptureDestinationType `json:"destinationType"`
	Protocol        ChangeDataCaptureProtocol        `json:"protocol"`
}

func (s BaseChangeDataCaptureDestinationDetailsImpl) ChangeDataCaptureDestinationDetails() BaseChangeDataCaptureDestinationDetailsImpl {
	return s
}

var _ ChangeDataCaptureDestinationDetails = RawChangeDataCaptureDestinationDetailsImpl{}

// RawChangeDataCaptureDestinationDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawChangeDataCaptureDestinationDetailsImpl struct {
	changeDataCaptureDestinationDetails BaseChangeDataCaptureDestinationDetailsImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawChangeDataCaptureDestinationDetailsImpl) ChangeDataCaptureDestinationDetails() BaseChangeDataCaptureDestinationDetailsImpl {
	return s.changeDataCaptureDestinationDetails
}

func (s RawChangeDataCaptureDestinationDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalChangeDataCaptureDestinationDetailsImplementation(input []byte) (ChangeDataCaptureDestinationDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChangeDataCaptureDestinationDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["destinationType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Fabric") {
		var out FabricDestinationDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FabricDestinationDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseChangeDataCaptureDestinationDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseChangeDataCaptureDestinationDetailsImpl: %+v", err)
	}

	return RawChangeDataCaptureDestinationDetailsImpl{
		changeDataCaptureDestinationDetails: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
