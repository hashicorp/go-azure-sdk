package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LinkedIntegrationRuntimeType = LinkedIntegrationRuntimeKeyAuthorization{}

type LinkedIntegrationRuntimeKeyAuthorization struct {
	Key SecretBase `json:"key"`

	// Fields inherited from LinkedIntegrationRuntimeType
}

var _ json.Marshaler = LinkedIntegrationRuntimeKeyAuthorization{}

func (s LinkedIntegrationRuntimeKeyAuthorization) MarshalJSON() ([]byte, error) {
	type wrapper LinkedIntegrationRuntimeKeyAuthorization
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LinkedIntegrationRuntimeKeyAuthorization: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LinkedIntegrationRuntimeKeyAuthorization: %+v", err)
	}
	decoded["authorizationType"] = "Key"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LinkedIntegrationRuntimeKeyAuthorization: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LinkedIntegrationRuntimeKeyAuthorization{}

func (s *LinkedIntegrationRuntimeKeyAuthorization) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LinkedIntegrationRuntimeKeyAuthorization into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["key"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Key' for 'LinkedIntegrationRuntimeKeyAuthorization': %+v", err)
		}
		s.Key = impl
	}
	return nil
}
