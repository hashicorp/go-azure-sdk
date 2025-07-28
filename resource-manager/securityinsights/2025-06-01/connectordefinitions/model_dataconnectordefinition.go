package connectordefinitions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorDefinition interface {
	DataConnectorDefinition() BaseDataConnectorDefinitionImpl
}

var _ DataConnectorDefinition = BaseDataConnectorDefinitionImpl{}

type BaseDataConnectorDefinitionImpl struct {
	Etag       *string                     `json:"etag,omitempty"`
	Id         *string                     `json:"id,omitempty"`
	Kind       DataConnectorDefinitionKind `json:"kind"`
	Name       *string                     `json:"name,omitempty"`
	SystemData *systemdata.SystemData      `json:"systemData,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}

func (s BaseDataConnectorDefinitionImpl) DataConnectorDefinition() BaseDataConnectorDefinitionImpl {
	return s
}

var _ DataConnectorDefinition = RawDataConnectorDefinitionImpl{}

// RawDataConnectorDefinitionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataConnectorDefinitionImpl struct {
	dataConnectorDefinition BaseDataConnectorDefinitionImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawDataConnectorDefinitionImpl) DataConnectorDefinition() BaseDataConnectorDefinitionImpl {
	return s.dataConnectorDefinition
}

func UnmarshalDataConnectorDefinitionImplementation(input []byte) (DataConnectorDefinition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataConnectorDefinition into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Customizable") {
		var out CustomizableConnectorDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomizableConnectorDefinition: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataConnectorDefinitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataConnectorDefinitionImpl: %+v", err)
	}

	return RawDataConnectorDefinitionImpl{
		dataConnectorDefinition: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
