package connectordefinitions

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorDefinition = CustomizableConnectorDefinition{}

type CustomizableConnectorDefinition struct {
	Properties *CustomizableConnectorDefinitionProperties `json:"properties,omitempty"`

	// Fields inherited from DataConnectorDefinition

	Etag       *string                     `json:"etag,omitempty"`
	Id         *string                     `json:"id,omitempty"`
	Kind       DataConnectorDefinitionKind `json:"kind"`
	Name       *string                     `json:"name,omitempty"`
	SystemData *systemdata.SystemData      `json:"systemData,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}

func (s CustomizableConnectorDefinition) DataConnectorDefinition() BaseDataConnectorDefinitionImpl {
	return BaseDataConnectorDefinitionImpl{
		Etag:       s.Etag,
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = CustomizableConnectorDefinition{}

func (s CustomizableConnectorDefinition) MarshalJSON() ([]byte, error) {
	type wrapper CustomizableConnectorDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomizableConnectorDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomizableConnectorDefinition: %+v", err)
	}

	decoded["kind"] = "Customizable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomizableConnectorDefinition: %+v", err)
	}

	return encoded, nil
}
