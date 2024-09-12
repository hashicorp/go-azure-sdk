package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionAction interface {
	PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl
}

var _ PermissionsDefinitionAction = BasePermissionsDefinitionActionImpl{}

type BasePermissionsDefinitionActionImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BasePermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return s
}

var _ PermissionsDefinitionAction = RawPermissionsDefinitionActionImpl{}

// RawPermissionsDefinitionActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPermissionsDefinitionActionImpl struct {
	permissionsDefinitionAction BasePermissionsDefinitionActionImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawPermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return s.permissionsDefinitionAction
}

func UnmarshalPermissionsDefinitionActionImplementation(input []byte) (PermissionsDefinitionAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsDefinitionAction into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsPermissionsDefinitionAction") {
		var out AwsPermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsPermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azurePermissionsDefinitionAction") {
		var out AzurePermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzurePermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpPermissionsDefinitionAction") {
		var out GcpPermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpPermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	var parent BasePermissionsDefinitionActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePermissionsDefinitionActionImpl: %+v", err)
	}

	return RawPermissionsDefinitionActionImpl{
		permissionsDefinitionAction: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
