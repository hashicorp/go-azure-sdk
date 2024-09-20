package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnTokenIssuanceStartHandler interface {
	OnTokenIssuanceStartHandler() BaseOnTokenIssuanceStartHandlerImpl
}

var _ OnTokenIssuanceStartHandler = BaseOnTokenIssuanceStartHandlerImpl{}

type BaseOnTokenIssuanceStartHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseOnTokenIssuanceStartHandlerImpl) OnTokenIssuanceStartHandler() BaseOnTokenIssuanceStartHandlerImpl {
	return s
}

var _ OnTokenIssuanceStartHandler = RawOnTokenIssuanceStartHandlerImpl{}

// RawOnTokenIssuanceStartHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnTokenIssuanceStartHandlerImpl struct {
	onTokenIssuanceStartHandler BaseOnTokenIssuanceStartHandlerImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawOnTokenIssuanceStartHandlerImpl) OnTokenIssuanceStartHandler() BaseOnTokenIssuanceStartHandlerImpl {
	return s.onTokenIssuanceStartHandler
}

func UnmarshalOnTokenIssuanceStartHandlerImplementation(input []byte) (OnTokenIssuanceStartHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnTokenIssuanceStartHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onTokenIssuanceStartCustomExtensionHandler") {
		var out OnTokenIssuanceStartCustomExtensionHandler
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnTokenIssuanceStartCustomExtensionHandler: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnTokenIssuanceStartHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnTokenIssuanceStartHandlerImpl: %+v", err)
	}

	return RawOnTokenIssuanceStartHandlerImpl{
		onTokenIssuanceStartHandler: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
