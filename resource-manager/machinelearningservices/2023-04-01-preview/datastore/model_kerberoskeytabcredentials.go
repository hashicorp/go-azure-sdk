package datastore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatastoreCredentials = KerberosKeytabCredentials{}

type KerberosKeytabCredentials struct {
	KerberosKdcAddress string           `json:"kerberosKdcAddress"`
	KerberosPrincipal  string           `json:"kerberosPrincipal"`
	KerberosRealm      string           `json:"kerberosRealm"`
	Secrets            DatastoreSecrets `json:"secrets"`

	// Fields inherited from DatastoreCredentials
}

var _ json.Marshaler = KerberosKeytabCredentials{}

func (s KerberosKeytabCredentials) MarshalJSON() ([]byte, error) {
	type wrapper KerberosKeytabCredentials
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KerberosKeytabCredentials: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KerberosKeytabCredentials: %+v", err)
	}
	decoded["credentialsType"] = "KerberosKeytab"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KerberosKeytabCredentials: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &KerberosKeytabCredentials{}

func (s *KerberosKeytabCredentials) UnmarshalJSON(bytes []byte) error {
	type alias KerberosKeytabCredentials
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into KerberosKeytabCredentials: %+v", err)
	}

	s.KerberosKdcAddress = decoded.KerberosKdcAddress
	s.KerberosPrincipal = decoded.KerberosPrincipal
	s.KerberosRealm = decoded.KerberosRealm

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling KerberosKeytabCredentials into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["secrets"]; ok {
		impl, err := unmarshalDatastoreSecretsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Secrets' for 'KerberosKeytabCredentials': %+v", err)
		}
		s.Secrets = impl
	}
	return nil
}
