package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureIdentity interface {
	Entity
	AuthorizationSystemIdentity
	AzureIdentity() BaseAzureIdentityImpl
}

var _ AzureIdentity = BaseAzureIdentityImpl{}

type BaseAzureIdentityImpl struct {

	// Fields inherited from AuthorizationSystemIdentity

	// Navigation to the authorizationSystem object
	AuthorizationSystem *AuthorizationSystem `json:"authorizationSystem,omitempty"`

	// The name of the identity. Read-only. Supports $filter and (eq,contains).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique ID of the identity within the external system. Read-only.
	ExternalId *string `json:"externalId,omitempty"`

	// Represents details of the source of the identity.
	Source AuthorizationSystemIdentitySource `json:"source"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseAzureIdentityImpl) AzureIdentity() BaseAzureIdentityImpl {
	return s
}

func (s BaseAzureIdentityImpl) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
	return BaseAuthorizationSystemIdentityImpl{
		AuthorizationSystem: s.AuthorizationSystem,
		DisplayName:         s.DisplayName,
		ExternalId:          s.ExternalId,
		Source:              s.Source,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s BaseAzureIdentityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AzureIdentity = RawAzureIdentityImpl{}

// RawAzureIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAzureIdentityImpl struct {
	azureIdentity BaseAzureIdentityImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawAzureIdentityImpl) AzureIdentity() BaseAzureIdentityImpl {
	return s.azureIdentity
}

func (s RawAzureIdentityImpl) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
	return s.azureIdentity.AuthorizationSystemIdentity()
}

func (s RawAzureIdentityImpl) Entity() BaseEntityImpl {
	return s.azureIdentity.Entity()
}

var _ json.Marshaler = BaseAzureIdentityImpl{}

func (s BaseAzureIdentityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAzureIdentityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAzureIdentityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAzureIdentityImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.azureIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAzureIdentityImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseAzureIdentityImpl{}

func (s *BaseAzureIdentityImpl) UnmarshalJSON(bytes []byte) error {

	var decoded struct {
		AuthorizationSystem *AuthorizationSystem              `json:"authorizationSystem,omitempty"`
		DisplayName         nullable.Type[string]             `json:"displayName,omitempty"`
		ExternalId          *string                           `json:"externalId,omitempty"`
		Source              AuthorizationSystemIdentitySource `json:"source"`
		Id                  *string                           `json:"id,omitempty"`
		ODataId             *string                           `json:"@odata.id,omitempty"`
		ODataType           *string                           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAzureIdentityImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'BaseAzureIdentityImpl': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentitySourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'BaseAzureIdentityImpl': %+v", err)
		}
		s.Source = impl
	}

	return nil
}

func UnmarshalAzureIdentityImplementation(input []byte) (AzureIdentity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureIdentity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.azureGroup") {
		var out AzureGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureManagedIdentity") {
		var out AzureManagedIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureManagedIdentity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureServerlessFunction") {
		var out AzureServerlessFunction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureServerlessFunction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureServicePrincipal") {
		var out AzureServicePrincipal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureServicePrincipal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureUser") {
		var out AzureUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureUser: %+v", err)
		}
		return out, nil
	}

	var parent BaseAzureIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAzureIdentityImpl: %+v", err)
	}

	return RawAzureIdentityImpl{
		azureIdentity: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
