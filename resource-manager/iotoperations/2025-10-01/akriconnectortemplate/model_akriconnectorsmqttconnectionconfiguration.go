package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorsMqttConnectionConfiguration struct {
	Authentication       AkriConnectorsMqttAuthentication `json:"authentication"`
	Host                 *string                          `json:"host,omitempty"`
	KeepAliveSeconds     *int64                           `json:"keepAliveSeconds,omitempty"`
	MaxInflightMessages  *int64                           `json:"maxInflightMessages,omitempty"`
	Protocol             *AkriConnectorsMqttProtocolType  `json:"protocol,omitempty"`
	SessionExpirySeconds *int64                           `json:"sessionExpirySeconds,omitempty"`
	Tls                  *TlsProperties                   `json:"tls,omitempty"`
}

var _ json.Unmarshaler = &AkriConnectorsMqttConnectionConfiguration{}

func (s *AkriConnectorsMqttConnectionConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Host                 *string                         `json:"host,omitempty"`
		KeepAliveSeconds     *int64                          `json:"keepAliveSeconds,omitempty"`
		MaxInflightMessages  *int64                          `json:"maxInflightMessages,omitempty"`
		Protocol             *AkriConnectorsMqttProtocolType `json:"protocol,omitempty"`
		SessionExpirySeconds *int64                          `json:"sessionExpirySeconds,omitempty"`
		Tls                  *TlsProperties                  `json:"tls,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Host = decoded.Host
	s.KeepAliveSeconds = decoded.KeepAliveSeconds
	s.MaxInflightMessages = decoded.MaxInflightMessages
	s.Protocol = decoded.Protocol
	s.SessionExpirySeconds = decoded.SessionExpirySeconds
	s.Tls = decoded.Tls

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AkriConnectorsMqttConnectionConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authentication"]; ok {
		impl, err := UnmarshalAkriConnectorsMqttAuthenticationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Authentication' for 'AkriConnectorsMqttConnectionConfiguration': %+v", err)
		}
		s.Authentication = impl
	}

	return nil
}
