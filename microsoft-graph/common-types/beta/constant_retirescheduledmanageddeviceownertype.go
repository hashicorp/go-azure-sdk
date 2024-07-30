package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDeviceOwnerType string

const (
	RetireScheduledManagedDeviceOwnerType_Company  RetireScheduledManagedDeviceOwnerType = "company"
	RetireScheduledManagedDeviceOwnerType_Personal RetireScheduledManagedDeviceOwnerType = "personal"
	RetireScheduledManagedDeviceOwnerType_Unknown  RetireScheduledManagedDeviceOwnerType = "unknown"
)

func PossibleValuesForRetireScheduledManagedDeviceOwnerType() []string {
	return []string{
		string(RetireScheduledManagedDeviceOwnerType_Company),
		string(RetireScheduledManagedDeviceOwnerType_Personal),
		string(RetireScheduledManagedDeviceOwnerType_Unknown),
	}
}

func (s *RetireScheduledManagedDeviceOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetireScheduledManagedDeviceOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetireScheduledManagedDeviceOwnerType(input string) (*RetireScheduledManagedDeviceOwnerType, error) {
	vals := map[string]RetireScheduledManagedDeviceOwnerType{
		"company":  RetireScheduledManagedDeviceOwnerType_Company,
		"personal": RetireScheduledManagedDeviceOwnerType_Personal,
		"unknown":  RetireScheduledManagedDeviceOwnerType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetireScheduledManagedDeviceOwnerType(input)
	return &out, nil
}
