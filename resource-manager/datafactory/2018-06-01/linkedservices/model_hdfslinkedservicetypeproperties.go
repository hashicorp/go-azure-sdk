package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HdfsLinkedServiceTypeProperties struct {
	AuthenticationType  *interface{} `json:"authenticationType,omitempty"`
	EncryptedCredential *string      `json:"encryptedCredential,omitempty"`
	Password            SecretBase   `json:"password"`
	Url                 interface{}  `json:"url"`
	UserName            *interface{} `json:"userName,omitempty"`
}

var _ json.Unmarshaler = &HdfsLinkedServiceTypeProperties{}

func (s *HdfsLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias HdfsLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into HdfsLinkedServiceTypeProperties: %+v", err)
	}

	s.AuthenticationType = decoded.AuthenticationType
	s.EncryptedCredential = decoded.EncryptedCredential
	s.Url = decoded.Url
	s.UserName = decoded.UserName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling HdfsLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["password"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Password' for 'HdfsLinkedServiceTypeProperties': %+v", err)
		}
		s.Password = impl
	}
	return nil
}
