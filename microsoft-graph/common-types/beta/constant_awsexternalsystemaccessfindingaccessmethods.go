package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsExternalSystemAccessFindingAccessMethods string

const (
	AwsExternalSystemAccessFindingAccessMethods_Direct       AwsExternalSystemAccessFindingAccessMethods = "direct"
	AwsExternalSystemAccessFindingAccessMethods_RoleChaining AwsExternalSystemAccessFindingAccessMethods = "roleChaining"
)

func PossibleValuesForAwsExternalSystemAccessFindingAccessMethods() []string {
	return []string{
		string(AwsExternalSystemAccessFindingAccessMethods_Direct),
		string(AwsExternalSystemAccessFindingAccessMethods_RoleChaining),
	}
}

func (s *AwsExternalSystemAccessFindingAccessMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsExternalSystemAccessFindingAccessMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsExternalSystemAccessFindingAccessMethods(input string) (*AwsExternalSystemAccessFindingAccessMethods, error) {
	vals := map[string]AwsExternalSystemAccessFindingAccessMethods{
		"direct":       AwsExternalSystemAccessFindingAccessMethods_Direct,
		"rolechaining": AwsExternalSystemAccessFindingAccessMethods_RoleChaining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsExternalSystemAccessFindingAccessMethods(input)
	return &out, nil
}
