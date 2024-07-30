package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSystemExtensionTypeMappingAllowedTypes string

const (
	MacOSSystemExtensionTypeMappingAllowedTypes_DriverExtensionsAllowed           MacOSSystemExtensionTypeMappingAllowedTypes = "driverExtensionsAllowed"
	MacOSSystemExtensionTypeMappingAllowedTypes_EndpointSecurityExtensionsAllowed MacOSSystemExtensionTypeMappingAllowedTypes = "endpointSecurityExtensionsAllowed"
	MacOSSystemExtensionTypeMappingAllowedTypes_NetworkExtensionsAllowed          MacOSSystemExtensionTypeMappingAllowedTypes = "networkExtensionsAllowed"
)

func PossibleValuesForMacOSSystemExtensionTypeMappingAllowedTypes() []string {
	return []string{
		string(MacOSSystemExtensionTypeMappingAllowedTypes_DriverExtensionsAllowed),
		string(MacOSSystemExtensionTypeMappingAllowedTypes_EndpointSecurityExtensionsAllowed),
		string(MacOSSystemExtensionTypeMappingAllowedTypes_NetworkExtensionsAllowed),
	}
}

func (s *MacOSSystemExtensionTypeMappingAllowedTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSystemExtensionTypeMappingAllowedTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSystemExtensionTypeMappingAllowedTypes(input string) (*MacOSSystemExtensionTypeMappingAllowedTypes, error) {
	vals := map[string]MacOSSystemExtensionTypeMappingAllowedTypes{
		"driverextensionsallowed":           MacOSSystemExtensionTypeMappingAllowedTypes_DriverExtensionsAllowed,
		"endpointsecurityextensionsallowed": MacOSSystemExtensionTypeMappingAllowedTypes_EndpointSecurityExtensionsAllowed,
		"networkextensionsallowed":          MacOSSystemExtensionTypeMappingAllowedTypes_NetworkExtensionsAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSystemExtensionTypeMappingAllowedTypes(input)
	return &out, nil
}
