package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedInboundDataTransferSources string

const (
	TargetedManagedAppProtectionAllowedInboundDataTransferSources_AllApps     TargetedManagedAppProtectionAllowedInboundDataTransferSources = "allApps"
	TargetedManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps TargetedManagedAppProtectionAllowedInboundDataTransferSources = "managedApps"
	TargetedManagedAppProtectionAllowedInboundDataTransferSources_None        TargetedManagedAppProtectionAllowedInboundDataTransferSources = "none"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedInboundDataTransferSources_AllApps),
		string(TargetedManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps),
		string(TargetedManagedAppProtectionAllowedInboundDataTransferSources_None),
	}
}

func (s *TargetedManagedAppProtectionAllowedInboundDataTransferSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAllowedInboundDataTransferSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAllowedInboundDataTransferSources(input string) (*TargetedManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     TargetedManagedAppProtectionAllowedInboundDataTransferSources_AllApps,
		"managedapps": TargetedManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps,
		"none":        TargetedManagedAppProtectionAllowedInboundDataTransferSources_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
