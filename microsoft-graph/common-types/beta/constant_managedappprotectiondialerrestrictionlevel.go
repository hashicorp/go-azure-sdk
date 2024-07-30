package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionDialerRestrictionLevel string

const (
	ManagedAppProtectionDialerRestrictionLevel_AllApps     ManagedAppProtectionDialerRestrictionLevel = "allApps"
	ManagedAppProtectionDialerRestrictionLevel_Blocked     ManagedAppProtectionDialerRestrictionLevel = "blocked"
	ManagedAppProtectionDialerRestrictionLevel_CustomApp   ManagedAppProtectionDialerRestrictionLevel = "customApp"
	ManagedAppProtectionDialerRestrictionLevel_ManagedApps ManagedAppProtectionDialerRestrictionLevel = "managedApps"
)

func PossibleValuesForManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(ManagedAppProtectionDialerRestrictionLevel_AllApps),
		string(ManagedAppProtectionDialerRestrictionLevel_Blocked),
		string(ManagedAppProtectionDialerRestrictionLevel_CustomApp),
		string(ManagedAppProtectionDialerRestrictionLevel_ManagedApps),
	}
}

func (s *ManagedAppProtectionDialerRestrictionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionDialerRestrictionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionDialerRestrictionLevel(input string) (*ManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]ManagedAppProtectionDialerRestrictionLevel{
		"allapps":     ManagedAppProtectionDialerRestrictionLevel_AllApps,
		"blocked":     ManagedAppProtectionDialerRestrictionLevel_Blocked,
		"customapp":   ManagedAppProtectionDialerRestrictionLevel_CustomApp,
		"managedapps": ManagedAppProtectionDialerRestrictionLevel_ManagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}
