package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestMigrateInputProperties struct {
	ProviderSpecificDetails TestMigrateProviderSpecificInput `json:"providerSpecificDetails"`
}

var _ json.Unmarshaler = &TestMigrateInputProperties{}

func (s *TestMigrateInputProperties) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TestMigrateInputProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalTestMigrateProviderSpecificInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'TestMigrateInputProperties': %+v", err)
		}
		s.ProviderSpecificDetails = &impl
	}
	return nil
}
