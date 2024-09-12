package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosSingleSignOnExtension = IosCredentialSingleSignOnExtension{}

type IosCredentialSingleSignOnExtension struct {
	// Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain
	// a maximum of 500 elements.
	Configurations *[]KeyTypedValuePair `json:"configurations,omitempty"`

	// Gets or sets a list of hosts or domain names for which the app extension performs SSO.
	Domains *[]string `json:"domains,omitempty"`

	// Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
	ExtensionIdentifier *string `json:"extensionIdentifier,omitempty"`

	// Gets or sets the case-sensitive realm name for this profile.
	Realm *string `json:"realm,omitempty"`

	// Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
	TeamIdentifier nullable.Type[string] `json:"teamIdentifier,omitempty"`

	// Fields inherited from SingleSignOnExtension

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s IosCredentialSingleSignOnExtension) IosSingleSignOnExtension() BaseIosSingleSignOnExtensionImpl {
	return BaseIosSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s IosCredentialSingleSignOnExtension) SingleSignOnExtension() BaseSingleSignOnExtensionImpl {
	return BaseSingleSignOnExtensionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosCredentialSingleSignOnExtension{}

func (s IosCredentialSingleSignOnExtension) MarshalJSON() ([]byte, error) {
	type wrapper IosCredentialSingleSignOnExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosCredentialSingleSignOnExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosCredentialSingleSignOnExtension: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.iosCredentialSingleSignOnExtension"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosCredentialSingleSignOnExtension: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosCredentialSingleSignOnExtension{}

func (s *IosCredentialSingleSignOnExtension) UnmarshalJSON(bytes []byte) error {
	type alias IosCredentialSingleSignOnExtension
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into IosCredentialSingleSignOnExtension: %+v", err)
	}

	s.Domains = decoded.Domains
	s.ExtensionIdentifier = decoded.ExtensionIdentifier
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Realm = decoded.Realm
	s.TeamIdentifier = decoded.TeamIdentifier

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosCredentialSingleSignOnExtension into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["configurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Configurations into list []json.RawMessage: %+v", err)
		}

		output := make([]KeyTypedValuePair, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalKeyTypedValuePairImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Configurations' for 'IosCredentialSingleSignOnExtension': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Configurations = &output
	}
	return nil
}
