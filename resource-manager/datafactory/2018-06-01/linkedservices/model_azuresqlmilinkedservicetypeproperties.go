package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureSqlMILinkedServiceTypeProperties struct {
	AlwaysEncryptedSettings *SqlAlwaysEncryptedProperties `json:"alwaysEncryptedSettings,omitempty"`
	AzureCloudType          *interface{}                  `json:"azureCloudType,omitempty"`
	ConnectionString        interface{}                   `json:"connectionString"`
	Credential              *CredentialReference          `json:"credential,omitempty"`
	EncryptedCredential     *string                       `json:"encryptedCredential,omitempty"`
	Password                *AzureKeyVaultSecretReference `json:"password,omitempty"`
	ServicePrincipalId      *interface{}                  `json:"servicePrincipalId,omitempty"`
	ServicePrincipalKey     SecretBase                    `json:"servicePrincipalKey"`
	Tenant                  *interface{}                  `json:"tenant,omitempty"`
}

var _ json.Unmarshaler = &AzureSqlMILinkedServiceTypeProperties{}

func (s *AzureSqlMILinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias AzureSqlMILinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into AzureSqlMILinkedServiceTypeProperties: %+v", err)
	}

	s.AlwaysEncryptedSettings = decoded.AlwaysEncryptedSettings
	s.AzureCloudType = decoded.AzureCloudType
	s.ConnectionString = decoded.ConnectionString
	s.Credential = decoded.Credential
	s.EncryptedCredential = decoded.EncryptedCredential
	s.Password = decoded.Password
	s.ServicePrincipalId = decoded.ServicePrincipalId
	s.Tenant = decoded.Tenant

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AzureSqlMILinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["servicePrincipalKey"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ServicePrincipalKey' for 'AzureSqlMILinkedServiceTypeProperties': %+v", err)
		}
		s.ServicePrincipalKey = impl
	}
	return nil
}
