package roles

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Role interface {
	Role() BaseRoleImpl
}

var _ Role = BaseRoleImpl{}

type BaseRoleImpl struct {
	Id         *string                `json:"id,omitempty"`
	Kind       RoleTypes              `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseRoleImpl) Role() BaseRoleImpl {
	return s
}

var _ Role = RawRoleImpl{}

// RawRoleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRoleImpl struct {
	role   BaseRoleImpl
	Type   string
	Values map[string]interface{}
}

func (s RawRoleImpl) Role() BaseRoleImpl {
	return s.role
}

func UnmarshalRoleImplementation(input []byte) (Role, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Role into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseRoleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRoleImpl: %+v", err)
	}

	return RawRoleImpl{
		role:   parent,
		Type:   value,
		Values: temp,
	}, nil

}
