package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantConditionSetPermissionType string

const (
	PermissionGrantConditionSetPermissionType_Application              PermissionGrantConditionSetPermissionType = "application"
	PermissionGrantConditionSetPermissionType_Delegated                PermissionGrantConditionSetPermissionType = "delegated"
	PermissionGrantConditionSetPermissionType_DelegatedUserConsentable PermissionGrantConditionSetPermissionType = "delegatedUserConsentable"
)

func PossibleValuesForPermissionGrantConditionSetPermissionType() []string {
	return []string{
		string(PermissionGrantConditionSetPermissionType_Application),
		string(PermissionGrantConditionSetPermissionType_Delegated),
		string(PermissionGrantConditionSetPermissionType_DelegatedUserConsentable),
	}
}

func (s *PermissionGrantConditionSetPermissionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionGrantConditionSetPermissionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionGrantConditionSetPermissionType(input string) (*PermissionGrantConditionSetPermissionType, error) {
	vals := map[string]PermissionGrantConditionSetPermissionType{
		"application":              PermissionGrantConditionSetPermissionType_Application,
		"delegated":                PermissionGrantConditionSetPermissionType_Delegated,
		"delegateduserconsentable": PermissionGrantConditionSetPermissionType_DelegatedUserConsentable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionGrantConditionSetPermissionType(input)
	return &out, nil
}
