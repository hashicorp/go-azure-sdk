package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodConfiguration = ExternalAuthenticationMethodConfiguration{}

type ExternalAuthenticationMethodConfiguration struct {
	// appId for the app registration in Microsoft Entra ID representing the integration with the external provider.
	AppId *string `json:"appId,omitempty"`

	// Display name for the external authentication method. This name is shown to users during sign-in.
	DisplayName *string `json:"displayName,omitempty"`

	// A collection of groups that are enabled to use an authentication method as part of an authentication method policy in
	// Microsoft Entra ID.
	IncludeTargets *[]AuthenticationMethodTarget `json:"includeTargets,omitempty"`

	OpenIdConnectSetting *OpenIdConnectSetting `json:"openIdConnectSetting,omitempty"`

	// Fields inherited from AuthenticationMethodConfiguration

	// Groups of users that are excluded from a policy.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// The state of the policy. Possible values are: enabled, disabled.
	State *AuthenticationMethodState `json:"state,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s ExternalAuthenticationMethodConfiguration) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return BaseAuthenticationMethodConfigurationImpl{
		ExcludeTargets: s.ExcludeTargets,
		State:          s.State,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s ExternalAuthenticationMethodConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalAuthenticationMethodConfiguration{}

func (s ExternalAuthenticationMethodConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper ExternalAuthenticationMethodConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalAuthenticationMethodConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalAuthenticationMethodConfiguration: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.externalAuthenticationMethodConfiguration"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalAuthenticationMethodConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExternalAuthenticationMethodConfiguration{}

func (s *ExternalAuthenticationMethodConfiguration) UnmarshalJSON(bytes []byte) error {
	type alias ExternalAuthenticationMethodConfiguration
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ExternalAuthenticationMethodConfiguration: %+v", err)
	}

	s.AppId = decoded.AppId
	s.DisplayName = decoded.DisplayName
	s.ExcludeTargets = decoded.ExcludeTargets
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OpenIdConnectSetting = decoded.OpenIdConnectSetting
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExternalAuthenticationMethodConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["includeTargets"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IncludeTargets into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationMethodTarget, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationMethodTargetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IncludeTargets' for 'ExternalAuthenticationMethodConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IncludeTargets = &output
	}
	return nil
}
