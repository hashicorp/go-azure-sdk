package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTargetTargetType string

const (
	CrossTenantAccessPolicyTargetTargetType_Application CrossTenantAccessPolicyTargetTargetType = "application"
	CrossTenantAccessPolicyTargetTargetType_Group       CrossTenantAccessPolicyTargetTargetType = "group"
	CrossTenantAccessPolicyTargetTargetType_User        CrossTenantAccessPolicyTargetTargetType = "user"
)

func PossibleValuesForCrossTenantAccessPolicyTargetTargetType() []string {
	return []string{
		string(CrossTenantAccessPolicyTargetTargetType_Application),
		string(CrossTenantAccessPolicyTargetTargetType_Group),
		string(CrossTenantAccessPolicyTargetTargetType_User),
	}
}

func (s *CrossTenantAccessPolicyTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCrossTenantAccessPolicyTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCrossTenantAccessPolicyTargetTargetType(input string) (*CrossTenantAccessPolicyTargetTargetType, error) {
	vals := map[string]CrossTenantAccessPolicyTargetTargetType{
		"application": CrossTenantAccessPolicyTargetTargetType_Application,
		"group":       CrossTenantAccessPolicyTargetTargetType_Group,
		"user":        CrossTenantAccessPolicyTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CrossTenantAccessPolicyTargetTargetType(input)
	return &out, nil
}
