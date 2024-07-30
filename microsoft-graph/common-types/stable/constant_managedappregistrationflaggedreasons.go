package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppRegistrationFlaggedReasons string

const (
	ManagedAppRegistrationFlaggedReasons_None         ManagedAppRegistrationFlaggedReasons = "none"
	ManagedAppRegistrationFlaggedReasons_RootedDevice ManagedAppRegistrationFlaggedReasons = "rootedDevice"
)

func PossibleValuesForManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(ManagedAppRegistrationFlaggedReasons_None),
		string(ManagedAppRegistrationFlaggedReasons_RootedDevice),
	}
}

func (s *ManagedAppRegistrationFlaggedReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppRegistrationFlaggedReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppRegistrationFlaggedReasons(input string) (*ManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]ManagedAppRegistrationFlaggedReasons{
		"none":         ManagedAppRegistrationFlaggedReasons_None,
		"rooteddevice": ManagedAppRegistrationFlaggedReasons_RootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
