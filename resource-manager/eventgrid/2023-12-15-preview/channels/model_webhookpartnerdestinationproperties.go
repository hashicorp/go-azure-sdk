package channels

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebhookPartnerDestinationProperties struct {
	ClientAuthentication PartnerClientAuthentication `json:"clientAuthentication"`
	EndpointBaseUrl      *string                     `json:"endpointBaseUrl,omitempty"`
	EndpointUrl          *string                     `json:"endpointUrl,omitempty"`
}

var _ json.Unmarshaler = &WebhookPartnerDestinationProperties{}

func (s *WebhookPartnerDestinationProperties) UnmarshalJSON(bytes []byte) error {
	type alias WebhookPartnerDestinationProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into WebhookPartnerDestinationProperties: %+v", err)
	}

	s.EndpointBaseUrl = decoded.EndpointBaseUrl
	s.EndpointUrl = decoded.EndpointUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WebhookPartnerDestinationProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["clientAuthentication"]; ok {
		impl, err := unmarshalPartnerClientAuthenticationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ClientAuthentication' for 'WebhookPartnerDestinationProperties': %+v", err)
		}
		s.ClientAuthentication = impl
	}
	return nil
}
