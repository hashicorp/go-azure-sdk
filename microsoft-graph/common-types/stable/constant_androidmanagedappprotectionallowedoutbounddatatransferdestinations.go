package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps     AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations = "allApps"
	AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations = "managedApps"
	AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_None        AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations = "none"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps),
		string(AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps),
		string(AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_None),
	}
}

func (s *AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAllowedOutboundDataTransferDestinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps,
		"managedapps": AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps,
		"none":        AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
