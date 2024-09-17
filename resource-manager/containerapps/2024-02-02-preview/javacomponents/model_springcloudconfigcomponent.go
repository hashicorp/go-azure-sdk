package javacomponents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JavaComponentProperties = SpringCloudConfigComponent{}

type SpringCloudConfigComponent struct {

	// Fields inherited from JavaComponentProperties

	ComponentType     JavaComponentType                     `json:"componentType"`
	Configurations    *[]JavaComponentConfigurationProperty `json:"configurations,omitempty"`
	ProvisioningState *JavaComponentProvisioningState       `json:"provisioningState,omitempty"`
	ServiceBinds      *[]JavaComponentServiceBind           `json:"serviceBinds,omitempty"`
}

func (s SpringCloudConfigComponent) JavaComponentProperties() BaseJavaComponentPropertiesImpl {
	return BaseJavaComponentPropertiesImpl{
		ComponentType:     s.ComponentType,
		Configurations:    s.Configurations,
		ProvisioningState: s.ProvisioningState,
		ServiceBinds:      s.ServiceBinds,
	}
}

var _ json.Marshaler = SpringCloudConfigComponent{}

func (s SpringCloudConfigComponent) MarshalJSON() ([]byte, error) {
	type wrapper SpringCloudConfigComponent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SpringCloudConfigComponent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SpringCloudConfigComponent: %+v", err)
	}

	decoded["componentType"] = "SpringCloudConfig"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SpringCloudConfigComponent: %+v", err)
	}

	return encoded, nil
}
