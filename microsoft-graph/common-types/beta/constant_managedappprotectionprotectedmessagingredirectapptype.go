package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionProtectedMessagingRedirectAppType string

const (
	ManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp        ManagedAppProtectionProtectedMessagingRedirectAppType = "anyApp"
	ManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp ManagedAppProtectionProtectedMessagingRedirectAppType = "anyManagedApp"
	ManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps  ManagedAppProtectionProtectedMessagingRedirectAppType = "specificApps"
)

func PossibleValuesForManagedAppProtectionProtectedMessagingRedirectAppType() []string {
	return []string{
		string(ManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp),
		string(ManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp),
		string(ManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps),
	}
}

func (s *ManagedAppProtectionProtectedMessagingRedirectAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionProtectedMessagingRedirectAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionProtectedMessagingRedirectAppType(input string) (*ManagedAppProtectionProtectedMessagingRedirectAppType, error) {
	vals := map[string]ManagedAppProtectionProtectedMessagingRedirectAppType{
		"anyapp":        ManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp,
		"anymanagedapp": ManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp,
		"specificapps":  ManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionProtectedMessagingRedirectAppType(input)
	return &out, nil
}
