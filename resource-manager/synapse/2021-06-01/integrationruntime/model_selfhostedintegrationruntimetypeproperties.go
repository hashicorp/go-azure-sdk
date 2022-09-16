package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfHostedIntegrationRuntimeTypeProperties struct {
	LinkedInfo LinkedIntegrationRuntimeType `json:"linkedInfo"`
}

var _ json.Unmarshaler = &SelfHostedIntegrationRuntimeTypeProperties{}

func (s *SelfHostedIntegrationRuntimeTypeProperties) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SelfHostedIntegrationRuntimeTypeProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["linkedInfo"]; ok {
		impl, err := unmarshalLinkedIntegrationRuntimeTypeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LinkedInfo' for 'SelfHostedIntegrationRuntimeTypeProperties': %+v", err)
		}
		s.LinkedInfo = impl
	}
	return nil
}
