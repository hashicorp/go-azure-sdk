package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowStagingInfo struct {
	FolderPath    *interface{}            `json:"folderPath,omitempty"`
	LinkedService *LinkedServiceReference `json:"linkedService,omitempty"`
}

var _ json.Unmarshaler = &DataFlowStagingInfo{}

func (s *DataFlowStagingInfo) UnmarshalJSON(bytes []byte) error {
	type alias DataFlowStagingInfo
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DataFlowStagingInfo: %+v", err)
	}

	s.FolderPath = decoded.FolderPath

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataFlowStagingInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["linkedService"]; ok {
		impl, err := unmarshalLinkedServiceReferenceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LinkedService' for 'DataFlowStagingInfo': %+v", err)
		}
		s.LinkedService = impl
	}
	return nil
}
