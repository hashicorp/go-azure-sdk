package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SSISAccessCredential struct {
	Domain   string     `json:"domain"`
	Password SecretBase `json:"password"`
	UserName string     `json:"userName"`
}

var _ json.Unmarshaler = &SSISAccessCredential{}

func (s *SSISAccessCredential) UnmarshalJSON(bytes []byte) error {
	type alias SSISAccessCredential
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SSISAccessCredential: %+v", err)
	}

	s.Domain = decoded.Domain
	s.UserName = decoded.UserName

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SSISAccessCredential into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["password"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Password' for 'SSISAccessCredential': %+v", err)
		}
		s.Password = impl
	}
	return nil
}
