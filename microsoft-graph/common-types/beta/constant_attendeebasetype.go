package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeBaseType string

const (
	AttendeeBaseType_Optional AttendeeBaseType = "optional"
	AttendeeBaseType_Required AttendeeBaseType = "required"
	AttendeeBaseType_Resource AttendeeBaseType = "resource"
)

func PossibleValuesForAttendeeBaseType() []string {
	return []string{
		string(AttendeeBaseType_Optional),
		string(AttendeeBaseType_Required),
		string(AttendeeBaseType_Resource),
	}
}

func (s *AttendeeBaseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttendeeBaseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttendeeBaseType(input string) (*AttendeeBaseType, error) {
	vals := map[string]AttendeeBaseType{
		"optional": AttendeeBaseType_Optional,
		"required": AttendeeBaseType_Required,
		"resource": AttendeeBaseType_Resource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeBaseType(input)
	return &out, nil
}
