package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesFilterMode string

const (
	DevicesFilterMode_Allowed DevicesFilterMode = "allowed"
	DevicesFilterMode_Blocked DevicesFilterMode = "blocked"
)

func PossibleValuesForDevicesFilterMode() []string {
	return []string{
		string(DevicesFilterMode_Allowed),
		string(DevicesFilterMode_Blocked),
	}
}

func (s *DevicesFilterMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDevicesFilterMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDevicesFilterMode(input string) (*DevicesFilterMode, error) {
	vals := map[string]DevicesFilterMode{
		"allowed": DevicesFilterMode_Allowed,
		"blocked": DevicesFilterMode_Blocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DevicesFilterMode(input)
	return &out, nil
}
