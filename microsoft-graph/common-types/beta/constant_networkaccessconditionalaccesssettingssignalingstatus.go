package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConditionalAccessSettingsSignalingStatus string

const (
	NetworkaccessConditionalAccessSettingsSignalingStatus_Disabled NetworkaccessConditionalAccessSettingsSignalingStatus = "disabled"
	NetworkaccessConditionalAccessSettingsSignalingStatus_Enabled  NetworkaccessConditionalAccessSettingsSignalingStatus = "enabled"
)

func PossibleValuesForNetworkaccessConditionalAccessSettingsSignalingStatus() []string {
	return []string{
		string(NetworkaccessConditionalAccessSettingsSignalingStatus_Disabled),
		string(NetworkaccessConditionalAccessSettingsSignalingStatus_Enabled),
	}
}

func (s *NetworkaccessConditionalAccessSettingsSignalingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessConditionalAccessSettingsSignalingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessConditionalAccessSettingsSignalingStatus(input string) (*NetworkaccessConditionalAccessSettingsSignalingStatus, error) {
	vals := map[string]NetworkaccessConditionalAccessSettingsSignalingStatus{
		"disabled": NetworkaccessConditionalAccessSettingsSignalingStatus_Disabled,
		"enabled":  NetworkaccessConditionalAccessSettingsSignalingStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessConditionalAccessSettingsSignalingStatus(input)
	return &out, nil
}
