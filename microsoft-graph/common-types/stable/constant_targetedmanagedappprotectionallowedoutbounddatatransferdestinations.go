package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps     TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations = "allApps"
	TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations = "managedApps"
	TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_None        TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations = "none"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps),
		string(TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps),
		string(TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_None),
	}
}

func (s *TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAllowedOutboundDataTransferDestinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps,
		"managedapps": TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps,
		"none":        TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
