package datastore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatastoreCredentials = KerberosPasswordCredentials{}

type KerberosPasswordCredentials struct {
	KerberosKdcAddress string                  `json:"kerberosKdcAddress"`
	KerberosPrincipal  string                  `json:"kerberosPrincipal"`
	KerberosRealm      string                  `json:"kerberosRealm"`
	Secrets            KerberosPasswordSecrets `json:"secrets"`

	// Fields inherited from DatastoreCredentials
}

var _ json.Marshaler = KerberosPasswordCredentials{}

func (s KerberosPasswordCredentials) MarshalJSON() ([]byte, error) {
	type wrapper KerberosPasswordCredentials
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KerberosPasswordCredentials: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KerberosPasswordCredentials: %+v", err)
	}
	decoded["credentialsType"] = "KerberosPassword"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KerberosPasswordCredentials: %+v", err)
	}

	return encoded, nil
}
