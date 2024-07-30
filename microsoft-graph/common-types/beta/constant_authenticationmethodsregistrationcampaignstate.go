package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaignState string

const (
	AuthenticationMethodsRegistrationCampaignState_Default  AuthenticationMethodsRegistrationCampaignState = "default"
	AuthenticationMethodsRegistrationCampaignState_Disabled AuthenticationMethodsRegistrationCampaignState = "disabled"
	AuthenticationMethodsRegistrationCampaignState_Enabled  AuthenticationMethodsRegistrationCampaignState = "enabled"
)

func PossibleValuesForAuthenticationMethodsRegistrationCampaignState() []string {
	return []string{
		string(AuthenticationMethodsRegistrationCampaignState_Default),
		string(AuthenticationMethodsRegistrationCampaignState_Disabled),
		string(AuthenticationMethodsRegistrationCampaignState_Enabled),
	}
}

func (s *AuthenticationMethodsRegistrationCampaignState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodsRegistrationCampaignState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodsRegistrationCampaignState(input string) (*AuthenticationMethodsRegistrationCampaignState, error) {
	vals := map[string]AuthenticationMethodsRegistrationCampaignState{
		"default":  AuthenticationMethodsRegistrationCampaignState_Default,
		"disabled": AuthenticationMethodsRegistrationCampaignState_Disabled,
		"enabled":  AuthenticationMethodsRegistrationCampaignState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodsRegistrationCampaignState(input)
	return &out, nil
}
