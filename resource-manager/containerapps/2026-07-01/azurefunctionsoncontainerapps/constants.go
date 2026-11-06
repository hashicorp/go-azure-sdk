package azurefunctionsoncontainerapps

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsFunctionState string

const (
	ContainerAppsFunctionStateDisabled ContainerAppsFunctionState = "Disabled"
	ContainerAppsFunctionStateEnabled  ContainerAppsFunctionState = "Enabled"
)

func PossibleValuesForContainerAppsFunctionState() []string {
	return []string{
		string(ContainerAppsFunctionStateDisabled),
		string(ContainerAppsFunctionStateEnabled),
	}
}

func (s *ContainerAppsFunctionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContainerAppsFunctionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContainerAppsFunctionState(input string) (*ContainerAppsFunctionState, error) {
	vals := map[string]ContainerAppsFunctionState{
		"disabled": ContainerAppsFunctionStateDisabled,
		"enabled":  ContainerAppsFunctionStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerAppsFunctionState(input)
	return &out, nil
}
