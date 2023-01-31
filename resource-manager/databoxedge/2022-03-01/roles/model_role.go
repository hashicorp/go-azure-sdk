package roles

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Role interface {
}

func unmarshalRoleImplementation(input []byte) (Role, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Role into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "CloudEdgeManagement") {
		var out CloudEdgeManagementRole
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudEdgeManagementRole: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IOT") {
		var out IoTRole
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IoTRole: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Kubernetes") {
		var out KubernetesRole
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KubernetesRole: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MEC") {
		var out MECRole
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MECRole: %+v", err)
		}
		return out, nil
	}

	type RawRoleImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawRoleImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
