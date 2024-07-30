package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetailsAdminConfiguration string

const (
	AuthenticationAppPolicyDetailsAdminConfiguration_Disabled      AuthenticationAppPolicyDetailsAdminConfiguration = "disabled"
	AuthenticationAppPolicyDetailsAdminConfiguration_Enabled       AuthenticationAppPolicyDetailsAdminConfiguration = "enabled"
	AuthenticationAppPolicyDetailsAdminConfiguration_NotApplicable AuthenticationAppPolicyDetailsAdminConfiguration = "notApplicable"
)

func PossibleValuesForAuthenticationAppPolicyDetailsAdminConfiguration() []string {
	return []string{
		string(AuthenticationAppPolicyDetailsAdminConfiguration_Disabled),
		string(AuthenticationAppPolicyDetailsAdminConfiguration_Enabled),
		string(AuthenticationAppPolicyDetailsAdminConfiguration_NotApplicable),
	}
}

func (s *AuthenticationAppPolicyDetailsAdminConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAppPolicyDetailsAdminConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAppPolicyDetailsAdminConfiguration(input string) (*AuthenticationAppPolicyDetailsAdminConfiguration, error) {
	vals := map[string]AuthenticationAppPolicyDetailsAdminConfiguration{
		"disabled":      AuthenticationAppPolicyDetailsAdminConfiguration_Disabled,
		"enabled":       AuthenticationAppPolicyDetailsAdminConfiguration_Enabled,
		"notapplicable": AuthenticationAppPolicyDetailsAdminConfiguration_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppPolicyDetailsAdminConfiguration(input)
	return &out, nil
}
