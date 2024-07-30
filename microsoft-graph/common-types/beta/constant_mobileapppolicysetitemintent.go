package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppPolicySetItemIntent string

const (
	MobileAppPolicySetItemIntent_Available                  MobileAppPolicySetItemIntent = "available"
	MobileAppPolicySetItemIntent_AvailableWithoutEnrollment MobileAppPolicySetItemIntent = "availableWithoutEnrollment"
	MobileAppPolicySetItemIntent_Required                   MobileAppPolicySetItemIntent = "required"
	MobileAppPolicySetItemIntent_Uninstall                  MobileAppPolicySetItemIntent = "uninstall"
)

func PossibleValuesForMobileAppPolicySetItemIntent() []string {
	return []string{
		string(MobileAppPolicySetItemIntent_Available),
		string(MobileAppPolicySetItemIntent_AvailableWithoutEnrollment),
		string(MobileAppPolicySetItemIntent_Required),
		string(MobileAppPolicySetItemIntent_Uninstall),
	}
}

func (s *MobileAppPolicySetItemIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppPolicySetItemIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppPolicySetItemIntent(input string) (*MobileAppPolicySetItemIntent, error) {
	vals := map[string]MobileAppPolicySetItemIntent{
		"available":                  MobileAppPolicySetItemIntent_Available,
		"availablewithoutenrollment": MobileAppPolicySetItemIntent_AvailableWithoutEnrollment,
		"required":                   MobileAppPolicySetItemIntent_Required,
		"uninstall":                  MobileAppPolicySetItemIntent_Uninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppPolicySetItemIntent(input)
	return &out, nil
}
