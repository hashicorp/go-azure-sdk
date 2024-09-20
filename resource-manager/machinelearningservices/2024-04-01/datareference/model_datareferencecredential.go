package datareference

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataReferenceCredential interface {
	DataReferenceCredential() BaseDataReferenceCredentialImpl
}

var _ DataReferenceCredential = BaseDataReferenceCredentialImpl{}

type BaseDataReferenceCredentialImpl struct {
	CredentialType DataReferenceCredentialType `json:"credentialType"`
}

func (s BaseDataReferenceCredentialImpl) DataReferenceCredential() BaseDataReferenceCredentialImpl {
	return s
}

var _ DataReferenceCredential = RawDataReferenceCredentialImpl{}

// RawDataReferenceCredentialImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataReferenceCredentialImpl struct {
	dataReferenceCredential BaseDataReferenceCredentialImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawDataReferenceCredentialImpl) DataReferenceCredential() BaseDataReferenceCredentialImpl {
	return s.dataReferenceCredential
}

func UnmarshalDataReferenceCredentialImplementation(input []byte) (DataReferenceCredential, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataReferenceCredential into map[string]interface: %+v", err)
	}

	value, ok := temp["credentialType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "NoCredentials") {
		var out AnonymousAccessCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AnonymousAccessCredential: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DockerCredentials") {
		var out DockerCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DockerCredential: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ManagedIdentity") {
		var out ManagedIdentityCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIdentityCredential: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SAS") {
		var out SASCredential
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SASCredential: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataReferenceCredentialImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataReferenceCredentialImpl: %+v", err)
	}

	return RawDataReferenceCredentialImpl{
		dataReferenceCredential: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
