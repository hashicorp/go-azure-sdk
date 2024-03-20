package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmazonS3LinkedServiceTypeProperties struct {
	AccessKeyId         *interface{} `json:"accessKeyId,omitempty"`
	AuthenticationType  *interface{} `json:"authenticationType,omitempty"`
	EncryptedCredential *string      `json:"encryptedCredential,omitempty"`
	SecretAccessKey     SecretBase   `json:"secretAccessKey"`
	ServiceUrl          *interface{} `json:"serviceUrl,omitempty"`
	SessionToken        SecretBase   `json:"sessionToken"`
}

var _ json.Unmarshaler = &AmazonS3LinkedServiceTypeProperties{}

func (s *AmazonS3LinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias AmazonS3LinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AmazonS3LinkedServiceTypeProperties: %+v", err)
	}

	s.AccessKeyId = decoded.AccessKeyId
	s.AuthenticationType = decoded.AuthenticationType
	s.EncryptedCredential = decoded.EncryptedCredential
	s.ServiceUrl = decoded.ServiceUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AmazonS3LinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["secretAccessKey"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SecretAccessKey' for 'AmazonS3LinkedServiceTypeProperties': %+v", err)
		}
		s.SecretAccessKey = impl
	}

	if v, ok := temp["sessionToken"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SessionToken' for 'AmazonS3LinkedServiceTypeProperties': %+v", err)
		}
		s.SessionToken = impl
	}
	return nil
}
