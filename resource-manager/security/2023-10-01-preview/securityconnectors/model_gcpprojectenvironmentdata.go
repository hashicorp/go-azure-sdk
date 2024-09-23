package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnvironmentData = GcpProjectEnvironmentData{}

type GcpProjectEnvironmentData struct {
	OrganizationalData GcpOrganizationalData `json:"organizationalData"`
	ProjectDetails     *GcpProjectDetails    `json:"projectDetails,omitempty"`
	ScanInterval       *int64                `json:"scanInterval,omitempty"`

	// Fields inherited from EnvironmentData

	EnvironmentType EnvironmentType `json:"environmentType"`
}

func (s GcpProjectEnvironmentData) EnvironmentData() BaseEnvironmentDataImpl {
	return BaseEnvironmentDataImpl{
		EnvironmentType: s.EnvironmentType,
	}
}

var _ json.Marshaler = GcpProjectEnvironmentData{}

func (s GcpProjectEnvironmentData) MarshalJSON() ([]byte, error) {
	type wrapper GcpProjectEnvironmentData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpProjectEnvironmentData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpProjectEnvironmentData: %+v", err)
	}

	decoded["environmentType"] = "GcpProject"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpProjectEnvironmentData: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GcpProjectEnvironmentData{}

func (s *GcpProjectEnvironmentData) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ProjectDetails  *GcpProjectDetails `json:"projectDetails,omitempty"`
		ScanInterval    *int64             `json:"scanInterval,omitempty"`
		EnvironmentType EnvironmentType    `json:"environmentType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ProjectDetails = decoded.ProjectDetails
	s.ScanInterval = decoded.ScanInterval
	s.EnvironmentType = decoded.EnvironmentType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GcpProjectEnvironmentData into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["organizationalData"]; ok {
		impl, err := UnmarshalGcpOrganizationalDataImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OrganizationalData' for 'GcpProjectEnvironmentData': %+v", err)
		}
		s.OrganizationalData = impl
	}

	return nil
}
