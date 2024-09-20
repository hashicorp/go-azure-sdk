package datareference

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataReferenceCredential = DockerCredential{}

type DockerCredential struct {
	Password *string `json:"password,omitempty"`
	UserName *string `json:"userName,omitempty"`

	// Fields inherited from DataReferenceCredential

	CredentialType DataReferenceCredentialType `json:"credentialType"`
}

func (s DockerCredential) DataReferenceCredential() BaseDataReferenceCredentialImpl {
	return BaseDataReferenceCredentialImpl{
		CredentialType: s.CredentialType,
	}
}

var _ json.Marshaler = DockerCredential{}

func (s DockerCredential) MarshalJSON() ([]byte, error) {
	type wrapper DockerCredential
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DockerCredential: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DockerCredential: %+v", err)
	}

	decoded["credentialType"] = "DockerCredentials"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DockerCredential: %+v", err)
	}

	return encoded, nil
}
