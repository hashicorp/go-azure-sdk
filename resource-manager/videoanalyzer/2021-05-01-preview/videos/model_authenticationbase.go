package videos

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationBase interface {
	AuthenticationBase() BaseAuthenticationBaseImpl
}

var _ AuthenticationBase = BaseAuthenticationBaseImpl{}

type BaseAuthenticationBaseImpl struct {
	Type string `json:"@type"`
}

func (s BaseAuthenticationBaseImpl) AuthenticationBase() BaseAuthenticationBaseImpl {
	return s
}

var _ AuthenticationBase = RawAuthenticationBaseImpl{}

// RawAuthenticationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationBaseImpl struct {
	authenticationBase BaseAuthenticationBaseImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawAuthenticationBaseImpl) AuthenticationBase() BaseAuthenticationBaseImpl {
	return s.authenticationBase
}

func UnmarshalAuthenticationBaseImplementation(input []byte) (AuthenticationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.VideoAnalyzer.JwtAuthentication") {
		var out JwtAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JwtAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationBaseImpl: %+v", err)
	}

	return RawAuthenticationBaseImpl{
		authenticationBase: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
