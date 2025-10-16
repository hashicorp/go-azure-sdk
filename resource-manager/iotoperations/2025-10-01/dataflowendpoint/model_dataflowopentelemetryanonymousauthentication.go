package dataflowendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowOpenTelemetryAuthentication = DataflowOpenTelemetryAnonymousAuthentication{}

type DataflowOpenTelemetryAnonymousAuthentication struct {
	AnonymousSettings interface{} `json:"anonymousSettings"`

	// Fields inherited from DataflowOpenTelemetryAuthentication

	Method DataflowOpenTelemetryAuthenticationMethod `json:"method"`
}

func (s DataflowOpenTelemetryAnonymousAuthentication) DataflowOpenTelemetryAuthentication() BaseDataflowOpenTelemetryAuthenticationImpl {
	return BaseDataflowOpenTelemetryAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = DataflowOpenTelemetryAnonymousAuthentication{}

func (s DataflowOpenTelemetryAnonymousAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper DataflowOpenTelemetryAnonymousAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowOpenTelemetryAnonymousAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowOpenTelemetryAnonymousAuthentication: %+v", err)
	}

	decoded["method"] = "Anonymous"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowOpenTelemetryAnonymousAuthentication: %+v", err)
	}

	return encoded, nil
}
