package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationParticipantInfoEndpointType string

const (
	InvitationParticipantInfoEndpointTypedefault                   InvitationParticipantInfoEndpointType = "Default"
	InvitationParticipantInfoEndpointTypeskypeForBusiness          InvitationParticipantInfoEndpointType = "SkypeForBusiness"
	InvitationParticipantInfoEndpointTypeskypeForBusinessVoipPhone InvitationParticipantInfoEndpointType = "SkypeForBusinessVoipPhone"
	InvitationParticipantInfoEndpointTypevoicemail                 InvitationParticipantInfoEndpointType = "Voicemail"
)

func PossibleValuesForInvitationParticipantInfoEndpointType() []string {
	return []string{
		string(InvitationParticipantInfoEndpointTypedefault),
		string(InvitationParticipantInfoEndpointTypeskypeForBusiness),
		string(InvitationParticipantInfoEndpointTypeskypeForBusinessVoipPhone),
		string(InvitationParticipantInfoEndpointTypevoicemail),
	}
}

func parseInvitationParticipantInfoEndpointType(input string) (*InvitationParticipantInfoEndpointType, error) {
	vals := map[string]InvitationParticipantInfoEndpointType{
		"default":                   InvitationParticipantInfoEndpointTypedefault,
		"skypeforbusiness":          InvitationParticipantInfoEndpointTypeskypeForBusiness,
		"skypeforbusinessvoipphone": InvitationParticipantInfoEndpointTypeskypeForBusinessVoipPhone,
		"voicemail":                 InvitationParticipantInfoEndpointTypevoicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvitationParticipantInfoEndpointType(input)
	return &out, nil
}
