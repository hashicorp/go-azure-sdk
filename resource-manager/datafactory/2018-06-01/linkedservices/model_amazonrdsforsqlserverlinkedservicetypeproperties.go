package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmazonRdsForSqlServerLinkedServiceTypeProperties struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedProperties `json:"alwaysEncryptedSettings,omitempty"`
	ConnectionString        interface{}                   `json:"connectionString"`
	EncryptedCredential     *string                       `json:"encryptedCredential,omitempty"`
	Password                SecretBase                    `json:"password"`
	UserName                *interface{}                  `json:"userName,omitempty"`
}

var _ json.Unmarshaler = &AmazonRdsForSqlServerLinkedServiceTypeProperties{}

func (s *AmazonRdsForSqlServerLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias AmazonRdsForSqlServerLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AmazonRdsForSqlServerLinkedServiceTypeProperties: %+v", err)
	}

	s.AlwaysEncryptedSettings = decoded.AlwaysEncryptedSettings
	s.ConnectionString = decoded.ConnectionString
	s.EncryptedCredential = decoded.EncryptedCredential
	s.UserName = decoded.UserName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AmazonRdsForSqlServerLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["password"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Password' for 'AmazonRdsForSqlServerLinkedServiceTypeProperties': %+v", err)
		}
		s.Password = impl
	}
	return nil
}
