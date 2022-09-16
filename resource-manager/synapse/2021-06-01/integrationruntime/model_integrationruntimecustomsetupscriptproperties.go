package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeCustomSetupScriptProperties struct {
	BlobContainerUri *string    `json:"blobContainerUri,omitempty"`
	SasToken         SecretBase `json:"sasToken"`
}

var _ json.Unmarshaler = &IntegrationRuntimeCustomSetupScriptProperties{}

func (s *IntegrationRuntimeCustomSetupScriptProperties) UnmarshalJSON(bytes []byte) error {
	type alias IntegrationRuntimeCustomSetupScriptProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into IntegrationRuntimeCustomSetupScriptProperties: %+v", err)
	}

	s.BlobContainerUri = decoded.BlobContainerUri

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IntegrationRuntimeCustomSetupScriptProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["sasToken"]; ok {
		impl, err := unmarshalSecretBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SasToken' for 'IntegrationRuntimeCustomSetupScriptProperties': %+v", err)
		}
		s.SasToken = impl
	}
	return nil
}
