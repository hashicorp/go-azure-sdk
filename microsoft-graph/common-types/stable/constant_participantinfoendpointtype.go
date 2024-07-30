package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantInfoEndpointType string

const (
	ParticipantInfoEndpointType_Default                   ParticipantInfoEndpointType = "default"
	ParticipantInfoEndpointType_SkypeForBusiness          ParticipantInfoEndpointType = "skypeForBusiness"
	ParticipantInfoEndpointType_SkypeForBusinessVoipPhone ParticipantInfoEndpointType = "skypeForBusinessVoipPhone"
	ParticipantInfoEndpointType_Voicemail                 ParticipantInfoEndpointType = "voicemail"
)

func PossibleValuesForParticipantInfoEndpointType() []string {
	return []string{
		string(ParticipantInfoEndpointType_Default),
		string(ParticipantInfoEndpointType_SkypeForBusiness),
		string(ParticipantInfoEndpointType_SkypeForBusinessVoipPhone),
		string(ParticipantInfoEndpointType_Voicemail),
	}
}

func (s *ParticipantInfoEndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseParticipantInfoEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseParticipantInfoEndpointType(input string) (*ParticipantInfoEndpointType, error) {
	vals := map[string]ParticipantInfoEndpointType{
		"default":                   ParticipantInfoEndpointType_Default,
		"skypeforbusiness":          ParticipantInfoEndpointType_SkypeForBusiness,
		"skypeforbusinessvoipphone": ParticipantInfoEndpointType_SkypeForBusinessVoipPhone,
		"voicemail":                 ParticipantInfoEndpointType_Voicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParticipantInfoEndpointType(input)
	return &out, nil
}
