package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppGroupType string

const (
	AndroidManagedAppProtectionAppGroupType_AllApps              AndroidManagedAppProtectionAppGroupType = "allApps"
	AndroidManagedAppProtectionAppGroupType_AllCoreMicrosoftApps AndroidManagedAppProtectionAppGroupType = "allCoreMicrosoftApps"
	AndroidManagedAppProtectionAppGroupType_AllMicrosoftApps     AndroidManagedAppProtectionAppGroupType = "allMicrosoftApps"
	AndroidManagedAppProtectionAppGroupType_SelectedPublicApps   AndroidManagedAppProtectionAppGroupType = "selectedPublicApps"
)

func PossibleValuesForAndroidManagedAppProtectionAppGroupType() []string {
	return []string{
		string(AndroidManagedAppProtectionAppGroupType_AllApps),
		string(AndroidManagedAppProtectionAppGroupType_AllCoreMicrosoftApps),
		string(AndroidManagedAppProtectionAppGroupType_AllMicrosoftApps),
		string(AndroidManagedAppProtectionAppGroupType_SelectedPublicApps),
	}
}

func (s *AndroidManagedAppProtectionAppGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppGroupType(input string) (*AndroidManagedAppProtectionAppGroupType, error) {
	vals := map[string]AndroidManagedAppProtectionAppGroupType{
		"allapps":              AndroidManagedAppProtectionAppGroupType_AllApps,
		"allcoremicrosoftapps": AndroidManagedAppProtectionAppGroupType_AllCoreMicrosoftApps,
		"allmicrosoftapps":     AndroidManagedAppProtectionAppGroupType_AllMicrosoftApps,
		"selectedpublicapps":   AndroidManagedAppProtectionAppGroupType_SelectedPublicApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppGroupType(input)
	return &out, nil
}
