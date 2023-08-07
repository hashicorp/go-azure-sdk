package datastore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatastoreCredentials = KerberosKeytabCredentials{}

type KerberosKeytabCredentials struct {
	KerberosKdcAddress string                `json:"kerberosKdcAddress"`
	KerberosPrincipal  string                `json:"kerberosPrincipal"`
	KerberosRealm      string                `json:"kerberosRealm"`
	Secrets            KerberosKeytabSecrets `json:"secrets"`

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
