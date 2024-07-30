package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationProcessType string

const (
	PrivilegeManagementElevationProcessType_Child     PrivilegeManagementElevationProcessType = "child"
	PrivilegeManagementElevationProcessType_Parent    PrivilegeManagementElevationProcessType = "parent"
	PrivilegeManagementElevationProcessType_Undefined PrivilegeManagementElevationProcessType = "undefined"
)

func PossibleValuesForPrivilegeManagementElevationProcessType() []string {
	return []string{
		string(PrivilegeManagementElevationProcessType_Child),
		string(PrivilegeManagementElevationProcessType_Parent),
		string(PrivilegeManagementElevationProcessType_Undefined),
	}
}

func (s *PrivilegeManagementElevationProcessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegeManagementElevationProcessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegeManagementElevationProcessType(input string) (*PrivilegeManagementElevationProcessType, error) {
	vals := map[string]PrivilegeManagementElevationProcessType{
		"child":     PrivilegeManagementElevationProcessType_Child,
		"parent":    PrivilegeManagementElevationProcessType_Parent,
		"undefined": PrivilegeManagementElevationProcessType_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementElevationProcessType(input)
	return &out, nil
}
