package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionProtectedMessagingRedirectAppType string

const (
	DefaultManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp        DefaultManagedAppProtectionProtectedMessagingRedirectAppType = "anyApp"
	DefaultManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp DefaultManagedAppProtectionProtectedMessagingRedirectAppType = "anyManagedApp"
	DefaultManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps  DefaultManagedAppProtectionProtectedMessagingRedirectAppType = "specificApps"
)

func PossibleValuesForDefaultManagedAppProtectionProtectedMessagingRedirectAppType() []string {
	return []string{
		string(DefaultManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp),
		string(DefaultManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp),
		string(DefaultManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps),
	}
}

func (s *DefaultManagedAppProtectionProtectedMessagingRedirectAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionProtectedMessagingRedirectAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionProtectedMessagingRedirectAppType(input string) (*DefaultManagedAppProtectionProtectedMessagingRedirectAppType, error) {
	vals := map[string]DefaultManagedAppProtectionProtectedMessagingRedirectAppType{
		"anyapp":        DefaultManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp,
		"anymanagedapp": DefaultManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp,
		"specificapps":  DefaultManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionProtectedMessagingRedirectAppType(input)
	return &out, nil
}
