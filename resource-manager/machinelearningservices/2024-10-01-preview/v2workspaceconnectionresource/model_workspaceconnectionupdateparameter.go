package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceConnectionUpdateParameter struct {
	Properties WorkspaceConnectionPropertiesV2 `json:"properties"`
}

var _ json.Unmarshaler = &WorkspaceConnectionUpdateParameter{}

func (s *WorkspaceConnectionUpdateParameter) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WorkspaceConnectionUpdateParameter into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalWorkspaceConnectionPropertiesV2Implementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'WorkspaceConnectionUpdateParameter': %+v", err)
		}
		s.Properties = impl
	}

	return nil
}
