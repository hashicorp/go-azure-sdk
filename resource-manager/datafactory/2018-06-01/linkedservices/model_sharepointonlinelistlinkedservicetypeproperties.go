package linkedservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharePointOnlineListLinkedServiceTypeProperties struct {
	EncryptedCredential                  *string    `json:"encryptedCredential,omitempty"`
	ServicePrincipalCredentialType       *string    `json:"servicePrincipalCredentialType,omitempty"`
	ServicePrincipalEmbeddedCert         SecretBase `json:"servicePrincipalEmbeddedCert"`
	ServicePrincipalEmbeddedCertPassword SecretBase `json:"servicePrincipalEmbeddedCertPassword"`
	ServicePrincipalId                   string     `json:"servicePrincipalId"`
	ServicePrincipalKey                  SecretBase `json:"servicePrincipalKey"`
	SiteUrl                              string     `json:"siteUrl"`
	TenantId                             string     `json:"tenantId"`
}

var _ json.Unmarshaler = &SharePointOnlineListLinkedServiceTypeProperties{}

func (s *SharePointOnlineListLinkedServiceTypeProperties) UnmarshalJSON(bytes []byte) error {
	type alias SharePointOnlineListLinkedServiceTypeProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SharePointOnlineListLinkedServiceTypeProperties: %+v", err)
	}

	s.EncryptedCredential = decoded.EncryptedCredential
	s.ServicePrincipalCredentialType = decoded.ServicePrincipalCredentialType
	s.ServicePrincipalId = decoded.ServicePrincipalId
	s.SiteUrl = decoded.SiteUrl
	s.TenantId = decoded.TenantId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharePointOnlineListLinkedServiceTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["servicePrincipalEmbeddedCert"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ServicePrincipalEmbeddedCert' for 'SharePointOnlineListLinkedServiceTypeProperties': %+v", err)
		}
		s.ServicePrincipalEmbeddedCert = impl
	}

	if v, ok := temp["servicePrincipalEmbeddedCertPassword"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ServicePrincipalEmbeddedCertPassword' for 'SharePointOnlineListLinkedServiceTypeProperties': %+v", err)
		}
		s.ServicePrincipalEmbeddedCertPassword = impl
	}

	if v, ok := temp["servicePrincipalKey"]; ok {
		impl, err := UnmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ServicePrincipalKey' for 'SharePointOnlineListLinkedServiceTypeProperties': %+v", err)
		}
		s.ServicePrincipalKey = impl
	}
	return nil
}
