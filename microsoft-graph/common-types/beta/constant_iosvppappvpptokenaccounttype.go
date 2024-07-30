package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppVppTokenAccountType string

const (
	IosVppAppVppTokenAccountType_Business  IosVppAppVppTokenAccountType = "business"
	IosVppAppVppTokenAccountType_Education IosVppAppVppTokenAccountType = "education"
)

func PossibleValuesForIosVppAppVppTokenAccountType() []string {
	return []string{
		string(IosVppAppVppTokenAccountType_Business),
		string(IosVppAppVppTokenAccountType_Education),
	}
}

func (s *IosVppAppVppTokenAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVppAppVppTokenAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVppAppVppTokenAccountType(input string) (*IosVppAppVppTokenAccountType, error) {
	vals := map[string]IosVppAppVppTokenAccountType{
		"business":  IosVppAppVppTokenAccountType_Business,
		"education": IosVppAppVppTokenAccountType_Education,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppAppVppTokenAccountType(input)
	return &out, nil
}
