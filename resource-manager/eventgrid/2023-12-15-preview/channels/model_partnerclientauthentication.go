package channels

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerClientAuthentication interface {
	PartnerClientAuthentication() BasePartnerClientAuthenticationImpl
}

var _ PartnerClientAuthentication = BasePartnerClientAuthenticationImpl{}

type BasePartnerClientAuthenticationImpl struct {
	ClientAuthenticationType PartnerClientAuthenticationType `json:"clientAuthenticationType"`
}

func (s BasePartnerClientAuthenticationImpl) PartnerClientAuthentication() BasePartnerClientAuthenticationImpl {
	return s
}

var _ PartnerClientAuthentication = RawPartnerClientAuthenticationImpl{}

// RawPartnerClientAuthenticationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPartnerClientAuthenticationImpl struct {
	partnerClientAuthentication BasePartnerClientAuthenticationImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawPartnerClientAuthenticationImpl) PartnerClientAuthentication() BasePartnerClientAuthenticationImpl {
	return s.partnerClientAuthentication
}

func UnmarshalPartnerClientAuthenticationImplementation(input []byte) (PartnerClientAuthentication, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerClientAuthentication into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["clientAuthenticationType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureAD") {
		var out AzureADPartnerClientAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureADPartnerClientAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BasePartnerClientAuthenticationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePartnerClientAuthenticationImpl: %+v", err)
	}

	return RawPartnerClientAuthenticationImpl{
		partnerClientAuthentication: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
