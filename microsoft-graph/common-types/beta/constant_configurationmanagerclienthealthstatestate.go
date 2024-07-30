package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientHealthStateState string

const (
	ConfigurationManagerClientHealthStateState_CommunicationError ConfigurationManagerClientHealthStateState = "communicationError"
	ConfigurationManagerClientHealthStateState_Healthy            ConfigurationManagerClientHealthStateState = "healthy"
	ConfigurationManagerClientHealthStateState_InstallFailed      ConfigurationManagerClientHealthStateState = "installFailed"
	ConfigurationManagerClientHealthStateState_Installed          ConfigurationManagerClientHealthStateState = "installed"
	ConfigurationManagerClientHealthStateState_Unknown            ConfigurationManagerClientHealthStateState = "unknown"
	ConfigurationManagerClientHealthStateState_UpdateFailed       ConfigurationManagerClientHealthStateState = "updateFailed"
)

func PossibleValuesForConfigurationManagerClientHealthStateState() []string {
	return []string{
		string(ConfigurationManagerClientHealthStateState_CommunicationError),
		string(ConfigurationManagerClientHealthStateState_Healthy),
		string(ConfigurationManagerClientHealthStateState_InstallFailed),
		string(ConfigurationManagerClientHealthStateState_Installed),
		string(ConfigurationManagerClientHealthStateState_Unknown),
		string(ConfigurationManagerClientHealthStateState_UpdateFailed),
	}
}

func (s *ConfigurationManagerClientHealthStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerClientHealthStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerClientHealthStateState(input string) (*ConfigurationManagerClientHealthStateState, error) {
	vals := map[string]ConfigurationManagerClientHealthStateState{
		"communicationerror": ConfigurationManagerClientHealthStateState_CommunicationError,
		"healthy":            ConfigurationManagerClientHealthStateState_Healthy,
		"installfailed":      ConfigurationManagerClientHealthStateState_InstallFailed,
		"installed":          ConfigurationManagerClientHealthStateState_Installed,
		"unknown":            ConfigurationManagerClientHealthStateState_Unknown,
		"updatefailed":       ConfigurationManagerClientHealthStateState_UpdateFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerClientHealthStateState(input)
	return &out, nil
}
