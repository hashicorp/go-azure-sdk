package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GoogleCloudStorageLinkedServiceTypeProperties struct {
	AccessKeyId         *string    `json:"accessKeyId,omitempty"`
	EncryptedCredential *string    `json:"encryptedCredential,omitempty"`
	SecretAccessKey     SecretBase `json:"secretAccessKey"`
	ServiceUrl          *string    `json:"serviceUrl,omitempty"`
}

var _ json.Unmarshaler = &GoogleCloudStorageLinkedServiceTypeProperties{}

func (s *GoogleCloudStorageLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias GoogleCloudStorageLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into GoogleCloudStorageLinkedServiceTypeProperties: %+v", err)
	}

	s.AccessKeyId = decoded.AccessKeyId
	s.EncryptedCredential = decoded.EncryptedCredential
	s.ServiceUrl = decoded.ServiceUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GoogleCloudStorageLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["secretAccessKey"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SecretAccessKey' for 'GoogleCloudStorageLinkedServiceTypeProperties': %+v", err)
		}
		s.SecretAccessKey = impl
	}
	return nil
}
