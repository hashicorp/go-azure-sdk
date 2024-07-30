package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistanceSettingsRemoteAssistanceState string

const (
	RemoteAssistanceSettingsRemoteAssistanceState_Disabled RemoteAssistanceSettingsRemoteAssistanceState = "disabled"
	RemoteAssistanceSettingsRemoteAssistanceState_Enabled  RemoteAssistanceSettingsRemoteAssistanceState = "enabled"
)

func PossibleValuesForRemoteAssistanceSettingsRemoteAssistanceState() []string {
	return []string{
		string(RemoteAssistanceSettingsRemoteAssistanceState_Disabled),
		string(RemoteAssistanceSettingsRemoteAssistanceState_Enabled),
	}
}

func (s *RemoteAssistanceSettingsRemoteAssistanceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteAssistanceSettingsRemoteAssistanceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteAssistanceSettingsRemoteAssistanceState(input string) (*RemoteAssistanceSettingsRemoteAssistanceState, error) {
	vals := map[string]RemoteAssistanceSettingsRemoteAssistanceState{
		"disabled": RemoteAssistanceSettingsRemoteAssistanceState_Disabled,
		"enabled":  RemoteAssistanceSettingsRemoteAssistanceState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteAssistanceSettingsRemoteAssistanceState(input)
	return &out, nil
}
