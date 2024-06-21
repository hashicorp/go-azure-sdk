package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConnectorProperties struct {
	EnvironmentData                 EnvironmentData  `json:"environmentData"`
	EnvironmentName                 *CloudName       `json:"environmentName,omitempty"`
	HierarchyIdentifier             *string          `json:"hierarchyIdentifier,omitempty"`
	HierarchyIdentifierTrialEndDate *string          `json:"hierarchyIdentifierTrialEndDate,omitempty"`
	Offerings                       *[]CloudOffering `json:"offerings,omitempty"`
}

var _ json.Unmarshaler = &SecurityConnectorProperties{}

func (s *SecurityConnectorProperties) UnmarshalJSON(bytes []byte) error {
	type alias SecurityConnectorProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SecurityConnectorProperties: %+v", err)
	}

	s.EnvironmentName = decoded.EnvironmentName
	s.HierarchyIdentifier = decoded.HierarchyIdentifier
	s.HierarchyIdentifierTrialEndDate = decoded.HierarchyIdentifierTrialEndDate

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityConnectorProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["environmentData"]; ok {
		impl, err := unmarshalEnvironmentDataImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EnvironmentData' for 'SecurityConnectorProperties': %+v", err)
		}
		s.EnvironmentData = impl
	}

	if v, ok := temp["offerings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Offerings into list []json.RawMessage: %+v", err)
		}

		output := make([]CloudOffering, 0)
		for i, val := range listTemp {
			impl, err := unmarshalCloudOfferingImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Offerings' for 'SecurityConnectorProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Offerings = &output
	}
	return nil
}
