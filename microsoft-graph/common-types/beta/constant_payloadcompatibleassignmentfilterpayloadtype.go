package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadCompatibleAssignmentFilterPayloadType string

const (
	PayloadCompatibleAssignmentFilterPayloadType_EnrollmentRestrictions PayloadCompatibleAssignmentFilterPayloadType = "enrollmentRestrictions"
	PayloadCompatibleAssignmentFilterPayloadType_NotSet                 PayloadCompatibleAssignmentFilterPayloadType = "notSet"
)

func PossibleValuesForPayloadCompatibleAssignmentFilterPayloadType() []string {
	return []string{
		string(PayloadCompatibleAssignmentFilterPayloadType_EnrollmentRestrictions),
		string(PayloadCompatibleAssignmentFilterPayloadType_NotSet),
	}
}

func (s *PayloadCompatibleAssignmentFilterPayloadType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadCompatibleAssignmentFilterPayloadType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadCompatibleAssignmentFilterPayloadType(input string) (*PayloadCompatibleAssignmentFilterPayloadType, error) {
	vals := map[string]PayloadCompatibleAssignmentFilterPayloadType{
		"enrollmentrestrictions": PayloadCompatibleAssignmentFilterPayloadType_EnrollmentRestrictions,
		"notset":                 PayloadCompatibleAssignmentFilterPayloadType_NotSet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadCompatibleAssignmentFilterPayloadType(input)
	return &out, nil
}
