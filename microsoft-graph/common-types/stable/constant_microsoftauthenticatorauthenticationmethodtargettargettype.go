package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodTargetTargetType string

const (
	MicrosoftAuthenticatorAuthenticationMethodTargetTargetType_Group MicrosoftAuthenticatorAuthenticationMethodTargetTargetType = "group"
	MicrosoftAuthenticatorAuthenticationMethodTargetTargetType_User  MicrosoftAuthenticatorAuthenticationMethodTargetTargetType = "user"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodTargetTargetType() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodTargetTargetType_Group),
		string(MicrosoftAuthenticatorAuthenticationMethodTargetTargetType_User),
	}
}

func (s *MicrosoftAuthenticatorAuthenticationMethodTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftAuthenticatorAuthenticationMethodTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftAuthenticatorAuthenticationMethodTargetTargetType(input string) (*MicrosoftAuthenticatorAuthenticationMethodTargetTargetType, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodTargetTargetType{
		"group": MicrosoftAuthenticatorAuthenticationMethodTargetTargetType_Group,
		"user":  MicrosoftAuthenticatorAuthenticationMethodTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodTargetTargetType(input)
	return &out, nil
}
