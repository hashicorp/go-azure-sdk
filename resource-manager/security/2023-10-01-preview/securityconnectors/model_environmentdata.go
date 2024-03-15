package securityconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentData interface {
}

// RawEnvironmentDataImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEnvironmentDataImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalEnvironmentDataImplementation(input []byte) (EnvironmentData, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EnvironmentData into map[string]interface: %+v", err)
	}

	value, ok := temp["environmentType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AwsAccount") {
		var out AwsEnvironmentData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsEnvironmentData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureDevOpsScope") {
		var out AzureDevOpsScopeEnvironmentData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureDevOpsScopeEnvironmentData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GcpProject") {
		var out GcpProjectEnvironmentData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpProjectEnvironmentData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GithubScope") {
		var out GithubScopeEnvironmentData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GithubScopeEnvironmentData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GitlabScope") {
		var out GitlabScopeEnvironmentData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GitlabScopeEnvironmentData: %+v", err)
		}
		return out, nil
	}

	out := RawEnvironmentDataImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
