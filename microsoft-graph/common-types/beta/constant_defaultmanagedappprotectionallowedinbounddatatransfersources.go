package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedInboundDataTransferSources string

const (
	DefaultManagedAppProtectionAllowedInboundDataTransferSources_AllApps     DefaultManagedAppProtectionAllowedInboundDataTransferSources = "allApps"
	DefaultManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps DefaultManagedAppProtectionAllowedInboundDataTransferSources = "managedApps"
	DefaultManagedAppProtectionAllowedInboundDataTransferSources_None        DefaultManagedAppProtectionAllowedInboundDataTransferSources = "none"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedInboundDataTransferSources_AllApps),
		string(DefaultManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps),
		string(DefaultManagedAppProtectionAllowedInboundDataTransferSources_None),
	}
}

func (s *DefaultManagedAppProtectionAllowedInboundDataTransferSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAllowedInboundDataTransferSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAllowedInboundDataTransferSources(input string) (*DefaultManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     DefaultManagedAppProtectionAllowedInboundDataTransferSources_AllApps,
		"managedapps": DefaultManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps,
		"none":        DefaultManagedAppProtectionAllowedInboundDataTransferSources_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
