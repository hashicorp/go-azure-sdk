package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfigurationAppGroupType string

const (
	TargetedManagedAppConfigurationAppGroupType_AllApps              TargetedManagedAppConfigurationAppGroupType = "allApps"
	TargetedManagedAppConfigurationAppGroupType_AllCoreMicrosoftApps TargetedManagedAppConfigurationAppGroupType = "allCoreMicrosoftApps"
	TargetedManagedAppConfigurationAppGroupType_AllMicrosoftApps     TargetedManagedAppConfigurationAppGroupType = "allMicrosoftApps"
	TargetedManagedAppConfigurationAppGroupType_SelectedPublicApps   TargetedManagedAppConfigurationAppGroupType = "selectedPublicApps"
)

func PossibleValuesForTargetedManagedAppConfigurationAppGroupType() []string {
	return []string{
		string(TargetedManagedAppConfigurationAppGroupType_AllApps),
		string(TargetedManagedAppConfigurationAppGroupType_AllCoreMicrosoftApps),
		string(TargetedManagedAppConfigurationAppGroupType_AllMicrosoftApps),
		string(TargetedManagedAppConfigurationAppGroupType_SelectedPublicApps),
	}
}

func (s *TargetedManagedAppConfigurationAppGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppConfigurationAppGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppConfigurationAppGroupType(input string) (*TargetedManagedAppConfigurationAppGroupType, error) {
	vals := map[string]TargetedManagedAppConfigurationAppGroupType{
		"allapps":              TargetedManagedAppConfigurationAppGroupType_AllApps,
		"allcoremicrosoftapps": TargetedManagedAppConfigurationAppGroupType_AllCoreMicrosoftApps,
		"allmicrosoftapps":     TargetedManagedAppConfigurationAppGroupType_AllMicrosoftApps,
		"selectedpublicapps":   TargetedManagedAppConfigurationAppGroupType_SelectedPublicApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppConfigurationAppGroupType(input)
	return &out, nil
}
