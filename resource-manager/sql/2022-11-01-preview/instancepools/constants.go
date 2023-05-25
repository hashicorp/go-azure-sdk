package instancepools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstancePoolLicenseType string

const (
	InstancePoolLicenseTypeBasePrice       InstancePoolLicenseType = "BasePrice"
	InstancePoolLicenseTypeLicenseIncluded InstancePoolLicenseType = "LicenseIncluded"
)

func PossibleValuesForInstancePoolLicenseType() []string {
	return []string{
		string(InstancePoolLicenseTypeBasePrice),
		string(InstancePoolLicenseTypeLicenseIncluded),
	}
}

func (s *InstancePoolLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstancePoolLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstancePoolLicenseType(input string) (*InstancePoolLicenseType, error) {
	vals := map[string]InstancePoolLicenseType{
		"baseprice":       InstancePoolLicenseTypeBasePrice,
		"licenseincluded": InstancePoolLicenseTypeLicenseIncluded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstancePoolLicenseType(input)
	return &out, nil
}
