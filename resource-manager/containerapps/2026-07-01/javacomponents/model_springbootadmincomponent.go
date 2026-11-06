package javacomponents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JavaComponentProperties = SpringBootAdminComponent{}

type SpringBootAdminComponent struct {
	Ingress *JavaComponentIngress `json:"ingress,omitempty"`

	// Fields inherited from JavaComponentProperties

	ComponentType     JavaComponentType                     `json:"componentType"`
	Configurations    *[]JavaComponentConfigurationProperty `json:"configurations,omitempty"`
	ProvisioningState *JavaComponentProvisioningState       `json:"provisioningState,omitempty"`
	Scale             *JavaComponentPropertiesScale         `json:"scale,omitempty"`
	ServiceBinds      *[]JavaComponentServiceBind           `json:"serviceBinds,omitempty"`
}

func (s SpringBootAdminComponent) JavaComponentProperties() BaseJavaComponentPropertiesImpl {
	return BaseJavaComponentPropertiesImpl{
		ComponentType:     s.ComponentType,
		Configurations:    s.Configurations,
		ProvisioningState: s.ProvisioningState,
		Scale:             s.Scale,
		ServiceBinds:      s.ServiceBinds,
	}
}

var _ json.Marshaler = SpringBootAdminComponent{}

func (s SpringBootAdminComponent) MarshalJSON() ([]byte, error) {
	type wrapper SpringBootAdminComponent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SpringBootAdminComponent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SpringBootAdminComponent: %+v", err)
	}

	decoded["componentType"] = "SpringBootAdmin"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SpringBootAdminComponent: %+v", err)
	}

	return encoded, nil
}
