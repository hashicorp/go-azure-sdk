package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	ManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps     ManagedAppProtectionAllowedOutboundDataTransferDestinations = "allApps"
	ManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps ManagedAppProtectionAllowedOutboundDataTransferDestinations = "managedApps"
	ManagedAppProtectionAllowedOutboundDataTransferDestinations_None        ManagedAppProtectionAllowedOutboundDataTransferDestinations = "none"
)

func PossibleValuesForManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(ManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps),
		string(ManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps),
		string(ManagedAppProtectionAllowedOutboundDataTransferDestinations_None),
	}
}

func (s *ManagedAppProtectionAllowedOutboundDataTransferDestinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAllowedOutboundDataTransferDestinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*ManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]ManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     ManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps,
		"managedapps": ManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps,
		"none":        ManagedAppProtectionAllowedOutboundDataTransferDestinations_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
