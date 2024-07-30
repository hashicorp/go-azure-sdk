package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationParticipantInfoEndpointType string

const (
	InvitationParticipantInfoEndpointType_Default                   InvitationParticipantInfoEndpointType = "default"
	InvitationParticipantInfoEndpointType_SkypeForBusiness          InvitationParticipantInfoEndpointType = "skypeForBusiness"
	InvitationParticipantInfoEndpointType_SkypeForBusinessVoipPhone InvitationParticipantInfoEndpointType = "skypeForBusinessVoipPhone"
	InvitationParticipantInfoEndpointType_Voicemail                 InvitationParticipantInfoEndpointType = "voicemail"
)

func PossibleValuesForInvitationParticipantInfoEndpointType() []string {
	return []string{
		string(InvitationParticipantInfoEndpointType_Default),
		string(InvitationParticipantInfoEndpointType_SkypeForBusiness),
		string(InvitationParticipantInfoEndpointType_SkypeForBusinessVoipPhone),
		string(InvitationParticipantInfoEndpointType_Voicemail),
	}
}

func (s *InvitationParticipantInfoEndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvitationParticipantInfoEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvitationParticipantInfoEndpointType(input string) (*InvitationParticipantInfoEndpointType, error) {
	vals := map[string]InvitationParticipantInfoEndpointType{
		"default":                   InvitationParticipantInfoEndpointType_Default,
		"skypeforbusiness":          InvitationParticipantInfoEndpointType_SkypeForBusiness,
		"skypeforbusinessvoipphone": InvitationParticipantInfoEndpointType_SkypeForBusinessVoipPhone,
		"voicemail":                 InvitationParticipantInfoEndpointType_Voicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvitationParticipantInfoEndpointType(input)
	return &out, nil
}
