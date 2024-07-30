package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOsVppAppVppTokenAccountType string

const (
	MacOsVppAppVppTokenAccountType_Business  MacOsVppAppVppTokenAccountType = "business"
	MacOsVppAppVppTokenAccountType_Education MacOsVppAppVppTokenAccountType = "education"
)

func PossibleValuesForMacOsVppAppVppTokenAccountType() []string {
	return []string{
		string(MacOsVppAppVppTokenAccountType_Business),
		string(MacOsVppAppVppTokenAccountType_Education),
	}
}

func (s *MacOsVppAppVppTokenAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOsVppAppVppTokenAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOsVppAppVppTokenAccountType(input string) (*MacOsVppAppVppTokenAccountType, error) {
	vals := map[string]MacOsVppAppVppTokenAccountType{
		"business":  MacOsVppAppVppTokenAccountType_Business,
		"education": MacOsVppAppVppTokenAccountType_Education,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOsVppAppVppTokenAccountType(input)
	return &out, nil
}
