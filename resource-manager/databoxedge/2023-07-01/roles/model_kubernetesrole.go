package roles

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Role = KubernetesRole{}

type KubernetesRole struct {
	Properties *KubernetesRoleProperties `json:"properties,omitempty"`

	// Fields inherited from Role

	Id         *string                `json:"id,omitempty"`
	Kind       RoleTypes              `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s KubernetesRole) Role() BaseRoleImpl {
	return BaseRoleImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = KubernetesRole{}

func (s KubernetesRole) MarshalJSON() ([]byte, error) {
	type wrapper KubernetesRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KubernetesRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KubernetesRole: %+v", err)
	}

	decoded["kind"] = "Kubernetes"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KubernetesRole: %+v", err)
	}

	return encoded, nil
}
