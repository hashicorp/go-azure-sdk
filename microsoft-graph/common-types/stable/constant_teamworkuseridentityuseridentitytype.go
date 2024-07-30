package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkUserIdentityUserIdentityType string

const (
	TeamworkUserIdentityUserIdentityType_AadUser                      TeamworkUserIdentityUserIdentityType = "aadUser"
	TeamworkUserIdentityUserIdentityType_AnonymousGuest               TeamworkUserIdentityUserIdentityType = "anonymousGuest"
	TeamworkUserIdentityUserIdentityType_EmailUser                    TeamworkUserIdentityUserIdentityType = "emailUser"
	TeamworkUserIdentityUserIdentityType_FederatedUser                TeamworkUserIdentityUserIdentityType = "federatedUser"
	TeamworkUserIdentityUserIdentityType_OnPremiseAadUser             TeamworkUserIdentityUserIdentityType = "onPremiseAadUser"
	TeamworkUserIdentityUserIdentityType_PersonalMicrosoftAccountUser TeamworkUserIdentityUserIdentityType = "personalMicrosoftAccountUser"
	TeamworkUserIdentityUserIdentityType_PhoneUser                    TeamworkUserIdentityUserIdentityType = "phoneUser"
	TeamworkUserIdentityUserIdentityType_SkypeUser                    TeamworkUserIdentityUserIdentityType = "skypeUser"
)

func PossibleValuesForTeamworkUserIdentityUserIdentityType() []string {
	return []string{
		string(TeamworkUserIdentityUserIdentityType_AadUser),
		string(TeamworkUserIdentityUserIdentityType_AnonymousGuest),
		string(TeamworkUserIdentityUserIdentityType_EmailUser),
		string(TeamworkUserIdentityUserIdentityType_FederatedUser),
		string(TeamworkUserIdentityUserIdentityType_OnPremiseAadUser),
		string(TeamworkUserIdentityUserIdentityType_PersonalMicrosoftAccountUser),
		string(TeamworkUserIdentityUserIdentityType_PhoneUser),
		string(TeamworkUserIdentityUserIdentityType_SkypeUser),
	}
}

func (s *TeamworkUserIdentityUserIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkUserIdentityUserIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkUserIdentityUserIdentityType(input string) (*TeamworkUserIdentityUserIdentityType, error) {
	vals := map[string]TeamworkUserIdentityUserIdentityType{
		"aaduser":                      TeamworkUserIdentityUserIdentityType_AadUser,
		"anonymousguest":               TeamworkUserIdentityUserIdentityType_AnonymousGuest,
		"emailuser":                    TeamworkUserIdentityUserIdentityType_EmailUser,
		"federateduser":                TeamworkUserIdentityUserIdentityType_FederatedUser,
		"onpremiseaaduser":             TeamworkUserIdentityUserIdentityType_OnPremiseAadUser,
		"personalmicrosoftaccountuser": TeamworkUserIdentityUserIdentityType_PersonalMicrosoftAccountUser,
		"phoneuser":                    TeamworkUserIdentityUserIdentityType_PhoneUser,
		"skypeuser":                    TeamworkUserIdentityUserIdentityType_SkypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkUserIdentityUserIdentityType(input)
	return &out, nil
}
