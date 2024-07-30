package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AssociatedAppsAppType string

const (
	Windows10AssociatedAppsAppType_Desktop   Windows10AssociatedAppsAppType = "desktop"
	Windows10AssociatedAppsAppType_Universal Windows10AssociatedAppsAppType = "universal"
)

func PossibleValuesForWindows10AssociatedAppsAppType() []string {
	return []string{
		string(Windows10AssociatedAppsAppType_Desktop),
		string(Windows10AssociatedAppsAppType_Universal),
	}
}

func (s *Windows10AssociatedAppsAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10AssociatedAppsAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10AssociatedAppsAppType(input string) (*Windows10AssociatedAppsAppType, error) {
	vals := map[string]Windows10AssociatedAppsAppType{
		"desktop":   Windows10AssociatedAppsAppType_Desktop,
		"universal": Windows10AssociatedAppsAppType_Universal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10AssociatedAppsAppType(input)
	return &out, nil
}
