package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionDialerRestrictionLevel string

const (
	AndroidManagedAppProtectionDialerRestrictionLevel_AllApps     AndroidManagedAppProtectionDialerRestrictionLevel = "allApps"
	AndroidManagedAppProtectionDialerRestrictionLevel_Blocked     AndroidManagedAppProtectionDialerRestrictionLevel = "blocked"
	AndroidManagedAppProtectionDialerRestrictionLevel_CustomApp   AndroidManagedAppProtectionDialerRestrictionLevel = "customApp"
	AndroidManagedAppProtectionDialerRestrictionLevel_ManagedApps AndroidManagedAppProtectionDialerRestrictionLevel = "managedApps"
)

func PossibleValuesForAndroidManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(AndroidManagedAppProtectionDialerRestrictionLevel_AllApps),
		string(AndroidManagedAppProtectionDialerRestrictionLevel_Blocked),
		string(AndroidManagedAppProtectionDialerRestrictionLevel_CustomApp),
		string(AndroidManagedAppProtectionDialerRestrictionLevel_ManagedApps),
	}
}

func (s *AndroidManagedAppProtectionDialerRestrictionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionDialerRestrictionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionDialerRestrictionLevel(input string) (*AndroidManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]AndroidManagedAppProtectionDialerRestrictionLevel{
		"allapps":     AndroidManagedAppProtectionDialerRestrictionLevel_AllApps,
		"blocked":     AndroidManagedAppProtectionDialerRestrictionLevel_Blocked,
		"customapp":   AndroidManagedAppProtectionDialerRestrictionLevel_CustomApp,
		"managedapps": AndroidManagedAppProtectionDialerRestrictionLevel_ManagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}
