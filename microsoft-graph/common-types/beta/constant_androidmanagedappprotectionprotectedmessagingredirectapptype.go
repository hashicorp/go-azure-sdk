package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionProtectedMessagingRedirectAppType string

const (
	AndroidManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp        AndroidManagedAppProtectionProtectedMessagingRedirectAppType = "anyApp"
	AndroidManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp AndroidManagedAppProtectionProtectedMessagingRedirectAppType = "anyManagedApp"
	AndroidManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps  AndroidManagedAppProtectionProtectedMessagingRedirectAppType = "specificApps"
)

func PossibleValuesForAndroidManagedAppProtectionProtectedMessagingRedirectAppType() []string {
	return []string{
		string(AndroidManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp),
		string(AndroidManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp),
		string(AndroidManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps),
	}
}

func (s *AndroidManagedAppProtectionProtectedMessagingRedirectAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionProtectedMessagingRedirectAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionProtectedMessagingRedirectAppType(input string) (*AndroidManagedAppProtectionProtectedMessagingRedirectAppType, error) {
	vals := map[string]AndroidManagedAppProtectionProtectedMessagingRedirectAppType{
		"anyapp":        AndroidManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp,
		"anymanagedapp": AndroidManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp,
		"specificapps":  AndroidManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionProtectedMessagingRedirectAppType(input)
	return &out, nil
}
