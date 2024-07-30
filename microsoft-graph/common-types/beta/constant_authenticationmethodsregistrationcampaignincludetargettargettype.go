package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType string

const (
	AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType_Group AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType = "group"
	AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType_User  AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType = "user"
)

func PossibleValuesForAuthenticationMethodsRegistrationCampaignIncludeTargetTargetType() []string {
	return []string{
		string(AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType_Group),
		string(AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType_User),
	}
}

func (s *AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodsRegistrationCampaignIncludeTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodsRegistrationCampaignIncludeTargetTargetType(input string) (*AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType, error) {
	vals := map[string]AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType{
		"group": AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType_Group,
		"user":  AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType(input)
	return &out, nil
}
