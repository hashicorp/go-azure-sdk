package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedOutboundDataTransferDestinations string

const (
	IosManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps     IosManagedAppProtectionAllowedOutboundDataTransferDestinations = "allApps"
	IosManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps IosManagedAppProtectionAllowedOutboundDataTransferDestinations = "managedApps"
	IosManagedAppProtectionAllowedOutboundDataTransferDestinations_None        IosManagedAppProtectionAllowedOutboundDataTransferDestinations = "none"
)

func PossibleValuesForIosManagedAppProtectionAllowedOutboundDataTransferDestinations() []string {
	return []string{
		string(IosManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps),
		string(IosManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps),
		string(IosManagedAppProtectionAllowedOutboundDataTransferDestinations_None),
	}
}

func (s *IosManagedAppProtectionAllowedOutboundDataTransferDestinations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAllowedOutboundDataTransferDestinations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAllowedOutboundDataTransferDestinations(input string) (*IosManagedAppProtectionAllowedOutboundDataTransferDestinations, error) {
	vals := map[string]IosManagedAppProtectionAllowedOutboundDataTransferDestinations{
		"allapps":     IosManagedAppProtectionAllowedOutboundDataTransferDestinations_AllApps,
		"managedapps": IosManagedAppProtectionAllowedOutboundDataTransferDestinations_ManagedApps,
		"none":        IosManagedAppProtectionAllowedOutboundDataTransferDestinations_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedOutboundDataTransferDestinations(input)
	return &out, nil
}
