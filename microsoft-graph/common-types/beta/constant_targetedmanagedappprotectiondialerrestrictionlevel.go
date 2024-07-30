package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionDialerRestrictionLevel string

const (
	TargetedManagedAppProtectionDialerRestrictionLevel_AllApps     TargetedManagedAppProtectionDialerRestrictionLevel = "allApps"
	TargetedManagedAppProtectionDialerRestrictionLevel_Blocked     TargetedManagedAppProtectionDialerRestrictionLevel = "blocked"
	TargetedManagedAppProtectionDialerRestrictionLevel_CustomApp   TargetedManagedAppProtectionDialerRestrictionLevel = "customApp"
	TargetedManagedAppProtectionDialerRestrictionLevel_ManagedApps TargetedManagedAppProtectionDialerRestrictionLevel = "managedApps"
)

func PossibleValuesForTargetedManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(TargetedManagedAppProtectionDialerRestrictionLevel_AllApps),
		string(TargetedManagedAppProtectionDialerRestrictionLevel_Blocked),
		string(TargetedManagedAppProtectionDialerRestrictionLevel_CustomApp),
		string(TargetedManagedAppProtectionDialerRestrictionLevel_ManagedApps),
	}
}

func (s *TargetedManagedAppProtectionDialerRestrictionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionDialerRestrictionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionDialerRestrictionLevel(input string) (*TargetedManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]TargetedManagedAppProtectionDialerRestrictionLevel{
		"allapps":     TargetedManagedAppProtectionDialerRestrictionLevel_AllApps,
		"blocked":     TargetedManagedAppProtectionDialerRestrictionLevel_Blocked,
		"customapp":   TargetedManagedAppProtectionDialerRestrictionLevel_CustomApp,
		"managedapps": TargetedManagedAppProtectionDialerRestrictionLevel_ManagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}
