package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesforceLinkedServiceTypeProperties struct {
	ApiVersion          *string    `json:"apiVersion,omitempty"`
	EncryptedCredential *string    `json:"encryptedCredential,omitempty"`
	EnvironmentUrl      *string    `json:"environmentUrl,omitempty"`
	Password            SecretBase `json:"password"`
	SecurityToken       SecretBase `json:"securityToken"`
	Username            *string    `json:"username,omitempty"`
}

var _ json.Unmarshaler = &SalesforceLinkedServiceTypeProperties{}

func (s *SalesforceLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias SalesforceLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SalesforceLinkedServiceTypeProperties: %+v", err)
	}

	s.ApiVersion = decoded.ApiVersion
	s.EncryptedCredential = decoded.EncryptedCredential
	s.EnvironmentUrl = decoded.EnvironmentUrl
	s.Username = decoded.Username

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SalesforceLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["password"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Password' for 'SalesforceLinkedServiceTypeProperties': %+v", err)
		}
		s.Password = impl
	}

	if v, ok := temp["securityToken"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SecurityToken' for 'SalesforceLinkedServiceTypeProperties': %+v", err)
		}
		s.SecurityToken = impl
	}
	return nil
}
