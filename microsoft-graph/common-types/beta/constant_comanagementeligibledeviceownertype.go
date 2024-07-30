package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceOwnerType string

const (
	ComanagementEligibleDeviceOwnerType_Company  ComanagementEligibleDeviceOwnerType = "company"
	ComanagementEligibleDeviceOwnerType_Personal ComanagementEligibleDeviceOwnerType = "personal"
	ComanagementEligibleDeviceOwnerType_Unknown  ComanagementEligibleDeviceOwnerType = "unknown"
)

func PossibleValuesForComanagementEligibleDeviceOwnerType() []string {
	return []string{
		string(ComanagementEligibleDeviceOwnerType_Company),
		string(ComanagementEligibleDeviceOwnerType_Personal),
		string(ComanagementEligibleDeviceOwnerType_Unknown),
	}
}

func (s *ComanagementEligibleDeviceOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComanagementEligibleDeviceOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComanagementEligibleDeviceOwnerType(input string) (*ComanagementEligibleDeviceOwnerType, error) {
	vals := map[string]ComanagementEligibleDeviceOwnerType{
		"company":  ComanagementEligibleDeviceOwnerType_Company,
		"personal": ComanagementEligibleDeviceOwnerType_Personal,
		"unknown":  ComanagementEligibleDeviceOwnerType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceOwnerType(input)
	return &out, nil
}
