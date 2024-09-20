package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventHandlerResult interface {
	AuthenticationEventHandlerResult() BaseAuthenticationEventHandlerResultImpl
}

var _ AuthenticationEventHandlerResult = BaseAuthenticationEventHandlerResultImpl{}

type BaseAuthenticationEventHandlerResultImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseAuthenticationEventHandlerResultImpl) AuthenticationEventHandlerResult() BaseAuthenticationEventHandlerResultImpl {
	return s
}

var _ AuthenticationEventHandlerResult = RawAuthenticationEventHandlerResultImpl{}

// RawAuthenticationEventHandlerResultImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationEventHandlerResultImpl struct {
	authenticationEventHandlerResult BaseAuthenticationEventHandlerResultImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawAuthenticationEventHandlerResultImpl) AuthenticationEventHandlerResult() BaseAuthenticationEventHandlerResultImpl {
	return s.authenticationEventHandlerResult
}

func UnmarshalAuthenticationEventHandlerResultImplementation(input []byte) (AuthenticationEventHandlerResult, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationEventHandlerResult into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customExtensionCalloutResult") {
		var out CustomExtensionCalloutResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomExtensionCalloutResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationEventHandlerResultImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationEventHandlerResultImpl: %+v", err)
	}

	return RawAuthenticationEventHandlerResultImpl{
		authenticationEventHandlerResult: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
