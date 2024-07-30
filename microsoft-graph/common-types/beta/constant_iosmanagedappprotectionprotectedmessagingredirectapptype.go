package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionProtectedMessagingRedirectAppType string

const (
	IosManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp        IosManagedAppProtectionProtectedMessagingRedirectAppType = "anyApp"
	IosManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp IosManagedAppProtectionProtectedMessagingRedirectAppType = "anyManagedApp"
	IosManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps  IosManagedAppProtectionProtectedMessagingRedirectAppType = "specificApps"
)

func PossibleValuesForIosManagedAppProtectionProtectedMessagingRedirectAppType() []string {
	return []string{
		string(IosManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp),
		string(IosManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp),
		string(IosManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps),
	}
}

func (s *IosManagedAppProtectionProtectedMessagingRedirectAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionProtectedMessagingRedirectAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionProtectedMessagingRedirectAppType(input string) (*IosManagedAppProtectionProtectedMessagingRedirectAppType, error) {
	vals := map[string]IosManagedAppProtectionProtectedMessagingRedirectAppType{
		"anyapp":        IosManagedAppProtectionProtectedMessagingRedirectAppType_AnyApp,
		"anymanagedapp": IosManagedAppProtectionProtectedMessagingRedirectAppType_AnyManagedApp,
		"specificapps":  IosManagedAppProtectionProtectedMessagingRedirectAppType_SpecificApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionProtectedMessagingRedirectAppType(input)
	return &out, nil
}
