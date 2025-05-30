package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoenixLinkedServiceTypeProperties struct {
	AllowHostNameCNMismatch   *bool                     `json:"allowHostNameCNMismatch,omitempty"`
	AllowSelfSignedServerCert *bool                     `json:"allowSelfSignedServerCert,omitempty"`
	AuthenticationType        PhoenixAuthenticationType `json:"authenticationType"`
	EnableSsl                 *bool                     `json:"enableSsl,omitempty"`
	EncryptedCredential       *string                   `json:"encryptedCredential,omitempty"`
	HTTPPath                  *interface{}              `json:"httpPath,omitempty"`
	Host                      interface{}               `json:"host"`
	Password                  SecretBase                `json:"password"`
	Port                      *int64                    `json:"port,omitempty"`
	TrustedCertPath           *interface{}              `json:"trustedCertPath,omitempty"`
	UseSystemTrustStore       *bool                     `json:"useSystemTrustStore,omitempty"`
	Username                  *interface{}              `json:"username,omitempty"`
}

var _ json.Unmarshaler = &PhoenixLinkedServiceTypeProperties{}

func (s *PhoenixLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowHostNameCNMismatch   *bool                     `json:"allowHostNameCNMismatch,omitempty"`
		AllowSelfSignedServerCert *bool                     `json:"allowSelfSignedServerCert,omitempty"`
		AuthenticationType        PhoenixAuthenticationType `json:"authenticationType"`
		EnableSsl                 *bool                     `json:"enableSsl,omitempty"`
		EncryptedCredential       *string                   `json:"encryptedCredential,omitempty"`
		HTTPPath                  *interface{}              `json:"httpPath,omitempty"`
		Host                      interface{}               `json:"host"`
		Port                      *int64                    `json:"port,omitempty"`
		TrustedCertPath           *interface{}              `json:"trustedCertPath,omitempty"`
		UseSystemTrustStore       *bool                     `json:"useSystemTrustStore,omitempty"`
		Username                  *interface{}              `json:"username,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowHostNameCNMismatch = decoded.AllowHostNameCNMismatch
	s.AllowSelfSignedServerCert = decoded.AllowSelfSignedServerCert
	s.AuthenticationType = decoded.AuthenticationType
	s.EnableSsl = decoded.EnableSsl
	s.EncryptedCredential = decoded.EncryptedCredential
	s.HTTPPath = decoded.HTTPPath
	s.Host = decoded.Host
	s.Port = decoded.Port
	s.TrustedCertPath = decoded.TrustedCertPath
	s.UseSystemTrustStore = decoded.UseSystemTrustStore
	s.Username = decoded.Username

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PhoenixLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["password"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Password' for 'PhoenixLinkedServiceTypeProperties': %+v", err)
		}
		s.Password = impl
	}

	return nil
}
