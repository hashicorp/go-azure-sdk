package dataflowendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowOpenTelemetryAuthentication = DataflowOpenTelemetryX509CertificateAuthentication{}

type DataflowOpenTelemetryX509CertificateAuthentication struct {
	X509CertificateSettings DataflowEndpointAuthenticationX509 `json:"x509CertificateSettings"`

	// Fields inherited from DataflowOpenTelemetryAuthentication

	Method DataflowOpenTelemetryAuthenticationMethod `json:"method"`
}

func (s DataflowOpenTelemetryX509CertificateAuthentication) DataflowOpenTelemetryAuthentication() BaseDataflowOpenTelemetryAuthenticationImpl {
	return BaseDataflowOpenTelemetryAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = DataflowOpenTelemetryX509CertificateAuthentication{}

func (s DataflowOpenTelemetryX509CertificateAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper DataflowOpenTelemetryX509CertificateAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowOpenTelemetryX509CertificateAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowOpenTelemetryX509CertificateAuthentication: %+v", err)
	}

	decoded["method"] = "X509Certificate"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowOpenTelemetryX509CertificateAuthentication: %+v", err)
	}

	return encoded, nil
}
