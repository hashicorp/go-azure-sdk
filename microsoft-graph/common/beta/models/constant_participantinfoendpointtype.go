package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantInfoEndpointType string

const (
	ParticipantInfoEndpointTypedefault                   ParticipantInfoEndpointType = "Default"
	ParticipantInfoEndpointTypeskypeForBusiness          ParticipantInfoEndpointType = "SkypeForBusiness"
	ParticipantInfoEndpointTypeskypeForBusinessVoipPhone ParticipantInfoEndpointType = "SkypeForBusinessVoipPhone"
	ParticipantInfoEndpointTypevoicemail                 ParticipantInfoEndpointType = "Voicemail"
)

func PossibleValuesForParticipantInfoEndpointType() []string {
	return []string{
		string(ParticipantInfoEndpointTypedefault),
		string(ParticipantInfoEndpointTypeskypeForBusiness),
		string(ParticipantInfoEndpointTypeskypeForBusinessVoipPhone),
		string(ParticipantInfoEndpointTypevoicemail),
	}
}

func parseParticipantInfoEndpointType(input string) (*ParticipantInfoEndpointType, error) {
	vals := map[string]ParticipantInfoEndpointType{
		"default":                   ParticipantInfoEndpointTypedefault,
		"skypeforbusiness":          ParticipantInfoEndpointTypeskypeForBusiness,
		"skypeforbusinessvoipphone": ParticipantInfoEndpointTypeskypeForBusinessVoipPhone,
		"voicemail":                 ParticipantInfoEndpointTypevoicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ParticipantInfoEndpointType(input)
	return &out, nil
}
