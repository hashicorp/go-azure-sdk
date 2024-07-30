package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceAuthenticationMethodTargetTargetType string

const (
	VoiceAuthenticationMethodTargetTargetType_Group VoiceAuthenticationMethodTargetTargetType = "group"
	VoiceAuthenticationMethodTargetTargetType_User  VoiceAuthenticationMethodTargetTargetType = "user"
)

func PossibleValuesForVoiceAuthenticationMethodTargetTargetType() []string {
	return []string{
		string(VoiceAuthenticationMethodTargetTargetType_Group),
		string(VoiceAuthenticationMethodTargetTargetType_User),
	}
}

func (s *VoiceAuthenticationMethodTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVoiceAuthenticationMethodTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVoiceAuthenticationMethodTargetTargetType(input string) (*VoiceAuthenticationMethodTargetTargetType, error) {
	vals := map[string]VoiceAuthenticationMethodTargetTargetType{
		"group": VoiceAuthenticationMethodTargetTargetType_Group,
		"user":  VoiceAuthenticationMethodTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VoiceAuthenticationMethodTargetTargetType(input)
	return &out, nil
}
