package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeSsisCatalogInfo struct {
	CatalogAdminPassword  SecretBase                                `json:"catalogAdminPassword"`
	CatalogAdminUserName  *string                                   `json:"catalogAdminUserName,omitempty"`
	CatalogPricingTier    *IntegrationRuntimeSsisCatalogPricingTier `json:"catalogPricingTier,omitempty"`
	CatalogServerEndpoint *string                                   `json:"catalogServerEndpoint,omitempty"`
}

var _ json.Unmarshaler = &IntegrationRuntimeSsisCatalogInfo{}

func (s *IntegrationRuntimeSsisCatalogInfo) UnmarshalJSON(bytes []byte) error {
	type alias IntegrationRuntimeSsisCatalogInfo
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into IntegrationRuntimeSsisCatalogInfo: %+v", err)
	}

	s.CatalogAdminUserName = decoded.CatalogAdminUserName
	s.CatalogPricingTier = decoded.CatalogPricingTier
	s.CatalogServerEndpoint = decoded.CatalogServerEndpoint

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IntegrationRuntimeSsisCatalogInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["catalogAdminPassword"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CatalogAdminPassword' for 'IntegrationRuntimeSsisCatalogInfo': %+v", err)
		}
		s.CatalogAdminPassword = impl
	}
	return nil
}
