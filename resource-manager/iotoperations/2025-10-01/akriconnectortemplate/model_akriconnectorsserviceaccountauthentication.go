package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AkriConnectorsMqttAuthentication = AkriConnectorsServiceAccountAuthentication{}

type AkriConnectorsServiceAccountAuthentication struct {
	ServiceAccountTokenSettings AkriConnectorsServiceAccountTokenSettings `json:"serviceAccountTokenSettings"`

	// Fields inherited from AkriConnectorsMqttAuthentication

	Method AkriConnectorsMqttAuthenticationMethod `json:"method"`
}

func (s AkriConnectorsServiceAccountAuthentication) AkriConnectorsMqttAuthentication() BaseAkriConnectorsMqttAuthenticationImpl {
	return BaseAkriConnectorsMqttAuthenticationImpl{
		Method: s.Method,
	}
}

var _ json.Marshaler = AkriConnectorsServiceAccountAuthentication{}

func (s AkriConnectorsServiceAccountAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper AkriConnectorsServiceAccountAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AkriConnectorsServiceAccountAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsServiceAccountAuthentication: %+v", err)
	}

	decoded["method"] = "ServiceAccountToken"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AkriConnectorsServiceAccountAuthentication: %+v", err)
	}

	return encoded, nil
}
