package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeSsisProperties struct {
	CatalogInfo                  *IntegrationRuntimeSsisCatalogInfo             `json:"catalogInfo,omitempty"`
	CustomSetupScriptProperties  *IntegrationRuntimeCustomSetupScriptProperties `json:"customSetupScriptProperties,omitempty"`
	DataProxyProperties          *IntegrationRuntimeDataProxyProperties         `json:"dataProxyProperties,omitempty"`
	Edition                      *IntegrationRuntimeEdition                     `json:"edition,omitempty"`
	ExpressCustomSetupProperties *[]CustomSetupBase                             `json:"expressCustomSetupProperties,omitempty"`
	LicenseType                  *IntegrationRuntimeLicenseType                 `json:"licenseType,omitempty"`
}

var _ json.Unmarshaler = &IntegrationRuntimeSsisProperties{}

func (s *IntegrationRuntimeSsisProperties) UnmarshalJSON(bytes []byte) error {
	type alias IntegrationRuntimeSsisProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into IntegrationRuntimeSsisProperties: %+v", err)
	}

	s.CatalogInfo = decoded.CatalogInfo
	s.CustomSetupScriptProperties = decoded.CustomSetupScriptProperties
	s.DataProxyProperties = decoded.DataProxyProperties
	s.Edition = decoded.Edition
	s.LicenseType = decoded.LicenseType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IntegrationRuntimeSsisProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["expressCustomSetupProperties"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ExpressCustomSetupProperties into list []json.RawMessage: %+v", err)
		}

		output := make([]CustomSetupBase, 0)
		for i, val := range listTemp {
			impl, err := unmarshalCustomSetupBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ExpressCustomSetupProperties' for 'IntegrationRuntimeSsisProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ExpressCustomSetupProperties = &output
	}
	return nil
}
