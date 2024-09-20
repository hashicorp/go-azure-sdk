package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitorComputeConfigurationBase = MonitorServerlessSparkCompute{}

type MonitorServerlessSparkCompute struct {
	ComputeIdentity MonitorComputeIdentityBase `json:"computeIdentity"`
	InstanceType    string                     `json:"instanceType"`
	RuntimeVersion  string                     `json:"runtimeVersion"`

	// Fields inherited from MonitorComputeConfigurationBase

	ComputeType MonitorComputeType `json:"computeType"`
}

func (s MonitorServerlessSparkCompute) MonitorComputeConfigurationBase() BaseMonitorComputeConfigurationBaseImpl {
	return BaseMonitorComputeConfigurationBaseImpl{
		ComputeType: s.ComputeType,
	}
}

var _ json.Marshaler = MonitorServerlessSparkCompute{}

func (s MonitorServerlessSparkCompute) MarshalJSON() ([]byte, error) {
	type wrapper MonitorServerlessSparkCompute
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MonitorServerlessSparkCompute: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MonitorServerlessSparkCompute: %+v", err)
	}

	decoded["computeType"] = "ServerlessSpark"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MonitorServerlessSparkCompute: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MonitorServerlessSparkCompute{}

func (s *MonitorServerlessSparkCompute) UnmarshalJSON(bytes []byte) error {
	type alias MonitorServerlessSparkCompute
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MonitorServerlessSparkCompute: %+v", err)
	}

	s.ComputeType = decoded.ComputeType
	s.InstanceType = decoded.InstanceType
	s.RuntimeVersion = decoded.RuntimeVersion

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MonitorServerlessSparkCompute into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["computeIdentity"]; ok {
		impl, err := UnmarshalMonitorComputeIdentityBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ComputeIdentity' for 'MonitorServerlessSparkCompute': %+v", err)
		}
		s.ComputeIdentity = impl
	}
	return nil
}
