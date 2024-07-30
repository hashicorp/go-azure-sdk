package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppAppAvailability string

const (
	ManagedAppAppAvailability_Global         ManagedAppAppAvailability = "global"
	ManagedAppAppAvailability_LineOfBusiness ManagedAppAppAvailability = "lineOfBusiness"
)

func PossibleValuesForManagedAppAppAvailability() []string {
	return []string{
		string(ManagedAppAppAvailability_Global),
		string(ManagedAppAppAvailability_LineOfBusiness),
	}
}

func (s *ManagedAppAppAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppAppAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppAppAvailability(input string) (*ManagedAppAppAvailability, error) {
	vals := map[string]ManagedAppAppAvailability{
		"global":         ManagedAppAppAvailability_Global,
		"lineofbusiness": ManagedAppAppAvailability_LineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppAppAvailability(input)
	return &out, nil
}
