package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionNotificationRestriction string

const (
	ManagedAppProtectionNotificationRestriction_Allow                   ManagedAppProtectionNotificationRestriction = "allow"
	ManagedAppProtectionNotificationRestriction_Block                   ManagedAppProtectionNotificationRestriction = "block"
	ManagedAppProtectionNotificationRestriction_BlockOrganizationalData ManagedAppProtectionNotificationRestriction = "blockOrganizationalData"
)

func PossibleValuesForManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(ManagedAppProtectionNotificationRestriction_Allow),
		string(ManagedAppProtectionNotificationRestriction_Block),
		string(ManagedAppProtectionNotificationRestriction_BlockOrganizationalData),
	}
}

func (s *ManagedAppProtectionNotificationRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionNotificationRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionNotificationRestriction(input string) (*ManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]ManagedAppProtectionNotificationRestriction{
		"allow":                   ManagedAppProtectionNotificationRestriction_Allow,
		"block":                   ManagedAppProtectionNotificationRestriction_Block,
		"blockorganizationaldata": ManagedAppProtectionNotificationRestriction_BlockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
