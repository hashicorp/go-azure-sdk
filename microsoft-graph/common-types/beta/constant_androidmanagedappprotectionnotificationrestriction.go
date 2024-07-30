package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionNotificationRestriction string

const (
	AndroidManagedAppProtectionNotificationRestriction_Allow                   AndroidManagedAppProtectionNotificationRestriction = "allow"
	AndroidManagedAppProtectionNotificationRestriction_Block                   AndroidManagedAppProtectionNotificationRestriction = "block"
	AndroidManagedAppProtectionNotificationRestriction_BlockOrganizationalData AndroidManagedAppProtectionNotificationRestriction = "blockOrganizationalData"
)

func PossibleValuesForAndroidManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(AndroidManagedAppProtectionNotificationRestriction_Allow),
		string(AndroidManagedAppProtectionNotificationRestriction_Block),
		string(AndroidManagedAppProtectionNotificationRestriction_BlockOrganizationalData),
	}
}

func (s *AndroidManagedAppProtectionNotificationRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionNotificationRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionNotificationRestriction(input string) (*AndroidManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]AndroidManagedAppProtectionNotificationRestriction{
		"allow":                   AndroidManagedAppProtectionNotificationRestriction_Allow,
		"block":                   AndroidManagedAppProtectionNotificationRestriction_Block,
		"blockorganizationaldata": AndroidManagedAppProtectionNotificationRestriction_BlockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
