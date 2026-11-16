package authenticationsettings

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationSettingProperties = ManagedIdentityAuthenticationSettingProperties{}

type ManagedIdentityAuthenticationSettingProperties struct {
	ManagedIdentityName string `json:"managedIdentityName"`

	// Fields inherited from AuthenticationSettingProperties

	AuthenticationKind AuthenticationKind            `json:"authenticationKind"`
	DisplayName        *string                       `json:"displayName,omitempty"`
	ProvisioningState  *HealthModelProvisioningState `json:"provisioningState,omitempty"`
}

func (s ManagedIdentityAuthenticationSettingProperties) AuthenticationSettingProperties() BaseAuthenticationSettingPropertiesImpl {
	return BaseAuthenticationSettingPropertiesImpl{
		AuthenticationKind: s.AuthenticationKind,
		DisplayName:        s.DisplayName,
		ProvisioningState:  s.ProvisioningState,
	}
}

var _ json.Marshaler = ManagedIdentityAuthenticationSettingProperties{}

func (s ManagedIdentityAuthenticationSettingProperties) MarshalJSON() ([]byte, error) {
	type wrapper ManagedIdentityAuthenticationSettingProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedIdentityAuthenticationSettingProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedIdentityAuthenticationSettingProperties: %+v", err)
	}

	decoded["authenticationKind"] = "ManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedIdentityAuthenticationSettingProperties: %+v", err)
	}

	return encoded, nil
}
