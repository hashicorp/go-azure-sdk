package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionAllowedInboundDataTransferSources string

const (
	WindowsManagedAppProtectionAllowedInboundDataTransferSources_AllApps WindowsManagedAppProtectionAllowedInboundDataTransferSources = "allApps"
	WindowsManagedAppProtectionAllowedInboundDataTransferSources_None    WindowsManagedAppProtectionAllowedInboundDataTransferSources = "none"
)

func PossibleValuesForWindowsManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(WindowsManagedAppProtectionAllowedInboundDataTransferSources_AllApps),
		string(WindowsManagedAppProtectionAllowedInboundDataTransferSources_None),
	}
}

func (s *WindowsManagedAppProtectionAllowedInboundDataTransferSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppProtectionAllowedInboundDataTransferSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppProtectionAllowedInboundDataTransferSources(input string) (*WindowsManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]WindowsManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps": WindowsManagedAppProtectionAllowedInboundDataTransferSources_AllApps,
		"none":    WindowsManagedAppProtectionAllowedInboundDataTransferSources_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
