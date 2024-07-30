package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusDetailsErrorCategory string

const (
	StatusDetailsErrorCategory_Failure           StatusDetailsErrorCategory = "failure"
	StatusDetailsErrorCategory_NonServiceFailure StatusDetailsErrorCategory = "nonServiceFailure"
	StatusDetailsErrorCategory_Success           StatusDetailsErrorCategory = "success"
)

func PossibleValuesForStatusDetailsErrorCategory() []string {
	return []string{
		string(StatusDetailsErrorCategory_Failure),
		string(StatusDetailsErrorCategory_NonServiceFailure),
		string(StatusDetailsErrorCategory_Success),
	}
}

func (s *StatusDetailsErrorCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusDetailsErrorCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusDetailsErrorCategory(input string) (*StatusDetailsErrorCategory, error) {
	vals := map[string]StatusDetailsErrorCategory{
		"failure":           StatusDetailsErrorCategory_Failure,
		"nonservicefailure": StatusDetailsErrorCategory_NonServiceFailure,
		"success":           StatusDetailsErrorCategory_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusDetailsErrorCategory(input)
	return &out, nil
}
