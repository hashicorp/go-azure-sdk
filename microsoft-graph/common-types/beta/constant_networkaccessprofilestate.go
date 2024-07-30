package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessProfileState string

const (
	NetworkaccessProfileState_Disabled NetworkaccessProfileState = "disabled"
	NetworkaccessProfileState_Enabled  NetworkaccessProfileState = "enabled"
)

func PossibleValuesForNetworkaccessProfileState() []string {
	return []string{
		string(NetworkaccessProfileState_Disabled),
		string(NetworkaccessProfileState_Enabled),
	}
}

func (s *NetworkaccessProfileState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessProfileState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessProfileState(input string) (*NetworkaccessProfileState, error) {
	vals := map[string]NetworkaccessProfileState{
		"disabled": NetworkaccessProfileState_Disabled,
		"enabled":  NetworkaccessProfileState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessProfileState(input)
	return &out, nil
}
