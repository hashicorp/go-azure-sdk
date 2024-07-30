package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludeAllAccountTargetContentType string

const (
	IncludeAllAccountTargetContentType_AddressBook IncludeAllAccountTargetContentType = "addressBook"
	IncludeAllAccountTargetContentType_IncludeAll  IncludeAllAccountTargetContentType = "includeAll"
	IncludeAllAccountTargetContentType_Unknown     IncludeAllAccountTargetContentType = "unknown"
)

func PossibleValuesForIncludeAllAccountTargetContentType() []string {
	return []string{
		string(IncludeAllAccountTargetContentType_AddressBook),
		string(IncludeAllAccountTargetContentType_IncludeAll),
		string(IncludeAllAccountTargetContentType_Unknown),
	}
}

func (s *IncludeAllAccountTargetContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncludeAllAccountTargetContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncludeAllAccountTargetContentType(input string) (*IncludeAllAccountTargetContentType, error) {
	vals := map[string]IncludeAllAccountTargetContentType{
		"addressbook": IncludeAllAccountTargetContentType_AddressBook,
		"includeall":  IncludeAllAccountTargetContentType_IncludeAll,
		"unknown":     IncludeAllAccountTargetContentType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncludeAllAccountTargetContentType(input)
	return &out, nil
}
