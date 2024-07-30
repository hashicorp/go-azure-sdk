package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionProtectedMessagingRedirectAppType string

const (
	TargetedManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp        TargetedManagedAppProtectionProtectedMessagingRedirectAppType = "anyApp"
	TargetedManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp TargetedManagedAppProtectionProtectedMessagingRedirectAppType = "anyManagedApp"
	TargetedManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps  TargetedManagedAppProtectionProtectedMessagingRedirectAppType = "specificApps"
)

func PossibleValuesForTargetedManagedAppProtectionProtectedMessagingRedirectAppType() []string {
	return []string{
		string(TargetedManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp),
		string(TargetedManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp),
		string(TargetedManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps),
	}
}

func (s *TargetedManagedAppProtectionProtectedMessagingRedirectAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionProtectedMessagingRedirectAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionProtectedMessagingRedirectAppType(input string) (*TargetedManagedAppProtectionProtectedMessagingRedirectAppType, error) {
	vals := map[string]TargetedManagedAppProtectionProtectedMessagingRedirectAppType{
		"anyapp":        TargetedManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp,
		"anymanagedapp": TargetedManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp,
		"specificapps":  TargetedManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionProtectedMessagingRedirectAppType(input)
	return &out, nil
}
