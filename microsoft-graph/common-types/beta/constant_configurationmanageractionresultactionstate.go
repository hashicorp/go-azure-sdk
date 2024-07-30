package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionResultActionState string

const (
	ConfigurationManagerActionResultActionState_Active       ConfigurationManagerActionResultActionState = "active"
	ConfigurationManagerActionResultActionState_Canceled     ConfigurationManagerActionResultActionState = "canceled"
	ConfigurationManagerActionResultActionState_Done         ConfigurationManagerActionResultActionState = "done"
	ConfigurationManagerActionResultActionState_Failed       ConfigurationManagerActionResultActionState = "failed"
	ConfigurationManagerActionResultActionState_None         ConfigurationManagerActionResultActionState = "none"
	ConfigurationManagerActionResultActionState_NotSupported ConfigurationManagerActionResultActionState = "notSupported"
	ConfigurationManagerActionResultActionState_Pending      ConfigurationManagerActionResultActionState = "pending"
)

func PossibleValuesForConfigurationManagerActionResultActionState() []string {
	return []string{
		string(ConfigurationManagerActionResultActionState_Active),
		string(ConfigurationManagerActionResultActionState_Canceled),
		string(ConfigurationManagerActionResultActionState_Done),
		string(ConfigurationManagerActionResultActionState_Failed),
		string(ConfigurationManagerActionResultActionState_None),
		string(ConfigurationManagerActionResultActionState_NotSupported),
		string(ConfigurationManagerActionResultActionState_Pending),
	}
}

func (s *ConfigurationManagerActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerActionResultActionState(input string) (*ConfigurationManagerActionResultActionState, error) {
	vals := map[string]ConfigurationManagerActionResultActionState{
		"active":       ConfigurationManagerActionResultActionState_Active,
		"canceled":     ConfigurationManagerActionResultActionState_Canceled,
		"done":         ConfigurationManagerActionResultActionState_Done,
		"failed":       ConfigurationManagerActionResultActionState_Failed,
		"none":         ConfigurationManagerActionResultActionState_None,
		"notsupported": ConfigurationManagerActionResultActionState_NotSupported,
		"pending":      ConfigurationManagerActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionResultActionState(input)
	return &out, nil
}
