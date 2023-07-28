package datastore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatastoreSecrets = KerberosPasswordSecrets{}

type KerberosPasswordSecrets struct {
	KerberosPassword *string `json:"kerberosPassword,omitempty"`

	// Fields inherited from DatastoreSecrets
}

var _ json.Marshaler = KerberosPasswordSecrets{}

func (s KerberosPasswordSecrets) MarshalJSON() ([]byte, error) {
	type wrapper KerberosPasswordSecrets
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KerberosPasswordSecrets: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KerberosPasswordSecrets: %+v", err)
	}
	decoded["secretsType"] = "KerberosPassword"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KerberosPasswordSecrets: %+v", err)
	}

	return encoded, nil
}
