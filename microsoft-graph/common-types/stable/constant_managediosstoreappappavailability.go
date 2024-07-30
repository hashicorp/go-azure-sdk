package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIOSStoreAppAppAvailability string

const (
	ManagedIOSStoreAppAppAvailability_Global         ManagedIOSStoreAppAppAvailability = "global"
	ManagedIOSStoreAppAppAvailability_LineOfBusiness ManagedIOSStoreAppAppAvailability = "lineOfBusiness"
)

func PossibleValuesForManagedIOSStoreAppAppAvailability() []string {
	return []string{
		string(ManagedIOSStoreAppAppAvailability_Global),
		string(ManagedIOSStoreAppAppAvailability_LineOfBusiness),
	}
}

func (s *ManagedIOSStoreAppAppAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedIOSStoreAppAppAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedIOSStoreAppAppAvailability(input string) (*ManagedIOSStoreAppAppAvailability, error) {
	vals := map[string]ManagedIOSStoreAppAppAvailability{
		"global":         ManagedIOSStoreAppAppAvailability_Global,
		"lineofbusiness": ManagedIOSStoreAppAppAvailability_LineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedIOSStoreAppAppAvailability(input)
	return &out, nil
}
