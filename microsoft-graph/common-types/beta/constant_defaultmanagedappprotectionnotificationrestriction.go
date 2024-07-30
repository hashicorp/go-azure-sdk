package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionNotificationRestriction string

const (
	DefaultManagedAppProtectionNotificationRestriction_Allow                   DefaultManagedAppProtectionNotificationRestriction = "allow"
	DefaultManagedAppProtectionNotificationRestriction_Block                   DefaultManagedAppProtectionNotificationRestriction = "block"
	DefaultManagedAppProtectionNotificationRestriction_BlockOrganizationalData DefaultManagedAppProtectionNotificationRestriction = "blockOrganizationalData"
)

func PossibleValuesForDefaultManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(DefaultManagedAppProtectionNotificationRestriction_Allow),
		string(DefaultManagedAppProtectionNotificationRestriction_Block),
		string(DefaultManagedAppProtectionNotificationRestriction_BlockOrganizationalData),
	}
}

func (s *DefaultManagedAppProtectionNotificationRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionNotificationRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionNotificationRestriction(input string) (*DefaultManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]DefaultManagedAppProtectionNotificationRestriction{
		"allow":                   DefaultManagedAppProtectionNotificationRestriction_Allow,
		"block":                   DefaultManagedAppProtectionNotificationRestriction_Block,
		"blockorganizationaldata": DefaultManagedAppProtectionNotificationRestriction_BlockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
