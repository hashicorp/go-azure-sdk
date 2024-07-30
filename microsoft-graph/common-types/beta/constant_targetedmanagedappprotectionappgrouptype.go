package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAppGroupType string

const (
	TargetedManagedAppProtectionAppGroupType_AllApps              TargetedManagedAppProtectionAppGroupType = "allApps"
	TargetedManagedAppProtectionAppGroupType_AllCoreMicrosoftApps TargetedManagedAppProtectionAppGroupType = "allCoreMicrosoftApps"
	TargetedManagedAppProtectionAppGroupType_AllMicrosoftApps     TargetedManagedAppProtectionAppGroupType = "allMicrosoftApps"
	TargetedManagedAppProtectionAppGroupType_SelectedPublicApps   TargetedManagedAppProtectionAppGroupType = "selectedPublicApps"
)

func PossibleValuesForTargetedManagedAppProtectionAppGroupType() []string {
	return []string{
		string(TargetedManagedAppProtectionAppGroupType_AllApps),
		string(TargetedManagedAppProtectionAppGroupType_AllCoreMicrosoftApps),
		string(TargetedManagedAppProtectionAppGroupType_AllMicrosoftApps),
		string(TargetedManagedAppProtectionAppGroupType_SelectedPublicApps),
	}
}

func (s *TargetedManagedAppProtectionAppGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAppGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAppGroupType(input string) (*TargetedManagedAppProtectionAppGroupType, error) {
	vals := map[string]TargetedManagedAppProtectionAppGroupType{
		"allapps":              TargetedManagedAppProtectionAppGroupType_AllApps,
		"allcoremicrosoftapps": TargetedManagedAppProtectionAppGroupType_AllCoreMicrosoftApps,
		"allmicrosoftapps":     TargetedManagedAppProtectionAppGroupType_AllMicrosoftApps,
		"selectedpublicapps":   TargetedManagedAppProtectionAppGroupType_SelectedPublicApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAppGroupType(input)
	return &out, nil
}
