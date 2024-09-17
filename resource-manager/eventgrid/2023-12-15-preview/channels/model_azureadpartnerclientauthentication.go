package channels

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartnerClientAuthentication = AzureADPartnerClientAuthentication{}

type AzureADPartnerClientAuthentication struct {
	Properties *AzureADPartnerClientAuthenticationProperties `json:"properties,omitempty"`

	// Fields inherited from PartnerClientAuthentication

	ClientAuthenticationType PartnerClientAuthenticationType `json:"clientAuthenticationType"`
}

func (s AzureADPartnerClientAuthentication) PartnerClientAuthentication() BasePartnerClientAuthenticationImpl {
	return BasePartnerClientAuthenticationImpl{
		ClientAuthenticationType: s.ClientAuthenticationType,
	}
}

var _ json.Marshaler = AzureADPartnerClientAuthentication{}

func (s AzureADPartnerClientAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper AzureADPartnerClientAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureADPartnerClientAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureADPartnerClientAuthentication: %+v", err)
	}

	decoded["clientAuthenticationType"] = "AzureAD"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureADPartnerClientAuthentication: %+v", err)
	}

	return encoded, nil
}
