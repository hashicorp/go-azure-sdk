package javacomponents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JavaComponentProperties = SpringCloudEurekaComponent{}

type SpringCloudEurekaComponent struct {
	Ingress *JavaComponentIngress `json:"ingress,omitempty"`

	// Fields inherited from JavaComponentProperties

	ComponentType     JavaComponentType                     `json:"componentType"`
	Configurations    *[]JavaComponentConfigurationProperty `json:"configurations,omitempty"`
	ProvisioningState *JavaComponentProvisioningState       `json:"provisioningState,omitempty"`
	ServiceBinds      *[]JavaComponentServiceBind           `json:"serviceBinds,omitempty"`
}

func (s SpringCloudEurekaComponent) JavaComponentProperties() BaseJavaComponentPropertiesImpl {
	return BaseJavaComponentPropertiesImpl{
		ComponentType:     s.ComponentType,
		Configurations:    s.Configurations,
		ProvisioningState: s.ProvisioningState,
		ServiceBinds:      s.ServiceBinds,
	}
}

var _ json.Marshaler = SpringCloudEurekaComponent{}

func (s SpringCloudEurekaComponent) MarshalJSON() ([]byte, error) {
	type wrapper SpringCloudEurekaComponent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SpringCloudEurekaComponent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SpringCloudEurekaComponent: %+v", err)
	}

	decoded["componentType"] = "SpringCloudEureka"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SpringCloudEurekaComponent: %+v", err)
	}

	return encoded, nil
}
