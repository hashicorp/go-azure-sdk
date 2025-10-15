package dataflowendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowEndpointOpenTelemetry struct {
	Authentication DataflowOpenTelemetryAuthentication `json:"authentication"`
	Batching       *BatchingConfiguration              `json:"batching,omitempty"`
	Host           string                              `json:"host"`
	Tls            *TlsProperties                      `json:"tls,omitempty"`
}

var _ json.Unmarshaler = &DataflowEndpointOpenTelemetry{}

func (s *DataflowEndpointOpenTelemetry) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Batching *BatchingConfiguration `json:"batching,omitempty"`
		Host     string                 `json:"host"`
		Tls      *TlsProperties         `json:"tls,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Batching = decoded.Batching
	s.Host = decoded.Host
	s.Tls = decoded.Tls

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataflowEndpointOpenTelemetry into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authentication"]; ok {
		impl, err := UnmarshalDataflowOpenTelemetryAuthenticationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Authentication' for 'DataflowEndpointOpenTelemetry': %+v", err)
		}
		s.Authentication = impl
	}

	return nil
}
