package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpalaLinkedServiceTypeProperties struct {
	AllowHostNameCNMismatch   *interface{}             `json:"allowHostNameCNMismatch,omitempty"`
	AllowSelfSignedServerCert *interface{}             `json:"allowSelfSignedServerCert,omitempty"`
	AuthenticationType        ImpalaAuthenticationType `json:"authenticationType"`
	EnableSsl                 *interface{}             `json:"enableSsl,omitempty"`
	EncryptedCredential       *string                  `json:"encryptedCredential,omitempty"`
	Host                      interface{}              `json:"host"`
	Password                  SecretBase               `json:"password"`
	Port                      *interface{}             `json:"port,omitempty"`
	TrustedCertPath           *interface{}             `json:"trustedCertPath,omitempty"`
	UseSystemTrustStore       *interface{}             `json:"useSystemTrustStore,omitempty"`
	Username                  *interface{}             `json:"username,omitempty"`
}

var _ json.Unmarshaler = &ImpalaLinkedServiceTypeProperties{}

func (s *ImpalaLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias ImpalaLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ImpalaLinkedServiceTypeProperties: %+v", err)
	}

	s.AllowHostNameCNMismatch = decoded.AllowHostNameCNMismatch
	s.AllowSelfSignedServerCert = decoded.AllowSelfSignedServerCert
	s.AuthenticationType = decoded.AuthenticationType
	s.EnableSsl = decoded.EnableSsl
	s.EncryptedCredential = decoded.EncryptedCredential
	s.Host = decoded.Host
	s.Port = decoded.Port
	s.TrustedCertPath = decoded.TrustedCertPath
	s.UseSystemTrustStore = decoded.UseSystemTrustStore
	s.Username = decoded.Username

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ImpalaLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["password"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Password' for 'ImpalaLinkedServiceTypeProperties': %+v", err)
		}
		s.Password = impl
	}
	return nil
}
