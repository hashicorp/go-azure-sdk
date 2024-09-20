package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateProtectionIntentProperties struct {
	ProviderSpecificDetails CreateProtectionIntentProviderSpecificDetails `json:"providerSpecificDetails"`
}

var _ json.Unmarshaler = &CreateProtectionIntentProperties{}

func (s *CreateProtectionIntentProperties) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CreateProtectionIntentProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := UnmarshalCreateProtectionIntentProviderSpecificDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'CreateProtectionIntentProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
