package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareOathAuthenticationMethodConfigurationState string

const (
	SoftwareOathAuthenticationMethodConfigurationState_Disabled SoftwareOathAuthenticationMethodConfigurationState = "disabled"
	SoftwareOathAuthenticationMethodConfigurationState_Enabled  SoftwareOathAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForSoftwareOathAuthenticationMethodConfigurationState() []string {
	return []string{
		string(SoftwareOathAuthenticationMethodConfigurationState_Disabled),
		string(SoftwareOathAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *SoftwareOathAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSoftwareOathAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSoftwareOathAuthenticationMethodConfigurationState(input string) (*SoftwareOathAuthenticationMethodConfigurationState, error) {
	vals := map[string]SoftwareOathAuthenticationMethodConfigurationState{
		"disabled": SoftwareOathAuthenticationMethodConfigurationState_Disabled,
		"enabled":  SoftwareOathAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SoftwareOathAuthenticationMethodConfigurationState(input)
	return &out, nil
}
