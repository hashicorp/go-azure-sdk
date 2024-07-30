package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedMobileLobAppAppAvailability string

const (
	ManagedMobileLobAppAppAvailability_Global         ManagedMobileLobAppAppAvailability = "global"
	ManagedMobileLobAppAppAvailability_LineOfBusiness ManagedMobileLobAppAppAvailability = "lineOfBusiness"
)

func PossibleValuesForManagedMobileLobAppAppAvailability() []string {
	return []string{
		string(ManagedMobileLobAppAppAvailability_Global),
		string(ManagedMobileLobAppAppAvailability_LineOfBusiness),
	}
}

func (s *ManagedMobileLobAppAppAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedMobileLobAppAppAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedMobileLobAppAppAvailability(input string) (*ManagedMobileLobAppAppAvailability, error) {
	vals := map[string]ManagedMobileLobAppAppAvailability{
		"global":         ManagedMobileLobAppAppAvailability_Global,
		"lineofbusiness": ManagedMobileLobAppAppAvailability_LineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedMobileLobAppAppAvailability(input)
	return &out, nil
}
