package replicationappliances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplianceSpecificDetails interface {
	ApplianceSpecificDetails() BaseApplianceSpecificDetailsImpl
}

var _ ApplianceSpecificDetails = BaseApplianceSpecificDetailsImpl{}

type BaseApplianceSpecificDetailsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseApplianceSpecificDetailsImpl) ApplianceSpecificDetails() BaseApplianceSpecificDetailsImpl {
	return s
}

var _ ApplianceSpecificDetails = RawApplianceSpecificDetailsImpl{}

// RawApplianceSpecificDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawApplianceSpecificDetailsImpl struct {
	applianceSpecificDetails BaseApplianceSpecificDetailsImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawApplianceSpecificDetailsImpl) ApplianceSpecificDetails() BaseApplianceSpecificDetailsImpl {
	return s.applianceSpecificDetails
}

func UnmarshalApplianceSpecificDetailsImplementation(input []byte) (ApplianceSpecificDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplianceSpecificDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "InMageRcm") {
		var out InMageRcmApplianceSpecificDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageRcmApplianceSpecificDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseApplianceSpecificDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseApplianceSpecificDetailsImpl: %+v", err)
	}

	return RawApplianceSpecificDetailsImpl{
		applianceSpecificDetails: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
