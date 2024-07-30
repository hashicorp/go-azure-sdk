package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionNotificationRestriction string

const (
	TargetedManagedAppProtectionNotificationRestriction_Allow                   TargetedManagedAppProtectionNotificationRestriction = "allow"
	TargetedManagedAppProtectionNotificationRestriction_Block                   TargetedManagedAppProtectionNotificationRestriction = "block"
	TargetedManagedAppProtectionNotificationRestriction_BlockOrganizationalData TargetedManagedAppProtectionNotificationRestriction = "blockOrganizationalData"
)

func PossibleValuesForTargetedManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(TargetedManagedAppProtectionNotificationRestriction_Allow),
		string(TargetedManagedAppProtectionNotificationRestriction_Block),
		string(TargetedManagedAppProtectionNotificationRestriction_BlockOrganizationalData),
	}
}

func (s *TargetedManagedAppProtectionNotificationRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionNotificationRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionNotificationRestriction(input string) (*TargetedManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]TargetedManagedAppProtectionNotificationRestriction{
		"allow":                   TargetedManagedAppProtectionNotificationRestriction_Allow,
		"block":                   TargetedManagedAppProtectionNotificationRestriction_Block,
		"blockorganizationaldata": TargetedManagedAppProtectionNotificationRestriction_BlockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
