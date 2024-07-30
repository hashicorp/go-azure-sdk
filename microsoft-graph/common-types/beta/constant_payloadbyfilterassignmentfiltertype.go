package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadByFilterAssignmentFilterType string

const (
	PayloadByFilterAssignmentFilterType_Exclude PayloadByFilterAssignmentFilterType = "exclude"
	PayloadByFilterAssignmentFilterType_Include PayloadByFilterAssignmentFilterType = "include"
	PayloadByFilterAssignmentFilterType_None    PayloadByFilterAssignmentFilterType = "none"
)

func PossibleValuesForPayloadByFilterAssignmentFilterType() []string {
	return []string{
		string(PayloadByFilterAssignmentFilterType_Exclude),
		string(PayloadByFilterAssignmentFilterType_Include),
		string(PayloadByFilterAssignmentFilterType_None),
	}
}

func (s *PayloadByFilterAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadByFilterAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadByFilterAssignmentFilterType(input string) (*PayloadByFilterAssignmentFilterType, error) {
	vals := map[string]PayloadByFilterAssignmentFilterType{
		"exclude": PayloadByFilterAssignmentFilterType_Exclude,
		"include": PayloadByFilterAssignmentFilterType_Include,
		"none":    PayloadByFilterAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadByFilterAssignmentFilterType(input)
	return &out, nil
}
