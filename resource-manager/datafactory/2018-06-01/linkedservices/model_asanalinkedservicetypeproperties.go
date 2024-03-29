package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AsanaLinkedServiceTypeProperties struct {
	ApiToken            SecretBase `json:"apiToken"`
	EncryptedCredential *string    `json:"encryptedCredential,omitempty"`
}

var _ json.Unmarshaler = &AsanaLinkedServiceTypeProperties{}

func (s *AsanaLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias AsanaLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AsanaLinkedServiceTypeProperties: %+v", err)
	}

	s.EncryptedCredential = decoded.EncryptedCredential

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AsanaLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["apiToken"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ApiToken' for 'AsanaLinkedServiceTypeProperties': %+v", err)
		}
		s.ApiToken = impl
	}
	return nil
}
