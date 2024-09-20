package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceConnectionPropertiesV2 interface {
	WorkspaceConnectionPropertiesV2() BaseWorkspaceConnectionPropertiesV2Impl
}

var _ WorkspaceConnectionPropertiesV2 = BaseWorkspaceConnectionPropertiesV2Impl{}

type BaseWorkspaceConnectionPropertiesV2Impl struct {
	AuthType    ConnectionAuthType  `json:"authType"`
	Category    *ConnectionCategory `json:"category,omitempty"`
	Target      *string             `json:"target,omitempty"`
	Value       *string             `json:"value,omitempty"`
	ValueFormat *ValueFormat        `json:"valueFormat,omitempty"`
}

func (s BaseWorkspaceConnectionPropertiesV2Impl) WorkspaceConnectionPropertiesV2() BaseWorkspaceConnectionPropertiesV2Impl {
	return s
}

var _ WorkspaceConnectionPropertiesV2 = RawWorkspaceConnectionPropertiesV2Impl{}

// RawWorkspaceConnectionPropertiesV2Impl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWorkspaceConnectionPropertiesV2Impl struct {
	workspaceConnectionPropertiesV2 BaseWorkspaceConnectionPropertiesV2Impl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawWorkspaceConnectionPropertiesV2Impl) WorkspaceConnectionPropertiesV2() BaseWorkspaceConnectionPropertiesV2Impl {
	return s.workspaceConnectionPropertiesV2
}

func UnmarshalWorkspaceConnectionPropertiesV2Implementation(input []byte) (WorkspaceConnectionPropertiesV2, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkspaceConnectionPropertiesV2 into map[string]interface: %+v", err)
	}

	value, ok := temp["authType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "ManagedIdentity") {
		var out ManagedIdentityAuthTypeWorkspaceConnectionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIdentityAuthTypeWorkspaceConnectionProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "None") {
		var out NoneAuthTypeWorkspaceConnectionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoneAuthTypeWorkspaceConnectionProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PAT") {
		var out PATAuthTypeWorkspaceConnectionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PATAuthTypeWorkspaceConnectionProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SAS") {
		var out SASAuthTypeWorkspaceConnectionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SASAuthTypeWorkspaceConnectionProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UsernamePassword") {
		var out UsernamePasswordAuthTypeWorkspaceConnectionProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UsernamePasswordAuthTypeWorkspaceConnectionProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseWorkspaceConnectionPropertiesV2Impl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWorkspaceConnectionPropertiesV2Impl: %+v", err)
	}

	return RawWorkspaceConnectionPropertiesV2Impl{
		workspaceConnectionPropertiesV2: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
