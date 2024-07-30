package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps     DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations = "allApps"
	DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations = "managedApps"
	DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_None        DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations = "none"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps),
		string(DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps),
		string(DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_None),
	}
}

func (s *DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAllowedOutboundDataTransferDestinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps,
		"managedapps": DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps,
		"none":        DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
