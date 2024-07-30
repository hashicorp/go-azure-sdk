package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedInboundDataTransferSources string

const (
	AndroidManagedAppProtectionAllowedInboundDataTransferSources_AllApps     AndroidManagedAppProtectionAllowedInboundDataTransferSources = "allApps"
	AndroidManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps AndroidManagedAppProtectionAllowedInboundDataTransferSources = "managedApps"
	AndroidManagedAppProtectionAllowedInboundDataTransferSources_None        AndroidManagedAppProtectionAllowedInboundDataTransferSources = "none"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedInboundDataTransferSources_AllApps),
		string(AndroidManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps),
		string(AndroidManagedAppProtectionAllowedInboundDataTransferSources_None),
	}
}

func (s *AndroidManagedAppProtectionAllowedInboundDataTransferSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAllowedInboundDataTransferSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAllowedInboundDataTransferSources(input string) (*AndroidManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     AndroidManagedAppProtectionAllowedInboundDataTransferSources_AllApps,
		"managedapps": AndroidManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps,
		"none":        AndroidManagedAppProtectionAllowedInboundDataTransferSources_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
