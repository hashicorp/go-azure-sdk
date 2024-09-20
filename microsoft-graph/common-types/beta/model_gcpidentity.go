package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpIdentity interface {
	Entity
	AuthorizationSystemIdentity
	GcpIdentity() BaseGcpIdentityImpl
}

var _ GcpIdentity = BaseGcpIdentityImpl{}

type BaseGcpIdentityImpl struct {

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

func (s BaseGcpIdentityImpl) GcpIdentity() BaseGcpIdentityImpl {
	return s
}

func (s BaseGcpIdentityImpl) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
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

func (s BaseGcpIdentityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GcpIdentity = RawGcpIdentityImpl{}

// RawGcpIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGcpIdentityImpl struct {
	gcpIdentity BaseGcpIdentityImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawGcpIdentityImpl) GcpIdentity() BaseGcpIdentityImpl {
	return s.gcpIdentity
}

func (s RawGcpIdentityImpl) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
	return s.gcpIdentity.AuthorizationSystemIdentity()
}

func (s RawGcpIdentityImpl) Entity() BaseEntityImpl {
	return s.gcpIdentity.Entity()
}

var _ json.Marshaler = BaseGcpIdentityImpl{}

func (s BaseGcpIdentityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGcpIdentityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGcpIdentityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGcpIdentityImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.gcpIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGcpIdentityImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseGcpIdentityImpl{}

func (s *BaseGcpIdentityImpl) UnmarshalJSON(bytes []byte) error {

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
		return fmt.Errorf("unmarshaling BaseGcpIdentityImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'BaseGcpIdentityImpl': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentitySourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'BaseGcpIdentityImpl': %+v", err)
		}
		s.Source = impl
	}

	return nil
}

func UnmarshalGcpIdentityImplementation(input []byte) (GcpIdentity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpIdentity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpCloudFunction") {
		var out GcpCloudFunction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpCloudFunction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpGroup") {
		var out GcpGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpServiceAccount") {
		var out GcpServiceAccount
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpServiceAccount: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpUser") {
		var out GcpUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpUser: %+v", err)
		}
		return out, nil
	}

	var parent BaseGcpIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGcpIdentityImpl: %+v", err)
	}

	return RawGcpIdentityImpl{
		gcpIdentity: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
