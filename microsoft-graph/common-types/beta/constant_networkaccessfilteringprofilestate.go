package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringProfileState string

const (
	NetworkaccessFilteringProfileState_Disabled NetworkaccessFilteringProfileState = "disabled"
	NetworkaccessFilteringProfileState_Enabled  NetworkaccessFilteringProfileState = "enabled"
)

func PossibleValuesForNetworkaccessFilteringProfileState() []string {
	return []string{
		string(NetworkaccessFilteringProfileState_Disabled),
		string(NetworkaccessFilteringProfileState_Enabled),
	}
}

func (s *NetworkaccessFilteringProfileState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessFilteringProfileState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessFilteringProfileState(input string) (*NetworkaccessFilteringProfileState, error) {
	vals := map[string]NetworkaccessFilteringProfileState{
		"disabled": NetworkaccessFilteringProfileState_Disabled,
		"enabled":  NetworkaccessFilteringProfileState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessFilteringProfileState(input)
	return &out, nil
}
