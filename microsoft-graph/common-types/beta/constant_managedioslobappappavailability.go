package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSLobAppAppAvailability string

const (
	ManagedIOSLobAppAppAvailability_Global         ManagedIOSLobAppAppAvailability = "global"
	ManagedIOSLobAppAppAvailability_LineOfBusiness ManagedIOSLobAppAppAvailability = "lineOfBusiness"
)

func PossibleValuesForManagedIOSLobAppAppAvailability() []string {
	return []string{
		string(ManagedIOSLobAppAppAvailability_Global),
		string(ManagedIOSLobAppAppAvailability_LineOfBusiness),
	}
}

func (s *ManagedIOSLobAppAppAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedIOSLobAppAppAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedIOSLobAppAppAvailability(input string) (*ManagedIOSLobAppAppAvailability, error) {
	vals := map[string]ManagedIOSLobAppAppAvailability{
		"global":         ManagedIOSLobAppAppAvailability_Global,
		"lineofbusiness": ManagedIOSLobAppAppAvailability_LineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIOSLobAppAppAvailability(input)
	return &out, nil
}
