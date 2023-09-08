package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationsIdentitySetEndpointType string

const (
	CommunicationsIdentitySetEndpointTypedefault                   CommunicationsIdentitySetEndpointType = "Default"
	CommunicationsIdentitySetEndpointTypeskypeForBusiness          CommunicationsIdentitySetEndpointType = "SkypeForBusiness"
	CommunicationsIdentitySetEndpointTypeskypeForBusinessVoipPhone CommunicationsIdentitySetEndpointType = "SkypeForBusinessVoipPhone"
	CommunicationsIdentitySetEndpointTypevoicemail                 CommunicationsIdentitySetEndpointType = "Voicemail"
)

func PossibleValuesForCommunicationsIdentitySetEndpointType() []string {
	return []string{
		string(CommunicationsIdentitySetEndpointTypedefault),
		string(CommunicationsIdentitySetEndpointTypeskypeForBusiness),
		string(CommunicationsIdentitySetEndpointTypeskypeForBusinessVoipPhone),
		string(CommunicationsIdentitySetEndpointTypevoicemail),
	}
}

func parseCommunicationsIdentitySetEndpointType(input string) (*CommunicationsIdentitySetEndpointType, error) {
	vals := map[string]CommunicationsIdentitySetEndpointType{
		"default":                   CommunicationsIdentitySetEndpointTypedefault,
		"skypeforbusiness":          CommunicationsIdentitySetEndpointTypeskypeForBusiness,
		"skypeforbusinessvoipphone": CommunicationsIdentitySetEndpointTypeskypeForBusinessVoipPhone,
		"voicemail":                 CommunicationsIdentitySetEndpointTypevoicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommunicationsIdentitySetEndpointType(input)
	return &out, nil
}
