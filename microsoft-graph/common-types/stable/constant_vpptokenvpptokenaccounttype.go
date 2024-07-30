package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenVppTokenAccountType string

const (
	VppTokenVppTokenAccountType_Business  VppTokenVppTokenAccountType = "business"
	VppTokenVppTokenAccountType_Education VppTokenVppTokenAccountType = "education"
)

func PossibleValuesForVppTokenVppTokenAccountType() []string {
	return []string{
		string(VppTokenVppTokenAccountType_Business),
		string(VppTokenVppTokenAccountType_Education),
	}
}

func (s *VppTokenVppTokenAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenVppTokenAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenVppTokenAccountType(input string) (*VppTokenVppTokenAccountType, error) {
	vals := map[string]VppTokenVppTokenAccountType{
		"business":  VppTokenVppTokenAccountType_Business,
		"education": VppTokenVppTokenAccountType_Education,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenVppTokenAccountType(input)
	return &out, nil
}
