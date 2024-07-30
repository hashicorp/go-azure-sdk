package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedInboundDataTransferSources string

const (
	IosManagedAppProtectionAllowedInboundDataTransferSources_AllApps     IosManagedAppProtectionAllowedInboundDataTransferSources = "allApps"
	IosManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps IosManagedAppProtectionAllowedInboundDataTransferSources = "managedApps"
	IosManagedAppProtectionAllowedInboundDataTransferSources_None        IosManagedAppProtectionAllowedInboundDataTransferSources = "none"
)

func PossibleValuesForIosManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(IosManagedAppProtectionAllowedInboundDataTransferSources_AllApps),
		string(IosManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps),
		string(IosManagedAppProtectionAllowedInboundDataTransferSources_None),
	}
}

func (s *IosManagedAppProtectionAllowedInboundDataTransferSources) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAllowedInboundDataTransferSources(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAllowedInboundDataTransferSources(input string) (*IosManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]IosManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     IosManagedAppProtectionAllowedInboundDataTransferSources_AllApps,
		"managedapps": IosManagedAppProtectionAllowedInboundDataTransferSources_ManagedApps,
		"none":        IosManagedAppProtectionAllowedInboundDataTransferSources_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
