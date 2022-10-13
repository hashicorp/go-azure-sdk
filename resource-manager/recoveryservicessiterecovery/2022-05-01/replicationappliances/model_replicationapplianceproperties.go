package replicationappliances

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationApplianceProperties struct {
	ProviderSpecificDetails ApplianceSpecificDetails `json:"providerSpecificDetails"`
}

var _ json.Unmarshaler = &ReplicationApplianceProperties{}

func (s *ReplicationApplianceProperties) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ReplicationApplianceProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalApplianceSpecificDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'ReplicationApplianceProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
