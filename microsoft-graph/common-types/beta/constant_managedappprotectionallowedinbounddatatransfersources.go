package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedInboundDataTransferSources string

const (
	ManagedAppProtectionAllowedInboundDataTransferSources_AllApps     ManagedAppProtectionAllowedInboundDataTransferSources = "allApps"
	ManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps ManagedAppProtectionAllowedInboundDataTransferSources = "managedApps"
	ManagedAppProtectionAllowedInboundDataTransferSources_None        ManagedAppProtectionAllowedInboundDataTransferSources = "none"
)

func PossibleValuesForManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(ManagedAppProtectionAllowedInboundDataTransferSources_AllApps),
		string(ManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps),
		string(ManagedAppProtectionAllowedInboundDataTransferSources_None),
	}
}

func (s *ManagedAppProtectionAllowedInboundDataTransferSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAllowedInboundDataTransferSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAllowedInboundDataTransferSources(input string) (*ManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]ManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     ManagedAppProtectionAllowedInboundDataTransferSources_AllApps,
		"managedapps": ManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps,
		"none":        ManagedAppProtectionAllowedInboundDataTransferSources_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
