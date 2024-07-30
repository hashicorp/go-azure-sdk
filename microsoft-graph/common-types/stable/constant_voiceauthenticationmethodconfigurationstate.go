package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceAuthenticationMethodConfigurationState string

const (
	VoiceAuthenticationMethodConfigurationState_Disabled VoiceAuthenticationMethodConfigurationState = "disabled"
	VoiceAuthenticationMethodConfigurationState_Enabled  VoiceAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForVoiceAuthenticationMethodConfigurationState() []string {
	return []string{
		string(VoiceAuthenticationMethodConfigurationState_Disabled),
		string(VoiceAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *VoiceAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVoiceAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVoiceAuthenticationMethodConfigurationState(input string) (*VoiceAuthenticationMethodConfigurationState, error) {
	vals := map[string]VoiceAuthenticationMethodConfigurationState{
		"disabled": VoiceAuthenticationMethodConfigurationState_Disabled,
		"enabled":  VoiceAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VoiceAuthenticationMethodConfigurationState(input)
	return &out, nil
}
