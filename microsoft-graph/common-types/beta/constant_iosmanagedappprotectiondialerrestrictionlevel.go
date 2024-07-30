package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionDialerRestrictionLevel string

const (
	IosManagedAppProtectionDialerRestrictionLevel_AllApps     IosManagedAppProtectionDialerRestrictionLevel = "allApps"
	IosManagedAppProtectionDialerRestrictionLevel_Blocked     IosManagedAppProtectionDialerRestrictionLevel = "blocked"
	IosManagedAppProtectionDialerRestrictionLevel_CustomApp   IosManagedAppProtectionDialerRestrictionLevel = "customApp"
	IosManagedAppProtectionDialerRestrictionLevel_ManagedApps IosManagedAppProtectionDialerRestrictionLevel = "managedApps"
)

func PossibleValuesForIosManagedAppProtectionDialerRestrictionLevel() []string {
	return []string{
		string(IosManagedAppProtectionDialerRestrictionLevel_AllApps),
		string(IosManagedAppProtectionDialerRestrictionLevel_Blocked),
		string(IosManagedAppProtectionDialerRestrictionLevel_CustomApp),
		string(IosManagedAppProtectionDialerRestrictionLevel_ManagedApps),
	}
}

func (s *IosManagedAppProtectionDialerRestrictionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionDialerRestrictionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionDialerRestrictionLevel(input string) (*IosManagedAppProtectionDialerRestrictionLevel, error) {
	vals := map[string]IosManagedAppProtectionDialerRestrictionLevel{
		"allapps":     IosManagedAppProtectionDialerRestrictionLevel_AllApps,
		"blocked":     IosManagedAppProtectionDialerRestrictionLevel_Blocked,
		"customapp":   IosManagedAppProtectionDialerRestrictionLevel_CustomApp,
		"managedapps": IosManagedAppProtectionDialerRestrictionLevel_ManagedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionDialerRestrictionLevel(input)
	return &out, nil
}
