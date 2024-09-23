package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnvironmentData = AwsEnvironmentData{}

type AwsEnvironmentData struct {
	AccountName        *string               `json:"accountName,omitempty"`
	OrganizationalData AwsOrganizationalData `json:"organizationalData"`
	Regions            *[]string             `json:"regions,omitempty"`
	ScanInterval       *int64                `json:"scanInterval,omitempty"`

	// Fields inherited from EnvironmentData

	EnvironmentType EnvironmentType `json:"environmentType"`
}

func (s AwsEnvironmentData) EnvironmentData() BaseEnvironmentDataImpl {
	return BaseEnvironmentDataImpl{
		EnvironmentType: s.EnvironmentType,
	}
}

var _ json.Marshaler = AwsEnvironmentData{}

func (s AwsEnvironmentData) MarshalJSON() ([]byte, error) {
	type wrapper AwsEnvironmentData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsEnvironmentData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsEnvironmentData: %+v", err)
	}

	decoded["environmentType"] = "AwsAccount"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsEnvironmentData: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AwsEnvironmentData{}

func (s *AwsEnvironmentData) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountName     *string         `json:"accountName,omitempty"`
		Regions         *[]string       `json:"regions,omitempty"`
		ScanInterval    *int64          `json:"scanInterval,omitempty"`
		EnvironmentType EnvironmentType `json:"environmentType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountName = decoded.AccountName
	s.Regions = decoded.Regions
	s.ScanInterval = decoded.ScanInterval
	s.EnvironmentType = decoded.EnvironmentType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AwsEnvironmentData into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["organizationalData"]; ok {
		impl, err := UnmarshalAwsOrganizationalDataImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OrganizationalData' for 'AwsEnvironmentData': %+v", err)
		}
		s.OrganizationalData = impl
	}

	return nil
}
