package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadCompatibleAssignmentFilterAssignmentFilterManagementType string

const (
	PayloadCompatibleAssignmentFilterAssignmentFilterManagementType_Apps    PayloadCompatibleAssignmentFilterAssignmentFilterManagementType = "apps"
	PayloadCompatibleAssignmentFilterAssignmentFilterManagementType_Devices PayloadCompatibleAssignmentFilterAssignmentFilterManagementType = "devices"
)

func PossibleValuesForPayloadCompatibleAssignmentFilterAssignmentFilterManagementType() []string {
	return []string{
		string(PayloadCompatibleAssignmentFilterAssignmentFilterManagementType_Apps),
		string(PayloadCompatibleAssignmentFilterAssignmentFilterManagementType_Devices),
	}
}

func (s *PayloadCompatibleAssignmentFilterAssignmentFilterManagementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadCompatibleAssignmentFilterAssignmentFilterManagementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadCompatibleAssignmentFilterAssignmentFilterManagementType(input string) (*PayloadCompatibleAssignmentFilterAssignmentFilterManagementType, error) {
	vals := map[string]PayloadCompatibleAssignmentFilterAssignmentFilterManagementType{
		"apps":    PayloadCompatibleAssignmentFilterAssignmentFilterManagementType_Apps,
		"devices": PayloadCompatibleAssignmentFilterAssignmentFilterManagementType_Devices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadCompatibleAssignmentFilterAssignmentFilterManagementType(input)
	return &out, nil
}
