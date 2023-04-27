package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableMigrationInputProperties struct {
	PolicyId                string                               `json:"policyId"`
	ProviderSpecificDetails EnableMigrationProviderSpecificInput `json:"providerSpecificDetails"`
}

var _ json.Unmarshaler = &EnableMigrationInputProperties{}

func (s *EnableMigrationInputProperties) UnmarshalJSON(bytes []byte) error {
	type alias EnableMigrationInputProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EnableMigrationInputProperties: %+v", err)
	}

	s.PolicyId = decoded.PolicyId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EnableMigrationInputProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalEnableMigrationProviderSpecificInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'EnableMigrationInputProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
