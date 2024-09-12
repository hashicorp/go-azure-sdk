package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityUserFlow = B2xIdentityUserFlow{}

type B2xIdentityUserFlow struct {
	// Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only
	// obtain the value of this object using Get userFlowApiConnectorConfiguration.
	ApiConnectorConfiguration *UserFlowApiConnectorConfiguration `json:"apiConnectorConfiguration,omitempty"`

	// The identity providers included in the user flow.
	IdentityProviders *[]IdentityProvider `json:"identityProviders,omitempty"`

	// The languages supported for customization within the user flow. Language customization is enabled by default in
	// self-service sign-up user flow. You can't create custom languages in self-service sign-up user flows.
	Languages *[]UserFlowLanguageConfiguration `json:"languages,omitempty"`

	// The user attribute assignments included in the user flow.
	UserAttributeAssignments *[]IdentityUserFlowAttributeAssignment `json:"userAttributeAssignments,omitempty"`

	UserFlowIdentityProviders *[]IdentityProviderBase `json:"userFlowIdentityProviders,omitempty"`

	// Fields inherited from IdentityUserFlow

	UserFlowType *UserFlowType `json:"userFlowType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s B2xIdentityUserFlow) IdentityUserFlow() BaseIdentityUserFlowImpl {
	return BaseIdentityUserFlowImpl{
		UserFlowType: s.UserFlowType,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s B2xIdentityUserFlow) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = B2xIdentityUserFlow{}

func (s B2xIdentityUserFlow) MarshalJSON() ([]byte, error) {
	type wrapper B2xIdentityUserFlow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling B2xIdentityUserFlow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling B2xIdentityUserFlow: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.b2xIdentityUserFlow"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling B2xIdentityUserFlow: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &B2xIdentityUserFlow{}

func (s *B2xIdentityUserFlow) UnmarshalJSON(bytes []byte) error {
	type alias B2xIdentityUserFlow
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into B2xIdentityUserFlow: %+v", err)
	}

	s.ApiConnectorConfiguration = decoded.ApiConnectorConfiguration
	s.Id = decoded.Id
	s.Languages = decoded.Languages
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.UserAttributeAssignments = decoded.UserAttributeAssignments
	s.UserFlowType = decoded.UserFlowType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling B2xIdentityUserFlow into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityProviders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IdentityProviders into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityProvider, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityProviderImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IdentityProviders' for 'B2xIdentityUserFlow': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IdentityProviders = &output
	}

	if v, ok := temp["userFlowIdentityProviders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling UserFlowIdentityProviders into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityProviderBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityProviderBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'UserFlowIdentityProviders' for 'B2xIdentityUserFlow': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.UserFlowIdentityProviders = &output
	}
	return nil
}
