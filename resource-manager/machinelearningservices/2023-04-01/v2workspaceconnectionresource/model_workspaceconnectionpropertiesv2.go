package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceConnectionPropertiesV2 interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWorkspaceConnectionPropertiesV2Impl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalWorkspaceConnectionPropertiesV2Implementation(input []byte) (WorkspaceConnectionPropertiesV2, error) {
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

	out := RawWorkspaceConnectionPropertiesV2Impl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
