package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationsIdentitySetEndpointType string

const (
	CommunicationsIdentitySetEndpointType_Default                   CommunicationsIdentitySetEndpointType = "default"
	CommunicationsIdentitySetEndpointType_SkypeForBusiness          CommunicationsIdentitySetEndpointType = "skypeForBusiness"
	CommunicationsIdentitySetEndpointType_SkypeForBusinessVoipPhone CommunicationsIdentitySetEndpointType = "skypeForBusinessVoipPhone"
	CommunicationsIdentitySetEndpointType_Voicemail                 CommunicationsIdentitySetEndpointType = "voicemail"
)

func PossibleValuesForCommunicationsIdentitySetEndpointType() []string {
	return []string{
		string(CommunicationsIdentitySetEndpointType_Default),
		string(CommunicationsIdentitySetEndpointType_SkypeForBusiness),
		string(CommunicationsIdentitySetEndpointType_SkypeForBusinessVoipPhone),
		string(CommunicationsIdentitySetEndpointType_Voicemail),
	}
}

func (s *CommunicationsIdentitySetEndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommunicationsIdentitySetEndpointType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommunicationsIdentitySetEndpointType(input string) (*CommunicationsIdentitySetEndpointType, error) {
	vals := map[string]CommunicationsIdentitySetEndpointType{
		"default":                   CommunicationsIdentitySetEndpointType_Default,
		"skypeforbusiness":          CommunicationsIdentitySetEndpointType_SkypeForBusiness,
		"skypeforbusinessvoipphone": CommunicationsIdentitySetEndpointType_SkypeForBusinessVoipPhone,
		"voicemail":                 CommunicationsIdentitySetEndpointType_Voicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommunicationsIdentitySetEndpointType(input)
	return &out, nil
}
