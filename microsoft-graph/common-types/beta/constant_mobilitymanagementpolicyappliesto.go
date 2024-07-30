package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobilityManagementPolicyAppliesTo string

const (
	MobilityManagementPolicyAppliesTo_All      MobilityManagementPolicyAppliesTo = "all"
	MobilityManagementPolicyAppliesTo_None     MobilityManagementPolicyAppliesTo = "none"
	MobilityManagementPolicyAppliesTo_Selected MobilityManagementPolicyAppliesTo = "selected"
)

func PossibleValuesForMobilityManagementPolicyAppliesTo() []string {
	return []string{
		string(MobilityManagementPolicyAppliesTo_All),
		string(MobilityManagementPolicyAppliesTo_None),
		string(MobilityManagementPolicyAppliesTo_Selected),
	}
}

func (s *MobilityManagementPolicyAppliesTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobilityManagementPolicyAppliesTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobilityManagementPolicyAppliesTo(input string) (*MobilityManagementPolicyAppliesTo, error) {
	vals := map[string]MobilityManagementPolicyAppliesTo{
		"all":      MobilityManagementPolicyAppliesTo_All,
		"none":     MobilityManagementPolicyAppliesTo_None,
		"selected": MobilityManagementPolicyAppliesTo_Selected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobilityManagementPolicyAppliesTo(input)
	return &out, nil
}
