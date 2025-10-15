package dataflowendpoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowOpenTelemetryAuthentication interface {
	DataflowOpenTelemetryAuthentication() BaseDataflowOpenTelemetryAuthenticationImpl
}

var _ DataflowOpenTelemetryAuthentication = BaseDataflowOpenTelemetryAuthenticationImpl{}

type BaseDataflowOpenTelemetryAuthenticationImpl struct {
	Method DataflowOpenTelemetryAuthenticationMethod `json:"method"`
}

func (s BaseDataflowOpenTelemetryAuthenticationImpl) DataflowOpenTelemetryAuthentication() BaseDataflowOpenTelemetryAuthenticationImpl {
	return s
}

var _ DataflowOpenTelemetryAuthentication = RawDataflowOpenTelemetryAuthenticationImpl{}

// RawDataflowOpenTelemetryAuthenticationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataflowOpenTelemetryAuthenticationImpl struct {
	dataflowOpenTelemetryAuthentication BaseDataflowOpenTelemetryAuthenticationImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawDataflowOpenTelemetryAuthenticationImpl) DataflowOpenTelemetryAuthentication() BaseDataflowOpenTelemetryAuthenticationImpl {
	return s.dataflowOpenTelemetryAuthentication
}

func UnmarshalDataflowOpenTelemetryAuthenticationImplementation(input []byte) (DataflowOpenTelemetryAuthentication, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowOpenTelemetryAuthentication into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["method"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Anonymous") {
		var out DataflowOpenTelemetryAnonymousAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowOpenTelemetryAnonymousAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ServiceAccountToken") {
		var out DataflowOpenTelemetryServiceAccountAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowOpenTelemetryServiceAccountAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "X509Certificate") {
		var out DataflowOpenTelemetryX509CertificateAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowOpenTelemetryX509CertificateAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataflowOpenTelemetryAuthenticationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataflowOpenTelemetryAuthenticationImpl: %+v", err)
	}

	return RawDataflowOpenTelemetryAuthenticationImpl{
		dataflowOpenTelemetryAuthentication: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
