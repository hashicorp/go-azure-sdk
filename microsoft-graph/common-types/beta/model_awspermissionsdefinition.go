package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PermissionsDefinition = AwsPermissionsDefinition{}

type AwsPermissionsDefinition struct {
	ActionInfo *AwsPermissionsDefinitionAction `json:"actionInfo,omitempty"`

	// Fields inherited from PermissionsDefinition

	AuthorizationSystemInfo *PermissionsDefinitionAuthorizationSystem         `json:"authorizationSystemInfo,omitempty"`
	IdentityInfo            *PermissionsDefinitionAuthorizationSystemIdentity `json:"identityInfo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s AwsPermissionsDefinition) PermissionsDefinition() BasePermissionsDefinitionImpl {
	return BasePermissionsDefinitionImpl{
		AuthorizationSystemInfo: s.AuthorizationSystemInfo,
		IdentityInfo:            s.IdentityInfo,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

var _ json.Marshaler = AwsPermissionsDefinition{}

func (s AwsPermissionsDefinition) MarshalJSON() ([]byte, error) {
	type wrapper AwsPermissionsDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsPermissionsDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsPermissionsDefinition: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.awsPermissionsDefinition"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsPermissionsDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AwsPermissionsDefinition{}

func (s *AwsPermissionsDefinition) UnmarshalJSON(bytes []byte) error {
	type alias AwsPermissionsDefinition
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AwsPermissionsDefinition: %+v", err)
	}

	s.AuthorizationSystemInfo = decoded.AuthorizationSystemInfo
	s.IdentityInfo = decoded.IdentityInfo
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AwsPermissionsDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actionInfo"]; ok {
		impl, err := UnmarshalAwsPermissionsDefinitionActionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ActionInfo' for 'AwsPermissionsDefinition': %+v", err)
		}
		s.ActionInfo = &impl
	}
	return nil
}
