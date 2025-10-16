package dataflowendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowOpenTelemetryAuthentication = DataflowOpenTelemetryServiceAccountAuthentication{}

type DataflowOpenTelemetryServiceAccountAuthentication struct {
	ServiceAccountTokenSettings DataflowEndpointAuthenticationServiceAccountToken `json:"serviceAccountTokenSettings"`

	// Fields inherited from DataflowOpenTelemetryAuthentication

	Method DataflowOpenTelemetryAuthenticationMethod `json:"method"`
}

func (s DataflowOpenTelemetryServiceAccountAuthentication) DataflowOpenTelemetryAuthentication() BaseDataflowOpenTelemetryAuthenticationImpl {
	return BaseDataflowOpenTelemetryAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = DataflowOpenTelemetryServiceAccountAuthentication{}

func (s DataflowOpenTelemetryServiceAccountAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper DataflowOpenTelemetryServiceAccountAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowOpenTelemetryServiceAccountAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowOpenTelemetryServiceAccountAuthentication: %+v", err)
	}

	decoded["method"] = "ServiceAccountToken"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowOpenTelemetryServiceAccountAuthentication: %+v", err)
	}

	return encoded, nil
}
