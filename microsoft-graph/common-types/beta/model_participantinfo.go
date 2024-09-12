package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantInfo struct {
	// The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call.
	// Read-only.
	CountryCode nullable.Type[string] `json:"countryCode,omitempty"`

	// The type of endpoint the participant is using. Possible values are: default, voicemail, skypeForBusiness,
	// skypeForBusinessVoipPhone and unknownFutureValue. Read-only.
	EndpointType *EndpointType `json:"endpointType,omitempty"`

	Identity IdentitySet `json:"identity"`

	// The language culture string. Read-only.
	LanguageId nullable.Type[string] `json:"languageId,omitempty"`

	NonAnonymizedIdentity IdentitySet `json:"nonAnonymizedIdentity"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The participant ID of the participant. Read-only.
	ParticipantId nullable.Type[string] `json:"participantId,omitempty"`

	// The client platform ID of the participant. Read-only.
	PlatformId nullable.Type[string] `json:"platformId,omitempty"`

	// The home region of the participant, and can be a country, a continent, or a larger geographic region. The region
	// doesn't change based on the participant's current physical location, unlike countryCode. Read-only.
	Region nullable.Type[string] `json:"region,omitempty"`
}

var _ json.Marshaler = ParticipantInfo{}

func (s ParticipantInfo) MarshalJSON() ([]byte, error) {
	type wrapper ParticipantInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ParticipantInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ParticipantInfo: %+v", err)
	}

	delete(decoded, "countryCode")
	delete(decoded, "endpointType")
	delete(decoded, "languageId")
	delete(decoded, "participantId")
	delete(decoded, "platformId")
	delete(decoded, "region")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ParticipantInfo: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ParticipantInfo{}

func (s *ParticipantInfo) UnmarshalJSON(bytes []byte) error {
	type alias ParticipantInfo
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ParticipantInfo: %+v", err)
	}

	s.CountryCode = decoded.CountryCode
	s.EndpointType = decoded.EndpointType
	s.LanguageId = decoded.LanguageId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ParticipantId = decoded.ParticipantId
	s.PlatformId = decoded.PlatformId
	s.Region = decoded.Region

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ParticipantInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'ParticipantInfo': %+v", err)
		}
		s.Identity = impl
	}

	if v, ok := temp["nonAnonymizedIdentity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'NonAnonymizedIdentity' for 'ParticipantInfo': %+v", err)
		}
		s.NonAnonymizedIdentity = impl
	}
	return nil
}
