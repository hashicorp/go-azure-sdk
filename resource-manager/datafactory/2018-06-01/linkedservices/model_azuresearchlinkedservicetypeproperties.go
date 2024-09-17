package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureSearchLinkedServiceTypeProperties struct {
	EncryptedCredential *string    `json:"encryptedCredential,omitempty"`
	Key                 SecretBase `json:"key"`
	Url                 string     `json:"url"`
}

var _ json.Unmarshaler = &AzureSearchLinkedServiceTypeProperties{}

func (s *AzureSearchLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias AzureSearchLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AzureSearchLinkedServiceTypeProperties: %+v", err)
	}

	s.EncryptedCredential = decoded.EncryptedCredential
	s.Url = decoded.Url

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AzureSearchLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["key"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Key' for 'AzureSearchLinkedServiceTypeProperties': %+v", err)
		}
		s.Key = impl
	}
	return nil
}
