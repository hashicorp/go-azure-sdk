package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionDialerRestrictionLevel string

const (
	DefaultManagedAppProtectionDialerRestrictionLevel_AllApps     DefaultManagedAppProtectionDialerRestrictionLevel = "allApps"
	DefaultManagedAppProtectionDialerRestrictionLevel_Blocked     DefaultManagedAppProtectionDialerRestrictionLevel = "blocked"
	DefaultManagedAppProtectionDialerRestrictionLevel_CustomApp   DefaultManagedAppProtectionDialerRestrictionLevel = "customApp"
	DefaultManagedAppProtectionDialerRestrictionLevel_ManagedApps DefaultManagedAppProtectionDialerRestrictionLevel = "managedApps"
)

func PossibleValuesForDefaultManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(DefaultManagedAppProtectionDialerRestrictionLevel_AllApps),
		string(DefaultManagedAppProtectionDialerRestrictionLevel_Blocked),
		string(DefaultManagedAppProtectionDialerRestrictionLevel_CustomApp),
		string(DefaultManagedAppProtectionDialerRestrictionLevel_ManagedApps),
	}
}

func (s *DefaultManagedAppProtectionDialerRestrictionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionDialerRestrictionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionDialerRestrictionLevel(input string) (*DefaultManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]DefaultManagedAppProtectionDialerRestrictionLevel{
		"allapps":     DefaultManagedAppProtectionDialerRestrictionLevel_AllApps,
		"blocked":     DefaultManagedAppProtectionDialerRestrictionLevel_Blocked,
		"customapp":   DefaultManagedAppProtectionDialerRestrictionLevel_CustomApp,
		"managedapps": DefaultManagedAppProtectionDialerRestrictionLevel_ManagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}
