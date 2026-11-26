package endpoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Credentials = AzureKeyVaultS3WithHmacCredentials{}

type AzureKeyVaultS3WithHmacCredentials struct {
	AccessKeyUri *string `json:"accessKeyUri,omitempty"`
	SecretKeyUri *string `json:"secretKeyUri,omitempty"`

	// Fields inherited from Credentials

	Type CredentialType `json:"type"`
}

func (s AzureKeyVaultS3WithHmacCredentials) Credentials() BaseCredentialsImpl {
	return BaseCredentialsImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = AzureKeyVaultS3WithHmacCredentials{}

func (s AzureKeyVaultS3WithHmacCredentials) MarshalJSON() ([]byte, error) {
	type wrapper AzureKeyVaultS3WithHmacCredentials
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureKeyVaultS3WithHmacCredentials: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureKeyVaultS3WithHmacCredentials: %+v", err)
	}

	decoded["type"] = "AzureKeyVaultS3WithHMAC"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureKeyVaultS3WithHmacCredentials: %+v", err)
	}

	return encoded, nil
}
