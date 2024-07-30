package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityResourceResourceType string

const (
	SecurityResourceResourceType_Attacked SecurityResourceResourceType = "attacked"
	SecurityResourceResourceType_Related  SecurityResourceResourceType = "related"
	SecurityResourceResourceType_Unknown  SecurityResourceResourceType = "unknown"
)

func PossibleValuesForSecurityResourceResourceType() []string {
	return []string{
		string(SecurityResourceResourceType_Attacked),
		string(SecurityResourceResourceType_Related),
		string(SecurityResourceResourceType_Unknown),
	}
}

func (s *SecurityResourceResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityResourceResourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityResourceResourceType(input string) (*SecurityResourceResourceType, error) {
	vals := map[string]SecurityResourceResourceType{
		"attacked": SecurityResourceResourceType_Attacked,
		"related":  SecurityResourceResourceType_Related,
		"unknown":  SecurityResourceResourceType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityResourceResourceType(input)
	return &out, nil
}
