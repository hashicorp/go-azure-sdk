package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidLobAppAppAvailability string

const (
	ManagedAndroidLobAppAppAvailability_Global         ManagedAndroidLobAppAppAvailability = "global"
	ManagedAndroidLobAppAppAvailability_LineOfBusiness ManagedAndroidLobAppAppAvailability = "lineOfBusiness"
)

func PossibleValuesForManagedAndroidLobAppAppAvailability() []string {
	return []string{
		string(ManagedAndroidLobAppAppAvailability_Global),
		string(ManagedAndroidLobAppAppAvailability_LineOfBusiness),
	}
}

func (s *ManagedAndroidLobAppAppAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAndroidLobAppAppAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAndroidLobAppAppAvailability(input string) (*ManagedAndroidLobAppAppAvailability, error) {
	vals := map[string]ManagedAndroidLobAppAppAvailability{
		"global":         ManagedAndroidLobAppAppAvailability_Global,
		"lineofbusiness": ManagedAndroidLobAppAppAvailability_LineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidLobAppAppAvailability(input)
	return &out, nil
}
