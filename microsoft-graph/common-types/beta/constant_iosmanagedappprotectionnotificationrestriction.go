package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionNotificationRestriction string

const (
	IosManagedAppProtectionNotificationRestriction_Allow                   IosManagedAppProtectionNotificationRestriction = "allow"
	IosManagedAppProtectionNotificationRestriction_Block                   IosManagedAppProtectionNotificationRestriction = "block"
	IosManagedAppProtectionNotificationRestriction_BlockOrganizationalData IosManagedAppProtectionNotificationRestriction = "blockOrganizationalData"
)

func PossibleValuesForIosManagedAppProtectionNotificationRestriction() []string {
	return []string{
		string(IosManagedAppProtectionNotificationRestriction_Allow),
		string(IosManagedAppProtectionNotificationRestriction_Block),
		string(IosManagedAppProtectionNotificationRestriction_BlockOrganizationalData),
	}
}

func (s *IosManagedAppProtectionNotificationRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionNotificationRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionNotificationRestriction(input string) (*IosManagedAppProtectionNotificationRestriction, error) {
	vals := map[string]IosManagedAppProtectionNotificationRestriction{
		"allow":                   IosManagedAppProtectionNotificationRestriction_Allow,
		"block":                   IosManagedAppProtectionNotificationRestriction_Block,
		"blockorganizationaldata": IosManagedAppProtectionNotificationRestriction_BlockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionNotificationRestriction(input)
	return &out, nil
}
