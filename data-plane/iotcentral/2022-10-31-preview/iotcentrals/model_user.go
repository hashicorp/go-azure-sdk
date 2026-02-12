package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type User interface {
	User() BaseUserImpl
}

var _ User = BaseUserImpl{}

type BaseUserImpl struct {
	Id    *string          `json:"id,omitempty"`
	Roles []RoleAssignment `json:"roles"`
	Type  string           `json:"type"`
}

func (s BaseUserImpl) User() BaseUserImpl {
	return s
}

var _ User = RawUserImpl{}

// RawUserImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawUserImpl struct {
	user   BaseUserImpl
	Type   string
	Values map[string]interface{}
}

func (s RawUserImpl) User() BaseUserImpl {
	return s.user
}

func UnmarshalUserImplementation(input []byte) (User, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling User into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "adGroup") {
		var out ADGroupUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ADGroupUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "email") {
		var out EmailUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "servicePrincipal") {
		var out ServicePrincipalUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalUser: %+v", err)
		}
		return out, nil
	}

	var parent BaseUserImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUserImpl: %+v", err)
	}

	return RawUserImpl{
		user:   parent,
		Type:   value,
		Values: temp,
	}, nil

}
